package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcLock struct {
	OfcCT_Lock
}

func NewOfcLock() *OfcLock {
	ret := &OfcLock{}
	ret.OfcCT_Lock = *NewOfcCT_Lock()
	return ret
}

func (m *OfcLock) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.OfcCT_Lock.MarshalXML(e, start)
}

func (m *OfcLock) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_Lock = *NewOfcCT_Lock()
	for _, attr := range start.Attr {
		if attr.Name.Local == "position" {
			m.PositionAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "selection" {
			m.SelectionAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "grouping" {
			m.GroupingAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ungrouping" {
			m.UngroupingAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "rotation" {
			m.RotationAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "cropping" {
			m.CroppingAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "verticies" {
			m.VerticiesAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "adjusthandles" {
			m.AdjusthandlesAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "text" {
			m.TextAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "aspectratio" {
			m.AspectratioAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "shapetype" {
			m.ShapetypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcLock: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcLock and its children
func (m *OfcLock) Validate() error {
	return m.ValidateWithPath("OfcLock")
}

// ValidateWithPath validates the OfcLock and its children, prefixing error messages with path
func (m *OfcLock) ValidateWithPath(path string) error {
	if err := m.OfcCT_Lock.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
