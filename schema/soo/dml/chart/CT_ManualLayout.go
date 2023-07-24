package chart

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_ManualLayout struct {
	LayoutTarget *CT_LayoutTarget
	XMode        *CT_LayoutMode
	YMode        *CT_LayoutMode
	WMode        *CT_LayoutMode
	HMode        *CT_LayoutMode
	X            *CT_Double
	Y            *CT_Double
	W            *CT_Double
	H            *CT_Double
	ExtLst       *CT_ExtensionList
}

func NewCT_ManualLayout() *CT_ManualLayout {
	ret := &CT_ManualLayout{}
	return ret
}

func (m *CT_ManualLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.LayoutTarget != nil {
		selayoutTarget := xml.StartElement{Name: xml.Name{Local: "c:layoutTarget"}}
		e.EncodeElement(m.LayoutTarget, selayoutTarget)
	}
	if m.XMode != nil {
		sexMode := xml.StartElement{Name: xml.Name{Local: "c:xMode"}}
		e.EncodeElement(m.XMode, sexMode)
	}
	if m.YMode != nil {
		seyMode := xml.StartElement{Name: xml.Name{Local: "c:yMode"}}
		e.EncodeElement(m.YMode, seyMode)
	}
	if m.WMode != nil {
		sewMode := xml.StartElement{Name: xml.Name{Local: "c:wMode"}}
		e.EncodeElement(m.WMode, sewMode)
	}
	if m.HMode != nil {
		sehMode := xml.StartElement{Name: xml.Name{Local: "c:hMode"}}
		e.EncodeElement(m.HMode, sehMode)
	}
	if m.X != nil {
		sex := xml.StartElement{Name: xml.Name{Local: "c:x"}}
		e.EncodeElement(m.X, sex)
	}
	if m.Y != nil {
		sey := xml.StartElement{Name: xml.Name{Local: "c:y"}}
		e.EncodeElement(m.Y, sey)
	}
	if m.W != nil {
		sew := xml.StartElement{Name: xml.Name{Local: "c:w"}}
		e.EncodeElement(m.W, sew)
	}
	if m.H != nil {
		seh := xml.StartElement{Name: xml.Name{Local: "c:h"}}
		e.EncodeElement(m.H, seh)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ManualLayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_ManualLayout:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "layoutTarget"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "layoutTarget"}:
				m.LayoutTarget = NewCT_LayoutTarget()
				if err := d.DecodeElement(m.LayoutTarget, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "xMode"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "xMode"}:
				m.XMode = NewCT_LayoutMode()
				if err := d.DecodeElement(m.XMode, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "yMode"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "yMode"}:
				m.YMode = NewCT_LayoutMode()
				if err := d.DecodeElement(m.YMode, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "wMode"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "wMode"}:
				m.WMode = NewCT_LayoutMode()
				if err := d.DecodeElement(m.WMode, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "hMode"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "hMode"}:
				m.HMode = NewCT_LayoutMode()
				if err := d.DecodeElement(m.HMode, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "x"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "x"}:
				m.X = NewCT_Double()
				if err := d.DecodeElement(m.X, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "y"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "y"}:
				m.Y = NewCT_Double()
				if err := d.DecodeElement(m.Y, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "w"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "w"}:
				m.W = NewCT_Double()
				if err := d.DecodeElement(m.W, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "h"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "h"}:
				m.H = NewCT_Double()
				if err := d.DecodeElement(m.H, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_ManualLayout %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ManualLayout
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ManualLayout and its children
func (m *CT_ManualLayout) Validate() error {
	return m.ValidateWithPath("CT_ManualLayout")
}

// ValidateWithPath validates the CT_ManualLayout and its children, prefixing error messages with path
func (m *CT_ManualLayout) ValidateWithPath(path string) error {
	if m.LayoutTarget != nil {
		if err := m.LayoutTarget.ValidateWithPath(path + "/LayoutTarget"); err != nil {
			return err
		}
	}
	if m.XMode != nil {
		if err := m.XMode.ValidateWithPath(path + "/XMode"); err != nil {
			return err
		}
	}
	if m.YMode != nil {
		if err := m.YMode.ValidateWithPath(path + "/YMode"); err != nil {
			return err
		}
	}
	if m.WMode != nil {
		if err := m.WMode.ValidateWithPath(path + "/WMode"); err != nil {
			return err
		}
	}
	if m.HMode != nil {
		if err := m.HMode.ValidateWithPath(path + "/HMode"); err != nil {
			return err
		}
	}
	if m.X != nil {
		if err := m.X.ValidateWithPath(path + "/X"); err != nil {
			return err
		}
	}
	if m.Y != nil {
		if err := m.Y.ValidateWithPath(path + "/Y"); err != nil {
			return err
		}
	}
	if m.W != nil {
		if err := m.W.ValidateWithPath(path + "/W"); err != nil {
			return err
		}
	}
	if m.H != nil {
		if err := m.H.ValidateWithPath(path + "/H"); err != nil {
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
