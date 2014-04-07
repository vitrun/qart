/**
 * Copyright ©2014-04-07 Alex <zhirun.yu@duitang.com>
 */
package qart

import (
	"bytes"
	"image/png"
	"io/ioutil"
	"fmt"
	"os"
	"testing"
)

// ReadWrite test
func ReadWrite() {
	i := loadSize("/tmp/in.png", 48)
	var buf bytes.Buffer
	png.Encode(&buf, i)
	ioutil.WriteFile("/tmp/out.png", buf.Bytes(), (os.FileMode)(0644))
	fmt.Printf("Hello world!")
}

// Image test
func TestEncode(t *testing.T) {
	srcImg := "/tmp/in.png"
	dstImg := "/tmp/out.png"
	url := "http://www.baidu.com/"
	Encode(srcImg, dstImg, url)
}

