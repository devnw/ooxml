package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_Coordinate32 is a union type
type ST_Coordinate32 struct {
	ST_Coordinate32Unqualified *int32
	ST_UniversalMeasure        *string
}

func (m *ST_Coordinate32) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Coordinate32) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_Coordinate32Unqualified != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_Coordinate32Unqualified)))
	}
	if m.ST_UniversalMeasure != nil {
		e.EncodeToken(xml.CharData(*m.ST_UniversalMeasure))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_Coordinate32) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_Coordinate32Unqualified != nil {
		mems = append(mems, "ST_Coordinate32Unqualified")
	}
	if m.ST_UniversalMeasure != nil {
		mems = append(mems, "ST_UniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_Coordinate32) String() string {
	if m.ST_Coordinate32Unqualified != nil {
		return fmt.Sprintf("%v", *m.ST_Coordinate32Unqualified)
	}
	if m.ST_UniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_UniversalMeasure)
	}
	return ""
}
