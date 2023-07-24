package document

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
	"go.devnw.com/ooxml/schema/soo/wml"
)

// Numbering is the document wide numbering styles contained in numbering.xml.
type Numbering struct {
	x *wml.Numbering
}

// NewNumbering constructs a new numbering.
func NewNumbering() Numbering {
	n := wml.NewNumbering()
	return Numbering{n}
}

// X returns the inner wrapped XML type.
func (n Numbering) X() *wml.Numbering {
	return n.x
}

// Clear resets the numbering.
func (n Numbering) Clear() {
	n.x.AbstractNum = nil
	n.x.Num = nil
	n.x.NumIdMacAtCleanup = nil
	n.x.NumPicBullet = nil
}

// InitializeDefault constructs a default numbering.
func (n Numbering) InitializeDefault() {
	abs := wml.NewCT_AbstractNum()
	abs.MultiLevelType = wml.NewCT_MultiLevelType()
	abs.MultiLevelType.ValAttr = wml.ST_MultiLevelTypeHybridMultilevel

	n.x.AbstractNum = append(n.x.AbstractNum, abs)
	abs.AbstractNumIdAttr = 1
	const indentStart = 720
	const indentDelta = 720
	const hangingIndent = 360
	for i := 0; i < 9; i++ {
		lvl := wml.NewCT_Lvl()
		lvl.IlvlAttr = int64(i)
		lvl.Start = wml.NewCT_DecimalNumber()
		lvl.Start.ValAttr = 1

		lvl.NumFmt = wml.NewCT_NumFmt()
		lvl.NumFmt.ValAttr = wml.ST_NumberFormatBullet

		lvl.Suff = wml.NewCT_LevelSuffix()
		lvl.Suff.ValAttr = wml.ST_LevelSuffixNothing

		lvl.LvlText = wml.NewCT_LevelText()
		lvl.LvlText.ValAttr = office.String("")

		lvl.LvlJc = wml.NewCT_Jc()
		lvl.LvlJc.ValAttr = wml.ST_JcLeft

		lvl.RPr = wml.NewCT_RPr()
		lvl.RPr.RFonts = wml.NewCT_Fonts()
		lvl.RPr.RFonts.AsciiAttr = office.String("Symbol")
		lvl.RPr.RFonts.HAnsiAttr = office.String("Symbol")
		lvl.RPr.RFonts.HintAttr = wml.ST_HintDefault

		lvl.PPr = wml.NewCT_PPrGeneral()

		indent := int64(i*indentDelta + indentStart)
		lvl.PPr.Ind = wml.NewCT_Ind()
		lvl.PPr.Ind.LeftAttr = &wml.ST_SignedTwipsMeasure{}
		lvl.PPr.Ind.LeftAttr.Int64 = office.Int64(indent)
		lvl.PPr.Ind.HangingAttr = &sharedTypes.ST_TwipsMeasure{}
		lvl.PPr.Ind.HangingAttr.ST_UnsignedDecimalNumber = office.Uint64(uint64(hangingIndent))

		abs.Lvl = append(abs.Lvl, lvl)
	}
	num := wml.NewCT_Num()
	num.NumIdAttr = 1
	num.AbstractNumId = wml.NewCT_DecimalNumber()
	num.AbstractNumId.ValAttr = 1
	n.x.Num = append(n.x.Num, num)
}

// Definitions returns the defined numbering definitions.
func (n Numbering) Definitions() []NumberingDefinition {
	ret := []NumberingDefinition{}
	for _, n := range n.x.AbstractNum {
		ret = append(ret, NumberingDefinition{n})
	}
	return ret
}

// AddDefinition adds a new numbering definition.
func (n Numbering) AddDefinition() NumberingDefinition {
	nx := wml.NewCT_Num()

	nextAbstractNumID := int64(1)
	for _, nd := range n.Definitions() {
		if nd.AbstractNumberID() >= nextAbstractNumID {
			nextAbstractNumID = nd.AbstractNumberID() + 1
		}
	}
	nextNumID := int64(1)
	for _, n := range n.X().Num {
		if n.NumIdAttr >= nextNumID {
			nextNumID = n.NumIdAttr + 1
		}
	}
	nx.NumIdAttr = nextNumID
	nx.AbstractNumId = wml.NewCT_DecimalNumber()
	nx.AbstractNumId.ValAttr = nextAbstractNumID

	an := wml.NewCT_AbstractNum()
	an.AbstractNumIdAttr = nextAbstractNumID

	n.x.AbstractNum = append(n.x.AbstractNum, an)
	n.x.Num = append(n.x.Num, nx)
	return NumberingDefinition{an}
}
