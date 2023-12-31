package wml

import (
	"encoding/xml"
	"fmt"
)

// ST_HexColor is a union type
type ST_HexColor struct {
	ST_HexColorAuto ST_HexColorAuto
	ST_HexColorRGB  *string
}

func (m *ST_HexColor) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HexColor) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_HexColorAuto != ST_HexColorAutoUnset {
		e.EncodeToken(xml.CharData(m.ST_HexColorAuto.String()))
	}
	if m.ST_HexColorRGB != nil {
		e.EncodeToken(xml.CharData(*m.ST_HexColorRGB))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_HexColor) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_HexColorAuto != ST_HexColorAutoUnset {
		mems = append(mems, "ST_HexColorAuto")
	}
	if m.ST_HexColorRGB != nil {
		mems = append(mems, "ST_HexColorRGB")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_HexColor) String() string {
	if m.ST_HexColorAuto != ST_HexColorAutoUnset {
		return m.ST_HexColorAuto.String()
	}
	if m.ST_HexColorRGB != nil {
		return fmt.Sprintf("%v", *m.ST_HexColorRGB)
	}
	return ""
}
