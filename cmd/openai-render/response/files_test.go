package response

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_FilesFromChat(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		lang     string
		want     []string
	}{
		{
			name:     "extract go files from code blocks like In `hello/ping.go`: ```...",
			filePath: "fixtures/sample.output",
			lang:     "go",
			want: []string{
				"hello/ping.go",
				"cmd/main.go",
			},
		},
		{
			name:     "extract go files from code blocks like ping.go: ``` // hello/ping.go...",
			filePath: "fixtures/sample2.output",
			lang:     "go",
			want: []string{
				"hello/hello_service.go",
				"cmd/main.go",
			},
		},
		{
			name:     "extract go files from code blocks like ``` // hello/ping.go...",
			filePath: "fixtures/sample3.output",
			lang:     "go",
			want: []string{
				"cmd/main.go",
				"src/ping.go",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fixture := readFileAsString(t, test.filePath)
			got := FilesFromChat(fixture, test.lang)
			assert.ElementsMatch(t, test.want, got.Filenames())
		})
	}
}

func readFileAsString(t *testing.T, filePath string) string {
	bytes, err := os.ReadFile(filePath)
	require.NoError(t, err)
	return string(bytes)
}
