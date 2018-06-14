package vo

import "github.com/Myra-Security-GmbH/myra-shell/api/types"

//
// BaseEntityVO ...
//
type BaseEntityVO struct {
	Modified *types.DateTime `json:"modified,omitempty"`
	Created  *types.DateTime `json:"created,omitempty"`
	ID       *uint64         `json:"id,omitempty"`
}
