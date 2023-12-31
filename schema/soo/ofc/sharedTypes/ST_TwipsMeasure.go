package sharedTypes

import (
	"encoding/xml"
	"fmt"
)

// ST_TwipsMeasure is a union type
type ST_TwipsMeasure struct {
	ST_UnsignedDecimalNumber    *uint64
	ST_PositiveUniversalMeasure *string
}

func (m *ST_TwipsMeasure) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TwipsMeasure) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_UnsignedDecimalNumber != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_UnsignedDecimalNumber)))
	}
	if m.ST_PositiveUniversalMeasure != nil {
		e.EncodeToken(xml.CharData(*m.ST_PositiveUniversalMeasure))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_TwipsMeasure) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_UnsignedDecimalNumber != nil {
		mems = append(mems, "ST_UnsignedDecimalNumber")
	}
	if m.ST_PositiveUniversalMeasure != nil {
		mems = append(mems, "ST_PositiveUniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_TwipsMeasure) String() string {
	if m.ST_UnsignedDecimalNumber != nil {
		return fmt.Sprintf("%v", *m.ST_UnsignedDecimalNumber)
	}
	if m.ST_PositiveUniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_PositiveUniversalMeasure)
	}
	return ""
}
