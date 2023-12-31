package vml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type Imagedata struct {
	CT_ImageData
}

func NewImagedata() *Imagedata {
	ret := &Imagedata{}
	ret.CT_ImageData = *NewCT_ImageData()
	return ret
}

func (m *Imagedata) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_ImageData.MarshalXML(e, start)
}

func (m *Imagedata) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_ImageData = *NewCT_ImageData()
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "pict" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "pict" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.PictAttr = &parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "href" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RHrefAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "althref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlthrefAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "title" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TitleAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "oleid" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.OleidAttr = &pt
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "detectmouseclick" {
			m.DetectmouseclickAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "movie" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.MovieAttr = &pt
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "relid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RelidAttr = &parsed
			continue
		}
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SIdAttr = &parsed
			continue
		}
		if attr.Name.Local == "cropbottom" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CropbottomAttr = &parsed
			continue
		}
		if attr.Name.Local == "embosscolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EmbosscolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "src" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SrcAttr = &parsed
			continue
		}
		if attr.Name.Local == "cropleft" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CropleftAttr = &parsed
			continue
		}
		if attr.Name.Local == "croptop" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CroptopAttr = &parsed
			continue
		}
		if attr.Name.Local == "cropright" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CroprightAttr = &parsed
			continue
		}
		if attr.Name.Local == "recolortarget" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RecolortargetAttr = &parsed
			continue
		}
		if attr.Name.Local == "gain" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GainAttr = &parsed
			continue
		}
		if attr.Name.Local == "blacklevel" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BlacklevelAttr = &parsed
			continue
		}
		if attr.Name.Local == "gamma" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.GammaAttr = &parsed
			continue
		}
		if attr.Name.Local == "grayscale" {
			m.GrayscaleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "bilevel" {
			m.BilevelAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "chromakey" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ChromakeyAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing Imagedata: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the Imagedata and its children
func (m *Imagedata) Validate() error {
	return m.ValidateWithPath("Imagedata")
}

// ValidateWithPath validates the Imagedata and its children, prefixing error messages with path
func (m *Imagedata) ValidateWithPath(path string) error {
	if err := m.CT_ImageData.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
