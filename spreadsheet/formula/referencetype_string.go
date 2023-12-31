package formula

import "fmt"

const _ReferenceType_name = "ReferenceTypeInvalidReferenceTypeCellReferenceTypeNamedRangeReferenceTypeRangeReferenceTypeSheet"

var _ReferenceType_index = [...]uint8{0, 20, 37, 60, 78, 96}

func (i ReferenceType) String() string {
	if i >= ReferenceType(len(_ReferenceType_index)-1) {
		return fmt.Sprintf("ReferenceType(%d)", i)
	}
	return _ReferenceType_name[_ReferenceType_index[i]:_ReferenceType_index[i+1]]
}
