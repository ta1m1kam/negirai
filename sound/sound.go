package sound

import (
	_ "github.com/TaigaMikami/negirai/statik"
	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto"
	"github.com/rakyll/statik/fs"
	"io"
)

func Play() error {
	statikFs, err := fs.New()
	if err != nil {
		return err
	}

	f, err := statikFs.Open("/yarujanaika.mp3")
	if err != nil {
		return err
	}
	defer f.Close()

	d, err := mp3.NewDecoder(f)
	if err != nil {
		return err
	}

	c, err := oto.NewContext(d.SampleRate(), 2, 2, 8192)
	if err != nil {
		return err
	}
	defer c.Close()

	p := c.NewPlayer()
	defer p.Close()

	if _, err := io.Copy(p, d); err != nil {
		return err
	}

	return nil
}
