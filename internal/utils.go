package internal

import (
	"io/fs"
	"os"
)

func FileExists(path string) (bool, fs.FileInfo, error) {
    stat, err := os.Stat(path)
    if err == nil { return true, stat, nil }
    if os.IsNotExist(err) { return false, stat, nil }
    return false, nil, err
}
