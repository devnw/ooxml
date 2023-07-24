package document

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/measurement"
	"go.devnw.com/ooxml/schema/soo/ofc/sharedTypes"
	"go.devnw.com/ooxml/schema/soo/wml"
)

// ParagraphSpacing controls the spacing for a paragraph and its lines.
type ParagraphSpacing struct {
	x *wml.CT_Spacing
}

// SetBefore sets the spacing that comes before the paragraph.
func (p ParagraphSpacing) SetBefore(before measurement.Distance) {
	p.x.BeforeAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.BeforeAttr.ST_UnsignedDecimalNumber = office.Uint64(uint64(before / measurement.Twips))
}

// SetAfter sets the spacing that comes after the paragraph.
func (p ParagraphSpacing) SetAfter(after measurement.Distance) {
	p.x.AfterAttr = &sharedTypes.ST_TwipsMeasure{}
	p.x.AfterAttr.ST_UnsignedDecimalNumber = office.Uint64(uint64(after / measurement.Twips))
}

// SetLineSpacing sets the spacing between lines in a paragraph.
func (p ParagraphSpacing) SetLineSpacing(d measurement.Distance, rule wml.ST_LineSpacingRule) {
	if rule == wml.ST_LineSpacingRuleUnset {
		p.x.LineRuleAttr = wml.ST_LineSpacingRuleUnset
		p.x.LineAttr = nil
	} else {
		p.x.LineRuleAttr = rule
		p.x.LineAttr = &wml.ST_SignedTwipsMeasure{}
		p.x.LineAttr.Int64 = office.Int64(int64(d / measurement.Twips))
	}
}

// SetBeforeAuto controls if spacing before a paragraph is automatically determined.
func (p ParagraphSpacing) SetBeforeAuto(b bool) {
	if b {
		p.x.BeforeAutospacingAttr = &sharedTypes.ST_OnOff{}
		p.x.BeforeAutospacingAttr.Bool = office.Bool(true)
	} else {
		p.x.BeforeAutospacingAttr = nil
	}
}

// SetAfterAuto controls if spacing after a paragraph is automatically determined.
func (p ParagraphSpacing) SetAfterAuto(b bool) {
	if b {
		p.x.AfterAutospacingAttr = &sharedTypes.ST_OnOff{}
		p.x.AfterAutospacingAttr.Bool = office.Bool(true)
	} else {
		p.x.AfterAutospacingAttr = nil
	}
}
