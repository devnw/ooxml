package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type CT_Ratio struct {
	NAttr int64
	DAttr int64
}

func NewCT_Ratio() *CT_Ratio {
	ret := &CT_Ratio{}
	return ret
}

func (m *CT_Ratio) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "n"},
		Value: fmt.Sprintf("%v", m.NAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "d"},
		Value: fmt.Sprintf("%v", m.DAttr)})
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Ratio) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "n" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.NAttr = parsed
			continue
		}
		if attr.Name.Local == "d" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.DAttr = parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_Ratio: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_Ratio and its children
func (m *CT_Ratio) Validate() error {
	return m.ValidateWithPath("CT_Ratio")
}

// ValidateWithPath validates the CT_Ratio and its children, prefixing error messages with path
func (m *CT_Ratio) ValidateWithPath(path string) error {
	return nil
}
