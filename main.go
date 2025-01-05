package cm2obj

import (
	"errors"
	"os"
	"path"
	"strconv"

	"github.com/g3n/engine/loader/obj"
	"github.com/nameless9000/cm2go/block"
	"github.com/nameless9000/cm2go/build"
)

func Gen(FileObj, FileMtl string) (string, error) {
	//var Mode = os.Args[1]

	objReader, err := os.Open(FileObj)

	if err != nil {
		return "", err
	}
	defer objReader.Close()

	mtlReader, err := os.Open(FileMtl)

	if err != nil {
		return "", err
	}
	defer mtlReader.Close()

	if path.Ext(FileMtl) != ".mtl" {
		return "", errors.New("ERROR: Please use a mtl file with the extention '.mtl'")
	}

	obj, err := getObj(objReader, mtlReader)

	if err != nil {
		return "", err
	}
	var collection block.Collection
	var thingieidk *block.Base
	var color2 block.Color
	color2.R = 255
	color2.G = 255
	color2.B = 255

	var count float32 = 1
	for vertCount, _ := range obj.Vertices {
		if (float32(vertCount)+1)/3 == count {
			thingieidk = collection.Append(block.TILE(color2, 2))
			thingieidk.Offset.X = obj.Vertices[vertCount]
			thingieidk.Offset.Y = obj.Vertices[vertCount-1] + 1
			thingieidk.Offset.Z = obj.Vertices[vertCount-2]
			count++
		}
	}
	out, err := build.Compile([]block.Collection{collection})

	if err != nil {
		return "", err
	}
	return out, nil
}

func getObj(objReader, mtlReader *os.File) (*obj.Decoder, error) {
	obj, err := obj.DecodeReader(objReader, mtlReader)

	if err != nil {
		return nil, err
	}

	return obj, nil
}

func f32toStr(float float32) string {
	return strconv.FormatFloat(float64(float), 'g', -1, 32)
}
