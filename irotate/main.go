package main

import (
    "flag"
    "fmt"
    "github.com/disintegration/imaging"
    "image"
    "image/color"
    "math"
)

type RotateConfig struct {
    Input  string
    Output string
    X      int 
    Y      int 
    Size   int 
    Angle  float64
}

func parseArgs() RotateConfig {
    config := RotateConfig{}
    pi := flag.String("i", "ok", "input file name")
    po := flag.String("o", "ok_output", "output file name")
    px := flag.Int("x", 1, "center x")
    py := flag.Int("y", 1, "center y")
    ps := flag.Int("s", 1, "size")
    pa := flag.Float64("p", 1, "angle")
    flag.Parse()
    config.Input = *pi 
    config.Output = *po 
    config.X = *px 
    config.Y = *py 
    config.Size = *ps 
    config.Angle = *pa 

    return config
}

func rotate(config RotateConfig) {
    if im, err := imaging.Open(config.Input); err != nil {
        Fatal(fmt.Sprintf("can't open image %s\n", config.Input))
    } else {
        // new a 3x image
        width := im.Bounds().Max.X
        height := im.Bounds().Max.Y
        background := imaging.New(width*3, height*3, color.Black)
        lim := imaging.Paste(background, im, image.Point{width, height})
        config.X = config.X + width
               config.Y = config.Y + height

        fout := int(float64(config.Size)*math.Sqrt(2)) + 1

        sim := imaging.Crop(lim, image.Rect(config.X-fout, config.Y-fout, config.X+fout, config.Y+fout))
        rim := imaging.Rotate(sim, config.Angle, color.Black)

        dst := imaging.CropCenter(rim, config.Size*2, config.Size*2)

        if err1 := imaging.Save(dst, config.Output); err1 != nil {
            Fatal(fmt.Sprintf("can't save image %s\n", config.Output))
        }
    }
}

func main() {
    config := parseArgs()
    rotate(config)
}
