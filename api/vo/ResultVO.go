package vo

//
// ResultVO ...
//
type ResultVO struct {
	Error         bool          `json:"error"`
	ViolationList []ViolationVO `json:"violationList"`
}
