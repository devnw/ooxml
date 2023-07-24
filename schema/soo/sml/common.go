package sml

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"

	"go.devnw.com/ooxml"
)

func ParseStdlibTime(s string) (time.Time, error) {
	return time.Time{}, nil
}
func ParseSliceST_Sqref(s string) (ST_Sqref, error) {
	return ST_Sqref(strings.Fields(s)), nil
}
func ParseSliceST_CellSpans(s string) (ST_CellSpans, error) {
	return ST_CellSpans(strings.Fields(s)), nil
}
func (s ST_Sqref) String() string {
	return strings.Join(s, " ")
}
func (s ST_CellSpans) String() string {
	return strings.Join(s, " ")
}

func b2i(b bool) uint8 {
	if b {
		return 1
	}
	return 0
}

type ST_FilterOperator byte

const (
	ST_FilterOperatorUnset              ST_FilterOperator = 0
	ST_FilterOperatorEqual              ST_FilterOperator = 1
	ST_FilterOperatorLessThan           ST_FilterOperator = 2
	ST_FilterOperatorLessThanOrEqual    ST_FilterOperator = 3
	ST_FilterOperatorNotEqual           ST_FilterOperator = 4
	ST_FilterOperatorGreaterThanOrEqual ST_FilterOperator = 5
	ST_FilterOperatorGreaterThan        ST_FilterOperator = 6
)

func (e ST_FilterOperator) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FilterOperatorUnset:
		attr.Value = ""
	case ST_FilterOperatorEqual:
		attr.Value = "equal"
	case ST_FilterOperatorLessThan:
		attr.Value = "lessThan"
	case ST_FilterOperatorLessThanOrEqual:
		attr.Value = "lessThanOrEqual"
	case ST_FilterOperatorNotEqual:
		attr.Value = "notEqual"
	case ST_FilterOperatorGreaterThanOrEqual:
		attr.Value = "greaterThanOrEqual"
	case ST_FilterOperatorGreaterThan:
		attr.Value = "greaterThan"
	}
	return attr, nil
}

func (e *ST_FilterOperator) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "equal":
		*e = 1
	case "lessThan":
		*e = 2
	case "lessThanOrEqual":
		*e = 3
	case "notEqual":
		*e = 4
	case "greaterThanOrEqual":
		*e = 5
	case "greaterThan":
		*e = 6
	}
	return nil
}

func (m ST_FilterOperator) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FilterOperator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "equal":
			*m = 1
		case "lessThan":
			*m = 2
		case "lessThanOrEqual":
			*m = 3
		case "notEqual":
			*m = 4
		case "greaterThanOrEqual":
			*m = 5
		case "greaterThan":
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

func (m ST_FilterOperator) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "equal"
	case 2:
		return "lessThan"
	case 3:
		return "lessThanOrEqual"
	case 4:
		return "notEqual"
	case 5:
		return "greaterThanOrEqual"
	case 6:
		return "greaterThan"
	}
	return ""
}

func (m ST_FilterOperator) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FilterOperator) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DynamicFilterType byte

const (
	ST_DynamicFilterTypeUnset        ST_DynamicFilterType = 0
	ST_DynamicFilterTypeNull         ST_DynamicFilterType = 1
	ST_DynamicFilterTypeAboveAverage ST_DynamicFilterType = 2
	ST_DynamicFilterTypeBelowAverage ST_DynamicFilterType = 3
	ST_DynamicFilterTypeTomorrow     ST_DynamicFilterType = 4
	ST_DynamicFilterTypeToday        ST_DynamicFilterType = 5
	ST_DynamicFilterTypeYesterday    ST_DynamicFilterType = 6
	ST_DynamicFilterTypeNextWeek     ST_DynamicFilterType = 7
	ST_DynamicFilterTypeThisWeek     ST_DynamicFilterType = 8
	ST_DynamicFilterTypeLastWeek     ST_DynamicFilterType = 9
	ST_DynamicFilterTypeNextMonth    ST_DynamicFilterType = 10
	ST_DynamicFilterTypeThisMonth    ST_DynamicFilterType = 11
	ST_DynamicFilterTypeLastMonth    ST_DynamicFilterType = 12
	ST_DynamicFilterTypeNextQuarter  ST_DynamicFilterType = 13
	ST_DynamicFilterTypeThisQuarter  ST_DynamicFilterType = 14
	ST_DynamicFilterTypeLastQuarter  ST_DynamicFilterType = 15
	ST_DynamicFilterTypeNextYear     ST_DynamicFilterType = 16
	ST_DynamicFilterTypeThisYear     ST_DynamicFilterType = 17
	ST_DynamicFilterTypeLastYear     ST_DynamicFilterType = 18
	ST_DynamicFilterTypeYearToDate   ST_DynamicFilterType = 19
	ST_DynamicFilterTypeQ1           ST_DynamicFilterType = 20
	ST_DynamicFilterTypeQ2           ST_DynamicFilterType = 21
	ST_DynamicFilterTypeQ3           ST_DynamicFilterType = 22
	ST_DynamicFilterTypeQ4           ST_DynamicFilterType = 23
	ST_DynamicFilterTypeM1           ST_DynamicFilterType = 24
	ST_DynamicFilterTypeM2           ST_DynamicFilterType = 25
	ST_DynamicFilterTypeM3           ST_DynamicFilterType = 26
	ST_DynamicFilterTypeM4           ST_DynamicFilterType = 27
	ST_DynamicFilterTypeM5           ST_DynamicFilterType = 28
	ST_DynamicFilterTypeM6           ST_DynamicFilterType = 29
	ST_DynamicFilterTypeM7           ST_DynamicFilterType = 30
	ST_DynamicFilterTypeM8           ST_DynamicFilterType = 31
	ST_DynamicFilterTypeM9           ST_DynamicFilterType = 32
	ST_DynamicFilterTypeM10          ST_DynamicFilterType = 33
	ST_DynamicFilterTypeM11          ST_DynamicFilterType = 34
	ST_DynamicFilterTypeM12          ST_DynamicFilterType = 35
)

func (e ST_DynamicFilterType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DynamicFilterTypeUnset:
		attr.Value = ""
	case ST_DynamicFilterTypeNull:
		attr.Value = "null"
	case ST_DynamicFilterTypeAboveAverage:
		attr.Value = "aboveAverage"
	case ST_DynamicFilterTypeBelowAverage:
		attr.Value = "belowAverage"
	case ST_DynamicFilterTypeTomorrow:
		attr.Value = "tomorrow"
	case ST_DynamicFilterTypeToday:
		attr.Value = "today"
	case ST_DynamicFilterTypeYesterday:
		attr.Value = "yesterday"
	case ST_DynamicFilterTypeNextWeek:
		attr.Value = "nextWeek"
	case ST_DynamicFilterTypeThisWeek:
		attr.Value = "thisWeek"
	case ST_DynamicFilterTypeLastWeek:
		attr.Value = "lastWeek"
	case ST_DynamicFilterTypeNextMonth:
		attr.Value = "nextMonth"
	case ST_DynamicFilterTypeThisMonth:
		attr.Value = "thisMonth"
	case ST_DynamicFilterTypeLastMonth:
		attr.Value = "lastMonth"
	case ST_DynamicFilterTypeNextQuarter:
		attr.Value = "nextQuarter"
	case ST_DynamicFilterTypeThisQuarter:
		attr.Value = "thisQuarter"
	case ST_DynamicFilterTypeLastQuarter:
		attr.Value = "lastQuarter"
	case ST_DynamicFilterTypeNextYear:
		attr.Value = "nextYear"
	case ST_DynamicFilterTypeThisYear:
		attr.Value = "thisYear"
	case ST_DynamicFilterTypeLastYear:
		attr.Value = "lastYear"
	case ST_DynamicFilterTypeYearToDate:
		attr.Value = "yearToDate"
	case ST_DynamicFilterTypeQ1:
		attr.Value = "Q1"
	case ST_DynamicFilterTypeQ2:
		attr.Value = "Q2"
	case ST_DynamicFilterTypeQ3:
		attr.Value = "Q3"
	case ST_DynamicFilterTypeQ4:
		attr.Value = "Q4"
	case ST_DynamicFilterTypeM1:
		attr.Value = "M1"
	case ST_DynamicFilterTypeM2:
		attr.Value = "M2"
	case ST_DynamicFilterTypeM3:
		attr.Value = "M3"
	case ST_DynamicFilterTypeM4:
		attr.Value = "M4"
	case ST_DynamicFilterTypeM5:
		attr.Value = "M5"
	case ST_DynamicFilterTypeM6:
		attr.Value = "M6"
	case ST_DynamicFilterTypeM7:
		attr.Value = "M7"
	case ST_DynamicFilterTypeM8:
		attr.Value = "M8"
	case ST_DynamicFilterTypeM9:
		attr.Value = "M9"
	case ST_DynamicFilterTypeM10:
		attr.Value = "M10"
	case ST_DynamicFilterTypeM11:
		attr.Value = "M11"
	case ST_DynamicFilterTypeM12:
		attr.Value = "M12"
	}
	return attr, nil
}

func (e *ST_DynamicFilterType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "null":
		*e = 1
	case "aboveAverage":
		*e = 2
	case "belowAverage":
		*e = 3
	case "tomorrow":
		*e = 4
	case "today":
		*e = 5
	case "yesterday":
		*e = 6
	case "nextWeek":
		*e = 7
	case "thisWeek":
		*e = 8
	case "lastWeek":
		*e = 9
	case "nextMonth":
		*e = 10
	case "thisMonth":
		*e = 11
	case "lastMonth":
		*e = 12
	case "nextQuarter":
		*e = 13
	case "thisQuarter":
		*e = 14
	case "lastQuarter":
		*e = 15
	case "nextYear":
		*e = 16
	case "thisYear":
		*e = 17
	case "lastYear":
		*e = 18
	case "yearToDate":
		*e = 19
	case "Q1":
		*e = 20
	case "Q2":
		*e = 21
	case "Q3":
		*e = 22
	case "Q4":
		*e = 23
	case "M1":
		*e = 24
	case "M2":
		*e = 25
	case "M3":
		*e = 26
	case "M4":
		*e = 27
	case "M5":
		*e = 28
	case "M6":
		*e = 29
	case "M7":
		*e = 30
	case "M8":
		*e = 31
	case "M9":
		*e = 32
	case "M10":
		*e = 33
	case "M11":
		*e = 34
	case "M12":
		*e = 35
	}
	return nil
}

func (m ST_DynamicFilterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DynamicFilterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "null":
			*m = 1
		case "aboveAverage":
			*m = 2
		case "belowAverage":
			*m = 3
		case "tomorrow":
			*m = 4
		case "today":
			*m = 5
		case "yesterday":
			*m = 6
		case "nextWeek":
			*m = 7
		case "thisWeek":
			*m = 8
		case "lastWeek":
			*m = 9
		case "nextMonth":
			*m = 10
		case "thisMonth":
			*m = 11
		case "lastMonth":
			*m = 12
		case "nextQuarter":
			*m = 13
		case "thisQuarter":
			*m = 14
		case "lastQuarter":
			*m = 15
		case "nextYear":
			*m = 16
		case "thisYear":
			*m = 17
		case "lastYear":
			*m = 18
		case "yearToDate":
			*m = 19
		case "Q1":
			*m = 20
		case "Q2":
			*m = 21
		case "Q3":
			*m = 22
		case "Q4":
			*m = 23
		case "M1":
			*m = 24
		case "M2":
			*m = 25
		case "M3":
			*m = 26
		case "M4":
			*m = 27
		case "M5":
			*m = 28
		case "M6":
			*m = 29
		case "M7":
			*m = 30
		case "M8":
			*m = 31
		case "M9":
			*m = 32
		case "M10":
			*m = 33
		case "M11":
			*m = 34
		case "M12":
			*m = 35
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

func (m ST_DynamicFilterType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "null"
	case 2:
		return "aboveAverage"
	case 3:
		return "belowAverage"
	case 4:
		return "tomorrow"
	case 5:
		return "today"
	case 6:
		return "yesterday"
	case 7:
		return "nextWeek"
	case 8:
		return "thisWeek"
	case 9:
		return "lastWeek"
	case 10:
		return "nextMonth"
	case 11:
		return "thisMonth"
	case 12:
		return "lastMonth"
	case 13:
		return "nextQuarter"
	case 14:
		return "thisQuarter"
	case 15:
		return "lastQuarter"
	case 16:
		return "nextYear"
	case 17:
		return "thisYear"
	case 18:
		return "lastYear"
	case 19:
		return "yearToDate"
	case 20:
		return "Q1"
	case 21:
		return "Q2"
	case 22:
		return "Q3"
	case 23:
		return "Q4"
	case 24:
		return "M1"
	case 25:
		return "M2"
	case 26:
		return "M3"
	case 27:
		return "M4"
	case 28:
		return "M5"
	case 29:
		return "M6"
	case 30:
		return "M7"
	case 31:
		return "M8"
	case 32:
		return "M9"
	case 33:
		return "M10"
	case 34:
		return "M11"
	case 35:
		return "M12"
	}
	return ""
}

func (m ST_DynamicFilterType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DynamicFilterType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_IconSetType byte

const (
	ST_IconSetTypeUnset           ST_IconSetType = 0
	ST_IconSetType3Arrows         ST_IconSetType = 1
	ST_IconSetType3ArrowsGray     ST_IconSetType = 2
	ST_IconSetType3Flags          ST_IconSetType = 3
	ST_IconSetType3TrafficLights1 ST_IconSetType = 4
	ST_IconSetType3TrafficLights2 ST_IconSetType = 5
	ST_IconSetType3Signs          ST_IconSetType = 6
	ST_IconSetType3Symbols        ST_IconSetType = 7
	ST_IconSetType3Symbols2       ST_IconSetType = 8
	ST_IconSetType4Arrows         ST_IconSetType = 9
	ST_IconSetType4ArrowsGray     ST_IconSetType = 10
	ST_IconSetType4RedToBlack     ST_IconSetType = 11
	ST_IconSetType4Rating         ST_IconSetType = 12
	ST_IconSetType4TrafficLights  ST_IconSetType = 13
	ST_IconSetType5Arrows         ST_IconSetType = 14
	ST_IconSetType5ArrowsGray     ST_IconSetType = 15
	ST_IconSetType5Rating         ST_IconSetType = 16
	ST_IconSetType5Quarters       ST_IconSetType = 17
)

func (e ST_IconSetType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_IconSetTypeUnset:
		attr.Value = ""
	case ST_IconSetType3Arrows:
		attr.Value = "3Arrows"
	case ST_IconSetType3ArrowsGray:
		attr.Value = "3ArrowsGray"
	case ST_IconSetType3Flags:
		attr.Value = "3Flags"
	case ST_IconSetType3TrafficLights1:
		attr.Value = "3TrafficLights1"
	case ST_IconSetType3TrafficLights2:
		attr.Value = "3TrafficLights2"
	case ST_IconSetType3Signs:
		attr.Value = "3Signs"
	case ST_IconSetType3Symbols:
		attr.Value = "3Symbols"
	case ST_IconSetType3Symbols2:
		attr.Value = "3Symbols2"
	case ST_IconSetType4Arrows:
		attr.Value = "4Arrows"
	case ST_IconSetType4ArrowsGray:
		attr.Value = "4ArrowsGray"
	case ST_IconSetType4RedToBlack:
		attr.Value = "4RedToBlack"
	case ST_IconSetType4Rating:
		attr.Value = "4Rating"
	case ST_IconSetType4TrafficLights:
		attr.Value = "4TrafficLights"
	case ST_IconSetType5Arrows:
		attr.Value = "5Arrows"
	case ST_IconSetType5ArrowsGray:
		attr.Value = "5ArrowsGray"
	case ST_IconSetType5Rating:
		attr.Value = "5Rating"
	case ST_IconSetType5Quarters:
		attr.Value = "5Quarters"
	}
	return attr, nil
}

func (e *ST_IconSetType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "3Arrows":
		*e = 1
	case "3ArrowsGray":
		*e = 2
	case "3Flags":
		*e = 3
	case "3TrafficLights1":
		*e = 4
	case "3TrafficLights2":
		*e = 5
	case "3Signs":
		*e = 6
	case "3Symbols":
		*e = 7
	case "3Symbols2":
		*e = 8
	case "4Arrows":
		*e = 9
	case "4ArrowsGray":
		*e = 10
	case "4RedToBlack":
		*e = 11
	case "4Rating":
		*e = 12
	case "4TrafficLights":
		*e = 13
	case "5Arrows":
		*e = 14
	case "5ArrowsGray":
		*e = 15
	case "5Rating":
		*e = 16
	case "5Quarters":
		*e = 17
	}
	return nil
}

func (m ST_IconSetType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_IconSetType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "3Arrows":
			*m = 1
		case "3ArrowsGray":
			*m = 2
		case "3Flags":
			*m = 3
		case "3TrafficLights1":
			*m = 4
		case "3TrafficLights2":
			*m = 5
		case "3Signs":
			*m = 6
		case "3Symbols":
			*m = 7
		case "3Symbols2":
			*m = 8
		case "4Arrows":
			*m = 9
		case "4ArrowsGray":
			*m = 10
		case "4RedToBlack":
			*m = 11
		case "4Rating":
			*m = 12
		case "4TrafficLights":
			*m = 13
		case "5Arrows":
			*m = 14
		case "5ArrowsGray":
			*m = 15
		case "5Rating":
			*m = 16
		case "5Quarters":
			*m = 17
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

func (m ST_IconSetType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "3Arrows"
	case 2:
		return "3ArrowsGray"
	case 3:
		return "3Flags"
	case 4:
		return "3TrafficLights1"
	case 5:
		return "3TrafficLights2"
	case 6:
		return "3Signs"
	case 7:
		return "3Symbols"
	case 8:
		return "3Symbols2"
	case 9:
		return "4Arrows"
	case 10:
		return "4ArrowsGray"
	case 11:
		return "4RedToBlack"
	case 12:
		return "4Rating"
	case 13:
		return "4TrafficLights"
	case 14:
		return "5Arrows"
	case 15:
		return "5ArrowsGray"
	case 16:
		return "5Rating"
	case 17:
		return "5Quarters"
	}
	return ""
}

func (m ST_IconSetType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_IconSetType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SortBy byte

const (
	ST_SortByUnset     ST_SortBy = 0
	ST_SortByValue     ST_SortBy = 1
	ST_SortByCellColor ST_SortBy = 2
	ST_SortByFontColor ST_SortBy = 3
	ST_SortByIcon      ST_SortBy = 4
)

func (e ST_SortBy) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SortByUnset:
		attr.Value = ""
	case ST_SortByValue:
		attr.Value = "value"
	case ST_SortByCellColor:
		attr.Value = "cellColor"
	case ST_SortByFontColor:
		attr.Value = "fontColor"
	case ST_SortByIcon:
		attr.Value = "icon"
	}
	return attr, nil
}

func (e *ST_SortBy) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "value":
		*e = 1
	case "cellColor":
		*e = 2
	case "fontColor":
		*e = 3
	case "icon":
		*e = 4
	}
	return nil
}

func (m ST_SortBy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SortBy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "value":
			*m = 1
		case "cellColor":
			*m = 2
		case "fontColor":
			*m = 3
		case "icon":
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

func (m ST_SortBy) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "value"
	case 2:
		return "cellColor"
	case 3:
		return "fontColor"
	case 4:
		return "icon"
	}
	return ""
}

func (m ST_SortBy) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SortBy) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SortMethod byte

const (
	ST_SortMethodUnset  ST_SortMethod = 0
	ST_SortMethodStroke ST_SortMethod = 1
	ST_SortMethodPinYin ST_SortMethod = 2
	ST_SortMethodNone   ST_SortMethod = 3
)

func (e ST_SortMethod) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SortMethodUnset:
		attr.Value = ""
	case ST_SortMethodStroke:
		attr.Value = "stroke"
	case ST_SortMethodPinYin:
		attr.Value = "pinYin"
	case ST_SortMethodNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_SortMethod) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "stroke":
		*e = 1
	case "pinYin":
		*e = 2
	case "none":
		*e = 3
	}
	return nil
}

func (m ST_SortMethod) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SortMethod) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "stroke":
			*m = 1
		case "pinYin":
			*m = 2
		case "none":
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

func (m ST_SortMethod) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "stroke"
	case 2:
		return "pinYin"
	case 3:
		return "none"
	}
	return ""
}

func (m ST_SortMethod) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SortMethod) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DateTimeGrouping byte

const (
	ST_DateTimeGroupingUnset  ST_DateTimeGrouping = 0
	ST_DateTimeGroupingYear   ST_DateTimeGrouping = 1
	ST_DateTimeGroupingMonth  ST_DateTimeGrouping = 2
	ST_DateTimeGroupingDay    ST_DateTimeGrouping = 3
	ST_DateTimeGroupingHour   ST_DateTimeGrouping = 4
	ST_DateTimeGroupingMinute ST_DateTimeGrouping = 5
	ST_DateTimeGroupingSecond ST_DateTimeGrouping = 6
)

func (e ST_DateTimeGrouping) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DateTimeGroupingUnset:
		attr.Value = ""
	case ST_DateTimeGroupingYear:
		attr.Value = "year"
	case ST_DateTimeGroupingMonth:
		attr.Value = "month"
	case ST_DateTimeGroupingDay:
		attr.Value = "day"
	case ST_DateTimeGroupingHour:
		attr.Value = "hour"
	case ST_DateTimeGroupingMinute:
		attr.Value = "minute"
	case ST_DateTimeGroupingSecond:
		attr.Value = "second"
	}
	return attr, nil
}

func (e *ST_DateTimeGrouping) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "year":
		*e = 1
	case "month":
		*e = 2
	case "day":
		*e = 3
	case "hour":
		*e = 4
	case "minute":
		*e = 5
	case "second":
		*e = 6
	}
	return nil
}

func (m ST_DateTimeGrouping) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DateTimeGrouping) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "year":
			*m = 1
		case "month":
			*m = 2
		case "day":
			*m = 3
		case "hour":
			*m = 4
		case "minute":
			*m = 5
		case "second":
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

func (m ST_DateTimeGrouping) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "year"
	case 2:
		return "month"
	case 3:
		return "day"
	case 4:
		return "hour"
	case 5:
		return "minute"
	case 6:
		return "second"
	}
	return ""
}

func (m ST_DateTimeGrouping) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DateTimeGrouping) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextHAlign byte

const (
	ST_TextHAlignUnset       ST_TextHAlign = 0
	ST_TextHAlignLeft        ST_TextHAlign = 1
	ST_TextHAlignCenter      ST_TextHAlign = 2
	ST_TextHAlignRight       ST_TextHAlign = 3
	ST_TextHAlignJustify     ST_TextHAlign = 4
	ST_TextHAlignDistributed ST_TextHAlign = 5
)

func (e ST_TextHAlign) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextHAlignUnset:
		attr.Value = ""
	case ST_TextHAlignLeft:
		attr.Value = "left"
	case ST_TextHAlignCenter:
		attr.Value = "center"
	case ST_TextHAlignRight:
		attr.Value = "right"
	case ST_TextHAlignJustify:
		attr.Value = "justify"
	case ST_TextHAlignDistributed:
		attr.Value = "distributed"
	}
	return attr, nil
}

func (e *ST_TextHAlign) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "left":
		*e = 1
	case "center":
		*e = 2
	case "right":
		*e = 3
	case "justify":
		*e = 4
	case "distributed":
		*e = 5
	}
	return nil
}

func (m ST_TextHAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextHAlign) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "justify":
			*m = 4
		case "distributed":
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

func (m ST_TextHAlign) String() string {
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
		return "justify"
	case 5:
		return "distributed"
	}
	return ""
}

func (m ST_TextHAlign) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextHAlign) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TextVAlign byte

const (
	ST_TextVAlignUnset       ST_TextVAlign = 0
	ST_TextVAlignTop         ST_TextVAlign = 1
	ST_TextVAlignCenter      ST_TextVAlign = 2
	ST_TextVAlignBottom      ST_TextVAlign = 3
	ST_TextVAlignJustify     ST_TextVAlign = 4
	ST_TextVAlignDistributed ST_TextVAlign = 5
)

func (e ST_TextVAlign) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TextVAlignUnset:
		attr.Value = ""
	case ST_TextVAlignTop:
		attr.Value = "top"
	case ST_TextVAlignCenter:
		attr.Value = "center"
	case ST_TextVAlignBottom:
		attr.Value = "bottom"
	case ST_TextVAlignJustify:
		attr.Value = "justify"
	case ST_TextVAlignDistributed:
		attr.Value = "distributed"
	}
	return attr, nil
}

func (e *ST_TextVAlign) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "top":
		*e = 1
	case "center":
		*e = 2
	case "bottom":
		*e = 3
	case "justify":
		*e = 4
	case "distributed":
		*e = 5
	}
	return nil
}

func (m ST_TextVAlign) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TextVAlign) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "justify":
			*m = 4
		case "distributed":
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

func (m ST_TextVAlign) String() string {
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
		return "justify"
	case 5:
		return "distributed"
	}
	return ""
}

func (m ST_TextVAlign) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextVAlign) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CredMethod byte

const (
	ST_CredMethodUnset      ST_CredMethod = 0
	ST_CredMethodIntegrated ST_CredMethod = 1
	ST_CredMethodNone       ST_CredMethod = 2
	ST_CredMethodStored     ST_CredMethod = 3
	ST_CredMethodPrompt     ST_CredMethod = 4
)

func (e ST_CredMethod) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CredMethodUnset:
		attr.Value = ""
	case ST_CredMethodIntegrated:
		attr.Value = "integrated"
	case ST_CredMethodNone:
		attr.Value = "none"
	case ST_CredMethodStored:
		attr.Value = "stored"
	case ST_CredMethodPrompt:
		attr.Value = "prompt"
	}
	return attr, nil
}

func (e *ST_CredMethod) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "integrated":
		*e = 1
	case "none":
		*e = 2
	case "stored":
		*e = 3
	case "prompt":
		*e = 4
	}
	return nil
}

func (m ST_CredMethod) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CredMethod) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "integrated":
			*m = 1
		case "none":
			*m = 2
		case "stored":
			*m = 3
		case "prompt":
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

func (m ST_CredMethod) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "integrated"
	case 2:
		return "none"
	case 3:
		return "stored"
	case 4:
		return "prompt"
	}
	return ""
}

func (m ST_CredMethod) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CredMethod) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HtmlFmt byte

const (
	ST_HtmlFmtUnset ST_HtmlFmt = 0
	ST_HtmlFmtNone  ST_HtmlFmt = 1
	ST_HtmlFmtRtf   ST_HtmlFmt = 2
	ST_HtmlFmtAll   ST_HtmlFmt = 3
)

func (e ST_HtmlFmt) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HtmlFmtUnset:
		attr.Value = ""
	case ST_HtmlFmtNone:
		attr.Value = "none"
	case ST_HtmlFmtRtf:
		attr.Value = "rtf"
	case ST_HtmlFmtAll:
		attr.Value = "all"
	}
	return attr, nil
}

func (e *ST_HtmlFmt) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "rtf":
		*e = 2
	case "all":
		*e = 3
	}
	return nil
}

func (m ST_HtmlFmt) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HtmlFmt) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "rtf":
			*m = 2
		case "all":
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

func (m ST_HtmlFmt) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "rtf"
	case 3:
		return "all"
	}
	return ""
}

func (m ST_HtmlFmt) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HtmlFmt) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ParameterType byte

const (
	ST_ParameterTypeUnset  ST_ParameterType = 0
	ST_ParameterTypePrompt ST_ParameterType = 1
	ST_ParameterTypeValue  ST_ParameterType = 2
	ST_ParameterTypeCell   ST_ParameterType = 3
)

func (e ST_ParameterType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ParameterTypeUnset:
		attr.Value = ""
	case ST_ParameterTypePrompt:
		attr.Value = "prompt"
	case ST_ParameterTypeValue:
		attr.Value = "value"
	case ST_ParameterTypeCell:
		attr.Value = "cell"
	}
	return attr, nil
}

func (e *ST_ParameterType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "prompt":
		*e = 1
	case "value":
		*e = 2
	case "cell":
		*e = 3
	}
	return nil
}

func (m ST_ParameterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ParameterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "prompt":
			*m = 1
		case "value":
			*m = 2
		case "cell":
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

func (m ST_ParameterType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "prompt"
	case 2:
		return "value"
	case 3:
		return "cell"
	}
	return ""
}

func (m ST_ParameterType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ParameterType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FileType byte

const (
	ST_FileTypeUnset ST_FileType = 0
	ST_FileTypeMac   ST_FileType = 1
	ST_FileTypeWin   ST_FileType = 2
	ST_FileTypeDos   ST_FileType = 3
	ST_FileTypeLin   ST_FileType = 4
	ST_FileTypeOther ST_FileType = 5
)

func (e ST_FileType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FileTypeUnset:
		attr.Value = ""
	case ST_FileTypeMac:
		attr.Value = "mac"
	case ST_FileTypeWin:
		attr.Value = "win"
	case ST_FileTypeDos:
		attr.Value = "dos"
	case ST_FileTypeLin:
		attr.Value = "lin"
	case ST_FileTypeOther:
		attr.Value = "other"
	}
	return attr, nil
}

func (e *ST_FileType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "mac":
		*e = 1
	case "win":
		*e = 2
	case "dos":
		*e = 3
	case "lin":
		*e = 4
	case "other":
		*e = 5
	}
	return nil
}

func (m ST_FileType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FileType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "mac":
			*m = 1
		case "win":
			*m = 2
		case "dos":
			*m = 3
		case "lin":
			*m = 4
		case "other":
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

func (m ST_FileType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "mac"
	case 2:
		return "win"
	case 3:
		return "dos"
	case 4:
		return "lin"
	case 5:
		return "other"
	}
	return ""
}

func (m ST_FileType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FileType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Qualifier byte

const (
	ST_QualifierUnset       ST_Qualifier = 0
	ST_QualifierDoubleQuote ST_Qualifier = 1
	ST_QualifierSingleQuote ST_Qualifier = 2
	ST_QualifierNone        ST_Qualifier = 3
)

func (e ST_Qualifier) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_QualifierUnset:
		attr.Value = ""
	case ST_QualifierDoubleQuote:
		attr.Value = "doubleQuote"
	case ST_QualifierSingleQuote:
		attr.Value = "singleQuote"
	case ST_QualifierNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_Qualifier) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "doubleQuote":
		*e = 1
	case "singleQuote":
		*e = 2
	case "none":
		*e = 3
	}
	return nil
}

func (m ST_Qualifier) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Qualifier) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "doubleQuote":
			*m = 1
		case "singleQuote":
			*m = 2
		case "none":
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

func (m ST_Qualifier) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "doubleQuote"
	case 2:
		return "singleQuote"
	case 3:
		return "none"
	}
	return ""
}

func (m ST_Qualifier) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Qualifier) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ExternalConnectionType byte

const (
	ST_ExternalConnectionTypeUnset   ST_ExternalConnectionType = 0
	ST_ExternalConnectionTypeGeneral ST_ExternalConnectionType = 1
	ST_ExternalConnectionTypeText    ST_ExternalConnectionType = 2
	ST_ExternalConnectionTypeMDY     ST_ExternalConnectionType = 3
	ST_ExternalConnectionTypeDMY     ST_ExternalConnectionType = 4
	ST_ExternalConnectionTypeYMD     ST_ExternalConnectionType = 5
	ST_ExternalConnectionTypeMYD     ST_ExternalConnectionType = 6
	ST_ExternalConnectionTypeDYM     ST_ExternalConnectionType = 7
	ST_ExternalConnectionTypeYDM     ST_ExternalConnectionType = 8
	ST_ExternalConnectionTypeSkip    ST_ExternalConnectionType = 9
	ST_ExternalConnectionTypeEMD     ST_ExternalConnectionType = 10
)

func (e ST_ExternalConnectionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ExternalConnectionTypeUnset:
		attr.Value = ""
	case ST_ExternalConnectionTypeGeneral:
		attr.Value = "general"
	case ST_ExternalConnectionTypeText:
		attr.Value = "text"
	case ST_ExternalConnectionTypeMDY:
		attr.Value = "MDY"
	case ST_ExternalConnectionTypeDMY:
		attr.Value = "DMY"
	case ST_ExternalConnectionTypeYMD:
		attr.Value = "YMD"
	case ST_ExternalConnectionTypeMYD:
		attr.Value = "MYD"
	case ST_ExternalConnectionTypeDYM:
		attr.Value = "DYM"
	case ST_ExternalConnectionTypeYDM:
		attr.Value = "YDM"
	case ST_ExternalConnectionTypeSkip:
		attr.Value = "skip"
	case ST_ExternalConnectionTypeEMD:
		attr.Value = "EMD"
	}
	return attr, nil
}

func (e *ST_ExternalConnectionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "general":
		*e = 1
	case "text":
		*e = 2
	case "MDY":
		*e = 3
	case "DMY":
		*e = 4
	case "YMD":
		*e = 5
	case "MYD":
		*e = 6
	case "DYM":
		*e = 7
	case "YDM":
		*e = 8
	case "skip":
		*e = 9
	case "EMD":
		*e = 10
	}
	return nil
}

func (m ST_ExternalConnectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ExternalConnectionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "general":
			*m = 1
		case "text":
			*m = 2
		case "MDY":
			*m = 3
		case "DMY":
			*m = 4
		case "YMD":
			*m = 5
		case "MYD":
			*m = 6
		case "DYM":
			*m = 7
		case "YDM":
			*m = 8
		case "skip":
			*m = 9
		case "EMD":
			*m = 10
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

func (m ST_ExternalConnectionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "general"
	case 2:
		return "text"
	case 3:
		return "MDY"
	case 4:
		return "DMY"
	case 5:
		return "YMD"
	case 6:
		return "MYD"
	case 7:
		return "DYM"
	case 8:
		return "YDM"
	case 9:
		return "skip"
	case 10:
		return "EMD"
	}
	return ""
}

func (m ST_ExternalConnectionType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ExternalConnectionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SourceType byte

const (
	ST_SourceTypeUnset         ST_SourceType = 0
	ST_SourceTypeWorksheet     ST_SourceType = 1
	ST_SourceTypeExternal      ST_SourceType = 2
	ST_SourceTypeConsolidation ST_SourceType = 3
	ST_SourceTypeScenario      ST_SourceType = 4
)

func (e ST_SourceType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SourceTypeUnset:
		attr.Value = ""
	case ST_SourceTypeWorksheet:
		attr.Value = "worksheet"
	case ST_SourceTypeExternal:
		attr.Value = "external"
	case ST_SourceTypeConsolidation:
		attr.Value = "consolidation"
	case ST_SourceTypeScenario:
		attr.Value = "scenario"
	}
	return attr, nil
}

func (e *ST_SourceType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "worksheet":
		*e = 1
	case "external":
		*e = 2
	case "consolidation":
		*e = 3
	case "scenario":
		*e = 4
	}
	return nil
}

func (m ST_SourceType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SourceType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "worksheet":
			*m = 1
		case "external":
			*m = 2
		case "consolidation":
			*m = 3
		case "scenario":
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

func (m ST_SourceType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "worksheet"
	case 2:
		return "external"
	case 3:
		return "consolidation"
	case 4:
		return "scenario"
	}
	return ""
}

func (m ST_SourceType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SourceType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_GroupBy byte

const (
	ST_GroupByUnset    ST_GroupBy = 0
	ST_GroupByRange    ST_GroupBy = 1
	ST_GroupBySeconds  ST_GroupBy = 2
	ST_GroupByMinutes  ST_GroupBy = 3
	ST_GroupByHours    ST_GroupBy = 4
	ST_GroupByDays     ST_GroupBy = 5
	ST_GroupByMonths   ST_GroupBy = 6
	ST_GroupByQuarters ST_GroupBy = 7
	ST_GroupByYears    ST_GroupBy = 8
)

func (e ST_GroupBy) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_GroupByUnset:
		attr.Value = ""
	case ST_GroupByRange:
		attr.Value = "range"
	case ST_GroupBySeconds:
		attr.Value = "seconds"
	case ST_GroupByMinutes:
		attr.Value = "minutes"
	case ST_GroupByHours:
		attr.Value = "hours"
	case ST_GroupByDays:
		attr.Value = "days"
	case ST_GroupByMonths:
		attr.Value = "months"
	case ST_GroupByQuarters:
		attr.Value = "quarters"
	case ST_GroupByYears:
		attr.Value = "years"
	}
	return attr, nil
}

func (e *ST_GroupBy) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "range":
		*e = 1
	case "seconds":
		*e = 2
	case "minutes":
		*e = 3
	case "hours":
		*e = 4
	case "days":
		*e = 5
	case "months":
		*e = 6
	case "quarters":
		*e = 7
	case "years":
		*e = 8
	}
	return nil
}

func (m ST_GroupBy) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_GroupBy) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "range":
			*m = 1
		case "seconds":
			*m = 2
		case "minutes":
			*m = 3
		case "hours":
			*m = 4
		case "days":
			*m = 5
		case "months":
			*m = 6
		case "quarters":
			*m = 7
		case "years":
			*m = 8
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

func (m ST_GroupBy) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "range"
	case 2:
		return "seconds"
	case 3:
		return "minutes"
	case 4:
		return "hours"
	case 5:
		return "days"
	case 6:
		return "months"
	case 7:
		return "quarters"
	case 8:
		return "years"
	}
	return ""
}

func (m ST_GroupBy) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_GroupBy) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SortType byte

const (
	ST_SortTypeUnset             ST_SortType = 0
	ST_SortTypeNone              ST_SortType = 1
	ST_SortTypeAscending         ST_SortType = 2
	ST_SortTypeDescending        ST_SortType = 3
	ST_SortTypeAscendingAlpha    ST_SortType = 4
	ST_SortTypeDescendingAlpha   ST_SortType = 5
	ST_SortTypeAscendingNatural  ST_SortType = 6
	ST_SortTypeDescendingNatural ST_SortType = 7
)

func (e ST_SortType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SortTypeUnset:
		attr.Value = ""
	case ST_SortTypeNone:
		attr.Value = "none"
	case ST_SortTypeAscending:
		attr.Value = "ascending"
	case ST_SortTypeDescending:
		attr.Value = "descending"
	case ST_SortTypeAscendingAlpha:
		attr.Value = "ascendingAlpha"
	case ST_SortTypeDescendingAlpha:
		attr.Value = "descendingAlpha"
	case ST_SortTypeAscendingNatural:
		attr.Value = "ascendingNatural"
	case ST_SortTypeDescendingNatural:
		attr.Value = "descendingNatural"
	}
	return attr, nil
}

func (e *ST_SortType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "ascending":
		*e = 2
	case "descending":
		*e = 3
	case "ascendingAlpha":
		*e = 4
	case "descendingAlpha":
		*e = 5
	case "ascendingNatural":
		*e = 6
	case "descendingNatural":
		*e = 7
	}
	return nil
}

func (m ST_SortType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SortType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "ascending":
			*m = 2
		case "descending":
			*m = 3
		case "ascendingAlpha":
			*m = 4
		case "descendingAlpha":
			*m = 5
		case "ascendingNatural":
			*m = 6
		case "descendingNatural":
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

func (m ST_SortType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "ascending"
	case 3:
		return "descending"
	case 4:
		return "ascendingAlpha"
	case 5:
		return "descendingAlpha"
	case 6:
		return "ascendingNatural"
	case 7:
		return "descendingNatural"
	}
	return ""
}

func (m ST_SortType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SortType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Scope byte

const (
	ST_ScopeUnset     ST_Scope = 0
	ST_ScopeSelection ST_Scope = 1
	ST_ScopeData      ST_Scope = 2
	ST_ScopeField     ST_Scope = 3
)

func (e ST_Scope) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ScopeUnset:
		attr.Value = ""
	case ST_ScopeSelection:
		attr.Value = "selection"
	case ST_ScopeData:
		attr.Value = "data"
	case ST_ScopeField:
		attr.Value = "field"
	}
	return attr, nil
}

func (e *ST_Scope) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "selection":
		*e = 1
	case "data":
		*e = 2
	case "field":
		*e = 3
	}
	return nil
}

func (m ST_Scope) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Scope) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "selection":
			*m = 1
		case "data":
			*m = 2
		case "field":
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

func (m ST_Scope) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "selection"
	case 2:
		return "data"
	case 3:
		return "field"
	}
	return ""
}

func (m ST_Scope) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Scope) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Type byte

const (
	ST_TypeUnset  ST_Type = 0
	ST_TypeNone   ST_Type = 1
	ST_TypeAll    ST_Type = 2
	ST_TypeRow    ST_Type = 3
	ST_TypeColumn ST_Type = 4
)

func (e ST_Type) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TypeUnset:
		attr.Value = ""
	case ST_TypeNone:
		attr.Value = "none"
	case ST_TypeAll:
		attr.Value = "all"
	case ST_TypeRow:
		attr.Value = "row"
	case ST_TypeColumn:
		attr.Value = "column"
	}
	return attr, nil
}

func (e *ST_Type) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "all":
		*e = 2
	case "row":
		*e = 3
	case "column":
		*e = 4
	}
	return nil
}

func (m ST_Type) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Type) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "all":
			*m = 2
		case "row":
			*m = 3
		case "column":
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

func (m ST_Type) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "all"
	case 3:
		return "row"
	case 4:
		return "column"
	}
	return ""
}

func (m ST_Type) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Type) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ShowDataAs byte

const (
	ST_ShowDataAsUnset          ST_ShowDataAs = 0
	ST_ShowDataAsNormal         ST_ShowDataAs = 1
	ST_ShowDataAsDifference     ST_ShowDataAs = 2
	ST_ShowDataAsPercent        ST_ShowDataAs = 3
	ST_ShowDataAsPercentDiff    ST_ShowDataAs = 4
	ST_ShowDataAsRunTotal       ST_ShowDataAs = 5
	ST_ShowDataAsPercentOfRow   ST_ShowDataAs = 6
	ST_ShowDataAsPercentOfCol   ST_ShowDataAs = 7
	ST_ShowDataAsPercentOfTotal ST_ShowDataAs = 8
	ST_ShowDataAsIndex          ST_ShowDataAs = 9
)

func (e ST_ShowDataAs) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ShowDataAsUnset:
		attr.Value = ""
	case ST_ShowDataAsNormal:
		attr.Value = "normal"
	case ST_ShowDataAsDifference:
		attr.Value = "difference"
	case ST_ShowDataAsPercent:
		attr.Value = "percent"
	case ST_ShowDataAsPercentDiff:
		attr.Value = "percentDiff"
	case ST_ShowDataAsRunTotal:
		attr.Value = "runTotal"
	case ST_ShowDataAsPercentOfRow:
		attr.Value = "percentOfRow"
	case ST_ShowDataAsPercentOfCol:
		attr.Value = "percentOfCol"
	case ST_ShowDataAsPercentOfTotal:
		attr.Value = "percentOfTotal"
	case ST_ShowDataAsIndex:
		attr.Value = "index"
	}
	return attr, nil
}

func (e *ST_ShowDataAs) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "normal":
		*e = 1
	case "difference":
		*e = 2
	case "percent":
		*e = 3
	case "percentDiff":
		*e = 4
	case "runTotal":
		*e = 5
	case "percentOfRow":
		*e = 6
	case "percentOfCol":
		*e = 7
	case "percentOfTotal":
		*e = 8
	case "index":
		*e = 9
	}
	return nil
}

func (m ST_ShowDataAs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ShowDataAs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "normal":
			*m = 1
		case "difference":
			*m = 2
		case "percent":
			*m = 3
		case "percentDiff":
			*m = 4
		case "runTotal":
			*m = 5
		case "percentOfRow":
			*m = 6
		case "percentOfCol":
			*m = 7
		case "percentOfTotal":
			*m = 8
		case "index":
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

func (m ST_ShowDataAs) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "normal"
	case 2:
		return "difference"
	case 3:
		return "percent"
	case 4:
		return "percentDiff"
	case 5:
		return "runTotal"
	case 6:
		return "percentOfRow"
	case 7:
		return "percentOfCol"
	case 8:
		return "percentOfTotal"
	case 9:
		return "index"
	}
	return ""
}

func (m ST_ShowDataAs) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ShowDataAs) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ItemType byte

const (
	ST_ItemTypeUnset   ST_ItemType = 0
	ST_ItemTypeData    ST_ItemType = 1
	ST_ItemTypeDefault ST_ItemType = 2
	ST_ItemTypeSum     ST_ItemType = 3
	ST_ItemTypeCountA  ST_ItemType = 4
	ST_ItemTypeAvg     ST_ItemType = 5
	ST_ItemTypeMax     ST_ItemType = 6
	ST_ItemTypeMin     ST_ItemType = 7
	ST_ItemTypeProduct ST_ItemType = 8
	ST_ItemTypeCount   ST_ItemType = 9
	ST_ItemTypeStdDev  ST_ItemType = 10
	ST_ItemTypeStdDevP ST_ItemType = 11
	ST_ItemTypeVar     ST_ItemType = 12
	ST_ItemTypeVarP    ST_ItemType = 13
	ST_ItemTypeGrand   ST_ItemType = 14
	ST_ItemTypeBlank   ST_ItemType = 15
)

func (e ST_ItemType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ItemTypeUnset:
		attr.Value = ""
	case ST_ItemTypeData:
		attr.Value = "data"
	case ST_ItemTypeDefault:
		attr.Value = "default"
	case ST_ItemTypeSum:
		attr.Value = "sum"
	case ST_ItemTypeCountA:
		attr.Value = "countA"
	case ST_ItemTypeAvg:
		attr.Value = "avg"
	case ST_ItemTypeMax:
		attr.Value = "max"
	case ST_ItemTypeMin:
		attr.Value = "min"
	case ST_ItemTypeProduct:
		attr.Value = "product"
	case ST_ItemTypeCount:
		attr.Value = "count"
	case ST_ItemTypeStdDev:
		attr.Value = "stdDev"
	case ST_ItemTypeStdDevP:
		attr.Value = "stdDevP"
	case ST_ItemTypeVar:
		attr.Value = "var"
	case ST_ItemTypeVarP:
		attr.Value = "varP"
	case ST_ItemTypeGrand:
		attr.Value = "grand"
	case ST_ItemTypeBlank:
		attr.Value = "blank"
	}
	return attr, nil
}

func (e *ST_ItemType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "data":
		*e = 1
	case "default":
		*e = 2
	case "sum":
		*e = 3
	case "countA":
		*e = 4
	case "avg":
		*e = 5
	case "max":
		*e = 6
	case "min":
		*e = 7
	case "product":
		*e = 8
	case "count":
		*e = 9
	case "stdDev":
		*e = 10
	case "stdDevP":
		*e = 11
	case "var":
		*e = 12
	case "varP":
		*e = 13
	case "grand":
		*e = 14
	case "blank":
		*e = 15
	}
	return nil
}

func (m ST_ItemType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ItemType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "data":
			*m = 1
		case "default":
			*m = 2
		case "sum":
			*m = 3
		case "countA":
			*m = 4
		case "avg":
			*m = 5
		case "max":
			*m = 6
		case "min":
			*m = 7
		case "product":
			*m = 8
		case "count":
			*m = 9
		case "stdDev":
			*m = 10
		case "stdDevP":
			*m = 11
		case "var":
			*m = 12
		case "varP":
			*m = 13
		case "grand":
			*m = 14
		case "blank":
			*m = 15
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

func (m ST_ItemType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "data"
	case 2:
		return "default"
	case 3:
		return "sum"
	case 4:
		return "countA"
	case 5:
		return "avg"
	case 6:
		return "max"
	case 7:
		return "min"
	case 8:
		return "product"
	case 9:
		return "count"
	case 10:
		return "stdDev"
	case 11:
		return "stdDevP"
	case 12:
		return "var"
	case 13:
		return "varP"
	case 14:
		return "grand"
	case 15:
		return "blank"
	}
	return ""
}

func (m ST_ItemType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ItemType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FormatAction byte

const (
	ST_FormatActionUnset      ST_FormatAction = 0
	ST_FormatActionBlank      ST_FormatAction = 1
	ST_FormatActionFormatting ST_FormatAction = 2
	ST_FormatActionDrill      ST_FormatAction = 3
	ST_FormatActionFormula    ST_FormatAction = 4
)

func (e ST_FormatAction) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FormatActionUnset:
		attr.Value = ""
	case ST_FormatActionBlank:
		attr.Value = "blank"
	case ST_FormatActionFormatting:
		attr.Value = "formatting"
	case ST_FormatActionDrill:
		attr.Value = "drill"
	case ST_FormatActionFormula:
		attr.Value = "formula"
	}
	return attr, nil
}

func (e *ST_FormatAction) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "blank":
		*e = 1
	case "formatting":
		*e = 2
	case "drill":
		*e = 3
	case "formula":
		*e = 4
	}
	return nil
}

func (m ST_FormatAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FormatAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "blank":
			*m = 1
		case "formatting":
			*m = 2
		case "drill":
			*m = 3
		case "formula":
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

func (m ST_FormatAction) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "blank"
	case 2:
		return "formatting"
	case 3:
		return "drill"
	case 4:
		return "formula"
	}
	return ""
}

func (m ST_FormatAction) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FormatAction) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FieldSortType byte

const (
	ST_FieldSortTypeUnset      ST_FieldSortType = 0
	ST_FieldSortTypeManual     ST_FieldSortType = 1
	ST_FieldSortTypeAscending  ST_FieldSortType = 2
	ST_FieldSortTypeDescending ST_FieldSortType = 3
)

func (e ST_FieldSortType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FieldSortTypeUnset:
		attr.Value = ""
	case ST_FieldSortTypeManual:
		attr.Value = "manual"
	case ST_FieldSortTypeAscending:
		attr.Value = "ascending"
	case ST_FieldSortTypeDescending:
		attr.Value = "descending"
	}
	return attr, nil
}

func (e *ST_FieldSortType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "manual":
		*e = 1
	case "ascending":
		*e = 2
	case "descending":
		*e = 3
	}
	return nil
}

func (m ST_FieldSortType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FieldSortType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "manual":
			*m = 1
		case "ascending":
			*m = 2
		case "descending":
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

func (m ST_FieldSortType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "manual"
	case 2:
		return "ascending"
	case 3:
		return "descending"
	}
	return ""
}

func (m ST_FieldSortType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FieldSortType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PivotFilterType byte

const (
	ST_PivotFilterTypeUnset                     ST_PivotFilterType = 0
	ST_PivotFilterTypeUnknown                   ST_PivotFilterType = 1
	ST_PivotFilterTypeCount                     ST_PivotFilterType = 2
	ST_PivotFilterTypePercent                   ST_PivotFilterType = 3
	ST_PivotFilterTypeSum                       ST_PivotFilterType = 4
	ST_PivotFilterTypeCaptionEqual              ST_PivotFilterType = 5
	ST_PivotFilterTypeCaptionNotEqual           ST_PivotFilterType = 6
	ST_PivotFilterTypeCaptionBeginsWith         ST_PivotFilterType = 7
	ST_PivotFilterTypeCaptionNotBeginsWith      ST_PivotFilterType = 8
	ST_PivotFilterTypeCaptionEndsWith           ST_PivotFilterType = 9
	ST_PivotFilterTypeCaptionNotEndsWith        ST_PivotFilterType = 10
	ST_PivotFilterTypeCaptionContains           ST_PivotFilterType = 11
	ST_PivotFilterTypeCaptionNotContains        ST_PivotFilterType = 12
	ST_PivotFilterTypeCaptionGreaterThan        ST_PivotFilterType = 13
	ST_PivotFilterTypeCaptionGreaterThanOrEqual ST_PivotFilterType = 14
	ST_PivotFilterTypeCaptionLessThan           ST_PivotFilterType = 15
	ST_PivotFilterTypeCaptionLessThanOrEqual    ST_PivotFilterType = 16
	ST_PivotFilterTypeCaptionBetween            ST_PivotFilterType = 17
	ST_PivotFilterTypeCaptionNotBetween         ST_PivotFilterType = 18
	ST_PivotFilterTypeValueEqual                ST_PivotFilterType = 19
	ST_PivotFilterTypeValueNotEqual             ST_PivotFilterType = 20
	ST_PivotFilterTypeValueGreaterThan          ST_PivotFilterType = 21
	ST_PivotFilterTypeValueGreaterThanOrEqual   ST_PivotFilterType = 22
	ST_PivotFilterTypeValueLessThan             ST_PivotFilterType = 23
	ST_PivotFilterTypeValueLessThanOrEqual      ST_PivotFilterType = 24
	ST_PivotFilterTypeValueBetween              ST_PivotFilterType = 25
	ST_PivotFilterTypeValueNotBetween           ST_PivotFilterType = 26
	ST_PivotFilterTypeDateEqual                 ST_PivotFilterType = 27
	ST_PivotFilterTypeDateNotEqual              ST_PivotFilterType = 28
	ST_PivotFilterTypeDateOlderThan             ST_PivotFilterType = 29
	ST_PivotFilterTypeDateOlderThanOrEqual      ST_PivotFilterType = 30
	ST_PivotFilterTypeDateNewerThan             ST_PivotFilterType = 31
	ST_PivotFilterTypeDateNewerThanOrEqual      ST_PivotFilterType = 32
	ST_PivotFilterTypeDateBetween               ST_PivotFilterType = 33
	ST_PivotFilterTypeDateNotBetween            ST_PivotFilterType = 34
	ST_PivotFilterTypeTomorrow                  ST_PivotFilterType = 35
	ST_PivotFilterTypeToday                     ST_PivotFilterType = 36
	ST_PivotFilterTypeYesterday                 ST_PivotFilterType = 37
	ST_PivotFilterTypeNextWeek                  ST_PivotFilterType = 38
	ST_PivotFilterTypeThisWeek                  ST_PivotFilterType = 39
	ST_PivotFilterTypeLastWeek                  ST_PivotFilterType = 40
	ST_PivotFilterTypeNextMonth                 ST_PivotFilterType = 41
	ST_PivotFilterTypeThisMonth                 ST_PivotFilterType = 42
	ST_PivotFilterTypeLastMonth                 ST_PivotFilterType = 43
	ST_PivotFilterTypeNextQuarter               ST_PivotFilterType = 44
	ST_PivotFilterTypeThisQuarter               ST_PivotFilterType = 45
	ST_PivotFilterTypeLastQuarter               ST_PivotFilterType = 46
	ST_PivotFilterTypeNextYear                  ST_PivotFilterType = 47
	ST_PivotFilterTypeThisYear                  ST_PivotFilterType = 48
	ST_PivotFilterTypeLastYear                  ST_PivotFilterType = 49
	ST_PivotFilterTypeYearToDate                ST_PivotFilterType = 50
	ST_PivotFilterTypeQ1                        ST_PivotFilterType = 51
	ST_PivotFilterTypeQ2                        ST_PivotFilterType = 52
	ST_PivotFilterTypeQ3                        ST_PivotFilterType = 53
	ST_PivotFilterTypeQ4                        ST_PivotFilterType = 54
	ST_PivotFilterTypeM1                        ST_PivotFilterType = 55
	ST_PivotFilterTypeM2                        ST_PivotFilterType = 56
	ST_PivotFilterTypeM3                        ST_PivotFilterType = 57
	ST_PivotFilterTypeM4                        ST_PivotFilterType = 58
	ST_PivotFilterTypeM5                        ST_PivotFilterType = 59
	ST_PivotFilterTypeM6                        ST_PivotFilterType = 60
	ST_PivotFilterTypeM7                        ST_PivotFilterType = 61
	ST_PivotFilterTypeM8                        ST_PivotFilterType = 62
	ST_PivotFilterTypeM9                        ST_PivotFilterType = 63
	ST_PivotFilterTypeM10                       ST_PivotFilterType = 64
	ST_PivotFilterTypeM11                       ST_PivotFilterType = 65
	ST_PivotFilterTypeM12                       ST_PivotFilterType = 66
)

func (e ST_PivotFilterType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PivotFilterTypeUnset:
		attr.Value = ""
	case ST_PivotFilterTypeUnknown:
		attr.Value = "unknown"
	case ST_PivotFilterTypeCount:
		attr.Value = "count"
	case ST_PivotFilterTypePercent:
		attr.Value = "percent"
	case ST_PivotFilterTypeSum:
		attr.Value = "sum"
	case ST_PivotFilterTypeCaptionEqual:
		attr.Value = "captionEqual"
	case ST_PivotFilterTypeCaptionNotEqual:
		attr.Value = "captionNotEqual"
	case ST_PivotFilterTypeCaptionBeginsWith:
		attr.Value = "captionBeginsWith"
	case ST_PivotFilterTypeCaptionNotBeginsWith:
		attr.Value = "captionNotBeginsWith"
	case ST_PivotFilterTypeCaptionEndsWith:
		attr.Value = "captionEndsWith"
	case ST_PivotFilterTypeCaptionNotEndsWith:
		attr.Value = "captionNotEndsWith"
	case ST_PivotFilterTypeCaptionContains:
		attr.Value = "captionContains"
	case ST_PivotFilterTypeCaptionNotContains:
		attr.Value = "captionNotContains"
	case ST_PivotFilterTypeCaptionGreaterThan:
		attr.Value = "captionGreaterThan"
	case ST_PivotFilterTypeCaptionGreaterThanOrEqual:
		attr.Value = "captionGreaterThanOrEqual"
	case ST_PivotFilterTypeCaptionLessThan:
		attr.Value = "captionLessThan"
	case ST_PivotFilterTypeCaptionLessThanOrEqual:
		attr.Value = "captionLessThanOrEqual"
	case ST_PivotFilterTypeCaptionBetween:
		attr.Value = "captionBetween"
	case ST_PivotFilterTypeCaptionNotBetween:
		attr.Value = "captionNotBetween"
	case ST_PivotFilterTypeValueEqual:
		attr.Value = "valueEqual"
	case ST_PivotFilterTypeValueNotEqual:
		attr.Value = "valueNotEqual"
	case ST_PivotFilterTypeValueGreaterThan:
		attr.Value = "valueGreaterThan"
	case ST_PivotFilterTypeValueGreaterThanOrEqual:
		attr.Value = "valueGreaterThanOrEqual"
	case ST_PivotFilterTypeValueLessThan:
		attr.Value = "valueLessThan"
	case ST_PivotFilterTypeValueLessThanOrEqual:
		attr.Value = "valueLessThanOrEqual"
	case ST_PivotFilterTypeValueBetween:
		attr.Value = "valueBetween"
	case ST_PivotFilterTypeValueNotBetween:
		attr.Value = "valueNotBetween"
	case ST_PivotFilterTypeDateEqual:
		attr.Value = "dateEqual"
	case ST_PivotFilterTypeDateNotEqual:
		attr.Value = "dateNotEqual"
	case ST_PivotFilterTypeDateOlderThan:
		attr.Value = "dateOlderThan"
	case ST_PivotFilterTypeDateOlderThanOrEqual:
		attr.Value = "dateOlderThanOrEqual"
	case ST_PivotFilterTypeDateNewerThan:
		attr.Value = "dateNewerThan"
	case ST_PivotFilterTypeDateNewerThanOrEqual:
		attr.Value = "dateNewerThanOrEqual"
	case ST_PivotFilterTypeDateBetween:
		attr.Value = "dateBetween"
	case ST_PivotFilterTypeDateNotBetween:
		attr.Value = "dateNotBetween"
	case ST_PivotFilterTypeTomorrow:
		attr.Value = "tomorrow"
	case ST_PivotFilterTypeToday:
		attr.Value = "today"
	case ST_PivotFilterTypeYesterday:
		attr.Value = "yesterday"
	case ST_PivotFilterTypeNextWeek:
		attr.Value = "nextWeek"
	case ST_PivotFilterTypeThisWeek:
		attr.Value = "thisWeek"
	case ST_PivotFilterTypeLastWeek:
		attr.Value = "lastWeek"
	case ST_PivotFilterTypeNextMonth:
		attr.Value = "nextMonth"
	case ST_PivotFilterTypeThisMonth:
		attr.Value = "thisMonth"
	case ST_PivotFilterTypeLastMonth:
		attr.Value = "lastMonth"
	case ST_PivotFilterTypeNextQuarter:
		attr.Value = "nextQuarter"
	case ST_PivotFilterTypeThisQuarter:
		attr.Value = "thisQuarter"
	case ST_PivotFilterTypeLastQuarter:
		attr.Value = "lastQuarter"
	case ST_PivotFilterTypeNextYear:
		attr.Value = "nextYear"
	case ST_PivotFilterTypeThisYear:
		attr.Value = "thisYear"
	case ST_PivotFilterTypeLastYear:
		attr.Value = "lastYear"
	case ST_PivotFilterTypeYearToDate:
		attr.Value = "yearToDate"
	case ST_PivotFilterTypeQ1:
		attr.Value = "Q1"
	case ST_PivotFilterTypeQ2:
		attr.Value = "Q2"
	case ST_PivotFilterTypeQ3:
		attr.Value = "Q3"
	case ST_PivotFilterTypeQ4:
		attr.Value = "Q4"
	case ST_PivotFilterTypeM1:
		attr.Value = "M1"
	case ST_PivotFilterTypeM2:
		attr.Value = "M2"
	case ST_PivotFilterTypeM3:
		attr.Value = "M3"
	case ST_PivotFilterTypeM4:
		attr.Value = "M4"
	case ST_PivotFilterTypeM5:
		attr.Value = "M5"
	case ST_PivotFilterTypeM6:
		attr.Value = "M6"
	case ST_PivotFilterTypeM7:
		attr.Value = "M7"
	case ST_PivotFilterTypeM8:
		attr.Value = "M8"
	case ST_PivotFilterTypeM9:
		attr.Value = "M9"
	case ST_PivotFilterTypeM10:
		attr.Value = "M10"
	case ST_PivotFilterTypeM11:
		attr.Value = "M11"
	case ST_PivotFilterTypeM12:
		attr.Value = "M12"
	}
	return attr, nil
}

func (e *ST_PivotFilterType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "unknown":
		*e = 1
	case "count":
		*e = 2
	case "percent":
		*e = 3
	case "sum":
		*e = 4
	case "captionEqual":
		*e = 5
	case "captionNotEqual":
		*e = 6
	case "captionBeginsWith":
		*e = 7
	case "captionNotBeginsWith":
		*e = 8
	case "captionEndsWith":
		*e = 9
	case "captionNotEndsWith":
		*e = 10
	case "captionContains":
		*e = 11
	case "captionNotContains":
		*e = 12
	case "captionGreaterThan":
		*e = 13
	case "captionGreaterThanOrEqual":
		*e = 14
	case "captionLessThan":
		*e = 15
	case "captionLessThanOrEqual":
		*e = 16
	case "captionBetween":
		*e = 17
	case "captionNotBetween":
		*e = 18
	case "valueEqual":
		*e = 19
	case "valueNotEqual":
		*e = 20
	case "valueGreaterThan":
		*e = 21
	case "valueGreaterThanOrEqual":
		*e = 22
	case "valueLessThan":
		*e = 23
	case "valueLessThanOrEqual":
		*e = 24
	case "valueBetween":
		*e = 25
	case "valueNotBetween":
		*e = 26
	case "dateEqual":
		*e = 27
	case "dateNotEqual":
		*e = 28
	case "dateOlderThan":
		*e = 29
	case "dateOlderThanOrEqual":
		*e = 30
	case "dateNewerThan":
		*e = 31
	case "dateNewerThanOrEqual":
		*e = 32
	case "dateBetween":
		*e = 33
	case "dateNotBetween":
		*e = 34
	case "tomorrow":
		*e = 35
	case "today":
		*e = 36
	case "yesterday":
		*e = 37
	case "nextWeek":
		*e = 38
	case "thisWeek":
		*e = 39
	case "lastWeek":
		*e = 40
	case "nextMonth":
		*e = 41
	case "thisMonth":
		*e = 42
	case "lastMonth":
		*e = 43
	case "nextQuarter":
		*e = 44
	case "thisQuarter":
		*e = 45
	case "lastQuarter":
		*e = 46
	case "nextYear":
		*e = 47
	case "thisYear":
		*e = 48
	case "lastYear":
		*e = 49
	case "yearToDate":
		*e = 50
	case "Q1":
		*e = 51
	case "Q2":
		*e = 52
	case "Q3":
		*e = 53
	case "Q4":
		*e = 54
	case "M1":
		*e = 55
	case "M2":
		*e = 56
	case "M3":
		*e = 57
	case "M4":
		*e = 58
	case "M5":
		*e = 59
	case "M6":
		*e = 60
	case "M7":
		*e = 61
	case "M8":
		*e = 62
	case "M9":
		*e = 63
	case "M10":
		*e = 64
	case "M11":
		*e = 65
	case "M12":
		*e = 66
	}
	return nil
}

func (m ST_PivotFilterType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PivotFilterType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "unknown":
			*m = 1
		case "count":
			*m = 2
		case "percent":
			*m = 3
		case "sum":
			*m = 4
		case "captionEqual":
			*m = 5
		case "captionNotEqual":
			*m = 6
		case "captionBeginsWith":
			*m = 7
		case "captionNotBeginsWith":
			*m = 8
		case "captionEndsWith":
			*m = 9
		case "captionNotEndsWith":
			*m = 10
		case "captionContains":
			*m = 11
		case "captionNotContains":
			*m = 12
		case "captionGreaterThan":
			*m = 13
		case "captionGreaterThanOrEqual":
			*m = 14
		case "captionLessThan":
			*m = 15
		case "captionLessThanOrEqual":
			*m = 16
		case "captionBetween":
			*m = 17
		case "captionNotBetween":
			*m = 18
		case "valueEqual":
			*m = 19
		case "valueNotEqual":
			*m = 20
		case "valueGreaterThan":
			*m = 21
		case "valueGreaterThanOrEqual":
			*m = 22
		case "valueLessThan":
			*m = 23
		case "valueLessThanOrEqual":
			*m = 24
		case "valueBetween":
			*m = 25
		case "valueNotBetween":
			*m = 26
		case "dateEqual":
			*m = 27
		case "dateNotEqual":
			*m = 28
		case "dateOlderThan":
			*m = 29
		case "dateOlderThanOrEqual":
			*m = 30
		case "dateNewerThan":
			*m = 31
		case "dateNewerThanOrEqual":
			*m = 32
		case "dateBetween":
			*m = 33
		case "dateNotBetween":
			*m = 34
		case "tomorrow":
			*m = 35
		case "today":
			*m = 36
		case "yesterday":
			*m = 37
		case "nextWeek":
			*m = 38
		case "thisWeek":
			*m = 39
		case "lastWeek":
			*m = 40
		case "nextMonth":
			*m = 41
		case "thisMonth":
			*m = 42
		case "lastMonth":
			*m = 43
		case "nextQuarter":
			*m = 44
		case "thisQuarter":
			*m = 45
		case "lastQuarter":
			*m = 46
		case "nextYear":
			*m = 47
		case "thisYear":
			*m = 48
		case "lastYear":
			*m = 49
		case "yearToDate":
			*m = 50
		case "Q1":
			*m = 51
		case "Q2":
			*m = 52
		case "Q3":
			*m = 53
		case "Q4":
			*m = 54
		case "M1":
			*m = 55
		case "M2":
			*m = 56
		case "M3":
			*m = 57
		case "M4":
			*m = 58
		case "M5":
			*m = 59
		case "M6":
			*m = 60
		case "M7":
			*m = 61
		case "M8":
			*m = 62
		case "M9":
			*m = 63
		case "M10":
			*m = 64
		case "M11":
			*m = 65
		case "M12":
			*m = 66
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

func (m ST_PivotFilterType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "unknown"
	case 2:
		return "count"
	case 3:
		return "percent"
	case 4:
		return "sum"
	case 5:
		return "captionEqual"
	case 6:
		return "captionNotEqual"
	case 7:
		return "captionBeginsWith"
	case 8:
		return "captionNotBeginsWith"
	case 9:
		return "captionEndsWith"
	case 10:
		return "captionNotEndsWith"
	case 11:
		return "captionContains"
	case 12:
		return "captionNotContains"
	case 13:
		return "captionGreaterThan"
	case 14:
		return "captionGreaterThanOrEqual"
	case 15:
		return "captionLessThan"
	case 16:
		return "captionLessThanOrEqual"
	case 17:
		return "captionBetween"
	case 18:
		return "captionNotBetween"
	case 19:
		return "valueEqual"
	case 20:
		return "valueNotEqual"
	case 21:
		return "valueGreaterThan"
	case 22:
		return "valueGreaterThanOrEqual"
	case 23:
		return "valueLessThan"
	case 24:
		return "valueLessThanOrEqual"
	case 25:
		return "valueBetween"
	case 26:
		return "valueNotBetween"
	case 27:
		return "dateEqual"
	case 28:
		return "dateNotEqual"
	case 29:
		return "dateOlderThan"
	case 30:
		return "dateOlderThanOrEqual"
	case 31:
		return "dateNewerThan"
	case 32:
		return "dateNewerThanOrEqual"
	case 33:
		return "dateBetween"
	case 34:
		return "dateNotBetween"
	case 35:
		return "tomorrow"
	case 36:
		return "today"
	case 37:
		return "yesterday"
	case 38:
		return "nextWeek"
	case 39:
		return "thisWeek"
	case 40:
		return "lastWeek"
	case 41:
		return "nextMonth"
	case 42:
		return "thisMonth"
	case 43:
		return "lastMonth"
	case 44:
		return "nextQuarter"
	case 45:
		return "thisQuarter"
	case 46:
		return "lastQuarter"
	case 47:
		return "nextYear"
	case 48:
		return "thisYear"
	case 49:
		return "lastYear"
	case 50:
		return "yearToDate"
	case 51:
		return "Q1"
	case 52:
		return "Q2"
	case 53:
		return "Q3"
	case 54:
		return "Q4"
	case 55:
		return "M1"
	case 56:
		return "M2"
	case 57:
		return "M3"
	case 58:
		return "M4"
	case 59:
		return "M5"
	case 60:
		return "M6"
	case 61:
		return "M7"
	case 62:
		return "M8"
	case 63:
		return "M9"
	case 64:
		return "M10"
	case 65:
		return "M11"
	case 66:
		return "M12"
	}
	return ""
}

func (m ST_PivotFilterType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PivotFilterType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PivotAreaType byte

const (
	ST_PivotAreaTypeUnset    ST_PivotAreaType = 0
	ST_PivotAreaTypeNone     ST_PivotAreaType = 1
	ST_PivotAreaTypeNormal   ST_PivotAreaType = 2
	ST_PivotAreaTypeData     ST_PivotAreaType = 3
	ST_PivotAreaTypeAll      ST_PivotAreaType = 4
	ST_PivotAreaTypeOrigin   ST_PivotAreaType = 5
	ST_PivotAreaTypeButton   ST_PivotAreaType = 6
	ST_PivotAreaTypeTopEnd   ST_PivotAreaType = 7
	ST_PivotAreaTypeTopRight ST_PivotAreaType = 8
)

func (e ST_PivotAreaType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PivotAreaTypeUnset:
		attr.Value = ""
	case ST_PivotAreaTypeNone:
		attr.Value = "none"
	case ST_PivotAreaTypeNormal:
		attr.Value = "normal"
	case ST_PivotAreaTypeData:
		attr.Value = "data"
	case ST_PivotAreaTypeAll:
		attr.Value = "all"
	case ST_PivotAreaTypeOrigin:
		attr.Value = "origin"
	case ST_PivotAreaTypeButton:
		attr.Value = "button"
	case ST_PivotAreaTypeTopEnd:
		attr.Value = "topEnd"
	case ST_PivotAreaTypeTopRight:
		attr.Value = "topRight"
	}
	return attr, nil
}

func (e *ST_PivotAreaType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "normal":
		*e = 2
	case "data":
		*e = 3
	case "all":
		*e = 4
	case "origin":
		*e = 5
	case "button":
		*e = 6
	case "topEnd":
		*e = 7
	case "topRight":
		*e = 8
	}
	return nil
}

func (m ST_PivotAreaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PivotAreaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "normal":
			*m = 2
		case "data":
			*m = 3
		case "all":
			*m = 4
		case "origin":
			*m = 5
		case "button":
			*m = 6
		case "topEnd":
			*m = 7
		case "topRight":
			*m = 8
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

func (m ST_PivotAreaType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "normal"
	case 3:
		return "data"
	case 4:
		return "all"
	case 5:
		return "origin"
	case 6:
		return "button"
	case 7:
		return "topEnd"
	case 8:
		return "topRight"
	}
	return ""
}

func (m ST_PivotAreaType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PivotAreaType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Axis byte

const (
	ST_AxisUnset      ST_Axis = 0
	ST_AxisAxisRow    ST_Axis = 1
	ST_AxisAxisCol    ST_Axis = 2
	ST_AxisAxisPage   ST_Axis = 3
	ST_AxisAxisValues ST_Axis = 4
)

func (e ST_Axis) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_AxisUnset:
		attr.Value = ""
	case ST_AxisAxisRow:
		attr.Value = "axisRow"
	case ST_AxisAxisCol:
		attr.Value = "axisCol"
	case ST_AxisAxisPage:
		attr.Value = "axisPage"
	case ST_AxisAxisValues:
		attr.Value = "axisValues"
	}
	return attr, nil
}

func (e *ST_Axis) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "axisRow":
		*e = 1
	case "axisCol":
		*e = 2
	case "axisPage":
		*e = 3
	case "axisValues":
		*e = 4
	}
	return nil
}

func (m ST_Axis) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Axis) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "axisRow":
			*m = 1
		case "axisCol":
			*m = 2
		case "axisPage":
			*m = 3
		case "axisValues":
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

func (m ST_Axis) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "axisRow"
	case 2:
		return "axisCol"
	case 3:
		return "axisPage"
	case 4:
		return "axisValues"
	}
	return ""
}

func (m ST_Axis) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Axis) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_GrowShrinkType byte

const (
	ST_GrowShrinkTypeUnset          ST_GrowShrinkType = 0
	ST_GrowShrinkTypeInsertDelete   ST_GrowShrinkType = 1
	ST_GrowShrinkTypeInsertClear    ST_GrowShrinkType = 2
	ST_GrowShrinkTypeOverwriteClear ST_GrowShrinkType = 3
)

func (e ST_GrowShrinkType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_GrowShrinkTypeUnset:
		attr.Value = ""
	case ST_GrowShrinkTypeInsertDelete:
		attr.Value = "insertDelete"
	case ST_GrowShrinkTypeInsertClear:
		attr.Value = "insertClear"
	case ST_GrowShrinkTypeOverwriteClear:
		attr.Value = "overwriteClear"
	}
	return attr, nil
}

func (e *ST_GrowShrinkType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "insertDelete":
		*e = 1
	case "insertClear":
		*e = 2
	case "overwriteClear":
		*e = 3
	}
	return nil
}

func (m ST_GrowShrinkType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_GrowShrinkType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "insertDelete":
			*m = 1
		case "insertClear":
			*m = 2
		case "overwriteClear":
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

func (m ST_GrowShrinkType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "insertDelete"
	case 2:
		return "insertClear"
	case 3:
		return "overwriteClear"
	}
	return ""
}

func (m ST_GrowShrinkType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_GrowShrinkType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PhoneticType byte

const (
	ST_PhoneticTypeUnset             ST_PhoneticType = 0
	ST_PhoneticTypeHalfwidthKatakana ST_PhoneticType = 1
	ST_PhoneticTypeFullwidthKatakana ST_PhoneticType = 2
	ST_PhoneticTypeHiragana          ST_PhoneticType = 3
	ST_PhoneticTypeNoConversion      ST_PhoneticType = 4
)

func (e ST_PhoneticType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PhoneticTypeUnset:
		attr.Value = ""
	case ST_PhoneticTypeHalfwidthKatakana:
		attr.Value = "halfwidthKatakana"
	case ST_PhoneticTypeFullwidthKatakana:
		attr.Value = "fullwidthKatakana"
	case ST_PhoneticTypeHiragana:
		attr.Value = "Hiragana"
	case ST_PhoneticTypeNoConversion:
		attr.Value = "noConversion"
	}
	return attr, nil
}

func (e *ST_PhoneticType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "halfwidthKatakana":
		*e = 1
	case "fullwidthKatakana":
		*e = 2
	case "Hiragana":
		*e = 3
	case "noConversion":
		*e = 4
	}
	return nil
}

func (m ST_PhoneticType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PhoneticType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "halfwidthKatakana":
			*m = 1
		case "fullwidthKatakana":
			*m = 2
		case "Hiragana":
			*m = 3
		case "noConversion":
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

func (m ST_PhoneticType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "halfwidthKatakana"
	case 2:
		return "fullwidthKatakana"
	case 3:
		return "Hiragana"
	case 4:
		return "noConversion"
	}
	return ""
}

func (m ST_PhoneticType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PhoneticType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PhoneticAlignment byte

const (
	ST_PhoneticAlignmentUnset       ST_PhoneticAlignment = 0
	ST_PhoneticAlignmentNoControl   ST_PhoneticAlignment = 1
	ST_PhoneticAlignmentLeft        ST_PhoneticAlignment = 2
	ST_PhoneticAlignmentCenter      ST_PhoneticAlignment = 3
	ST_PhoneticAlignmentDistributed ST_PhoneticAlignment = 4
)

func (e ST_PhoneticAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PhoneticAlignmentUnset:
		attr.Value = ""
	case ST_PhoneticAlignmentNoControl:
		attr.Value = "noControl"
	case ST_PhoneticAlignmentLeft:
		attr.Value = "left"
	case ST_PhoneticAlignmentCenter:
		attr.Value = "center"
	case ST_PhoneticAlignmentDistributed:
		attr.Value = "distributed"
	}
	return attr, nil
}

func (e *ST_PhoneticAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "noControl":
		*e = 1
	case "left":
		*e = 2
	case "center":
		*e = 3
	case "distributed":
		*e = 4
	}
	return nil
}

func (m ST_PhoneticAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PhoneticAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "noControl":
			*m = 1
		case "left":
			*m = 2
		case "center":
			*m = 3
		case "distributed":
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

func (m ST_PhoneticAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "noControl"
	case 2:
		return "left"
	case 3:
		return "center"
	case 4:
		return "distributed"
	}
	return ""
}

func (m ST_PhoneticAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PhoneticAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_rwColActionType byte

const (
	ST_rwColActionTypeUnset     ST_rwColActionType = 0
	ST_rwColActionTypeInsertRow ST_rwColActionType = 1
	ST_rwColActionTypeDeleteRow ST_rwColActionType = 2
	ST_rwColActionTypeInsertCol ST_rwColActionType = 3
	ST_rwColActionTypeDeleteCol ST_rwColActionType = 4
)

func (e ST_rwColActionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_rwColActionTypeUnset:
		attr.Value = ""
	case ST_rwColActionTypeInsertRow:
		attr.Value = "insertRow"
	case ST_rwColActionTypeDeleteRow:
		attr.Value = "deleteRow"
	case ST_rwColActionTypeInsertCol:
		attr.Value = "insertCol"
	case ST_rwColActionTypeDeleteCol:
		attr.Value = "deleteCol"
	}
	return attr, nil
}

func (e *ST_rwColActionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "insertRow":
		*e = 1
	case "deleteRow":
		*e = 2
	case "insertCol":
		*e = 3
	case "deleteCol":
		*e = 4
	}
	return nil
}

func (m ST_rwColActionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_rwColActionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "insertRow":
			*m = 1
		case "deleteRow":
			*m = 2
		case "insertCol":
			*m = 3
		case "deleteCol":
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

func (m ST_rwColActionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "insertRow"
	case 2:
		return "deleteRow"
	case 3:
		return "insertCol"
	case 4:
		return "deleteCol"
	}
	return ""
}

func (m ST_rwColActionType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_rwColActionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_RevisionAction byte

const (
	ST_RevisionActionUnset  ST_RevisionAction = 0
	ST_RevisionActionAdd    ST_RevisionAction = 1
	ST_RevisionActionDelete ST_RevisionAction = 2
)

func (e ST_RevisionAction) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_RevisionActionUnset:
		attr.Value = ""
	case ST_RevisionActionAdd:
		attr.Value = "add"
	case ST_RevisionActionDelete:
		attr.Value = "delete"
	}
	return attr, nil
}

func (e *ST_RevisionAction) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "add":
		*e = 1
	case "delete":
		*e = 2
	}
	return nil
}

func (m ST_RevisionAction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_RevisionAction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "add":
			*m = 1
		case "delete":
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

func (m ST_RevisionAction) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "add"
	case 2:
		return "delete"
	}
	return ""
}

func (m ST_RevisionAction) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_RevisionAction) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FormulaExpression byte

const (
	ST_FormulaExpressionUnset        ST_FormulaExpression = 0
	ST_FormulaExpressionRef          ST_FormulaExpression = 1
	ST_FormulaExpressionRefError     ST_FormulaExpression = 2
	ST_FormulaExpressionArea         ST_FormulaExpression = 3
	ST_FormulaExpressionAreaError    ST_FormulaExpression = 4
	ST_FormulaExpressionComputedArea ST_FormulaExpression = 5
)

func (e ST_FormulaExpression) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FormulaExpressionUnset:
		attr.Value = ""
	case ST_FormulaExpressionRef:
		attr.Value = "ref"
	case ST_FormulaExpressionRefError:
		attr.Value = "refError"
	case ST_FormulaExpressionArea:
		attr.Value = "area"
	case ST_FormulaExpressionAreaError:
		attr.Value = "areaError"
	case ST_FormulaExpressionComputedArea:
		attr.Value = "computedArea"
	}
	return attr, nil
}

func (e *ST_FormulaExpression) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "ref":
		*e = 1
	case "refError":
		*e = 2
	case "area":
		*e = 3
	case "areaError":
		*e = 4
	case "computedArea":
		*e = 5
	}
	return nil
}

func (m ST_FormulaExpression) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FormulaExpression) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "ref":
			*m = 1
		case "refError":
			*m = 2
		case "area":
			*m = 3
		case "areaError":
			*m = 4
		case "computedArea":
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

func (m ST_FormulaExpression) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "ref"
	case 2:
		return "refError"
	case 3:
		return "area"
	case 4:
		return "areaError"
	case 5:
		return "computedArea"
	}
	return ""
}

func (m ST_FormulaExpression) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FormulaExpression) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CellType byte

const (
	ST_CellTypeUnset     ST_CellType = 0
	ST_CellTypeB         ST_CellType = 1
	ST_CellTypeN         ST_CellType = 2
	ST_CellTypeE         ST_CellType = 3
	ST_CellTypeS         ST_CellType = 4
	ST_CellTypeStr       ST_CellType = 5
	ST_CellTypeInlineStr ST_CellType = 6
)

func (e ST_CellType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CellTypeUnset:
		attr.Value = ""
	case ST_CellTypeB:
		attr.Value = "b"
	case ST_CellTypeN:
		attr.Value = "n"
	case ST_CellTypeE:
		attr.Value = "e"
	case ST_CellTypeS:
		attr.Value = "s"
	case ST_CellTypeStr:
		attr.Value = "str"
	case ST_CellTypeInlineStr:
		attr.Value = "inlineStr"
	}
	return attr, nil
}

func (e *ST_CellType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "b":
		*e = 1
	case "n":
		*e = 2
	case "e":
		*e = 3
	case "s":
		*e = 4
	case "str":
		*e = 5
	case "inlineStr":
		*e = 6
	}
	return nil
}

func (m ST_CellType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CellType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "b":
			*m = 1
		case "n":
			*m = 2
		case "e":
			*m = 3
		case "s":
			*m = 4
		case "str":
			*m = 5
		case "inlineStr":
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

func (m ST_CellType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "b"
	case 2:
		return "n"
	case 3:
		return "e"
	case 4:
		return "s"
	case 5:
		return "str"
	case 6:
		return "inlineStr"
	}
	return ""
}

func (m ST_CellType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CellType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CellFormulaType byte

const (
	ST_CellFormulaTypeUnset     ST_CellFormulaType = 0
	ST_CellFormulaTypeNormal    ST_CellFormulaType = 1
	ST_CellFormulaTypeArray     ST_CellFormulaType = 2
	ST_CellFormulaTypeDataTable ST_CellFormulaType = 3
	ST_CellFormulaTypeShared    ST_CellFormulaType = 4
)

func (e ST_CellFormulaType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CellFormulaTypeUnset:
		attr.Value = ""
	case ST_CellFormulaTypeNormal:
		attr.Value = "normal"
	case ST_CellFormulaTypeArray:
		attr.Value = "array"
	case ST_CellFormulaTypeDataTable:
		attr.Value = "dataTable"
	case ST_CellFormulaTypeShared:
		attr.Value = "shared"
	}
	return attr, nil
}

func (e *ST_CellFormulaType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "normal":
		*e = 1
	case "array":
		*e = 2
	case "dataTable":
		*e = 3
	case "shared":
		*e = 4
	}
	return nil
}

func (m ST_CellFormulaType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CellFormulaType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "normal":
			*m = 1
		case "array":
			*m = 2
		case "dataTable":
			*m = 3
		case "shared":
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

func (m ST_CellFormulaType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "normal"
	case 2:
		return "array"
	case 3:
		return "dataTable"
	case 4:
		return "shared"
	}
	return ""
}

func (m ST_CellFormulaType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CellFormulaType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Pane byte

const (
	ST_PaneUnset       ST_Pane = 0
	ST_PaneBottomRight ST_Pane = 1
	ST_PaneTopRight    ST_Pane = 2
	ST_PaneBottomLeft  ST_Pane = 3
	ST_PaneTopLeft     ST_Pane = 4
)

func (e ST_Pane) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PaneUnset:
		attr.Value = ""
	case ST_PaneBottomRight:
		attr.Value = "bottomRight"
	case ST_PaneTopRight:
		attr.Value = "topRight"
	case ST_PaneBottomLeft:
		attr.Value = "bottomLeft"
	case ST_PaneTopLeft:
		attr.Value = "topLeft"
	}
	return attr, nil
}

func (e *ST_Pane) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "bottomRight":
		*e = 1
	case "topRight":
		*e = 2
	case "bottomLeft":
		*e = 3
	case "topLeft":
		*e = 4
	}
	return nil
}

func (m ST_Pane) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Pane) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "bottomRight":
			*m = 1
		case "topRight":
			*m = 2
		case "bottomLeft":
			*m = 3
		case "topLeft":
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

func (m ST_Pane) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "bottomRight"
	case 2:
		return "topRight"
	case 3:
		return "bottomLeft"
	case 4:
		return "topLeft"
	}
	return ""
}

func (m ST_Pane) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Pane) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SheetViewType byte

const (
	ST_SheetViewTypeUnset            ST_SheetViewType = 0
	ST_SheetViewTypeNormal           ST_SheetViewType = 1
	ST_SheetViewTypePageBreakPreview ST_SheetViewType = 2
	ST_SheetViewTypePageLayout       ST_SheetViewType = 3
)

func (e ST_SheetViewType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SheetViewTypeUnset:
		attr.Value = ""
	case ST_SheetViewTypeNormal:
		attr.Value = "normal"
	case ST_SheetViewTypePageBreakPreview:
		attr.Value = "pageBreakPreview"
	case ST_SheetViewTypePageLayout:
		attr.Value = "pageLayout"
	}
	return attr, nil
}

func (e *ST_SheetViewType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "normal":
		*e = 1
	case "pageBreakPreview":
		*e = 2
	case "pageLayout":
		*e = 3
	}
	return nil
}

func (m ST_SheetViewType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SheetViewType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "normal":
			*m = 1
		case "pageBreakPreview":
			*m = 2
		case "pageLayout":
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

func (m ST_SheetViewType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "normal"
	case 2:
		return "pageBreakPreview"
	case 3:
		return "pageLayout"
	}
	return ""
}

func (m ST_SheetViewType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SheetViewType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DataConsolidateFunction byte

const (
	ST_DataConsolidateFunctionUnset     ST_DataConsolidateFunction = 0
	ST_DataConsolidateFunctionAverage   ST_DataConsolidateFunction = 1
	ST_DataConsolidateFunctionCount     ST_DataConsolidateFunction = 2
	ST_DataConsolidateFunctionCountNums ST_DataConsolidateFunction = 3
	ST_DataConsolidateFunctionMax       ST_DataConsolidateFunction = 4
	ST_DataConsolidateFunctionMin       ST_DataConsolidateFunction = 5
	ST_DataConsolidateFunctionProduct   ST_DataConsolidateFunction = 6
	ST_DataConsolidateFunctionStdDev    ST_DataConsolidateFunction = 7
	ST_DataConsolidateFunctionStdDevp   ST_DataConsolidateFunction = 8
	ST_DataConsolidateFunctionSum       ST_DataConsolidateFunction = 9
	ST_DataConsolidateFunctionVar       ST_DataConsolidateFunction = 10
	ST_DataConsolidateFunctionVarp      ST_DataConsolidateFunction = 11
)

func (e ST_DataConsolidateFunction) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DataConsolidateFunctionUnset:
		attr.Value = ""
	case ST_DataConsolidateFunctionAverage:
		attr.Value = "average"
	case ST_DataConsolidateFunctionCount:
		attr.Value = "count"
	case ST_DataConsolidateFunctionCountNums:
		attr.Value = "countNums"
	case ST_DataConsolidateFunctionMax:
		attr.Value = "max"
	case ST_DataConsolidateFunctionMin:
		attr.Value = "min"
	case ST_DataConsolidateFunctionProduct:
		attr.Value = "product"
	case ST_DataConsolidateFunctionStdDev:
		attr.Value = "stdDev"
	case ST_DataConsolidateFunctionStdDevp:
		attr.Value = "stdDevp"
	case ST_DataConsolidateFunctionSum:
		attr.Value = "sum"
	case ST_DataConsolidateFunctionVar:
		attr.Value = "var"
	case ST_DataConsolidateFunctionVarp:
		attr.Value = "varp"
	}
	return attr, nil
}

func (e *ST_DataConsolidateFunction) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "average":
		*e = 1
	case "count":
		*e = 2
	case "countNums":
		*e = 3
	case "max":
		*e = 4
	case "min":
		*e = 5
	case "product":
		*e = 6
	case "stdDev":
		*e = 7
	case "stdDevp":
		*e = 8
	case "sum":
		*e = 9
	case "var":
		*e = 10
	case "varp":
		*e = 11
	}
	return nil
}

func (m ST_DataConsolidateFunction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DataConsolidateFunction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "average":
			*m = 1
		case "count":
			*m = 2
		case "countNums":
			*m = 3
		case "max":
			*m = 4
		case "min":
			*m = 5
		case "product":
			*m = 6
		case "stdDev":
			*m = 7
		case "stdDevp":
			*m = 8
		case "sum":
			*m = 9
		case "var":
			*m = 10
		case "varp":
			*m = 11
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

func (m ST_DataConsolidateFunction) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "average"
	case 2:
		return "count"
	case 3:
		return "countNums"
	case 4:
		return "max"
	case 5:
		return "min"
	case 6:
		return "product"
	case 7:
		return "stdDev"
	case 8:
		return "stdDevp"
	case 9:
		return "sum"
	case 10:
		return "var"
	case 11:
		return "varp"
	}
	return ""
}

func (m ST_DataConsolidateFunction) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DataConsolidateFunction) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DataValidationType byte

const (
	ST_DataValidationTypeUnset      ST_DataValidationType = 0
	ST_DataValidationTypeNone       ST_DataValidationType = 1
	ST_DataValidationTypeWhole      ST_DataValidationType = 2
	ST_DataValidationTypeDecimal    ST_DataValidationType = 3
	ST_DataValidationTypeList       ST_DataValidationType = 4
	ST_DataValidationTypeDate       ST_DataValidationType = 5
	ST_DataValidationTypeTime       ST_DataValidationType = 6
	ST_DataValidationTypeTextLength ST_DataValidationType = 7
	ST_DataValidationTypeCustom     ST_DataValidationType = 8
)

func (e ST_DataValidationType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DataValidationTypeUnset:
		attr.Value = ""
	case ST_DataValidationTypeNone:
		attr.Value = "none"
	case ST_DataValidationTypeWhole:
		attr.Value = "whole"
	case ST_DataValidationTypeDecimal:
		attr.Value = "decimal"
	case ST_DataValidationTypeList:
		attr.Value = "list"
	case ST_DataValidationTypeDate:
		attr.Value = "date"
	case ST_DataValidationTypeTime:
		attr.Value = "time"
	case ST_DataValidationTypeTextLength:
		attr.Value = "textLength"
	case ST_DataValidationTypeCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *ST_DataValidationType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "whole":
		*e = 2
	case "decimal":
		*e = 3
	case "list":
		*e = 4
	case "date":
		*e = 5
	case "time":
		*e = 6
	case "textLength":
		*e = 7
	case "custom":
		*e = 8
	}
	return nil
}

func (m ST_DataValidationType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DataValidationType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "whole":
			*m = 2
		case "decimal":
			*m = 3
		case "list":
			*m = 4
		case "date":
			*m = 5
		case "time":
			*m = 6
		case "textLength":
			*m = 7
		case "custom":
			*m = 8
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

func (m ST_DataValidationType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "whole"
	case 3:
		return "decimal"
	case 4:
		return "list"
	case 5:
		return "date"
	case 6:
		return "time"
	case 7:
		return "textLength"
	case 8:
		return "custom"
	}
	return ""
}

func (m ST_DataValidationType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DataValidationType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DataValidationOperator byte

const (
	ST_DataValidationOperatorUnset              ST_DataValidationOperator = 0
	ST_DataValidationOperatorBetween            ST_DataValidationOperator = 1
	ST_DataValidationOperatorNotBetween         ST_DataValidationOperator = 2
	ST_DataValidationOperatorEqual              ST_DataValidationOperator = 3
	ST_DataValidationOperatorNotEqual           ST_DataValidationOperator = 4
	ST_DataValidationOperatorLessThan           ST_DataValidationOperator = 5
	ST_DataValidationOperatorLessThanOrEqual    ST_DataValidationOperator = 6
	ST_DataValidationOperatorGreaterThan        ST_DataValidationOperator = 7
	ST_DataValidationOperatorGreaterThanOrEqual ST_DataValidationOperator = 8
)

func (e ST_DataValidationOperator) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DataValidationOperatorUnset:
		attr.Value = ""
	case ST_DataValidationOperatorBetween:
		attr.Value = "between"
	case ST_DataValidationOperatorNotBetween:
		attr.Value = "notBetween"
	case ST_DataValidationOperatorEqual:
		attr.Value = "equal"
	case ST_DataValidationOperatorNotEqual:
		attr.Value = "notEqual"
	case ST_DataValidationOperatorLessThan:
		attr.Value = "lessThan"
	case ST_DataValidationOperatorLessThanOrEqual:
		attr.Value = "lessThanOrEqual"
	case ST_DataValidationOperatorGreaterThan:
		attr.Value = "greaterThan"
	case ST_DataValidationOperatorGreaterThanOrEqual:
		attr.Value = "greaterThanOrEqual"
	}
	return attr, nil
}

func (e *ST_DataValidationOperator) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "between":
		*e = 1
	case "notBetween":
		*e = 2
	case "equal":
		*e = 3
	case "notEqual":
		*e = 4
	case "lessThan":
		*e = 5
	case "lessThanOrEqual":
		*e = 6
	case "greaterThan":
		*e = 7
	case "greaterThanOrEqual":
		*e = 8
	}
	return nil
}

func (m ST_DataValidationOperator) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DataValidationOperator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "between":
			*m = 1
		case "notBetween":
			*m = 2
		case "equal":
			*m = 3
		case "notEqual":
			*m = 4
		case "lessThan":
			*m = 5
		case "lessThanOrEqual":
			*m = 6
		case "greaterThan":
			*m = 7
		case "greaterThanOrEqual":
			*m = 8
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

func (m ST_DataValidationOperator) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "between"
	case 2:
		return "notBetween"
	case 3:
		return "equal"
	case 4:
		return "notEqual"
	case 5:
		return "lessThan"
	case 6:
		return "lessThanOrEqual"
	case 7:
		return "greaterThan"
	case 8:
		return "greaterThanOrEqual"
	}
	return ""
}

func (m ST_DataValidationOperator) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DataValidationOperator) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DataValidationErrorStyle byte

const (
	ST_DataValidationErrorStyleUnset       ST_DataValidationErrorStyle = 0
	ST_DataValidationErrorStyleStop        ST_DataValidationErrorStyle = 1
	ST_DataValidationErrorStyleWarning     ST_DataValidationErrorStyle = 2
	ST_DataValidationErrorStyleInformation ST_DataValidationErrorStyle = 3
)

func (e ST_DataValidationErrorStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DataValidationErrorStyleUnset:
		attr.Value = ""
	case ST_DataValidationErrorStyleStop:
		attr.Value = "stop"
	case ST_DataValidationErrorStyleWarning:
		attr.Value = "warning"
	case ST_DataValidationErrorStyleInformation:
		attr.Value = "information"
	}
	return attr, nil
}

func (e *ST_DataValidationErrorStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "stop":
		*e = 1
	case "warning":
		*e = 2
	case "information":
		*e = 3
	}
	return nil
}

func (m ST_DataValidationErrorStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DataValidationErrorStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "stop":
			*m = 1
		case "warning":
			*m = 2
		case "information":
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

func (m ST_DataValidationErrorStyle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "stop"
	case 2:
		return "warning"
	case 3:
		return "information"
	}
	return ""
}

func (m ST_DataValidationErrorStyle) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DataValidationErrorStyle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DataValidationImeMode byte

const (
	ST_DataValidationImeModeUnset        ST_DataValidationImeMode = 0
	ST_DataValidationImeModeNoControl    ST_DataValidationImeMode = 1
	ST_DataValidationImeModeOff          ST_DataValidationImeMode = 2
	ST_DataValidationImeModeOn           ST_DataValidationImeMode = 3
	ST_DataValidationImeModeDisabled     ST_DataValidationImeMode = 4
	ST_DataValidationImeModeHiragana     ST_DataValidationImeMode = 5
	ST_DataValidationImeModeFullKatakana ST_DataValidationImeMode = 6
	ST_DataValidationImeModeHalfKatakana ST_DataValidationImeMode = 7
	ST_DataValidationImeModeFullAlpha    ST_DataValidationImeMode = 8
	ST_DataValidationImeModeHalfAlpha    ST_DataValidationImeMode = 9
	ST_DataValidationImeModeFullHangul   ST_DataValidationImeMode = 10
	ST_DataValidationImeModeHalfHangul   ST_DataValidationImeMode = 11
)

func (e ST_DataValidationImeMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DataValidationImeModeUnset:
		attr.Value = ""
	case ST_DataValidationImeModeNoControl:
		attr.Value = "noControl"
	case ST_DataValidationImeModeOff:
		attr.Value = "off"
	case ST_DataValidationImeModeOn:
		attr.Value = "on"
	case ST_DataValidationImeModeDisabled:
		attr.Value = "disabled"
	case ST_DataValidationImeModeHiragana:
		attr.Value = "hiragana"
	case ST_DataValidationImeModeFullKatakana:
		attr.Value = "fullKatakana"
	case ST_DataValidationImeModeHalfKatakana:
		attr.Value = "halfKatakana"
	case ST_DataValidationImeModeFullAlpha:
		attr.Value = "fullAlpha"
	case ST_DataValidationImeModeHalfAlpha:
		attr.Value = "halfAlpha"
	case ST_DataValidationImeModeFullHangul:
		attr.Value = "fullHangul"
	case ST_DataValidationImeModeHalfHangul:
		attr.Value = "halfHangul"
	}
	return attr, nil
}

func (e *ST_DataValidationImeMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "noControl":
		*e = 1
	case "off":
		*e = 2
	case "on":
		*e = 3
	case "disabled":
		*e = 4
	case "hiragana":
		*e = 5
	case "fullKatakana":
		*e = 6
	case "halfKatakana":
		*e = 7
	case "fullAlpha":
		*e = 8
	case "halfAlpha":
		*e = 9
	case "fullHangul":
		*e = 10
	case "halfHangul":
		*e = 11
	}
	return nil
}

func (m ST_DataValidationImeMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DataValidationImeMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "noControl":
			*m = 1
		case "off":
			*m = 2
		case "on":
			*m = 3
		case "disabled":
			*m = 4
		case "hiragana":
			*m = 5
		case "fullKatakana":
			*m = 6
		case "halfKatakana":
			*m = 7
		case "fullAlpha":
			*m = 8
		case "halfAlpha":
			*m = 9
		case "fullHangul":
			*m = 10
		case "halfHangul":
			*m = 11
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

func (m ST_DataValidationImeMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "noControl"
	case 2:
		return "off"
	case 3:
		return "on"
	case 4:
		return "disabled"
	case 5:
		return "hiragana"
	case 6:
		return "fullKatakana"
	case 7:
		return "halfKatakana"
	case 8:
		return "fullAlpha"
	case 9:
		return "halfAlpha"
	case 10:
		return "fullHangul"
	case 11:
		return "halfHangul"
	}
	return ""
}

func (m ST_DataValidationImeMode) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DataValidationImeMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CfType byte

const (
	ST_CfTypeUnset             ST_CfType = 0
	ST_CfTypeExpression        ST_CfType = 1
	ST_CfTypeCellIs            ST_CfType = 2
	ST_CfTypeColorScale        ST_CfType = 3
	ST_CfTypeDataBar           ST_CfType = 4
	ST_CfTypeIconSet           ST_CfType = 5
	ST_CfTypeTop10             ST_CfType = 6
	ST_CfTypeUniqueValues      ST_CfType = 7
	ST_CfTypeDuplicateValues   ST_CfType = 8
	ST_CfTypeContainsText      ST_CfType = 9
	ST_CfTypeNotContainsText   ST_CfType = 10
	ST_CfTypeBeginsWith        ST_CfType = 11
	ST_CfTypeEndsWith          ST_CfType = 12
	ST_CfTypeContainsBlanks    ST_CfType = 13
	ST_CfTypeNotContainsBlanks ST_CfType = 14
	ST_CfTypeContainsErrors    ST_CfType = 15
	ST_CfTypeNotContainsErrors ST_CfType = 16
	ST_CfTypeTimePeriod        ST_CfType = 17
	ST_CfTypeAboveAverage      ST_CfType = 18
)

func (e ST_CfType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CfTypeUnset:
		attr.Value = ""
	case ST_CfTypeExpression:
		attr.Value = "expression"
	case ST_CfTypeCellIs:
		attr.Value = "cellIs"
	case ST_CfTypeColorScale:
		attr.Value = "colorScale"
	case ST_CfTypeDataBar:
		attr.Value = "dataBar"
	case ST_CfTypeIconSet:
		attr.Value = "iconSet"
	case ST_CfTypeTop10:
		attr.Value = "top10"
	case ST_CfTypeUniqueValues:
		attr.Value = "uniqueValues"
	case ST_CfTypeDuplicateValues:
		attr.Value = "duplicateValues"
	case ST_CfTypeContainsText:
		attr.Value = "containsText"
	case ST_CfTypeNotContainsText:
		attr.Value = "notContainsText"
	case ST_CfTypeBeginsWith:
		attr.Value = "beginsWith"
	case ST_CfTypeEndsWith:
		attr.Value = "endsWith"
	case ST_CfTypeContainsBlanks:
		attr.Value = "containsBlanks"
	case ST_CfTypeNotContainsBlanks:
		attr.Value = "notContainsBlanks"
	case ST_CfTypeContainsErrors:
		attr.Value = "containsErrors"
	case ST_CfTypeNotContainsErrors:
		attr.Value = "notContainsErrors"
	case ST_CfTypeTimePeriod:
		attr.Value = "timePeriod"
	case ST_CfTypeAboveAverage:
		attr.Value = "aboveAverage"
	}
	return attr, nil
}

func (e *ST_CfType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "expression":
		*e = 1
	case "cellIs":
		*e = 2
	case "colorScale":
		*e = 3
	case "dataBar":
		*e = 4
	case "iconSet":
		*e = 5
	case "top10":
		*e = 6
	case "uniqueValues":
		*e = 7
	case "duplicateValues":
		*e = 8
	case "containsText":
		*e = 9
	case "notContainsText":
		*e = 10
	case "beginsWith":
		*e = 11
	case "endsWith":
		*e = 12
	case "containsBlanks":
		*e = 13
	case "notContainsBlanks":
		*e = 14
	case "containsErrors":
		*e = 15
	case "notContainsErrors":
		*e = 16
	case "timePeriod":
		*e = 17
	case "aboveAverage":
		*e = 18
	}
	return nil
}

func (m ST_CfType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CfType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "expression":
			*m = 1
		case "cellIs":
			*m = 2
		case "colorScale":
			*m = 3
		case "dataBar":
			*m = 4
		case "iconSet":
			*m = 5
		case "top10":
			*m = 6
		case "uniqueValues":
			*m = 7
		case "duplicateValues":
			*m = 8
		case "containsText":
			*m = 9
		case "notContainsText":
			*m = 10
		case "beginsWith":
			*m = 11
		case "endsWith":
			*m = 12
		case "containsBlanks":
			*m = 13
		case "notContainsBlanks":
			*m = 14
		case "containsErrors":
			*m = 15
		case "notContainsErrors":
			*m = 16
		case "timePeriod":
			*m = 17
		case "aboveAverage":
			*m = 18
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

func (m ST_CfType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "expression"
	case 2:
		return "cellIs"
	case 3:
		return "colorScale"
	case 4:
		return "dataBar"
	case 5:
		return "iconSet"
	case 6:
		return "top10"
	case 7:
		return "uniqueValues"
	case 8:
		return "duplicateValues"
	case 9:
		return "containsText"
	case 10:
		return "notContainsText"
	case 11:
		return "beginsWith"
	case 12:
		return "endsWith"
	case 13:
		return "containsBlanks"
	case 14:
		return "notContainsBlanks"
	case 15:
		return "containsErrors"
	case 16:
		return "notContainsErrors"
	case 17:
		return "timePeriod"
	case 18:
		return "aboveAverage"
	}
	return ""
}

func (m ST_CfType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CfType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TimePeriod byte

const (
	ST_TimePeriodUnset     ST_TimePeriod = 0
	ST_TimePeriodToday     ST_TimePeriod = 1
	ST_TimePeriodYesterday ST_TimePeriod = 2
	ST_TimePeriodTomorrow  ST_TimePeriod = 3
	ST_TimePeriodLast7Days ST_TimePeriod = 4
	ST_TimePeriodThisMonth ST_TimePeriod = 5
	ST_TimePeriodLastMonth ST_TimePeriod = 6
	ST_TimePeriodNextMonth ST_TimePeriod = 7
	ST_TimePeriodThisWeek  ST_TimePeriod = 8
	ST_TimePeriodLastWeek  ST_TimePeriod = 9
	ST_TimePeriodNextWeek  ST_TimePeriod = 10
)

func (e ST_TimePeriod) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TimePeriodUnset:
		attr.Value = ""
	case ST_TimePeriodToday:
		attr.Value = "today"
	case ST_TimePeriodYesterday:
		attr.Value = "yesterday"
	case ST_TimePeriodTomorrow:
		attr.Value = "tomorrow"
	case ST_TimePeriodLast7Days:
		attr.Value = "last7Days"
	case ST_TimePeriodThisMonth:
		attr.Value = "thisMonth"
	case ST_TimePeriodLastMonth:
		attr.Value = "lastMonth"
	case ST_TimePeriodNextMonth:
		attr.Value = "nextMonth"
	case ST_TimePeriodThisWeek:
		attr.Value = "thisWeek"
	case ST_TimePeriodLastWeek:
		attr.Value = "lastWeek"
	case ST_TimePeriodNextWeek:
		attr.Value = "nextWeek"
	}
	return attr, nil
}

func (e *ST_TimePeriod) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "today":
		*e = 1
	case "yesterday":
		*e = 2
	case "tomorrow":
		*e = 3
	case "last7Days":
		*e = 4
	case "thisMonth":
		*e = 5
	case "lastMonth":
		*e = 6
	case "nextMonth":
		*e = 7
	case "thisWeek":
		*e = 8
	case "lastWeek":
		*e = 9
	case "nextWeek":
		*e = 10
	}
	return nil
}

func (m ST_TimePeriod) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TimePeriod) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "today":
			*m = 1
		case "yesterday":
			*m = 2
		case "tomorrow":
			*m = 3
		case "last7Days":
			*m = 4
		case "thisMonth":
			*m = 5
		case "lastMonth":
			*m = 6
		case "nextMonth":
			*m = 7
		case "thisWeek":
			*m = 8
		case "lastWeek":
			*m = 9
		case "nextWeek":
			*m = 10
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

func (m ST_TimePeriod) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "today"
	case 2:
		return "yesterday"
	case 3:
		return "tomorrow"
	case 4:
		return "last7Days"
	case 5:
		return "thisMonth"
	case 6:
		return "lastMonth"
	case 7:
		return "nextMonth"
	case 8:
		return "thisWeek"
	case 9:
		return "lastWeek"
	case 10:
		return "nextWeek"
	}
	return ""
}

func (m ST_TimePeriod) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TimePeriod) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_ConditionalFormattingOperator byte

const (
	ST_ConditionalFormattingOperatorUnset              ST_ConditionalFormattingOperator = 0
	ST_ConditionalFormattingOperatorLessThan           ST_ConditionalFormattingOperator = 1
	ST_ConditionalFormattingOperatorLessThanOrEqual    ST_ConditionalFormattingOperator = 2
	ST_ConditionalFormattingOperatorEqual              ST_ConditionalFormattingOperator = 3
	ST_ConditionalFormattingOperatorNotEqual           ST_ConditionalFormattingOperator = 4
	ST_ConditionalFormattingOperatorGreaterThanOrEqual ST_ConditionalFormattingOperator = 5
	ST_ConditionalFormattingOperatorGreaterThan        ST_ConditionalFormattingOperator = 6
	ST_ConditionalFormattingOperatorBetween            ST_ConditionalFormattingOperator = 7
	ST_ConditionalFormattingOperatorNotBetween         ST_ConditionalFormattingOperator = 8
	ST_ConditionalFormattingOperatorContainsText       ST_ConditionalFormattingOperator = 9
	ST_ConditionalFormattingOperatorNotContains        ST_ConditionalFormattingOperator = 10
	ST_ConditionalFormattingOperatorBeginsWith         ST_ConditionalFormattingOperator = 11
	ST_ConditionalFormattingOperatorEndsWith           ST_ConditionalFormattingOperator = 12
)

func (e ST_ConditionalFormattingOperator) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ConditionalFormattingOperatorUnset:
		attr.Value = ""
	case ST_ConditionalFormattingOperatorLessThan:
		attr.Value = "lessThan"
	case ST_ConditionalFormattingOperatorLessThanOrEqual:
		attr.Value = "lessThanOrEqual"
	case ST_ConditionalFormattingOperatorEqual:
		attr.Value = "equal"
	case ST_ConditionalFormattingOperatorNotEqual:
		attr.Value = "notEqual"
	case ST_ConditionalFormattingOperatorGreaterThanOrEqual:
		attr.Value = "greaterThanOrEqual"
	case ST_ConditionalFormattingOperatorGreaterThan:
		attr.Value = "greaterThan"
	case ST_ConditionalFormattingOperatorBetween:
		attr.Value = "between"
	case ST_ConditionalFormattingOperatorNotBetween:
		attr.Value = "notBetween"
	case ST_ConditionalFormattingOperatorContainsText:
		attr.Value = "containsText"
	case ST_ConditionalFormattingOperatorNotContains:
		attr.Value = "notContains"
	case ST_ConditionalFormattingOperatorBeginsWith:
		attr.Value = "beginsWith"
	case ST_ConditionalFormattingOperatorEndsWith:
		attr.Value = "endsWith"
	}
	return attr, nil
}

func (e *ST_ConditionalFormattingOperator) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "lessThan":
		*e = 1
	case "lessThanOrEqual":
		*e = 2
	case "equal":
		*e = 3
	case "notEqual":
		*e = 4
	case "greaterThanOrEqual":
		*e = 5
	case "greaterThan":
		*e = 6
	case "between":
		*e = 7
	case "notBetween":
		*e = 8
	case "containsText":
		*e = 9
	case "notContains":
		*e = 10
	case "beginsWith":
		*e = 11
	case "endsWith":
		*e = 12
	}
	return nil
}

func (m ST_ConditionalFormattingOperator) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_ConditionalFormattingOperator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "lessThan":
			*m = 1
		case "lessThanOrEqual":
			*m = 2
		case "equal":
			*m = 3
		case "notEqual":
			*m = 4
		case "greaterThanOrEqual":
			*m = 5
		case "greaterThan":
			*m = 6
		case "between":
			*m = 7
		case "notBetween":
			*m = 8
		case "containsText":
			*m = 9
		case "notContains":
			*m = 10
		case "beginsWith":
			*m = 11
		case "endsWith":
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

func (m ST_ConditionalFormattingOperator) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "lessThan"
	case 2:
		return "lessThanOrEqual"
	case 3:
		return "equal"
	case 4:
		return "notEqual"
	case 5:
		return "greaterThanOrEqual"
	case 6:
		return "greaterThan"
	case 7:
		return "between"
	case 8:
		return "notBetween"
	case 9:
		return "containsText"
	case 10:
		return "notContains"
	case 11:
		return "beginsWith"
	case 12:
		return "endsWith"
	}
	return ""
}

func (m ST_ConditionalFormattingOperator) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_ConditionalFormattingOperator) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CfvoType byte

const (
	ST_CfvoTypeUnset      ST_CfvoType = 0
	ST_CfvoTypeNum        ST_CfvoType = 1
	ST_CfvoTypePercent    ST_CfvoType = 2
	ST_CfvoTypeMax        ST_CfvoType = 3
	ST_CfvoTypeMin        ST_CfvoType = 4
	ST_CfvoTypeFormula    ST_CfvoType = 5
	ST_CfvoTypePercentile ST_CfvoType = 6
)

func (e ST_CfvoType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CfvoTypeUnset:
		attr.Value = ""
	case ST_CfvoTypeNum:
		attr.Value = "num"
	case ST_CfvoTypePercent:
		attr.Value = "percent"
	case ST_CfvoTypeMax:
		attr.Value = "max"
	case ST_CfvoTypeMin:
		attr.Value = "min"
	case ST_CfvoTypeFormula:
		attr.Value = "formula"
	case ST_CfvoTypePercentile:
		attr.Value = "percentile"
	}
	return attr, nil
}

func (e *ST_CfvoType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "num":
		*e = 1
	case "percent":
		*e = 2
	case "max":
		*e = 3
	case "min":
		*e = 4
	case "formula":
		*e = 5
	case "percentile":
		*e = 6
	}
	return nil
}

func (m ST_CfvoType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CfvoType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "num":
			*m = 1
		case "percent":
			*m = 2
		case "max":
			*m = 3
		case "min":
			*m = 4
		case "formula":
			*m = 5
		case "percentile":
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

func (m ST_CfvoType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "num"
	case 2:
		return "percent"
	case 3:
		return "max"
	case 4:
		return "min"
	case 5:
		return "formula"
	case 6:
		return "percentile"
	}
	return ""
}

func (m ST_CfvoType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CfvoType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PageOrder byte

const (
	ST_PageOrderUnset        ST_PageOrder = 0
	ST_PageOrderDownThenOver ST_PageOrder = 1
	ST_PageOrderOverThenDown ST_PageOrder = 2
)

func (e ST_PageOrder) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PageOrderUnset:
		attr.Value = ""
	case ST_PageOrderDownThenOver:
		attr.Value = "downThenOver"
	case ST_PageOrderOverThenDown:
		attr.Value = "overThenDown"
	}
	return attr, nil
}

func (e *ST_PageOrder) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "downThenOver":
		*e = 1
	case "overThenDown":
		*e = 2
	}
	return nil
}

func (m ST_PageOrder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PageOrder) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "downThenOver":
			*m = 1
		case "overThenDown":
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

func (m ST_PageOrder) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "downThenOver"
	case 2:
		return "overThenDown"
	}
	return ""
}

func (m ST_PageOrder) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PageOrder) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Orientation byte

const (
	ST_OrientationUnset     ST_Orientation = 0
	ST_OrientationDefault   ST_Orientation = 1
	ST_OrientationPortrait  ST_Orientation = 2
	ST_OrientationLandscape ST_Orientation = 3
)

func (e ST_Orientation) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_OrientationUnset:
		attr.Value = ""
	case ST_OrientationDefault:
		attr.Value = "default"
	case ST_OrientationPortrait:
		attr.Value = "portrait"
	case ST_OrientationLandscape:
		attr.Value = "landscape"
	}
	return attr, nil
}

func (e *ST_Orientation) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "default":
		*e = 1
	case "portrait":
		*e = 2
	case "landscape":
		*e = 3
	}
	return nil
}

func (m ST_Orientation) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Orientation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "default":
			*m = 1
		case "portrait":
			*m = 2
		case "landscape":
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

func (m ST_Orientation) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "default"
	case 2:
		return "portrait"
	case 3:
		return "landscape"
	}
	return ""
}

func (m ST_Orientation) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Orientation) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CellComments byte

const (
	ST_CellCommentsUnset       ST_CellComments = 0
	ST_CellCommentsNone        ST_CellComments = 1
	ST_CellCommentsAsDisplayed ST_CellComments = 2
	ST_CellCommentsAtEnd       ST_CellComments = 3
)

func (e ST_CellComments) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CellCommentsUnset:
		attr.Value = ""
	case ST_CellCommentsNone:
		attr.Value = "none"
	case ST_CellCommentsAsDisplayed:
		attr.Value = "asDisplayed"
	case ST_CellCommentsAtEnd:
		attr.Value = "atEnd"
	}
	return attr, nil
}

func (e *ST_CellComments) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "asDisplayed":
		*e = 2
	case "atEnd":
		*e = 3
	}
	return nil
}

func (m ST_CellComments) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CellComments) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "asDisplayed":
			*m = 2
		case "atEnd":
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

func (m ST_CellComments) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "asDisplayed"
	case 3:
		return "atEnd"
	}
	return ""
}

func (m ST_CellComments) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CellComments) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PrintError byte

const (
	ST_PrintErrorUnset     ST_PrintError = 0
	ST_PrintErrorDisplayed ST_PrintError = 1
	ST_PrintErrorBlank     ST_PrintError = 2
	ST_PrintErrorDash      ST_PrintError = 3
	ST_PrintErrorNA        ST_PrintError = 4
)

func (e ST_PrintError) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PrintErrorUnset:
		attr.Value = ""
	case ST_PrintErrorDisplayed:
		attr.Value = "displayed"
	case ST_PrintErrorBlank:
		attr.Value = "blank"
	case ST_PrintErrorDash:
		attr.Value = "dash"
	case ST_PrintErrorNA:
		attr.Value = "NA"
	}
	return attr, nil
}

func (e *ST_PrintError) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "displayed":
		*e = 1
	case "blank":
		*e = 2
	case "dash":
		*e = 3
	case "NA":
		*e = 4
	}
	return nil
}

func (m ST_PrintError) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PrintError) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "displayed":
			*m = 1
		case "blank":
			*m = 2
		case "dash":
			*m = 3
		case "NA":
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

func (m ST_PrintError) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "displayed"
	case 2:
		return "blank"
	case 3:
		return "dash"
	case 4:
		return "NA"
	}
	return ""
}

func (m ST_PrintError) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PrintError) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DvAspect byte

const (
	ST_DvAspectUnset            ST_DvAspect = 0
	ST_DvAspectDVASPECT_CONTENT ST_DvAspect = 1
	ST_DvAspectDVASPECT_ICON    ST_DvAspect = 2
)

func (e ST_DvAspect) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DvAspectUnset:
		attr.Value = ""
	case ST_DvAspectDVASPECT_CONTENT:
		attr.Value = "DVASPECT_CONTENT"
	case ST_DvAspectDVASPECT_ICON:
		attr.Value = "DVASPECT_ICON"
	}
	return attr, nil
}

func (e *ST_DvAspect) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "DVASPECT_CONTENT":
		*e = 1
	case "DVASPECT_ICON":
		*e = 2
	}
	return nil
}

func (m ST_DvAspect) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DvAspect) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "DVASPECT_CONTENT":
			*m = 1
		case "DVASPECT_ICON":
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

func (m ST_DvAspect) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "DVASPECT_CONTENT"
	case 2:
		return "DVASPECT_ICON"
	}
	return ""
}

func (m ST_DvAspect) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DvAspect) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_OleUpdate byte

const (
	ST_OleUpdateUnset            ST_OleUpdate = 0
	ST_OleUpdateOLEUPDATE_ALWAYS ST_OleUpdate = 1
	ST_OleUpdateOLEUPDATE_ONCALL ST_OleUpdate = 2
)

func (e ST_OleUpdate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_OleUpdateUnset:
		attr.Value = ""
	case ST_OleUpdateOLEUPDATE_ALWAYS:
		attr.Value = "OLEUPDATE_ALWAYS"
	case ST_OleUpdateOLEUPDATE_ONCALL:
		attr.Value = "OLEUPDATE_ONCALL"
	}
	return attr, nil
}

func (e *ST_OleUpdate) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "OLEUPDATE_ALWAYS":
		*e = 1
	case "OLEUPDATE_ONCALL":
		*e = 2
	}
	return nil
}

func (m ST_OleUpdate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_OleUpdate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "OLEUPDATE_ALWAYS":
			*m = 1
		case "OLEUPDATE_ONCALL":
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

func (m ST_OleUpdate) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "OLEUPDATE_ALWAYS"
	case 2:
		return "OLEUPDATE_ONCALL"
	}
	return ""
}

func (m ST_OleUpdate) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_OleUpdate) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_WebSourceType byte

const (
	ST_WebSourceTypeUnset      ST_WebSourceType = 0
	ST_WebSourceTypeSheet      ST_WebSourceType = 1
	ST_WebSourceTypePrintArea  ST_WebSourceType = 2
	ST_WebSourceTypeAutoFilter ST_WebSourceType = 3
	ST_WebSourceTypeRange      ST_WebSourceType = 4
	ST_WebSourceTypeChart      ST_WebSourceType = 5
	ST_WebSourceTypePivotTable ST_WebSourceType = 6
	ST_WebSourceTypeQuery      ST_WebSourceType = 7
	ST_WebSourceTypeLabel      ST_WebSourceType = 8
)

func (e ST_WebSourceType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_WebSourceTypeUnset:
		attr.Value = ""
	case ST_WebSourceTypeSheet:
		attr.Value = "sheet"
	case ST_WebSourceTypePrintArea:
		attr.Value = "printArea"
	case ST_WebSourceTypeAutoFilter:
		attr.Value = "autoFilter"
	case ST_WebSourceTypeRange:
		attr.Value = "range"
	case ST_WebSourceTypeChart:
		attr.Value = "chart"
	case ST_WebSourceTypePivotTable:
		attr.Value = "pivotTable"
	case ST_WebSourceTypeQuery:
		attr.Value = "query"
	case ST_WebSourceTypeLabel:
		attr.Value = "label"
	}
	return attr, nil
}

func (e *ST_WebSourceType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "sheet":
		*e = 1
	case "printArea":
		*e = 2
	case "autoFilter":
		*e = 3
	case "range":
		*e = 4
	case "chart":
		*e = 5
	case "pivotTable":
		*e = 6
	case "query":
		*e = 7
	case "label":
		*e = 8
	}
	return nil
}

func (m ST_WebSourceType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_WebSourceType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "sheet":
			*m = 1
		case "printArea":
			*m = 2
		case "autoFilter":
			*m = 3
		case "range":
			*m = 4
		case "chart":
			*m = 5
		case "pivotTable":
			*m = 6
		case "query":
			*m = 7
		case "label":
			*m = 8
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

func (m ST_WebSourceType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "sheet"
	case 2:
		return "printArea"
	case 3:
		return "autoFilter"
	case 4:
		return "range"
	case 5:
		return "chart"
	case 6:
		return "pivotTable"
	case 7:
		return "query"
	case 8:
		return "label"
	}
	return ""
}

func (m ST_WebSourceType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_WebSourceType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PaneState byte

const (
	ST_PaneStateUnset       ST_PaneState = 0
	ST_PaneStateSplit       ST_PaneState = 1
	ST_PaneStateFrozen      ST_PaneState = 2
	ST_PaneStateFrozenSplit ST_PaneState = 3
)

func (e ST_PaneState) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PaneStateUnset:
		attr.Value = ""
	case ST_PaneStateSplit:
		attr.Value = "split"
	case ST_PaneStateFrozen:
		attr.Value = "frozen"
	case ST_PaneStateFrozenSplit:
		attr.Value = "frozenSplit"
	}
	return attr, nil
}

func (e *ST_PaneState) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "split":
		*e = 1
	case "frozen":
		*e = 2
	case "frozenSplit":
		*e = 3
	}
	return nil
}

func (m ST_PaneState) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PaneState) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "split":
			*m = 1
		case "frozen":
			*m = 2
		case "frozenSplit":
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

func (m ST_PaneState) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "split"
	case 2:
		return "frozen"
	case 3:
		return "frozenSplit"
	}
	return ""
}

func (m ST_PaneState) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PaneState) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_MdxFunctionType byte

const (
	ST_MdxFunctionTypeUnset ST_MdxFunctionType = 0
	ST_MdxFunctionTypeM     ST_MdxFunctionType = 1
	ST_MdxFunctionTypeV     ST_MdxFunctionType = 2
	ST_MdxFunctionTypeS     ST_MdxFunctionType = 3
	ST_MdxFunctionTypeC     ST_MdxFunctionType = 4
	ST_MdxFunctionTypeR     ST_MdxFunctionType = 5
	ST_MdxFunctionTypeP     ST_MdxFunctionType = 6
	ST_MdxFunctionTypeK     ST_MdxFunctionType = 7
)

func (e ST_MdxFunctionType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MdxFunctionTypeUnset:
		attr.Value = ""
	case ST_MdxFunctionTypeM:
		attr.Value = "m"
	case ST_MdxFunctionTypeV:
		attr.Value = "v"
	case ST_MdxFunctionTypeS:
		attr.Value = "s"
	case ST_MdxFunctionTypeC:
		attr.Value = "c"
	case ST_MdxFunctionTypeR:
		attr.Value = "r"
	case ST_MdxFunctionTypeP:
		attr.Value = "p"
	case ST_MdxFunctionTypeK:
		attr.Value = "k"
	}
	return attr, nil
}

func (e *ST_MdxFunctionType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "m":
		*e = 1
	case "v":
		*e = 2
	case "s":
		*e = 3
	case "c":
		*e = 4
	case "r":
		*e = 5
	case "p":
		*e = 6
	case "k":
		*e = 7
	}
	return nil
}

func (m ST_MdxFunctionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_MdxFunctionType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "m":
			*m = 1
		case "v":
			*m = 2
		case "s":
			*m = 3
		case "c":
			*m = 4
		case "r":
			*m = 5
		case "p":
			*m = 6
		case "k":
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

func (m ST_MdxFunctionType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "m"
	case 2:
		return "v"
	case 3:
		return "s"
	case 4:
		return "c"
	case 5:
		return "r"
	case 6:
		return "p"
	case 7:
		return "k"
	}
	return ""
}

func (m ST_MdxFunctionType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_MdxFunctionType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_MdxSetOrder byte

const (
	ST_MdxSetOrderUnset ST_MdxSetOrder = 0
	ST_MdxSetOrderU     ST_MdxSetOrder = 1
	ST_MdxSetOrderA     ST_MdxSetOrder = 2
	ST_MdxSetOrderD     ST_MdxSetOrder = 3
	ST_MdxSetOrderAa    ST_MdxSetOrder = 4
	ST_MdxSetOrderAd    ST_MdxSetOrder = 5
	ST_MdxSetOrderNa    ST_MdxSetOrder = 6
	ST_MdxSetOrderNd    ST_MdxSetOrder = 7
)

func (e ST_MdxSetOrder) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MdxSetOrderUnset:
		attr.Value = ""
	case ST_MdxSetOrderU:
		attr.Value = "u"
	case ST_MdxSetOrderA:
		attr.Value = "a"
	case ST_MdxSetOrderD:
		attr.Value = "d"
	case ST_MdxSetOrderAa:
		attr.Value = "aa"
	case ST_MdxSetOrderAd:
		attr.Value = "ad"
	case ST_MdxSetOrderNa:
		attr.Value = "na"
	case ST_MdxSetOrderNd:
		attr.Value = "nd"
	}
	return attr, nil
}

func (e *ST_MdxSetOrder) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "u":
		*e = 1
	case "a":
		*e = 2
	case "d":
		*e = 3
	case "aa":
		*e = 4
	case "ad":
		*e = 5
	case "na":
		*e = 6
	case "nd":
		*e = 7
	}
	return nil
}

func (m ST_MdxSetOrder) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_MdxSetOrder) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "u":
			*m = 1
		case "a":
			*m = 2
		case "d":
			*m = 3
		case "aa":
			*m = 4
		case "ad":
			*m = 5
		case "na":
			*m = 6
		case "nd":
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

func (m ST_MdxSetOrder) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "u"
	case 2:
		return "a"
	case 3:
		return "d"
	case 4:
		return "aa"
	case 5:
		return "ad"
	case 6:
		return "na"
	case 7:
		return "nd"
	}
	return ""
}

func (m ST_MdxSetOrder) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_MdxSetOrder) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_MdxKPIProperty byte

const (
	ST_MdxKPIPropertyUnset ST_MdxKPIProperty = 0
	ST_MdxKPIPropertyV     ST_MdxKPIProperty = 1
	ST_MdxKPIPropertyG     ST_MdxKPIProperty = 2
	ST_MdxKPIPropertyS     ST_MdxKPIProperty = 3
	ST_MdxKPIPropertyT     ST_MdxKPIProperty = 4
	ST_MdxKPIPropertyW     ST_MdxKPIProperty = 5
	ST_MdxKPIPropertyM     ST_MdxKPIProperty = 6
)

func (e ST_MdxKPIProperty) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_MdxKPIPropertyUnset:
		attr.Value = ""
	case ST_MdxKPIPropertyV:
		attr.Value = "v"
	case ST_MdxKPIPropertyG:
		attr.Value = "g"
	case ST_MdxKPIPropertyS:
		attr.Value = "s"
	case ST_MdxKPIPropertyT:
		attr.Value = "t"
	case ST_MdxKPIPropertyW:
		attr.Value = "w"
	case ST_MdxKPIPropertyM:
		attr.Value = "m"
	}
	return attr, nil
}

func (e *ST_MdxKPIProperty) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "v":
		*e = 1
	case "g":
		*e = 2
	case "s":
		*e = 3
	case "t":
		*e = 4
	case "w":
		*e = 5
	case "m":
		*e = 6
	}
	return nil
}

func (m ST_MdxKPIProperty) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_MdxKPIProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "v":
			*m = 1
		case "g":
			*m = 2
		case "s":
			*m = 3
		case "t":
			*m = 4
		case "w":
			*m = 5
		case "m":
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

func (m ST_MdxKPIProperty) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "v"
	case 2:
		return "g"
	case 3:
		return "s"
	case 4:
		return "t"
	case 5:
		return "w"
	case 6:
		return "m"
	}
	return ""
}

func (m ST_MdxKPIProperty) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_MdxKPIProperty) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_BorderStyle byte

const (
	ST_BorderStyleUnset            ST_BorderStyle = 0
	ST_BorderStyleNone             ST_BorderStyle = 1
	ST_BorderStyleThin             ST_BorderStyle = 2
	ST_BorderStyleMedium           ST_BorderStyle = 3
	ST_BorderStyleDashed           ST_BorderStyle = 4
	ST_BorderStyleDotted           ST_BorderStyle = 5
	ST_BorderStyleThick            ST_BorderStyle = 6
	ST_BorderStyleDouble           ST_BorderStyle = 7
	ST_BorderStyleHair             ST_BorderStyle = 8
	ST_BorderStyleMediumDashed     ST_BorderStyle = 9
	ST_BorderStyleDashDot          ST_BorderStyle = 10
	ST_BorderStyleMediumDashDot    ST_BorderStyle = 11
	ST_BorderStyleDashDotDot       ST_BorderStyle = 12
	ST_BorderStyleMediumDashDotDot ST_BorderStyle = 13
	ST_BorderStyleSlantDashDot     ST_BorderStyle = 14
)

func (e ST_BorderStyle) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_BorderStyleUnset:
		attr.Value = ""
	case ST_BorderStyleNone:
		attr.Value = "none"
	case ST_BorderStyleThin:
		attr.Value = "thin"
	case ST_BorderStyleMedium:
		attr.Value = "medium"
	case ST_BorderStyleDashed:
		attr.Value = "dashed"
	case ST_BorderStyleDotted:
		attr.Value = "dotted"
	case ST_BorderStyleThick:
		attr.Value = "thick"
	case ST_BorderStyleDouble:
		attr.Value = "double"
	case ST_BorderStyleHair:
		attr.Value = "hair"
	case ST_BorderStyleMediumDashed:
		attr.Value = "mediumDashed"
	case ST_BorderStyleDashDot:
		attr.Value = "dashDot"
	case ST_BorderStyleMediumDashDot:
		attr.Value = "mediumDashDot"
	case ST_BorderStyleDashDotDot:
		attr.Value = "dashDotDot"
	case ST_BorderStyleMediumDashDotDot:
		attr.Value = "mediumDashDotDot"
	case ST_BorderStyleSlantDashDot:
		attr.Value = "slantDashDot"
	}
	return attr, nil
}

func (e *ST_BorderStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "thin":
		*e = 2
	case "medium":
		*e = 3
	case "dashed":
		*e = 4
	case "dotted":
		*e = 5
	case "thick":
		*e = 6
	case "double":
		*e = 7
	case "hair":
		*e = 8
	case "mediumDashed":
		*e = 9
	case "dashDot":
		*e = 10
	case "mediumDashDot":
		*e = 11
	case "dashDotDot":
		*e = 12
	case "mediumDashDotDot":
		*e = 13
	case "slantDashDot":
		*e = 14
	}
	return nil
}

func (m ST_BorderStyle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_BorderStyle) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "thin":
			*m = 2
		case "medium":
			*m = 3
		case "dashed":
			*m = 4
		case "dotted":
			*m = 5
		case "thick":
			*m = 6
		case "double":
			*m = 7
		case "hair":
			*m = 8
		case "mediumDashed":
			*m = 9
		case "dashDot":
			*m = 10
		case "mediumDashDot":
			*m = 11
		case "dashDotDot":
			*m = 12
		case "mediumDashDotDot":
			*m = 13
		case "slantDashDot":
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

func (m ST_BorderStyle) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "thin"
	case 3:
		return "medium"
	case 4:
		return "dashed"
	case 5:
		return "dotted"
	case 6:
		return "thick"
	case 7:
		return "double"
	case 8:
		return "hair"
	case 9:
		return "mediumDashed"
	case 10:
		return "dashDot"
	case 11:
		return "mediumDashDot"
	case 12:
		return "dashDotDot"
	case 13:
		return "mediumDashDotDot"
	case 14:
		return "slantDashDot"
	}
	return ""
}

func (m ST_BorderStyle) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_BorderStyle) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_PatternType byte

const (
	ST_PatternTypeUnset           ST_PatternType = 0
	ST_PatternTypeNone            ST_PatternType = 1
	ST_PatternTypeSolid           ST_PatternType = 2
	ST_PatternTypeMediumGray      ST_PatternType = 3
	ST_PatternTypeDarkGray        ST_PatternType = 4
	ST_PatternTypeLightGray       ST_PatternType = 5
	ST_PatternTypeDarkHorizontal  ST_PatternType = 6
	ST_PatternTypeDarkVertical    ST_PatternType = 7
	ST_PatternTypeDarkDown        ST_PatternType = 8
	ST_PatternTypeDarkUp          ST_PatternType = 9
	ST_PatternTypeDarkGrid        ST_PatternType = 10
	ST_PatternTypeDarkTrellis     ST_PatternType = 11
	ST_PatternTypeLightHorizontal ST_PatternType = 12
	ST_PatternTypeLightVertical   ST_PatternType = 13
	ST_PatternTypeLightDown       ST_PatternType = 14
	ST_PatternTypeLightUp         ST_PatternType = 15
	ST_PatternTypeLightGrid       ST_PatternType = 16
	ST_PatternTypeLightTrellis    ST_PatternType = 17
	ST_PatternTypeGray125         ST_PatternType = 18
	ST_PatternTypeGray0625        ST_PatternType = 19
)

func (e ST_PatternType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_PatternTypeUnset:
		attr.Value = ""
	case ST_PatternTypeNone:
		attr.Value = "none"
	case ST_PatternTypeSolid:
		attr.Value = "solid"
	case ST_PatternTypeMediumGray:
		attr.Value = "mediumGray"
	case ST_PatternTypeDarkGray:
		attr.Value = "darkGray"
	case ST_PatternTypeLightGray:
		attr.Value = "lightGray"
	case ST_PatternTypeDarkHorizontal:
		attr.Value = "darkHorizontal"
	case ST_PatternTypeDarkVertical:
		attr.Value = "darkVertical"
	case ST_PatternTypeDarkDown:
		attr.Value = "darkDown"
	case ST_PatternTypeDarkUp:
		attr.Value = "darkUp"
	case ST_PatternTypeDarkGrid:
		attr.Value = "darkGrid"
	case ST_PatternTypeDarkTrellis:
		attr.Value = "darkTrellis"
	case ST_PatternTypeLightHorizontal:
		attr.Value = "lightHorizontal"
	case ST_PatternTypeLightVertical:
		attr.Value = "lightVertical"
	case ST_PatternTypeLightDown:
		attr.Value = "lightDown"
	case ST_PatternTypeLightUp:
		attr.Value = "lightUp"
	case ST_PatternTypeLightGrid:
		attr.Value = "lightGrid"
	case ST_PatternTypeLightTrellis:
		attr.Value = "lightTrellis"
	case ST_PatternTypeGray125:
		attr.Value = "gray125"
	case ST_PatternTypeGray0625:
		attr.Value = "gray0625"
	}
	return attr, nil
}

func (e *ST_PatternType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "solid":
		*e = 2
	case "mediumGray":
		*e = 3
	case "darkGray":
		*e = 4
	case "lightGray":
		*e = 5
	case "darkHorizontal":
		*e = 6
	case "darkVertical":
		*e = 7
	case "darkDown":
		*e = 8
	case "darkUp":
		*e = 9
	case "darkGrid":
		*e = 10
	case "darkTrellis":
		*e = 11
	case "lightHorizontal":
		*e = 12
	case "lightVertical":
		*e = 13
	case "lightDown":
		*e = 14
	case "lightUp":
		*e = 15
	case "lightGrid":
		*e = 16
	case "lightTrellis":
		*e = 17
	case "gray125":
		*e = 18
	case "gray0625":
		*e = 19
	}
	return nil
}

func (m ST_PatternType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_PatternType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "solid":
			*m = 2
		case "mediumGray":
			*m = 3
		case "darkGray":
			*m = 4
		case "lightGray":
			*m = 5
		case "darkHorizontal":
			*m = 6
		case "darkVertical":
			*m = 7
		case "darkDown":
			*m = 8
		case "darkUp":
			*m = 9
		case "darkGrid":
			*m = 10
		case "darkTrellis":
			*m = 11
		case "lightHorizontal":
			*m = 12
		case "lightVertical":
			*m = 13
		case "lightDown":
			*m = 14
		case "lightUp":
			*m = 15
		case "lightGrid":
			*m = 16
		case "lightTrellis":
			*m = 17
		case "gray125":
			*m = 18
		case "gray0625":
			*m = 19
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

func (m ST_PatternType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "solid"
	case 3:
		return "mediumGray"
	case 4:
		return "darkGray"
	case 5:
		return "lightGray"
	case 6:
		return "darkHorizontal"
	case 7:
		return "darkVertical"
	case 8:
		return "darkDown"
	case 9:
		return "darkUp"
	case 10:
		return "darkGrid"
	case 11:
		return "darkTrellis"
	case 12:
		return "lightHorizontal"
	case 13:
		return "lightVertical"
	case 14:
		return "lightDown"
	case 15:
		return "lightUp"
	case 16:
		return "lightGrid"
	case 17:
		return "lightTrellis"
	case 18:
		return "gray125"
	case 19:
		return "gray0625"
	}
	return ""
}

func (m ST_PatternType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_PatternType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_GradientType byte

const (
	ST_GradientTypeUnset  ST_GradientType = 0
	ST_GradientTypeLinear ST_GradientType = 1
	ST_GradientTypePath   ST_GradientType = 2
)

func (e ST_GradientType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_GradientTypeUnset:
		attr.Value = ""
	case ST_GradientTypeLinear:
		attr.Value = "linear"
	case ST_GradientTypePath:
		attr.Value = "path"
	}
	return attr, nil
}

func (e *ST_GradientType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "linear":
		*e = 1
	case "path":
		*e = 2
	}
	return nil
}

func (m ST_GradientType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_GradientType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "linear":
			*m = 1
		case "path":
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

func (m ST_GradientType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "linear"
	case 2:
		return "path"
	}
	return ""
}

func (m ST_GradientType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_GradientType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_HorizontalAlignment byte

const (
	ST_HorizontalAlignmentUnset            ST_HorizontalAlignment = 0
	ST_HorizontalAlignmentGeneral          ST_HorizontalAlignment = 1
	ST_HorizontalAlignmentLeft             ST_HorizontalAlignment = 2
	ST_HorizontalAlignmentCenter           ST_HorizontalAlignment = 3
	ST_HorizontalAlignmentRight            ST_HorizontalAlignment = 4
	ST_HorizontalAlignmentFill             ST_HorizontalAlignment = 5
	ST_HorizontalAlignmentJustify          ST_HorizontalAlignment = 6
	ST_HorizontalAlignmentCenterContinuous ST_HorizontalAlignment = 7
	ST_HorizontalAlignmentDistributed      ST_HorizontalAlignment = 8
)

func (e ST_HorizontalAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_HorizontalAlignmentUnset:
		attr.Value = ""
	case ST_HorizontalAlignmentGeneral:
		attr.Value = "general"
	case ST_HorizontalAlignmentLeft:
		attr.Value = "left"
	case ST_HorizontalAlignmentCenter:
		attr.Value = "center"
	case ST_HorizontalAlignmentRight:
		attr.Value = "right"
	case ST_HorizontalAlignmentFill:
		attr.Value = "fill"
	case ST_HorizontalAlignmentJustify:
		attr.Value = "justify"
	case ST_HorizontalAlignmentCenterContinuous:
		attr.Value = "centerContinuous"
	case ST_HorizontalAlignmentDistributed:
		attr.Value = "distributed"
	}
	return attr, nil
}

func (e *ST_HorizontalAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "general":
		*e = 1
	case "left":
		*e = 2
	case "center":
		*e = 3
	case "right":
		*e = 4
	case "fill":
		*e = 5
	case "justify":
		*e = 6
	case "centerContinuous":
		*e = 7
	case "distributed":
		*e = 8
	}
	return nil
}

func (m ST_HorizontalAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_HorizontalAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "general":
			*m = 1
		case "left":
			*m = 2
		case "center":
			*m = 3
		case "right":
			*m = 4
		case "fill":
			*m = 5
		case "justify":
			*m = 6
		case "centerContinuous":
			*m = 7
		case "distributed":
			*m = 8
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

func (m ST_HorizontalAlignment) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "general"
	case 2:
		return "left"
	case 3:
		return "center"
	case 4:
		return "right"
	case 5:
		return "fill"
	case 6:
		return "justify"
	case 7:
		return "centerContinuous"
	case 8:
		return "distributed"
	}
	return ""
}

func (m ST_HorizontalAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_HorizontalAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_VerticalAlignment byte

const (
	ST_VerticalAlignmentUnset       ST_VerticalAlignment = 0
	ST_VerticalAlignmentTop         ST_VerticalAlignment = 1
	ST_VerticalAlignmentCenter      ST_VerticalAlignment = 2
	ST_VerticalAlignmentBottom      ST_VerticalAlignment = 3
	ST_VerticalAlignmentJustify     ST_VerticalAlignment = 4
	ST_VerticalAlignmentDistributed ST_VerticalAlignment = 5
)

func (e ST_VerticalAlignment) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VerticalAlignmentUnset:
		attr.Value = ""
	case ST_VerticalAlignmentTop:
		attr.Value = "top"
	case ST_VerticalAlignmentCenter:
		attr.Value = "center"
	case ST_VerticalAlignmentBottom:
		attr.Value = "bottom"
	case ST_VerticalAlignmentJustify:
		attr.Value = "justify"
	case ST_VerticalAlignmentDistributed:
		attr.Value = "distributed"
	}
	return attr, nil
}

func (e *ST_VerticalAlignment) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "top":
		*e = 1
	case "center":
		*e = 2
	case "bottom":
		*e = 3
	case "justify":
		*e = 4
	case "distributed":
		*e = 5
	}
	return nil
}

func (m ST_VerticalAlignment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VerticalAlignment) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "justify":
			*m = 4
		case "distributed":
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

func (m ST_VerticalAlignment) String() string {
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
		return "justify"
	case 5:
		return "distributed"
	}
	return ""
}

func (m ST_VerticalAlignment) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VerticalAlignment) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TableStyleType byte

const (
	ST_TableStyleTypeUnset                  ST_TableStyleType = 0
	ST_TableStyleTypeWholeTable             ST_TableStyleType = 1
	ST_TableStyleTypeHeaderRow              ST_TableStyleType = 2
	ST_TableStyleTypeTotalRow               ST_TableStyleType = 3
	ST_TableStyleTypeFirstColumn            ST_TableStyleType = 4
	ST_TableStyleTypeLastColumn             ST_TableStyleType = 5
	ST_TableStyleTypeFirstRowStripe         ST_TableStyleType = 6
	ST_TableStyleTypeSecondRowStripe        ST_TableStyleType = 7
	ST_TableStyleTypeFirstColumnStripe      ST_TableStyleType = 8
	ST_TableStyleTypeSecondColumnStripe     ST_TableStyleType = 9
	ST_TableStyleTypeFirstHeaderCell        ST_TableStyleType = 10
	ST_TableStyleTypeLastHeaderCell         ST_TableStyleType = 11
	ST_TableStyleTypeFirstTotalCell         ST_TableStyleType = 12
	ST_TableStyleTypeLastTotalCell          ST_TableStyleType = 13
	ST_TableStyleTypeFirstSubtotalColumn    ST_TableStyleType = 14
	ST_TableStyleTypeSecondSubtotalColumn   ST_TableStyleType = 15
	ST_TableStyleTypeThirdSubtotalColumn    ST_TableStyleType = 16
	ST_TableStyleTypeFirstSubtotalRow       ST_TableStyleType = 17
	ST_TableStyleTypeSecondSubtotalRow      ST_TableStyleType = 18
	ST_TableStyleTypeThirdSubtotalRow       ST_TableStyleType = 19
	ST_TableStyleTypeBlankRow               ST_TableStyleType = 20
	ST_TableStyleTypeFirstColumnSubheading  ST_TableStyleType = 21
	ST_TableStyleTypeSecondColumnSubheading ST_TableStyleType = 22
	ST_TableStyleTypeThirdColumnSubheading  ST_TableStyleType = 23
	ST_TableStyleTypeFirstRowSubheading     ST_TableStyleType = 24
	ST_TableStyleTypeSecondRowSubheading    ST_TableStyleType = 25
	ST_TableStyleTypeThirdRowSubheading     ST_TableStyleType = 26
	ST_TableStyleTypePageFieldLabels        ST_TableStyleType = 27
	ST_TableStyleTypePageFieldValues        ST_TableStyleType = 28
)

func (e ST_TableStyleType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TableStyleTypeUnset:
		attr.Value = ""
	case ST_TableStyleTypeWholeTable:
		attr.Value = "wholeTable"
	case ST_TableStyleTypeHeaderRow:
		attr.Value = "headerRow"
	case ST_TableStyleTypeTotalRow:
		attr.Value = "totalRow"
	case ST_TableStyleTypeFirstColumn:
		attr.Value = "firstColumn"
	case ST_TableStyleTypeLastColumn:
		attr.Value = "lastColumn"
	case ST_TableStyleTypeFirstRowStripe:
		attr.Value = "firstRowStripe"
	case ST_TableStyleTypeSecondRowStripe:
		attr.Value = "secondRowStripe"
	case ST_TableStyleTypeFirstColumnStripe:
		attr.Value = "firstColumnStripe"
	case ST_TableStyleTypeSecondColumnStripe:
		attr.Value = "secondColumnStripe"
	case ST_TableStyleTypeFirstHeaderCell:
		attr.Value = "firstHeaderCell"
	case ST_TableStyleTypeLastHeaderCell:
		attr.Value = "lastHeaderCell"
	case ST_TableStyleTypeFirstTotalCell:
		attr.Value = "firstTotalCell"
	case ST_TableStyleTypeLastTotalCell:
		attr.Value = "lastTotalCell"
	case ST_TableStyleTypeFirstSubtotalColumn:
		attr.Value = "firstSubtotalColumn"
	case ST_TableStyleTypeSecondSubtotalColumn:
		attr.Value = "secondSubtotalColumn"
	case ST_TableStyleTypeThirdSubtotalColumn:
		attr.Value = "thirdSubtotalColumn"
	case ST_TableStyleTypeFirstSubtotalRow:
		attr.Value = "firstSubtotalRow"
	case ST_TableStyleTypeSecondSubtotalRow:
		attr.Value = "secondSubtotalRow"
	case ST_TableStyleTypeThirdSubtotalRow:
		attr.Value = "thirdSubtotalRow"
	case ST_TableStyleTypeBlankRow:
		attr.Value = "blankRow"
	case ST_TableStyleTypeFirstColumnSubheading:
		attr.Value = "firstColumnSubheading"
	case ST_TableStyleTypeSecondColumnSubheading:
		attr.Value = "secondColumnSubheading"
	case ST_TableStyleTypeThirdColumnSubheading:
		attr.Value = "thirdColumnSubheading"
	case ST_TableStyleTypeFirstRowSubheading:
		attr.Value = "firstRowSubheading"
	case ST_TableStyleTypeSecondRowSubheading:
		attr.Value = "secondRowSubheading"
	case ST_TableStyleTypeThirdRowSubheading:
		attr.Value = "thirdRowSubheading"
	case ST_TableStyleTypePageFieldLabels:
		attr.Value = "pageFieldLabels"
	case ST_TableStyleTypePageFieldValues:
		attr.Value = "pageFieldValues"
	}
	return attr, nil
}

func (e *ST_TableStyleType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "wholeTable":
		*e = 1
	case "headerRow":
		*e = 2
	case "totalRow":
		*e = 3
	case "firstColumn":
		*e = 4
	case "lastColumn":
		*e = 5
	case "firstRowStripe":
		*e = 6
	case "secondRowStripe":
		*e = 7
	case "firstColumnStripe":
		*e = 8
	case "secondColumnStripe":
		*e = 9
	case "firstHeaderCell":
		*e = 10
	case "lastHeaderCell":
		*e = 11
	case "firstTotalCell":
		*e = 12
	case "lastTotalCell":
		*e = 13
	case "firstSubtotalColumn":
		*e = 14
	case "secondSubtotalColumn":
		*e = 15
	case "thirdSubtotalColumn":
		*e = 16
	case "firstSubtotalRow":
		*e = 17
	case "secondSubtotalRow":
		*e = 18
	case "thirdSubtotalRow":
		*e = 19
	case "blankRow":
		*e = 20
	case "firstColumnSubheading":
		*e = 21
	case "secondColumnSubheading":
		*e = 22
	case "thirdColumnSubheading":
		*e = 23
	case "firstRowSubheading":
		*e = 24
	case "secondRowSubheading":
		*e = 25
	case "thirdRowSubheading":
		*e = 26
	case "pageFieldLabels":
		*e = 27
	case "pageFieldValues":
		*e = 28
	}
	return nil
}

func (m ST_TableStyleType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TableStyleType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "wholeTable":
			*m = 1
		case "headerRow":
			*m = 2
		case "totalRow":
			*m = 3
		case "firstColumn":
			*m = 4
		case "lastColumn":
			*m = 5
		case "firstRowStripe":
			*m = 6
		case "secondRowStripe":
			*m = 7
		case "firstColumnStripe":
			*m = 8
		case "secondColumnStripe":
			*m = 9
		case "firstHeaderCell":
			*m = 10
		case "lastHeaderCell":
			*m = 11
		case "firstTotalCell":
			*m = 12
		case "lastTotalCell":
			*m = 13
		case "firstSubtotalColumn":
			*m = 14
		case "secondSubtotalColumn":
			*m = 15
		case "thirdSubtotalColumn":
			*m = 16
		case "firstSubtotalRow":
			*m = 17
		case "secondSubtotalRow":
			*m = 18
		case "thirdSubtotalRow":
			*m = 19
		case "blankRow":
			*m = 20
		case "firstColumnSubheading":
			*m = 21
		case "secondColumnSubheading":
			*m = 22
		case "thirdColumnSubheading":
			*m = 23
		case "firstRowSubheading":
			*m = 24
		case "secondRowSubheading":
			*m = 25
		case "thirdRowSubheading":
			*m = 26
		case "pageFieldLabels":
			*m = 27
		case "pageFieldValues":
			*m = 28
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

func (m ST_TableStyleType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "wholeTable"
	case 2:
		return "headerRow"
	case 3:
		return "totalRow"
	case 4:
		return "firstColumn"
	case 5:
		return "lastColumn"
	case 6:
		return "firstRowStripe"
	case 7:
		return "secondRowStripe"
	case 8:
		return "firstColumnStripe"
	case 9:
		return "secondColumnStripe"
	case 10:
		return "firstHeaderCell"
	case 11:
		return "lastHeaderCell"
	case 12:
		return "firstTotalCell"
	case 13:
		return "lastTotalCell"
	case 14:
		return "firstSubtotalColumn"
	case 15:
		return "secondSubtotalColumn"
	case 16:
		return "thirdSubtotalColumn"
	case 17:
		return "firstSubtotalRow"
	case 18:
		return "secondSubtotalRow"
	case 19:
		return "thirdSubtotalRow"
	case 20:
		return "blankRow"
	case 21:
		return "firstColumnSubheading"
	case 22:
		return "secondColumnSubheading"
	case 23:
		return "thirdColumnSubheading"
	case 24:
		return "firstRowSubheading"
	case 25:
		return "secondRowSubheading"
	case 26:
		return "thirdRowSubheading"
	case 27:
		return "pageFieldLabels"
	case 28:
		return "pageFieldValues"
	}
	return ""
}

func (m ST_TableStyleType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TableStyleType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_FontScheme byte

const (
	ST_FontSchemeUnset ST_FontScheme = 0
	ST_FontSchemeNone  ST_FontScheme = 1
	ST_FontSchemeMajor ST_FontScheme = 2
	ST_FontSchemeMinor ST_FontScheme = 3
)

func (e ST_FontScheme) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_FontSchemeUnset:
		attr.Value = ""
	case ST_FontSchemeNone:
		attr.Value = "none"
	case ST_FontSchemeMajor:
		attr.Value = "major"
	case ST_FontSchemeMinor:
		attr.Value = "minor"
	}
	return attr, nil
}

func (e *ST_FontScheme) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "major":
		*e = 2
	case "minor":
		*e = 3
	}
	return nil
}

func (m ST_FontScheme) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_FontScheme) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "major":
			*m = 2
		case "minor":
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

func (m ST_FontScheme) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "major"
	case 3:
		return "minor"
	}
	return ""
}

func (m ST_FontScheme) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FontScheme) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_UnderlineValues byte

const (
	ST_UnderlineValuesUnset            ST_UnderlineValues = 0
	ST_UnderlineValuesSingle           ST_UnderlineValues = 1
	ST_UnderlineValuesDouble           ST_UnderlineValues = 2
	ST_UnderlineValuesSingleAccounting ST_UnderlineValues = 3
	ST_UnderlineValuesDoubleAccounting ST_UnderlineValues = 4
	ST_UnderlineValuesNone             ST_UnderlineValues = 5
)

func (e ST_UnderlineValues) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_UnderlineValuesUnset:
		attr.Value = ""
	case ST_UnderlineValuesSingle:
		attr.Value = "single"
	case ST_UnderlineValuesDouble:
		attr.Value = "double"
	case ST_UnderlineValuesSingleAccounting:
		attr.Value = "singleAccounting"
	case ST_UnderlineValuesDoubleAccounting:
		attr.Value = "doubleAccounting"
	case ST_UnderlineValuesNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_UnderlineValues) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "single":
		*e = 1
	case "double":
		*e = 2
	case "singleAccounting":
		*e = 3
	case "doubleAccounting":
		*e = 4
	case "none":
		*e = 5
	}
	return nil
}

func (m ST_UnderlineValues) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_UnderlineValues) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "singleAccounting":
			*m = 3
		case "doubleAccounting":
			*m = 4
		case "none":
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

func (m ST_UnderlineValues) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "single"
	case 2:
		return "double"
	case 3:
		return "singleAccounting"
	case 4:
		return "doubleAccounting"
	case 5:
		return "none"
	}
	return ""
}

func (m ST_UnderlineValues) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_UnderlineValues) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_DdeValueType byte

const (
	ST_DdeValueTypeUnset ST_DdeValueType = 0
	ST_DdeValueTypeNil   ST_DdeValueType = 1
	ST_DdeValueTypeB     ST_DdeValueType = 2
	ST_DdeValueTypeN     ST_DdeValueType = 3
	ST_DdeValueTypeE     ST_DdeValueType = 4
	ST_DdeValueTypeStr   ST_DdeValueType = 5
)

func (e ST_DdeValueType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_DdeValueTypeUnset:
		attr.Value = ""
	case ST_DdeValueTypeNil:
		attr.Value = "nil"
	case ST_DdeValueTypeB:
		attr.Value = "b"
	case ST_DdeValueTypeN:
		attr.Value = "n"
	case ST_DdeValueTypeE:
		attr.Value = "e"
	case ST_DdeValueTypeStr:
		attr.Value = "str"
	}
	return attr, nil
}

func (e *ST_DdeValueType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "nil":
		*e = 1
	case "b":
		*e = 2
	case "n":
		*e = 3
	case "e":
		*e = 4
	case "str":
		*e = 5
	}
	return nil
}

func (m ST_DdeValueType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_DdeValueType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "nil":
			*m = 1
		case "b":
			*m = 2
		case "n":
			*m = 3
		case "e":
			*m = 4
		case "str":
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

func (m ST_DdeValueType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "nil"
	case 2:
		return "b"
	case 3:
		return "n"
	case 4:
		return "e"
	case 5:
		return "str"
	}
	return ""
}

func (m ST_DdeValueType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_DdeValueType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TableType byte

const (
	ST_TableTypeUnset      ST_TableType = 0
	ST_TableTypeWorksheet  ST_TableType = 1
	ST_TableTypeXml        ST_TableType = 2
	ST_TableTypeQueryTable ST_TableType = 3
)

func (e ST_TableType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TableTypeUnset:
		attr.Value = ""
	case ST_TableTypeWorksheet:
		attr.Value = "worksheet"
	case ST_TableTypeXml:
		attr.Value = "xml"
	case ST_TableTypeQueryTable:
		attr.Value = "queryTable"
	}
	return attr, nil
}

func (e *ST_TableType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "worksheet":
		*e = 1
	case "xml":
		*e = 2
	case "queryTable":
		*e = 3
	}
	return nil
}

func (m ST_TableType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TableType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "worksheet":
			*m = 1
		case "xml":
			*m = 2
		case "queryTable":
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

func (m ST_TableType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "worksheet"
	case 2:
		return "xml"
	case 3:
		return "queryTable"
	}
	return ""
}

func (m ST_TableType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TableType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TotalsRowFunction byte

const (
	ST_TotalsRowFunctionUnset     ST_TotalsRowFunction = 0
	ST_TotalsRowFunctionNone      ST_TotalsRowFunction = 1
	ST_TotalsRowFunctionSum       ST_TotalsRowFunction = 2
	ST_TotalsRowFunctionMin       ST_TotalsRowFunction = 3
	ST_TotalsRowFunctionMax       ST_TotalsRowFunction = 4
	ST_TotalsRowFunctionAverage   ST_TotalsRowFunction = 5
	ST_TotalsRowFunctionCount     ST_TotalsRowFunction = 6
	ST_TotalsRowFunctionCountNums ST_TotalsRowFunction = 7
	ST_TotalsRowFunctionStdDev    ST_TotalsRowFunction = 8
	ST_TotalsRowFunctionVar       ST_TotalsRowFunction = 9
	ST_TotalsRowFunctionCustom    ST_TotalsRowFunction = 10
)

func (e ST_TotalsRowFunction) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TotalsRowFunctionUnset:
		attr.Value = ""
	case ST_TotalsRowFunctionNone:
		attr.Value = "none"
	case ST_TotalsRowFunctionSum:
		attr.Value = "sum"
	case ST_TotalsRowFunctionMin:
		attr.Value = "min"
	case ST_TotalsRowFunctionMax:
		attr.Value = "max"
	case ST_TotalsRowFunctionAverage:
		attr.Value = "average"
	case ST_TotalsRowFunctionCount:
		attr.Value = "count"
	case ST_TotalsRowFunctionCountNums:
		attr.Value = "countNums"
	case ST_TotalsRowFunctionStdDev:
		attr.Value = "stdDev"
	case ST_TotalsRowFunctionVar:
		attr.Value = "var"
	case ST_TotalsRowFunctionCustom:
		attr.Value = "custom"
	}
	return attr, nil
}

func (e *ST_TotalsRowFunction) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "none":
		*e = 1
	case "sum":
		*e = 2
	case "min":
		*e = 3
	case "max":
		*e = 4
	case "average":
		*e = 5
	case "count":
		*e = 6
	case "countNums":
		*e = 7
	case "stdDev":
		*e = 8
	case "var":
		*e = 9
	case "custom":
		*e = 10
	}
	return nil
}

func (m ST_TotalsRowFunction) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TotalsRowFunction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "sum":
			*m = 2
		case "min":
			*m = 3
		case "max":
			*m = 4
		case "average":
			*m = 5
		case "count":
			*m = 6
		case "countNums":
			*m = 7
		case "stdDev":
			*m = 8
		case "var":
			*m = 9
		case "custom":
			*m = 10
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

func (m ST_TotalsRowFunction) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "none"
	case 2:
		return "sum"
	case 3:
		return "min"
	case 4:
		return "max"
	case 5:
		return "average"
	case 6:
		return "count"
	case 7:
		return "countNums"
	case 8:
		return "stdDev"
	case 9:
		return "var"
	case 10:
		return "custom"
	}
	return ""
}

func (m ST_TotalsRowFunction) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TotalsRowFunction) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_VolDepType byte

const (
	ST_VolDepTypeUnset         ST_VolDepType = 0
	ST_VolDepTypeRealTimeData  ST_VolDepType = 1
	ST_VolDepTypeOlapFunctions ST_VolDepType = 2
)

func (e ST_VolDepType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VolDepTypeUnset:
		attr.Value = ""
	case ST_VolDepTypeRealTimeData:
		attr.Value = "realTimeData"
	case ST_VolDepTypeOlapFunctions:
		attr.Value = "olapFunctions"
	}
	return attr, nil
}

func (e *ST_VolDepType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "realTimeData":
		*e = 1
	case "olapFunctions":
		*e = 2
	}
	return nil
}

func (m ST_VolDepType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VolDepType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "realTimeData":
			*m = 1
		case "olapFunctions":
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

func (m ST_VolDepType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "realTimeData"
	case 2:
		return "olapFunctions"
	}
	return ""
}

func (m ST_VolDepType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VolDepType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_VolValueType byte

const (
	ST_VolValueTypeUnset ST_VolValueType = 0
	ST_VolValueTypeB     ST_VolValueType = 1
	ST_VolValueTypeN     ST_VolValueType = 2
	ST_VolValueTypeE     ST_VolValueType = 3
	ST_VolValueTypeS     ST_VolValueType = 4
)

func (e ST_VolValueType) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VolValueTypeUnset:
		attr.Value = ""
	case ST_VolValueTypeB:
		attr.Value = "b"
	case ST_VolValueTypeN:
		attr.Value = "n"
	case ST_VolValueTypeE:
		attr.Value = "e"
	case ST_VolValueTypeS:
		attr.Value = "s"
	}
	return attr, nil
}

func (e *ST_VolValueType) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "b":
		*e = 1
	case "n":
		*e = 2
	case "e":
		*e = 3
	case "s":
		*e = 4
	}
	return nil
}

func (m ST_VolValueType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_VolValueType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "b":
			*m = 1
		case "n":
			*m = 2
		case "e":
			*m = 3
		case "s":
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

func (m ST_VolValueType) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "b"
	case 2:
		return "n"
	case 3:
		return "e"
	case 4:
		return "s"
	}
	return ""
}

func (m ST_VolValueType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_VolValueType) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Visibility byte

const (
	ST_VisibilityUnset      ST_Visibility = 0
	ST_VisibilityVisible    ST_Visibility = 1
	ST_VisibilityHidden     ST_Visibility = 2
	ST_VisibilityVeryHidden ST_Visibility = 3
)

func (e ST_Visibility) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_VisibilityUnset:
		attr.Value = ""
	case ST_VisibilityVisible:
		attr.Value = "visible"
	case ST_VisibilityHidden:
		attr.Value = "hidden"
	case ST_VisibilityVeryHidden:
		attr.Value = "veryHidden"
	}
	return attr, nil
}

func (e *ST_Visibility) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "visible":
		*e = 1
	case "hidden":
		*e = 2
	case "veryHidden":
		*e = 3
	}
	return nil
}

func (m ST_Visibility) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Visibility) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "visible":
			*m = 1
		case "hidden":
			*m = 2
		case "veryHidden":
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

func (m ST_Visibility) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "visible"
	case 2:
		return "hidden"
	case 3:
		return "veryHidden"
	}
	return ""
}

func (m ST_Visibility) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Visibility) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Comments byte

const (
	ST_CommentsUnset             ST_Comments = 0
	ST_CommentsCommNone          ST_Comments = 1
	ST_CommentsCommIndicator     ST_Comments = 2
	ST_CommentsCommIndAndComment ST_Comments = 3
)

func (e ST_Comments) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CommentsUnset:
		attr.Value = ""
	case ST_CommentsCommNone:
		attr.Value = "commNone"
	case ST_CommentsCommIndicator:
		attr.Value = "commIndicator"
	case ST_CommentsCommIndAndComment:
		attr.Value = "commIndAndComment"
	}
	return attr, nil
}

func (e *ST_Comments) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "commNone":
		*e = 1
	case "commIndicator":
		*e = 2
	case "commIndAndComment":
		*e = 3
	}
	return nil
}

func (m ST_Comments) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Comments) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "commNone":
			*m = 1
		case "commIndicator":
			*m = 2
		case "commIndAndComment":
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

func (m ST_Comments) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "commNone"
	case 2:
		return "commIndicator"
	case 3:
		return "commIndAndComment"
	}
	return ""
}

func (m ST_Comments) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Comments) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_Objects byte

const (
	ST_ObjectsUnset        ST_Objects = 0
	ST_ObjectsAll          ST_Objects = 1
	ST_ObjectsPlaceholders ST_Objects = 2
	ST_ObjectsNone         ST_Objects = 3
)

func (e ST_Objects) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_ObjectsUnset:
		attr.Value = ""
	case ST_ObjectsAll:
		attr.Value = "all"
	case ST_ObjectsPlaceholders:
		attr.Value = "placeholders"
	case ST_ObjectsNone:
		attr.Value = "none"
	}
	return attr, nil
}

func (e *ST_Objects) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "all":
		*e = 1
	case "placeholders":
		*e = 2
	case "none":
		*e = 3
	}
	return nil
}

func (m ST_Objects) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_Objects) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "all":
			*m = 1
		case "placeholders":
			*m = 2
		case "none":
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

func (m ST_Objects) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "all"
	case 2:
		return "placeholders"
	case 3:
		return "none"
	}
	return ""
}

func (m ST_Objects) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_Objects) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SheetState byte

const (
	ST_SheetStateUnset      ST_SheetState = 0
	ST_SheetStateVisible    ST_SheetState = 1
	ST_SheetStateHidden     ST_SheetState = 2
	ST_SheetStateVeryHidden ST_SheetState = 3
)

func (e ST_SheetState) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SheetStateUnset:
		attr.Value = ""
	case ST_SheetStateVisible:
		attr.Value = "visible"
	case ST_SheetStateHidden:
		attr.Value = "hidden"
	case ST_SheetStateVeryHidden:
		attr.Value = "veryHidden"
	}
	return attr, nil
}

func (e *ST_SheetState) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "visible":
		*e = 1
	case "hidden":
		*e = 2
	case "veryHidden":
		*e = 3
	}
	return nil
}

func (m ST_SheetState) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SheetState) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "visible":
			*m = 1
		case "hidden":
			*m = 2
		case "veryHidden":
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

func (m ST_SheetState) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "visible"
	case 2:
		return "hidden"
	case 3:
		return "veryHidden"
	}
	return ""
}

func (m ST_SheetState) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SheetState) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_UpdateLinks byte

const (
	ST_UpdateLinksUnset   ST_UpdateLinks = 0
	ST_UpdateLinksUserSet ST_UpdateLinks = 1
	ST_UpdateLinksNever   ST_UpdateLinks = 2
	ST_UpdateLinksAlways  ST_UpdateLinks = 3
)

func (e ST_UpdateLinks) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_UpdateLinksUnset:
		attr.Value = ""
	case ST_UpdateLinksUserSet:
		attr.Value = "userSet"
	case ST_UpdateLinksNever:
		attr.Value = "never"
	case ST_UpdateLinksAlways:
		attr.Value = "always"
	}
	return attr, nil
}

func (e *ST_UpdateLinks) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "userSet":
		*e = 1
	case "never":
		*e = 2
	case "always":
		*e = 3
	}
	return nil
}

func (m ST_UpdateLinks) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_UpdateLinks) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "userSet":
			*m = 1
		case "never":
			*m = 2
		case "always":
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

func (m ST_UpdateLinks) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "userSet"
	case 2:
		return "never"
	case 3:
		return "always"
	}
	return ""
}

func (m ST_UpdateLinks) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_UpdateLinks) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_SmartTagShow byte

const (
	ST_SmartTagShowUnset       ST_SmartTagShow = 0
	ST_SmartTagShowAll         ST_SmartTagShow = 1
	ST_SmartTagShowNone        ST_SmartTagShow = 2
	ST_SmartTagShowNoIndicator ST_SmartTagShow = 3
)

func (e ST_SmartTagShow) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_SmartTagShowUnset:
		attr.Value = ""
	case ST_SmartTagShowAll:
		attr.Value = "all"
	case ST_SmartTagShowNone:
		attr.Value = "none"
	case ST_SmartTagShowNoIndicator:
		attr.Value = "noIndicator"
	}
	return attr, nil
}

func (e *ST_SmartTagShow) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "all":
		*e = 1
	case "none":
		*e = 2
	case "noIndicator":
		*e = 3
	}
	return nil
}

func (m ST_SmartTagShow) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_SmartTagShow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "all":
			*m = 1
		case "none":
			*m = 2
		case "noIndicator":
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

func (m ST_SmartTagShow) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "all"
	case 2:
		return "none"
	case 3:
		return "noIndicator"
	}
	return ""
}

func (m ST_SmartTagShow) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_SmartTagShow) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_CalcMode byte

const (
	ST_CalcModeUnset       ST_CalcMode = 0
	ST_CalcModeManual      ST_CalcMode = 1
	ST_CalcModeAuto        ST_CalcMode = 2
	ST_CalcModeAutoNoTable ST_CalcMode = 3
)

func (e ST_CalcMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_CalcModeUnset:
		attr.Value = ""
	case ST_CalcModeManual:
		attr.Value = "manual"
	case ST_CalcModeAuto:
		attr.Value = "auto"
	case ST_CalcModeAutoNoTable:
		attr.Value = "autoNoTable"
	}
	return attr, nil
}

func (e *ST_CalcMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "manual":
		*e = 1
	case "auto":
		*e = 2
	case "autoNoTable":
		*e = 3
	}
	return nil
}

func (m ST_CalcMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_CalcMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "manual":
			*m = 1
		case "auto":
			*m = 2
		case "autoNoTable":
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

func (m ST_CalcMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "manual"
	case 2:
		return "auto"
	case 3:
		return "autoNoTable"
	}
	return ""
}

func (m ST_CalcMode) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_CalcMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_RefMode byte

const (
	ST_RefModeUnset ST_RefMode = 0
	ST_RefModeA1    ST_RefMode = 1
	ST_RefModeR1C1  ST_RefMode = 2
)

func (e ST_RefMode) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_RefModeUnset:
		attr.Value = ""
	case ST_RefModeA1:
		attr.Value = "A1"
	case ST_RefModeR1C1:
		attr.Value = "R1C1"
	}
	return attr, nil
}

func (e *ST_RefMode) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "A1":
		*e = 1
	case "R1C1":
		*e = 2
	}
	return nil
}

func (m ST_RefMode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_RefMode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "A1":
			*m = 1
		case "R1C1":
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

func (m ST_RefMode) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "A1"
	case 2:
		return "R1C1"
	}
	return ""
}

func (m ST_RefMode) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_RefMode) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

type ST_TargetScreenSize byte

const (
	ST_TargetScreenSizeUnset     ST_TargetScreenSize = 0
	ST_TargetScreenSize544x376   ST_TargetScreenSize = 1
	ST_TargetScreenSize640x480   ST_TargetScreenSize = 2
	ST_TargetScreenSize720x512   ST_TargetScreenSize = 3
	ST_TargetScreenSize800x600   ST_TargetScreenSize = 4
	ST_TargetScreenSize1024x768  ST_TargetScreenSize = 5
	ST_TargetScreenSize1152x882  ST_TargetScreenSize = 6
	ST_TargetScreenSize1152x900  ST_TargetScreenSize = 7
	ST_TargetScreenSize1280x1024 ST_TargetScreenSize = 8
	ST_TargetScreenSize1600x1200 ST_TargetScreenSize = 9
	ST_TargetScreenSize1800x1440 ST_TargetScreenSize = 10
	ST_TargetScreenSize1920x1200 ST_TargetScreenSize = 11
)

func (e ST_TargetScreenSize) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	attr := xml.Attr{}
	attr.Name = name
	switch e {
	case ST_TargetScreenSizeUnset:
		attr.Value = ""
	case ST_TargetScreenSize544x376:
		attr.Value = "544x376"
	case ST_TargetScreenSize640x480:
		attr.Value = "640x480"
	case ST_TargetScreenSize720x512:
		attr.Value = "720x512"
	case ST_TargetScreenSize800x600:
		attr.Value = "800x600"
	case ST_TargetScreenSize1024x768:
		attr.Value = "1024x768"
	case ST_TargetScreenSize1152x882:
		attr.Value = "1152x882"
	case ST_TargetScreenSize1152x900:
		attr.Value = "1152x900"
	case ST_TargetScreenSize1280x1024:
		attr.Value = "1280x1024"
	case ST_TargetScreenSize1600x1200:
		attr.Value = "1600x1200"
	case ST_TargetScreenSize1800x1440:
		attr.Value = "1800x1440"
	case ST_TargetScreenSize1920x1200:
		attr.Value = "1920x1200"
	}
	return attr, nil
}

func (e *ST_TargetScreenSize) UnmarshalXMLAttr(attr xml.Attr) error {
	switch attr.Value {
	case "":
		*e = 0
	case "544x376":
		*e = 1
	case "640x480":
		*e = 2
	case "720x512":
		*e = 3
	case "800x600":
		*e = 4
	case "1024x768":
		*e = 5
	case "1152x882":
		*e = 6
	case "1152x900":
		*e = 7
	case "1280x1024":
		*e = 8
	case "1600x1200":
		*e = 9
	case "1800x1440":
		*e = 10
	case "1920x1200":
		*e = 11
	}
	return nil
}

func (m ST_TargetScreenSize) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return e.EncodeElement(m.String(), start)
}

func (m *ST_TargetScreenSize) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
		case "544x376":
			*m = 1
		case "640x480":
			*m = 2
		case "720x512":
			*m = 3
		case "800x600":
			*m = 4
		case "1024x768":
			*m = 5
		case "1152x882":
			*m = 6
		case "1152x900":
			*m = 7
		case "1280x1024":
			*m = 8
		case "1600x1200":
			*m = 9
		case "1800x1440":
			*m = 10
		case "1920x1200":
			*m = 11
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

func (m ST_TargetScreenSize) String() string {
	switch m {
	case 0:
		return ""
	case 1:
		return "544x376"
	case 2:
		return "640x480"
	case 3:
		return "720x512"
	case 4:
		return "800x600"
	case 5:
		return "1024x768"
	case 6:
		return "1152x882"
	case 7:
		return "1152x900"
	case 8:
		return "1280x1024"
	case 9:
		return "1600x1200"
	case 10:
		return "1800x1440"
	case 11:
		return "1920x1200"
	}
	return ""
}

func (m ST_TargetScreenSize) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TargetScreenSize) ValidateWithPath(path string) error {
	switch m {
	case 0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11:
	default:
		return fmt.Errorf("%s: out of range value %d", path, int(m))
	}
	return nil
}

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_AutoFilter", NewCT_AutoFilter)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FilterColumn", NewCT_FilterColumn)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Filters", NewCT_Filters)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Filter", NewCT_Filter)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomFilters", NewCT_CustomFilters)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomFilter", NewCT_CustomFilter)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Top10", NewCT_Top10)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ColorFilter", NewCT_ColorFilter)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_IconFilter", NewCT_IconFilter)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DynamicFilter", NewCT_DynamicFilter)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SortState", NewCT_SortState)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SortCondition", NewCT_SortCondition)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DateGroupItem", NewCT_DateGroupItem)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_XStringElement", NewCT_XStringElement)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Extension", NewCT_Extension)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ObjectAnchor", NewCT_ObjectAnchor)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExtensionList", NewCT_ExtensionList)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CalcChain", NewCT_CalcChain)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CalcCell", NewCT_CalcCell)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Comments", NewCT_Comments)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Authors", NewCT_Authors)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CommentList", NewCT_CommentList)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Comment", NewCT_Comment)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CommentPr", NewCT_CommentPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MapInfo", NewCT_MapInfo)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Schema", NewCT_Schema)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Map", NewCT_Map)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DataBinding", NewCT_DataBinding)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Connections", NewCT_Connections)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Connection", NewCT_Connection)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DbPr", NewCT_DbPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_OlapPr", NewCT_OlapPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_WebPr", NewCT_WebPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Parameters", NewCT_Parameters)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Parameter", NewCT_Parameter)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Tables", NewCT_Tables)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TableMissing", NewCT_TableMissing)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TextPr", NewCT_TextPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TextFields", NewCT_TextFields)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TextField", NewCT_TextField)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotCacheDefinition", NewCT_PivotCacheDefinition)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CacheFields", NewCT_CacheFields)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CacheField", NewCT_CacheField)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CacheSource", NewCT_CacheSource)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_WorksheetSource", NewCT_WorksheetSource)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Consolidation", NewCT_Consolidation)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Pages", NewCT_Pages)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PCDSCPage", NewCT_PCDSCPage)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PageItem", NewCT_PageItem)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RangeSets", NewCT_RangeSets)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RangeSet", NewCT_RangeSet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SharedItems", NewCT_SharedItems)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Missing", NewCT_Missing)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Number", NewCT_Number)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Boolean", NewCT_Boolean)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Error", NewCT_Error)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_String", NewCT_String)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DateTime", NewCT_DateTime)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FieldGroup", NewCT_FieldGroup)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RangePr", NewCT_RangePr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DiscretePr", NewCT_DiscretePr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_GroupItems", NewCT_GroupItems)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotCacheRecords", NewCT_PivotCacheRecords)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Record", NewCT_Record)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PCDKPIs", NewCT_PCDKPIs)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PCDKPI", NewCT_PCDKPI)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CacheHierarchies", NewCT_CacheHierarchies)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CacheHierarchy", NewCT_CacheHierarchy)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FieldsUsage", NewCT_FieldsUsage)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FieldUsage", NewCT_FieldUsage)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_GroupLevels", NewCT_GroupLevels)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_GroupLevel", NewCT_GroupLevel)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Groups", NewCT_Groups)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_LevelGroup", NewCT_LevelGroup)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_GroupMembers", NewCT_GroupMembers)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_GroupMember", NewCT_GroupMember)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TupleCache", NewCT_TupleCache)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ServerFormat", NewCT_ServerFormat)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ServerFormats", NewCT_ServerFormats)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PCDSDTCEntries", NewCT_PCDSDTCEntries)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Tuples", NewCT_Tuples)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Tuple", NewCT_Tuple)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Sets", NewCT_Sets)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Set", NewCT_Set)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_QueryCache", NewCT_QueryCache)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Query", NewCT_Query)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CalculatedItems", NewCT_CalculatedItems)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CalculatedItem", NewCT_CalculatedItem)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CalculatedMembers", NewCT_CalculatedMembers)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CalculatedMember", NewCT_CalculatedMember)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_pivotTableDefinition", NewCT_pivotTableDefinition)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Location", NewCT_Location)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotFields", NewCT_PivotFields)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotField", NewCT_PivotField)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_AutoSortScope", NewCT_AutoSortScope)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Items", NewCT_Items)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Item", NewCT_Item)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PageFields", NewCT_PageFields)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PageField", NewCT_PageField)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DataFields", NewCT_DataFields)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DataField", NewCT_DataField)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_rowItems", NewCT_rowItems)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_colItems", NewCT_colItems)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_I", NewCT_I)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_X", NewCT_X)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RowFields", NewCT_RowFields)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ColFields", NewCT_ColFields)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Field", NewCT_Field)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Formats", NewCT_Formats)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Format", NewCT_Format)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ConditionalFormats", NewCT_ConditionalFormats)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ConditionalFormat", NewCT_ConditionalFormat)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotAreas", NewCT_PivotAreas)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ChartFormats", NewCT_ChartFormats)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ChartFormat", NewCT_ChartFormat)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotHierarchies", NewCT_PivotHierarchies)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotHierarchy", NewCT_PivotHierarchy)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RowHierarchiesUsage", NewCT_RowHierarchiesUsage)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ColHierarchiesUsage", NewCT_ColHierarchiesUsage)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_HierarchyUsage", NewCT_HierarchyUsage)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MemberProperties", NewCT_MemberProperties)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MemberProperty", NewCT_MemberProperty)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Members", NewCT_Members)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Member", NewCT_Member)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Dimensions", NewCT_Dimensions)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotDimension", NewCT_PivotDimension)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MeasureGroups", NewCT_MeasureGroups)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MeasureDimensionMaps", NewCT_MeasureDimensionMaps)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MeasureGroup", NewCT_MeasureGroup)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MeasureDimensionMap", NewCT_MeasureDimensionMap)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotTableStyle", NewCT_PivotTableStyle)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotFilters", NewCT_PivotFilters)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotFilter", NewCT_PivotFilter)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotArea", NewCT_PivotArea)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotAreaReferences", NewCT_PivotAreaReferences)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotAreaReference", NewCT_PivotAreaReference)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Index", NewCT_Index)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_QueryTable", NewCT_QueryTable)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_QueryTableRefresh", NewCT_QueryTableRefresh)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_QueryTableDeletedFields", NewCT_QueryTableDeletedFields)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DeletedField", NewCT_DeletedField)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_QueryTableFields", NewCT_QueryTableFields)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_QueryTableField", NewCT_QueryTableField)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Sst", NewCT_Sst)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PhoneticRun", NewCT_PhoneticRun)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RElt", NewCT_RElt)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RPrElt", NewCT_RPrElt)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Rst", NewCT_Rst)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PhoneticPr", NewCT_PhoneticPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionHeaders", NewCT_RevisionHeaders)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Revisions", NewCT_Revisions)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionHeader", NewCT_RevisionHeader)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetIdMap", NewCT_SheetIdMap)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetId", NewCT_SheetId)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ReviewedRevisions", NewCT_ReviewedRevisions)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Reviewed", NewCT_Reviewed)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_UndoInfo", NewCT_UndoInfo)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionRowColumn", NewCT_RevisionRowColumn)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionMove", NewCT_RevisionMove)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionCustomView", NewCT_RevisionCustomView)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionSheetRename", NewCT_RevisionSheetRename)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionInsertSheet", NewCT_RevisionInsertSheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionCellChange", NewCT_RevisionCellChange)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionFormatting", NewCT_RevisionFormatting)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionAutoFormatting", NewCT_RevisionAutoFormatting)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionComment", NewCT_RevisionComment)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionDefinedName", NewCT_RevisionDefinedName)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionConflict", NewCT_RevisionConflict)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RevisionQueryTableField", NewCT_RevisionQueryTableField)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Users", NewCT_Users)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SharedUser", NewCT_SharedUser)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Macrosheet", NewCT_Macrosheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Dialogsheet", NewCT_Dialogsheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Worksheet", NewCT_Worksheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetData", NewCT_SheetData)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetCalcPr", NewCT_SheetCalcPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetFormatPr", NewCT_SheetFormatPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Cols", NewCT_Cols)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Col", NewCT_Col)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Row", NewCT_Row)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Cell", NewCT_Cell)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetPr", NewCT_SheetPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetDimension", NewCT_SheetDimension)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetViews", NewCT_SheetViews)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetView", NewCT_SheetView)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Pane", NewCT_Pane)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotSelection", NewCT_PivotSelection)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Selection", NewCT_Selection)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PageBreak", NewCT_PageBreak)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Break", NewCT_Break)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_OutlinePr", NewCT_OutlinePr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PageSetUpPr", NewCT_PageSetUpPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DataConsolidate", NewCT_DataConsolidate)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DataRefs", NewCT_DataRefs)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DataRef", NewCT_DataRef)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MergeCells", NewCT_MergeCells)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MergeCell", NewCT_MergeCell)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SmartTags", NewCT_SmartTags)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellSmartTags", NewCT_CellSmartTags)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellSmartTag", NewCT_CellSmartTag)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellSmartTagPr", NewCT_CellSmartTagPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Drawing", NewCT_Drawing)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_LegacyDrawing", NewCT_LegacyDrawing)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DrawingHF", NewCT_DrawingHF)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomSheetViews", NewCT_CustomSheetViews)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomSheetView", NewCT_CustomSheetView)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DataValidations", NewCT_DataValidations)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DataValidation", NewCT_DataValidation)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ConditionalFormatting", NewCT_ConditionalFormatting)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CfRule", NewCT_CfRule)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Hyperlinks", NewCT_Hyperlinks)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Hyperlink", NewCT_Hyperlink)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellFormula", NewCT_CellFormula)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ColorScale", NewCT_ColorScale)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DataBar", NewCT_DataBar)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_IconSet", NewCT_IconSet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Cfvo", NewCT_Cfvo)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PageMargins", NewCT_PageMargins)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PrintOptions", NewCT_PrintOptions)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PageSetup", NewCT_PageSetup)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_HeaderFooter", NewCT_HeaderFooter)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Scenarios", NewCT_Scenarios)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetProtection", NewCT_SheetProtection)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ProtectedRanges", NewCT_ProtectedRanges)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ProtectedRange", NewCT_ProtectedRange)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Scenario", NewCT_Scenario)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_InputCells", NewCT_InputCells)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellWatches", NewCT_CellWatches)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellWatch", NewCT_CellWatch)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Chartsheet", NewCT_Chartsheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ChartsheetPr", NewCT_ChartsheetPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ChartsheetViews", NewCT_ChartsheetViews)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ChartsheetView", NewCT_ChartsheetView)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ChartsheetProtection", NewCT_ChartsheetProtection)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CsPageSetup", NewCT_CsPageSetup)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomChartsheetViews", NewCT_CustomChartsheetViews)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomChartsheetView", NewCT_CustomChartsheetView)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomProperties", NewCT_CustomProperties)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomProperty", NewCT_CustomProperty)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_OleObjects", NewCT_OleObjects)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_OleObject", NewCT_OleObject)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ObjectPr", NewCT_ObjectPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_WebPublishItems", NewCT_WebPublishItems)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_WebPublishItem", NewCT_WebPublishItem)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Controls", NewCT_Controls)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Control", NewCT_Control)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ControlPr", NewCT_ControlPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_IgnoredErrors", NewCT_IgnoredErrors)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_IgnoredError", NewCT_IgnoredError)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TableParts", NewCT_TableParts)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TablePart", NewCT_TablePart)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Metadata", NewCT_Metadata)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MetadataTypes", NewCT_MetadataTypes)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MetadataType", NewCT_MetadataType)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MetadataBlocks", NewCT_MetadataBlocks)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MetadataBlock", NewCT_MetadataBlock)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MetadataRecord", NewCT_MetadataRecord)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FutureMetadata", NewCT_FutureMetadata)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FutureMetadataBlock", NewCT_FutureMetadataBlock)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MdxMetadata", NewCT_MdxMetadata)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Mdx", NewCT_Mdx)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MdxTuple", NewCT_MdxTuple)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MdxSet", NewCT_MdxSet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MdxMemeberProp", NewCT_MdxMemeberProp)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MdxKPI", NewCT_MdxKPI)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MetadataStringIndex", NewCT_MetadataStringIndex)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MetadataStrings", NewCT_MetadataStrings)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SingleXmlCells", NewCT_SingleXmlCells)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SingleXmlCell", NewCT_SingleXmlCell)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_XmlCellPr", NewCT_XmlCellPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_XmlPr", NewCT_XmlPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Stylesheet", NewCT_Stylesheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellAlignment", NewCT_CellAlignment)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Borders", NewCT_Borders)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Border", NewCT_Border)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_BorderPr", NewCT_BorderPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellProtection", NewCT_CellProtection)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Fonts", NewCT_Fonts)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Fills", NewCT_Fills)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Fill", NewCT_Fill)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PatternFill", NewCT_PatternFill)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Color", NewCT_Color)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_GradientFill", NewCT_GradientFill)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_GradientStop", NewCT_GradientStop)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_NumFmts", NewCT_NumFmts)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_NumFmt", NewCT_NumFmt)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellStyleXfs", NewCT_CellStyleXfs)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellXfs", NewCT_CellXfs)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Xf", NewCT_Xf)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellStyles", NewCT_CellStyles)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CellStyle", NewCT_CellStyle)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Dxfs", NewCT_Dxfs)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Dxf", NewCT_Dxf)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Colors", NewCT_Colors)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_IndexedColors", NewCT_IndexedColors)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_MRUColors", NewCT_MRUColors)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_RgbColor", NewCT_RgbColor)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TableStyles", NewCT_TableStyles)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TableStyle", NewCT_TableStyle)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TableStyleElement", NewCT_TableStyleElement)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_BooleanProperty", NewCT_BooleanProperty)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FontSize", NewCT_FontSize)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_IntProperty", NewCT_IntProperty)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FontName", NewCT_FontName)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_VerticalAlignFontProperty", NewCT_VerticalAlignFontProperty)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FontScheme", NewCT_FontScheme)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_UnderlineProperty", NewCT_UnderlineProperty)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Font", NewCT_Font)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FontFamily", NewCT_FontFamily)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalLink", NewCT_ExternalLink)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalBook", NewCT_ExternalBook)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalSheetNames", NewCT_ExternalSheetNames)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalSheetName", NewCT_ExternalSheetName)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalDefinedNames", NewCT_ExternalDefinedNames)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalDefinedName", NewCT_ExternalDefinedName)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalSheetDataSet", NewCT_ExternalSheetDataSet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalSheetData", NewCT_ExternalSheetData)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalRow", NewCT_ExternalRow)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalCell", NewCT_ExternalCell)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DdeLink", NewCT_DdeLink)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DdeItems", NewCT_DdeItems)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DdeItem", NewCT_DdeItem)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DdeValues", NewCT_DdeValues)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DdeValue", NewCT_DdeValue)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_OleLink", NewCT_OleLink)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_OleItems", NewCT_OleItems)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_OleItem", NewCT_OleItem)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Table", NewCT_Table)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TableStyleInfo", NewCT_TableStyleInfo)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TableColumns", NewCT_TableColumns)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TableColumn", NewCT_TableColumn)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_TableFormula", NewCT_TableFormula)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_XmlColumnPr", NewCT_XmlColumnPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_VolTypes", NewCT_VolTypes)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_VolType", NewCT_VolType)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_VolMain", NewCT_VolMain)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_VolTopic", NewCT_VolTopic)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_VolTopicRef", NewCT_VolTopicRef)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Workbook", NewCT_Workbook)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FileVersion", NewCT_FileVersion)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_BookViews", NewCT_BookViews)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_BookView", NewCT_BookView)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomWorkbookViews", NewCT_CustomWorkbookViews)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CustomWorkbookView", NewCT_CustomWorkbookView)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Sheets", NewCT_Sheets)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_Sheet", NewCT_Sheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_WorkbookPr", NewCT_WorkbookPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SmartTagPr", NewCT_SmartTagPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SmartTagTypes", NewCT_SmartTagTypes)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SmartTagType", NewCT_SmartTagType)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FileRecoveryPr", NewCT_FileRecoveryPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_CalcPr", NewCT_CalcPr)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DefinedNames", NewCT_DefinedNames)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_DefinedName", NewCT_DefinedName)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalReferences", NewCT_ExternalReferences)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_ExternalReference", NewCT_ExternalReference)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_SheetBackgroundPicture", NewCT_SheetBackgroundPicture)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotCaches", NewCT_PivotCaches)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_PivotCache", NewCT_PivotCache)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FileSharing", NewCT_FileSharing)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_OleSize", NewCT_OleSize)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_WorkbookProtection", NewCT_WorkbookProtection)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_WebPublishing", NewCT_WebPublishing)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FunctionGroups", NewCT_FunctionGroups)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_FunctionGroup", NewCT_FunctionGroup)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_WebPublishObjects", NewCT_WebPublishObjects)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "CT_WebPublishObject", NewCT_WebPublishObject)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "calcChain", NewCalcChain)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "comments", NewComments)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "MapInfo", NewMapInfo)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "connections", NewConnections)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "pivotCacheDefinition", NewPivotCacheDefinition)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "pivotCacheRecords", NewPivotCacheRecords)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "pivotTableDefinition", NewPivotTableDefinition)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "queryTable", NewQueryTable)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "sst", NewSst)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "headers", NewHeaders)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "revisions", NewRevisions)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "users", NewUsers)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "worksheet", NewWorksheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "chartsheet", NewChartsheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "dialogsheet", NewDialogsheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "metadata", NewMetadata)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "singleXmlCells", NewSingleXmlCells)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "styleSheet", NewStyleSheet)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "externalLink", NewExternalLink)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "table", NewTable)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "volTypes", NewVolTypes)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "workbook", NewWorkbook)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "EG_ExtensionList", NewEG_ExtensionList)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "AG_RevData", NewAG_RevData)
	office.RegisterConstructor("http://schemas.openxmlformats.org/spreadsheetml/2006/main", "AG_AutoFormat", NewAG_AutoFormat)
}

type ST_Sqref []string
type ST_CellSpans []string
