package presentation

import (
	"fmt"
	"os"
)

// OpenTemplate opens a template file.
func OpenTemplate(fn string) (*Presentation, error) {
	p, err := Open(fn)
	if err != nil {
		return nil, err
	}
	return p, nil
}

// Open opens and reads a document from a file (.pptx).
func Open(filename string) (*Presentation, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	defer f.Close()
	fi, err := os.Stat(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening %s: %s", filename, err)
	}
	_ = fi
	return Read(f, fi.Size())
}
