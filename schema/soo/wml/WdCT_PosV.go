package wml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type WdCT_PosV struct {
	RelativeFromAttr WdST_RelFromV
	Choice           *WdCT_PosVChoice
}

func NewWdCT_PosV() *WdCT_PosV {
	ret := &WdCT_PosV{}
	ret.RelativeFromAttr = WdST_RelFromV(1)
	ret.Choice = NewWdCT_PosVChoice()
	return ret
}

func (m *WdCT_PosV) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.RelativeFromAttr.MarshalXMLAttr(xml.Name{Local: "relativeFrom"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	m.Choice.MarshalXML(e, xml.StartElement{})
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *WdCT_PosV) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RelativeFromAttr = WdST_RelFromV(1)
	m.Choice = NewWdCT_PosVChoice()
	for _, attr := range start.Attr {
		if attr.Name.Local == "relativeFrom" {
			m.RelativeFromAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lWdCT_PosV:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "align"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "align"}:
				m.Choice = NewWdCT_PosVChoice()
				if err := d.DecodeElement(&m.Choice.Align, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing", Local: "posOffset"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/wordprocessingDrawing", Local: "posOffset"}:
				m.Choice = NewWdCT_PosVChoice()
				if err := d.DecodeElement(&m.Choice.PosOffset, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on WdCT_PosV %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lWdCT_PosV
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the WdCT_PosV and its children
func (m *WdCT_PosV) Validate() error {
	return m.ValidateWithPath("WdCT_PosV")
}

// ValidateWithPath validates the WdCT_PosV and its children, prefixing error messages with path
func (m *WdCT_PosV) ValidateWithPath(path string) error {
	if m.RelativeFromAttr == WdST_RelFromVUnset {
		return fmt.Errorf("%s/RelativeFromAttr is a mandatory field", path)
	}
	if err := m.RelativeFromAttr.ValidateWithPath(path + "/RelativeFromAttr"); err != nil {
		return err
	}
	if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
		return err
	}
	return nil
}
