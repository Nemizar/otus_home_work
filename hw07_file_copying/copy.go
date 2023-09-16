package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/cheggaaa/pb/v3"
)

var (
	ErrUnsupportedFile       = errors.New("unsupported file")
	ErrOffsetExceedsFileSize = errors.New("offset exceeds file size")
	ErrFileNotExists         = errors.New("file not exists")
	ErrInvalidLimit          = errors.New("limit cannot be a negative number")
	ErrInvalidOffset         = errors.New("offset cannot be a negative number")
)

func Copy(fromPath, toPath string, offset, limit int64) error {
	if limit < 0 {
		return ErrInvalidLimit
	}

	if offset < 0 {
		return ErrInvalidOffset
	}

	from, err := os.OpenFile(fromPath, os.O_RDONLY, 0o644)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrFileNotExists
		}

		return fmt.Errorf("open file %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("close file %s: %w", fromPath, err))
		}
	}(from)

	fi, err := os.Stat(from.Name())
	if err != nil {
		return fmt.Errorf("stat file %s: %w", from.Name(), err)
	}

	fSize := fi.Size()

	if fi.Size() == 0 {
		return ErrUnsupportedFile
	}

	if offset > fSize {
		return ErrOffsetExceedsFileSize
	}

	limit = getLimit(limit, fSize, offset)

	_, err = from.Seek(offset, io.SeekStart)
	if err != nil {
		return fmt.Errorf("seek %w", err)
	}

	to, err := os.Create(toPath)
	if err != nil {
		return fmt.Errorf("create file %s: %w", toPath, err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(fmt.Errorf("close file %s: %w", fromPath, err))
		}
	}(to)

	bar := pb.Full.Start64(limit)
	barReader := bar.NewProxyReader(from)
	defer bar.Finish()

	_, err = io.CopyN(to, barReader, limit)
	if err != nil {
		err := os.Remove(to.Name())
		if err != nil {
			return fmt.Errorf("delete file %w", err)
		}

		return fmt.Errorf("copy %w", err)
	}

	return nil
}

func getLimit(limit int64, fSize int64, offset int64) int64 {
	if limit == 0 || limit > fSize {
		limit = fSize
	} else if fSize < offset+limit {
		limit = fSize - offset
	}

	return limit
}
