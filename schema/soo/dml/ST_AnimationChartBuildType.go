package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_AnimationChartBuildType is a union type
type ST_AnimationChartBuildType struct {
	ST_AnimationBuildType          ST_AnimationBuildType
	ST_AnimationChartOnlyBuildType ST_AnimationChartOnlyBuildType
}

func (m *ST_AnimationChartBuildType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AnimationChartBuildType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		e.EncodeToken(xml.CharData(m.ST_AnimationBuildType.String()))
	}
	if m.ST_AnimationChartOnlyBuildType != ST_AnimationChartOnlyBuildTypeUnset {
		e.EncodeToken(xml.CharData(m.ST_AnimationChartOnlyBuildType.String()))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_AnimationChartBuildType) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		mems = append(mems, "ST_AnimationBuildType")
	}
	if m.ST_AnimationChartOnlyBuildType != ST_AnimationChartOnlyBuildTypeUnset {
		mems = append(mems, "ST_AnimationChartOnlyBuildType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_AnimationChartBuildType) String() string {
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		return m.ST_AnimationBuildType.String()
	}
	if m.ST_AnimationChartOnlyBuildType != ST_AnimationChartOnlyBuildTypeUnset {
		return m.ST_AnimationChartOnlyBuildType.String()
	}
	return ""
}
