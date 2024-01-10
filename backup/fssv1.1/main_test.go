package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestRun(t *testing.T) {
	testCases := []struct {
		name     string
		root     string
		cfg      config
		expected string
	}{
		{
			name: "NoFilter",
			root: "testdata",
			cfg: config{
				ext:  "",
				size: 0,
				list: true,
			},
			expected: "testdata/dir.log\ntestdata/dir2/script.sh\ntestdata/log.gz\n",
		},
		{
			name: "FilterExtensionMatch",
			root: "testdata",
			cfg: config{
				ext:  ".log",
				size: 0,
				list: true,
			},
			expected: "testdata/dir.log\n",
		},
		{
			name: "FilterExtensionSizeMatch",
			root: "testdata",
			cfg: config{
				ext:  ".log",
				size: 10,
				list: true,
			},
			expected: "testdata/dir.log\n",
		},
		{
			name: "FilterExtensionSizeNoMatch",
			root: "testdata",
			cfg: config{
				ext:  ".log",
				size: 20,
				list: true,
			},
			expected: "",
		},
		{
			name: "FilterExtensionNoMatch",
			root: "testdata",
			cfg: config{
				ext:  ".gz",
				size: 0,
				list: true,
			},
			expected: "testdata/log.gz\n",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer
			if err := run(tc.root, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}
			res := buffer.String()

			if tc.expected != res {
				t.Errorf("expected %q, got %q instead\n", tc.expected, res)
			}
		})
	}
}

// TestRunDelExtension
func TestRunDelExtension(t *testing.T) {
	testCases := []struct {
		name        string
		cfg         config
		extNoDelete string
		nDelete     int
		nNodelete   int
		expected    string
	}{
		{
			name: "DeleteExtensionNoMatch",
			cfg: config{
				ext: ".log",
				del: true,
			},
			extNoDelete: ".gz",
			nDelete:     0,
			nNodelete:   10,
			expected:    "",
		},
		{
			name: "DeleteExtensionMatch",
			cfg: config{
				ext: ".log",
				del: true,
			},
			extNoDelete: "",
			nDelete:     10,
			nNodelete:   0,
			expected:    "",
		},
		{
			name: "DeleteExtensionMixed",
			cfg: config{
				ext: ".log",
				del: true,
			},
			extNoDelete: ".gz",
			nDelete:     5,
			nNodelete:   5,
			expected:    "",
		},
	}

	// Execute RunDel Test test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var buffer bytes.Buffer

			tempDir, cleanup := createTempDir(t, map[string]int{
				tc.cfg.ext:     tc.nDelete,
				tc.extNoDelete: tc.nNodelete,
			})
			defer cleanup()
			if err := run(tempDir, &buffer, tc.cfg); err != nil {
				t.Fatal(err)
			}
			res := buffer.String()
			if tc.expected != res {
				t.Errorf("expected %q, go %q instead\n", tc.expected, res)
			}
			filesLeft, err := ioutil.ReadDir(tempDir)
			if err != nil {
				t.Error(err)
			}

			if len(filesLeft) != tc.nNodelete {
				t.Errorf("expected %d files left, got %d instead\n", tc.nNodelete, len(filesLeft))
			}
		})
	}
}

//createTestDir
func createTempDir(t *testing.T, files map[string]int) (dirname string, cleanup func()) {
	t.Helper()

	tempDir, err := ioutil.TempDir("/tmp", "walktest")
	// fmt.Println(tempDir)
	if err != nil {
		t.Fatal(err)
	}

	for i, n := range files {
		for j := 1; j <= n; j++ {
			fname := fmt.Sprintf("file%d%s", j, i)
			fpath := filepath.Join(tempDir, fname)
			if err := ioutil.WriteFile(fpath, []byte("dummy"), 0644); err != nil {
				t.Fatal(err)
			}
		}
	}
	return tempDir, func() { os.RemoveAll(tempDir) }
}
