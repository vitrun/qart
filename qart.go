package qart

import (
	"fmt"
	"io/ioutil"
	"image"
	"image/png"
	"os"
	"bytes"
)

// Encode encodes a string with an image as the background
func Encode(url string, src []byte, seed int64, version, scale, mask, x, y int,
	randCtrl, dither, onlyData, saveCtrl bool) []byte {
	size, rotate := 0, 0
	if version > 8 {
		version = 8
	}
	if scale == 0 {
		scale = 8
	}
	if version >= 12 && scale >= 4 {
		scale /= 2
	}
	// TODO. support more than png
	decodedImg, _, err := image.Decode(bytes.NewBuffer(src))
	if err != nil {
		return nil
	}
	writer := new(bytes.Buffer)
	err = png.Encode(writer, decodedImg)
	if err != nil {
		return nil
	}

	target := makeTarg(writer.Bytes(), 17+4*version+size)
	img := &Image{
		Dx:           x,
		Dy:           y,
		URL:          url,
		Version:      version,
		Mask:         mask,
		RandControl:  randCtrl,
		Dither:       dither,
		OnlyDataBits: onlyData,
		SaveControl:  saveCtrl,
		Scale:        scale,
		Target:       target,
		Seed:         seed,
		Rotation:     rotate,
		Size:         size,
	}

	if err := img.Encode(); err != nil {
		fmt.Printf("error: %s\n", err)
		return nil
	}
	var dat []byte
	switch {
	case img.SaveControl:
		dat = img.Control
	default:
		dat = img.Code.PNG()
	}
	return dat
}

// EncodeByFile encodes the given url with a specific image
func EncodeByFile(url, srcImg, dstImg string, version int) {
	data, err := ioutil.ReadFile(srcImg)
	if err != nil {
		fmt.Printf("err: %s\n", err)
		return
	}
	dst := Encode(url, data, 879633355, version, 4, 2, 4, 4, false, false, false, false)
	ioutil.WriteFile(dstImg, dst, (os.FileMode)(0644))
}
