package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_EffectList struct {
	Blur        *CT_BlurEffect
	FillOverlay *CT_FillOverlayEffect
	Glow        *CT_GlowEffect
	InnerShdw   *CT_InnerShadowEffect
	OuterShdw   *CT_OuterShadowEffect
	PrstShdw    *CT_PresetShadowEffect
	Reflection  *CT_ReflectionEffect
	SoftEdge    *CT_SoftEdgesEffect
}

func NewCT_EffectList() *CT_EffectList {
	ret := &CT_EffectList{}
	return ret
}

func (m *CT_EffectList) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Blur != nil {
		seblur := xml.StartElement{Name: xml.Name{Local: "a:blur"}}
		e.EncodeElement(m.Blur, seblur)
	}
	if m.FillOverlay != nil {
		sefillOverlay := xml.StartElement{Name: xml.Name{Local: "a:fillOverlay"}}
		e.EncodeElement(m.FillOverlay, sefillOverlay)
	}
	if m.Glow != nil {
		seglow := xml.StartElement{Name: xml.Name{Local: "a:glow"}}
		e.EncodeElement(m.Glow, seglow)
	}
	if m.InnerShdw != nil {
		seinnerShdw := xml.StartElement{Name: xml.Name{Local: "a:innerShdw"}}
		e.EncodeElement(m.InnerShdw, seinnerShdw)
	}
	if m.OuterShdw != nil {
		seouterShdw := xml.StartElement{Name: xml.Name{Local: "a:outerShdw"}}
		e.EncodeElement(m.OuterShdw, seouterShdw)
	}
	if m.PrstShdw != nil {
		seprstShdw := xml.StartElement{Name: xml.Name{Local: "a:prstShdw"}}
		e.EncodeElement(m.PrstShdw, seprstShdw)
	}
	if m.Reflection != nil {
		sereflection := xml.StartElement{Name: xml.Name{Local: "a:reflection"}}
		e.EncodeElement(m.Reflection, sereflection)
	}
	if m.SoftEdge != nil {
		sesoftEdge := xml.StartElement{Name: xml.Name{Local: "a:softEdge"}}
		e.EncodeElement(m.SoftEdge, sesoftEdge)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EffectList) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_EffectList:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "blur"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "blur"}:
				m.Blur = NewCT_BlurEffect()
				if err := d.DecodeElement(m.Blur, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fillOverlay"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fillOverlay"}:
				m.FillOverlay = NewCT_FillOverlayEffect()
				if err := d.DecodeElement(m.FillOverlay, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "glow"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "glow"}:
				m.Glow = NewCT_GlowEffect()
				if err := d.DecodeElement(m.Glow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "innerShdw"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "innerShdw"}:
				m.InnerShdw = NewCT_InnerShadowEffect()
				if err := d.DecodeElement(m.InnerShdw, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "outerShdw"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "outerShdw"}:
				m.OuterShdw = NewCT_OuterShadowEffect()
				if err := d.DecodeElement(m.OuterShdw, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstShdw"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "prstShdw"}:
				m.PrstShdw = NewCT_PresetShadowEffect()
				if err := d.DecodeElement(m.PrstShdw, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "reflection"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "reflection"}:
				m.Reflection = NewCT_ReflectionEffect()
				if err := d.DecodeElement(m.Reflection, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "softEdge"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "softEdge"}:
				m.SoftEdge = NewCT_SoftEdgesEffect()
				if err := d.DecodeElement(m.SoftEdge, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_EffectList %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_EffectList
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_EffectList and its children
func (m *CT_EffectList) Validate() error {
	return m.ValidateWithPath("CT_EffectList")
}

// ValidateWithPath validates the CT_EffectList and its children, prefixing error messages with path
func (m *CT_EffectList) ValidateWithPath(path string) error {
	if m.Blur != nil {
		if err := m.Blur.ValidateWithPath(path + "/Blur"); err != nil {
			return err
		}
	}
	if m.FillOverlay != nil {
		if err := m.FillOverlay.ValidateWithPath(path + "/FillOverlay"); err != nil {
			return err
		}
	}
	if m.Glow != nil {
		if err := m.Glow.ValidateWithPath(path + "/Glow"); err != nil {
			return err
		}
	}
	if m.InnerShdw != nil {
		if err := m.InnerShdw.ValidateWithPath(path + "/InnerShdw"); err != nil {
			return err
		}
	}
	if m.OuterShdw != nil {
		if err := m.OuterShdw.ValidateWithPath(path + "/OuterShdw"); err != nil {
			return err
		}
	}
	if m.PrstShdw != nil {
		if err := m.PrstShdw.ValidateWithPath(path + "/PrstShdw"); err != nil {
			return err
		}
	}
	if m.Reflection != nil {
		if err := m.Reflection.ValidateWithPath(path + "/Reflection"); err != nil {
			return err
		}
	}
	if m.SoftEdge != nil {
		if err := m.SoftEdge.ValidateWithPath(path + "/SoftEdge"); err != nil {
			return err
		}
	}
	return nil
}
