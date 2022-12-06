# VideoFullAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Clips** | Pointer to [**[]VideoWithChannel**](VideoWithChannel.md) | Included when &#39;includes&#39; contains &#39;clips&#39; | [optional] 
**Sources** | Pointer to [**[]VideoWithChannel**](VideoWithChannel.md) | Included when &#39;includes&#39; contains &#39;sources&#39; | [optional] 
**Refers** | Pointer to [**[]VideoWithChannel**](VideoWithChannel.md) | Included when &#39;includes&#39; contains &#39;refers&#39; | [optional] 
**Simulcasts** | Pointer to [**[]VideoWithChannel**](VideoWithChannel.md) | Included when &#39;includes&#39; contains &#39;simulcasts&#39; | [optional] 
**Mentions** | Pointer to [**[]ChannelMin**](ChannelMin.md) | VTubers mentioned by this video, Included when &#39;includes&#39; contains &#39;mentions&#39; | [optional] 
**Songs** | Pointer to **NullableFloat32** | Number of songs | [optional] 

## Methods

### NewVideoFullAllOf

`func NewVideoFullAllOf() *VideoFullAllOf`

NewVideoFullAllOf instantiates a new VideoFullAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoFullAllOfWithDefaults

`func NewVideoFullAllOfWithDefaults() *VideoFullAllOf`

NewVideoFullAllOfWithDefaults instantiates a new VideoFullAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClips

`func (o *VideoFullAllOf) GetClips() []VideoWithChannel`

GetClips returns the Clips field if non-nil, zero value otherwise.

### GetClipsOk

`func (o *VideoFullAllOf) GetClipsOk() (*[]VideoWithChannel, bool)`

GetClipsOk returns a tuple with the Clips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClips

`func (o *VideoFullAllOf) SetClips(v []VideoWithChannel)`

SetClips sets Clips field to given value.

### HasClips

`func (o *VideoFullAllOf) HasClips() bool`

HasClips returns a boolean if a field has been set.

### SetClipsNil

`func (o *VideoFullAllOf) SetClipsNil(b bool)`

 SetClipsNil sets the value for Clips to be an explicit nil

### UnsetClips
`func (o *VideoFullAllOf) UnsetClips()`

UnsetClips ensures that no value is present for Clips, not even an explicit nil
### GetSources

`func (o *VideoFullAllOf) GetSources() []VideoWithChannel`

GetSources returns the Sources field if non-nil, zero value otherwise.

### GetSourcesOk

`func (o *VideoFullAllOf) GetSourcesOk() (*[]VideoWithChannel, bool)`

GetSourcesOk returns a tuple with the Sources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSources

`func (o *VideoFullAllOf) SetSources(v []VideoWithChannel)`

SetSources sets Sources field to given value.

### HasSources

`func (o *VideoFullAllOf) HasSources() bool`

HasSources returns a boolean if a field has been set.

### SetSourcesNil

`func (o *VideoFullAllOf) SetSourcesNil(b bool)`

 SetSourcesNil sets the value for Sources to be an explicit nil

### UnsetSources
`func (o *VideoFullAllOf) UnsetSources()`

UnsetSources ensures that no value is present for Sources, not even an explicit nil
### GetRefers

`func (o *VideoFullAllOf) GetRefers() []VideoWithChannel`

GetRefers returns the Refers field if non-nil, zero value otherwise.

### GetRefersOk

`func (o *VideoFullAllOf) GetRefersOk() (*[]VideoWithChannel, bool)`

GetRefersOk returns a tuple with the Refers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefers

`func (o *VideoFullAllOf) SetRefers(v []VideoWithChannel)`

SetRefers sets Refers field to given value.

### HasRefers

`func (o *VideoFullAllOf) HasRefers() bool`

HasRefers returns a boolean if a field has been set.

### SetRefersNil

`func (o *VideoFullAllOf) SetRefersNil(b bool)`

 SetRefersNil sets the value for Refers to be an explicit nil

### UnsetRefers
`func (o *VideoFullAllOf) UnsetRefers()`

UnsetRefers ensures that no value is present for Refers, not even an explicit nil
### GetSimulcasts

`func (o *VideoFullAllOf) GetSimulcasts() []VideoWithChannel`

GetSimulcasts returns the Simulcasts field if non-nil, zero value otherwise.

### GetSimulcastsOk

`func (o *VideoFullAllOf) GetSimulcastsOk() (*[]VideoWithChannel, bool)`

GetSimulcastsOk returns a tuple with the Simulcasts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulcasts

`func (o *VideoFullAllOf) SetSimulcasts(v []VideoWithChannel)`

SetSimulcasts sets Simulcasts field to given value.

### HasSimulcasts

`func (o *VideoFullAllOf) HasSimulcasts() bool`

HasSimulcasts returns a boolean if a field has been set.

### SetSimulcastsNil

`func (o *VideoFullAllOf) SetSimulcastsNil(b bool)`

 SetSimulcastsNil sets the value for Simulcasts to be an explicit nil

### UnsetSimulcasts
`func (o *VideoFullAllOf) UnsetSimulcasts()`

UnsetSimulcasts ensures that no value is present for Simulcasts, not even an explicit nil
### GetMentions

`func (o *VideoFullAllOf) GetMentions() []ChannelMin`

GetMentions returns the Mentions field if non-nil, zero value otherwise.

### GetMentionsOk

`func (o *VideoFullAllOf) GetMentionsOk() (*[]ChannelMin, bool)`

GetMentionsOk returns a tuple with the Mentions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMentions

`func (o *VideoFullAllOf) SetMentions(v []ChannelMin)`

SetMentions sets Mentions field to given value.

### HasMentions

`func (o *VideoFullAllOf) HasMentions() bool`

HasMentions returns a boolean if a field has been set.

### SetMentionsNil

`func (o *VideoFullAllOf) SetMentionsNil(b bool)`

 SetMentionsNil sets the value for Mentions to be an explicit nil

### UnsetMentions
`func (o *VideoFullAllOf) UnsetMentions()`

UnsetMentions ensures that no value is present for Mentions, not even an explicit nil
### GetSongs

`func (o *VideoFullAllOf) GetSongs() float32`

GetSongs returns the Songs field if non-nil, zero value otherwise.

### GetSongsOk

`func (o *VideoFullAllOf) GetSongsOk() (*float32, bool)`

GetSongsOk returns a tuple with the Songs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSongs

`func (o *VideoFullAllOf) SetSongs(v float32)`

SetSongs sets Songs field to given value.

### HasSongs

`func (o *VideoFullAllOf) HasSongs() bool`

HasSongs returns a boolean if a field has been set.

### SetSongsNil

`func (o *VideoFullAllOf) SetSongsNil(b bool)`

 SetSongsNil sets the value for Songs to be an explicit nil

### UnsetSongs
`func (o *VideoFullAllOf) UnsetSongs()`

UnsetSongs ensures that no value is present for Songs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


