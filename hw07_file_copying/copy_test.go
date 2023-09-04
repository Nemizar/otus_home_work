package main

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCopy(t *testing.T) {
	testCases := []struct {
		name     string
		fromPath string
		toPath   string
		limit    int64
		offset   int64
		error    error
	}{
		{
			name:     "simple case",
			fromPath: "testdata/input.txt",
			toPath:   "out.txt",
			limit:    0,
			offset:   0,
			error:    nil,
		},
		{
			name:     "offset is over file size",
			fromPath: "testdata/input.txt",
			toPath:   "out.txt",
			limit:    0,
			offset:   7000,
			error:    ErrOffsetExceedsFileSize,
		},
		{
			name:     "unsupported file",
			fromPath: "/dev/urandom",
			toPath:   "out.txt",
			limit:    0,
			offset:   0,
			error:    ErrUnsupportedFile,
		},
		{
			name:     "open file error",
			fromPath: "testdata/input",
			toPath:   "out.txt",
			limit:    0,
			offset:   0,
			error:    errors.New("file not exists"),
		},
		{
			name:     "negative offset",
			fromPath: "testdata/input.txt",
			toPath:   "out.txt",
			limit:    0,
			offset:   -2,
			error:    ErrInvalidOffset,
		},
		{
			name:     "negative limit",
			fromPath: "testdata/input.txt",
			toPath:   "out.txt",
			limit:    -2,
			offset:   0,
			error:    ErrInvalidLimit,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := Copy(tc.fromPath, tc.toPath, tc.offset, tc.limit)
			if tc.error == nil {
				require.NoError(t, err)

				return
			}
			require.EqualError(t, err, tc.error.Error())

			errRm := os.Remove(tc.toPath)
			_ = errRm
		})
	}
}
