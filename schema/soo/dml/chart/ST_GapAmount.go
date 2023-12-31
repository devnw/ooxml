package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_GapAmount is a union type
type ST_GapAmount struct {
	ST_GapAmountPercent *string
	ST_GapAmountUShort  *uint16
}

func (m *ST_GapAmount) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_GapAmount) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_GapAmountPercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_GapAmountPercent))
	}
	if m.ST_GapAmountUShort != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_GapAmountUShort)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_GapAmount) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_GapAmountPercent != nil {
		mems = append(mems, "ST_GapAmountPercent")
	}
	if m.ST_GapAmountUShort != nil {
		mems = append(mems, "ST_GapAmountUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_GapAmount) String() string {
	if m.ST_GapAmountPercent != nil {
		return fmt.Sprintf("%v", *m.ST_GapAmountPercent)
	}
	if m.ST_GapAmountUShort != nil {
		return fmt.Sprintf("%v", *m.ST_GapAmountUShort)
	}
	return ""
}
