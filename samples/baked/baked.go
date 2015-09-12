package main

import (
	"fmt"
	"github.com/gonutz/fontstash.go/truetype"
	"io/ioutil"
	"path/filepath"
)

func main() {
	data, err := ioutil.ReadFile(filepath.Join("..", "ClearSans-Regular.ttf"))
	if err != nil {
		panic(err)
	}

	tmpBitmap := make([]byte, 512*512)
	cdata, err, _, tmpBitmap := truetype.BakeFontBitmap(data, 0, 32, tmpBitmap, 512, 512, 32, 96)
	var x, y float64
	b := 'b'
	x, q := truetype.GetBakedQuad(cdata, 512, 512, int(b)-32, x, y, true)

	fmt.Println(q)
}
