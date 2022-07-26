package controller

import (
	"context"

	"github.com/lreimer/testkube-watch-controller/config"
	"github.com/lreimer/testkube-watch-controller/pkg/client"
	"github.com/lreimer/testkube-watch-controller/pkg/utils"
	"github.com/sirupsen/logrus"
	apps_v1 "k8s.io/api/apps/v1"
	api_v1 "k8s.io/api/core/v1"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func Start(conf *config.Config) {
	var kubeClient kubernetes.Interface

	if _, err := rest.InClusterConfig(); err != nil {
		kubeClient = utils.GetClientOutOfCluster()
	} else {
		kubeClient = utils.GetClient()
	}

	l := map[string]string{"testkube.io/enabled": "true"}
	listOptions := meta_v1.ListOptions{
		LabelSelector: labels.SelectorFromSet(l).String(),
	}

	if conf.Resource.Deployment {
		logrus.Info("Watching for Deployment changes ...")
		watcher, err := kubeClient.AppsV1().Deployments(conf.Namespace).Watch(context.TODO(), listOptions)
		if err != nil {
			logrus.Fatalf("Unable to watch for deployment changes %s", err)
		}
		go func() {
			ch := watcher.ResultChan()
			for event := range ch {
				d, ok := event.Object.(*apps_v1.Deployment)
				if !ok {
					logrus.Errorf("Unexpected type %s", event.Object)
				}
				switch event.Type {
				case watch.Modified:
					logrus.Infof("Deployment %s modified. Processing annotations.", d.Name)
					processTestkubeAnnotations(conf, d.Annotations)
				}
			}
		}()
	}

	if conf.Resource.Services {
		logrus.Info("Watching for Service changes ...")
		watcher, err := kubeClient.CoreV1().Services(conf.Namespace).Watch(context.TODO(), listOptions)
		if err != nil {
			logrus.Fatalf("Unable to watch for service changes %s", err)
		}
		go func() {
			ch := watcher.ResultChan()
			for event := range ch {
				s, ok := event.Object.(*api_v1.Service)
				if !ok {
					logrus.Errorf("Unexpected type %s", event.Object)
				}
				switch event.Type {
				case watch.Added:
					logrus.Infof("Service %s added. Processing annotations.", s.Name)
					processTestkubeAnnotations(conf, s.Annotations)
				case watch.Modified:
					logrus.Infof("Service %s modified. Processing annotations.", s.Name)
					processTestkubeAnnotations(conf, s.Annotations)
				}
			}
		}()
	}
}

func processTestkubeAnnotations(conf *config.Config, annotations map[string]string) {
	test := annotations["testkube.io/test"]
	testSuite := annotations["testkube.io/test-suite"]
	namespace := annotations["testkube.io/namespace"]

	if len(namespace) == 0 {
		namespace = "testkube"
	}

	if len(test) != 0 {
		logrus.Infof("Executing Testkube test %s/%s", namespace, test)
		client.ExecuteTest(conf, test, namespace)
	}

	if len(testSuite) != 0 {
		logrus.Infof("Executing Testkube suite %s/%s", namespace, testSuite)
		client.ExecuteTestSuite(conf, testSuite, namespace)
	}
}
