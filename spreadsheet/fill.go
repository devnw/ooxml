package spreadsheet

import (
	"go.devnw.com/ooxml/schema/soo/sml"
)

type Fill struct {
	x     *sml.CT_Fill
	fills *sml.CT_Fills
}

func (f Fill) Index() uint32 {
	// in differential formats, CT_Fill is not held in a CT_Fills and index
	// doesn't mean anything
	if f.fills == nil {
		return 0
	}

	for i, sf := range f.fills.Fill {
		if f.x == sf {
			return uint32(i)
		}
	}
	return 0
}

func (f Fill) SetPatternFill() PatternFill {
	f.x.GradientFill = nil
	f.x.PatternFill = sml.NewCT_PatternFill()
	f.x.PatternFill.PatternTypeAttr = sml.ST_PatternTypeSolid
	return PatternFill{f.x.PatternFill, f.x}
}
