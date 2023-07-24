package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_Missing struct {
	// Unused Item
	UAttr *bool
	// Calculated Item
	FAttr *bool
	// Caption
	CAttr *string
	// Member Property Count
	CpAttr *uint32
	// Format Index
	InAttr *uint32
	// background Color
	BcAttr *string
	// Foreground Color
	FcAttr *string
	// Italic
	IAttr *bool
	// Underline
	UnAttr *bool
	// Strikethrough
	StAttr *bool
	// Bold
	BAttr *bool
	// Tuples
	Tpls []*CT_Tuples
	// Member Property Indexes
	X []*CT_X
}

func NewCT_Missing() *CT_Missing {
	ret := &CT_Missing{}
	return ret
}

func (m *CT_Missing) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "u"},
			Value: fmt.Sprintf("%d", b2i(*m.UAttr))})
	}
	if m.FAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "f"},
			Value: fmt.Sprintf("%d", b2i(*m.FAttr))})
	}
	if m.CAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "c"},
			Value: fmt.Sprintf("%v", *m.CAttr)})
	}
	if m.CpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "cp"},
			Value: fmt.Sprintf("%v", *m.CpAttr)})
	}
	if m.InAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "in"},
			Value: fmt.Sprintf("%v", *m.InAttr)})
	}
	if m.BcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bc"},
			Value: fmt.Sprintf("%v", *m.BcAttr)})
	}
	if m.FcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fc"},
			Value: fmt.Sprintf("%v", *m.FcAttr)})
	}
	if m.IAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "i"},
			Value: fmt.Sprintf("%d", b2i(*m.IAttr))})
	}
	if m.UnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "un"},
			Value: fmt.Sprintf("%d", b2i(*m.UnAttr))})
	}
	if m.StAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "st"},
			Value: fmt.Sprintf("%d", b2i(*m.StAttr))})
	}
	if m.BAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "b"},
			Value: fmt.Sprintf("%d", b2i(*m.BAttr))})
	}
	e.EncodeToken(start)
	if m.Tpls != nil {
		setpls := xml.StartElement{Name: xml.Name{Local: "ma:tpls"}}
		for _, c := range m.Tpls {
			e.EncodeElement(c, setpls)
		}
	}
	if m.X != nil {
		sex := xml.StartElement{Name: xml.Name{Local: "ma:x"}}
		for _, c := range m.X {
			e.EncodeElement(c, sex)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Missing) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fc" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FcAttr = &parsed
			continue
		}
		if attr.Name.Local == "f" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FAttr = &parsed
			continue
		}
		if attr.Name.Local == "c" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.CAttr = &parsed
			continue
		}
		if attr.Name.Local == "cp" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.CpAttr = &pt
			continue
		}
		if attr.Name.Local == "in" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.InAttr = &pt
			continue
		}
		if attr.Name.Local == "bc" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BcAttr = &parsed
			continue
		}
		if attr.Name.Local == "u" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UAttr = &parsed
			continue
		}
		if attr.Name.Local == "i" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IAttr = &parsed
			continue
		}
		if attr.Name.Local == "un" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UnAttr = &parsed
			continue
		}
		if attr.Name.Local == "st" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StAttr = &parsed
			continue
		}
		if attr.Name.Local == "b" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BAttr = &parsed
			continue
		}
	}
lCT_Missing:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "tpls"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "tpls"}:
				tmp := NewCT_Tuples()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Tpls = append(m.Tpls, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "x"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "x"}:
				tmp := NewCT_X()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.X = append(m.X, tmp)
			default:
				office.Log("skipping unsupported element on CT_Missing %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Missing
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Missing and its children
func (m *CT_Missing) Validate() error {
	return m.ValidateWithPath("CT_Missing")
}

// ValidateWithPath validates the CT_Missing and its children, prefixing error messages with path
func (m *CT_Missing) ValidateWithPath(path string) error {
	for i, v := range m.Tpls {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Tpls[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.X {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/X[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
