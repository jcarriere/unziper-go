package unziper

import (
	"archive/zip"
	"os"
	"path/filepath"
	"io"
	"fmt"
)

type Unziper struct {
	Src string
	Dest string
}

func New(src string, dest string) Unziper {
	return Unziper{src, dest}
}

func (uz Unziper) Extract() error {

	fmt.Println("Extraction of " + uz.Src + " started!")

	r, err := zip.OpenReader(uz.Src)
	if err != nil {
		return err
	}
	defer func() {
		if err := r.Close(); err != nil {
			panic(err)
		}
	}()

	os.MkdirAll(uz.Dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer func() {
			if err := rc.Close(); err != nil {
				panic(err)
			}
		}()

		path := filepath.Join(uz.Dest, f.Name)

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer func() {
				if err := f.Close(); err != nil {
					panic(err)
				}
			}()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
		fmt.Println("Extracting file: " + f.Name)
	}

	fmt.Println("Extraction of " + uz.Src + " finished!")

	return nil
}