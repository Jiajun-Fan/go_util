package main

import (
	"flag"
	"fmt"
	"github.com/disintegration/imaging"
)

type ResizeConfig struct {
	Input  string
	Output string
	Width  int
	Height int
}

func parseArgs() ResizeConfig {
	config := ResizeConfig{}
	pi := flag.String("i", "ok", "input file name")
	po := flag.String("o", "ok_output", "output file name")
	pw := flag.Int("w", 256, "width")
	ph := flag.Int("h", 256, "height")
	flag.Parse()
	config.Input = *pi
	config.Output = *po
	config.Width = *pw
	config.Height = *ph

	return config
}

func resize(config ResizeConfig) {
	if im, err := imaging.Open(config.Input); err != nil {
		Fatal(fmt.Sprintf("can't open image %s\n", config.Input))
	} else {
		dst := imaging.Resize(im, config.Width, config.Height, imaging.Box)
		if err1 := imaging.Save(dst, config.Output); err1 != nil {
			Fatal(fmt.Sprintf("can't save image %s\n", config.Output))
		}
	}
}

func main() {
	config := parseArgs()
	resize(config)
}
