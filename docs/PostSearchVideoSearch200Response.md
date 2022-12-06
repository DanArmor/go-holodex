# PostSearchVideoSearch200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** |  | [optional] 
**Items** | Pointer to [**[]VideoWithChannel**](VideoWithChannel.md) |  | [optional] 

## Methods

### NewPostSearchVideoSearch200Response

`func NewPostSearchVideoSearch200Response() *PostSearchVideoSearch200Response`

NewPostSearchVideoSearch200Response instantiates a new PostSearchVideoSearch200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSearchVideoSearch200ResponseWithDefaults

`func NewPostSearchVideoSearch200ResponseWithDefaults() *PostSearchVideoSearch200Response`

NewPostSearchVideoSearch200ResponseWithDefaults instantiates a new PostSearchVideoSearch200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *PostSearchVideoSearch200Response) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PostSearchVideoSearch200Response) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PostSearchVideoSearch200Response) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PostSearchVideoSearch200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetItems

`func (o *PostSearchVideoSearch200Response) GetItems() []VideoWithChannel`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PostSearchVideoSearch200Response) GetItemsOk() (*[]VideoWithChannel, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PostSearchVideoSearch200Response) SetItems(v []VideoWithChannel)`

SetItems sets Items field to given value.

### HasItems

`func (o *PostSearchVideoSearch200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


