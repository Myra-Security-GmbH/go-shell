// +build !testing

package context

import (
	"testing"

	"myracloud.com/myra-shell/event"

	"github.com/stretchr/testify/require"
)

func buildMyraContainer(cc *myraContext) *myraContextContainer {
	return &myraContextContainer{
		GenericSubscriber: event.NewGenericSubscriber(),
		currentContext:    cc,
	}
}

func TestBuildPrompt(t *testing.T) {
	ctx := buildMyraContainer(&myraContext{
		id:        0,
		name:      "IPFilter",
		selection: AreaIPFilter,
		parent: &myraContext{
			id:        7331,
			name:      "SubDomain",
			selection: AreaSubDomain,
			parent: &myraContext{
				id:        1337,
				name:      "DOMAIN",
				selection: AreaDomain,
				parent: &myraContext{
					id:        0,
					name:      "NONE",
					selection: AreaNone,
				},
			},
		},
	})

	require.Equal(t, "/DOMAIN/SubDomain/IPFilter> ", ctx.BuildPrompt(false))
	require.Equal(t, "/DOMAIN> ", ctx.FindSelection(AreaDomain).BuildPrompt(false))

	require.Equal(t,
		"\x1b[1;32m/\x1b[0m\x1b[1;34mDOMAIN\x1b[0m"+
			"\x1b[1;32m/\x1b[0m\x1b[1;36mSubDomain\x1b[0m"+
			"\x1b[1;32m/\x1b[0m\x1b[1;33mIPFilter\x1b[0m> ",
		ctx.BuildPrompt(true),
	)

	ctx = buildMyraContainer(nil)

	require.Equal(t, "> ", ctx.BuildPrompt(false))
	require.Equal(t, "> ", ctx.BuildPrompt(true))
}

func TestStringer(t *testing.T) {
	ctx := buildMyraContainer(&myraContext{
		id:        7331,
		name:      "SubDomain",
		selection: AreaSubDomain,
		parent: &myraContext{
			id:        1337,
			name:      "DOMAIN",
			selection: AreaDomain,
			parent: &myraContext{
				id:        0,
				name:      "NONE",
				selection: AreaNone,
			},
		},
	})

	require.Equal(
		t,
		"[id=1337, selection=4] DOMAIN",
		ctx.FindSelection(AreaDomain).String(),
	)

	require.NotEmpty(t, ctx.String())
}

func TestIdentifier(t *testing.T) {
	ctx := buildMyraContainer(&myraContext{
		id:        1337,
		name:      "domain.de",
		selection: AreaDomain,
	})

	require.Equal(t, "ALL:domain.de", ctx.Identifier())
	require.Equal(t, "ALL:domain.de", ctx.FindSelection(AreaDomain).Identifier())

	ctx.SwitchDown(7331, "www.domain.de", AreaSubDomain)

	require.Equal(t, "www.domain.de", ctx.Identifier())
	require.Equal(t, "www.domain.de", ctx.FindSelection(AreaSubDomain).Identifier())

	ctx = buildMyraContainer(nil)
	require.Empty(t, ctx.Identifier())
}

func TestSwitchError(t *testing.T) {
	ctx := buildMyraContainer(nil)

	e, err := ctx.SwitchUp()

	require.Nil(t, e)
	require.Error(t, err)

	ctx.SwitchDown(0, "TEST", AreaDomain)

	e, err = ctx.SwitchUp()

	require.Nil(t, e)
	require.Error(t, err)
}

func TestSwitch(t *testing.T) {
	ctx := buildMyraContainer(nil)
	ctx.SwitchDown(0, "DOMAIN", AreaDomain)

	require.Nil(t, ctx.GetParent())
	require.Equal(t, uint64(0), ctx.GetID())
	require.Equal(t, "DOMAIN", ctx.GetName())
	require.Equal(t, uint(AreaDomain), ctx.GetSelection())

	e, err := ctx.SwitchDown(1, "IPFILTER", AreaIPFilter)

	require.Nil(t, err)
	require.NotNil(t, ctx.GetParent())
	require.Equal(t, uint64(1), ctx.GetID())
	require.Equal(t, "IPFILTER", ctx.GetName())
	require.Equal(t, uint(AreaIPFilter), ctx.GetSelection())
	require.Equal(t, e.GetID(), ctx.GetID())
	require.Equal(t, e.GetSelection(), ctx.GetSelection())
	require.Equal(t, e.GetName(), ctx.GetName())

	e, err = ctx.SwitchUp()

	require.Nil(t, err)
	require.Nil(t, ctx.GetParent())
	require.Equal(t, uint64(0), ctx.GetID())
	require.Equal(t, "DOMAIN", ctx.GetName())
	require.Equal(t, uint(AreaDomain), ctx.GetSelection())
	require.Equal(t, e.GetID(), ctx.GetID())
	require.Equal(t, e.GetSelection(), ctx.GetSelection())
	require.Equal(t, e.GetName(), ctx.GetName())
}

func TestFindIdentifier(t *testing.T) {
	ctx := buildMyraContainer(&myraContext{
		selection: AreaIPFilter,
		name:      "IPFILTER",
		parent: &myraContext{
			selection: AreaDomain,
			name:      "DOMAIN",
			parent: &myraContext{
				selection: AreaNone,
				name:      "NONE",
			},
		},
	})

	e := ctx.FindSelection(AreaDomain)
	require.Equal(t, ctx.GetParent(), e)
}

func TestFindIdentifier2(t *testing.T) {
	ctx := buildMyraContainer(&myraContext{
		selection: AreaIPFilter,
		name:      "IPFILTER",
		parent: &myraContext{
			selection: AreaDomain,
			name:      "DOMAIN",
		},
	})

	e := ctx.FindSelection(AreaDomain)

	require.Equal(t, ctx.GetParent(), e)
}

func TestFindOne(t *testing.T) {
	ctx := buildMyraContainer(&myraContext{
		selection: AreaIPFilter,
		parent: &myraContext{
			selection: AreaSubDomain,
			name:      "SUBDOMAIN",
			parent: &myraContext{
				selection: AreaDomain,
				name:      "DOMAIN",
				parent: &myraContext{
					selection: AreaNone,
					name:      "NONE",
				},
			},
		},
	})

	require.Equal(t, ctx.GetParent(), ctx.FindSelection(AreaSubDomain, AreaDomain))
	require.Equal(t, ctx.GetParent(), ctx.FindSelection(AreaDomain, AreaSubDomain))
	require.Equal(t, ctx.GetParent().GetParent(), ctx.FindSelection(AreaDomain))
	require.Nil(t, ctx.FindSelection(AreaRedirect))
}
