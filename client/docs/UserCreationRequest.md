# UserCreationRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Status of a user | [optional] [default to false]
**Addresses** | [**[]AddressInfo**](AddressInfo.md) | User addresses | [optional] 
**Emails** | [**[]EmailInfo**](EmailInfo.md) | User email addresses | 
**ExternalId** | **string** | External identifier of a user | [optional] 
**Id** | **string** | Internal identifier of a user | [optional] 
**Name** | [**NameInfo**](NameInfo.md) |  | 
**PhoneNumbers** | [**[]PhoneNumberInfoRequest**](PhoneNumberInfoRequest.md) | User phone numbers | [optional] 
**Photos** | [**[]PhotoInfo**](PhotoInfo.md) |  | [optional] 
**Schemas** | **[]string** | Specification links | 
**Urnietfparamsscimschemasextensionenterprise20User** | [**EnterpriseUser**](EnterpriseUser.md) |  | [optional] 
**UserName** | **string** | User mailbox. Must be same as work type email address | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


