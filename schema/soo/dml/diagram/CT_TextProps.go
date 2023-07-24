package diagram

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_TextProps struct {
	Sp3d   *dml.CT_Shape3D
	FlatTx *dml.CT_FlatText
}

func NewCT_TextProps() *CT_TextProps {
	ret := &CT_TextProps{}
	return ret
}

func (m *CT_TextProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Sp3d != nil {
		sesp3d := xml.StartElement{Name: xml.Name{Local: "sp3d"}}
		e.EncodeElement(m.Sp3d, sesp3d)
	}
	if m.FlatTx != nil {
		seflatTx := xml.StartElement{Name: xml.Name{Local: "flatTx"}}
		e.EncodeElement(m.FlatTx, seflatTx)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_TextProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_TextProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "sp3d"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "sp3d"}:
				m.Sp3d = dml.NewCT_Shape3D()
				if err := d.DecodeElement(m.Sp3d, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "flatTx"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "flatTx"}:
				m.FlatTx = dml.NewCT_FlatText()
				if err := d.DecodeElement(m.FlatTx, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_TextProps %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_TextProps
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_TextProps and its children
func (m *CT_TextProps) Validate() error {
	return m.ValidateWithPath("CT_TextProps")
}

// ValidateWithPath validates the CT_TextProps and its children, prefixing error messages with path
func (m *CT_TextProps) ValidateWithPath(path string) error {
	if m.Sp3d != nil {
		if err := m.Sp3d.ValidateWithPath(path + "/Sp3d"); err != nil {
			return err
		}
	}
	if m.FlatTx != nil {
		if err := m.FlatTx.ValidateWithPath(path + "/FlatTx"); err != nil {
			return err
		}
	}
	return nil
}
