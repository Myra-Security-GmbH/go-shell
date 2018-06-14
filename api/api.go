package api

import (
	"github.com/Myra-Security-GmbH/myra-shell/api/types"
	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
)

//
// API ...
//
type API interface {
	GetLanguage() string
	SetLanguage(language string)

	GetEndpoint() string
	SetEndpoint(endpoint string)

	GetAPIKey() string
	SetAPIKey(APIKey string)

	GetSecret() string
	SetSecret(secret string)

	SaveEntity(identifer string, entity interface{}) error
	RemoveEntity(identifier string, entity interface{}) error
	CacheSettingList(domain string, search string) ([]vo.CacheSettingVO, error)
	DomainList(search string) ([]vo.DomainVO, error)
	DomainByName(name string) (*vo.DomainVO, error)
	IPFilterList(domain string, search string) ([]vo.IPFilterVO, error)
	RedirectList(domain string, filter string) ([]vo.RedirectVO, error)
	Settings(domain string) (vo.SettingsVO, error)

	GetStatistic(domain string,
		interval string,
		startDate *types.DateTime,
		endDate *types.DateTime,
		source []vo.DatasourceVO,
	) (*vo.StatisticVO, error)

	CacheClear(domain string,
		fqdn string,
		pattern string,
		recursive bool,
	) ([]vo.CacheClearVO, error)

	DNSRecordList(domain string,
		filter *string,
		types *[]string,
		activeOnly bool,
	) ([]vo.DNSRecordVO, error)

	ErrorPage(domain string, search string) ([]vo.ErrorPageVOList, error)
	SslCertList(domain string, search string) ([]vo.SslCertVO, error)
}
