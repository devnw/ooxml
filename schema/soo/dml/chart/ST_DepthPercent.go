package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_DepthPercent is a union type
type ST_DepthPercent struct {
	ST_DepthPercentWithSymbol *string
	ST_DepthPercentUShort     *uint16
}

func (m *ST_DepthPercent) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DepthPercent) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_DepthPercentWithSymbol != nil {
		e.EncodeToken(xml.CharData(*m.ST_DepthPercentWithSymbol))
	}
	if m.ST_DepthPercentUShort != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_DepthPercentUShort)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_DepthPercent) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_DepthPercentWithSymbol != nil {
		mems = append(mems, "ST_DepthPercentWithSymbol")
	}
	if m.ST_DepthPercentUShort != nil {
		mems = append(mems, "ST_DepthPercentUShort")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_DepthPercent) String() string {
	if m.ST_DepthPercentWithSymbol != nil {
		return fmt.Sprintf("%v", *m.ST_DepthPercentWithSymbol)
	}
	if m.ST_DepthPercentUShort != nil {
		return fmt.Sprintf("%v", *m.ST_DepthPercentUShort)
	}
	return ""
}
