// +build !testing

package storage

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TT struct {
	Name string
	Val  string
}

func TestAdd(t *testing.T) {
	store := &Generic{
		data: make(map[string]interface{}),
	}

	store.Set("list", []interface{}{
		TT{Name: "t1", Val: "v1"},
		TT{Name: "t2", Val: "v2"},
		TT{Name: "t3", Val: "v3"},
		TT{Name: "t4", Val: "v4"},
	})

	require.NotZero(t, store.Len("list"))

	e, _ := store.IsStructElemAvailable("list", "Name", "t1")
	require.NotNil(t, e)
	require.Equal(t, "v1", e.(TT).Val)

	e, _ = store.IsStructElemAvailable("list", "Name", "t38")
	require.Nil(t, e)

	store.Add("list", TT{Name: "t38", Val: "v38"})

	e, _ = store.IsStructElemAvailable("list", "Name", "t38")
	require.NotNil(t, e)
	require.Equal(t, "v38", e.(TT).Val)

	store.Set("list2", "MUH")
	require.NotNil(t, store.Get("list2"))

	store.Clear("list")
	e, _ = store.IsStructElemAvailable("list", "Name", "t38")
	require.Nil(t, e)
	require.NotNil(t, store.Get("list2"))

	store.ClearAll()
	require.Nil(t, store.Get("list2"))
}
