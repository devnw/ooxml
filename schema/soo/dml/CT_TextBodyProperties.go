package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_TextBodyProperties struct {
	RotAttr              *int32
	SpcFirstLastParaAttr *bool
	VertOverflowAttr     ST_TextVertOverflowType
	HorzOverflowAttr     ST_TextHorzOverflowType
	VertAttr             ST_TextVerticalType
	WrapAttr             ST_TextWrappingType
	LInsAttr             *ST_Coordinate32
	TInsAttr             *ST_Coordinate32
	RInsAttr             *ST_Coordinate32
	BInsAttr             *ST_Coordinate32
	NumColAttr           *int32
	SpcColAttr           *int32
	RtlColAttr           *bool
	FromWordArtAttr      *bool
	AnchorAttr           ST_TextAnchoringType
	AnchorCtrAttr        *bool
	ForceAAAttr          *bool
	UprightAttr          *bool
	CompatLnSpcAttr      *bool
	PrstTxWarp           *CT_PresetTextShape
	NoAutofit            *CT_TextNoAutofit
	NormAutofit          *CT_TextNormalAutofit
	SpAutoFit            *CT_TextShapeAutofit
	Scene3d              *CT_Scene3D
	Sp3d                 *CT_Shape3D
	FlatTx               *CT_FlatText
	ExtLst               *CT_OfficeArtExtensionList
}

func NewCT_TextBodyProperties() *CT_TextBodyProperties {
	ret := &CT_TextBodyProperties{}
	return ret
}

func (m *CT_TextBodyProperties) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.RotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rot"},
			Value: fmt.Sprintf("%v", *m.RotAttr)})
	}
	if m.SpcFirstLastParaAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spcFirstLastPara"},
			Value: fmt.Sprintf("%d", b2i(*m.SpcFirstLastParaAttr))})
	}
	if m.VertOverflowAttr != ST_TextVertOverflowTypeUnset {
		attr, err := m.VertOverflowAttr.MarshalXMLAttr(xml.Name{Local: "vertOverflow"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.HorzOverflowAttr != ST_TextHorzOverflowTypeUnset {
		attr, err := m.HorzOverflowAttr.MarshalXMLAttr(xml.Name{Local: "horzOverflow"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.VertAttr != ST_TextVerticalTypeUnset {
		attr, err := m.VertAttr.MarshalXMLAttr(xml.Name{Local: "vert"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.WrapAttr != ST_TextWrappingTypeUnset {
		attr, err := m.WrapAttr.MarshalXMLAttr(xml.Name{Local: "wrap"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.LInsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lIns"},
			Value: fmt.Sprintf("%v", *m.LInsAttr)})
	}
	if m.TInsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "tIns"},
			Value: fmt.Sprintf("%v", *m.TInsAttr)})
	}
	if m.RInsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rIns"},
			Value: fmt.Sprintf("%v", *m.RInsAttr)})
	}
	if m.BInsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bIns"},
			Value: fmt.Sprintf("%v", *m.BInsAttr)})
	}
	if m.NumColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "numCol"},
			Value: fmt.Sprintf("%v", *m.NumColAttr)})
	}
	if m.SpcColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "spcCol"},
			Value: fmt.Sprintf("%v", *m.SpcColAttr)})
	}
	if m.RtlColAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rtlCol"},
			Value: fmt.Sprintf("%d", b2i(*m.RtlColAttr))})
	}
	if m.FromWordArtAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "fromWordArt"},
			Value: fmt.Sprintf("%d", b2i(*m.FromWordArtAttr))})
	}
	if m.AnchorAttr != ST_TextAnchoringTypeUnset {
		attr, err := m.AnchorAttr.MarshalXMLAttr(xml.Name{Local: "anchor"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.AnchorCtrAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "anchorCtr"},
			Value: fmt.Sprintf("%d", b2i(*m.AnchorCtrAttr))})
	}
	if m.ForceAAAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "forceAA"},
			Value: fmt.Sprintf("%d", b2i(*m.ForceAAAttr))})
	}
	if m.UprightAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "upright"},
			Value: fmt.Sprintf("%d", b2i(*m.UprightAttr))})
	}
	if m.CompatLnSpcAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "compatLnSpc"},
			Value: fmt.Sprintf("%d", b2i(*m.CompatLnSpcAttr))})
	}
	e.EncodeToken(start)
	if m.PrstTxWarp != nil {
		seprstTxWarp := xml.StartElement{Name: xml.Name{Local: "a:prstTxWarp"}}
		e.EncodeElement(m.PrstTxWarp, seprstTxWarp)
	}
	if m.NoAutofit != nil {
		senoAutofit := xml.StartElement{Name: xml.Name{Local: "a:noAutofit"}}
		e.EncodeElement(m.NoAutofit, senoAutofit)
	}
	if m.NormAutofit != nil {
		senormAutofit := xml.StartElement{Name: xml.Name{Local: "a:normAutofit"}}
		e.EncodeElement(m.NormAutofit, senormAutofit)
	}
	if m.SpAutoFit != nil {
		sespAutoFit := xml.StartElement{Name: xml.Name{Local: "a:spAutoFit"}}
		e.EncodeElement(m.SpAutoFit, sespAutoFit)
	}
	if m.Scene3d != nil {
		sescene3d := xml.StartElement{Name: xml.Name{Local: "a:scene3d"}}
		e.EncodeElement(m.Scene3d, sescene3d)
	}
	if m.Sp3d != nil {
		sesp3d := xml.StartElement{Name: xml.Name{Local: "a:sp3d"}}
		e.EncodeElement(m.Sp3d, sesp3d)
	}
	if m.FlatTx != nil {
		seflatTx := xml.StartElement{Name: xml.Name{Local: "a:flatTx"}}
		e.EncodeElement(m.FlatTx, seflatTx)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextBodyProperties) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "fromWordArt" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.FromWordArtAttr = &parsed
			continue
		}
		if attr.Name.Local == "anchor" {
			m.AnchorAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "spcFirstLastPara" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.SpcFirstLastParaAttr = &parsed
			continue
		}
		if attr.Name.Local == "anchorCtr" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AnchorCtrAttr = &parsed
			continue
		}
		if attr.Name.Local == "horzOverflow" {
			m.HorzOverflowAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "forceAA" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ForceAAAttr = &parsed
			continue
		}
		if attr.Name.Local == "wrap" {
			m.WrapAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "upright" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UprightAttr = &parsed
			continue
		}
		if attr.Name.Local == "tIns" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.TInsAttr = &parsed
			continue
		}
		if attr.Name.Local == "compatLnSpc" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.CompatLnSpcAttr = &parsed
			continue
		}
		if attr.Name.Local == "bIns" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.BInsAttr = &parsed
			continue
		}
		if attr.Name.Local == "vertOverflow" {
			m.VertOverflowAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "rot" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.RotAttr = &pt
			continue
		}
		if attr.Name.Local == "spcCol" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.SpcColAttr = &pt
			continue
		}
		if attr.Name.Local == "vert" {
			m.VertAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "rIns" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.RInsAttr = &parsed
			continue
		}
		if attr.Name.Local == "numCol" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.NumColAttr = &pt
			continue
		}
		if attr.Name.Local == "rtlCol" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.RtlColAttr = &parsed
			continue
		}
		if attr.Name.Local == "lIns" {
			parsed, err := ParseUnionST_Coordinate32(attr.Value)
			if err != nil {
				return err
			}
			m.LInsAttr = &parsed
			continue
		}
	}
lCT_TextBodyProperties:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "prstTxWarp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "prstTxWarp"}:
				m.PrstTxWarp = NewCT_PresetTextShape()
				if err := d.DecodeElement(m.PrstTxWarp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "noAutofit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "noAutofit"}:
				m.NoAutofit = NewCT_TextNoAutofit()
				if err := d.DecodeElement(m.NoAutofit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "normAutofit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "normAutofit"}:
				m.NormAutofit = NewCT_TextNormalAutofit()
				if err := d.DecodeElement(m.NormAutofit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "spAutoFit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "spAutoFit"}:
				m.SpAutoFit = NewCT_TextShapeAutofit()
				if err := d.DecodeElement(m.SpAutoFit, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "scene3d"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "scene3d"}:
				m.Scene3d = NewCT_Scene3D()
				if err := d.DecodeElement(m.Scene3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sp3d"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "sp3d"}:
				m.Sp3d = NewCT_Shape3D()
				if err := d.DecodeElement(m.Sp3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "flatTx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "flatTx"}:
				m.FlatTx = NewCT_FlatText()
				if err := d.DecodeElement(m.FlatTx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_TextBodyProperties %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextBodyProperties
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextBodyProperties and its children
func (m *CT_TextBodyProperties) Validate() error {
	return m.ValidateWithPath("CT_TextBodyProperties")
}

// ValidateWithPath validates the CT_TextBodyProperties and its children, prefixing error messages with path
func (m *CT_TextBodyProperties) ValidateWithPath(path string) error {
	if err := m.VertOverflowAttr.ValidateWithPath(path + "/VertOverflowAttr"); err != nil {
		return err
	}
	if err := m.HorzOverflowAttr.ValidateWithPath(path + "/HorzOverflowAttr"); err != nil {
		return err
	}
	if err := m.VertAttr.ValidateWithPath(path + "/VertAttr"); err != nil {
		return err
	}
	if err := m.WrapAttr.ValidateWithPath(path + "/WrapAttr"); err != nil {
		return err
	}
	if m.LInsAttr != nil {
		if err := m.LInsAttr.ValidateWithPath(path + "/LInsAttr"); err != nil {
			return err
		}
	}
	if m.TInsAttr != nil {
		if err := m.TInsAttr.ValidateWithPath(path + "/TInsAttr"); err != nil {
			return err
		}
	}
	if m.RInsAttr != nil {
		if err := m.RInsAttr.ValidateWithPath(path + "/RInsAttr"); err != nil {
			return err
		}
	}
	if m.BInsAttr != nil {
		if err := m.BInsAttr.ValidateWithPath(path + "/BInsAttr"); err != nil {
			return err
		}
	}
	if m.NumColAttr != nil {
		if *m.NumColAttr < 1 {
			return fmt.Errorf("%s/m.NumColAttr must be >= 1 (have %v)", path, *m.NumColAttr)
		}
		if *m.NumColAttr > 16 {
			return fmt.Errorf("%s/m.NumColAttr must be <= 16 (have %v)", path, *m.NumColAttr)
		}
	}
	if m.SpcColAttr != nil {
		if *m.SpcColAttr < 0 {
			return fmt.Errorf("%s/m.SpcColAttr must be >= 0 (have %v)", path, *m.SpcColAttr)
		}
	}
	if err := m.AnchorAttr.ValidateWithPath(path + "/AnchorAttr"); err != nil {
		return err
	}
	if m.PrstTxWarp != nil {
		if err := m.PrstTxWarp.ValidateWithPath(path + "/PrstTxWarp"); err != nil {
			return err
		}
	}
	if m.NoAutofit != nil {
		if err := m.NoAutofit.ValidateWithPath(path + "/NoAutofit"); err != nil {
			return err
		}
	}
	if m.NormAutofit != nil {
		if err := m.NormAutofit.ValidateWithPath(path + "/NormAutofit"); err != nil {
			return err
		}
	}
	if m.SpAutoFit != nil {
		if err := m.SpAutoFit.ValidateWithPath(path + "/SpAutoFit"); err != nil {
			return err
		}
	}
	if m.Scene3d != nil {
		if err := m.Scene3d.ValidateWithPath(path + "/Scene3d"); err != nil {
			return err
		}
	}
	if m.Sp3d != nil {
		if err := m.Sp3d.ValidateWithPath(path + "/Sp3d"); err != nil {
			return err
		}
	}
	if m.FlatTx != nil {
		if err := m.FlatTx.ValidateWithPath(path + "/FlatTx"); err != nil {
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
