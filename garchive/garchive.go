package garchive

import (
	"archive/zip"
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"time"
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

//Zip 压缩文件
func Zip(zipfile string, md5 string, bs []byte) error {

	tempfile := randname(md5)
	// 创建 zip 包文件
	fw, err := os.Create(zipfile)
	if err != nil {
		return err
	}

	// 实例化新的 zip.Writer
	zw := zip.NewWriter(fw)

rename:
	f, err := os.Create(tempfile)
	if err != nil {
		tempfile = randname(md5)
		fmt.Println("create file err:", err, tempfile)
		goto rename
	}

	fi, err := f.Stat()
	if err != nil {
		tempfile = randname(md5)
		fmt.Println("Stat file err:", err, tempfile)
		goto rename
	}

	fh, err := zip.FileInfoHeader(fi)
	if err != nil {
		return err
	}
	w, err := zw.CreateHeader(fh)
	if err != nil {
		return err
	}
	buf := bytes.NewBuffer(bs)
	// 写入文件内容
	_, err = io.Copy(w, buf)
	if err != nil {
		return err
	}

	defer fw.Close()
	defer zw.Close()
	defer func() {
		f.Close()
		os.Remove(tempfile)
	}()

	return nil
}

//UnZip 解压文件
func UnZip(path string) ([]byte, string, error) {
	rc, err := zip.OpenReader(path)
	defer rc.Close()
	bs := make([]byte, 0)
	if err != nil {
		return bs, "", err
	}
	for _, file := range rc.File {
		f, err := file.Open()
		if err != nil {
			return bs, "", err
		}
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return bs, "", err
		}

		bs = append(bs, b...)
	}
	md5, err := GetFileMD5(bs)
	return bs, md5, err
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

func randname(md5 string) string {
	rand.Seed(time.Now().UnixNano())
	ri := rand.Intn(1000000)
	rs := strconv.Itoa(ri)
	tempfile := md5 + "-" + rs + ".txt"
	return tempfile
}
