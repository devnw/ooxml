package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_SecondPieSize is a union type
type ST_SecondPieSize struct {
	ST_SecondPieSizePercent *string
	ST_SecondPieSizeUShort  *uint16
}

func (m *ST_SecondPieSize) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SecondPieSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_SecondPieSizePercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_SecondPieSizePercent))
	}
	if m.ST_SecondPieSizeUShort != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_SecondPieSizeUShort)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_SecondPieSize) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_SecondPieSizePercent != nil {
		mems = append(mems, "ST_SecondPieSizePercent")
	}
	if m.ST_SecondPieSizeUShort != nil {
		mems = append(mems, "ST_SecondPieSizeUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_SecondPieSize) String() string {
	if m.ST_SecondPieSizePercent != nil {
		return fmt.Sprintf("%v", *m.ST_SecondPieSizePercent)
	}
	if m.ST_SecondPieSizeUShort != nil {
		return fmt.Sprintf("%v", *m.ST_SecondPieSizeUShort)
	}
	return ""
}
