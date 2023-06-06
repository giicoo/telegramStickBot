package images

import (
	"fmt"
	"image/png"
	"os"

	"github.com/nfnt/resize"
)

func ResizeImage(chat_id int64, msg_id int64) error {
	path_in := fmt.Sprintf("internal/images/files/in/%v_%v.png", chat_id, msg_id)
	path_out := fmt.Sprintf("internal/images/files/out/%v_%v.png", chat_id, msg_id)

	file, err := os.Open(path_in)
	if err != nil {
		return err
	}

	outFile, err := os.Create(path_out)
	if err != nil {
		return err
	}

	defer outFile.Close()
	defer file.Close()

	src, err := png.Decode(file)
	if err != nil {
		return err
	}
	out := resize.Resize(512, 512, src, resize.Lanczos2)
	png.Encode(outFile, out)

	return nil
}
