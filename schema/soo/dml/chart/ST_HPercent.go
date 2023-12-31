package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_HPercent is a union type
type ST_HPercent struct {
	ST_HPercentWithSymbol *string
	ST_HPercentUShort     *uint16
}

func (m *ST_HPercent) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HPercent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_HPercentWithSymbol != nil {
		e.EncodeToken(xml.CharData(*m.ST_HPercentWithSymbol))
	}
	if m.ST_HPercentUShort != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_HPercentUShort)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_HPercent) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_HPercentWithSymbol != nil {
		mems = append(mems, "ST_HPercentWithSymbol")
	}
	if m.ST_HPercentUShort != nil {
		mems = append(mems, "ST_HPercentUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_HPercent) String() string {
	if m.ST_HPercentWithSymbol != nil {
		return fmt.Sprintf("%v", *m.ST_HPercentWithSymbol)
	}
	if m.ST_HPercentUShort != nil {
		return fmt.Sprintf("%v", *m.ST_HPercentUShort)
	}
	return ""
}
