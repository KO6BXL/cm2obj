package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ko6bxl/cm2obj"
)

func main() {

	var FileObj = os.Args[1]
	var FileMtl = os.Args[2]

	out, err := cm2obj.Gen(FileObj, FileMtl)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out)
}
