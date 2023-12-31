package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_PositiveFixedPercentage is a union type
type ST_PositiveFixedPercentage struct {
	ST_PositiveFixedPercentageDecimal *int32
	ST_PositiveFixedPercentage        *ST_Percentage
}

func (m *ST_PositiveFixedPercentage) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PositiveFixedPercentage) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_PositiveFixedPercentageDecimal != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_PositiveFixedPercentageDecimal)))
	}
	if m.ST_PositiveFixedPercentage != nil {
		e.Encode(m.ST_PositiveFixedPercentage)
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_PositiveFixedPercentage) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_PositiveFixedPercentageDecimal != nil {
		mems = append(mems, "ST_PositiveFixedPercentageDecimal")
	}
	if m.ST_PositiveFixedPercentage != nil {
		if err := m.ST_PositiveFixedPercentage.ValidateWithPath(path + "/ST_PositiveFixedPercentage"); err != nil {
			return err
		}
		mems = append(mems, "ST_PositiveFixedPercentage")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_PositiveFixedPercentage) String() string {
	if m.ST_PositiveFixedPercentageDecimal != nil {
		return fmt.Sprintf("%v", *m.ST_PositiveFixedPercentageDecimal)
	}
	if m.ST_PositiveFixedPercentage != nil {
		return m.ST_PositiveFixedPercentage.String()
	}
	return ""
}
