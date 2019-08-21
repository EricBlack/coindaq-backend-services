package utils

import (
	"image"
	"bytes"
	"github.com/satori/go.uuid"
	"os"
	"image/png"
)

func GenerateImage(filePath string, imageByte []byte) (string, error) {
	image, format, err := image.Decode(bytes.NewReader(imageByte))
	if err != nil {
		return "", err
	}else{
		uuidInfo, _ := uuid.NewV4()
		fileName := uuidInfo.String() + "." +format
		file, err := os.Create(filePath + "/" + fileName)
		if err != nil {
			return "", err
		}
		defer file.Close()

		err = png.Encode(file, image)
		return fileName, err
	}
}
