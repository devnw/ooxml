package pml

import (
	"encoding/xml"
	"fmt"
)

// ST_TLTime is a union type
type ST_TLTime struct {
	Uint32              *uint32
	ST_TLTimeIndefinite ST_TLTimeIndefinite
}

func (m *ST_TLTime) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TLTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Uint32 != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.Uint32)))
	}
	if m.ST_TLTimeIndefinite != ST_TLTimeIndefiniteUnset {
		e.EncodeToken(xml.CharData(m.ST_TLTimeIndefinite.String()))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_TLTime) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Uint32 != nil {
		mems = append(mems, "Uint32")
	}
	if m.ST_TLTimeIndefinite != ST_TLTimeIndefiniteUnset {
		mems = append(mems, "ST_TLTimeIndefinite")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_TLTime) String() string {
	if m.Uint32 != nil {
		return fmt.Sprintf("%v", *m.Uint32)
	}
	if m.ST_TLTimeIndefinite != ST_TLTimeIndefiniteUnset {
		return m.ST_TLTimeIndefinite.String()
	}
	return ""
}
