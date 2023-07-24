package sml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_Comments struct {
	// Authors
	Authors *CT_Authors
	// List of Comments
	CommentList *CT_CommentList
	ExtLst      *CT_ExtensionList
}

func NewCT_Comments() *CT_Comments {
	ret := &CT_Comments{}
	ret.Authors = NewCT_Authors()
	ret.CommentList = NewCT_CommentList()
	return ret
}

func (m *CT_Comments) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	seauthors := xml.StartElement{Name: xml.Name{Local: "ma:authors"}}
	e.EncodeElement(m.Authors, seauthors)
	secommentList := xml.StartElement{Name: xml.Name{Local: "ma:commentList"}}
	e.EncodeElement(m.CommentList, secommentList)
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "ma:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_Comments) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.Authors = NewCT_Authors()
	m.CommentList = NewCT_CommentList()
lCT_Comments:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "authors"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "authors"}:
				if err := d.DecodeElement(m.Authors, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "commentList"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "commentList"}:
				if err := d.DecodeElement(m.CommentList, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/spreadsheetml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/spreadsheetml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_Comments %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_Comments
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_Comments and its children
func (m *CT_Comments) Validate() error {
	return m.ValidateWithPath("CT_Comments")
}

// ValidateWithPath validates the CT_Comments and its children, prefixing error messages with path
func (m *CT_Comments) ValidateWithPath(path string) error {
	if err := m.Authors.ValidateWithPath(path + "/Authors"); err != nil {
		return err
	}
	if err := m.CommentList.ValidateWithPath(path + "/CommentList"); err != nil {
		return err
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
