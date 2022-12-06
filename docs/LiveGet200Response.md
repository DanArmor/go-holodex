# LiveGet200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** |  | [optional] 
**Items** | Pointer to [**[]Video**](Video.md) |  | [optional] 

## Methods

### NewLiveGet200Response

`func NewLiveGet200Response() *LiveGet200Response`

NewLiveGet200Response instantiates a new LiveGet200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLiveGet200ResponseWithDefaults

`func NewLiveGet200ResponseWithDefaults() *LiveGet200Response`

NewLiveGet200ResponseWithDefaults instantiates a new LiveGet200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *LiveGet200Response) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LiveGet200Response) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LiveGet200Response) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LiveGet200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *LiveGet200Response) GetItems() []Video`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *LiveGet200Response) GetItemsOk() (*[]Video, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *LiveGet200Response) SetItems(v []Video)`

SetItems sets Items field to given value.

### HasItems

`func (o *LiveGet200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


