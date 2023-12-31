package wml

import (
	"encoding/xml"
	"fmt"
)

// ST_SignedHpsMeasure is a union type
type ST_SignedHpsMeasure struct {
	Int64               *int64
	ST_UniversalMeasure *string
}

func (m *ST_SignedHpsMeasure) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SignedHpsMeasure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.Int64 != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.Int64)))
	}
	if m.ST_UniversalMeasure != nil {
		e.EncodeToken(xml.CharData(*m.ST_UniversalMeasure))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_SignedHpsMeasure) ValidateWithPath(path string) error {
	mems := []string{}
	if m.Int64 != nil {
		mems = append(mems, "Int64")
	}
	if m.ST_UniversalMeasure != nil {
		mems = append(mems, "ST_UniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_SignedHpsMeasure) String() string {
	if m.Int64 != nil {
		return fmt.Sprintf("%v", *m.Int64)
	}
	if m.ST_UniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_UniversalMeasure)
	}
	return ""
}
