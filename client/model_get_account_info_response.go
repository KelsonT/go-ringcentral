/*
 * RingCentral Connect Platform API Explorer
 *
 * <p>This is a beta interactive API explorer for the RingCentral Connect Platform. To use this service, you will need to have an account with the proper credentials to generate an OAuth2 access token.</p><p><h2>Quick Start</h2></p><ol><li>1) Go to <b>Authentication > /oauth/token</b></li><li>2) Enter <b>app_key, app_secret, username, password</b> fields and then click \"Try it out!\"</li><li>3) Upon success, your access_token is loaded and you can access any form requiring authorization.</li></ol><h2>Links</h2><ul><li><a href=\"https://github.com/ringcentral\" target=\"_blank\">RingCentral SDKs on Github</a></li><li><a href=\"mailto:devsupport@ringcentral.com\">RingCentral Developer Support Email</a></li></ul>
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package ringcentral

type GetAccountInfoResponse struct {
	// Internal identifier of an account
	Id string `json:"id,omitempty"`
	// Canonical URI of an account
	Uri string `json:"uri,omitempty"`
	// Main phone number of the current account
	MainNumber string                   `json:"mainNumber,omitempty"`
	Operator   GetExtensionInfoResponse `json:"operator,omitempty"`
	// Additional account identifier, developed and applied by the client
	PartnerId   string      `json:"partnerId,omitempty"`
	ServiceInfo ServiceInfo `json:"serviceInfo,omitempty"`
	// Specifies account configuration wizard state (web service setup). The default value is 'NotStarted'
	SetupWizardState string `json:"setupWizardState,omitempty"`
	// Status of the current account
	Status           string            `json:"status,omitempty"`
	StatusInfo       AccountStatusInfo `json:"statusInfo,omitempty"`
	RegionalSettings RegionalSettings  `json:"regionalSettings,omitempty"`
	// Specifies whether an account is included into any federation of accounts or not
	Federated bool `json:"federated,omitempty"`
}
