package listing

import "github.com/Myra-Security-GmbH/myra-shell/api/types"

type genericRow struct {
	name           string
	cols           []interface{}
	contextSwitch  bool
	amountLongCols int
	formatFlags    string
}

func (r *genericRow) GetID() uint64 {
	return 0
}

func (r *genericRow) FormatFlags() string {
	if r.formatFlags == "" {
		return "--" +
			formatFlag('x', '-', r.contextSwitch)
	}

	return r.formatFlags
}

func (r *genericRow) IsAvailableForContextSwitch() bool {
	return r.contextSwitch
}

func (r *genericRow) GetName() string {
	return r.name
}

func (r *genericRow) GetColumns(long bool, verbose bool) []interface{} {
	ret := r.cols

	if long {
		ret = make([]interface{}, r.amountLongCols)

		for i := 0; i < r.amountLongCols; i++ {
			ret[i] = ""
		}

		for idx, i := range r.cols {
			ret[idx] = i
		}
	}

	return ret
}

func (r *genericRow) GetEntity() interface{} {
	return nil
}

//
// NewContextSwitchRow ...
//
func newContextSwitchRow(name string, amountLongCols int, ctxSwitch bool, cols ...interface{}) *genericRow {
	if len(cols) == 0 {
		cols = []interface{}{
			"0",
			types.DateTimeNow().ToUnixDate(),
		}
	}

	return &genericRow{
		name:           name,
		contextSwitch:  ctxSwitch,
		amountLongCols: amountLongCols,
		cols:           cols,
	}
}

func newContextSwitchRowEx(name string, amountLongCols int, ctxSwitch bool, flags string, cols ...interface{}) *genericRow {
	e := newContextSwitchRow(name, amountLongCols, ctxSwitch, cols...)

	e.formatFlags = flags

	return e
}
