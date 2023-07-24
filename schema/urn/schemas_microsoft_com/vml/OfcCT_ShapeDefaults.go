package vml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
)

type OfcCT_ShapeDefaults struct {
	SpidmaxAttr     *int64
	StyleAttr       *string
	FillAttr        sharedTypes.ST_TrueFalse
	FillcolorAttr   *string
	StrokeAttr      sharedTypes.ST_TrueFalse
	StrokecolorAttr *string
	AllowincellAttr sharedTypes.ST_TrueFalse
	Fill            *Fill
	Stroke          *Stroke
	Textbox         *Textbox
	Shadow          *Shadow
	Skew            *OfcSkew
	Extrusion       *OfcExtrusion
	Callout         *OfcCallout
	Lock            *OfcLock
	Colormru        *OfcCT_ColorMru
	Colormenu       *OfcCT_ColorMenu
	ExtAttr         ST_Ext
}

func NewOfcCT_ShapeDefaults() *OfcCT_ShapeDefaults {
	ret := &OfcCT_ShapeDefaults{}
	return ret
}

func (m *OfcCT_ShapeDefaults) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.SpidmaxAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spidmax"},
			Value: fmt.Sprintf("%v", *m.SpidmaxAttr)})
	}
	if m.StyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "style"},
			Value: fmt.Sprintf("%v", *m.StyleAttr)})
	}
	if m.FillAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.FillAttr.MarshalXMLAttr(xml.Name{Local: "fill"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.FillcolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fillcolor"},
			Value: fmt.Sprintf("%v", *m.FillcolorAttr)})
	}
	if m.StrokeAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.StrokeAttr.MarshalXMLAttr(xml.Name{Local: "stroke"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.StrokecolorAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "strokecolor"},
			Value: fmt.Sprintf("%v", *m.StrokecolorAttr)})
	}
	if m.AllowincellAttr != sharedTypes.ST_TrueFalseUnset {
		attr, err := m.AllowincellAttr.MarshalXMLAttr(xml.Name{Local: "allowincell"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.ExtAttr != ST_ExtUnset {
		attr, err := m.ExtAttr.MarshalXMLAttr(xml.Name{Local: "ext"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.Fill != nil {
		sefill := xml.StartElement{Name: xml.Name{Local: "v:fill"}}
		e.EncodeElement(m.Fill, sefill)
	}
	if m.Stroke != nil {
		sestroke := xml.StartElement{Name: xml.Name{Local: "v:stroke"}}
		e.EncodeElement(m.Stroke, sestroke)
	}
	if m.Textbox != nil {
		setextbox := xml.StartElement{Name: xml.Name{Local: "v:textbox"}}
		e.EncodeElement(m.Textbox, setextbox)
	}
	if m.Shadow != nil {
		seshadow := xml.StartElement{Name: xml.Name{Local: "v:shadow"}}
		e.EncodeElement(m.Shadow, seshadow)
	}
	if m.Skew != nil {
		seskew := xml.StartElement{Name: xml.Name{Local: "o:skew"}}
		e.EncodeElement(m.Skew, seskew)
	}
	if m.Extrusion != nil {
		seextrusion := xml.StartElement{Name: xml.Name{Local: "o:extrusion"}}
		e.EncodeElement(m.Extrusion, seextrusion)
	}
	if m.Callout != nil {
		secallout := xml.StartElement{Name: xml.Name{Local: "o:callout"}}
		e.EncodeElement(m.Callout, secallout)
	}
	if m.Lock != nil {
		selock := xml.StartElement{Name: xml.Name{Local: "o:lock"}}
		e.EncodeElement(m.Lock, selock)
	}
	if m.Colormru != nil {
		secolormru := xml.StartElement{Name: xml.Name{Local: "o:colormru"}}
		e.EncodeElement(m.Colormru, secolormru)
	}
	if m.Colormenu != nil {
		secolormenu := xml.StartElement{Name: xml.Name{Local: "o:colormenu"}}
		e.EncodeElement(m.Colormenu, secolormenu)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *OfcCT_ShapeDefaults) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "spidmax" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 64)
			if err != nil {
				return err
			}
			m.SpidmaxAttr = &parsed
			continue
		}
		if attr.Name.Local == "allowincell" {
			m.AllowincellAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "strokecolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StrokecolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "stroke" {
			m.StrokeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "fillcolor" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FillcolorAttr = &parsed
			continue
		}
		if attr.Name.Local == "fill" {
			m.FillAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "style" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.StyleAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lOfcCT_ShapeDefaults:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "fill"}:
				m.Fill = NewFill()
				if err := d.DecodeElement(m.Fill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "stroke"}:
				m.Stroke = NewStroke()
				if err := d.DecodeElement(m.Stroke, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "textbox"}:
				m.Textbox = NewTextbox()
				if err := d.DecodeElement(m.Textbox, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:vml", Local: "shadow"}:
				m.Shadow = NewShadow()
				if err := d.DecodeElement(m.Shadow, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "skew"}:
				m.Skew = NewOfcSkew()
				if err := d.DecodeElement(m.Skew, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "extrusion"}:
				m.Extrusion = NewOfcExtrusion()
				if err := d.DecodeElement(m.Extrusion, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "callout"}:
				m.Callout = NewOfcCallout()
				if err := d.DecodeElement(m.Callout, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "lock"}:
				m.Lock = NewOfcLock()
				if err := d.DecodeElement(m.Lock, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "colormru"}:
				m.Colormru = NewOfcCT_ColorMru()
				if err := d.DecodeElement(m.Colormru, &el); err != nil {
					return err
				}
			case xml.Name{Space: "urn:schemas-microsoft-com:office:office", Local: "colormenu"}:
				m.Colormenu = NewOfcCT_ColorMenu()
				if err := d.DecodeElement(m.Colormenu, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on OfcCT_ShapeDefaults %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lOfcCT_ShapeDefaults
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the OfcCT_ShapeDefaults and its children
func (m *OfcCT_ShapeDefaults) Validate() error {
	return m.ValidateWithPath("OfcCT_ShapeDefaults")
}

// ValidateWithPath validates the OfcCT_ShapeDefaults and its children, prefixing error messages with path
func (m *OfcCT_ShapeDefaults) ValidateWithPath(path string) error {
	if err := m.FillAttr.ValidateWithPath(path + "/FillAttr"); err != nil {
		return err
	}
	if err := m.StrokeAttr.ValidateWithPath(path + "/StrokeAttr"); err != nil {
		return err
	}
	if err := m.AllowincellAttr.ValidateWithPath(path + "/AllowincellAttr"); err != nil {
		return err
	}
	if m.Fill != nil {
		if err := m.Fill.ValidateWithPath(path + "/Fill"); err != nil {
			return err
		}
	}
	if m.Stroke != nil {
		if err := m.Stroke.ValidateWithPath(path + "/Stroke"); err != nil {
			return err
		}
	}
	if m.Textbox != nil {
		if err := m.Textbox.ValidateWithPath(path + "/Textbox"); err != nil {
			return err
		}
	}
	if m.Shadow != nil {
		if err := m.Shadow.ValidateWithPath(path + "/Shadow"); err != nil {
			return err
		}
	}
	if m.Skew != nil {
		if err := m.Skew.ValidateWithPath(path + "/Skew"); err != nil {
			return err
		}
	}
	if m.Extrusion != nil {
		if err := m.Extrusion.ValidateWithPath(path + "/Extrusion"); err != nil {
			return err
		}
	}
	if m.Callout != nil {
		if err := m.Callout.ValidateWithPath(path + "/Callout"); err != nil {
			return err
		}
	}
	if m.Lock != nil {
		if err := m.Lock.ValidateWithPath(path + "/Lock"); err != nil {
			return err
		}
	}
	if m.Colormru != nil {
		if err := m.Colormru.ValidateWithPath(path + "/Colormru"); err != nil {
			return err
		}
	}
	if m.Colormenu != nil {
		if err := m.Colormenu.ValidateWithPath(path + "/Colormenu"); err != nil {
			return err
		}
	}
	if err := m.ExtAttr.ValidateWithPath(path + "/ExtAttr"); err != nil {
		return err
	}
	return nil
}
