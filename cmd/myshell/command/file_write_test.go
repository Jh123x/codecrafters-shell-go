package command

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWriteToFil(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "test_create_dir")
	assert.Nil(t, os.Mkdir(tmpDir, 0644))
	defer func() { assert.Nil(t, os.RemoveAll(tmpDir)) }()

	tests := map[string]struct {
		setUp    func(*testing.T)
		filePath string
		contents string
	}{
		"file which exists": {
			setUp: func(t *testing.T) {
				path := filepath.Join(tmpDir, "exists.txt")
				fmt.Println(path)
				assert.Nil(t, os.WriteFile(path, []byte("some other create data"), 0644))
			},
			filePath: filepath.Join(tmpDir, "exists.txt"),
			contents: "some create content",
		},
		"file which does not exists": {
			filePath: filepath.Join(tmpDir, "other.txt"),
			contents: "some create not exists content",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if tc.setUp != nil {
				tc.setUp(t)
			}

			assert.Nil(t, WriteToFile(tc.filePath, tc.contents))
			res, err := os.ReadFile(tc.filePath)
			assert.Nil(t, err)
			assert.Equal(t, tc.contents, string(res))
		})
	}
}

func TestAppendToFile(t *testing.T) {
	tmpDir := filepath.Join(os.TempDir(), "test_append_dir")
	assert.Nil(t, os.Mkdir(tmpDir, 0644))
	defer func() { assert.Nil(t, os.RemoveAll(tmpDir)) }()

	tests := map[string]struct {
		setUp    func(*testing.T)
		filePath string
		contents string

		expectedContents string
	}{
		"file which exists": {
			setUp: func(t *testing.T) {
				path := filepath.Join(tmpDir, "exists.txt")
				fmt.Println(path)
				assert.Nil(t, os.WriteFile(path, []byte("some other append data\n"), 0644))
			},
			filePath:         filepath.Join(tmpDir, "exists.txt"),
			contents:         "some append content",
			expectedContents: "some other append data\nsome append content",
		},
		"file which does not exists": {
			filePath:         filepath.Join(tmpDir, "other.txt"),
			contents:         "some append content",
			expectedContents: "some append content",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if tc.setUp != nil {
				tc.setUp(t)
			}

			assert.Nil(t, AppendToFile(tc.filePath, tc.contents))
			res, err := os.ReadFile(tc.filePath)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectedContents, string(res))
		})
	}
}
