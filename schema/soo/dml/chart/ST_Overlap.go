package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_Overlap is a union type
type ST_Overlap struct {
	ST_OverlapPercent *string
	ST_OverlapByte    *int8
}

func (m *ST_Overlap) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Overlap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_OverlapPercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_OverlapPercent))
	}
	if m.ST_OverlapByte != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_OverlapByte)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_Overlap) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_OverlapPercent != nil {
		mems = append(mems, "ST_OverlapPercent")
	}
	if m.ST_OverlapByte != nil {
		mems = append(mems, "ST_OverlapByte")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_Overlap) String() string {
	if m.ST_OverlapPercent != nil {
		return fmt.Sprintf("%v", *m.ST_OverlapPercent)
	}
	if m.ST_OverlapByte != nil {
		return fmt.Sprintf("%v", *m.ST_OverlapByte)
	}
	return ""
}
