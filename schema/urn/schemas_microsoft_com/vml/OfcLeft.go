package vml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type OfcLeft struct {
	OfcCT_StrokeChild
}

func NewOfcLeft() *OfcLeft {
	ret := &OfcLeft{}
	ret.OfcCT_StrokeChild = *NewOfcCT_StrokeChild()
	return ret
}

func (m *OfcLeft) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:left"
	return m.OfcCT_StrokeChild.MarshalXML(e, start)
}

func (m *OfcLeft) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_StrokeChild = *NewOfcCT_StrokeChild()
	for _, attr := range start.Attr {
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "href" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.HrefAttr = &parsed
			continue
		}
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "forcedash" {
			m.ForcedashAttr.UnmarshalXMLAttr(attr)
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
		if attr.Name.Space == "urn:schemas-microsoft-com:office:office" && attr.Name.Local == "althref" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AlthrefAttr = &parsed
			continue
		}
		if attr.Name.Local == "imageaspect" {
			m.ImageaspectAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "startarrow" {
			m.StartarrowAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "linestyle" {
			m.LinestyleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "startarrowwidth" {
			m.StartarrowwidthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "joinstyle" {
			m.JoinstyleAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "startarrowlength" {
			m.StartarrowlengthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dashstyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DashstyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "endarrow" {
			m.EndarrowAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "filltype" {
			m.FilltypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "endarrowwidth" {
			m.EndarrowwidthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "opacity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OpacityAttr = &parsed
			continue
		}
		if attr.Name.Local == "color" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
			continue
		}
		if attr.Name.Local == "insetpen" {
			m.InsetpenAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "endarrowlength" {
			m.EndarrowlengthAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "endcap" {
			m.EndcapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "color2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Color2Attr = &parsed
			continue
		}
		if attr.Name.Local == "imagealignshape" {
			m.ImagealignshapeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "weight" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.WeightAttr = &parsed
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
		if attr.Name.Local == "imagesize" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ImagesizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "miterlimit" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			m.MiterlimitAttr = &parsed
			continue
		}
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcLeft: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcLeft and its children
func (m *OfcLeft) Validate() error {
	return m.ValidateWithPath("OfcLeft")
}

// ValidateWithPath validates the OfcLeft and its children, prefixing error messages with path
func (m *OfcLeft) ValidateWithPath(path string) error {
	if err := m.OfcCT_StrokeChild.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
