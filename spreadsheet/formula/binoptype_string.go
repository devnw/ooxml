package formula

import "fmt"

const _BinOpType_name = "BinOpTypeUnknownBinOpTypePlusBinOpTypeMinusBinOpTypeMultBinOpTypeDivBinOpTypeExpBinOpTypeLTBinOpTypeGTBinOpTypeEQBinOpTypeLEQBinOpTypeGEQBinOpTypeNEBinOpTypeConcat"

var _BinOpType_index = [...]uint8{0, 16, 29, 43, 56, 68, 80, 91, 102, 113, 125, 137, 148, 163}

func (i BinOpType) String() string {
	if i >= BinOpType(len(_BinOpType_index)-1) {
		return fmt.Sprintf("BinOpType(%d)", i)
	}
	return _BinOpType_name[_BinOpType_index[i]:_BinOpType_index[i+1]]
}
