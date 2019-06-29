package fzip

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

func UnZipPath(zipFile string, deleteZip ...bool) error {
	r, err := zip.OpenReader(zipFile)
	if err != nil {
		return err
	}
	defer r.Close()

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
	//удаление архива
	if len(deleteZip) > 0 && deleteZip[0] {
		if err := os.Remove(zipFile); err != nil {
			return err
		}
	}
	return nil
}
