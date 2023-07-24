package formula

// Evaluator is the interface for a formula evaluator.  This is needed so we can
// pass it to the spreadsheet to let it evaluate formula cells before returning
// the results.
// NOTE: in order to implement Evaluator without cache embed noCache in it.
type Evaluator interface {
	Eval(ctx Context, formula string) Result
	SetCache(key string, value Result)
	GetFromCache(key string) (Result, bool)
	LastEvalIsRef() bool
}

// NewEvaluator constructs a new defEval object which is the default formula evaluator.
func NewEvaluator() Evaluator {
	ev := &defEval{}
	ev.evCache = newEvCache()
	return ev
}
