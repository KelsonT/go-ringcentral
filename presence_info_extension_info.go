/* 
 * Ringcentral API
 *
 * RingCentral Connect Platform API
 *
 * OpenAPI spec version: v1.0
 * 
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package ringcentral

type PresenceInfoExtensionInfo struct {

	// Internal identifier of an extension
	Id string `json:"id,omitempty"`

	// Canonical URI of an extension
	Uri string `json:"uri,omitempty"`

	// Extension number (usually 3 or 4 digits)
	ExtensionNumber string `json:"extensionNumber,omitempty"`
}