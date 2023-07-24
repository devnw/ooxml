package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_TableStyle struct {
	// Table Style Name
	NameAttr string
	// Pivot Style
	PivotAttr *bool
	// Table
	TableAttr *bool
	// Table Style Count
	CountAttr *uint32
	// Table Style
	TableStyleElement []*CT_TableStyleElement
}

func NewCT_TableStyle() *CT_TableStyle {
	ret := &CT_TableStyle{}
	return ret
}

func (m *CT_TableStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	if m.PivotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "pivot"},
			Value: fmt.Sprintf("%d", b2i(*m.PivotAttr))})
	}
	if m.TableAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "table"},
			Value: fmt.Sprintf("%d", b2i(*m.TableAttr))})
	}
	if m.CountAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "count"},
			Value: fmt.Sprintf("%v", *m.CountAttr)})
	}
	e.EncodeToken(start)
	if m.TableStyleElement != nil {
		setableStyleElement := xml.StartElement{Name: xml.Name{Local: "ma:tableStyleElement"}}
		for _, c := range m.TableStyleElement {
			e.EncodeElement(c, setableStyleElement)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "pivot" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PivotAttr = &parsed
			continue
		}
		if attr.Name.Local == "table" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.TableAttr = &parsed
			continue
		}
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
lCT_TableStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tableStyleElement"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "tableStyleElement"}:
				tmp := NewCT_TableStyleElement()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TableStyleElement = append(m.TableStyleElement, tmp)
			default:
				office.Log("skipping unsupported element on CT_TableStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableStyle and its children
func (m *CT_TableStyle) Validate() error {
	return m.ValidateWithPath("CT_TableStyle")
}

// ValidateWithPath validates the CT_TableStyle and its children, prefixing error messages with path
func (m *CT_TableStyle) ValidateWithPath(path string) error {
	for i, v := range m.TableStyleElement {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TableStyleElement[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
