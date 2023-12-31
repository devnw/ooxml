package spreadsheet

import "fmt"

const _SortOrder_name = "SortOrderAscendingSortOrderDescending"

var _SortOrder_index = [...]uint8{0, 18, 37}

func (i SortOrder) String() string {
	if i >= SortOrder(len(_SortOrder_index)-1) {
		return fmt.Sprintf("SortOrder(%d)", i)
	}
	return _SortOrder_name[_SortOrder_index[i]:_SortOrder_index[i+1]]
}
