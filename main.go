package main

import (
	"fmt"

	"github.com/xupin/bitmap/bitmap"
)

func main() {
	bm := bitmap.NewBitMap(8)
	fmt.Printf("%08b\n", bm.GetBits())
	bm.Add(8)
	fmt.Printf("%08b\n", bm.GetBits())
	fmt.Println(bm.Lookup(8))
	bm.Del(8)
	fmt.Printf("%08b\n", bm.GetBits())
}
