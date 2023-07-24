package chart

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_PictureOptions struct {
	ApplyToFront     *CT_Boolean
	ApplyToSides     *CT_Boolean
	ApplyToEnd       *CT_Boolean
	PictureFormat    *CT_PictureFormat
	PictureStackUnit *CT_PictureStackUnit
}

func NewCT_PictureOptions() *CT_PictureOptions {
	ret := &CT_PictureOptions{}
	return ret
}

func (m *CT_PictureOptions) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ApplyToFront != nil {
		seapplyToFront := xml.StartElement{Name: xml.Name{Local: "c:applyToFront"}}
		e.EncodeElement(m.ApplyToFront, seapplyToFront)
	}
	if m.ApplyToSides != nil {
		seapplyToSides := xml.StartElement{Name: xml.Name{Local: "c:applyToSides"}}
		e.EncodeElement(m.ApplyToSides, seapplyToSides)
	}
	if m.ApplyToEnd != nil {
		seapplyToEnd := xml.StartElement{Name: xml.Name{Local: "c:applyToEnd"}}
		e.EncodeElement(m.ApplyToEnd, seapplyToEnd)
	}
	if m.PictureFormat != nil {
		sepictureFormat := xml.StartElement{Name: xml.Name{Local: "c:pictureFormat"}}
		e.EncodeElement(m.PictureFormat, sepictureFormat)
	}
	if m.PictureStackUnit != nil {
		sepictureStackUnit := xml.StartElement{Name: xml.Name{Local: "c:pictureStackUnit"}}
		e.EncodeElement(m.PictureStackUnit, sepictureStackUnit)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PictureOptions) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_PictureOptions:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "applyToFront"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "applyToFront"}:
				m.ApplyToFront = NewCT_Boolean()
				if err := d.DecodeElement(m.ApplyToFront, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "applyToSides"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "applyToSides"}:
				m.ApplyToSides = NewCT_Boolean()
				if err := d.DecodeElement(m.ApplyToSides, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "applyToEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "applyToEnd"}:
				m.ApplyToEnd = NewCT_Boolean()
				if err := d.DecodeElement(m.ApplyToEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pictureFormat"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "pictureFormat"}:
				m.PictureFormat = NewCT_PictureFormat()
				if err := d.DecodeElement(m.PictureFormat, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "pictureStackUnit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "pictureStackUnit"}:
				m.PictureStackUnit = NewCT_PictureStackUnit()
				if err := d.DecodeElement(m.PictureStackUnit, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_PictureOptions %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PictureOptions
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PictureOptions and its children
func (m *CT_PictureOptions) Validate() error {
	return m.ValidateWithPath("CT_PictureOptions")
}

// ValidateWithPath validates the CT_PictureOptions and its children, prefixing error messages with path
func (m *CT_PictureOptions) ValidateWithPath(path string) error {
	if m.ApplyToFront != nil {
		if err := m.ApplyToFront.ValidateWithPath(path + "/ApplyToFront"); err != nil {
			return err
		}
	}
	if m.ApplyToSides != nil {
		if err := m.ApplyToSides.ValidateWithPath(path + "/ApplyToSides"); err != nil {
			return err
		}
	}
	if m.ApplyToEnd != nil {
		if err := m.ApplyToEnd.ValidateWithPath(path + "/ApplyToEnd"); err != nil {
			return err
		}
	}
	if m.PictureFormat != nil {
		if err := m.PictureFormat.ValidateWithPath(path + "/PictureFormat"); err != nil {
			return err
		}
	}
	if m.PictureStackUnit != nil {
		if err := m.PictureStackUnit.ValidateWithPath(path + "/PictureStackUnit"); err != nil {
			return err
		}
	}
	return nil
}
