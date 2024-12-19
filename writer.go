package githuboutput

import (
	"fmt"
	"log/slog"
	"os"
)

type Writer struct {
	env Env
}

func NewWriter(env Env) *Writer {
	return &Writer{env: env}
}

func (w *Writer) WriteMap(values map[string]string) error {
	output, err := w.env.Get(EnvName)
	if err != nil {
		return fmt.Errorf("failed to get output filename from env variable %s: %w", EnvName, err)
	}

	outputFile, err := os.OpenFile(output, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("failed to open output file: %w", err)
	}

	defer func(outputFile *os.File) {
		ferr := outputFile.Close()
		if ferr != nil {
			slog.With(slog.Any("err", ferr)).Error("failed to close output file")
		}
	}(outputFile)

	for k, v := range values {
		_, err = outputFile.WriteString(fmt.Sprintf("%s=%s\n", k, v))
		if err != nil {
			return fmt.Errorf("failed to write key %q to output file: %w", k, err)
		}
	}

	return nil
}
