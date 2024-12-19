package githuboutput

import "errors"

const EnvName = "GITHUB_OUTPUT"

var (
	defaultEnv    = NewLocalEnv()
	defaultWriter = NewWriter(defaultEnv)
)

func WhenAvailable(write func() error) error {
	if _, err := defaultEnv.Get(EnvName); err != nil {
		if errors.Is(err, ErrEnvVarNotFound) {
			return nil
		}
		return err
	}

	return write()
}

func Write(key string, value string) error {
	return WriteMap(map[string]string{key: value})
}

func WriteMap(values map[string]string) error {
	return defaultWriter.WriteMap(values)
}
