package utils

import (
	"fmt"

	"github.com/paypal/dce-go/types"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type ConditionFunc func() (string, error)

func PluginPanicHandler(condition ConditionFunc) (res string, err error) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recover : %v \n", r)
			err = errors.New(fmt.Sprintf("Recover : %v \n", r))
		}
	}()

	if res, err = condition(); err != nil {
		log.Errorf("Error executing plugins: %v \n", err)
		return res, err
	}
	return res, err
}

func ToPodStatus(s string) types.PodStatus {
	switch s {
	case "POD_STAGING":
		return types.POD_STAGING
	case "POD_STARTING":
		return types.POD_STARTING
	case "POD_RUNNING":
		return types.POD_RUNNING
	case "POD_FAILED":
		return types.POD_FAILED
	case "POD_KILLED":
		return types.POD_KILLED
	case "POD_FINISHED":
		return types.POD_FINISHED
	case "POD_PULL_FAILED":
		return types.POD_PULL_FAILED
	}

	return types.POD_EMPTY
}

func ToHealthStatus(s string) types.HealthStatus {
	switch s {
	case "starting":
		return types.STARTING
	case "healthy":
		return types.HEALTHY
	case "unhealthy":
		return types.UNHEALTHY
	}

	return types.UNKNOWN_HEALTH_STATUS
}
