package main

import (
	"os"

	"github.com/ko6bxl/cm2obj"
)

func main() {

	var FileObj = os.Args[1]
	var FileMtl = os.Args[2]

	cm2obj.Gen(FileObj, FileMtl)
}
