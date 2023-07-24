package sml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type Comments struct {
	CT_Comments
}

func NewComments() *Comments {
	ret := &Comments{}
	ret.CT_Comments = *NewCT_Comments()
	return ret
}

func (m *Comments) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:ma"}, Value: "http://schemas.openxmlformats.org/spreadsheetml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xdr"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "ma:comments"
	return m.CT_Comments.MarshalXML(e, start)
}

func (m *Comments) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Comments = *NewCT_Comments()
lComments:
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
				office.Log("skipping unsupported element on Comments %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lComments
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Comments and its children
func (m *Comments) Validate() error {
	return m.ValidateWithPath("Comments")
}

// ValidateWithPath validates the Comments and its children, prefixing error messages with path
func (m *Comments) ValidateWithPath(path string) error {
	if err := m.CT_Comments.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
