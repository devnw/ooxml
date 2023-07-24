package wml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type WdCT_LinkedTextboxInformation struct {
	IdAttr  uint16
	SeqAttr uint16
	ExtLst  *dml.CT_OfficeArtExtensionList
}

func NewWdCT_LinkedTextboxInformation() *WdCT_LinkedTextboxInformation {
	ret := &WdCT_LinkedTextboxInformation{}
	return ret
}

func (m *WdCT_LinkedTextboxInformation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "seq"},
		Value: fmt.Sprintf("%v", m.SeqAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "wp:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_LinkedTextboxInformation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			m.IdAttr = uint16(parsed)
			continue
		}
		if attr.Name.Local == "seq" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 16)
			if err != nil {
				return err
			}
			m.SeqAttr = uint16(parsed)
			continue
		}
	}
lWdCT_LinkedTextboxInformation:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on WdCT_LinkedTextboxInformation %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_LinkedTextboxInformation
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_LinkedTextboxInformation and its children
func (m *WdCT_LinkedTextboxInformation) Validate() error {
	return m.ValidateWithPath("WdCT_LinkedTextboxInformation")
}

// ValidateWithPath validates the WdCT_LinkedTextboxInformation and its children, prefixing error messages with path
func (m *WdCT_LinkedTextboxInformation) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
