package wml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_Odso struct {
	// UDL Connection String
	Udl *CT_String
	// Data Source Table Name
	Table *CT_String
	// ODSO Data Source File Path
	Src *CT_Rel
	// Column Delimiter for Data Source
	ColDelim *CT_DecimalNumber
	// ODSO Data Source Type
	Type *CT_MailMergeSourceType
	// First Row of Data Source Contains Column Names
	FHdr *CT_OnOff
	// External Data Source to Merge Field Mapping
	FieldMapData []*CT_OdsoFieldMapData
	// Reference to Inclusion/Exclusion Data for Data Source
	RecipientData []*CT_Rel
}

func NewCT_Odso() *CT_Odso {
	ret := &CT_Odso{}
	return ret
}

func (m *CT_Odso) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Udl != nil {
		seudl := xml.StartElement{Name: xml.Name{Local: "w:udl"}}
		e.EncodeElement(m.Udl, seudl)
	}
	if m.Table != nil {
		setable := xml.StartElement{Name: xml.Name{Local: "w:table"}}
		e.EncodeElement(m.Table, setable)
	}
	if m.Src != nil {
		sesrc := xml.StartElement{Name: xml.Name{Local: "w:src"}}
		e.EncodeElement(m.Src, sesrc)
	}
	if m.ColDelim != nil {
		secolDelim := xml.StartElement{Name: xml.Name{Local: "w:colDelim"}}
		e.EncodeElement(m.ColDelim, secolDelim)
	}
	if m.Type != nil {
		setype := xml.StartElement{Name: xml.Name{Local: "w:type"}}
		e.EncodeElement(m.Type, setype)
	}
	if m.FHdr != nil {
		sefHdr := xml.StartElement{Name: xml.Name{Local: "w:fHdr"}}
		e.EncodeElement(m.FHdr, sefHdr)
	}
	if m.FieldMapData != nil {
		sefieldMapData := xml.StartElement{Name: xml.Name{Local: "w:fieldMapData"}}
		for _, c := range m.FieldMapData {
			e.EncodeElement(c, sefieldMapData)
		}
	}
	if m.RecipientData != nil {
		serecipientData := xml.StartElement{Name: xml.Name{Local: "w:recipientData"}}
		for _, c := range m.RecipientData {
			e.EncodeElement(c, serecipientData)
		}
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Odso) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_Odso:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "udl"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "udl"}:
				m.Udl = NewCT_String()
				if err := d.DecodeElement(m.Udl, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "table"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "table"}:
				m.Table = NewCT_String()
				if err := d.DecodeElement(m.Table, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "src"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "src"}:
				m.Src = NewCT_Rel()
				if err := d.DecodeElement(m.Src, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "colDelim"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "colDelim"}:
				m.ColDelim = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.ColDelim, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "type"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "type"}:
				m.Type = NewCT_MailMergeSourceType()
				if err := d.DecodeElement(m.Type, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "fHdr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "fHdr"}:
				m.FHdr = NewCT_OnOff()
				if err := d.DecodeElement(m.FHdr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "fieldMapData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "fieldMapData"}:
				tmp := NewCT_OdsoFieldMapData()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.FieldMapData = append(m.FieldMapData, tmp)
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "recipientData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "recipientData"}:
				tmp := NewCT_Rel()
				if err := d.DecodeElement(tmp, &el); err != nil {
					return err
				}
				m.RecipientData = append(m.RecipientData, tmp)
			default:
				office.Log("skipping unsupported element on CT_Odso %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Odso
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Odso and its children
func (m *CT_Odso) Validate() error {
	return m.ValidateWithPath("CT_Odso")
}

// ValidateWithPath validates the CT_Odso and its children, prefixing error messages with path
func (m *CT_Odso) ValidateWithPath(path string) error {
	if m.Udl != nil {
		if err := m.Udl.ValidateWithPath(path + "/Udl"); err != nil {
			return err
		}
	}
	if m.Table != nil {
		if err := m.Table.ValidateWithPath(path + "/Table"); err != nil {
			return err
		}
	}
	if m.Src != nil {
		if err := m.Src.ValidateWithPath(path + "/Src"); err != nil {
			return err
		}
	}
	if m.ColDelim != nil {
		if err := m.ColDelim.ValidateWithPath(path + "/ColDelim"); err != nil {
			return err
		}
	}
	if m.Type != nil {
		if err := m.Type.ValidateWithPath(path + "/Type"); err != nil {
			return err
		}
	}
	if m.FHdr != nil {
		if err := m.FHdr.ValidateWithPath(path + "/FHdr"); err != nil {
			return err
		}
	}
	for i, v := range m.FieldMapData {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/FieldMapData[%d]", path, i)); err != nil {
			return err
		}
	}
	for i, v := range m.RecipientData {
		if err := v.ValidateWithPath(fmt.Sprintf("%s/RecipientData[%d]", path, i)); err != nil {
			return err
		}
	}
	return nil
}
