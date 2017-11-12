package main

import (
	"flag"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
)

type CropConfig struct {
	Input  string
	Output string
	Xmin   int
	Ymin   int
	Xmax   int
	Ymax   int
}

func parseArgs() CropConfig {
	config := CropConfig{}
	pi := flag.String("i", "ok", "input file name")
	po := flag.String("o", "ok_output", "output file name")
	pa := flag.Int("a", 0, "X min")
	pb := flag.Int("b", 0, "Y min")
	pc := flag.Int("c", 1, "X max")
	pd := flag.Int("d", 1, "Y max")
	flag.Parse()
	config.Input = *pi
	config.Output = *po
	config.Xmin = *pa
	config.Ymin = *pb
	config.Xmax = *pc
	config.Ymax = *pd

	return config
}

func crop(config CropConfig) {
	if config.Xmin >= config.Xmax || config.Ymin >= config.Ymax {
		Fatal(fmt.Sprintf("bad box %d %d %d %d\n", config.Xmin, config.Ymin, config.Xmax, config.Ymax))
	}
	if im, err := imaging.Open(config.Input); err != nil {
		Fatal(fmt.Sprintf("can't open image %s\n", config.Input))
	} else {
		dst := imaging.Crop(im, image.Rect(config.Xmin, config.Ymin, config.Xmax, config.Ymax))
		if err1 := imaging.Save(dst, config.Output); err1 != nil {
			Fatal(fmt.Sprintf("can't save image %s\n", config.Output))
		}
	}
}

func main() {
	config := parseArgs()
	crop(config)
}
