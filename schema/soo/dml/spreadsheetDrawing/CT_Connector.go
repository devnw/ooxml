package spreadsheetDrawing

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_Connector struct {
	MacroAttr      *string
	FPublishedAttr *bool
	NvCxnSpPr      *CT_ConnectorNonVisual
	SpPr           *dml.CT_ShapeProperties
	Style          *dml.CT_ShapeStyle
}

func NewCT_Connector() *CT_Connector {
	ret := &CT_Connector{}
	ret.NvCxnSpPr = NewCT_ConnectorNonVisual()
	ret.SpPr = dml.NewCT_ShapeProperties()
	return ret
}

func (m *CT_Connector) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.MacroAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "macro"},
			Value: fmt.Sprintf("%v", *m.MacroAttr)})
	}
	if m.FPublishedAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fPublished"},
			Value: fmt.Sprintf("%d", b2i(*m.FPublishedAttr))})
	}
	e.EncodeToken(start)
	senvCxnSpPr := xml.StartElement{Name: xml.Name{Local: "xdr:nvCxnSpPr"}}
	e.EncodeElement(m.NvCxnSpPr, senvCxnSpPr)
	sespPr := xml.StartElement{Name: xml.Name{Local: "xdr:spPr"}}
	e.EncodeElement(m.SpPr, sespPr)
	if m.Style != nil {
		sestyle := xml.StartElement{Name: xml.Name{Local: "xdr:style"}}
		e.EncodeElement(m.Style, sestyle)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Connector) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.NvCxnSpPr = NewCT_ConnectorNonVisual()
	m.SpPr = dml.NewCT_ShapeProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "macro" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MacroAttr = &parsed
			continue
		}
		if attr.Name.Local == "fPublished" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FPublishedAttr = &parsed
			continue
		}
	}
lCT_Connector:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "nvCxnSpPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "nvCxnSpPr"}:
				if err := d.DecodeElement(m.NvCxnSpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "spPr"}:
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "style"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "style"}:
				m.Style = dml.NewCT_ShapeStyle()
				if err := d.DecodeElement(m.Style, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_Connector %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Connector
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Connector and its children
func (m *CT_Connector) Validate() error {
	return m.ValidateWithPath("CT_Connector")
}

// ValidateWithPath validates the CT_Connector and its children, prefixing error messages with path
func (m *CT_Connector) ValidateWithPath(path string) error {
	if err := m.NvCxnSpPr.ValidateWithPath(path + "/NvCxnSpPr"); err != nil {
		return err
	}
	if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
		return err
	}
	if m.Style != nil {
		if err := m.Style.ValidateWithPath(path + "/Style"); err != nil {
			return err
		}
	}
	return nil
}
