# GetVideosVideoId200ResponseAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Comments** | Pointer to [**[]Comment**](Comment.md) | Comments are only returned if c &#x3D;&#x3D;&#x3D; &#39;1&#39; | [optional] 
**Recommendations** | Pointer to [**[]Video**](Video.md) |  | [optional] 

## Methods

### NewGetVideosVideoId200ResponseAllOf

`func NewGetVideosVideoId200ResponseAllOf() *GetVideosVideoId200ResponseAllOf`

NewGetVideosVideoId200ResponseAllOf instantiates a new GetVideosVideoId200ResponseAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideosVideoId200ResponseAllOfWithDefaults

`func NewGetVideosVideoId200ResponseAllOfWithDefaults() *GetVideosVideoId200ResponseAllOf`

NewGetVideosVideoId200ResponseAllOfWithDefaults instantiates a new GetVideosVideoId200ResponseAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComments

`func (o *GetVideosVideoId200ResponseAllOf) GetComments() []Comment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetVideosVideoId200ResponseAllOf) GetCommentsOk() (*[]Comment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetVideosVideoId200ResponseAllOf) SetComments(v []Comment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *GetVideosVideoId200ResponseAllOf) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetRecommendations

`func (o *GetVideosVideoId200ResponseAllOf) GetRecommendations() []Video`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *GetVideosVideoId200ResponseAllOf) GetRecommendationsOk() (*[]Video, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *GetVideosVideoId200ResponseAllOf) SetRecommendations(v []Video)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *GetVideosVideoId200ResponseAllOf) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


