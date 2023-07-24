package sml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_CfRule struct {
	// Type
	TypeAttr ST_CfType
	// Differential Formatting Id
	DxfIdAttr *uint32
	// Priority
	PriorityAttr int32
	// Stop If True
	StopIfTrueAttr *bool
	// Above Or Below Average
	AboveAverageAttr *bool
	// Top 10 Percent
	PercentAttr *bool
	// Bottom N
	BottomAttr *bool
	// Operator
	OperatorAttr ST_ConditionalFormattingOperator
	// Text
	TextAttr *string
	// Time Period
	TimePeriodAttr ST_TimePeriod
	// Rank
	RankAttr *uint32
	// StdDev
	StdDevAttr *int32
	// Equal Average
	EqualAverageAttr *bool
	// Formula
	Formula []string
	// Color Scale
	ColorScale *CT_ColorScale
	// Data Bar
	DataBar *CT_DataBar
	// Icon Set
	IconSet *CT_IconSet
	ExtLst  *CT_ExtensionList
}

func NewCT_CfRule() *CT_CfRule {
	ret := &CT_CfRule{}
	return ret
}

func (m *CT_CfRule) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.TypeAttr != ST_CfTypeUnset {
		attr, err := m.TypeAttr.MarshalXMLAttr(xml.Name{Local: "type"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.DxfIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "dxfId"},
			Value: fmt.Sprintf("%v", *m.DxfIdAttr)})
	}
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "priority"},
		Value: fmt.Sprintf("%v", m.PriorityAttr)})
	if m.StopIfTrueAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stopIfTrue"},
			Value: fmt.Sprintf("%d", b2i(*m.StopIfTrueAttr))})
	}
	if m.AboveAverageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "aboveAverage"},
			Value: fmt.Sprintf("%d", b2i(*m.AboveAverageAttr))})
	}
	if m.PercentAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "percent"},
			Value: fmt.Sprintf("%d", b2i(*m.PercentAttr))})
	}
	if m.BottomAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "bottom"},
			Value: fmt.Sprintf("%d", b2i(*m.BottomAttr))})
	}
	if m.OperatorAttr != ST_ConditionalFormattingOperatorUnset {
		attr, err := m.OperatorAttr.MarshalXMLAttr(xml.Name{Local: "operator"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.TextAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "text"},
			Value: fmt.Sprintf("%v", *m.TextAttr)})
	}
	if m.TimePeriodAttr != ST_TimePeriodUnset {
		attr, err := m.TimePeriodAttr.MarshalXMLAttr(xml.Name{Local: "timePeriod"})
		if err != nil {
			return err
		}
		start.Attr = append(start.Attr, attr)
	}
	if m.RankAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "rank"},
			Value: fmt.Sprintf("%v", *m.RankAttr)})
	}
	if m.StdDevAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "stdDev"},
			Value: fmt.Sprintf("%v", *m.StdDevAttr)})
	}
	if m.EqualAverageAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "equalAverage"},
			Value: fmt.Sprintf("%d", b2i(*m.EqualAverageAttr))})
	}
	e.EncodeToken(start)
	if m.Formula != nil {
		seformula := xml.StartElement{Name: xml.Name{Local: "ma:formula"}}
		for _, c := range m.Formula {
			e.EncodeElement(c, seformula)
		}
	}
	if m.ColorScale != nil {
		secolorScale := xml.StartElement{Name: xml.Name{Local: "ma:colorScale"}}
		e.EncodeElement(m.ColorScale, secolorScale)
	}
	if m.DataBar != nil {
		sedataBar := xml.StartElement{Name: xml.Name{Local: "ma:dataBar"}}
		e.EncodeElement(m.DataBar, sedataBar)
	}
	if m.IconSet != nil {
		seiconSet := xml.StartElement{Name: xml.Name{Local: "ma:iconSet"}}
		e.EncodeElement(m.IconSet, seiconSet)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CfRule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "timePeriod" {
			m.TimePeriodAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "rank" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.RankAttr = &pt
			continue
		}
		if attr.Name.Local == "priority" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.PriorityAttr = int32(parsed)
			continue
		}
		if attr.Name.Local == "stdDev" {
			parsed, err := strconv.ParseInt(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := int32(parsed)
			m.StdDevAttr = &pt
			continue
		}
		if attr.Name.Local == "aboveAverage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.AboveAverageAttr = &parsed
			continue
		}
		if attr.Name.Local == "bottom" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.BottomAttr = &parsed
			continue
		}
		if attr.Name.Local == "operator" {
			m.OperatorAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "text" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.TextAttr = &parsed
			continue
		}
		if attr.Name.Local == "dxfId" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			pt := uint32(parsed)
			m.DxfIdAttr = &pt
			continue
		}
		if attr.Name.Local == "stopIfTrue" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.StopIfTrueAttr = &parsed
			continue
		}
		if attr.Name.Local == "percent" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.PercentAttr = &parsed
			continue
		}
		if attr.Name.Local == "equalAverage" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.EqualAverageAttr = &parsed
			continue
		}
	}
lCT_CfRule:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "formula"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "formula"}:
				var tmp string
				if err := d.DecodeElement(&tmp, &el); err != nil {
					return err
				}
				m.Formula = append(m.Formula, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "colorScale"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "colorScale"}:
				m.ColorScale = NewCT_ColorScale()
				if err := d.DecodeElement(m.ColorScale, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "dataBar"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "dataBar"}:
				m.DataBar = NewCT_DataBar()
				if err := d.DecodeElement(m.DataBar, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "iconSet"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "iconSet"}:
				m.IconSet = NewCT_IconSet()
				if err := d.DecodeElement(m.IconSet, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_CfRule %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CfRule
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CfRule and its children
func (m *CT_CfRule) Validate() error {
	return m.ValidateWithPath("CT_CfRule")
}

// ValidateWithPath validates the CT_CfRule and its children, prefixing error messages with path
func (m *CT_CfRule) ValidateWithPath(path string) error {
	if err := m.TypeAttr.ValidateWithPath(path + "/TypeAttr"); err != nil {
		return err
	}
	if err := m.OperatorAttr.ValidateWithPath(path + "/OperatorAttr"); err != nil {
		return err
	}
	if err := m.TimePeriodAttr.ValidateWithPath(path + "/TimePeriodAttr"); err != nil {
		return err
	}
	if m.ColorScale != nil {
		if err := m.ColorScale.ValidateWithPath(path + "/ColorScale"); err != nil {
			return err
		}
	}
	if m.DataBar != nil {
		if err := m.DataBar.ValidateWithPath(path + "/DataBar"); err != nil {
			return err
		}
	}
	if m.IconSet != nil {
		if err := m.IconSet.ValidateWithPath(path + "/IconSet"); err != nil {
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
