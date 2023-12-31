package diagram

import (
	"encoding/xml"
	"fmt"
)

// ST_PrSetCustVal is a union type
type ST_PrSetCustVal struct {
	ST_Percentage *string
	Int32         *int32
}

func (m *ST_PrSetCustVal) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PrSetCustVal) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_Percentage != nil {
		e.EncodeToken(xml.CharData(*m.ST_Percentage))
	}
	if m.Int32 != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.Int32)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_PrSetCustVal) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_Percentage != nil {
		mems = append(mems, "ST_Percentage")
	}
	if m.Int32 != nil {
		mems = append(mems, "Int32")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_PrSetCustVal) String() string {
	if m.ST_Percentage != nil {
		return fmt.Sprintf("%v", *m.ST_Percentage)
	}
	if m.Int32 != nil {
		return fmt.Sprintf("%v", *m.Int32)
	}
	return ""
}
