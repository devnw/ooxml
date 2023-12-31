package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_Thickness is a union type
type ST_Thickness struct {
	ST_ThicknessPercent *string
	Uint32              *uint32
}

func (m *ST_Thickness) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Thickness) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_ThicknessPercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_ThicknessPercent))
	}
	if m.Uint32 != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.Uint32)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_Thickness) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_ThicknessPercent != nil {
		mems = append(mems, "ST_ThicknessPercent")
	}
	if m.Uint32 != nil {
		mems = append(mems, "Uint32")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_Thickness) String() string {
	if m.ST_ThicknessPercent != nil {
		return fmt.Sprintf("%v", *m.ST_ThicknessPercent)
	}
	if m.Uint32 != nil {
		return fmt.Sprintf("%v", *m.Uint32)
	}
	return ""
}
