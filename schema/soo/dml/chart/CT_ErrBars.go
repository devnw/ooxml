package chart

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_ErrBars struct {
	ErrDir     *CT_ErrDir
	ErrBarType *CT_ErrBarType
	ErrValType *CT_ErrValType
	NoEndCap   *CT_Boolean
	Plus       *CT_NumDataSource
	Minus      *CT_NumDataSource
	Val        *CT_Double
	SpPr       *dml.CT_ShapeProperties
	ExtLst     *CT_ExtensionList
}

func NewCT_ErrBars() *CT_ErrBars {
	ret := &CT_ErrBars{}
	ret.ErrBarType = NewCT_ErrBarType()
	ret.ErrValType = NewCT_ErrValType()
	return ret
}

func (m *CT_ErrBars) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ErrDir != nil {
		seerrDir := xml.StartElement{Name: xml.Name{Local: "c:errDir"}}
		e.EncodeElement(m.ErrDir, seerrDir)
	}
	seerrBarType := xml.StartElement{Name: xml.Name{Local: "c:errBarType"}}
	e.EncodeElement(m.ErrBarType, seerrBarType)
	seerrValType := xml.StartElement{Name: xml.Name{Local: "c:errValType"}}
	e.EncodeElement(m.ErrValType, seerrValType)
	if m.NoEndCap != nil {
		senoEndCap := xml.StartElement{Name: xml.Name{Local: "c:noEndCap"}}
		e.EncodeElement(m.NoEndCap, senoEndCap)
	}
	if m.Plus != nil {
		seplus := xml.StartElement{Name: xml.Name{Local: "c:plus"}}
		e.EncodeElement(m.Plus, seplus)
	}
	if m.Minus != nil {
		seminus := xml.StartElement{Name: xml.Name{Local: "c:minus"}}
		e.EncodeElement(m.Minus, seminus)
	}
	if m.Val != nil {
		seval := xml.StartElement{Name: xml.Name{Local: "c:val"}}
		e.EncodeElement(m.Val, seval)
	}
	if m.SpPr != nil {
		sespPr := xml.StartElement{Name: xml.Name{Local: "c:spPr"}}
		e.EncodeElement(m.SpPr, sespPr)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "c:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ErrBars) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.ErrBarType = NewCT_ErrBarType()
	m.ErrValType = NewCT_ErrValType()
lCT_ErrBars:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "errDir"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "errDir"}:
				m.ErrDir = NewCT_ErrDir()
				if err := d.DecodeElement(m.ErrDir, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "errBarType"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "errBarType"}:
				if err := d.DecodeElement(m.ErrBarType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "errValType"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "errValType"}:
				if err := d.DecodeElement(m.ErrValType, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "noEndCap"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "noEndCap"}:
				m.NoEndCap = NewCT_Boolean()
				if err := d.DecodeElement(m.NoEndCap, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "plus"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "plus"}:
				m.Plus = NewCT_NumDataSource()
				if err := d.DecodeElement(m.Plus, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "minus"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "minus"}:
				m.Minus = NewCT_NumDataSource()
				if err := d.DecodeElement(m.Minus, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "val"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "val"}:
				m.Val = NewCT_Double()
				if err := d.DecodeElement(m.Val, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "spPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "spPr"}:
				m.SpPr = dml.NewCT_ShapeProperties()
				if err := d.DecodeElement(m.SpPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/chart", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/chart", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_ErrBars %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ErrBars
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ErrBars and its children
func (m *CT_ErrBars) Validate() error {
	return m.ValidateWithPath("CT_ErrBars")
}

// ValidateWithPath validates the CT_ErrBars and its children, prefixing error messages with path
func (m *CT_ErrBars) ValidateWithPath(path string) error {
	if m.ErrDir != nil {
		if err := m.ErrDir.ValidateWithPath(path + "/ErrDir"); err != nil {
			return err
		}
	}
	if err := m.ErrBarType.ValidateWithPath(path + "/ErrBarType"); err != nil {
		return err
	}
	if err := m.ErrValType.ValidateWithPath(path + "/ErrValType"); err != nil {
		return err
	}
	if m.NoEndCap != nil {
		if err := m.NoEndCap.ValidateWithPath(path + "/NoEndCap"); err != nil {
			return err
		}
	}
	if m.Plus != nil {
		if err := m.Plus.ValidateWithPath(path + "/Plus"); err != nil {
			return err
		}
	}
	if m.Minus != nil {
		if err := m.Minus.ValidateWithPath(path + "/Minus"); err != nil {
			return err
		}
	}
	if m.Val != nil {
		if err := m.Val.ValidateWithPath(path + "/Val"); err != nil {
			return err
		}
	}
	if m.SpPr != nil {
		if err := m.SpPr.ValidateWithPath(path + "/SpPr"); err != nil {
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
