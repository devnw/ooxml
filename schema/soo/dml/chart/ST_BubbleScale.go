package chart

import (
	"encoding/xml"
	"fmt"
)

// ST_BubbleScale is a union type
type ST_BubbleScale struct {
	ST_BubbleScalePercent *string
	ST_BubbleScaleUInt    *uint32
}

func (m *ST_BubbleScale) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BubbleScale) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_BubbleScalePercent != nil {
		e.EncodeToken(xml.CharData(*m.ST_BubbleScalePercent))
	}
	if m.ST_BubbleScaleUInt != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_BubbleScaleUInt)))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_BubbleScale) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_BubbleScalePercent != nil {
		mems = append(mems, "ST_BubbleScalePercent")
	}
	if m.ST_BubbleScaleUInt != nil {
		mems = append(mems, "ST_BubbleScaleUInt")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_BubbleScale) String() string {
	if m.ST_BubbleScalePercent != nil {
		return fmt.Sprintf("%v", *m.ST_BubbleScalePercent)
	}
	if m.ST_BubbleScaleUInt != nil {
		return fmt.Sprintf("%v", *m.ST_BubbleScaleUInt)
	}
	return ""
}
