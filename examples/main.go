package main

import (
	"fmt"
	"github.com/jcarriere/unziper-go"
)

func main() {
	uz := unziper.New("file.zip", "directory/")
	err := uz.Extract()
	if err != nil {
		fmt.Println(err)
	}
}