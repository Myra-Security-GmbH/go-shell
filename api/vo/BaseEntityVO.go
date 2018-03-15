package vo

import "myracloud.com/myra-shell/api/types"

//
// BaseEntityVO ...
//
type BaseEntityVO struct {
	Modified *types.DateTime `json:"modified,omitempty"`
	Created  *types.DateTime `json:"created,omitempty"`
	ID       *uint64         `json:"id,omitempty"`
}
