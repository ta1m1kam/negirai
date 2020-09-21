package main

import (
	"fmt"
	"io"
	"log"

	"github.com/hajimehoshi/oto"
	"github.com/hajimehoshi/go-mp3"

	_ "github.com/TaigaMikami/negirai/statik"
	"github.com/rakyll/statik/fs"
)

func main() {
	statikFs, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}

	f, err := statikFs.Open("/yarujanaika.mp3")
	defer f.Close()
	d, err := mp3.NewDecoder(f)
	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	defer c.Close()
	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		fmt.Println(err)
	}
}
