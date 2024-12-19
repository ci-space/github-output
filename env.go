package githuboutput

import (
	"errors"
	"os"
)

type Env interface {
	// Get can return ErrEnvVarNotFound
	Get(key string) (string, error)
}

type LocalEnv struct{}

type MapEnv map[string]string

var ErrEnvVarNotFound = errors.New("environment variable not found")

func NewLocalEnv() *LocalEnv {
	return &LocalEnv{}
}

func (LocalEnv) Get(key string) (string, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return "", ErrEnvVarNotFound
	}

	return v, nil
}

func (e *MapEnv) Get(key string) (string, error) {
	v, ok := (*e)[key]
	if !ok {
		return "", ErrEnvVarNotFound
	}
	return v, nil
}
