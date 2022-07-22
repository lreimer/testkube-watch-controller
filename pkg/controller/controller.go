package controller

import (
	"github.com/lreimer/testkube-watch-controller/config"
	"github.com/sirupsen/logrus"
)

func Start(conf *config.Config) {
	if conf.Resource.Deployment {
		logrus.Info("Watching for Deployment changes")
	}
	if conf.Resource.Services {
		logrus.Info("Watching for Service changes")
	}
}
