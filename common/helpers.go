package common

import (
	"archive/zip"
	"fmt"
	"strings"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/zippkg"
)

// AddImageToZip adds an image (either from bytes or from disk) and adds it to the zip file.
func AddImageToZip(z *zip.Writer, img ImageRef, imageNum int, dt office.DocType) error {
	filename := office.AbsoluteImageFilename(dt, imageNum, strings.ToLower(img.Format()))
	if img.Data() != nil && len(*img.Data()) > 0 {
		if err := zippkg.AddFileFromBytes(z, filename, *img.Data()); err != nil {
			return err
		}
	} else if img.Path() != "" {
		if err := zippkg.AddFileFromDisk(z, filename, img.Path()); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unsupported image source: %+v", img)
	}
	return nil
}
