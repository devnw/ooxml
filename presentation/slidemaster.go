package presentation

import (
	"strconv"
	"strings"

	"go.devnw.com/ooxml/common"
	"go.devnw.com/ooxml/schema/soo/pml"
)

// SlideMaster is the slide master for a presentation.
type SlideMaster struct {
	p    *Presentation
	rels common.Relationships
	x    *pml.SldMaster
}

// X returns the inner wrapped XML type.
func (s SlideMaster) X() *pml.SldMaster {
	return s.x
}

// SlideLayouts returns a slice of all layouts in SlideMaster.
func (s SlideMaster) SlideLayouts() []SlideLayout {
	nameToLayoutIdx := map[string]int{}
	layouts := []SlideLayout{}
	for _, r := range s.rels.Relationships() {
		idxTxt := strings.Replace(r.Target(), "../slideLayouts/slideLayout", "", -1)
		idxTxt = strings.Replace(idxTxt, ".xml", "", -1)
		if idx, err := strconv.ParseInt(idxTxt, 10, 32); err == nil {
			nameToLayoutIdx[r.ID()] = int(idx)
		}
	}

	for _, lid := range s.x.SldLayoutIdLst.SldLayoutId {
		if idx, ok := nameToLayoutIdx[lid.RIdAttr]; ok {
			lout := s.p.layouts[idx-1]
			layouts = append(layouts, SlideLayout{lout})
		}
	}
	return layouts
}
