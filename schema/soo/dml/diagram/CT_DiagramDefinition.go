package diagram

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_DiagramDefinition struct {
	UniqueIdAttr *string
	MinVerAttr   *string
	DefStyleAttr *string
	Title        []*CT_Name
	Desc         []*CT_Description
	CatLst       *CT_Categories
	SampData     *CT_SampleData
	StyleData    *CT_SampleData
	ClrData      *CT_SampleData
	LayoutNode   *CT_LayoutNode
	ExtLst       *dml.CT_OfficeArtExtensionList
}

func NewCT_DiagramDefinition() *CT_DiagramDefinition {
	ret := &CT_DiagramDefinition{}
	ret.LayoutNode = NewCT_LayoutNode()
	return ret
}

func (m *CT_DiagramDefinition) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.UniqueIdAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "uniqueId"},
			Value: fmt.Sprintf("%v", *m.UniqueIdAttr)})
	}
	if m.MinVerAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "minVer"},
			Value: fmt.Sprintf("%v", *m.MinVerAttr)})
	}
	if m.DefStyleAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "defStyle"},
			Value: fmt.Sprintf("%v", *m.DefStyleAttr)})
	}
	e.EncodeToken(start)
	if m.Title != nil {
		setitle := xml.StartElement{Name: xml.Name{Local: "title"}}
		for _, c := range m.Title {
			e.EncodeElement(c, setitle)
		}
	}
	if m.Desc != nil {
		sedesc := xml.StartElement{Name: xml.Name{Local: "desc"}}
		for _, c := range m.Desc {
			e.EncodeElement(c, sedesc)
		}
	}
	if m.CatLst != nil {
		secatLst := xml.StartElement{Name: xml.Name{Local: "catLst"}}
		e.EncodeElement(m.CatLst, secatLst)
	}
	if m.SampData != nil {
		sesampData := xml.StartElement{Name: xml.Name{Local: "sampData"}}
		e.EncodeElement(m.SampData, sesampData)
	}
	if m.StyleData != nil {
		sestyleData := xml.StartElement{Name: xml.Name{Local: "styleData"}}
		e.EncodeElement(m.StyleData, sestyleData)
	}
	if m.ClrData != nil {
		seclrData := xml.StartElement{Name: xml.Name{Local: "clrData"}}
		e.EncodeElement(m.ClrData, seclrData)
	}
	selayoutNode := xml.StartElement{Name: xml.Name{Local: "layoutNode"}}
	e.EncodeElement(m.LayoutNode, selayoutNode)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_DiagramDefinition) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.LayoutNode = NewCT_LayoutNode()
	for _, attr := range start.Attr {
		if attr.Name.Local == "uniqueId" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.UniqueIdAttr = &parsed
			continue
		}
		if attr.Name.Local == "minVer" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.MinVerAttr = &parsed
			continue
		}
		if attr.Name.Local == "defStyle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DefStyleAttr = &parsed
			continue
		}
	}
lCT_DiagramDefinition:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "title"}:
				tmp := NewCT_Name()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Title = append(m.Title, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "desc"}:
				tmp := NewCT_Description()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Desc = append(m.Desc, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "catLst"}:
				m.CatLst = NewCT_Categories()
				if err := d.DecodeElement(m.CatLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "sampData"}:
				m.SampData = NewCT_SampleData()
				if err := d.DecodeElement(m.SampData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "styleData"}:
				m.StyleData = NewCT_SampleData()
				if err := d.DecodeElement(m.StyleData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "clrData"}:
				m.ClrData = NewCT_SampleData()
				if err := d.DecodeElement(m.ClrData, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "layoutNode"}:
				if err := d.DecodeElement(m.LayoutNode, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/diagram", Local: "extLst"}:
				m.ExtLst = dml.NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_DiagramDefinition %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_DiagramDefinition
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_DiagramDefinition and its children
func (m *CT_DiagramDefinition) Validate() error {
	return m.ValidateWithPath("CT_DiagramDefinition")
}

// ValidateWithPath validates the CT_DiagramDefinition and its children, prefixing error messages with path
func (m *CT_DiagramDefinition) ValidateWithPath(path string) error {
	for i, v := range m.Title {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Title[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Desc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Desc[%d]", path, i)); err != nil {
			return err
		}
	}
	if m.CatLst != nil {
		if err := m.CatLst.ValidateWithPath(path + "/CatLst"); err != nil {
			return err
		}
	}
	if m.SampData != nil {
		if err := m.SampData.ValidateWithPath(path + "/SampData"); err != nil {
			return err
		}
	}
	if m.StyleData != nil {
		if err := m.StyleData.ValidateWithPath(path + "/StyleData"); err != nil {
			return err
		}
	}
	if m.ClrData != nil {
		if err := m.ClrData.ValidateWithPath(path + "/ClrData"); err != nil {
			return err
		}
	}
	if err := m.LayoutNode.ValidateWithPath(path + "/LayoutNode"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
