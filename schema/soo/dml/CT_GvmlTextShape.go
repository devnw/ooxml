package dml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_GvmlTextShape struct {
	TxBody *CT_TextBody
	Choice *CT_GvmlTextShapeChoice
	ExtLst *CT_OfficeArtExtensionList
}

func NewCT_GvmlTextShape() *CT_GvmlTextShape {
	ret := &CT_GvmlTextShape{}
	ret.TxBody = NewCT_TextBody()
	return ret
}

func (m *CT_GvmlTextShape) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	setxBody := xml.StartElement{Name: xml.Name{Local: "a:txBody"}}
	e.EncodeElement(m.TxBody, setxBody)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "a:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_GvmlTextShape) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.TxBody = NewCT_TextBody()
lCT_GvmlTextShape:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "txBody"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "txBody"}:
				if err := d.DecodeElement(m.TxBody, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "useSpRect"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "useSpRect"}:
				m.Choice = NewCT_GvmlTextShapeChoice()
				if err := d.DecodeElement(&m.Choice.UseSpRect, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "xfrm"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "xfrm"}:
				m.Choice = NewCT_GvmlTextShapeChoice()
				if err := d.DecodeElement(&m.Choice.Xfrm, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "extLst"}:
				m.ExtLst = NewCT_OfficeArtExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_GvmlTextShape %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_GvmlTextShape
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_GvmlTextShape and its children
func (m *CT_GvmlTextShape) Validate() error {
	return m.ValidateWithPath("CT_GvmlTextShape")
}

// ValidateWithPath validates the CT_GvmlTextShape and its children, prefixing error messages with path
func (m *CT_GvmlTextShape) ValidateWithPath(path string) error {
	if err := m.TxBody.ValidateWithPath(path + "/TxBody"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
