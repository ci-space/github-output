package githuboutput

import (
	"fmt"
	"os"
)

type Env interface {
	Get(key string) (string, error)
}

type LocalEnv struct{}

type MapEnv map[string]string

func NewLocalEnv() *LocalEnv {
	return &LocalEnv{}
}

func (LocalEnv) Get(key string) (string, error) {
	v, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("key %s not set", key)
	}

	return v, nil
}

func (e *MapEnv) Get(key string) (string, error) {
	v, ok := (*e)[key]
	if !ok {
		return "", fmt.Errorf("key %s not set", key)
	}
	return v, nil
}
