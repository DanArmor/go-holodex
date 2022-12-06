# VideoWithChannel

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

## Methods

### NewVideoWithChannel

`func NewVideoWithChannel() *VideoWithChannel`

NewVideoWithChannel instantiates a new VideoWithChannel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoWithChannelWithDefaults

`func NewVideoWithChannelWithDefaults() *VideoWithChannel`

NewVideoWithChannelWithDefaults instantiates a new VideoWithChannel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoWithChannel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoWithChannel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoWithChannel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VideoWithChannel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *VideoWithChannel) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VideoWithChannel) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VideoWithChannel) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VideoWithChannel) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *VideoWithChannel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoWithChannel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoWithChannel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VideoWithChannel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTopicId

`func (o *VideoWithChannel) GetTopicId() string`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *VideoWithChannel) GetTopicIdOk() (*string, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *VideoWithChannel) SetTopicId(v string)`

SetTopicId sets TopicId field to given value.

### HasTopicId

`func (o *VideoWithChannel) HasTopicId() bool`

HasTopicId returns a boolean if a field has been set.

### SetTopicIdNil

`func (o *VideoWithChannel) SetTopicIdNil(b bool)`

 SetTopicIdNil sets the value for TopicId to be an explicit nil

### UnsetTopicId
`func (o *VideoWithChannel) UnsetTopicId()`

UnsetTopicId ensures that no value is present for TopicId, not even an explicit nil
### GetPublishedAt

`func (o *VideoWithChannel) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VideoWithChannel) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VideoWithChannel) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *VideoWithChannel) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### SetPublishedAtNil

`func (o *VideoWithChannel) SetPublishedAtNil(b bool)`

 SetPublishedAtNil sets the value for PublishedAt to be an explicit nil

### UnsetPublishedAt
`func (o *VideoWithChannel) UnsetPublishedAt()`

UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
### GetAvailableAt

`func (o *VideoWithChannel) GetAvailableAt() time.Time`

GetAvailableAt returns the AvailableAt field if non-nil, zero value otherwise.

### GetAvailableAtOk

`func (o *VideoWithChannel) GetAvailableAtOk() (*time.Time, bool)`

GetAvailableAtOk returns a tuple with the AvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAt

`func (o *VideoWithChannel) SetAvailableAt(v time.Time)`

SetAvailableAt sets AvailableAt field to given value.

### HasAvailableAt

`func (o *VideoWithChannel) HasAvailableAt() bool`

HasAvailableAt returns a boolean if a field has been set.

### GetDuration

`func (o *VideoWithChannel) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoWithChannel) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoWithChannel) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoWithChannel) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStatus

`func (o *VideoWithChannel) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoWithChannel) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoWithChannel) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoWithChannel) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartScheduled

`func (o *VideoWithChannel) GetStartScheduled() time.Time`

GetStartScheduled returns the StartScheduled field if non-nil, zero value otherwise.

### GetStartScheduledOk

`func (o *VideoWithChannel) GetStartScheduledOk() (*time.Time, bool)`

GetStartScheduledOk returns a tuple with the StartScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartScheduled

`func (o *VideoWithChannel) SetStartScheduled(v time.Time)`

SetStartScheduled sets StartScheduled field to given value.

### HasStartScheduled

`func (o *VideoWithChannel) HasStartScheduled() bool`

HasStartScheduled returns a boolean if a field has been set.

### SetStartScheduledNil

`func (o *VideoWithChannel) SetStartScheduledNil(b bool)`

 SetStartScheduledNil sets the value for StartScheduled to be an explicit nil

### UnsetStartScheduled
`func (o *VideoWithChannel) UnsetStartScheduled()`

UnsetStartScheduled ensures that no value is present for StartScheduled, not even an explicit nil
### GetStartActual

`func (o *VideoWithChannel) GetStartActual() time.Time`

GetStartActual returns the StartActual field if non-nil, zero value otherwise.

### GetStartActualOk

`func (o *VideoWithChannel) GetStartActualOk() (*time.Time, bool)`

GetStartActualOk returns a tuple with the StartActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActual

`func (o *VideoWithChannel) SetStartActual(v time.Time)`

SetStartActual sets StartActual field to given value.

### HasStartActual

`func (o *VideoWithChannel) HasStartActual() bool`

HasStartActual returns a boolean if a field has been set.

### SetStartActualNil

`func (o *VideoWithChannel) SetStartActualNil(b bool)`

 SetStartActualNil sets the value for StartActual to be an explicit nil

### UnsetStartActual
`func (o *VideoWithChannel) UnsetStartActual()`

UnsetStartActual ensures that no value is present for StartActual, not even an explicit nil
### GetEndActual

`func (o *VideoWithChannel) GetEndActual() time.Time`

GetEndActual returns the EndActual field if non-nil, zero value otherwise.

### GetEndActualOk

`func (o *VideoWithChannel) GetEndActualOk() (*time.Time, bool)`

GetEndActualOk returns a tuple with the EndActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndActual

`func (o *VideoWithChannel) SetEndActual(v time.Time)`

SetEndActual sets EndActual field to given value.

### HasEndActual

`func (o *VideoWithChannel) HasEndActual() bool`

HasEndActual returns a boolean if a field has been set.

### SetEndActualNil

`func (o *VideoWithChannel) SetEndActualNil(b bool)`

 SetEndActualNil sets the value for EndActual to be an explicit nil

### UnsetEndActual
`func (o *VideoWithChannel) UnsetEndActual()`

UnsetEndActual ensures that no value is present for EndActual, not even an explicit nil
### GetLiveViewers

`func (o *VideoWithChannel) GetLiveViewers() int32`

GetLiveViewers returns the LiveViewers field if non-nil, zero value otherwise.

### GetLiveViewersOk

`func (o *VideoWithChannel) GetLiveViewersOk() (*int32, bool)`

GetLiveViewersOk returns a tuple with the LiveViewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveViewers

`func (o *VideoWithChannel) SetLiveViewers(v int32)`

SetLiveViewers sets LiveViewers field to given value.

### HasLiveViewers

`func (o *VideoWithChannel) HasLiveViewers() bool`

HasLiveViewers returns a boolean if a field has been set.

### SetLiveViewersNil

`func (o *VideoWithChannel) SetLiveViewersNil(b bool)`

 SetLiveViewersNil sets the value for LiveViewers to be an explicit nil

### UnsetLiveViewers
`func (o *VideoWithChannel) UnsetLiveViewers()`

UnsetLiveViewers ensures that no value is present for LiveViewers, not even an explicit nil
### GetDescription

`func (o *VideoWithChannel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoWithChannel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoWithChannel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoWithChannel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSongcount

`func (o *VideoWithChannel) GetSongcount() float32`

GetSongcount returns the Songcount field if non-nil, zero value otherwise.

### GetSongcountOk

`func (o *VideoWithChannel) GetSongcountOk() (*float32, bool)`

GetSongcountOk returns a tuple with the Songcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongcount

`func (o *VideoWithChannel) SetSongcount(v float32)`

SetSongcount sets Songcount field to given value.

### HasSongcount

`func (o *VideoWithChannel) HasSongcount() bool`

HasSongcount returns a boolean if a field has been set.

### GetChannelId

`func (o *VideoWithChannel) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *VideoWithChannel) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *VideoWithChannel) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *VideoWithChannel) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetChannel

`func (o *VideoWithChannel) GetChannel() ChannelMin`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *VideoWithChannel) GetChannelOk() (*ChannelMin, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *VideoWithChannel) SetChannel(v ChannelMin)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *VideoWithChannel) HasChannel() bool`

HasChannel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


