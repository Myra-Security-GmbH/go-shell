package eventHandler

import (
	"errors"

	"myracloud.com/myra-shell/config"
	"myracloud.com/myra-shell/container"
	"myracloud.com/myra-shell/context"
	"myracloud.com/myra-shell/event"
)

//
// SwitchContextEvent ...
//
func SwitchContextEvent(source event.Subscriber, params ...interface{}) error {
	if len(params) != 3 {
		return errors.New("Parameter count does not match")
	}

	//from := params[0].(context.Context)
	to := params[2].(context.Context)
	ctxContainer := params[0].(context.Container)

	apiConnector := ctxContainer.GetAPIConnector()

	switch {
	case (to.GetSelection() == context.AreaNone):
		apiConnector.SetAPIKey("")
		apiConnector.SetSecret("")

	//
	// Switch to a user
	//
	case (to.GetSelection() == context.AreaUser):
		loginList := container.GetServiceEx("config").(*config.Config).Login

		for _, l := range loginList {
			if l.User == to.GetName() {
				apiConnector.SetAPIKey(l.APIKey)
				apiConnector.SetSecret(l.Secret)
				break
			}
		}

		list, err := apiConnector.DomainList("")

		if err != nil {
			return err
		}

		to.Set("domainList", list)

	//
	// Switch to a domain
	//
	case (to.GetSelection() == context.AreaDomain):
		filter := ""

		list, err := apiConnector.DNSRecordList(
			to.Identifier(),
			&filter,
			&[]string{"A", "AAAA", "CNAME"},
			false,
		)

		if err != nil {
			return err
		}

		to.Set("subDomainList", list)
	}

	return nil
}
