package pml

import (
	"encoding/xml"

	"go.devnw.com/ooxml"
)

type CT_SlideTiming struct {
	TnLst *CT_TimeNodeList
	// Build List
	BldLst *CT_BuildList
	ExtLst *CT_ExtensionListModify
}

func NewCT_SlideTiming() *CT_SlideTiming {
	ret := &CT_SlideTiming{}
	return ret
}

func (m *CT_SlideTiming) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.TnLst != nil {
		setnLst := xml.StartElement{Name: xml.Name{Local: "p:tnLst"}}
		e.EncodeElement(m.TnLst, setnLst)
	}
	if m.BldLst != nil {
		sebldLst := xml.StartElement{Name: xml.Name{Local: "p:bldLst"}}
		e.EncodeElement(m.BldLst, sebldLst)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_SlideTiming) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
lCT_SlideTiming:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "tnLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "tnLst"}:
				m.TnLst = NewCT_TimeNodeList()
				if err := d.DecodeElement(m.TnLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "bldLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "bldLst"}:
				m.BldLst = NewCT_BuildList()
				if err := d.DecodeElement(m.BldLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionListModify()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_SlideTiming %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_SlideTiming
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_SlideTiming and its children
func (m *CT_SlideTiming) Validate() error {
	return m.ValidateWithPath("CT_SlideTiming")
}

// ValidateWithPath validates the CT_SlideTiming and its children, prefixing error messages with path
func (m *CT_SlideTiming) ValidateWithPath(path string) error {
	if m.TnLst != nil {
		if err := m.TnLst.ValidateWithPath(path + "/TnLst"); err != nil {
			return err
		}
	}
	if m.BldLst != nil {
		if err := m.BldLst.ValidateWithPath(path + "/BldLst"); err != nil {
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
