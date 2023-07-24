package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_LineProperties struct {
	WAttr     *int32
	CapAttr   ST_LineCap
	CmpdAttr  ST_CompoundLine
	AlgnAttr  ST_PenAlignment
	NoFill    *CT_NoFillProperties
	SolidFill *CT_SolidColorFillProperties
	GradFill  *CT_GradientFillProperties
	PattFill  *CT_PatternFillProperties
	PrstDash  *CT_PresetLineDashProperties
	CustDash  *CT_DashStopList
	Round     *CT_LineJoinRound
	Bevel     *CT_LineJoinBevel
	Miter     *CT_LineJoinMiterProperties
	HeadEnd   *CT_LineEndProperties
	TailEnd   *CT_LineEndProperties
	ExtLst    *CT_OfficeArtExtensionList
}

func NewCT_LineProperties() *CT_LineProperties {
	ret := &CT_LineProperties{}
	return ret
}

func (m *CT_LineProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.WAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "w"},
			Value: fmt.Sprintf("%v", *m.WAttr)})
	}
	if m.CapAttr != ST_LineCapUnset {
		attr, err := m.CapAttr.MarshalXMLAttr(xml.Name{Local: "cap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.CmpdAttr != ST_CompoundLineUnset {
		attr, err := m.CmpdAttr.MarshalXMLAttr(xml.Name{Local: "cmpd"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AlgnAttr != ST_PenAlignmentUnset {
		attr, err := m.AlgnAttr.MarshalXMLAttr(xml.Name{Local: "algn"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	e.EncodeToken(start)
	if m.NoFill != nil {
		senoFill := xml.StartElement{Name: xml.Name{Local: "a:noFill"}}
		e.EncodeElement(m.NoFill, senoFill)
	}
	if m.SolidFill != nil {
		sesolidFill := xml.StartElement{Name: xml.Name{Local: "a:solidFill"}}
		e.EncodeElement(m.SolidFill, sesolidFill)
	}
	if m.GradFill != nil {
		segradFill := xml.StartElement{Name: xml.Name{Local: "a:gradFill"}}
		e.EncodeElement(m.GradFill, segradFill)
	}
	if m.PattFill != nil {
		sepattFill := xml.StartElement{Name: xml.Name{Local: "a:pattFill"}}
		e.EncodeElement(m.PattFill, sepattFill)
	}
	if m.PrstDash != nil {
		seprstDash := xml.StartElement{Name: xml.Name{Local: "a:prstDash"}}
		e.EncodeElement(m.PrstDash, seprstDash)
	}
	if m.CustDash != nil {
		secustDash := xml.StartElement{Name: xml.Name{Local: "a:custDash"}}
		e.EncodeElement(m.CustDash, secustDash)
	}
	if m.Round != nil {
		seround := xml.StartElement{Name: xml.Name{Local: "a:round"}}
		e.EncodeElement(m.Round, seround)
	}
	if m.Bevel != nil {
		sebevel := xml.StartElement{Name: xml.Name{Local: "a:bevel"}}
		e.EncodeElement(m.Bevel, sebevel)
	}
	if m.Miter != nil {
		semiter := xml.StartElement{Name: xml.Name{Local: "a:miter"}}
		e.EncodeElement(m.Miter, semiter)
	}
	if m.HeadEnd != nil {
		seheadEnd := xml.StartElement{Name: xml.Name{Local: "a:headEnd"}}
		e.EncodeElement(m.HeadEnd, seheadEnd)
	}
	if m.TailEnd != nil {
		setailEnd := xml.StartElement{Name: xml.Name{Local: "a:tailEnd"}}
		e.EncodeElement(m.TailEnd, setailEnd)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LineProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "w" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.WAttr = &pt
			continue
		}
		if attr.Name.Local == "cmpd" {
			m.CmpdAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "cap" {
			m.CapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "algn" {
			m.AlgnAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_LineProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "noFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "noFill"}:
				m.NoFill = NewCT_NoFillProperties()
				if err := d.DecodeElement(m.NoFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "solidFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "solidFill"}:
				m.SolidFill = NewCT_SolidColorFillProperties()
				if err := d.DecodeElement(m.SolidFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "gradFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "gradFill"}:
				m.GradFill = NewCT_GradientFillProperties()
				if err := d.DecodeElement(m.GradFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "pattFill"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "pattFill"}:
				m.PattFill = NewCT_PatternFillProperties()
				if err := d.DecodeElement(m.PattFill, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstDash"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "prstDash"}:
				m.PrstDash = NewCT_PresetLineDashProperties()
				if err := d.DecodeElement(m.PrstDash, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "custDash"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "custDash"}:
				m.CustDash = NewCT_DashStopList()
				if err := d.DecodeElement(m.CustDash, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "round"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "round"}:
				m.Round = NewCT_LineJoinRound()
				if err := d.DecodeElement(m.Round, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "bevel"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "bevel"}:
				m.Bevel = NewCT_LineJoinBevel()
				if err := d.DecodeElement(m.Bevel, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "miter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "miter"}:
				m.Miter = NewCT_LineJoinMiterProperties()
				if err := d.DecodeElement(m.Miter, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "headEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "headEnd"}:
				m.HeadEnd = NewCT_LineEndProperties()
				if err := d.DecodeElement(m.HeadEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "tailEnd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "tailEnd"}:
				m.TailEnd = NewCT_LineEndProperties()
				if err := d.DecodeElement(m.TailEnd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_LineProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LineProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LineProperties and its children
func (m *CT_LineProperties) Validate() error {
	return m.ValidateWithPath("CT_LineProperties")
}

// ValidateWithPath validates the CT_LineProperties and its children, prefixing error messages with path
func (m *CT_LineProperties) ValidateWithPath(path string) error {
	if m.WAttr != nil {
		if *m.WAttr < 0 {
			return fmt.Errorf("%s/m.WAttr must be >= 0 (have %v)", path, *m.WAttr)
		}
		if *m.WAttr > 20116800 {
			return fmt.Errorf("%s/m.WAttr must be <= 20116800 (have %v)", path, *m.WAttr)
		}
	}
	if err := m.CapAttr.ValidateWithPath(path + "/CapAttr"); err != nil {
		return err
	}
	if err := m.CmpdAttr.ValidateWithPath(path + "/CmpdAttr"); err != nil {
		return err
	}
	if err := m.AlgnAttr.ValidateWithPath(path + "/AlgnAttr"); err != nil {
		return err
	}
	if m.NoFill != nil {
		if err := m.NoFill.ValidateWithPath(path + "/NoFill"); err != nil {
			return err
		}
	}
	if m.SolidFill != nil {
		if err := m.SolidFill.ValidateWithPath(path + "/SolidFill"); err != nil {
			return err
		}
	}
	if m.GradFill != nil {
		if err := m.GradFill.ValidateWithPath(path + "/GradFill"); err != nil {
			return err
		}
	}
	if m.PattFill != nil {
		if err := m.PattFill.ValidateWithPath(path + "/PattFill"); err != nil {
			return err
		}
	}
	if m.PrstDash != nil {
		if err := m.PrstDash.ValidateWithPath(path + "/PrstDash"); err != nil {
			return err
		}
	}
	if m.CustDash != nil {
		if err := m.CustDash.ValidateWithPath(path + "/CustDash"); err != nil {
			return err
		}
	}
	if m.Round != nil {
		if err := m.Round.ValidateWithPath(path + "/Round"); err != nil {
			return err
		}
	}
	if m.Bevel != nil {
		if err := m.Bevel.ValidateWithPath(path + "/Bevel"); err != nil {
			return err
		}
	}
	if m.Miter != nil {
		if err := m.Miter.ValidateWithPath(path + "/Miter"); err != nil {
			return err
		}
	}
	if m.HeadEnd != nil {
		if err := m.HeadEnd.ValidateWithPath(path + "/HeadEnd"); err != nil {
			return err
		}
	}
	if m.TailEnd != nil {
		if err := m.TailEnd.ValidateWithPath(path + "/TailEnd"); err != nil {
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
