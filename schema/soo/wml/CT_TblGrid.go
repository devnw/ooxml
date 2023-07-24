package wml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_TblGrid struct {
	// Grid Column Definition
	GridCol       []*CT_TblGridCol
	TblGridChange *CT_TblGridChange
}

func NewCT_TblGrid() *CT_TblGrid {
	ret := &CT_TblGrid{}
	return ret
}

func (m *CT_TblGrid) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.GridCol != nil {
		segridCol := xml.StartElement{Name: xml.Name{Local: "w:gridCol"}}
		for _, c := range m.GridCol {
			e.EncodeElement(c, segridCol)
		}
	}
	if m.TblGridChange != nil {
		setblGridChange := xml.StartElement{Name: xml.Name{Local: "w:tblGridChange"}}
		e.EncodeElement(m.TblGridChange, setblGridChange)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TblGrid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TblGrid:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "gridCol"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "gridCol"}:
				tmp := NewCT_TblGridCol()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GridCol = append(m.GridCol, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tblGridChange"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "tblGridChange"}:
				m.TblGridChange = NewCT_TblGridChange()
				if err := d.DecodeElement(m.TblGridChange, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_TblGrid %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TblGrid
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TblGrid and its children
func (m *CT_TblGrid) Validate() error {
	return m.ValidateWithPath("CT_TblGrid")
}

// ValidateWithPath validates the CT_TblGrid and its children, prefixing error messages with path
func (m *CT_TblGrid) ValidateWithPath(path string) error {
	for i, v := range m.GridCol {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GridCol[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.TblGridChange != nil {
		if err := m.TblGridChange.ValidateWithPath(path + "/TblGridChange"); err != nil {
			return err
		}
	}
	return nil
}
