package output

//
// Success ...
//
type Success interface {
	Success() string
}

type successMessage struct {
	s string
}

func (e *successMessage) Success() string {
	return e.s
}

//
// NewSuccess ...
//
func NewSuccess(text string) Success {
	return &successMessage{text}
}
