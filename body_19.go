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

type Body19 struct {

	// Phone numbers to be reserved/un-reserved
	Records []ReservePhoneNumberRequestReserveRecord `json:"records,omitempty"`
}