package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Sheet struct {
	// Sheet Name
	NameAttr string
	// Sheet Tab Id
	SheetIdAttr uint32
	// Visible State
	StateAttr ST_SheetState
	IdAttr    string
}

func NewCT_Sheet() *CT_Sheet {
	ret := &CT_Sheet{}
	return ret
}

func (m *CT_Sheet) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "sheetId"},
		Value: fmt.Sprintf("%v", m.SheetIdAttr)})
	if m.StateAttr != ST_SheetStateUnset {
		attr, err := m.StateAttr.MarshalXMLAttr(xml.Name{Local: "state"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Sheet) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "id" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "sheetId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SheetIdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "state" {
			m.StateAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Sheet: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Sheet and its children
func (m *CT_Sheet) Validate() error {
	return m.ValidateWithPath("CT_Sheet")
}

// ValidateWithPath validates the CT_Sheet and its children, prefixing error messages with path
func (m *CT_Sheet) ValidateWithPath(path string) error {
	if err := m.StateAttr.ValidateWithPath(path + "/StateAttr"); err != nil {
		return err
	}
	return nil
}
