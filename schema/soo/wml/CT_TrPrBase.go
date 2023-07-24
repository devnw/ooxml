package wml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_TrPrBase struct {
	// Table Row Conditional Formatting
	CnfStyle []*CT_Cnf
	// Associated HTML div ID
	DivId []*CT_DecimalNumber
	// Grid Columns Before First Cell
	GridBefore []*CT_DecimalNumber
	// Grid Columns After Last Cell
	GridAfter []*CT_DecimalNumber
	// Preferred Width Before Table Row
	WBefore []*CT_TblWidth
	// Preferred Width After Table Row
	WAfter []*CT_TblWidth
	// Table Row Cannot Break Across Pages
	CantSplit []*CT_OnOff
	// Table Row Height
	TrHeight []*CT_Height
	// Repeat Table Row on Every New Page
	TblHeader []*CT_OnOff
	// Table Row Cell Spacing
	TblCellSpacing []*CT_TblWidth
	// Table Row Alignment
	Jc []*CT_JcTable
	// Hidden Table Row Marker
	Hidden []*CT_OnOff
}

func NewCT_TrPrBase() *CT_TrPrBase {
	ret := &CT_TrPrBase{}
	return ret
}

func (m *CT_TrPrBase) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.CnfStyle != nil {
		secnfStyle := xml.StartElement{Name: xml.Name{Local: "w:cnfStyle"}}
		for _, c := range m.CnfStyle {
			e.EncodeElement(c, secnfStyle)
		}
	}
	if m.DivId != nil {
		sedivId := xml.StartElement{Name: xml.Name{Local: "w:divId"}}
		for _, c := range m.DivId {
			e.EncodeElement(c, sedivId)
		}
	}
	if m.GridBefore != nil {
		segridBefore := xml.StartElement{Name: xml.Name{Local: "w:gridBefore"}}
		for _, c := range m.GridBefore {
			e.EncodeElement(c, segridBefore)
		}
	}
	if m.GridAfter != nil {
		segridAfter := xml.StartElement{Name: xml.Name{Local: "w:gridAfter"}}
		for _, c := range m.GridAfter {
			e.EncodeElement(c, segridAfter)
		}
	}
	if m.WBefore != nil {
		sewBefore := xml.StartElement{Name: xml.Name{Local: "w:wBefore"}}
		for _, c := range m.WBefore {
			e.EncodeElement(c, sewBefore)
		}
	}
	if m.WAfter != nil {
		sewAfter := xml.StartElement{Name: xml.Name{Local: "w:wAfter"}}
		for _, c := range m.WAfter {
			e.EncodeElement(c, sewAfter)
		}
	}
	if m.CantSplit != nil {
		secantSplit := xml.StartElement{Name: xml.Name{Local: "w:cantSplit"}}
		for _, c := range m.CantSplit {
			e.EncodeElement(c, secantSplit)
		}
	}
	if m.TrHeight != nil {
		setrHeight := xml.StartElement{Name: xml.Name{Local: "w:trHeight"}}
		for _, c := range m.TrHeight {
			e.EncodeElement(c, setrHeight)
		}
	}
	if m.TblHeader != nil {
		setblHeader := xml.StartElement{Name: xml.Name{Local: "w:tblHeader"}}
		for _, c := range m.TblHeader {
			e.EncodeElement(c, setblHeader)
		}
	}
	if m.TblCellSpacing != nil {
		setblCellSpacing := xml.StartElement{Name: xml.Name{Local: "w:tblCellSpacing"}}
		for _, c := range m.TblCellSpacing {
			e.EncodeElement(c, setblCellSpacing)
		}
	}
	if m.Jc != nil {
		sejc := xml.StartElement{Name: xml.Name{Local: "w:jc"}}
		for _, c := range m.Jc {
			e.EncodeElement(c, sejc)
		}
	}
	if m.Hidden != nil {
		sehidden := xml.StartElement{Name: xml.Name{Local: "w:hidden"}}
		for _, c := range m.Hidden {
			e.EncodeElement(c, sehidden)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TrPrBase) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TrPrBase:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "cnfStyle"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "cnfStyle"}:
				tmp := NewCT_Cnf()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CnfStyle = append(m.CnfStyle, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "divId"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "divId"}:
				tmp := NewCT_DecimalNumber()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.DivId = append(m.DivId, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "gridBefore"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "gridBefore"}:
				tmp := NewCT_DecimalNumber()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GridBefore = append(m.GridBefore, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "gridAfter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "gridAfter"}:
				tmp := NewCT_DecimalNumber()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.GridAfter = append(m.GridAfter, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "wBefore"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "wBefore"}:
				tmp := NewCT_TblWidth()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.WBefore = append(m.WBefore, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "wAfter"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "wAfter"}:
				tmp := NewCT_TblWidth()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.WAfter = append(m.WAfter, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "cantSplit"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "cantSplit"}:
				tmp := NewCT_OnOff()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.CantSplit = append(m.CantSplit, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "trHeight"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "trHeight"}:
				tmp := NewCT_Height()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TrHeight = append(m.TrHeight, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tblHeader"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "tblHeader"}:
				tmp := NewCT_OnOff()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TblHeader = append(m.TblHeader, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "tblCellSpacing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "tblCellSpacing"}:
				tmp := NewCT_TblWidth()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.TblCellSpacing = append(m.TblCellSpacing, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "jc"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "jc"}:
				tmp := NewCT_JcTable()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Jc = append(m.Jc, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "hidden"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "hidden"}:
				tmp := NewCT_OnOff()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.Hidden = append(m.Hidden, tmp)
			default:
				office.Log("skipping unsupported element on CT_TrPrBase %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TrPrBase
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TrPrBase and its children
func (m *CT_TrPrBase) Validate() error {
	return m.ValidateWithPath("CT_TrPrBase")
}

// ValidateWithPath validates the CT_TrPrBase and its children, prefixing error messages with path
func (m *CT_TrPrBase) ValidateWithPath(path string) error {
	for i, v := range m.CnfStyle {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CnfStyle[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.DivId {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/DivId[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GridBefore {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GridBefore[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.GridAfter {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/GridAfter[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.WBefore {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/WBefore[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.WAfter {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/WAfter[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.CantSplit {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/CantSplit[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.TrHeight {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TrHeight[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.TblHeader {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TblHeader[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.TblCellSpacing {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/TblCellSpacing[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Jc {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Jc[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.Hidden {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/Hidden[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
