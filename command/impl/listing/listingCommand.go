package listing

import (
	"errors"
	"fmt"

	"github.com/danwakefield/fnmatch"
	"github.com/logrusorgru/aurora"

	"myracloud.com/myra-shell/api/types"
	"myracloud.com/myra-shell/command"
	"myracloud.com/myra-shell/console"
	"myracloud.com/myra-shell/context"
)

//
// ListingCommand ...
//
type ListingCommand struct {
	*command.BaseCommand
}

//
// Completer ..,
//
func (cmd *ListingCommand) Completer(ctx context.Container, input string) []string {
	return []string{}
}

//
// BuildListing ...
//
func (cmd *ListingCommand) BuildListing(
	ctx context.Container,
	filter string,
	apiFilter string,
	additionalItems bool,
) (RowListing, error) {
	var rows RowListing
	var dors = ctx.FindSelection(context.AreaSubDomain, context.AreaDomain)
	connector := ctx.GetAPIConnector()

	switch ctx.GetSelection() {
	case context.AreaNone:
		tmp := ctx.Get("userList").([]console.Login)

		for _, d := range tmp {
			r := &loginRow{
				data: d,
			}

			if fnmatch.Match(filter, r.GetName(), 0) {
				rows = append(rows, r)
			}
		}

	case context.AreaUser:
		tmp, err := connector.DomainList(apiFilter)

		if err != nil {
			return rows, err
		}

		for _, d := range tmp {
			r := &domainRow{data: d}

			if fnmatch.Match(filter, r.GetName(), 0) {
				rows = append(rows, r)
			}
		}

	case context.AreaSubDomain:
		if additionalItems {
			rows = append(rows,
				newContextSwitchRowEx(context.ContextUp, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextCacheSettings, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextRedirect, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextIPFilter, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextSettings, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextStatistics, 3, true, "r-x--"),
			)
		}

	case context.AreaDomain:
		tmp, err := connector.DNSRecordList(ctx.Identifier(), &apiFilter, nil, false)

		if err != nil {
			return rows, err
		}

		if additionalItems {
			rows = append(rows,
				newContextSwitchRowEx(context.ContextUp, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextCacheSettings, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextRedirect, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextErrorPages, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextIPFilter, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextSettings, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextSsl, 3, true, "r-x--"),
				newContextSwitchRowEx(context.ContextStatistics, 3, true, "r-x--"),
			)
		}

		for _, d := range tmp {
			r := &dnsrecordrow{data: d}

			if fnmatch.Match(filter, r.GetName(), 0) {
				rows = append(rows, r)
			}
		}

	case context.AreaIPFilter:
		tmp, err := connector.IPFilterList(dors.Identifier(), apiFilter)

		if err != nil {
			return rows, err
		}

		if additionalItems {
			rows = append(rows, newContextSwitchRow(
				context.ContextUp, 3, true, "0", types.DateTimeNow().ToUnixDate(), "",
			))
		}

		for _, d := range tmp {
			r := &ipfilterrow{d}

			if fnmatch.Match(filter, r.GetName(), 0) {
				rows = append(rows, r)
			}
		}

	case context.AreaCache:
		tmp, err := connector.CacheSettingList(dors.Identifier(), apiFilter)

		if err != nil {
			return rows, err
		}

		if additionalItems {
			rows = append(rows, newContextSwitchRow(context.ContextUp, 4, true))
		}

		for _, d := range tmp {
			r := &cacheSettingsRow{d}

			if fnmatch.Match(filter, r.GetName(), 0) {
				rows = append(rows, r)
			}
		}

	case context.AreaRedirect:
		tmp, err := connector.RedirectList(dors.Identifier(), apiFilter)

		if err != nil {
			return rows, err
		}

		if additionalItems {
			rows = append(rows, newContextSwitchRow(context.ContextUp, 4, true))
		}

		for _, d := range tmp {
			r := &redirectRow{d}

			if fnmatch.Match(filter, r.GetName(), 0) {
				rows = append(rows, r)
			}
		}

	case context.AreaErrorPage:
		tmp, err := connector.ErrorPage(dors.Identifier(), apiFilter)

		if err != nil {
			return rows, err
		}

		if additionalItems {
			rows = append(rows, newContextSwitchRow(
				context.ContextUp,
				3,
				true,
				"0",
				types.DateTimeNow().ToUnixDate(),
				"",
			))
		}

		for _, ep := range tmp {
			if fnmatch.Match(filter, ep.SubDomainName, 0) {
				for _, er := range ep.Error {
					rows = append(
						rows,
						&errorpageRow{
							subDomainName: ep.SubDomainName,
							data:          er,
						},
					)
				}
			}
		}

	case context.AreaSettings:
		tmp, err := connector.Settings(dors.Identifier())

		if err != nil {
			return rows, err
		}

		if additionalItems {
			rows = append(rows, newContextSwitchRow(
				context.ContextUp,
				3,
				true,
				"0",
				types.DateTimeNow().ToUnixDate(),
				"",
			))
		}

		for key, val := range tmp {
			if fnmatch.Match(filter, key, 0) {
				rows = append(
					rows,
					&settingsRow{
						name:  key,
						value: val,
					},
				)
			}
		}

	case context.AreaSsl:
		tmp, err := connector.SslCertList(dors.GetName(), "")
		if err != nil {
			return rows, err
		}

		if additionalItems {
			rows = append(rows, newContextSwitchRow(
				context.ContextUp,
				3,
				true,
				"0",
				types.DateTimeNow().ToUnixDate(),
				"",
			))
		}

		for _, val := range tmp {
			r := &sslrow{val}

			if fnmatch.Match(filter, r.GetName(), 0) {
				rows = append(rows, r)
			}

		}

	default:
		return rows, errors.New("Do not know how to handle given context")
	}

	return rows, nil
}

//
// Execute ...
//
func (cmd *ListingCommand) Execute(ctx context.Container, buffer *string) (uint, error) {
	var filter = cmd.GetParameter(0, "*")
	var apiFilter = cmd.GetParameter(0, "")

	rows, err := cmd.BuildListing(ctx, filter, apiFilter, true)

	if err != nil {
		return 1, err
	}

	*buffer = cmd.OutputRows(rows)

	return 0, nil
}

//
// OutputRows ...
//
func (cmd *ListingCommand) OutputRows(rows RowListing) string {
	var ret string
	var long = cmd.IsFlagSet("long")
	var verbose = cmd.IsFlagSet("verbose")

	var colTpl = rows.buildFormatingTemplate(long, verbose)

	for _, row := range rows {
		cols := row.GetColumns(long, verbose)

		ret += fmt.Sprint(row.FormatFlags(), " ")
		ret += fmt.Sprintf(colTpl, cols...)

		if row.IsAvailableForContextSwitch() {
			ret += fmt.Sprintln(aurora.Colorize(row.GetName(), aurora.BlueFg|aurora.BoldFm))
		} else {
			ret += fmt.Sprintln(row.GetName())
		}
	}

	return ret
}
