package githuboutput

const EnvName = "GITHUB_OUTPUT"

var defaultWriter = NewWriter(NewLocalEnv())

func Write(key string, value string) error {
	return WriteMap(map[string]string{key: value})
}

func WriteMap(values map[string]string) error {
	return defaultWriter.WriteMap(values)
}
