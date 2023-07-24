package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_SingleXmlCells struct {
	// Table Properties
	SingleXmlCell []*CT_SingleXmlCell
}

func NewCT_SingleXmlCells() *CT_SingleXmlCells {
	ret := &CT_SingleXmlCells{}
	return ret
}

func (m *CT_SingleXmlCells) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sesingleXmlCell := xml.StartElement{Name: xml.Name{Local: "ma:singleXmlCell"}}
	for _, c := range m.SingleXmlCell {
		e.EncodeElement(c, sesingleXmlCell)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SingleXmlCells) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SingleXmlCells:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "singleXmlCell"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "singleXmlCell"}:
				tmp := NewCT_SingleXmlCell()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.SingleXmlCell = append(m.SingleXmlCell, tmp)
			default:
				office.Log("skipping unsupported element on CT_SingleXmlCells %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SingleXmlCells
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SingleXmlCells and its children
func (m *CT_SingleXmlCells) Validate() error {
	return m.ValidateWithPath("CT_SingleXmlCells")
}

// ValidateWithPath validates the CT_SingleXmlCells and its children, prefixing error messages with path
func (m *CT_SingleXmlCells) ValidateWithPath(path string) error {
	for i, v := range m.SingleXmlCell {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/SingleXmlCell[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
