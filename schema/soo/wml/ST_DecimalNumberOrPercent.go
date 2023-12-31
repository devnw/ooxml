package wml

import (
	"encoding/xml"
	"fmt"
)

// ST_DecimalNumberOrPercent is a union type
type ST_DecimalNumberOrPercent struct {
	ST_UnqualifiedPercentage *int64
	ST_Percentage            *string
}

func (m *ST_DecimalNumberOrPercent) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DecimalNumberOrPercent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_UnqualifiedPercentage != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_UnqualifiedPercentage)))
	}
	if m.ST_Percentage != nil {
		e.EncodeToken(xml.CharData(*m.ST_Percentage))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_DecimalNumberOrPercent) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_UnqualifiedPercentage != nil {
		mems = append(mems, "ST_UnqualifiedPercentage")
	}
	if m.ST_Percentage != nil {
		mems = append(mems, "ST_Percentage")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_DecimalNumberOrPercent) String() string {
	if m.ST_UnqualifiedPercentage != nil {
		return fmt.Sprintf("%v", *m.ST_UnqualifiedPercentage)
	}
	if m.ST_Percentage != nil {
		return fmt.Sprintf("%v", *m.ST_Percentage)
	}
	return ""
}
