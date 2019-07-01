package fzip

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func UnZipFile(f *os.File) (*zip.Reader, error) {
	fi, err := f.Stat()
	if err != nil {
		return nil, err
	}
	return zip.NewReader(f, fi.Size())
}

func UnZipPath(zipFile string, deleteZip ...bool) error {
	f, err := os.Open(zipFile)
	if err != nil {
		return err
	}
	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return err
	}
	r, err := zip.NewReader(f, fi.Size())
	if err != nil {
		return err
	}
	if err := readZip(r); err != nil {
		return err
	}
	if len(deleteZip) > 0 && deleteZip[0] {
		if err := os.Remove(zipFile); err != nil {
			return err
		}
	}
	return nil
}

func readZip(r *zip.Reader) error {
	if r == nil {
		return fmt.Errorf("не иницилизированный Reader")
	}
	for _, f := range r.File {
		if err := os.MkdirAll(filepath.Dir(f.Name), 0766); err != nil {
			return err
		}
		if f.FileInfo().IsDir() {
			continue
		}
		nf, err := os.OpenFile(f.Name, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer nf.Close()
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()
		_, err = io.Copy(nf, rc)
		if err != nil {
			return err
		}
	}
	return nil
}
