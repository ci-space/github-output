# github-output

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/ci-space/github-output) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/ci-space/github-output/master/LICENSE)

```
go get github.com/ci-space/github-output@v0.1.0
```

github-output - is simple GO library for writing values into $GITHUB_OUTPUT

## Usage

### Write single value

```go
package main

import githuboutput "github.com/ci-space/github-output"

func main() {
	githuboutput.Write("name", "Ivan")
}
```

### Write many values

```go
package main

import githuboutput "github.com/ci-space/github-output"

func main() {
	githuboutput.WriteMap(map[string]string{
		"name": "Ivan",
		"last_name": "Ivanov",
    })
}
```

### Write when output is available

```go
package main

import githuboutput "github.com/ci-space/github-output"

func main() {
	githuboutput.WhenAvailable(func() error {
        return githuboutput.Write("key", "value")
	})
}
```
