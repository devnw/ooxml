package dml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type CT_LightRig struct {
	RigAttr ST_LightRigType
	DirAttr ST_LightRigDirection
	Rot     *CT_SphereCoords
}

func NewCT_LightRig() *CT_LightRig {
	ret := &CT_LightRig{}
	ret.RigAttr = ST_LightRigType(1)
	ret.DirAttr = ST_LightRigDirection(1)
	return ret
}

func (m *CT_LightRig) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	attr, err := m.RigAttr.MarshalXMLAttr(xml.Name{Local: "rig"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	attr, err = m.DirAttr.MarshalXMLAttr(xml.Name{Local: "dir"})
	if err != nil {
		return err
	}
	start.Attr = append(start.Attr, attr)
	e.EncodeToken(start)
	if m.Rot != nil {
		serot := xml.StartElement{Name: xml.Name{Local: "a:rot"}}
		e.EncodeElement(m.Rot, serot)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_LightRig) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.RigAttr = ST_LightRigType(1)
	m.DirAttr = ST_LightRigDirection(1)
	for _, attr := range start.Attr {
		if attr.Name.Local == "rig" {
			m.RigAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "dir" {
			m.DirAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
lCT_LightRig:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "rot"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "rot"}:
				m.Rot = NewCT_SphereCoords()
				if err := d.DecodeElement(m.Rot, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_LightRig %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_LightRig
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_LightRig and its children
func (m *CT_LightRig) Validate() error {
	return m.ValidateWithPath("CT_LightRig")
}

// ValidateWithPath validates the CT_LightRig and its children, prefixing error messages with path
func (m *CT_LightRig) ValidateWithPath(path string) error {
	if m.RigAttr == ST_LightRigTypeUnset {
		return fmt.Errorf("%s/RigAttr is a mandatory field", path)
	}
	if err := m.RigAttr.ValidateWithPath(path + "/RigAttr"); err != nil {
		return err
	}
	if m.DirAttr == ST_LightRigDirectionUnset {
		return fmt.Errorf("%s/DirAttr is a mandatory field", path)
	}
	if err := m.DirAttr.ValidateWithPath(path + "/DirAttr"); err != nil {
		return err
	}
	if m.Rot != nil {
		if err := m.Rot.ValidateWithPath(path + "/Rot"); err != nil {
			return err
		}
	}
	return nil
}
