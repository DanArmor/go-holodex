# PostSearchCommentSearch200ResponseOneOfInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**TopicId** | Pointer to **NullableString** | corresponds to a Topic ID, Videos of type &#x60;clip&#x60; cannot not have topic. Streams may or may not have topic. | [optional] 
**PublishedAt** | Pointer to **NullableTime** |  | [optional] 
**AvailableAt** | Pointer to **time.Time** | Takes on the first non-null value of end_actual, start_actual, start_scheduled, or published_at | [optional] 
**Duration** | Pointer to **int32** | Duration of the video in seconds | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StartScheduled** | Pointer to **NullableTime** | Included when includes contains &#39;live_info&#39; | [optional] 
**StartActual** | Pointer to **NullableTime** | Included when includes contains &#39;live_info&#39; | [optional] 
**EndActual** | Pointer to **NullableTime** | Included when includes contains &#39;live_info&#39; | [optional] 
**LiveViewers** | Pointer to **NullableInt32** | Included when includes contains &#39;live_info&#39; | [optional] 
**Description** | Pointer to **string** | Included when includes contains &#39;description&#39; | [optional] 
**Songcount** | Pointer to **float32** | Number of tagged songs for this video | [optional] 
**ChannelId** | Pointer to **string** |  | [optional] 
**Channel** | Pointer to [**ChannelMin**](ChannelMin.md) |  | [optional] 
**Comments** | Pointer to [**[]Comment**](Comment.md) |  | [optional] 

## Methods

### NewPostSearchCommentSearch200ResponseOneOfInner

`func NewPostSearchCommentSearch200ResponseOneOfInner() *PostSearchCommentSearch200ResponseOneOfInner`

NewPostSearchCommentSearch200ResponseOneOfInner instantiates a new PostSearchCommentSearch200ResponseOneOfInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostSearchCommentSearch200ResponseOneOfInnerWithDefaults

`func NewPostSearchCommentSearch200ResponseOneOfInnerWithDefaults() *PostSearchCommentSearch200ResponseOneOfInner`

NewPostSearchCommentSearch200ResponseOneOfInnerWithDefaults instantiates a new PostSearchCommentSearch200ResponseOneOfInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTopicId

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetTopicId() string`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetTopicIdOk() (*string, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetTopicId(v string)`

SetTopicId sets TopicId field to given value.

### HasTopicId

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasTopicId() bool`

HasTopicId returns a boolean if a field has been set.

### SetTopicIdNil

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetTopicIdNil(b bool)`

 SetTopicIdNil sets the value for TopicId to be an explicit nil

### UnsetTopicId
`func (o *PostSearchCommentSearch200ResponseOneOfInner) UnsetTopicId()`

UnsetTopicId ensures that no value is present for TopicId, not even an explicit nil
### GetPublishedAt

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### SetPublishedAtNil

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetPublishedAtNil(b bool)`

 SetPublishedAtNil sets the value for PublishedAt to be an explicit nil

### UnsetPublishedAt
`func (o *PostSearchCommentSearch200ResponseOneOfInner) UnsetPublishedAt()`

UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
### GetAvailableAt

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetAvailableAt() time.Time`

GetAvailableAt returns the AvailableAt field if non-nil, zero value otherwise.

### GetAvailableAtOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetAvailableAtOk() (*time.Time, bool)`

GetAvailableAtOk returns a tuple with the AvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAt

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetAvailableAt(v time.Time)`

SetAvailableAt sets AvailableAt field to given value.

### HasAvailableAt

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasAvailableAt() bool`

HasAvailableAt returns a boolean if a field has been set.

### GetDuration

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStatus

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartScheduled

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetStartScheduled() time.Time`

GetStartScheduled returns the StartScheduled field if non-nil, zero value otherwise.

### GetStartScheduledOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetStartScheduledOk() (*time.Time, bool)`

GetStartScheduledOk returns a tuple with the StartScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartScheduled

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetStartScheduled(v time.Time)`

SetStartScheduled sets StartScheduled field to given value.

### HasStartScheduled

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasStartScheduled() bool`

HasStartScheduled returns a boolean if a field has been set.

### SetStartScheduledNil

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetStartScheduledNil(b bool)`

 SetStartScheduledNil sets the value for StartScheduled to be an explicit nil

### UnsetStartScheduled
`func (o *PostSearchCommentSearch200ResponseOneOfInner) UnsetStartScheduled()`

UnsetStartScheduled ensures that no value is present for StartScheduled, not even an explicit nil
### GetStartActual

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetStartActual() time.Time`

GetStartActual returns the StartActual field if non-nil, zero value otherwise.

### GetStartActualOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetStartActualOk() (*time.Time, bool)`

GetStartActualOk returns a tuple with the StartActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActual

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetStartActual(v time.Time)`

SetStartActual sets StartActual field to given value.

### HasStartActual

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasStartActual() bool`

HasStartActual returns a boolean if a field has been set.

### SetStartActualNil

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetStartActualNil(b bool)`

 SetStartActualNil sets the value for StartActual to be an explicit nil

### UnsetStartActual
`func (o *PostSearchCommentSearch200ResponseOneOfInner) UnsetStartActual()`

UnsetStartActual ensures that no value is present for StartActual, not even an explicit nil
### GetEndActual

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetEndActual() time.Time`

GetEndActual returns the EndActual field if non-nil, zero value otherwise.

### GetEndActualOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetEndActualOk() (*time.Time, bool)`

GetEndActualOk returns a tuple with the EndActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndActual

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetEndActual(v time.Time)`

SetEndActual sets EndActual field to given value.

### HasEndActual

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasEndActual() bool`

HasEndActual returns a boolean if a field has been set.

### SetEndActualNil

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetEndActualNil(b bool)`

 SetEndActualNil sets the value for EndActual to be an explicit nil

### UnsetEndActual
`func (o *PostSearchCommentSearch200ResponseOneOfInner) UnsetEndActual()`

UnsetEndActual ensures that no value is present for EndActual, not even an explicit nil
### GetLiveViewers

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetLiveViewers() int32`

GetLiveViewers returns the LiveViewers field if non-nil, zero value otherwise.

### GetLiveViewersOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetLiveViewersOk() (*int32, bool)`

GetLiveViewersOk returns a tuple with the LiveViewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveViewers

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetLiveViewers(v int32)`

SetLiveViewers sets LiveViewers field to given value.

### HasLiveViewers

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasLiveViewers() bool`

HasLiveViewers returns a boolean if a field has been set.

### SetLiveViewersNil

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetLiveViewersNil(b bool)`

 SetLiveViewersNil sets the value for LiveViewers to be an explicit nil

### UnsetLiveViewers
`func (o *PostSearchCommentSearch200ResponseOneOfInner) UnsetLiveViewers()`

UnsetLiveViewers ensures that no value is present for LiveViewers, not even an explicit nil
### GetDescription

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSongcount

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetSongcount() float32`

GetSongcount returns the Songcount field if non-nil, zero value otherwise.

### GetSongcountOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetSongcountOk() (*float32, bool)`

GetSongcountOk returns a tuple with the Songcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongcount

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetSongcount(v float32)`

SetSongcount sets Songcount field to given value.

### HasSongcount

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasSongcount() bool`

HasSongcount returns a boolean if a field has been set.

### GetChannelId

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannel

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetChannel() ChannelMin`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetChannelOk() (*ChannelMin, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetChannel(v ChannelMin)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetComments

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetComments() []Comment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *PostSearchCommentSearch200ResponseOneOfInner) GetCommentsOk() (*[]Comment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *PostSearchCommentSearch200ResponseOneOfInner) SetComments(v []Comment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *PostSearchCommentSearch200ResponseOneOfInner) HasComments() bool`

HasComments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


