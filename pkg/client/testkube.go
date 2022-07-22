package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/lreimer/testkube-watch-controller/config"
	"github.com/sirupsen/logrus"
)

type testExecutionRequest struct {
	Name string `json:"name"`
	// TestSuiteName   string            `json:"testSuiteName"`
	Namespace       string            `json:"namespace,omitempty"`
	ExecutionLabels map[string]string `json:"executionLabels,omitempty"`
}

func ExecuteTest(conf *config.Config, id string, namespace string) {
	uri := fmt.Sprintf("%s/v1/tests/%s/executions", conf.TestkubeApiServer, id)
	now := time.Now().String()
	executionName := fmt.Sprintf("%s-%s", id, now)
	payload := testExecutionRequest{Name: executionName, Namespace: namespace,
		ExecutionLabels: map[string]string{"date": now, "automatic": "true"}}
	execute(uri, namespace, payload)
}

type testSuiteExecutionRequest struct {
	Name            string            `json:"name"`
	Namespace       string            `json:"namespace,omitempty"`
	ExecutionLabels map[string]string `json:"executionLabels,omitempty"`
}

func ExecuteTestSuite(conf *config.Config, id string, namespace string) {
	uri := fmt.Sprintf("%s/v1/test-suites/%s/executions", conf.TestkubeApiServer, id)
	now := time.Now().String()
	executionName := fmt.Sprintf("%s-%s", id, now)
	payload := testSuiteExecutionRequest{Name: executionName, Namespace: namespace,
		ExecutionLabels: map[string]string{"date": now, "automatic": "true"}}
	execute(uri, namespace, payload)
}

func execute(uri string, namespace string, payload interface{}) {
	c := http.Client{Timeout: time.Second * 30}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		logrus.Errorf("Error marshalling execution request payload %s", err)
	}

	req, err := http.NewRequest("POST", uri, bytes.NewBuffer(jsonData))
	if err != nil {
		logrus.Errorf("Error creating HTTP request %s", err)
		return
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("namespace", namespace)
	req.URL.RawQuery = q.Encode()

	resp, err := c.Do(req)
	if err != nil {
		logrus.Errorf("Error during HTTP request %s", err)
		return
	}
	defer resp.Body.Close()
}
