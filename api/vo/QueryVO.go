package vo

//
// QueryVO ...
//
type QueryVO struct {
	Error    bool `json:"error"`
	Page     uint `json:"page"`
	Count    uint `json:"count"`
	PageSize uint `json:"pageSize"`
}
