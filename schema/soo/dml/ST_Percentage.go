package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_Percentage is a union type
type ST_Percentage struct {
	ST_PercentageDecimal *int32
	ST_Percentage        *string
}

func (m *ST_Percentage) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Percentage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_PercentageDecimal != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_PercentageDecimal)))
	}
	if m.ST_Percentage != nil {
		e.EncodeToken(xml.CharData(*m.ST_Percentage))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_Percentage) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_PercentageDecimal != nil {
		mems = append(mems, "ST_PercentageDecimal")
	}
	if m.ST_Percentage != nil {
		mems = append(mems, "ST_Percentage")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_Percentage) String() string {
	if m.ST_PercentageDecimal != nil {
		return fmt.Sprintf("%v", *m.ST_PercentageDecimal)
	}
	if m.ST_Percentage != nil {
		return fmt.Sprintf("%v", *m.ST_Percentage)
	}
	return ""
}
