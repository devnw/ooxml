package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_TableStyleTextStyle struct {
	BAttr     ST_OnOffStyleType
	IAttr     ST_OnOffStyleType
	Font      *CT_FontCollection
	FontRef   *CT_FontReference
	ScrgbClr  *CT_ScRgbColor
	SrgbClr   *CT_SRgbColor
	HslClr    *CT_HslColor
	SysClr    *CT_SystemColor
	SchemeClr *CT_SchemeColor
	PrstClr   *CT_PresetColor
	ExtLst    *CT_OfficeArtExtensionList
}

func NewCT_TableStyleTextStyle() *CT_TableStyleTextStyle {
	ret := &CT_TableStyleTextStyle{}
	return ret
}

func (m *CT_TableStyleTextStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.BAttr != ST_OnOffStyleTypeUnset {
		attr, err := m.BAttr.MarshalXMLAttr(xml.Name{Local: "b"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.IAttr != ST_OnOffStyleTypeUnset {
		attr, err := m.IAttr.MarshalXMLAttr(xml.Name{Local: "i"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.Font != nil {
		sefont := xml.StartElement{Name: xml.Name{Local: "a:font"}}
		e.EncodeElement(m.Font, sefont)
	}
	if m.FontRef != nil {
		sefontRef := xml.StartElement{Name: xml.Name{Local: "a:fontRef"}}
		e.EncodeElement(m.FontRef, sefontRef)
	}
	if m.ScrgbClr != nil {
		sescrgbClr := xml.StartElement{Name: xml.Name{Local: "a:scrgbClr"}}
		e.EncodeElement(m.ScrgbClr, sescrgbClr)
	}
	if m.SrgbClr != nil {
		sesrgbClr := xml.StartElement{Name: xml.Name{Local: "a:srgbClr"}}
		e.EncodeElement(m.SrgbClr, sesrgbClr)
	}
	if m.HslClr != nil {
		sehslClr := xml.StartElement{Name: xml.Name{Local: "a:hslClr"}}
		e.EncodeElement(m.HslClr, sehslClr)
	}
	if m.SysClr != nil {
		sesysClr := xml.StartElement{Name: xml.Name{Local: "a:sysClr"}}
		e.EncodeElement(m.SysClr, sesysClr)
	}
	if m.SchemeClr != nil {
		seschemeClr := xml.StartElement{Name: xml.Name{Local: "a:schemeClr"}}
		e.EncodeElement(m.SchemeClr, seschemeClr)
	}
	if m.PrstClr != nil {
		seprstClr := xml.StartElement{Name: xml.Name{Local: "a:prstClr"}}
		e.EncodeElement(m.PrstClr, seprstClr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TableStyleTextStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "b" {
			m.BAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "i" {
			m.IAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_TableStyleTextStyle:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "font"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "font"}:
				m.Font = NewCT_FontCollection()
				if err := d.DecodeElement(m.Font, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "fontRef"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "fontRef"}:
				m.FontRef = NewCT_FontReference()
				if err := d.DecodeElement(m.FontRef, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "scrgbClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "scrgbClr"}:
				m.ScrgbClr = NewCT_ScRgbColor()
				if err := d.DecodeElement(m.ScrgbClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "srgbClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "srgbClr"}:
				m.SrgbClr = NewCT_SRgbColor()
				if err := d.DecodeElement(m.SrgbClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "hslClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "hslClr"}:
				m.HslClr = NewCT_HslColor()
				if err := d.DecodeElement(m.HslClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sysClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "sysClr"}:
				m.SysClr = NewCT_SystemColor()
				if err := d.DecodeElement(m.SysClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "schemeClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "schemeClr"}:
				m.SchemeClr = NewCT_SchemeColor()
				if err := d.DecodeElement(m.SchemeClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstClr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "prstClr"}:
				m.PrstClr = NewCT_PresetColor()
				if err := d.DecodeElement(m.PrstClr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_TableStyleTextStyle %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TableStyleTextStyle
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TableStyleTextStyle and its children
func (m *CT_TableStyleTextStyle) Validate() error {
	return m.ValidateWithPath("CT_TableStyleTextStyle")
}

// ValidateWithPath validates the CT_TableStyleTextStyle and its children, prefixing error messages with path
func (m *CT_TableStyleTextStyle) ValidateWithPath(path string) error {
	if err := m.BAttr.ValidateWithPath(path + "/BAttr"); err != nil {
		return err
	}
	if err := m.IAttr.ValidateWithPath(path + "/IAttr"); err != nil {
		return err
	}
	if m.Font != nil {
		if err := m.Font.ValidateWithPath(path + "/Font"); err != nil {
			return err
		}
	}
	if m.FontRef != nil {
		if err := m.FontRef.ValidateWithPath(path + "/FontRef"); err != nil {
			return err
		}
	}
	if m.ScrgbClr != nil {
		if err := m.ScrgbClr.ValidateWithPath(path + "/ScrgbClr"); err != nil {
			return err
		}
	}
	if m.SrgbClr != nil {
		if err := m.SrgbClr.ValidateWithPath(path + "/SrgbClr"); err != nil {
			return err
		}
	}
	if m.HslClr != nil {
		if err := m.HslClr.ValidateWithPath(path + "/HslClr"); err != nil {
			return err
		}
	}
	if m.SysClr != nil {
		if err := m.SysClr.ValidateWithPath(path + "/SysClr"); err != nil {
			return err
		}
	}
	if m.SchemeClr != nil {
		if err := m.SchemeClr.ValidateWithPath(path + "/SchemeClr"); err != nil {
			return err
		}
	}
	if m.PrstClr != nil {
		if err := m.PrstClr.ValidateWithPath(path + "/PrstClr"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
