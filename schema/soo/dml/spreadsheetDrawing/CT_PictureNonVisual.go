package spreadsheetDrawing

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_PictureNonVisual struct {
	CNvPr    *dml.CT_NonVisualDrawingProps
	CNvPicPr *dml.CT_NonVisualPictureProperties
}

func NewCT_PictureNonVisual() *CT_PictureNonVisual {
	ret := &CT_PictureNonVisual{}
	ret.CNvPr = dml.NewCT_NonVisualDrawingProps()
	ret.CNvPicPr = dml.NewCT_NonVisualPictureProperties()
	return ret
}

func (m *CT_PictureNonVisual) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	secNvPr := xml.StartElement{Name: xml.Name{Local: "xdr:cNvPr"}}
	e.EncodeElement(m.CNvPr, secNvPr)
	secNvPicPr := xml.StartElement{Name: xml.Name{Local: "xdr:cNvPicPr"}}
	e.EncodeElement(m.CNvPicPr, secNvPicPr)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_PictureNonVisual) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CNvPr = dml.NewCT_NonVisualDrawingProps()
	m.CNvPicPr = dml.NewCT_NonVisualPictureProperties()
lCT_PictureNonVisual:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cNvPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "cNvPr"}:
				if err := d.DecodeElement(m.CNvPr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cNvPicPr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "cNvPicPr"}:
				if err := d.DecodeElement(m.CNvPicPr, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_PictureNonVisual %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_PictureNonVisual
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_PictureNonVisual and its children
func (m *CT_PictureNonVisual) Validate() error {
	return m.ValidateWithPath("CT_PictureNonVisual")
}

// ValidateWithPath validates the CT_PictureNonVisual and its children, prefixing error messages with path
func (m *CT_PictureNonVisual) ValidateWithPath(path string) error {
	if err := m.CNvPr.ValidateWithPath(path + "/CNvPr"); err != nil {
		return err
	}
	if err := m.CNvPicPr.ValidateWithPath(path + "/CNvPicPr"); err != nil {
		return err
	}
	return nil
}
