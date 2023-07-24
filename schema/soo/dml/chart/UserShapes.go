package chart

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml/schema/soo/dml/chartDrawing"
)

type UserShapes struct {
	chartDrawing.CT_Drawing
}

func NewUserShapes() *UserShapes {
	ret := &UserShapes{}
	ret.CT_Drawing = *chartDrawing.NewCT_Drawing()
	return ret
}

func (m *UserShapes) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/chart"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:a"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/main"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:c"}, Value: "http://schemas.openxmlformats.org/drawingml/2006/chart"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "c:userShapes"
	return m.CT_Drawing.MarshalXML(e, start)
}

func (m *UserShapes) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Drawing = *chartDrawing.NewCT_Drawing()
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing UserShapes: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the UserShapes and its children
func (m *UserShapes) Validate() error {
	return m.ValidateWithPath("UserShapes")
}

// ValidateWithPath validates the UserShapes and its children, prefixing error messages with path
func (m *UserShapes) ValidateWithPath(path string) error {
	if err := m.CT_Drawing.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
