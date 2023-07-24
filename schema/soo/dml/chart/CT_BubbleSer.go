package chart

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_BubbleSer struct {
	Idx              *CT_UnsignedInt
	Order            *CT_UnsignedInt
	Tx               *CT_SerTx
	SpPr             *dml.CT_ShapeProperties
	InvertIfNegative *CT_Boolean
	DPt              []*CT_DPt
	DLbls            *CT_DLbls
	Trendline        []*CT_Trendline
	ErrBars          []*CT_ErrBars
	XVal             *CT_AxDataSource
	YVal             *CT_NumDataSource
	BubbleSize       *CT_NumDataSource
	Bubble3D         *CT_Boolean
	ExtLst           *CT_ExtensionList
}

func NewCT_BubbleSer() *CT_BubbleSer {
	ret := &CT_BubbleSer{}
	ret.Idx = NewCT_UnsignedInt()
	ret.Order = NewCT_UnsignedInt()
	return ret
}

func (m *CT_BubbleSer) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seidx := xml.StartElement{Name: xml.Name{Local: "c:idx"}}
	e.EncodeElement(m.Idx, seidx)
	seorder := xml.StartElement{Name: xml.Name{Local: "c:order"}}
	e.EncodeElement(m.Order, seorder)
	if m.Tx != nil {
		setx := xml.StartElement{Name: xml.Name{Local: "c:tx"}}
		e.EncodeElement(m.Tx, setx)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.InvertIfNegative != nil {
		seinvertIfNegative := xml.StartElement{Name: xml.Name{Local: "c:invertIfNegative"}}
		e.EncodeElement(m.InvertIfNegative, seinvertIfNegative)
	}
	if m.DPt != nil {
		sedPt := xml.StartElement{Name: xml.Name{Local: "c:dPt"}}
		for _, c := range m.DPt {
			e.EncodeElement(c, sedPt)
		}
	}
	if m.DLbls != nil {
		sedLbls := xml.StartElement{Name: xml.Name{Local: "c:dLbls"}}
		e.EncodeElement(m.DLbls, sedLbls)
	}
	if m.Trendline != nil {
		setrendline := xml.StartElement{Name: xml.Name{Local: "c:trendline"}}
		for _, c := range m.Trendline {
			e.EncodeElement(c, setrendline)
		}
	}
	if m.ErrBars != nil {
		seerrBars := xml.StartElement{Name: xml.Name{Local: "c:errBars"}}
		for _, c := range m.ErrBars {
			e.EncodeElement(c, seerrBars)
		}
	}
	if m.XVal != nil {
		sexVal := xml.StartElement{Name: xml.Name{Local: "c:xVal"}}
		e.EncodeElement(m.XVal, sexVal)
	}
	if m.YVal != nil {
		seyVal := xml.StartElement{Name: xml.Name{Local: "c:yVal"}}
		e.EncodeElement(m.YVal, seyVal)
	}
	if m.BubbleSize != nil {
		sebubbleSize := xml.StartElement{Name: xml.Name{Local: "c:bubbleSize"}}
		e.EncodeElement(m.BubbleSize, sebubbleSize)
	}
	if m.Bubble3D != nil {
		sebubble3D := xml.StartElement{Name: xml.Name{Local: "c:bubble3D"}}
		e.EncodeElement(m.Bubble3D, sebubble3D)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_BubbleSer) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Idx = NewCT_UnsignedInt()
	m.Order = NewCT_UnsignedInt()
lCT_BubbleSer:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "idx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "idx"}:
				if err := d.DecodeElement(m.Idx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "order"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "order"}:
				if err := d.DecodeElement(m.Order, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "tx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "tx"}:
				m.Tx = NewCT_SerTx()
				if err := d.DecodeElement(m.Tx, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "invertIfNegative"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "invertIfNegative"}:
				m.InvertIfNegative = NewCT_Boolean()
				if err := d.DecodeElement(m.InvertIfNegative, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dPt"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dPt"}:
				tmp := NewCT_DPt()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DPt = append(m.DPt, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "dLbls"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "dLbls"}:
				m.DLbls = NewCT_DLbls()
				if err := d.DecodeElement(m.DLbls, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "trendline"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "trendline"}:
				tmp := NewCT_Trendline()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Trendline = append(m.Trendline, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "errBars"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "errBars"}:
				tmp := NewCT_ErrBars()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.ErrBars = append(m.ErrBars, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "xVal"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "xVal"}:
				m.XVal = NewCT_AxDataSource()
				if err := d.DecodeElement(m.XVal, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "yVal"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "yVal"}:
				m.YVal = NewCT_NumDataSource()
				if err := d.DecodeElement(m.YVal, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "bubbleSize"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "bubbleSize"}:
				m.BubbleSize = NewCT_NumDataSource()
				if err := d.DecodeElement(m.BubbleSize, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "bubble3D"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "bubble3D"}:
				m.Bubble3D = NewCT_Boolean()
				if err := d.DecodeElement(m.Bubble3D, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_BubbleSer %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_BubbleSer
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_BubbleSer and its children
func (m *CT_BubbleSer) Validate() error {
	return m.ValidateWithPath("CT_BubbleSer")
}

// ValidateWithPath validates the CT_BubbleSer and its children, prefixing error messages with path
func (m *CT_BubbleSer) ValidateWithPath(path string) error {
	if err := m.Idx.ValidateWithPath(path + "/Idx"); err != nil {
		return err
	}
	if err := m.Order.ValidateWithPath(path + "/Order"); err != nil {
		return err
	}
	if m.Tx != nil {
		if err := m.Tx.ValidateWithPath(path + "/Tx"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
			return err
		}
	}
	if m.InvertIfNegative != nil {
		if err := m.InvertIfNegative.ValidateWithPath(path + "/InvertIfNegative"); err != nil {
			return err
		}
	}
	for i, v := range m.DPt {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DPt[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.DLbls != nil {
		if err := m.DLbls.ValidateWithPath(path + "/DLbls"); err != nil {
			return err
		}
	}
	for i, v := range m.Trendline {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Trendline[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.ErrBars {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/ErrBars[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.XVal != nil {
		if err := m.XVal.ValidateWithPath(path + "/XVal"); err != nil {
			return err
		}
	}
	if m.YVal != nil {
		if err := m.YVal.ValidateWithPath(path + "/YVal"); err != nil {
			return err
		}
	}
	if m.BubbleSize != nil {
		if err := m.BubbleSize.ValidateWithPath(path + "/BubbleSize"); err != nil {
			return err
		}
	}
	if m.Bubble3D != nil {
		if err := m.Bubble3D.ValidateWithPath(path + "/Bubble3D"); err != nil {
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
