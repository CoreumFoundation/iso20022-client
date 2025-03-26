package dtif

import "time"

type Header struct {
	DTI     string `json:"DTI"`
	DTIType int    `json:"DTIType"`
	DLTType int    `json:"DLTType,omitempty"`
}

type Record struct {
	Header Header `json:"Header"`
}

type JsonTime struct {
	time.Time
}

func (t *JsonTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}

type LoginConfig struct {
	Icon                       string `json:"icon"`
	AssetsUrl                  string `json:"assetsUrl"`
	Auth0Domain                string `json:"auth0Domain"`
	Auth0Tenant                string `json:"auth0Tenant"`
	ClientConfigurationBaseUrl string `json:"clientConfigurationBaseUrl"`
	CallbackOnLocationHash     bool   `json:"callbackOnLocationHash"`
	CallbackURL                string `json:"callbackURL"`
	Cdn                        string `json:"cdn"`
	ClientID                   string `json:"clientID"`
	Connection                 string `json:"connection"`
	Dict                       struct {
		Signin struct {
			Title string `json:"title"`
		} `json:"signin"`
	} `json:"dict"`
	ExtraParams struct {
		Protocol     string `json:"protocol"`
		ResponseType string `json:"response_type"`
		Reauth       string `json:"reauth"`
		Csrf         string `json:"_csrf"`
		Intstate     string `json:"_intstate"`
		State        string `json:"state"`
	} `json:"extraParams"`
	InternalOptions struct {
		Protocol     string `json:"protocol"`
		ResponseType string `json:"response_type"`
		Reauth       string `json:"reauth"`
		Csrf         string `json:"_csrf"`
		Intstate     string `json:"_intstate"`
		State        string `json:"state"`
	} `json:"internalOptions"`
	WidgetUrl           string `json:"widgetUrl"`
	IsThirdPartyClient  bool   `json:"isThirdPartyClient"`
	AuthorizationServer struct {
		Url    string `json:"url"`
		Issuer string `json:"issuer"`
	} `json:"authorizationServer"`
	Colors struct {
	} `json:"colors"`
}

//{"client_id":"eaYmwHhnZUgMhPUBuhC9GV867aqKDqon","redirect_uri":"https://dtif.org/wp-login.php?redirect_to=https://dtif.org/","tenant":"prod-dtif-org","response_type":"token id_token","_csrf":"cC5yZ0sK-iP7mE6VQO-y-QBWdc4G7HsnJu68","state":"hKFo2SBFU1Ryb0h2aFdfa0JwOFN2N3VKZ1FpWWhzNl92UDBsdaFupWxvZ2luo3RpZNkgUGxoN2xmQlJnejMxUW1zZWNIRU5CQzJNSzZLWTVZNzKjY2lk2SBlYVltd0hoblpVZ01oUFVCdWhDOUdWODY3YXFLRHFvbg","_intstate":"deprecated","username":"a","password":"a","nonce":"rC8xYDMd1ncElGwXCSRKzbvRvV52V1n8","connection":"DB-DTIF"}

type LoginPayload struct {
	ClientId     string `json:"client_id"`
	RedirectUri  string `json:"redirect_uri"`
	Tenant       string `json:"tenant"`
	ResponseType string `json:"response_type"`
	Csrf         string `json:"_csrf"`
	State        string `json:"state"`
	Intstate     string `json:"_intstate"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Nonce        string `json:"nonce"`
	Connection   string `json:"connection"`
}
