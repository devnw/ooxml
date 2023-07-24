package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_AlphaModulateEffect struct {
	Cont *CT_EffectContainer
}

func NewCT_AlphaModulateEffect() *CT_AlphaModulateEffect {
	ret := &CT_AlphaModulateEffect{}
	ret.Cont = NewCT_EffectContainer()
	return ret
}

func (m *CT_AlphaModulateEffect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secont := xml.StartElement{Name: xml.Name{Local: "a:cont"}}
	e.EncodeElement(m.Cont, secont)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_AlphaModulateEffect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Cont = NewCT_EffectContainer()
lCT_AlphaModulateEffect:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "cont"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "cont"}:
				if err := d.DecodeElement(m.Cont, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_AlphaModulateEffect %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_AlphaModulateEffect
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_AlphaModulateEffect and its children
func (m *CT_AlphaModulateEffect) Validate() error {
	return m.ValidateWithPath("CT_AlphaModulateEffect")
}

// ValidateWithPath validates the CT_AlphaModulateEffect and its children, prefixing error messages with path
func (m *CT_AlphaModulateEffect) ValidateWithPath(path string) error {
	if err := m.Cont.ValidateWithPath(path + "/Cont"); err != nil {
		return err
	}
	return nil
}
