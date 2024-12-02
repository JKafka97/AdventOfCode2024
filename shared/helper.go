package helper

import (
	"os"
	"strings"
	"testing"
)

func ReadFile(t *testing.T, filePath string) string {
	t.Helper()
	content, err := os.ReadFile(filePath)
	if err != nil {
		t.Fatalf("Error reading file %s: %v", filePath, err)
	}
	return strings.ReplaceAll(string(content), "\r\n", "\n")
}
