package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
)

type CT_CommentAuthor struct {
	// Comment Author ID
	IdAttr uint32
	// Comment Author Name
	NameAttr string
	// Comment Author Initials
	InitialsAttr string
	// Index of Comment Author's last comment
	LastIdxAttr uint32
	// Comment Author Color Index
	ClrIdxAttr uint32
	ExtLst     *CT_ExtensionList
}

func NewCT_CommentAuthor() *CT_CommentAuthor {
	ret := &CT_CommentAuthor{}
	return ret
}

func (m *CT_CommentAuthor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "id"},
		Value: fmt.Sprintf("%v", m.IdAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
		Value: fmt.Sprintf("%v", m.NameAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "initials"},
		Value: fmt.Sprintf("%v", m.InitialsAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "lastIdx"},
		Value: fmt.Sprintf("%v", m.LastIdxAttr)})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "clrIdx"},
		Value: fmt.Sprintf("%v", m.ClrIdxAttr)})
	e.EncodeToken(start)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_CommentAuthor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "id" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.IdAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = parsed
			continue
		}
		if attr.Name.Local == "initials" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.InitialsAttr = parsed
			continue
		}
		if attr.Name.Local == "lastIdx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.LastIdxAttr = uint32(parsed)
			continue
		}
		if attr.Name.Local == "clrIdx" {
			parsed, err := strconv.ParseUint(attr.Value, 10, 32)
			if err != nil {
				return err
			}
			m.ClrIdxAttr = uint32(parsed)
			continue
		}
	}
lCT_CommentAuthor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_CommentAuthor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_CommentAuthor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_CommentAuthor and its children
func (m *CT_CommentAuthor) Validate() error {
	return m.ValidateWithPath("CT_CommentAuthor")
}

// ValidateWithPath validates the CT_CommentAuthor and its children, prefixing error messages with path
func (m *CT_CommentAuthor) ValidateWithPath(path string) error {
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
