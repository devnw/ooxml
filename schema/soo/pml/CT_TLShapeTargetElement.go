package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_TLShapeTargetElement struct {
	// Shape ID
	SpidAttr uint32
	// Background
	Bg *CT_Empty
	// Subshape
	SubSp *CT_TLSubShapeId
	// Embedded Chart Element
	OleChartEl *CT_TLOleChartTargetElement
	// Text Element
	TxEl *CT_TLTextTargetElement
	// Graphic Element
	GraphicEl *dml.CT_AnimationElementChoice
}

func NewCT_TLShapeTargetElement() *CT_TLShapeTargetElement {
	ret := &CT_TLShapeTargetElement{}
	return ret
}

func (m *CT_TLShapeTargetElement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spid"},
		Value: fmt.Sprintf("%v", m.SpidAttr)})
	e.EncodeToken(start)
	if m.Bg != nil {
		sebg := xml.StartElement{Name: xml.Name{Local: "p:bg"}}
		e.EncodeElement(m.Bg, sebg)
	}
	if m.SubSp != nil {
		sesubSp := xml.StartElement{Name: xml.Name{Local: "p:subSp"}}
		e.EncodeElement(m.SubSp, sesubSp)
	}
	if m.OleChartEl != nil {
		seoleChartEl := xml.StartElement{Name: xml.Name{Local: "p:oleChartEl"}}
		e.EncodeElement(m.OleChartEl, seoleChartEl)
	}
	if m.TxEl != nil {
		setxEl := xml.StartElement{Name: xml.Name{Local: "p:txEl"}}
		e.EncodeElement(m.TxEl, setxEl)
	}
	if m.GraphicEl != nil {
		segraphicEl := xml.StartElement{Name: xml.Name{Local: "p:graphicEl"}}
		e.EncodeElement(m.GraphicEl, segraphicEl)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLShapeTargetElement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spid" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.SpidAttr = uint32(parsed)
			continue
		}
	}
lCT_TLShapeTargetElement:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bg"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bg"}:
				m.Bg = NewCT_Empty()
				if err := d.DecodeElement(m.Bg, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "subSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "subSp"}:
				m.SubSp = NewCT_TLSubShapeId()
				if err := d.DecodeElement(m.SubSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "oleChartEl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "oleChartEl"}:
				m.OleChartEl = NewCT_TLOleChartTargetElement()
				if err := d.DecodeElement(m.OleChartEl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "txEl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "txEl"}:
				m.TxEl = NewCT_TLTextTargetElement()
				if err := d.DecodeElement(m.TxEl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "graphicEl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "graphicEl"}:
				m.GraphicEl = dml.NewCT_AnimationElementChoice()
				if err := d.DecodeElement(m.GraphicEl, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_TLShapeTargetElement %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLShapeTargetElement
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLShapeTargetElement and its children
func (m *CT_TLShapeTargetElement) Validate() error {
	return m.ValidateWithPath("CT_TLShapeTargetElement")
}

// ValidateWithPath validates the CT_TLShapeTargetElement and its children, prefixing error messages with path
func (m *CT_TLShapeTargetElement) ValidateWithPath(path string) error {
	if m.Bg != nil {
		if err := m.Bg.ValidateWithPath(path + "/Bg"); err != nil {
			return err
		}
	}
	if m.SubSp != nil {
		if err := m.SubSp.ValidateWithPath(path + "/SubSp"); err != nil {
			return err
		}
	}
	if m.OleChartEl != nil {
		if err := m.OleChartEl.ValidateWithPath(path + "/OleChartEl"); err != nil {
			return err
		}
	}
	if m.TxEl != nil {
		if err := m.TxEl.ValidateWithPath(path + "/TxEl"); err != nil {
			return err
		}
	}
	if m.GraphicEl != nil {
		if err := m.GraphicEl.ValidateWithPath(path + "/GraphicEl"); err != nil {
			return err
		}
	}
	return nil
}
