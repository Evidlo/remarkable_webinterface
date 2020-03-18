package main

import (
	"flag"
	"os"
	"fmt"
	"github.com/evidlo/remarkable_webinterface/rm"
	"github.com/evidlo/remarkable_webinterface/boilerplate"
	"github.com/kaitai-io/kaitai_struct_go_runtime/kaitai"
)


func main() {

	input := flag.String("input", "", "input .rm")
	flag.Parse()

	f, err := os.Open(*input)
	s := kaitai.NewStream(f)
	var r rm.Rm
	err = r.Read(s, &r, &r)
	boilerplate.Check(err)

	fmt.Println("NumLayers:", r.NumLayers)
}
