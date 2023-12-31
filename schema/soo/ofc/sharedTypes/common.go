package sharedTypes

import (
	"encoding/xml"
	"fmt"
	"regexp"
)

const ST_GuidPattern = `\{[0-9A-F]{8}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{4}-[0-9A-F]{12}\}`

var ST_GuidPatternRe = regexp.MustCompile(ST_GuidPattern)

const ST_UniversalMeasurePattern = `-?[0-9]+(\.[0-9]+)?(mm|cm|in|pt|pc|pi)`

var ST_UniversalMeasurePatternRe = regexp.MustCompile(ST_UniversalMeasurePattern)

const ST_PositiveUniversalMeasurePattern = `[0-9]+(\.[0-9]+)?(mm|cm|in|pt|pc|pi)`

var ST_PositiveUniversalMeasurePatternRe = regexp.MustCompile(ST_PositiveUniversalMeasurePattern)

const ST_PercentagePattern = `-?[0-9]+(\.[0-9]+)?%`

var ST_PercentagePatternRe = regexp.MustCompile(ST_PercentagePattern)

const ST_FixedPercentagePattern = `-?((100)|([0-9][0-9]?))(\.[0-9][0-9]?)?%`

var ST_FixedPercentagePatternRe = regexp.MustCompile(ST_FixedPercentagePattern)

const ST_PositivePercentagePattern = `[0-9]+(\.[0-9]+)?%`

var ST_PositivePercentagePatternRe = regexp.MustCompile(ST_PositivePercentagePattern)

const ST_PositiveFixedPercentagePattern = `((100)|([0-9][0-9]?))(\.[0-9][0-9]?)?%`

var ST_PositiveFixedPercentagePatternRe = regexp.MustCompile(ST_PositiveFixedPercentagePattern)

func ParseUnionST_OnOff(s string) (ST_OnOff, error) {
	r := ST_OnOff{}
	switch s {
	case "true", "1", "on":
		tru := true
		r.Bool = &tru
	default:
		fals := false
		r.Bool = &fals
	}
	return r, nil
}

func b2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

type ST_CalendarType byte

const (
	ST_CalendarTypeUnset                ST_CalendarType = 0
	ST_CalendarTypeGregorian            ST_CalendarType = 1
	ST_CalendarTypeGregorianUs          ST_CalendarType = 2
	ST_CalendarTypeGregorianMeFrench    ST_CalendarType = 3
	ST_CalendarTypeGregorianArabic      ST_CalendarType = 4
	ST_CalendarTypeHijri                ST_CalendarType = 5
	ST_CalendarTypeHebrew               ST_CalendarType = 6
	ST_CalendarTypeTaiwan               ST_CalendarType = 7
	ST_CalendarTypeJapan                ST_CalendarType = 8
	ST_CalendarTypeThai                 ST_CalendarType = 9
	ST_CalendarTypeKorea                ST_CalendarType = 10
	ST_CalendarTypeSaka                 ST_CalendarType = 11
	ST_CalendarTypeGregorianXlitEnglish ST_CalendarType = 12
	ST_CalendarTypeGregorianXlitFrench  ST_CalendarType = 13
	ST_CalendarTypeNone                 ST_CalendarType = 14
)

func (e ST_CalendarType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CalendarTypeUnset:
		attr.Value = ""
	case ST_CalendarTypeGregorian:
		attr.Value = "gregorian"
	case ST_CalendarTypeGregorianUs:
		attr.Value = "gregorianUs"
	case ST_CalendarTypeGregorianMeFrench:
		attr.Value = "gregorianMeFrench"
	case ST_CalendarTypeGregorianArabic:
		attr.Value = "gregorianArabic"
	case ST_CalendarTypeHijri:
		attr.Value = "hijri"
	case ST_CalendarTypeHebrew:
		attr.Value = "hebrew"
	case ST_CalendarTypeTaiwan:
		attr.Value = "taiwan"
	case ST_CalendarTypeJapan:
		attr.Value = "japan"
	case ST_CalendarTypeThai:
		attr.Value = "thai"
	case ST_CalendarTypeKorea:
		attr.Value = "korea"
	case ST_CalendarTypeSaka:
		attr.Value = "saka"
	case ST_CalendarTypeGregorianXlitEnglish:
		attr.Value = "gregorianXlitEnglish"
	case ST_CalendarTypeGregorianXlitFrench:
		attr.Value = "gregorianXlitFrench"
	case ST_CalendarTypeNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_CalendarType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "gregorian":
		*e = 1
	case "gregorianUs":
		*e = 2
	case "gregorianMeFrench":
		*e = 3
	case "gregorianArabic":
		*e = 4
	case "hijri":
		*e = 5
	case "hebrew":
		*e = 6
	case "taiwan":
		*e = 7
	case "japan":
		*e = 8
	case "thai":
		*e = 9
	case "korea":
		*e = 10
	case "saka":
		*e = 11
	case "gregorianXlitEnglish":
		*e = 12
	case "gregorianXlitFrench":
		*e = 13
	case "none":
		*e = 14
	}
	return nil
}

func (m ST_CalendarType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CalendarType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "gregorian":
			*m = 1
		case "gregorianUs":
			*m = 2
		case "gregorianMeFrench":
			*m = 3
		case "gregorianArabic":
			*m = 4
		case "hijri":
			*m = 5
		case "hebrew":
			*m = 6
		case "taiwan":
			*m = 7
		case "japan":
			*m = 8
		case "thai":
			*m = 9
		case "korea":
			*m = 10
		case "saka":
			*m = 11
		case "gregorianXlitEnglish":
			*m = 12
		case "gregorianXlitFrench":
			*m = 13
		case "none":
			*m = 14
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

func (m ST_CalendarType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "gregorian"
	case 2:
		return "gregorianUs"
	case 3:
		return "gregorianMeFrench"
	case 4:
		return "gregorianArabic"
	case 5:
		return "hijri"
	case 6:
		return "hebrew"
	case 7:
		return "taiwan"
	case 8:
		return "japan"
	case 9:
		return "thai"
	case 10:
		return "korea"
	case 11:
		return "saka"
	case 12:
		return "gregorianXlitEnglish"
	case 13:
		return "gregorianXlitFrench"
	case 14:
		return "none"
	}
	return ""
}

func (m ST_CalendarType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CalendarType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AlgClass byte

const (
	ST_AlgClassUnset  ST_AlgClass = 0
	ST_AlgClassHash   ST_AlgClass = 1
	ST_AlgClassCustom ST_AlgClass = 2
)

func (e ST_AlgClass) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AlgClassUnset:
		attr.Value = ""
	case ST_AlgClassHash:
		attr.Value = "hash"
	case ST_AlgClassCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *ST_AlgClass) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "hash":
		*e = 1
	case "custom":
		*e = 2
	}
	return nil
}

func (m ST_AlgClass) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_AlgClass) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "hash":
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

func (m ST_AlgClass) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "hash"
	case 2:
		return "custom"
	}
	return ""
}

func (m ST_AlgClass) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AlgClass) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CryptProv byte

const (
	ST_CryptProvUnset   ST_CryptProv = 0
	ST_CryptProvRsaAES  ST_CryptProv = 1
	ST_CryptProvRsaFull ST_CryptProv = 2
	ST_CryptProvCustom  ST_CryptProv = 3
)

func (e ST_CryptProv) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CryptProvUnset:
		attr.Value = ""
	case ST_CryptProvRsaAES:
		attr.Value = "rsaAES"
	case ST_CryptProvRsaFull:
		attr.Value = "rsaFull"
	case ST_CryptProvCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *ST_CryptProv) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "rsaAES":
		*e = 1
	case "rsaFull":
		*e = 2
	case "custom":
		*e = 3
	}
	return nil
}

func (m ST_CryptProv) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CryptProv) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "rsaAES":
			*m = 1
		case "rsaFull":
			*m = 2
		case "custom":
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

func (m ST_CryptProv) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "rsaAES"
	case 2:
		return "rsaFull"
	case 3:
		return "custom"
	}
	return ""
}

func (m ST_CryptProv) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CryptProv) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_AlgType byte

const (
	ST_AlgTypeUnset   ST_AlgType = 0
	ST_AlgTypeTypeAny ST_AlgType = 1
	ST_AlgTypeCustom  ST_AlgType = 2
)

func (e ST_AlgType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AlgTypeUnset:
		attr.Value = ""
	case ST_AlgTypeTypeAny:
		attr.Value = "typeAny"
	case ST_AlgTypeCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *ST_AlgType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "typeAny":
		*e = 1
	case "custom":
		*e = 2
	}
	return nil
}

func (m ST_AlgType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_AlgType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "typeAny":
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

func (m ST_AlgType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "typeAny"
	case 2:
		return "custom"
	}
	return ""
}

func (m ST_AlgType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AlgType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_OnOff1 byte

const (
	ST_OnOff1Unset ST_OnOff1 = 0
	ST_OnOff1On    ST_OnOff1 = 1
	ST_OnOff1Off   ST_OnOff1 = 2
)

func (e ST_OnOff1) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_OnOff1Unset:
		attr.Value = ""
	case ST_OnOff1On:
		attr.Value = "on"
	case ST_OnOff1Off:
		attr.Value = "off"
	}
	return attr, nil
}

func (e *ST_OnOff1) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "on":
		*e = 1
	case "off":
		*e = 2
	}
	return nil
}

func (m ST_OnOff1) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_OnOff1) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "on":
			*m = 1
		case "off":
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

func (m ST_OnOff1) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "on"
	case 2:
		return "off"
	}
	return ""
}

func (m ST_OnOff1) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_OnOff1) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TrueFalse byte

const (
	ST_TrueFalseUnset ST_TrueFalse = 0
	ST_TrueFalseT     ST_TrueFalse = 1
	ST_TrueFalseF     ST_TrueFalse = 2
	ST_TrueFalseTrue  ST_TrueFalse = 3
	ST_TrueFalseFalse ST_TrueFalse = 4
)

func (e ST_TrueFalse) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TrueFalseUnset:
		attr.Value = ""
	case ST_TrueFalseT:
		attr.Value = "t"
	case ST_TrueFalseF:
		attr.Value = "f"
	case ST_TrueFalseTrue:
		attr.Value = "true"
	case ST_TrueFalseFalse:
		attr.Value = "false"
	}
	return attr, nil
}

func (e *ST_TrueFalse) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "t":
		*e = 1
	case "f":
		*e = 2
	case "true":
		*e = 3
	case "false":
		*e = 4
	}
	return nil
}

func (m ST_TrueFalse) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TrueFalse) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "t":
			*m = 1
		case "f":
			*m = 2
		case "true":
			*m = 3
		case "false":
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

func (m ST_TrueFalse) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "t"
	case 2:
		return "f"
	case 3:
		return "true"
	case 4:
		return "false"
	}
	return ""
}

func (m ST_TrueFalse) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TrueFalse) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TrueFalseBlank byte

const (
	ST_TrueFalseBlankUnset  ST_TrueFalseBlank = 0
	ST_TrueFalseBlankT      ST_TrueFalseBlank = 1
	ST_TrueFalseBlankF      ST_TrueFalseBlank = 2
	ST_TrueFalseBlankTrue   ST_TrueFalseBlank = 3
	ST_TrueFalseBlankFalse  ST_TrueFalseBlank = 4
	ST_TrueFalseBlankTrue_  ST_TrueFalseBlank = 6
	ST_TrueFalseBlankFalse_ ST_TrueFalseBlank = 7
)

func (e ST_TrueFalseBlank) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TrueFalseBlankUnset:
		attr.Value = ""
	case ST_TrueFalseBlankT:
		attr.Value = "t"
	case ST_TrueFalseBlankF:
		attr.Value = "f"
	case ST_TrueFalseBlankTrue:
		attr.Value = "true"
	case ST_TrueFalseBlankFalse:
		attr.Value = "false"
	case ST_TrueFalseBlankTrue_:
		attr.Value = "True"
	case ST_TrueFalseBlankFalse_:
		attr.Value = "False"
	}
	return attr, nil
}

func (e *ST_TrueFalseBlank) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "t":
		*e = 1
	case "f":
		*e = 2
	case "true":
		*e = 3
	case "false":
		*e = 4
	case "True":
		*e = 6
	case "False":
		*e = 7
	}
	return nil
}

func (m ST_TrueFalseBlank) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TrueFalseBlank) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "t":
			*m = 1
		case "f":
			*m = 2
		case "true":
			*m = 3
		case "false":
			*m = 4
		case "True":
			*m = 6
		case "False":
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

func (m ST_TrueFalseBlank) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "t"
	case 2:
		return "f"
	case 3:
		return "true"
	case 4:
		return "false"
	case 6:
		return "True"
	case 7:
		return "False"
	}
	return ""
}

func (m ST_TrueFalseBlank) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TrueFalseBlank) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_VerticalAlignRun byte

const (
	ST_VerticalAlignRunUnset       ST_VerticalAlignRun = 0
	ST_VerticalAlignRunBaseline    ST_VerticalAlignRun = 1
	ST_VerticalAlignRunSuperscript ST_VerticalAlignRun = 2
	ST_VerticalAlignRunSubscript   ST_VerticalAlignRun = 3
)

func (e ST_VerticalAlignRun) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VerticalAlignRunUnset:
		attr.Value = ""
	case ST_VerticalAlignRunBaseline:
		attr.Value = "baseline"
	case ST_VerticalAlignRunSuperscript:
		attr.Value = "superscript"
	case ST_VerticalAlignRunSubscript:
		attr.Value = "subscript"
	}
	return attr, nil
}

func (e *ST_VerticalAlignRun) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "baseline":
		*e = 1
	case "superscript":
		*e = 2
	case "subscript":
		*e = 3
	}
	return nil
}

func (m ST_VerticalAlignRun) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VerticalAlignRun) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "baseline":
			*m = 1
		case "superscript":
			*m = 2
		case "subscript":
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

func (m ST_VerticalAlignRun) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "baseline"
	case 2:
		return "superscript"
	case 3:
		return "subscript"
	}
	return ""
}

func (m ST_VerticalAlignRun) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VerticalAlignRun) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_XAlign byte

const (
	ST_XAlignUnset   ST_XAlign = 0
	ST_XAlignLeft    ST_XAlign = 1
	ST_XAlignCenter  ST_XAlign = 2
	ST_XAlignRight   ST_XAlign = 3
	ST_XAlignInside  ST_XAlign = 4
	ST_XAlignOutside ST_XAlign = 5
)

func (e ST_XAlign) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_XAlignUnset:
		attr.Value = ""
	case ST_XAlignLeft:
		attr.Value = "left"
	case ST_XAlignCenter:
		attr.Value = "center"
	case ST_XAlignRight:
		attr.Value = "right"
	case ST_XAlignInside:
		attr.Value = "inside"
	case ST_XAlignOutside:
		attr.Value = "outside"
	}
	return attr, nil
}

func (e *ST_XAlign) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "left":
		*e = 1
	case "center":
		*e = 2
	case "right":
		*e = 3
	case "inside":
		*e = 4
	case "outside":
		*e = 5
	}
	return nil
}

func (m ST_XAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_XAlign) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "center":
			*m = 2
		case "right":
			*m = 3
		case "inside":
			*m = 4
		case "outside":
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

func (m ST_XAlign) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "left"
	case 2:
		return "center"
	case 3:
		return "right"
	case 4:
		return "inside"
	case 5:
		return "outside"
	}
	return ""
}

func (m ST_XAlign) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_XAlign) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_YAlign byte

const (
	ST_YAlignUnset   ST_YAlign = 0
	ST_YAlignInline  ST_YAlign = 1
	ST_YAlignTop     ST_YAlign = 2
	ST_YAlignCenter  ST_YAlign = 3
	ST_YAlignBottom  ST_YAlign = 4
	ST_YAlignInside  ST_YAlign = 5
	ST_YAlignOutside ST_YAlign = 6
)

func (e ST_YAlign) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_YAlignUnset:
		attr.Value = ""
	case ST_YAlignInline:
		attr.Value = "inline"
	case ST_YAlignTop:
		attr.Value = "top"
	case ST_YAlignCenter:
		attr.Value = "center"
	case ST_YAlignBottom:
		attr.Value = "bottom"
	case ST_YAlignInside:
		attr.Value = "inside"
	case ST_YAlignOutside:
		attr.Value = "outside"
	}
	return attr, nil
}

func (e *ST_YAlign) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "inline":
		*e = 1
	case "top":
		*e = 2
	case "center":
		*e = 3
	case "bottom":
		*e = 4
	case "inside":
		*e = 5
	case "outside":
		*e = 6
	}
	return nil
}

func (m ST_YAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_YAlign) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "inline":
			*m = 1
		case "top":
			*m = 2
		case "center":
			*m = 3
		case "bottom":
			*m = 4
		case "inside":
			*m = 5
		case "outside":
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

func (m ST_YAlign) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "inline"
	case 2:
		return "top"
	case 3:
		return "center"
	case 4:
		return "bottom"
	case 5:
		return "inside"
	case 6:
		return "outside"
	}
	return ""
}

func (m ST_YAlign) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_YAlign) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ConformanceClass byte

const (
	ST_ConformanceClassUnset        ST_ConformanceClass = 0
	ST_ConformanceClassStrict       ST_ConformanceClass = 1
	ST_ConformanceClassTransitional ST_ConformanceClass = 2
)

func (e ST_ConformanceClass) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ConformanceClassUnset:
		attr.Value = ""
	case ST_ConformanceClassStrict:
		attr.Value = "strict"
	case ST_ConformanceClassTransitional:
		attr.Value = "transitional"
	}
	return attr, nil
}

func (e *ST_ConformanceClass) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "strict":
		*e = 1
	case "transitional":
		*e = 2
	}
	return nil
}

func (m ST_ConformanceClass) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ConformanceClass) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "strict":
			*m = 1
		case "transitional":
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

func (m ST_ConformanceClass) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "strict"
	case 2:
		return "transitional"
	}
	return ""
}

func (m ST_ConformanceClass) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ConformanceClass) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}
