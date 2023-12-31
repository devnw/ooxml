package diagram

import (
	"encoding/xml"
	"fmt"
)

// ST_ModelId is a union type
type ST_ModelId struct {
	Int32   *int32
	ST_Guid *string
}

func (m *ST_ModelId) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ModelId) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Int32 != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.Int32)))
	}
	if m.ST_Guid != nil {
		e.EncodeToken(xml.CharData(*m.ST_Guid))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_ModelId) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Int32 != nil {
		mems = append(mems, "Int32")
	}
	if m.ST_Guid != nil {
		mems = append(mems, "ST_Guid")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_ModelId) String() string {
	if m.Int32 != nil {
		return fmt.Sprintf("%v", *m.Int32)
	}
	if m.ST_Guid != nil {
		return fmt.Sprintf("%v", *m.ST_Guid)
	}
	return ""
}
