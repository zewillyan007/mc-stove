package util

import (
	"archive/tar"
	"bytes"
	"compress/bzip2"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func ExtractBz2(body []byte, location string) error {
	bodyCopy := make([]byte, len(body))
	copy(bodyCopy, body)
	tarFile := bzip2.NewReader(bytes.NewReader(body))
	tarReader := tar.NewReader(tarFile)

	var dirList []string

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		dirList = append(dirList, header.Name)
	}

	basedir := findBaseDir(dirList)

	tarFile = bzip2.NewReader(bytes.NewReader(bodyCopy))
	tarReader = tar.NewReader(tarFile)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(location, strings.Replace(header.Name, basedir, "", -1))
		info := header.FileInfo()

		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
			}
			continue
		}

		if header.Typeflag == tar.TypeSymlink {
			err = os.Symlink(header.Linkname, path)
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err == nil {
			defer file.Close()
		}
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return err
		}
	}
	return nil
}

func findBaseDir(dirList []string) string {
	if len(dirList) == 1 {
		return path.Dir(dirList[0]) + "/"
	}

	dontdiff := []string{"pax_global_header"}
	for _, v := range dontdiff {
		dirList = removeStringFromSlice(dirList, v)
	}

	commonBaseDir := commonPrefix('/', dirList)
	if commonBaseDir != "" {
		commonBaseDir = commonBaseDir + "/"
	}
	return commonBaseDir
}

func removeStringFromSlice(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func commonPrefix(sep byte, paths []string) string {

	switch len(paths) {
	case 0:
		return ""
	case 1:
		return path.Clean(paths[0])
	}

	c := []byte(path.Clean(paths[0]))

	c = append(c, sep)

	for _, v := range paths[1:] {

		v = path.Clean(v) + string(sep)

		if len(v) < len(c) {
			c = c[:len(v)]
		}
		for i := 0; i < len(c); i++ {
			if v[i] != c[i] {
				c = c[:i]
				break
			}
		}
	}

	for i := len(c) - 1; i >= 0; i-- {
		if c[i] == sep {
			c = c[:i]
			break
		}
	}

	return string(c)
}
