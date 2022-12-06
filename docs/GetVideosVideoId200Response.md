# GetVideosVideoId200Response

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
**Comments** | Pointer to [**[]Comment**](Comment.md) | Comments are only returned if c &#x3D;&#x3D;&#x3D; &#39;1&#39; | [optional] 
**Recommendations** | Pointer to [**[]Video**](Video.md) |  | [optional] 

## Methods

### NewGetVideosVideoId200Response

`func NewGetVideosVideoId200Response() *GetVideosVideoId200Response`

NewGetVideosVideoId200Response instantiates a new GetVideosVideoId200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetVideosVideoId200ResponseWithDefaults

`func NewGetVideosVideoId200ResponseWithDefaults() *GetVideosVideoId200Response`

NewGetVideosVideoId200ResponseWithDefaults instantiates a new GetVideosVideoId200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetVideosVideoId200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetVideosVideoId200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetVideosVideoId200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetVideosVideoId200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *GetVideosVideoId200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetVideosVideoId200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetVideosVideoId200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetVideosVideoId200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetType

`func (o *GetVideosVideoId200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetVideosVideoId200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetVideosVideoId200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetVideosVideoId200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetTopicId

`func (o *GetVideosVideoId200Response) GetTopicId() string`

GetTopicId returns the TopicId field if non-nil, zero value otherwise.

### GetTopicIdOk

`func (o *GetVideosVideoId200Response) GetTopicIdOk() (*string, bool)`

GetTopicIdOk returns a tuple with the TopicId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicId

`func (o *GetVideosVideoId200Response) SetTopicId(v string)`

SetTopicId sets TopicId field to given value.

### HasTopicId

`func (o *GetVideosVideoId200Response) HasTopicId() bool`

HasTopicId returns a boolean if a field has been set.

### SetTopicIdNil

`func (o *GetVideosVideoId200Response) SetTopicIdNil(b bool)`

 SetTopicIdNil sets the value for TopicId to be an explicit nil

### UnsetTopicId
`func (o *GetVideosVideoId200Response) UnsetTopicId()`

UnsetTopicId ensures that no value is present for TopicId, not even an explicit nil
### GetPublishedAt

`func (o *GetVideosVideoId200Response) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *GetVideosVideoId200Response) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *GetVideosVideoId200Response) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *GetVideosVideoId200Response) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### SetPublishedAtNil

`func (o *GetVideosVideoId200Response) SetPublishedAtNil(b bool)`

 SetPublishedAtNil sets the value for PublishedAt to be an explicit nil

### UnsetPublishedAt
`func (o *GetVideosVideoId200Response) UnsetPublishedAt()`

UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
### GetAvailableAt

`func (o *GetVideosVideoId200Response) GetAvailableAt() time.Time`

GetAvailableAt returns the AvailableAt field if non-nil, zero value otherwise.

### GetAvailableAtOk

`func (o *GetVideosVideoId200Response) GetAvailableAtOk() (*time.Time, bool)`

GetAvailableAtOk returns a tuple with the AvailableAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableAt

`func (o *GetVideosVideoId200Response) SetAvailableAt(v time.Time)`

SetAvailableAt sets AvailableAt field to given value.

### HasAvailableAt

`func (o *GetVideosVideoId200Response) HasAvailableAt() bool`

HasAvailableAt returns a boolean if a field has been set.

### GetDuration

`func (o *GetVideosVideoId200Response) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *GetVideosVideoId200Response) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *GetVideosVideoId200Response) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *GetVideosVideoId200Response) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetStatus

`func (o *GetVideosVideoId200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetVideosVideoId200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetVideosVideoId200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetVideosVideoId200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStartScheduled

`func (o *GetVideosVideoId200Response) GetStartScheduled() time.Time`

GetStartScheduled returns the StartScheduled field if non-nil, zero value otherwise.

### GetStartScheduledOk

`func (o *GetVideosVideoId200Response) GetStartScheduledOk() (*time.Time, bool)`

GetStartScheduledOk returns a tuple with the StartScheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartScheduled

`func (o *GetVideosVideoId200Response) SetStartScheduled(v time.Time)`

SetStartScheduled sets StartScheduled field to given value.

### HasStartScheduled

`func (o *GetVideosVideoId200Response) HasStartScheduled() bool`

HasStartScheduled returns a boolean if a field has been set.

### SetStartScheduledNil

`func (o *GetVideosVideoId200Response) SetStartScheduledNil(b bool)`

 SetStartScheduledNil sets the value for StartScheduled to be an explicit nil

### UnsetStartScheduled
`func (o *GetVideosVideoId200Response) UnsetStartScheduled()`

UnsetStartScheduled ensures that no value is present for StartScheduled, not even an explicit nil
### GetStartActual

`func (o *GetVideosVideoId200Response) GetStartActual() time.Time`

GetStartActual returns the StartActual field if non-nil, zero value otherwise.

### GetStartActualOk

`func (o *GetVideosVideoId200Response) GetStartActualOk() (*time.Time, bool)`

GetStartActualOk returns a tuple with the StartActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartActual

`func (o *GetVideosVideoId200Response) SetStartActual(v time.Time)`

SetStartActual sets StartActual field to given value.

### HasStartActual

`func (o *GetVideosVideoId200Response) HasStartActual() bool`

HasStartActual returns a boolean if a field has been set.

### SetStartActualNil

`func (o *GetVideosVideoId200Response) SetStartActualNil(b bool)`

 SetStartActualNil sets the value for StartActual to be an explicit nil

### UnsetStartActual
`func (o *GetVideosVideoId200Response) UnsetStartActual()`

UnsetStartActual ensures that no value is present for StartActual, not even an explicit nil
### GetEndActual

`func (o *GetVideosVideoId200Response) GetEndActual() time.Time`

GetEndActual returns the EndActual field if non-nil, zero value otherwise.

### GetEndActualOk

`func (o *GetVideosVideoId200Response) GetEndActualOk() (*time.Time, bool)`

GetEndActualOk returns a tuple with the EndActual field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndActual

`func (o *GetVideosVideoId200Response) SetEndActual(v time.Time)`

SetEndActual sets EndActual field to given value.

### HasEndActual

`func (o *GetVideosVideoId200Response) HasEndActual() bool`

HasEndActual returns a boolean if a field has been set.

### SetEndActualNil

`func (o *GetVideosVideoId200Response) SetEndActualNil(b bool)`

 SetEndActualNil sets the value for EndActual to be an explicit nil

### UnsetEndActual
`func (o *GetVideosVideoId200Response) UnsetEndActual()`

UnsetEndActual ensures that no value is present for EndActual, not even an explicit nil
### GetLiveViewers

`func (o *GetVideosVideoId200Response) GetLiveViewers() int32`

GetLiveViewers returns the LiveViewers field if non-nil, zero value otherwise.

### GetLiveViewersOk

`func (o *GetVideosVideoId200Response) GetLiveViewersOk() (*int32, bool)`

GetLiveViewersOk returns a tuple with the LiveViewers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveViewers

`func (o *GetVideosVideoId200Response) SetLiveViewers(v int32)`

SetLiveViewers sets LiveViewers field to given value.

### HasLiveViewers

`func (o *GetVideosVideoId200Response) HasLiveViewers() bool`

HasLiveViewers returns a boolean if a field has been set.

### SetLiveViewersNil

`func (o *GetVideosVideoId200Response) SetLiveViewersNil(b bool)`

 SetLiveViewersNil sets the value for LiveViewers to be an explicit nil

### UnsetLiveViewers
`func (o *GetVideosVideoId200Response) UnsetLiveViewers()`

UnsetLiveViewers ensures that no value is present for LiveViewers, not even an explicit nil
### GetDescription

`func (o *GetVideosVideoId200Response) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetVideosVideoId200Response) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetVideosVideoId200Response) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetVideosVideoId200Response) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSongcount

`func (o *GetVideosVideoId200Response) GetSongcount() float32`

GetSongcount returns the Songcount field if non-nil, zero value otherwise.

### GetSongcountOk

`func (o *GetVideosVideoId200Response) GetSongcountOk() (*float32, bool)`

GetSongcountOk returns a tuple with the Songcount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongcount

`func (o *GetVideosVideoId200Response) SetSongcount(v float32)`

SetSongcount sets Songcount field to given value.

### HasSongcount

`func (o *GetVideosVideoId200Response) HasSongcount() bool`

HasSongcount returns a boolean if a field has been set.

### GetChannelId

`func (o *GetVideosVideoId200Response) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *GetVideosVideoId200Response) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *GetVideosVideoId200Response) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.

### HasChannelId

`func (o *GetVideosVideoId200Response) HasChannelId() bool`

HasChannelId returns a boolean if a field has been set.

### GetClips

`func (o *GetVideosVideoId200Response) GetClips() []VideoWithChannel`

GetClips returns the Clips field if non-nil, zero value otherwise.

### GetClipsOk

`func (o *GetVideosVideoId200Response) GetClipsOk() (*[]VideoWithChannel, bool)`

GetClipsOk returns a tuple with the Clips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClips

`func (o *GetVideosVideoId200Response) SetClips(v []VideoWithChannel)`

SetClips sets Clips field to given value.

### HasClips

`func (o *GetVideosVideoId200Response) HasClips() bool`

HasClips returns a boolean if a field has been set.

### SetClipsNil

`func (o *GetVideosVideoId200Response) SetClipsNil(b bool)`

 SetClipsNil sets the value for Clips to be an explicit nil

### UnsetClips
`func (o *GetVideosVideoId200Response) UnsetClips()`

UnsetClips ensures that no value is present for Clips, not even an explicit nil
### GetSources

`func (o *GetVideosVideoId200Response) GetSources() []VideoWithChannel`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *GetVideosVideoId200Response) GetSourcesOk() (*[]VideoWithChannel, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *GetVideosVideoId200Response) SetSources(v []VideoWithChannel)`

SetSources sets Sources field to given value.

### HasSources

`func (o *GetVideosVideoId200Response) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *GetVideosVideoId200Response) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *GetVideosVideoId200Response) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetRefers

`func (o *GetVideosVideoId200Response) GetRefers() []VideoWithChannel`

GetRefers returns the Refers field if non-nil, zero value otherwise.

### GetRefersOk

`func (o *GetVideosVideoId200Response) GetRefersOk() (*[]VideoWithChannel, bool)`

GetRefersOk returns a tuple with the Refers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefers

`func (o *GetVideosVideoId200Response) SetRefers(v []VideoWithChannel)`

SetRefers sets Refers field to given value.

### HasRefers

`func (o *GetVideosVideoId200Response) HasRefers() bool`

HasRefers returns a boolean if a field has been set.

### SetRefersNil

`func (o *GetVideosVideoId200Response) SetRefersNil(b bool)`

 SetRefersNil sets the value for Refers to be an explicit nil

### UnsetRefers
`func (o *GetVideosVideoId200Response) UnsetRefers()`

UnsetRefers ensures that no value is present for Refers, not even an explicit nil
### GetSimulcasts

`func (o *GetVideosVideoId200Response) GetSimulcasts() []VideoWithChannel`

GetSimulcasts returns the Simulcasts field if non-nil, zero value otherwise.

### GetSimulcastsOk

`func (o *GetVideosVideoId200Response) GetSimulcastsOk() (*[]VideoWithChannel, bool)`

GetSimulcastsOk returns a tuple with the Simulcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulcasts

`func (o *GetVideosVideoId200Response) SetSimulcasts(v []VideoWithChannel)`

SetSimulcasts sets Simulcasts field to given value.

### HasSimulcasts

`func (o *GetVideosVideoId200Response) HasSimulcasts() bool`

HasSimulcasts returns a boolean if a field has been set.

### SetSimulcastsNil

`func (o *GetVideosVideoId200Response) SetSimulcastsNil(b bool)`

 SetSimulcastsNil sets the value for Simulcasts to be an explicit nil

### UnsetSimulcasts
`func (o *GetVideosVideoId200Response) UnsetSimulcasts()`

UnsetSimulcasts ensures that no value is present for Simulcasts, not even an explicit nil
### GetMentions

`func (o *GetVideosVideoId200Response) GetMentions() []ChannelMin`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *GetVideosVideoId200Response) GetMentionsOk() (*[]ChannelMin, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *GetVideosVideoId200Response) SetMentions(v []ChannelMin)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *GetVideosVideoId200Response) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### SetMentionsNil

`func (o *GetVideosVideoId200Response) SetMentionsNil(b bool)`

 SetMentionsNil sets the value for Mentions to be an explicit nil

### UnsetMentions
`func (o *GetVideosVideoId200Response) UnsetMentions()`

UnsetMentions ensures that no value is present for Mentions, not even an explicit nil
### GetSongs

`func (o *GetVideosVideoId200Response) GetSongs() float32`

GetSongs returns the Songs field if non-nil, zero value otherwise.

### GetSongsOk

`func (o *GetVideosVideoId200Response) GetSongsOk() (*float32, bool)`

GetSongsOk returns a tuple with the Songs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongs

`func (o *GetVideosVideoId200Response) SetSongs(v float32)`

SetSongs sets Songs field to given value.

### HasSongs

`func (o *GetVideosVideoId200Response) HasSongs() bool`

HasSongs returns a boolean if a field has been set.

### SetSongsNil

`func (o *GetVideosVideoId200Response) SetSongsNil(b bool)`

 SetSongsNil sets the value for Songs to be an explicit nil

### UnsetSongs
`func (o *GetVideosVideoId200Response) UnsetSongs()`

UnsetSongs ensures that no value is present for Songs, not even an explicit nil
### GetComments

`func (o *GetVideosVideoId200Response) GetComments() []Comment`

GetComments returns the Comments field if non-nil, zero value otherwise.

### GetCommentsOk

`func (o *GetVideosVideoId200Response) GetCommentsOk() (*[]Comment, bool)`

GetCommentsOk returns a tuple with the Comments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComments

`func (o *GetVideosVideoId200Response) SetComments(v []Comment)`

SetComments sets Comments field to given value.

### HasComments

`func (o *GetVideosVideoId200Response) HasComments() bool`

HasComments returns a boolean if a field has been set.

### GetRecommendations

`func (o *GetVideosVideoId200Response) GetRecommendations() []Video`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *GetVideosVideoId200Response) GetRecommendationsOk() (*[]Video, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *GetVideosVideoId200Response) SetRecommendations(v []Video)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *GetVideosVideoId200Response) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


