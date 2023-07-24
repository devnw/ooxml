package chart

// Axis is the interface implemented by different axes when assigning to a
// chart.
type Axis interface {
	AxisID() uint32
}

type nullAxis byte

func (n nullAxis) AxisID() uint32 {
	return 0
}

// NullAxis is a null axis with an ID of zero
var NullAxis Axis = nullAxis(0)
