package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_TLAnimateScaleBehavior struct {
	// Zoom Content
	ZoomContentsAttr *bool
	CBhvr            *CT_TLCommonBehaviorData
	// By
	By   *CT_TLPoint
	From *CT_TLPoint
	// To
	To *CT_TLPoint
}

func NewCT_TLAnimateScaleBehavior() *CT_TLAnimateScaleBehavior {
	ret := &CT_TLAnimateScaleBehavior{}
	ret.CBhvr = NewCT_TLCommonBehaviorData()
	return ret
}

func (m *CT_TLAnimateScaleBehavior) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.ZoomContentsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "zoomContents"},
			Value: fmt.Sprintf("%d", b2i(*m.ZoomContentsAttr))})
	}
	e.EncodeToken(start)
	secBhvr := xml.StartElement{Name: xml.Name{Local: "p:cBhvr"}}
	e.EncodeElement(m.CBhvr, secBhvr)
	if m.By != nil {
		seby := xml.StartElement{Name: xml.Name{Local: "p:by"}}
		e.EncodeElement(m.By, seby)
	}
	if m.From != nil {
		sefrom := xml.StartElement{Name: xml.Name{Local: "p:from"}}
		e.EncodeElement(m.From, sefrom)
	}
	if m.To != nil {
		seto := xml.StartElement{Name: xml.Name{Local: "p:to"}}
		e.EncodeElement(m.To, seto)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TLAnimateScaleBehavior) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CBhvr = NewCT_TLCommonBehaviorData()
	for _, attr := range start.Attr {
		if attr.Name.Local == "zoomContents" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ZoomContentsAttr = &parsed
			continue
		}
	}
lCT_TLAnimateScaleBehavior:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "cBhvr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "cBhvr"}:
				if err := d.DecodeElement(m.CBhvr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "by"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "by"}:
				m.By = NewCT_TLPoint()
				if err := d.DecodeElement(m.By, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "from"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "from"}:
				m.From = NewCT_TLPoint()
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "to"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "to"}:
				m.To = NewCT_TLPoint()
				if err := d.DecodeElement(m.To, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_TLAnimateScaleBehavior %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TLAnimateScaleBehavior
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TLAnimateScaleBehavior and its children
func (m *CT_TLAnimateScaleBehavior) Validate() error {
	return m.ValidateWithPath("CT_TLAnimateScaleBehavior")
}

// ValidateWithPath validates the CT_TLAnimateScaleBehavior and its children, prefixing error messages with path
func (m *CT_TLAnimateScaleBehavior) ValidateWithPath(path string) error {
	if err := m.CBhvr.ValidateWithPath(path + "/CBhvr"); err != nil {
		return err
	}
	if m.By != nil {
		if err := m.By.ValidateWithPath(path + "/By"); err != nil {
			return err
		}
	}
	if m.From != nil {
		if err := m.From.ValidateWithPath(path + "/From"); err != nil {
			return err
		}
	}
	if m.To != nil {
		if err := m.To.ValidateWithPath(path + "/To"); err != nil {
			return err
		}
	}
	return nil
}
