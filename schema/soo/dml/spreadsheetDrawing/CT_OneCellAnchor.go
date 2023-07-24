package spreadsheetDrawing

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_OneCellAnchor struct {
	From       *CT_Marker
	Ext        *dml.CT_PositiveSize2D
	Choice     *EG_ObjectChoicesChoice
	ClientData *CT_AnchorClientData
}

func NewCT_OneCellAnchor() *CT_OneCellAnchor {
	ret := &CT_OneCellAnchor{}
	ret.From = NewCT_Marker()
	ret.Ext = dml.NewCT_PositiveSize2D()
	ret.ClientData = NewCT_AnchorClientData()
	return ret
}

func (m *CT_OneCellAnchor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	sefrom := xml.StartElement{Name: xml.Name{Local: "xdr:from"}}
	e.EncodeElement(m.From, sefrom)
	seext := xml.StartElement{Name: xml.Name{Local: "xdr:ext"}}
	e.EncodeElement(m.Ext, seext)
	if m.Choice != nil {
		m.Choice.MarshalXML(e, xml.StartElement{})
	}
	seclientData := xml.StartElement{Name: xml.Name{Local: "xdr:clientData"}}
	e.EncodeElement(m.ClientData, seclientData)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_OneCellAnchor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.From = NewCT_Marker()
	m.Ext = dml.NewCT_PositiveSize2D()
	m.ClientData = NewCT_AnchorClientData()
lCT_OneCellAnchor:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "from"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "from"}:
				if err := d.DecodeElement(m.From, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "ext"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "ext"}:
				if err := d.DecodeElement(m.Ext, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "sp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "sp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Sp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "grpSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "grpSp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GrpSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "graphicFrame"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "graphicFrame"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.GraphicFrame, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "cxnSp"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "cxnSp"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.CxnSp, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "pic"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "pic"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.Pic, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "contentPart"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "contentPart"}:
				m.Choice = NewEG_ObjectChoicesChoice()
				if err := d.DecodeElement(&m.Choice.ContentPart, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/spreadsheetDrawing", Local: "clientData"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/spreadsheetDrawing", Local: "clientData"}:
				if err := d.DecodeElement(m.ClientData, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_OneCellAnchor %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_OneCellAnchor
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_OneCellAnchor and its children
func (m *CT_OneCellAnchor) Validate() error {
	return m.ValidateWithPath("CT_OneCellAnchor")
}

// ValidateWithPath validates the CT_OneCellAnchor and its children, prefixing error messages with path
func (m *CT_OneCellAnchor) ValidateWithPath(path string) error {
	if err := m.From.ValidateWithPath(path + "/From"); err != nil {
		return err
	}
	if err := m.Ext.ValidateWithPath(path + "/Ext"); err != nil {
		return err
	}
	if m.Choice != nil {
		if err := m.Choice.ValidateWithPath(path + "/Choice"); err != nil {
			return err
		}
	}
	if err := m.ClientData.ValidateWithPath(path + "/ClientData"); err != nil {
		return err
	}
	return nil
}
