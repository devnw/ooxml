package sml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_CellWatches struct {
	// Cell Watch Item
	CellWatch []*CT_CellWatch
}

func NewCT_CellWatches() *CT_CellWatches {
	ret := &CT_CellWatches{}
	return ret
}

func (m *CT_CellWatches) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secellWatch := xml.StartElement{Name: xml.Name{Local: "ma:cellWatch"}}
	for _, c := range m.CellWatch {
		e.EncodeElement(c, secellWatch)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CellWatches) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_CellWatches:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "cellWatch"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "cellWatch"}:
				tmp := NewCT_CellWatch()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CellWatch = append(m.CellWatch, tmp)
			default:
				office.Log("skipping unsupported element on CT_CellWatches %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CellWatches
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CellWatches and its children
func (m *CT_CellWatches) Validate() error {
	return m.ValidateWithPath("CT_CellWatches")
}

// ValidateWithPath validates the CT_CellWatches and its children, prefixing error messages with path
func (m *CT_CellWatches) ValidateWithPath(path string) error {
	for i, v := range m.CellWatch {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CellWatch[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
