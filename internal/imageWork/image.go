package imageWork

import (
	"os"

	"github.com/sunshineplan/imgconv"
	"github.com/yusukebe/go-pngquant"
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
	defer dstWriter.Close()

	dstImage512 := imgconv.Resize(src, &imgconv.ResizeOption{Width: 512, Height: 512})
	if err := imgconv.Write(dstWriter,
		dstImage512,
		&imgconv.FormatOption{Format: imgconv.PNG}); err != nil {
		return err
	}

	fi, err := dstWriter.Stat()
	if err != nil {
		return err
	}
	size := fi.Size()

	if size > 512_000 {
		src, err := imgconv.Open(path_in)
		if err != nil {
			return err
		}
		out, err := pngquant.Compress(src, "3")
		if err != nil {
			return err
		}
		if err := imgconv.Write(dstWriter,
			out,
			&imgconv.FormatOption{Format: imgconv.PNG}); err != nil {
			return err
		}
	}
	return nil
}
