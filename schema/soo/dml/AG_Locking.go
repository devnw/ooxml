package dml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type AG_Locking struct {
	NoGrpAttr              *bool
	NoSelectAttr           *bool
	NoRotAttr              *bool
	NoChangeAspectAttr     *bool
	NoMoveAttr             *bool
	NoResizeAttr           *bool
	NoEditPointsAttr       *bool
	NoAdjustHandlesAttr    *bool
	NoChangeArrowheadsAttr *bool
	NoChangeShapeTypeAttr  *bool
}

func NewAG_Locking() *AG_Locking {
	ret := &AG_Locking{}
	return ret
}

func (m *AG_Locking) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.NoGrpAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noGrp"},
			Value: fmt.Sprintf("%d", b2i(*m.NoGrpAttr))})
	}
	if m.NoSelectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noSelect"},
			Value: fmt.Sprintf("%d", b2i(*m.NoSelectAttr))})
	}
	if m.NoRotAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noRot"},
			Value: fmt.Sprintf("%d", b2i(*m.NoRotAttr))})
	}
	if m.NoChangeAspectAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeAspect"},
			Value: fmt.Sprintf("%d", b2i(*m.NoChangeAspectAttr))})
	}
	if m.NoMoveAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noMove"},
			Value: fmt.Sprintf("%d", b2i(*m.NoMoveAttr))})
	}
	if m.NoResizeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noResize"},
			Value: fmt.Sprintf("%d", b2i(*m.NoResizeAttr))})
	}
	if m.NoEditPointsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noEditPoints"},
			Value: fmt.Sprintf("%d", b2i(*m.NoEditPointsAttr))})
	}
	if m.NoAdjustHandlesAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noAdjustHandles"},
			Value: fmt.Sprintf("%d", b2i(*m.NoAdjustHandlesAttr))})
	}
	if m.NoChangeArrowheadsAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeArrowheads"},
			Value: fmt.Sprintf("%d", b2i(*m.NoChangeArrowheadsAttr))})
	}
	if m.NoChangeShapeTypeAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "noChangeShapeType"},
			Value: fmt.Sprintf("%d", b2i(*m.NoChangeShapeTypeAttr))})
	}
	return nil
}

func (m *AG_Locking) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "noGrp" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoGrpAttr = &parsed
			continue
		}
		if attr.Name.Local == "noSelect" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoSelectAttr = &parsed
			continue
		}
		if attr.Name.Local == "noRot" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoRotAttr = &parsed
			continue
		}
		if attr.Name.Local == "noChangeAspect" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoChangeAspectAttr = &parsed
			continue
		}
		if attr.Name.Local == "noMove" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoMoveAttr = &parsed
			continue
		}
		if attr.Name.Local == "noResize" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoResizeAttr = &parsed
			continue
		}
		if attr.Name.Local == "noEditPoints" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoEditPointsAttr = &parsed
			continue
		}
		if attr.Name.Local == "noAdjustHandles" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoAdjustHandlesAttr = &parsed
			continue
		}
		if attr.Name.Local == "noChangeArrowheads" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoChangeArrowheadsAttr = &parsed
			continue
		}
		if attr.Name.Local == "noChangeShapeType" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.NoChangeShapeTypeAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing AG_Locking: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the AG_Locking and its children
func (m *AG_Locking) Validate() error {
	return m.ValidateWithPath("AG_Locking")
}

// ValidateWithPath validates the AG_Locking and its children, prefixing error messages with path
func (m *AG_Locking) ValidateWithPath(path string) error {
	return nil
}
