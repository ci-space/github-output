package githuboutput_test

import (
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	githuboutput "github.com/ci-space/github-output"
)

func TestWriter_WriteMap(t *testing.T) {
	t.Run("error: env not found", func(t *testing.T) {
		writer := githuboutput.NewWriter(&githuboutput.MapEnv{})

		err := writer.WriteMap(map[string]string{
			"key": "value",
		})
		require.Error(t, err)
		require.Equal(t, "failed to get output filename: key GITHUB_OUTPUT not set", err.Error())
	})

	t.Run("success", func(t *testing.T) {
		writer := githuboutput.NewWriter(&githuboutput.MapEnv{
			githuboutput.EnvName: "./test.txt",
		})

		err := writer.WriteMap(map[string]string{
			"key": "value",
		})
		require.NoError(t, err)
		require.FileExists(t, "./test.txt")
		defer os.Remove("./test.txt")
		file, err := os.Open("./test.txt")
		require.NoError(t, err)
		defer file.Close()
		content, err := io.ReadAll(file)
		require.NoError(t, err)
		require.Equal(t, "key=value\n", string(content))
	})
}
