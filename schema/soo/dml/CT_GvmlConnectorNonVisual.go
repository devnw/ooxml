package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_GvmlConnectorNonVisual struct {
	CNvPr      *CT_NonVisualDrawingProps
	CNvCxnSpPr *CT_NonVisualConnectorProperties
}

func NewCT_GvmlConnectorNonVisual() *CT_GvmlConnectorNonVisual {
	ret := &CT_GvmlConnectorNonVisual{}
	ret.CNvPr = NewCT_NonVisualDrawingProps()
	ret.CNvCxnSpPr = NewCT_NonVisualConnectorProperties()
	return ret
}

func (m *CT_GvmlConnectorNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "a:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvCxnSpPr := xml.StartElement{Name: xml.Name{Local: "a:cNvCxnSpPr"}}
	e.EncodeElement(m.CNvCxnSpPr, secNvCxnSpPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlConnectorNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = NewCT_NonVisualDrawingProps()
	m.CNvCxnSpPr = NewCT_NonVisualConnectorProperties()
lCT_GvmlConnectorNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cNvCxnSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cNvCxnSpPr"}:
				if err := d.DecodeElement(m.CNvCxnSpPr, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_GvmlConnectorNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlConnectorNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlConnectorNonVisual and its children
func (m *CT_GvmlConnectorNonVisual) Validate() error {
	return m.ValidateWithPath("CT_GvmlConnectorNonVisual")
}

// ValidateWithPath validates the CT_GvmlConnectorNonVisual and its children, prefixing error messages with path
func (m *CT_GvmlConnectorNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvCxnSpPr.ValidateWithPath(path + "/CNvCxnSpPr"); err != nil {
		return err
	}
	return nil
}
