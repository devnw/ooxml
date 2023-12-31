package wml

import (
	"encoding/xml"
	"fmt"
)

// ST_MeasurementOrPercent is a union type
type ST_MeasurementOrPercent struct {
	ST_DecimalNumberOrPercent *ST_DecimalNumberOrPercent
	ST_UniversalMeasure       *string
}

func (m *ST_MeasurementOrPercent) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_MeasurementOrPercent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_DecimalNumberOrPercent != nil {
		e.Encode(m.ST_DecimalNumberOrPercent)
	}
	if m.ST_UniversalMeasure != nil {
		e.EncodeToken(xml.CharData(*m.ST_UniversalMeasure))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_MeasurementOrPercent) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_DecimalNumberOrPercent != nil {
		if err := m.ST_DecimalNumberOrPercent.ValidateWithPath(path + "/ST_DecimalNumberOrPercent"); err != nil {
			return err
		}
		mems = append(mems, "ST_DecimalNumberOrPercent")
	}
	if m.ST_UniversalMeasure != nil {
		mems = append(mems, "ST_UniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_MeasurementOrPercent) String() string {
	if m.ST_DecimalNumberOrPercent != nil {
		return m.ST_DecimalNumberOrPercent.String()
	}
	if m.ST_UniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_UniversalMeasure)
	}
	return ""
}
