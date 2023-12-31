package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_AnimationDgmBuildType is a union type
type ST_AnimationDgmBuildType struct {
	ST_AnimationBuildType        ST_AnimationBuildType
	ST_AnimationDgmOnlyBuildType ST_AnimationDgmOnlyBuildType
}

func (m *ST_AnimationDgmBuildType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AnimationDgmBuildType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		e.EncodeToken(xml.CharData(m.ST_AnimationBuildType.String()))
	}
	if m.ST_AnimationDgmOnlyBuildType != ST_AnimationDgmOnlyBuildTypeUnset {
		e.EncodeToken(xml.CharData(m.ST_AnimationDgmOnlyBuildType.String()))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_AnimationDgmBuildType) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		mems = append(mems, "ST_AnimationBuildType")
	}
	if m.ST_AnimationDgmOnlyBuildType != ST_AnimationDgmOnlyBuildTypeUnset {
		mems = append(mems, "ST_AnimationDgmOnlyBuildType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_AnimationDgmBuildType) String() string {
	if m.ST_AnimationBuildType != ST_AnimationBuildTypeUnset {
		return m.ST_AnimationBuildType.String()
	}
	if m.ST_AnimationDgmOnlyBuildType != ST_AnimationDgmOnlyBuildTypeUnset {
		return m.ST_AnimationDgmOnlyBuildType.String()
	}
	return ""
}
