package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_ServerFormats struct {
	// Format Count
	CountAttr *uint32
	// Server Format
	ServerFormat []*CT_ServerFormat
}

func NewCT_ServerFormats() *CT_ServerFormats {
	ret := &CT_ServerFormats{}
	return ret
}

func (m *CT_ServerFormats) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.ServerFormat != nil {
		seserverFormat := xml.StartElement{Name: xml.Name{Local: "ma:serverFormat"}}
		for _, c := range m.ServerFormat {
			e.EncodeElement(c, seserverFormat)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ServerFormats) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "count" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CountAttr = &pt
			continue
		}
	}
lCT_ServerFormats:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "serverFormat"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "serverFormat"}:
				tmp := NewCT_ServerFormat()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ServerFormat = append(m.ServerFormat, tmp)
			default:
				office.Log("skipping unsupported element on CT_ServerFormats %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ServerFormats
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ServerFormats and its children
func (m *CT_ServerFormats) Validate() error {
	return m.ValidateWithPath("CT_ServerFormats")
}

// ValidateWithPath validates the CT_ServerFormats and its children, prefixing error messages with path
func (m *CT_ServerFormats) ValidateWithPath(path string) error {
	for i, v := range m.ServerFormat {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ServerFormat[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
