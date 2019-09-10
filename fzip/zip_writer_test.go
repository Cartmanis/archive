package fzip

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestZip(t *testing.T) {
	fileZip, err := os.Create("test.zip")
	if err != nil {
		t.Fatal(err)
	}
	defer fileZip.Close()

	var r1 bytes.Buffer
	var r2 bytes.Buffer

	_, err = r1.Write([]byte("Данные первого файла"))
	if err != nil {
		t.Fatal(err)
	}
	_, err = r2.Write([]byte("Reader второго файла"))
	if err != nil {
		t.Fatal(err)
	}
	listNameReader := make([]*NameReader, 0)
	listNameReader = append(listNameReader, &NameReader{"1.txt", &r1})
	listNameReader = append(listNameReader, &NameReader{"2.txt", &r2})

	if err := Zip(listNameReader, fileZip); err != nil {
		t.Fatal(err)
	}
	fmt.Println(fileZip)
}
