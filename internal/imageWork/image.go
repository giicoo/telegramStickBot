package imageWork

import (
	"os"

	"github.com/sunshineplan/imgconv"
)

func ResizeImage(path_in, path_out string) error {
	src, err := imgconv.Open(path_in)
	if err != nil {
		return err
	}

	dstWriter, err := os.Create(path_out)
	if err != nil {
		return err
	}

	dstImage512 := imgconv.Resize(src, &imgconv.ResizeOption{Width: 512, Height: 512})
	if err := imgconv.Write(dstWriter,
		dstImage512,
		&imgconv.FormatOption{Format: imgconv.PNG}); err != nil {
		return err
	}
	return nil
}
