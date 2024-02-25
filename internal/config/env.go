package config

import (
	"errors"
	"fmt"
)

var ErrUndefinedEnvironmentType = fmt.Errorf("Undefined environment type.")

type Environment string

func (e Environment) String() string {
	return string(e)
}

const (
	Production  Environment = "prd"
	Development Environment = "dev"
	Staging     Environment = "stg"
)

func Parse(env string) (Environment, error) {
	switch env {
	case "prd", "production":
		return Production, nil
	case "stg", "staging":
		return Staging, nil
	case "dev", "development":
		return Development, nil
	}
	return "", errors.Join(ErrUndefinedEnvironmentType, fmt.Errorf(":Wrong environment: %s", env))
}
