package wml

import (
	"encoding/xml"
	"fmt"
)

// ST_TextScale is a union type
type ST_TextScale struct {
	ST_TextScalePercent *string
	ST_TextScaleDecimal *int64
}

func (m *ST_TextScale) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextScale) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_TextScalePercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_TextScalePercent))
	}
	if m.ST_TextScaleDecimal != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_TextScaleDecimal)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_TextScale) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_TextScalePercent != nil {
		mems = append(mems, "ST_TextScalePercent")
	}
	if m.ST_TextScaleDecimal != nil {
		mems = append(mems, "ST_TextScaleDecimal")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_TextScale) String() string {
	if m.ST_TextScalePercent != nil {
		return fmt.Sprintf("%v", *m.ST_TextScalePercent)
	}
	if m.ST_TextScaleDecimal != nil {
		return fmt.Sprintf("%v", *m.ST_TextScaleDecimal)
	}
	return ""
}
