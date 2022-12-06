# VideoFull

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
**Clips** | Pointer to [**[]VideoWithChannel**](VideoWithChannel.md) | Included when &#39;includes&#39; contains &#39;clips&#39; | [optional] 
**Sources** | Pointer to [**[]VideoWithChannel**](VideoWithChannel.md) | Included when &#39;includes&#39; contains &#39;sources&#39; | [optional] 
**Refers** | Pointer to [**[]VideoWithChannel**](VideoWithChannel.md) | Included when &#39;includes&#39; contains &#39;refers&#39; | [optional] 
**Simulcasts** | Pointer to [**[]VideoWithChannel**](VideoWithChannel.md) | Included when &#39;includes&#39; contains &#39;simulcasts&#39; | [optional] 
**Mentions** | Pointer to [**[]ChannelMin**](ChannelMin.md) | VTubers mentioned by this video, Included when &#39;includes&#39; contains &#39;mentions&#39; | [optional] 
**Songs** | Pointer to **NullableFloat32** | Number of songs | [optional] 

## Methods

### NewVideoFull

`func NewVideoFull() *VideoFull`

NewVideoFull instantiates a new VideoFull object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoFullWithDefaults

`func NewVideoFullWithDefaults() *VideoFull`

NewVideoFullWithDefaults instantiates a new VideoFull object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VideoFull) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VideoFull) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VideoFull) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VideoFull) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *VideoFull) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *VideoFull) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *VideoFull) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *VideoFull) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *VideoFull) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *VideoFull) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *VideoFull) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *VideoFull) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTopicId

`func (o *VideoFull) GetTopicId() string`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *VideoFull) GetTopicIdOk() (*string, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *VideoFull) SetTopicId(v string)`

SetTopicId sets TopicId field to given value.

### HasTopicId

`func (o *VideoFull) HasTopicId() bool`

HasTopicId returns a boolean if a field has been set.

### SetTopicIdNil

`func (o *VideoFull) SetTopicIdNil(b bool)`

 SetTopicIdNil sets the value for TopicId to be an explicit nil

### UnsetTopicId
`func (o *VideoFull) UnsetTopicId()`

UnsetTopicId ensures that no value is present for TopicId, not even an explicit nil
### GetPublishedAt

`func (o *VideoFull) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *VideoFull) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *VideoFull) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *VideoFull) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### SetPublishedAtNil

`func (o *VideoFull) SetPublishedAtNil(b bool)`

 SetPublishedAtNil sets the value for PublishedAt to be an explicit nil

### UnsetPublishedAt
`func (o *VideoFull) UnsetPublishedAt()`

UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
### GetAvailableAt

`func (o *VideoFull) GetAvailableAt() time.Time`

GetAvailableAt returns the AvailableAt field if non-nil, zero value otherwise.

### GetAvailableAtOk

`func (o *VideoFull) GetAvailableAtOk() (*time.Time, bool)`

GetAvailableAtOk returns a tuple with the AvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAt

`func (o *VideoFull) SetAvailableAt(v time.Time)`

SetAvailableAt sets AvailableAt field to given value.

### HasAvailableAt

`func (o *VideoFull) HasAvailableAt() bool`

HasAvailableAt returns a boolean if a field has been set.

### GetDuration

`func (o *VideoFull) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VideoFull) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VideoFull) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VideoFull) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStatus

`func (o *VideoFull) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VideoFull) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VideoFull) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VideoFull) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartScheduled

`func (o *VideoFull) GetStartScheduled() time.Time`

GetStartScheduled returns the StartScheduled field if non-nil, zero value otherwise.

### GetStartScheduledOk

`func (o *VideoFull) GetStartScheduledOk() (*time.Time, bool)`

GetStartScheduledOk returns a tuple with the StartScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartScheduled

`func (o *VideoFull) SetStartScheduled(v time.Time)`

SetStartScheduled sets StartScheduled field to given value.

### HasStartScheduled

`func (o *VideoFull) HasStartScheduled() bool`

HasStartScheduled returns a boolean if a field has been set.

### SetStartScheduledNil

`func (o *VideoFull) SetStartScheduledNil(b bool)`

 SetStartScheduledNil sets the value for StartScheduled to be an explicit nil

### UnsetStartScheduled
`func (o *VideoFull) UnsetStartScheduled()`

UnsetStartScheduled ensures that no value is present for StartScheduled, not even an explicit nil
### GetStartActual

`func (o *VideoFull) GetStartActual() time.Time`

GetStartActual returns the StartActual field if non-nil, zero value otherwise.

### GetStartActualOk

`func (o *VideoFull) GetStartActualOk() (*time.Time, bool)`

GetStartActualOk returns a tuple with the StartActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActual

`func (o *VideoFull) SetStartActual(v time.Time)`

SetStartActual sets StartActual field to given value.

### HasStartActual

`func (o *VideoFull) HasStartActual() bool`

HasStartActual returns a boolean if a field has been set.

### SetStartActualNil

`func (o *VideoFull) SetStartActualNil(b bool)`

 SetStartActualNil sets the value for StartActual to be an explicit nil

### UnsetStartActual
`func (o *VideoFull) UnsetStartActual()`

UnsetStartActual ensures that no value is present for StartActual, not even an explicit nil
### GetEndActual

`func (o *VideoFull) GetEndActual() time.Time`

GetEndActual returns the EndActual field if non-nil, zero value otherwise.

### GetEndActualOk

`func (o *VideoFull) GetEndActualOk() (*time.Time, bool)`

GetEndActualOk returns a tuple with the EndActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndActual

`func (o *VideoFull) SetEndActual(v time.Time)`

SetEndActual sets EndActual field to given value.

### HasEndActual

`func (o *VideoFull) HasEndActual() bool`

HasEndActual returns a boolean if a field has been set.

### SetEndActualNil

`func (o *VideoFull) SetEndActualNil(b bool)`

 SetEndActualNil sets the value for EndActual to be an explicit nil

### UnsetEndActual
`func (o *VideoFull) UnsetEndActual()`

UnsetEndActual ensures that no value is present for EndActual, not even an explicit nil
### GetLiveViewers

`func (o *VideoFull) GetLiveViewers() int32`

GetLiveViewers returns the LiveViewers field if non-nil, zero value otherwise.

### GetLiveViewersOk

`func (o *VideoFull) GetLiveViewersOk() (*int32, bool)`

GetLiveViewersOk returns a tuple with the LiveViewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveViewers

`func (o *VideoFull) SetLiveViewers(v int32)`

SetLiveViewers sets LiveViewers field to given value.

### HasLiveViewers

`func (o *VideoFull) HasLiveViewers() bool`

HasLiveViewers returns a boolean if a field has been set.

### SetLiveViewersNil

`func (o *VideoFull) SetLiveViewersNil(b bool)`

 SetLiveViewersNil sets the value for LiveViewers to be an explicit nil

### UnsetLiveViewers
`func (o *VideoFull) UnsetLiveViewers()`

UnsetLiveViewers ensures that no value is present for LiveViewers, not even an explicit nil
### GetDescription

`func (o *VideoFull) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VideoFull) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VideoFull) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VideoFull) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSongcount

`func (o *VideoFull) GetSongcount() float32`

GetSongcount returns the Songcount field if non-nil, zero value otherwise.

### GetSongcountOk

`func (o *VideoFull) GetSongcountOk() (*float32, bool)`

GetSongcountOk returns a tuple with the Songcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongcount

`func (o *VideoFull) SetSongcount(v float32)`

SetSongcount sets Songcount field to given value.

### HasSongcount

`func (o *VideoFull) HasSongcount() bool`

HasSongcount returns a boolean if a field has been set.

### GetChannelId

`func (o *VideoFull) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *VideoFull) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *VideoFull) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *VideoFull) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetClips

`func (o *VideoFull) GetClips() []VideoWithChannel`

GetClips returns the Clips field if non-nil, zero value otherwise.

### GetClipsOk

`func (o *VideoFull) GetClipsOk() (*[]VideoWithChannel, bool)`

GetClipsOk returns a tuple with the Clips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClips

`func (o *VideoFull) SetClips(v []VideoWithChannel)`

SetClips sets Clips field to given value.

### HasClips

`func (o *VideoFull) HasClips() bool`

HasClips returns a boolean if a field has been set.

### SetClipsNil

`func (o *VideoFull) SetClipsNil(b bool)`

 SetClipsNil sets the value for Clips to be an explicit nil

### UnsetClips
`func (o *VideoFull) UnsetClips()`

UnsetClips ensures that no value is present for Clips, not even an explicit nil
### GetSources

`func (o *VideoFull) GetSources() []VideoWithChannel`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *VideoFull) GetSourcesOk() (*[]VideoWithChannel, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *VideoFull) SetSources(v []VideoWithChannel)`

SetSources sets Sources field to given value.

### HasSources

`func (o *VideoFull) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *VideoFull) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *VideoFull) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetRefers

`func (o *VideoFull) GetRefers() []VideoWithChannel`

GetRefers returns the Refers field if non-nil, zero value otherwise.

### GetRefersOk

`func (o *VideoFull) GetRefersOk() (*[]VideoWithChannel, bool)`

GetRefersOk returns a tuple with the Refers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefers

`func (o *VideoFull) SetRefers(v []VideoWithChannel)`

SetRefers sets Refers field to given value.

### HasRefers

`func (o *VideoFull) HasRefers() bool`

HasRefers returns a boolean if a field has been set.

### SetRefersNil

`func (o *VideoFull) SetRefersNil(b bool)`

 SetRefersNil sets the value for Refers to be an explicit nil

### UnsetRefers
`func (o *VideoFull) UnsetRefers()`

UnsetRefers ensures that no value is present for Refers, not even an explicit nil
### GetSimulcasts

`func (o *VideoFull) GetSimulcasts() []VideoWithChannel`

GetSimulcasts returns the Simulcasts field if non-nil, zero value otherwise.

### GetSimulcastsOk

`func (o *VideoFull) GetSimulcastsOk() (*[]VideoWithChannel, bool)`

GetSimulcastsOk returns a tuple with the Simulcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulcasts

`func (o *VideoFull) SetSimulcasts(v []VideoWithChannel)`

SetSimulcasts sets Simulcasts field to given value.

### HasSimulcasts

`func (o *VideoFull) HasSimulcasts() bool`

HasSimulcasts returns a boolean if a field has been set.

### SetSimulcastsNil

`func (o *VideoFull) SetSimulcastsNil(b bool)`

 SetSimulcastsNil sets the value for Simulcasts to be an explicit nil

### UnsetSimulcasts
`func (o *VideoFull) UnsetSimulcasts()`

UnsetSimulcasts ensures that no value is present for Simulcasts, not even an explicit nil
### GetMentions

`func (o *VideoFull) GetMentions() []ChannelMin`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *VideoFull) GetMentionsOk() (*[]ChannelMin, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *VideoFull) SetMentions(v []ChannelMin)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *VideoFull) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### SetMentionsNil

`func (o *VideoFull) SetMentionsNil(b bool)`

 SetMentionsNil sets the value for Mentions to be an explicit nil

### UnsetMentions
`func (o *VideoFull) UnsetMentions()`

UnsetMentions ensures that no value is present for Mentions, not even an explicit nil
### GetSongs

`func (o *VideoFull) GetSongs() float32`

GetSongs returns the Songs field if non-nil, zero value otherwise.

### GetSongsOk

`func (o *VideoFull) GetSongsOk() (*float32, bool)`

GetSongsOk returns a tuple with the Songs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongs

`func (o *VideoFull) SetSongs(v float32)`

SetSongs sets Songs field to given value.

### HasSongs

`func (o *VideoFull) HasSongs() bool`

HasSongs returns a boolean if a field has been set.

### SetSongsNil

`func (o *VideoFull) SetSongsNil(b bool)`

 SetSongsNil sets the value for Songs to be an explicit nil

### UnsetSongs
`func (o *VideoFull) UnsetSongs()`

UnsetSongs ensures that no value is present for Songs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


