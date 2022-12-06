# PostSearchCommentSearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sort** | **string** |  | [default to "newest"]
**Lang** | Pointer to **[]string** | If set, will filter clips to only show clips with these langs + all vtuber streams (provided &#x60;target&#x60; is not set to filter out streams) | [optional] 
**Target** | Pointer to **[]string** | Target types of videos | [optional] 
**Comment** | **string** | Find videos with comments containing specified string (case insensitive) | 
**Topic** | Pointer to **[]string** | Return videos that match one of the provided topics | [optional] 
**Vch** | Pointer to **[]string** | Videos with all of the specified channel  ids. If two or more channel IDs are specified, will only return their collabs, or if one channel is a clipper, it will only show clips of the other vtubers made by this clipper. | [optional] 
**Org** | Pointer to **[]string** | Videos of channels in any of the specified orgs, or clips that involve a channel in the specified org. | [optional] 
**Paginated** | Pointer to **bool** | If set at all, responds with total and items wrapping the array of objects | [optional] 
**Offset** | **float32** |  | 
**Limit** | **float32** |  | 

## Methods

### NewPostSearchCommentSearchRequest

`func NewPostSearchCommentSearchRequest(sort string, comment string, offset float32, limit float32, ) *PostSearchCommentSearchRequest`

NewPostSearchCommentSearchRequest instantiates a new PostSearchCommentSearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSearchCommentSearchRequestWithDefaults

`func NewPostSearchCommentSearchRequestWithDefaults() *PostSearchCommentSearchRequest`

NewPostSearchCommentSearchRequestWithDefaults instantiates a new PostSearchCommentSearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSort

`func (o *PostSearchCommentSearchRequest) GetSort() string`

GetSort returns the Sort field if non-nil, zero value otherwise.

### GetSortOk

`func (o *PostSearchCommentSearchRequest) GetSortOk() (*string, bool)`

GetSortOk returns a tuple with the Sort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSort

`func (o *PostSearchCommentSearchRequest) SetSort(v string)`

SetSort sets Sort field to given value.


### GetLang

`func (o *PostSearchCommentSearchRequest) GetLang() []string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *PostSearchCommentSearchRequest) GetLangOk() (*[]string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *PostSearchCommentSearchRequest) SetLang(v []string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *PostSearchCommentSearchRequest) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetTarget

`func (o *PostSearchCommentSearchRequest) GetTarget() []string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PostSearchCommentSearchRequest) GetTargetOk() (*[]string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PostSearchCommentSearchRequest) SetTarget(v []string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PostSearchCommentSearchRequest) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetComment

`func (o *PostSearchCommentSearchRequest) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *PostSearchCommentSearchRequest) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *PostSearchCommentSearchRequest) SetComment(v string)`

SetComment sets Comment field to given value.


### GetTopic

`func (o *PostSearchCommentSearchRequest) GetTopic() []string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *PostSearchCommentSearchRequest) GetTopicOk() (*[]string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *PostSearchCommentSearchRequest) SetTopic(v []string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *PostSearchCommentSearchRequest) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetVch

`func (o *PostSearchCommentSearchRequest) GetVch() []string`

GetVch returns the Vch field if non-nil, zero value otherwise.

### GetVchOk

`func (o *PostSearchCommentSearchRequest) GetVchOk() (*[]string, bool)`

GetVchOk returns a tuple with the Vch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVch

`func (o *PostSearchCommentSearchRequest) SetVch(v []string)`

SetVch sets Vch field to given value.

### HasVch

`func (o *PostSearchCommentSearchRequest) HasVch() bool`

HasVch returns a boolean if a field has been set.

### GetOrg

`func (o *PostSearchCommentSearchRequest) GetOrg() []string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PostSearchCommentSearchRequest) GetOrgOk() (*[]string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PostSearchCommentSearchRequest) SetOrg(v []string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PostSearchCommentSearchRequest) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPaginated

`func (o *PostSearchCommentSearchRequest) GetPaginated() bool`

GetPaginated returns the Paginated field if non-nil, zero value otherwise.

### GetPaginatedOk

`func (o *PostSearchCommentSearchRequest) GetPaginatedOk() (*bool, bool)`

GetPaginatedOk returns a tuple with the Paginated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaginated

`func (o *PostSearchCommentSearchRequest) SetPaginated(v bool)`

SetPaginated sets Paginated field to given value.

### HasPaginated

`func (o *PostSearchCommentSearchRequest) HasPaginated() bool`

HasPaginated returns a boolean if a field has been set.

### GetOffset

`func (o *PostSearchCommentSearchRequest) GetOffset() float32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PostSearchCommentSearchRequest) GetOffsetOk() (*float32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PostSearchCommentSearchRequest) SetOffset(v float32)`

SetOffset sets Offset field to given value.


### GetLimit

`func (o *PostSearchCommentSearchRequest) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PostSearchCommentSearchRequest) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PostSearchCommentSearchRequest) SetLimit(v float32)`

SetLimit sets Limit field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


