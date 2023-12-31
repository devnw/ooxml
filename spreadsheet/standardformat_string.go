package spreadsheet

import "fmt"

const (
	_StandardFormat_name_0 = "StandardFormatGeneralStandardFormatWholeNumberStandardFormat2StandardFormat3StandardFormat4"
	_StandardFormat_name_1 = "StandardFormatPercentStandardFormat10StandardFormat11StandardFormat12StandardFormat13StandardFormatDateStandardFormat15StandardFormat16StandardFormat17StandardFormat18StandardFormatTimeStandardFormat20StandardFormat21StandardFormatDateTime"
	_StandardFormat_name_2 = "StandardFormat37StandardFormat38StandardFormat39StandardFormat40"
	_StandardFormat_name_3 = "StandardFormat45StandardFormat46StandardFormat47StandardFormat48StandardFormat49"
)

var (
	_StandardFormat_index_0 = [...]uint8{0, 21, 46, 61, 76, 91}
	_StandardFormat_index_1 = [...]uint8{0, 21, 37, 53, 69, 85, 103, 119, 135, 151, 167, 185, 201, 217, 239}
	_StandardFormat_index_2 = [...]uint8{0, 16, 32, 48, 64}
	_StandardFormat_index_3 = [...]uint8{0, 16, 32, 48, 64, 80}
)

func (i StandardFormat) String() string {
	switch {
	case 0 <= i && i <= 4:
		return _StandardFormat_name_0[_StandardFormat_index_0[i]:_StandardFormat_index_0[i+1]]
	case 9 <= i && i <= 22:
		i -= 9
		return _StandardFormat_name_1[_StandardFormat_index_1[i]:_StandardFormat_index_1[i+1]]
	case 37 <= i && i <= 40:
		i -= 37
		return _StandardFormat_name_2[_StandardFormat_index_2[i]:_StandardFormat_index_2[i+1]]
	case 45 <= i && i <= 49:
		i -= 45
		return _StandardFormat_name_3[_StandardFormat_index_3[i]:_StandardFormat_index_3[i+1]]
	default:
		return fmt.Sprintf("StandardFormat(%d)", i)
	}
}
