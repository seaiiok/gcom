package gfiles

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
)

type fc func(string) (map[string]struct{}, error)

func scanFiles() (f fc) {
	files := make(map[string]struct{}, 0)
	f = func(root string) (map[string]struct{}, error) {
		rd, err := ioutil.ReadDir(root)
		for _, fi := range rd {
			newRoot := root + "/" + fi.Name()
			if fi.IsDir() {
				f(newRoot)
			} else {
				files[newRoot] = struct{}{}
			}
		}

		return files, err
	}
	return f
}

//ScanFiles 获取目录下所有文件
func ScanFiles(root string) (map[string]struct{}, error) {
	fc := scanFiles()
	return fc(root)
}

// GetFileMD5 计算文件MD5值
func GetFileMD5(b []byte) (string, error) {
	buf := bytes.NewBuffer(b)
	md5hash := md5.New()
	if _, err := io.Copy(md5hash, buf); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5hash.Sum(nil)), nil
}
