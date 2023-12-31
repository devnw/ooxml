package format

import "fmt"

const _FmtType_name = "FmtTypeLiteralFmtTypeDigitFmtTypeDigitOptFmtTypeCommaFmtTypeDecimalFmtTypePercentFmtTypeDollarFmtTypeDigitOptThousandsFmtTypeUnderscoreFmtTypeDateFmtTypeTimeFmtTypeFractionFmtTypeText"

var _FmtType_index = [...]uint8{0, 14, 26, 41, 53, 67, 81, 94, 118, 135, 146, 157, 172, 183}

func (i FmtType) String() string {
	if i >= FmtType(len(_FmtType_index)-1) {
		return fmt.Sprintf("FmtType(%d)", i)
	}
	return _FmtType_name[_FmtType_index[i]:_FmtType_index[i+1]]
}
