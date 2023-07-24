package diagram

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml/schema/soo/dml"
)

// ST_LayoutShapeType is a union type
type ST_LayoutShapeType struct {
	ST_ShapeType       dml.ST_ShapeType
	ST_OutputShapeType ST_OutputShapeType
}

func (m *ST_LayoutShapeType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_LayoutShapeType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_ShapeType != dml.ST_ShapeTypeUnset {
		e.EncodeToken(xml.CharData(m.ST_ShapeType.String()))
	}
	if m.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		e.EncodeToken(xml.CharData(m.ST_OutputShapeType.String()))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_LayoutShapeType) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_ShapeType != dml.ST_ShapeTypeUnset {
		mems = append(mems, "ST_ShapeType")
	}
	if m.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		mems = append(mems, "ST_OutputShapeType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_LayoutShapeType) String() string {
	if m.ST_ShapeType != dml.ST_ShapeTypeUnset {
		return m.ST_ShapeType.String()
	}
	if m.ST_OutputShapeType != ST_OutputShapeTypeUnset {
		return m.ST_OutputShapeType.String()
	}
	return ""
}
