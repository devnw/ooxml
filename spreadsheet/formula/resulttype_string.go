package formula

import "fmt"

const _ResultType_name = "ResultTypeUnknownResultTypeNumberResultTypeStringResultTypeListResultTypeArrayResultTypeErrorResultTypeEmpty"

var _ResultType_index = [...]uint8{0, 17, 33, 49, 63, 78, 93, 108}

func (i ResultType) String() string {
	if i >= ResultType(len(_ResultType_index)-1) {
		return fmt.Sprintf("ResultType(%d)", i)
	}
	return _ResultType_name[_ResultType_index[i]:_ResultType_index[i+1]]
}
