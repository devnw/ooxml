package pml

import (
	"encoding/xml"
	"strconv"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type ViewPr struct {
	CT_ViewProperties
}

func NewViewPr() *ViewPr {
	ret := &ViewPr{}
	ret.CT_ViewProperties = *NewCT_ViewProperties()
	return ret
}

func (m *ViewPr) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:p"}, Value: "http://schemas.openxmlformats.org/presentationml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:sh"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "p:viewPr"
	return m.CT_ViewProperties.MarshalXML(e, start)
}

func (m *ViewPr) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_ViewProperties = *NewCT_ViewProperties()
	for _, attr := range start.Attr {
		if attr.Name.Local == "lastView" {
			m.LastViewAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "showComments" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.ShowCommentsAttr = &parsed
			continue
		}
	}
lViewPr:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "normalViewPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "normalViewPr"}:
				m.NormalViewPr = NewCT_NormalViewProperties()
				if err := d.DecodeElement(m.NormalViewPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "slideViewPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "slideViewPr"}:
				m.SlideViewPr = NewCT_SlideViewProperties()
				if err := d.DecodeElement(m.SlideViewPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "outlineViewPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "outlineViewPr"}:
				m.OutlineViewPr = NewCT_OutlineViewProperties()
				if err := d.DecodeElement(m.OutlineViewPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "notesTextViewPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "notesTextViewPr"}:
				m.NotesTextViewPr = NewCT_NotesTextViewProperties()
				if err := d.DecodeElement(m.NotesTextViewPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "sorterViewPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "sorterViewPr"}:
				m.SorterViewPr = NewCT_SlideSorterViewProperties()
				if err := d.DecodeElement(m.SorterViewPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "notesViewPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "notesViewPr"}:
				m.NotesViewPr = NewCT_NotesViewProperties()
				if err := d.DecodeElement(m.NotesViewPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "gridSpacing"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "gridSpacing"}:
				m.GridSpacing = dml.NewCT_PositiveSize2D()
				if err := d.DecodeElement(m.GridSpacing, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on ViewPr %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lViewPr
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the ViewPr and its children
func (m *ViewPr) Validate() error {
	return m.ValidateWithPath("ViewPr")
}

// ValidateWithPath validates the ViewPr and its children, prefixing error messages with path
func (m *ViewPr) ValidateWithPath(path string) error {
	if err := m.CT_ViewProperties.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
