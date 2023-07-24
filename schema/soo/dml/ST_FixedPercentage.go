package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_FixedPercentage is a union type
type ST_FixedPercentage struct {
	ST_FixedPercentageDecimal *int32
	ST_FixedPercentage        *ST_Percentage
}

func (m *ST_FixedPercentage) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FixedPercentage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_FixedPercentageDecimal != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_FixedPercentageDecimal)))
	}
	if m.ST_FixedPercentage != nil {
		e.Encode(m.ST_FixedPercentage)
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_FixedPercentage) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_FixedPercentageDecimal != nil {
		mems = append(mems, "ST_FixedPercentageDecimal")
	}
	if m.ST_FixedPercentage != nil {
		if err := m.ST_FixedPercentage.ValidateWithPath(path + "/ST_FixedPercentage"); err != nil {
			return err
		}
		mems = append(mems, "ST_FixedPercentage")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_FixedPercentage) String() string {
	if m.ST_FixedPercentageDecimal != nil {
		return fmt.Sprintf("%v", *m.ST_FixedPercentageDecimal)
	}
	if m.ST_FixedPercentage != nil {
		return m.ST_FixedPercentage.String()
	}
	return ""
}
