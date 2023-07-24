package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_GraphicalObjectFrameLocking struct {
	NoGrpAttr          *bool
	NoDrilldownAttr    *bool
	NoSelectAttr       *bool
	NoChangeAspectAttr *bool
	NoMoveAttr         *bool
	NoResizeAttr       *bool
	ExtLst             *CT_OfficeArtExtensionList
}

func NewCT_GraphicalObjectFrameLocking() *CT_GraphicalObjectFrameLocking {
	ret := &CT_GraphicalObjectFrameLocking{}
	return ret
}

func (m *CT_GraphicalObjectFrameLocking) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NoGrpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noGrp"},
			Value: fmt.Sprintf("%d", b2i(*m.NoGrpAttr))})
	}
	if m.NoDrilldownAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noDrilldown"},
			Value: fmt.Sprintf("%d", b2i(*m.NoDrilldownAttr))})
	}
	if m.NoSelectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noSelect"},
			Value: fmt.Sprintf("%d", b2i(*m.NoSelectAttr))})
	}
	if m.NoChangeAspectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeAspect"},
			Value: fmt.Sprintf("%d", b2i(*m.NoChangeAspectAttr))})
	}
	if m.NoMoveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noMove"},
			Value: fmt.Sprintf("%d", b2i(*m.NoMoveAttr))})
	}
	if m.NoResizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noResize"},
			Value: fmt.Sprintf("%d", b2i(*m.NoResizeAttr))})
	}
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GraphicalObjectFrameLocking) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "noGrp" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoGrpAttr = &parsed
			continue
		}
		if attr.Name.Local == "noDrilldown" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoDrilldownAttr = &parsed
			continue
		}
		if attr.Name.Local == "noSelect" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoSelectAttr = &parsed
			continue
		}
		if attr.Name.Local == "noChangeAspect" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoChangeAspectAttr = &parsed
			continue
		}
		if attr.Name.Local == "noMove" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoMoveAttr = &parsed
			continue
		}
		if attr.Name.Local == "noResize" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoResizeAttr = &parsed
			continue
		}
	}
lCT_GraphicalObjectFrameLocking:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_GraphicalObjectFrameLocking %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GraphicalObjectFrameLocking
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GraphicalObjectFrameLocking and its children
func (m *CT_GraphicalObjectFrameLocking) Validate() error {
	return m.ValidateWithPath("CT_GraphicalObjectFrameLocking")
}

// ValidateWithPath validates the CT_GraphicalObjectFrameLocking and its children, prefixing error messages with path
func (m *CT_GraphicalObjectFrameLocking) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
