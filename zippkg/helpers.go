package zippkg

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"go.devnw.com/ooxml"

	"go.devnw.com/ooxml/algo"
	"go.devnw.com/ooxml/schema/soo/pkg/relationships"
)

// RelationsPathFor returns the relations path for a given filename.
func RelationsPathFor(path string) string {
	sp := strings.Split(path, "/")
	pathPortion := strings.Join(sp[0:len(sp)-1], "/")
	filePortion := sp[len(sp)-1]
	pathPortion += "/_rels/"
	filePortion += ".rels"
	return pathPortion + filePortion
}

// Decode unmarshals the content of a *zip.File as XML to a given destination.
func Decode(f *zip.File, dest interface{}) error {
	rc, err := f.Open()
	if err != nil {
		return fmt.Errorf("error reading %s: %s", f.Name, err)
	}
	defer rc.Close()
	dec := xml.NewDecoder(rc)
	if err := dec.Decode(dest); err != nil {
		return fmt.Errorf("error decoding %s: %s", f.Name, err)
	}

	// this ensures that relationship ID is increasing, which we apparently rely
	// on....
	if ds, ok := dest.(*relationships.Relationships); ok {
		for i, r := range ds.Relationship {
			switch r.TypeAttr {
			// Common
			case office.OfficeDocumentTypeStrict:
				ds.Relationship[i].TypeAttr = office.OfficeDocumentType
			case office.StylesTypeStrict:
				ds.Relationship[i].TypeAttr = office.StylesType
			case office.ThemeTypeStrict:
				ds.Relationship[i].TypeAttr = office.ThemeType
			case office.SettingsTypeStrict:
				ds.Relationship[i].TypeAttr = office.SettingsType
			case office.ImageTypeStrict:
				ds.Relationship[i].TypeAttr = office.ImageType
			case office.CommentsTypeStrict:
				ds.Relationship[i].TypeAttr = office.CommentsType
			case office.ThumbnailTypeStrict:
				ds.Relationship[i].TypeAttr = office.ThumbnailType
			case office.DrawingTypeStrict:
				ds.Relationship[i].TypeAttr = office.DrawingType
			case office.ChartTypeStrict:
				ds.Relationship[i].TypeAttr = office.ChartType
			case office.ExtendedPropertiesTypeStrict:
				ds.Relationship[i].TypeAttr = office.ExtendedPropertiesType
			case office.CustomXMLTypeStrict:
				ds.Relationship[i].TypeAttr = office.CustomXMLType

			// SML
			case office.WorksheetTypeStrict:
				ds.Relationship[i].TypeAttr = office.WorksheetType
			case office.SharedStringsTypeStrict:
				ds.Relationship[i].TypeAttr = office.SharedStringsType
			case office.TableTypeStrict:
				ds.Relationship[i].TypeAttr = office.TableType

			// WML
			case office.HeaderTypeStrict:
				ds.Relationship[i].TypeAttr = office.HeaderType
			case office.FooterTypeStrict:
				ds.Relationship[i].TypeAttr = office.FooterType
			case office.NumberingTypeStrict:
				ds.Relationship[i].TypeAttr = office.NumberingType
			case office.FontTableTypeStrict:
				ds.Relationship[i].TypeAttr = office.FontTableType
			case office.WebSettingsTypeStrict:
				ds.Relationship[i].TypeAttr = office.WebSettingsType
			case office.FootNotesTypeStrict:
				ds.Relationship[i].TypeAttr = office.FootNotesType
			case office.EndNotesTypeStrict:
				ds.Relationship[i].TypeAttr = office.EndNotesType

			// PML
			case office.SlideTypeStrict:
				ds.Relationship[i].TypeAttr = office.SlideType

			// VML
			case office.VMLDrawingTypeStrict:
				ds.Relationship[i].TypeAttr = office.VMLDrawingType
			}
		}

		sort.Slice(ds.Relationship, func(i, j int) bool {
			lhs := ds.Relationship[i]
			rhs := ds.Relationship[j]
			return algo.NaturalLess(lhs.IdAttr, rhs.IdAttr)
		})
	}
	return nil
}

// AddFileFromDisk reads a file from disk and adds it at a given path to a zip file.
func AddFileFromDisk(z *zip.Writer, zipPath, diskPath string) error {
	w, err := z.Create(zipPath)
	if err != nil {
		return fmt.Errorf("error creating %s: %s", zipPath, err)
	}
	f, err := os.Open(diskPath)
	if err != nil {
		return fmt.Errorf("error opening %s: %s", diskPath, err)
	}
	_, err = io.Copy(w, f)
	return err
}

// AddFileFromBytes takes a byte array and adds it at a given path to a zip file.
func AddFileFromBytes(z *zip.Writer, zipPath string, data []byte) error {
	w, err := z.Create(zipPath)
	if err != nil {
		return fmt.Errorf("error creating %s: %s", zipPath, err)
	}
	_, err = io.Copy(w, bytes.NewReader(data))
	return err
}

// ExtractToDiskTmp extracts a zip file to a temporary file in a given path,
// returning the name of the file.
func ExtractToDiskTmp(f *zip.File, path string) (string, error) {
	tmpFile, err := ioutil.TempFile(path, "zz")
	if err != nil {
		return "", err
	}
	defer tmpFile.Close()
	rc, err := f.Open()
	if err != nil {
		return "", err
	}
	defer rc.Close()
	_, err = io.Copy(tmpFile, rc)
	if err != nil {
		return "", err
	}
	return tmpFile.Name(), nil
}
