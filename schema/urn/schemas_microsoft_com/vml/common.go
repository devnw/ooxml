package vml

import (
	"encoding/xml"
	"fmt"

	"go.devnw.com/ooxml"
)

type ST_Ext byte

const (
	ST_ExtUnset              ST_Ext = 0
	ST_ExtView               ST_Ext = 1
	ST_ExtEdit               ST_Ext = 2
	ST_ExtBackwardCompatible ST_Ext = 3
)

func (e ST_Ext) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ExtUnset:
		attr.Value = ""
	case ST_ExtView:
		attr.Value = "view"
	case ST_ExtEdit:
		attr.Value = "edit"
	case ST_ExtBackwardCompatible:
		attr.Value = "backwardCompatible"
	}
	return attr, nil
}

func (e *ST_Ext) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "view":
		*e = 1
	case "edit":
		*e = 2
	case "backwardCompatible":
		*e = 3
	}
	return nil
}

func (m ST_Ext) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Ext) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "view":
			*m = 1
		case "edit":
			*m = 2
		case "backwardCompatible":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_Ext) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "view"
	case 2:
		return "edit"
	case 3:
		return "backwardCompatible"
	}
	return ""
}

func (m ST_Ext) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Ext) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FillType byte

const (
	ST_FillTypeUnset          ST_FillType = 0
	ST_FillTypeSolid          ST_FillType = 1
	ST_FillTypeGradient       ST_FillType = 2
	ST_FillTypeGradientRadial ST_FillType = 3
	ST_FillTypeTile           ST_FillType = 4
	ST_FillTypePattern        ST_FillType = 5
	ST_FillTypeFrame          ST_FillType = 6
)

func (e ST_FillType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FillTypeUnset:
		attr.Value = ""
	case ST_FillTypeSolid:
		attr.Value = "solid"
	case ST_FillTypeGradient:
		attr.Value = "gradient"
	case ST_FillTypeGradientRadial:
		attr.Value = "gradientRadial"
	case ST_FillTypeTile:
		attr.Value = "tile"
	case ST_FillTypePattern:
		attr.Value = "pattern"
	case ST_FillTypeFrame:
		attr.Value = "frame"
	}
	return attr, nil
}

func (e *ST_FillType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "solid":
		*e = 1
	case "gradient":
		*e = 2
	case "gradientRadial":
		*e = 3
	case "tile":
		*e = 4
	case "pattern":
		*e = 5
	case "frame":
		*e = 6
	}
	return nil
}

func (m ST_FillType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FillType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "solid":
			*m = 1
		case "gradient":
			*m = 2
		case "gradientRadial":
			*m = 3
		case "tile":
			*m = 4
		case "pattern":
			*m = 5
		case "frame":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FillType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "solid"
	case 2:
		return "gradient"
	case 3:
		return "gradientRadial"
	case 4:
		return "tile"
	case 5:
		return "pattern"
	case 6:
		return "frame"
	}
	return ""
}

func (m ST_FillType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FillType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FillMethod byte

const (
	ST_FillMethodUnset       ST_FillMethod = 0
	ST_FillMethodNone        ST_FillMethod = 1
	ST_FillMethodLinear      ST_FillMethod = 2
	ST_FillMethodSigma       ST_FillMethod = 3
	ST_FillMethodAny         ST_FillMethod = 4
	ST_FillMethodLinearSigma ST_FillMethod = 5
)

func (e ST_FillMethod) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FillMethodUnset:
		attr.Value = ""
	case ST_FillMethodNone:
		attr.Value = "none"
	case ST_FillMethodLinear:
		attr.Value = "linear"
	case ST_FillMethodSigma:
		attr.Value = "sigma"
	case ST_FillMethodAny:
		attr.Value = "any"
	case ST_FillMethodLinearSigma:
		attr.Value = "linear sigma"
	}
	return attr, nil
}

func (e *ST_FillMethod) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "linear":
		*e = 2
	case "sigma":
		*e = 3
	case "any":
		*e = 4
	case "linear sigma":
		*e = 5
	}
	return nil
}

func (m ST_FillMethod) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FillMethod) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "linear":
			*m = 2
		case "sigma":
			*m = 3
		case "any":
			*m = 4
		case "linear sigma":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_FillMethod) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "linear"
	case 3:
		return "sigma"
	case 4:
		return "any"
	case 5:
		return "linear sigma"
	}
	return ""
}

func (m ST_FillMethod) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FillMethod) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ShadowType byte

const (
	ST_ShadowTypeUnset       ST_ShadowType = 0
	ST_ShadowTypeSingle      ST_ShadowType = 1
	ST_ShadowTypeDouble      ST_ShadowType = 2
	ST_ShadowTypeEmboss      ST_ShadowType = 3
	ST_ShadowTypePerspective ST_ShadowType = 4
)

func (e ST_ShadowType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ShadowTypeUnset:
		attr.Value = ""
	case ST_ShadowTypeSingle:
		attr.Value = "single"
	case ST_ShadowTypeDouble:
		attr.Value = "double"
	case ST_ShadowTypeEmboss:
		attr.Value = "emboss"
	case ST_ShadowTypePerspective:
		attr.Value = "perspective"
	}
	return attr, nil
}

func (e *ST_ShadowType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "single":
		*e = 1
	case "double":
		*e = 2
	case "emboss":
		*e = 3
	case "perspective":
		*e = 4
	}
	return nil
}

func (m ST_ShadowType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ShadowType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "single":
			*m = 1
		case "double":
			*m = 2
		case "emboss":
			*m = 3
		case "perspective":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ShadowType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "single"
	case 2:
		return "double"
	case 3:
		return "emboss"
	case 4:
		return "perspective"
	}
	return ""
}

func (m ST_ShadowType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ShadowType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_StrokeLineStyle byte

const (
	ST_StrokeLineStyleUnset            ST_StrokeLineStyle = 0
	ST_StrokeLineStyleSingle           ST_StrokeLineStyle = 1
	ST_StrokeLineStyleThinThin         ST_StrokeLineStyle = 2
	ST_StrokeLineStyleThinThick        ST_StrokeLineStyle = 3
	ST_StrokeLineStyleThickThin        ST_StrokeLineStyle = 4
	ST_StrokeLineStyleThickBetweenThin ST_StrokeLineStyle = 5
)

func (e ST_StrokeLineStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StrokeLineStyleUnset:
		attr.Value = ""
	case ST_StrokeLineStyleSingle:
		attr.Value = "single"
	case ST_StrokeLineStyleThinThin:
		attr.Value = "thinThin"
	case ST_StrokeLineStyleThinThick:
		attr.Value = "thinThick"
	case ST_StrokeLineStyleThickThin:
		attr.Value = "thickThin"
	case ST_StrokeLineStyleThickBetweenThin:
		attr.Value = "thickBetweenThin"
	}
	return attr, nil
}

func (e *ST_StrokeLineStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "single":
		*e = 1
	case "thinThin":
		*e = 2
	case "thinThick":
		*e = 3
	case "thickThin":
		*e = 4
	case "thickBetweenThin":
		*e = 5
	}
	return nil
}

func (m ST_StrokeLineStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_StrokeLineStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "single":
			*m = 1
		case "thinThin":
			*m = 2
		case "thinThick":
			*m = 3
		case "thickThin":
			*m = 4
		case "thickBetweenThin":
			*m = 5
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_StrokeLineStyle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "single"
	case 2:
		return "thinThin"
	case 3:
		return "thinThick"
	case 4:
		return "thickThin"
	case 5:
		return "thickBetweenThin"
	}
	return ""
}

func (m ST_StrokeLineStyle) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_StrokeLineStyle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_StrokeJoinStyle byte

const (
	ST_StrokeJoinStyleUnset ST_StrokeJoinStyle = 0
	ST_StrokeJoinStyleRound ST_StrokeJoinStyle = 1
	ST_StrokeJoinStyleBevel ST_StrokeJoinStyle = 2
	ST_StrokeJoinStyleMiter ST_StrokeJoinStyle = 3
)

func (e ST_StrokeJoinStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StrokeJoinStyleUnset:
		attr.Value = ""
	case ST_StrokeJoinStyleRound:
		attr.Value = "round"
	case ST_StrokeJoinStyleBevel:
		attr.Value = "bevel"
	case ST_StrokeJoinStyleMiter:
		attr.Value = "miter"
	}
	return attr, nil
}

func (e *ST_StrokeJoinStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "round":
		*e = 1
	case "bevel":
		*e = 2
	case "miter":
		*e = 3
	}
	return nil
}

func (m ST_StrokeJoinStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_StrokeJoinStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "round":
			*m = 1
		case "bevel":
			*m = 2
		case "miter":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_StrokeJoinStyle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "round"
	case 2:
		return "bevel"
	case 3:
		return "miter"
	}
	return ""
}

func (m ST_StrokeJoinStyle) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_StrokeJoinStyle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_StrokeEndCap byte

const (
	ST_StrokeEndCapUnset  ST_StrokeEndCap = 0
	ST_StrokeEndCapFlat   ST_StrokeEndCap = 1
	ST_StrokeEndCapSquare ST_StrokeEndCap = 2
	ST_StrokeEndCapRound  ST_StrokeEndCap = 3
)

func (e ST_StrokeEndCap) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StrokeEndCapUnset:
		attr.Value = ""
	case ST_StrokeEndCapFlat:
		attr.Value = "flat"
	case ST_StrokeEndCapSquare:
		attr.Value = "square"
	case ST_StrokeEndCapRound:
		attr.Value = "round"
	}
	return attr, nil
}

func (e *ST_StrokeEndCap) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "flat":
		*e = 1
	case "square":
		*e = 2
	case "round":
		*e = 3
	}
	return nil
}

func (m ST_StrokeEndCap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_StrokeEndCap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "flat":
			*m = 1
		case "square":
			*m = 2
		case "round":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_StrokeEndCap) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "flat"
	case 2:
		return "square"
	case 3:
		return "round"
	}
	return ""
}

func (m ST_StrokeEndCap) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_StrokeEndCap) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_StrokeArrowLength byte

const (
	ST_StrokeArrowLengthUnset  ST_StrokeArrowLength = 0
	ST_StrokeArrowLengthShort  ST_StrokeArrowLength = 1
	ST_StrokeArrowLengthMedium ST_StrokeArrowLength = 2
	ST_StrokeArrowLengthLong   ST_StrokeArrowLength = 3
)

func (e ST_StrokeArrowLength) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StrokeArrowLengthUnset:
		attr.Value = ""
	case ST_StrokeArrowLengthShort:
		attr.Value = "short"
	case ST_StrokeArrowLengthMedium:
		attr.Value = "medium"
	case ST_StrokeArrowLengthLong:
		attr.Value = "long"
	}
	return attr, nil
}

func (e *ST_StrokeArrowLength) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "short":
		*e = 1
	case "medium":
		*e = 2
	case "long":
		*e = 3
	}
	return nil
}

func (m ST_StrokeArrowLength) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_StrokeArrowLength) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "short":
			*m = 1
		case "medium":
			*m = 2
		case "long":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_StrokeArrowLength) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "short"
	case 2:
		return "medium"
	case 3:
		return "long"
	}
	return ""
}

func (m ST_StrokeArrowLength) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_StrokeArrowLength) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_StrokeArrowWidth byte

const (
	ST_StrokeArrowWidthUnset  ST_StrokeArrowWidth = 0
	ST_StrokeArrowWidthNarrow ST_StrokeArrowWidth = 1
	ST_StrokeArrowWidthMedium ST_StrokeArrowWidth = 2
	ST_StrokeArrowWidthWide   ST_StrokeArrowWidth = 3
)

func (e ST_StrokeArrowWidth) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StrokeArrowWidthUnset:
		attr.Value = ""
	case ST_StrokeArrowWidthNarrow:
		attr.Value = "narrow"
	case ST_StrokeArrowWidthMedium:
		attr.Value = "medium"
	case ST_StrokeArrowWidthWide:
		attr.Value = "wide"
	}
	return attr, nil
}

func (e *ST_StrokeArrowWidth) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "narrow":
		*e = 1
	case "medium":
		*e = 2
	case "wide":
		*e = 3
	}
	return nil
}

func (m ST_StrokeArrowWidth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_StrokeArrowWidth) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "narrow":
			*m = 1
		case "medium":
			*m = 2
		case "wide":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_StrokeArrowWidth) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "narrow"
	case 2:
		return "medium"
	case 3:
		return "wide"
	}
	return ""
}

func (m ST_StrokeArrowWidth) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_StrokeArrowWidth) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_StrokeArrowType byte

const (
	ST_StrokeArrowTypeUnset   ST_StrokeArrowType = 0
	ST_StrokeArrowTypeNone    ST_StrokeArrowType = 1
	ST_StrokeArrowTypeBlock   ST_StrokeArrowType = 2
	ST_StrokeArrowTypeClassic ST_StrokeArrowType = 3
	ST_StrokeArrowTypeOval    ST_StrokeArrowType = 4
	ST_StrokeArrowTypeDiamond ST_StrokeArrowType = 5
	ST_StrokeArrowTypeOpen    ST_StrokeArrowType = 6
)

func (e ST_StrokeArrowType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_StrokeArrowTypeUnset:
		attr.Value = ""
	case ST_StrokeArrowTypeNone:
		attr.Value = "none"
	case ST_StrokeArrowTypeBlock:
		attr.Value = "block"
	case ST_StrokeArrowTypeClassic:
		attr.Value = "classic"
	case ST_StrokeArrowTypeOval:
		attr.Value = "oval"
	case ST_StrokeArrowTypeDiamond:
		attr.Value = "diamond"
	case ST_StrokeArrowTypeOpen:
		attr.Value = "open"
	}
	return attr, nil
}

func (e *ST_StrokeArrowType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "block":
		*e = 2
	case "classic":
		*e = 3
	case "oval":
		*e = 4
	case "diamond":
		*e = 5
	case "open":
		*e = 6
	}
	return nil
}

func (m ST_StrokeArrowType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_StrokeArrowType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "block":
			*m = 2
		case "classic":
			*m = 3
		case "oval":
			*m = 4
		case "diamond":
			*m = 5
		case "open":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_StrokeArrowType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "block"
	case 3:
		return "classic"
	case 4:
		return "oval"
	case 5:
		return "diamond"
	case 6:
		return "open"
	}
	return ""
}

func (m ST_StrokeArrowType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_StrokeArrowType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ImageAspect byte

const (
	ST_ImageAspectUnset   ST_ImageAspect = 0
	ST_ImageAspectIgnore  ST_ImageAspect = 1
	ST_ImageAspectAtMost  ST_ImageAspect = 2
	ST_ImageAspectAtLeast ST_ImageAspect = 3
)

func (e ST_ImageAspect) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ImageAspectUnset:
		attr.Value = ""
	case ST_ImageAspectIgnore:
		attr.Value = "ignore"
	case ST_ImageAspectAtMost:
		attr.Value = "atMost"
	case ST_ImageAspectAtLeast:
		attr.Value = "atLeast"
	}
	return attr, nil
}

func (e *ST_ImageAspect) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "ignore":
		*e = 1
	case "atMost":
		*e = 2
	case "atLeast":
		*e = 3
	}
	return nil
}

func (m ST_ImageAspect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ImageAspect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "ignore":
			*m = 1
		case "atMost":
			*m = 2
		case "atLeast":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_ImageAspect) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "ignore"
	case 2:
		return "atMost"
	case 3:
		return "atLeast"
	}
	return ""
}

func (m ST_ImageAspect) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ImageAspect) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_EditAs byte

const (
	ST_EditAsUnset    ST_EditAs = 0
	ST_EditAsCanvas   ST_EditAs = 1
	ST_EditAsOrgchart ST_EditAs = 2
	ST_EditAsRadial   ST_EditAs = 3
	ST_EditAsCycle    ST_EditAs = 4
	ST_EditAsStacked  ST_EditAs = 5
	ST_EditAsVenn     ST_EditAs = 6
	ST_EditAsBullseye ST_EditAs = 7
)

func (e ST_EditAs) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_EditAsUnset:
		attr.Value = ""
	case ST_EditAsCanvas:
		attr.Value = "canvas"
	case ST_EditAsOrgchart:
		attr.Value = "orgchart"
	case ST_EditAsRadial:
		attr.Value = "radial"
	case ST_EditAsCycle:
		attr.Value = "cycle"
	case ST_EditAsStacked:
		attr.Value = "stacked"
	case ST_EditAsVenn:
		attr.Value = "venn"
	case ST_EditAsBullseye:
		attr.Value = "bullseye"
	}
	return attr, nil
}

func (e *ST_EditAs) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "canvas":
		*e = 1
	case "orgchart":
		*e = 2
	case "radial":
		*e = 3
	case "cycle":
		*e = 4
	case "stacked":
		*e = 5
	case "venn":
		*e = 6
	case "bullseye":
		*e = 7
	}
	return nil
}

func (m ST_EditAs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_EditAs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "canvas":
			*m = 1
		case "orgchart":
			*m = 2
		case "radial":
			*m = 3
		case "cycle":
			*m = 4
		case "stacked":
			*m = 5
		case "venn":
			*m = 6
		case "bullseye":
			*m = 7
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m ST_EditAs) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "canvas"
	case 2:
		return "orgchart"
	case 3:
		return "radial"
	case 4:
		return "cycle"
	case 5:
		return "stacked"
	case 6:
		return "venn"
	case 7:
		return "bullseye"
	}
	return ""
}

func (m ST_EditAs) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_EditAs) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Shape", NewCT_Shape)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Shapetype", NewCT_Shapetype)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Group", NewCT_Group)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Background", NewCT_Background)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Fill", NewCT_Fill)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Formulas", NewCT_Formulas)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_F", NewCT_F)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Handles", NewCT_Handles)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_H", NewCT_H)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_ImageData", NewCT_ImageData)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Path", NewCT_Path)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Shadow", NewCT_Shadow)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Stroke", NewCT_Stroke)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Textbox", NewCT_Textbox)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_TextPath", NewCT_TextPath)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Arc", NewCT_Arc)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Curve", NewCT_Curve)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Image", NewCT_Image)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Line", NewCT_Line)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Oval", NewCT_Oval)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_PolyLine", NewCT_PolyLine)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_Rect", NewCT_Rect)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "CT_RoundRect", NewCT_RoundRect)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "shape", NewShape)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "shapetype", NewShapetype)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "group", NewGroup)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "background", NewBackground)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "fill", NewFill)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "formulas", NewFormulas)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "handles", NewHandles)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "imagedata", NewImagedata)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "path", NewPath)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "textbox", NewTextbox)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "shadow", NewShadow)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "stroke", NewStroke)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "textpath", NewTextpath)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "arc", NewArc)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "curve", NewCurve)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "image", NewImage)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "line", NewLine)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "oval", NewOval)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "polyline", NewPolyline)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "rect", NewRect)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "roundrect", NewRoundrect)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "EG_ShapeElements", NewEG_ShapeElements)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_Id", NewAG_Id)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_Style", NewAG_Style)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_Type", NewAG_Type)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_Adj", NewAG_Adj)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_Path", NewAG_Path)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_Fill", NewAG_Fill)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_Chromakey", NewAG_Chromakey)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_Ext", NewAG_Ext)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_CoreAttributes", NewAG_CoreAttributes)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_ShapeAttributes", NewAG_ShapeAttributes)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_OfficeCoreAttributes", NewAG_OfficeCoreAttributes)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_OfficeShapeAttributes", NewAG_OfficeShapeAttributes)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_AllCoreAttributes", NewAG_AllCoreAttributes)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_AllShapeAttributes", NewAG_AllShapeAttributes)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_ImageAttributes", NewAG_ImageAttributes)
	office.RegisterConstructor("urn:schemas-microsoft-com:vml", "AG_StrokeAttributes", NewAG_StrokeAttributes)
}

type OfcST_RType byte

const (
	OfcST_RTypeUnset     OfcST_RType = 0
	OfcST_RTypeArc       OfcST_RType = 1
	OfcST_RTypeCallout   OfcST_RType = 2
	OfcST_RTypeConnector OfcST_RType = 3
	OfcST_RTypeAlign     OfcST_RType = 4
)

func (e OfcST_RType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_RTypeUnset:
		attr.Value = ""
	case OfcST_RTypeArc:
		attr.Value = "arc"
	case OfcST_RTypeCallout:
		attr.Value = "callout"
	case OfcST_RTypeConnector:
		attr.Value = "connector"
	case OfcST_RTypeAlign:
		attr.Value = "align"
	}
	return attr, nil
}

func (e *OfcST_RType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "arc":
		*e = 1
	case "callout":
		*e = 2
	case "connector":
		*e = 3
	case "align":
		*e = 4
	}
	return nil
}

func (m OfcST_RType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_RType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "arc":
			*m = 1
		case "callout":
			*m = 2
		case "connector":
			*m = 3
		case "align":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_RType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "arc"
	case 2:
		return "callout"
	case 3:
		return "connector"
	case 4:
		return "align"
	}
	return ""
}

func (m OfcST_RType) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_RType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_How byte

const (
	OfcST_HowUnset  OfcST_How = 0
	OfcST_HowTop    OfcST_How = 1
	OfcST_HowMiddle OfcST_How = 2
	OfcST_HowBottom OfcST_How = 3
	OfcST_HowLeft   OfcST_How = 4
	OfcST_HowCenter OfcST_How = 5
	OfcST_HowRight  OfcST_How = 6
)

func (e OfcST_How) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_HowUnset:
		attr.Value = ""
	case OfcST_HowTop:
		attr.Value = "top"
	case OfcST_HowMiddle:
		attr.Value = "middle"
	case OfcST_HowBottom:
		attr.Value = "bottom"
	case OfcST_HowLeft:
		attr.Value = "left"
	case OfcST_HowCenter:
		attr.Value = "center"
	case OfcST_HowRight:
		attr.Value = "right"
	}
	return attr, nil
}

func (e *OfcST_How) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "top":
		*e = 1
	case "middle":
		*e = 2
	case "bottom":
		*e = 3
	case "left":
		*e = 4
	case "center":
		*e = 5
	case "right":
		*e = 6
	}
	return nil
}

func (m OfcST_How) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_How) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "top":
			*m = 1
		case "middle":
			*m = 2
		case "bottom":
			*m = 3
		case "left":
			*m = 4
		case "center":
			*m = 5
		case "right":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_How) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "top"
	case 2:
		return "middle"
	case 3:
		return "bottom"
	case 4:
		return "left"
	case 5:
		return "center"
	case 6:
		return "right"
	}
	return ""
}

func (m OfcST_How) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_How) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_BWMode byte

const (
	OfcST_BWModeUnset             OfcST_BWMode = 0
	OfcST_BWModeColor             OfcST_BWMode = 1
	OfcST_BWModeAuto              OfcST_BWMode = 2
	OfcST_BWModeGrayScale         OfcST_BWMode = 3
	OfcST_BWModeLightGrayscale    OfcST_BWMode = 4
	OfcST_BWModeInverseGray       OfcST_BWMode = 5
	OfcST_BWModeGrayOutline       OfcST_BWMode = 6
	OfcST_BWModeHighContrast      OfcST_BWMode = 7
	OfcST_BWModeBlack             OfcST_BWMode = 8
	OfcST_BWModeWhite             OfcST_BWMode = 9
	OfcST_BWModeHide              OfcST_BWMode = 10
	OfcST_BWModeUndrawn           OfcST_BWMode = 11
	OfcST_BWModeBlackTextAndLines OfcST_BWMode = 12
)

func (e OfcST_BWMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_BWModeUnset:
		attr.Value = ""
	case OfcST_BWModeColor:
		attr.Value = "color"
	case OfcST_BWModeAuto:
		attr.Value = "auto"
	case OfcST_BWModeGrayScale:
		attr.Value = "grayScale"
	case OfcST_BWModeLightGrayscale:
		attr.Value = "lightGrayscale"
	case OfcST_BWModeInverseGray:
		attr.Value = "inverseGray"
	case OfcST_BWModeGrayOutline:
		attr.Value = "grayOutline"
	case OfcST_BWModeHighContrast:
		attr.Value = "highContrast"
	case OfcST_BWModeBlack:
		attr.Value = "black"
	case OfcST_BWModeWhite:
		attr.Value = "white"
	case OfcST_BWModeHide:
		attr.Value = "hide"
	case OfcST_BWModeUndrawn:
		attr.Value = "undrawn"
	case OfcST_BWModeBlackTextAndLines:
		attr.Value = "blackTextAndLines"
	}
	return attr, nil
}

func (e *OfcST_BWMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "color":
		*e = 1
	case "auto":
		*e = 2
	case "grayScale":
		*e = 3
	case "lightGrayscale":
		*e = 4
	case "inverseGray":
		*e = 5
	case "grayOutline":
		*e = 6
	case "highContrast":
		*e = 7
	case "black":
		*e = 8
	case "white":
		*e = 9
	case "hide":
		*e = 10
	case "undrawn":
		*e = 11
	case "blackTextAndLines":
		*e = 12
	}
	return nil
}

func (m OfcST_BWMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_BWMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "color":
			*m = 1
		case "auto":
			*m = 2
		case "grayScale":
			*m = 3
		case "lightGrayscale":
			*m = 4
		case "inverseGray":
			*m = 5
		case "grayOutline":
			*m = 6
		case "highContrast":
			*m = 7
		case "black":
			*m = 8
		case "white":
			*m = 9
		case "hide":
			*m = 10
		case "undrawn":
			*m = 11
		case "blackTextAndLines":
			*m = 12
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_BWMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "color"
	case 2:
		return "auto"
	case 3:
		return "grayScale"
	case 4:
		return "lightGrayscale"
	case 5:
		return "inverseGray"
	case 6:
		return "grayOutline"
	case 7:
		return "highContrast"
	case 8:
		return "black"
	case 9:
		return "white"
	case 10:
		return "hide"
	case 11:
		return "undrawn"
	case 12:
		return "blackTextAndLines"
	}
	return ""
}

func (m OfcST_BWMode) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_BWMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_ScreenSize byte

const (
	OfcST_ScreenSizeUnset    OfcST_ScreenSize = 0
	OfcST_ScreenSize544x376  OfcST_ScreenSize = 1
	OfcST_ScreenSize640x480  OfcST_ScreenSize = 2
	OfcST_ScreenSize720x512  OfcST_ScreenSize = 3
	OfcST_ScreenSize800x600  OfcST_ScreenSize = 4
	OfcST_ScreenSize1024x768 OfcST_ScreenSize = 5
	OfcST_ScreenSize1152x862 OfcST_ScreenSize = 6
)

func (e OfcST_ScreenSize) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_ScreenSizeUnset:
		attr.Value = ""
	case OfcST_ScreenSize544x376:
		attr.Value = "544,376"
	case OfcST_ScreenSize640x480:
		attr.Value = "640,480"
	case OfcST_ScreenSize720x512:
		attr.Value = "720,512"
	case OfcST_ScreenSize800x600:
		attr.Value = "800,600"
	case OfcST_ScreenSize1024x768:
		attr.Value = "1024,768"
	case OfcST_ScreenSize1152x862:
		attr.Value = "1152,862"
	}
	return attr, nil
}

func (e *OfcST_ScreenSize) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "544,376":
		*e = 1
	case "640,480":
		*e = 2
	case "720,512":
		*e = 3
	case "800,600":
		*e = 4
	case "1024,768":
		*e = 5
	case "1152,862":
		*e = 6
	}
	return nil
}

func (m OfcST_ScreenSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_ScreenSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "544,376":
			*m = 1
		case "640,480":
			*m = 2
		case "720,512":
			*m = 3
		case "800,600":
			*m = 4
		case "1024,768":
			*m = 5
		case "1152,862":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_ScreenSize) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "544,376"
	case 2:
		return "640,480"
	case 3:
		return "720,512"
	case 4:
		return "800,600"
	case 5:
		return "1024,768"
	case 6:
		return "1152,862"
	}
	return ""
}

func (m OfcST_ScreenSize) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_ScreenSize) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_InsetMode byte

const (
	OfcST_InsetModeUnset  OfcST_InsetMode = 0
	OfcST_InsetModeAuto   OfcST_InsetMode = 1
	OfcST_InsetModeCustom OfcST_InsetMode = 2
)

func (e OfcST_InsetMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_InsetModeUnset:
		attr.Value = ""
	case OfcST_InsetModeAuto:
		attr.Value = "auto"
	case OfcST_InsetModeCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *OfcST_InsetMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	case "custom":
		*e = 2
	}
	return nil
}

func (m OfcST_InsetMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_InsetMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "auto":
			*m = 1
		case "custom":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_InsetMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	case 2:
		return "custom"
	}
	return ""
}

func (m OfcST_InsetMode) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_InsetMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_ColorMode byte

const (
	OfcST_ColorModeUnset  OfcST_ColorMode = 0
	OfcST_ColorModeAuto   OfcST_ColorMode = 1
	OfcST_ColorModeCustom OfcST_ColorMode = 2
)

func (e OfcST_ColorMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_ColorModeUnset:
		attr.Value = ""
	case OfcST_ColorModeAuto:
		attr.Value = "auto"
	case OfcST_ColorModeCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *OfcST_ColorMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "auto":
		*e = 1
	case "custom":
		*e = 2
	}
	return nil
}

func (m OfcST_ColorMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_ColorMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "auto":
			*m = 1
		case "custom":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_ColorMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "auto"
	case 2:
		return "custom"
	}
	return ""
}

func (m OfcST_ColorMode) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_ColorMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_DiagramLayout byte

const (
	OfcST_DiagramLayoutUnset OfcST_DiagramLayout = 0
	OfcST_DiagramLayout0     OfcST_DiagramLayout = 1
	OfcST_DiagramLayout1     OfcST_DiagramLayout = 2
	OfcST_DiagramLayout2     OfcST_DiagramLayout = 3
	OfcST_DiagramLayout3     OfcST_DiagramLayout = 4
)

func (e OfcST_DiagramLayout) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_DiagramLayoutUnset:
		attr.Value = ""
	case OfcST_DiagramLayout0:
		attr.Value = "0"
	case OfcST_DiagramLayout1:
		attr.Value = "1"
	case OfcST_DiagramLayout2:
		attr.Value = "2"
	case OfcST_DiagramLayout3:
		attr.Value = "3"
	}
	return attr, nil
}

func (e *OfcST_DiagramLayout) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "0":
		*e = 1
	case "1":
		*e = 2
	case "2":
		*e = 3
	case "3":
		*e = 4
	}
	return nil
}

func (m OfcST_DiagramLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_DiagramLayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "0":
			*m = 1
		case "1":
			*m = 2
		case "2":
			*m = 3
		case "3":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_DiagramLayout) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "0"
	case 2:
		return "1"
	case 3:
		return "2"
	case 4:
		return "3"
	}
	return ""
}

func (m OfcST_DiagramLayout) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_DiagramLayout) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_ExtrusionType byte

const (
	OfcST_ExtrusionTypeUnset       OfcST_ExtrusionType = 0
	OfcST_ExtrusionTypePerspective OfcST_ExtrusionType = 1
	OfcST_ExtrusionTypeParallel    OfcST_ExtrusionType = 2
)

func (e OfcST_ExtrusionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_ExtrusionTypeUnset:
		attr.Value = ""
	case OfcST_ExtrusionTypePerspective:
		attr.Value = "perspective"
	case OfcST_ExtrusionTypeParallel:
		attr.Value = "parallel"
	}
	return attr, nil
}

func (e *OfcST_ExtrusionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "perspective":
		*e = 1
	case "parallel":
		*e = 2
	}
	return nil
}

func (m OfcST_ExtrusionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_ExtrusionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "perspective":
			*m = 1
		case "parallel":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_ExtrusionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "perspective"
	case 2:
		return "parallel"
	}
	return ""
}

func (m OfcST_ExtrusionType) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_ExtrusionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_ExtrusionRender byte

const (
	OfcST_ExtrusionRenderUnset        OfcST_ExtrusionRender = 0
	OfcST_ExtrusionRenderSolid        OfcST_ExtrusionRender = 1
	OfcST_ExtrusionRenderWireFrame    OfcST_ExtrusionRender = 2
	OfcST_ExtrusionRenderBoundingCube OfcST_ExtrusionRender = 3
)

func (e OfcST_ExtrusionRender) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_ExtrusionRenderUnset:
		attr.Value = ""
	case OfcST_ExtrusionRenderSolid:
		attr.Value = "solid"
	case OfcST_ExtrusionRenderWireFrame:
		attr.Value = "wireFrame"
	case OfcST_ExtrusionRenderBoundingCube:
		attr.Value = "boundingCube"
	}
	return attr, nil
}

func (e *OfcST_ExtrusionRender) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "solid":
		*e = 1
	case "wireFrame":
		*e = 2
	case "boundingCube":
		*e = 3
	}
	return nil
}

func (m OfcST_ExtrusionRender) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_ExtrusionRender) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "solid":
			*m = 1
		case "wireFrame":
			*m = 2
		case "boundingCube":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_ExtrusionRender) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "solid"
	case 2:
		return "wireFrame"
	case 3:
		return "boundingCube"
	}
	return ""
}

func (m OfcST_ExtrusionRender) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_ExtrusionRender) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_ExtrusionPlane byte

const (
	OfcST_ExtrusionPlaneUnset OfcST_ExtrusionPlane = 0
	OfcST_ExtrusionPlaneXY    OfcST_ExtrusionPlane = 1
	OfcST_ExtrusionPlaneZX    OfcST_ExtrusionPlane = 2
	OfcST_ExtrusionPlaneYZ    OfcST_ExtrusionPlane = 3
)

func (e OfcST_ExtrusionPlane) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_ExtrusionPlaneUnset:
		attr.Value = ""
	case OfcST_ExtrusionPlaneXY:
		attr.Value = "XY"
	case OfcST_ExtrusionPlaneZX:
		attr.Value = "ZX"
	case OfcST_ExtrusionPlaneYZ:
		attr.Value = "YZ"
	}
	return attr, nil
}

func (e *OfcST_ExtrusionPlane) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "XY":
		*e = 1
	case "ZX":
		*e = 2
	case "YZ":
		*e = 3
	}
	return nil
}

func (m OfcST_ExtrusionPlane) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_ExtrusionPlane) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "XY":
			*m = 1
		case "ZX":
			*m = 2
		case "YZ":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_ExtrusionPlane) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "XY"
	case 2:
		return "ZX"
	case 3:
		return "YZ"
	}
	return ""
}

func (m OfcST_ExtrusionPlane) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_ExtrusionPlane) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_Angle byte

const (
	OfcST_AngleUnset OfcST_Angle = 0
	OfcST_AngleAny   OfcST_Angle = 1
	OfcST_Angle30    OfcST_Angle = 2
	OfcST_Angle45    OfcST_Angle = 3
	OfcST_Angle60    OfcST_Angle = 4
	OfcST_Angle90    OfcST_Angle = 5
	OfcST_AngleAuto  OfcST_Angle = 6
)

func (e OfcST_Angle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_AngleUnset:
		attr.Value = ""
	case OfcST_AngleAny:
		attr.Value = "any"
	case OfcST_Angle30:
		attr.Value = "30"
	case OfcST_Angle45:
		attr.Value = "45"
	case OfcST_Angle60:
		attr.Value = "60"
	case OfcST_Angle90:
		attr.Value = "90"
	case OfcST_AngleAuto:
		attr.Value = "auto"
	}
	return attr, nil
}

func (e *OfcST_Angle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "any":
		*e = 1
	case "30":
		*e = 2
	case "45":
		*e = 3
	case "60":
		*e = 4
	case "90":
		*e = 5
	case "auto":
		*e = 6
	}
	return nil
}

func (m OfcST_Angle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_Angle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "any":
			*m = 1
		case "30":
			*m = 2
		case "45":
			*m = 3
		case "60":
			*m = 4
		case "90":
			*m = 5
		case "auto":
			*m = 6
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_Angle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "any"
	case 2:
		return "30"
	case 3:
		return "45"
	case 4:
		return "60"
	case 5:
		return "90"
	case 6:
		return "auto"
	}
	return ""
}

func (m OfcST_Angle) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_Angle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_CalloutPlacement byte

const (
	OfcST_CalloutPlacementUnset  OfcST_CalloutPlacement = 0
	OfcST_CalloutPlacementTop    OfcST_CalloutPlacement = 1
	OfcST_CalloutPlacementCenter OfcST_CalloutPlacement = 2
	OfcST_CalloutPlacementBottom OfcST_CalloutPlacement = 3
	OfcST_CalloutPlacementUser   OfcST_CalloutPlacement = 4
)

func (e OfcST_CalloutPlacement) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_CalloutPlacementUnset:
		attr.Value = ""
	case OfcST_CalloutPlacementTop:
		attr.Value = "top"
	case OfcST_CalloutPlacementCenter:
		attr.Value = "center"
	case OfcST_CalloutPlacementBottom:
		attr.Value = "bottom"
	case OfcST_CalloutPlacementUser:
		attr.Value = "user"
	}
	return attr, nil
}

func (e *OfcST_CalloutPlacement) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "top":
		*e = 1
	case "center":
		*e = 2
	case "bottom":
		*e = 3
	case "user":
		*e = 4
	}
	return nil
}

func (m OfcST_CalloutPlacement) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_CalloutPlacement) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "top":
			*m = 1
		case "center":
			*m = 2
		case "bottom":
			*m = 3
		case "user":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_CalloutPlacement) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "top"
	case 2:
		return "center"
	case 3:
		return "bottom"
	case 4:
		return "user"
	}
	return ""
}

func (m OfcST_CalloutPlacement) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_CalloutPlacement) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_ConnectorType byte

const (
	OfcST_ConnectorTypeUnset    OfcST_ConnectorType = 0
	OfcST_ConnectorTypeNone     OfcST_ConnectorType = 1
	OfcST_ConnectorTypeStraight OfcST_ConnectorType = 2
	OfcST_ConnectorTypeElbow    OfcST_ConnectorType = 3
	OfcST_ConnectorTypeCurved   OfcST_ConnectorType = 4
)

func (e OfcST_ConnectorType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_ConnectorTypeUnset:
		attr.Value = ""
	case OfcST_ConnectorTypeNone:
		attr.Value = "none"
	case OfcST_ConnectorTypeStraight:
		attr.Value = "straight"
	case OfcST_ConnectorTypeElbow:
		attr.Value = "elbow"
	case OfcST_ConnectorTypeCurved:
		attr.Value = "curved"
	}
	return attr, nil
}

func (e *OfcST_ConnectorType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "straight":
		*e = 2
	case "elbow":
		*e = 3
	case "curved":
		*e = 4
	}
	return nil
}

func (m OfcST_ConnectorType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_ConnectorType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "straight":
			*m = 2
		case "elbow":
			*m = 3
		case "curved":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_ConnectorType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "straight"
	case 3:
		return "elbow"
	case 4:
		return "curved"
	}
	return ""
}

func (m OfcST_ConnectorType) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_ConnectorType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_HrAlign byte

const (
	OfcST_HrAlignUnset  OfcST_HrAlign = 0
	OfcST_HrAlignLeft   OfcST_HrAlign = 1
	OfcST_HrAlignRight  OfcST_HrAlign = 2
	OfcST_HrAlignCenter OfcST_HrAlign = 3
)

func (e OfcST_HrAlign) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_HrAlignUnset:
		attr.Value = ""
	case OfcST_HrAlignLeft:
		attr.Value = "left"
	case OfcST_HrAlignRight:
		attr.Value = "right"
	case OfcST_HrAlignCenter:
		attr.Value = "center"
	}
	return attr, nil
}

func (e *OfcST_HrAlign) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "left":
		*e = 1
	case "right":
		*e = 2
	case "center":
		*e = 3
	}
	return nil
}

func (m OfcST_HrAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_HrAlign) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "left":
			*m = 1
		case "right":
			*m = 2
		case "center":
			*m = 3
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_HrAlign) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "left"
	case 2:
		return "right"
	case 3:
		return "center"
	}
	return ""
}

func (m OfcST_HrAlign) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_HrAlign) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_ConnectType byte

const (
	OfcST_ConnectTypeUnset    OfcST_ConnectType = 0
	OfcST_ConnectTypeNone     OfcST_ConnectType = 1
	OfcST_ConnectTypeRect     OfcST_ConnectType = 2
	OfcST_ConnectTypeSegments OfcST_ConnectType = 3
	OfcST_ConnectTypeCustom   OfcST_ConnectType = 4
)

func (e OfcST_ConnectType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_ConnectTypeUnset:
		attr.Value = ""
	case OfcST_ConnectTypeNone:
		attr.Value = "none"
	case OfcST_ConnectTypeRect:
		attr.Value = "rect"
	case OfcST_ConnectTypeSegments:
		attr.Value = "segments"
	case OfcST_ConnectTypeCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *OfcST_ConnectType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "rect":
		*e = 2
	case "segments":
		*e = 3
	case "custom":
		*e = 4
	}
	return nil
}

func (m OfcST_ConnectType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_ConnectType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "none":
			*m = 1
		case "rect":
			*m = 2
		case "segments":
			*m = 3
		case "custom":
			*m = 4
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_ConnectType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "rect"
	case 3:
		return "segments"
	case 4:
		return "custom"
	}
	return ""
}

func (m OfcST_ConnectType) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_ConnectType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_OLEType byte

const (
	OfcST_OLETypeUnset OfcST_OLEType = 0
	OfcST_OLETypeEmbed OfcST_OLEType = 1
	OfcST_OLETypeLink  OfcST_OLEType = 2
)

func (e OfcST_OLEType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_OLETypeUnset:
		attr.Value = ""
	case OfcST_OLETypeEmbed:
		attr.Value = "Embed"
	case OfcST_OLETypeLink:
		attr.Value = "Link"
	}
	return attr, nil
}

func (e *OfcST_OLEType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "Embed":
		*e = 1
	case "Link":
		*e = 2
	}
	return nil
}

func (m OfcST_OLEType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_OLEType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "Embed":
			*m = 1
		case "Link":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_OLEType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "Embed"
	case 2:
		return "Link"
	}
	return ""
}

func (m OfcST_OLEType) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_OLEType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_OLEDrawAspect byte

const (
	OfcST_OLEDrawAspectUnset   OfcST_OLEDrawAspect = 0
	OfcST_OLEDrawAspectContent OfcST_OLEDrawAspect = 1
	OfcST_OLEDrawAspectIcon    OfcST_OLEDrawAspect = 2
)

func (e OfcST_OLEDrawAspect) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_OLEDrawAspectUnset:
		attr.Value = ""
	case OfcST_OLEDrawAspectContent:
		attr.Value = "Content"
	case OfcST_OLEDrawAspectIcon:
		attr.Value = "Icon"
	}
	return attr, nil
}

func (e *OfcST_OLEDrawAspect) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "Content":
		*e = 1
	case "Icon":
		*e = 2
	}
	return nil
}

func (m OfcST_OLEDrawAspect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_OLEDrawAspect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "Content":
			*m = 1
		case "Icon":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_OLEDrawAspect) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "Content"
	case 2:
		return "Icon"
	}
	return ""
}

func (m OfcST_OLEDrawAspect) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_OLEDrawAspect) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_OLEUpdateMode byte

const (
	OfcST_OLEUpdateModeUnset  OfcST_OLEUpdateMode = 0
	OfcST_OLEUpdateModeAlways OfcST_OLEUpdateMode = 1
	OfcST_OLEUpdateModeOnCall OfcST_OLEUpdateMode = 2
)

func (e OfcST_OLEUpdateMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_OLEUpdateModeUnset:
		attr.Value = ""
	case OfcST_OLEUpdateModeAlways:
		attr.Value = "Always"
	case OfcST_OLEUpdateModeOnCall:
		attr.Value = "OnCall"
	}
	return attr, nil
}

func (e *OfcST_OLEUpdateMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "Always":
		*e = 1
	case "OnCall":
		*e = 2
	}
	return nil
}

func (m OfcST_OLEUpdateMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_OLEUpdateMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "Always":
			*m = 1
		case "OnCall":
			*m = 2
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_OLEUpdateMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "Always"
	case 2:
		return "OnCall"
	}
	return ""
}

func (m OfcST_OLEUpdateMode) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_OLEUpdateMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type OfcST_FillType byte

const (
	OfcST_FillTypeUnset            OfcST_FillType = 0
	OfcST_FillTypeGradientCenter   OfcST_FillType = 1
	OfcST_FillTypeSolid            OfcST_FillType = 2
	OfcST_FillTypePattern          OfcST_FillType = 3
	OfcST_FillTypeTile             OfcST_FillType = 4
	OfcST_FillTypeFrame            OfcST_FillType = 5
	OfcST_FillTypeGradientUnscaled OfcST_FillType = 6
	OfcST_FillTypeGradientRadial   OfcST_FillType = 7
	OfcST_FillTypeGradient         OfcST_FillType = 8
	OfcST_FillTypeBackground       OfcST_FillType = 9
)

func (e OfcST_FillType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case OfcST_FillTypeUnset:
		attr.Value = ""
	case OfcST_FillTypeGradientCenter:
		attr.Value = "gradientCenter"
	case OfcST_FillTypeSolid:
		attr.Value = "solid"
	case OfcST_FillTypePattern:
		attr.Value = "pattern"
	case OfcST_FillTypeTile:
		attr.Value = "tile"
	case OfcST_FillTypeFrame:
		attr.Value = "frame"
	case OfcST_FillTypeGradientUnscaled:
		attr.Value = "gradientUnscaled"
	case OfcST_FillTypeGradientRadial:
		attr.Value = "gradientRadial"
	case OfcST_FillTypeGradient:
		attr.Value = "gradient"
	case OfcST_FillTypeBackground:
		attr.Value = "background"
	}
	return attr, nil
}

func (e *OfcST_FillType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "gradientCenter":
		*e = 1
	case "solid":
		*e = 2
	case "pattern":
		*e = 3
	case "tile":
		*e = 4
	case "frame":
		*e = 5
	case "gradientUnscaled":
		*e = 6
	case "gradientRadial":
		*e = 7
	case "gradient":
		*e = 8
	case "background":
		*e = 9
	}
	return nil
}

func (m OfcST_FillType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *OfcST_FillType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	tok, err := d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		*m = 1
		return nil
	}
	if cd, ok := tok.(xml.CharData); !ok {
		return fmt.Errorf("expected char data, got %T", tok)
	} else {
		switch string(cd) {
		case "":
			*m = 0
		case "gradientCenter":
			*m = 1
		case "solid":
			*m = 2
		case "pattern":
			*m = 3
		case "tile":
			*m = 4
		case "frame":
			*m = 5
		case "gradientUnscaled":
			*m = 6
		case "gradientRadial":
			*m = 7
		case "gradient":
			*m = 8
		case "background":
			*m = 9
		}
	}
	tok, err = d.Token()
	if err != nil {
		return err
	}
	if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
		return nil
	}
	return fmt.Errorf("expected end element, got %v", tok)
}

func (m OfcST_FillType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "gradientCenter"
	case 2:
		return "solid"
	case 3:
		return "pattern"
	case 4:
		return "tile"
	case 5:
		return "frame"
	case 6:
		return "gradientUnscaled"
	case 7:
		return "gradientRadial"
	case 8:
		return "gradient"
	case 9:
		return "background"
	}
	return ""
}

func (m OfcST_FillType) Validate() error {
	return m.ValidateWithPath("")
}

func (m OfcST_FillType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_ShapeDefaults", NewOfcCT_ShapeDefaults)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Ink", NewOfcCT_Ink)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_SignatureLine", NewOfcCT_SignatureLine)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_ShapeLayout", NewOfcCT_ShapeLayout)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_IdMap", NewOfcCT_IdMap)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_RegroupTable", NewOfcCT_RegroupTable)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Entry", NewOfcCT_Entry)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Rules", NewOfcCT_Rules)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_R", NewOfcCT_R)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Proxy", NewOfcCT_Proxy)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Diagram", NewOfcCT_Diagram)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_EquationXml", NewOfcCT_EquationXml)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_RelationTable", NewOfcCT_RelationTable)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Relation", NewOfcCT_Relation)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_ColorMru", NewOfcCT_ColorMru)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_ColorMenu", NewOfcCT_ColorMenu)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Skew", NewOfcCT_Skew)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Extrusion", NewOfcCT_Extrusion)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Callout", NewOfcCT_Callout)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Lock", NewOfcCT_Lock)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_OLEObject", NewOfcCT_OLEObject)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Complex", NewOfcCT_Complex)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_StrokeChild", NewOfcCT_StrokeChild)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_ClipPath", NewOfcCT_ClipPath)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "CT_Fill", NewOfcCT_Fill)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "shapedefaults", NewOfcShapedefaults)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "shapelayout", NewOfcShapelayout)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "signatureline", NewOfcSignatureline)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "ink", NewOfcInk)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "diagram", NewOfcDiagram)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "equationxml", NewOfcEquationxml)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "skew", NewOfcSkew)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "extrusion", NewOfcExtrusion)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "callout", NewOfcCallout)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "lock", NewOfcLock)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "OLEObject", NewOfcOLEObject)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "complex", NewOfcComplex)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "left", NewOfcLeft)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "top", NewOfcTop)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "right", NewOfcRight)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "bottom", NewOfcBottom)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "column", NewOfcColumn)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "clippath", NewOfcClippath)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:office", "fill", NewOfcFill)
}
