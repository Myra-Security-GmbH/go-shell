package api

import (
	"net/http"

	auth "github.com/Myra-Security-GmbH/myra-shell/api/authentication"
	"github.com/Myra-Security-GmbH/myra-shell/api/vo"
	"gopkg.in/go-playground/validator.v9"
)

//
// myraAPI ...
//
type myraAPI struct {
	apiKey   string
	endpoint string
	language string

	signature *auth.MyraSignature
	client    *http.Client
	validator *validator.Validate
}

//
// NewMyraAPI creates a new instance of API
//
func NewMyraAPI(APIKey string, Secret string, BaseURI string, Language string) (API, error) {
	signature := &auth.MyraSignature{
		Secret: Secret,
	}

	client := &http.Client{
		Transport: &http.Transport{},
	}

	validate := validator.New()

	validate.RegisterStructValidation(validateVO,
		vo.IPFilterVO{},
		vo.DNSRecordVO{},
		vo.CacheSettingVO{},
		vo.RedirectVO{},
	)

	return &myraAPI{
		apiKey:    APIKey,
		endpoint:  BaseURI,
		language:  Language,
		signature: signature,
		client:    client,
		validator: validate,
	}, nil
}

func (ma *myraAPI) GetSecret() string {
	return ma.signature.Secret
}

func (ma *myraAPI) SetAPIKey(apiKey string) {
	ma.apiKey = apiKey
}

func (ma *myraAPI) SetEndpoint(endpoint string) {
	ma.endpoint = endpoint
}

func (ma *myraAPI) SetLanguage(lang string) {
	ma.language = lang
}

func (ma *myraAPI) SetSecret(secret string) {
	ma.signature.Secret = secret
}

func (ma *myraAPI) GetLanguage() string {
	return ma.language
}

func (ma *myraAPI) GetEndpoint() string {
	return ma.endpoint
}

func (ma *myraAPI) GetAPIKey() string {
	return ma.apiKey
}
