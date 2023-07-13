package main

import (
	"fmt"
	"os"
	"sample_kaitai/gifmodule"

	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)

func main() {
	file, err := os.Open("giphy.gif")
	fmt.Println("err r ", err)
	g := gifmodule.NewGifmodule()
	err = g.Read(kaitai.NewStream(file), nil, g)

	fmt.Printf("width = %d\n", g.LogicalScreen.ImageWidth)
	fmt.Printf("height = %d\n", g.LogicalScreen.ImageHeight)
}
