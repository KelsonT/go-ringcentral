/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ringcentral

type ParsedNumberResource struct {
	Uri                    string          `json:"uri,omitempty"`
	OriginalString         string          `json:"originalString,omitempty"`
	Country                CountryResource `json:"country,omitempty"`
	AreaCode               string          `json:"areaCode,omitempty"`
	SubscriberNumber       string          `json:"subscriberNumber,omitempty"`
	FormattedNational      string          `json:"formattedNational,omitempty"`
	FormattedInternational string          `json:"formattedInternational,omitempty"`
	Dialable               string          `json:"dialable,omitempty"`
	E164                   string          `json:"e164,omitempty"`
	Special                bool            `json:"special,omitempty"`
	Normalized             string          `json:"normalized,omitempty"`
	TollFree               bool            `json:"tollFree,omitempty"`
	SubAddress             string          `json:"subAddress,omitempty"`
	DtmfPostfix            string          `json:"dtmfPostfix,omitempty"`
}
