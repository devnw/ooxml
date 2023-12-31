package sharedTypes

import (
	"encoding/xml"
	"fmt"
)

// ST_OnOff is a union type
type ST_OnOff struct {
	Bool      *bool
	ST_OnOff1 ST_OnOff1
}

func (m *ST_OnOff) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_OnOff) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Bool != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", b2i(*m.Bool))))
	}
	if m.ST_OnOff1 != ST_OnOff1Unset {
		e.EncodeToken(xml.CharData(m.ST_OnOff1.String()))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_OnOff) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Bool != nil {
		mems = append(mems, "Bool")
	}
	if m.ST_OnOff1 != ST_OnOff1Unset {
		mems = append(mems, "ST_OnOff1")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_OnOff) String() string {
	if m.Bool != nil {
		return fmt.Sprintf("%v", *m.Bool)
	}
	if m.ST_OnOff1 != ST_OnOff1Unset {
		return m.ST_OnOff1.String()
	}
	return ""
}
