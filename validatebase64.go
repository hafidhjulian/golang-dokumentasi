package main

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"log"
	"os"
	"strings"

	"github.com/google/uuid"
)

func Validatefile(index, file string) (string, int, error) {
	uid := uuid.New().String()
	decodefile := base64.NewDecoder(base64.StdEncoding, strings.NewReader(file))
	buff := bytes.Buffer{}
	_, errbuff := buff.ReadFrom(decodefile)
	if errbuff != nil {
		log.Fatal(errbuff)
		return errbuff.Error(), 1, errbuff
	}
	filenamefile := index + "-" + uid + ".pdf"
	os.WriteFile(filenamefile, buff.Bytes(), 0777)
	return filenamefile, 0, nil
}

func Validatefotoblob(foto string) (string, int, error) {
	// uid := uuid.New().String()
	decodefoto := base64.NewDecoder(base64.StdEncoding, strings.NewReader(foto))
	buff := bytes.Buffer{}
	_, errbuff := buff.ReadFrom(decodefoto)
	if errbuff != nil {
		log.Fatal(errbuff)
		return errbuff.Error(), 1, errbuff
	}
	imgCfg, fm, errdecode := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if errdecode != nil {
		log.Fatal(errdecode)
		return errdecode.Error(), 1, errdecode
	}
	if imgCfg.Width > imgCfg.Height && imgCfg.Height > 500 {
		log.Fatal("Invalid Images!")
		return "Resolusi Gambar tidak sesuai", 1, errors.New("Invalid Images!")
	}
	if fm != "png" && fm != "jpg" && fm != "jpeg" {
		log.Fatal("Invalid Images!")
		return "Format Gambar tidak sesuai", 1, errors.New("Invalid Images!")
	}
	//"data:{$files["pas_photo"]["type"]};base64,"
	filefoto := "data:" + fm + ";base64," + foto
	return filefoto, 0, nil
}
