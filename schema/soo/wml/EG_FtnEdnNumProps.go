package wml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type EG_FtnEdnNumProps struct {
	// Footnote and Endnote Numbering Starting Value
	NumStart *CT_DecimalNumber
	// Footnote and Endnote Numbering Restart Location
	NumRestart *CT_NumRestart
}

func NewEG_FtnEdnNumProps() *EG_FtnEdnNumProps {
	ret := &EG_FtnEdnNumProps{}
	return ret
}

func (m *EG_FtnEdnNumProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NumStart != nil {
		senumStart := xml.StartElement{Name: xml.Name{Local: "w:numStart"}}
		e.EncodeElement(m.NumStart, senumStart)
	}
	if m.NumRestart != nil {
		senumRestart := xml.StartElement{Name: xml.Name{Local: "w:numRestart"}}
		e.EncodeElement(m.NumRestart, senumRestart)
	}
	return nil
}

func (m *EG_FtnEdnNumProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lEG_FtnEdnNumProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "numStart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "numStart"}:
				m.NumStart = NewCT_DecimalNumber()
				if err := d.DecodeElement(m.NumStart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/wordprocessingml/2006/main", Local: "numRestart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/wordprocessingml/main", Local: "numRestart"}:
				m.NumRestart = NewCT_NumRestart()
				if err := d.DecodeElement(m.NumRestart, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on EG_FtnEdnNumProps %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lEG_FtnEdnNumProps
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the EG_FtnEdnNumProps and its children
func (m *EG_FtnEdnNumProps) Validate() error {
	return m.ValidateWithPath("EG_FtnEdnNumProps")
}

// ValidateWithPath validates the EG_FtnEdnNumProps and its children, prefixing error messages with path
func (m *EG_FtnEdnNumProps) ValidateWithPath(path string) error {
	if m.NumStart != nil {
		if err := m.NumStart.ValidateWithPath(path + "/NumStart"); err != nil {
			return err
		}
	}
	if m.NumRestart != nil {
		if err := m.NumRestart.ValidateWithPath(path + "/NumRestart"); err != nil {
			return err
		}
	}
	return nil
}
