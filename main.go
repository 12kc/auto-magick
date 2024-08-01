package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Filetype converted from
const PREV_FILETYPE = ".png"

// Filetype to convert to
const CONV_FILETYPE = ".jpeg"

// Reduce quality (read
// https://imagemagick.org/script/command-line-options.php#quality
// to understand)
const QUALITY = "85%"

func main() {
	cf, err := os.ReadDir(".")

	if err != nil {
		panic(err)
	}

	for _, f := range cf {
		if strings.HasSuffix(f.Name(), PREV_FILETYPE) {
			pic := f.Name()

			fmt.Printf("Found %s\n", pic)
			fmt.Printf("Converting %s to %s, hold on... \n", pic, CONV_FILETYPE)

			if err := convertImage(pic); err != nil {
				panic(err)
			}

			fmt.Println("Converted successfully!")
		}
	}

}

func convertImage(picture string) error {
	bp := picture
	ap := fmt.Sprintf("%s_CONVERTED%s", bp, CONV_FILETYPE)
	convert := exec.Command("magick", bp, "-quality", QUALITY, ap)

	if err := convert.Run(); err != nil {
		return err
	} else {
		return nil
	}
}
