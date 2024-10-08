/*
Holodex/HoloAPI V2

Holodex Public API. Successor to the HoloAPI v1

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package holodex

import (
	"encoding/json"
	"time"
)

// Video struct for Video
type Video struct {
	Id *string `json:"id,omitempty"`
	Title *string `json:"title,omitempty"`
	Type *string `json:"type,omitempty"`
	// corresponds to a Topic ID, Videos of type `clip` cannot not have topic. Streams may or may not have topic.
	TopicId NullableString `json:"topic_id,omitempty"`
	PublishedAt NullableTime `json:"published_at,omitempty"`
	// Takes on the first non-null value of end_actual, start_actual, start_scheduled, or published_at
	AvailableAt *time.Time `json:"available_at,omitempty"`
	// Duration of the video in seconds
	Duration *int32 `json:"duration,omitempty"`
	Status *string `json:"status,omitempty"`
	// Included when includes contains 'live_info'
	StartScheduled NullableTime `json:"start_scheduled,omitempty"`
	// Included when includes contains 'live_info'
	StartActual NullableTime `json:"start_actual,omitempty"`
	// Included when includes contains 'live_info'
	EndActual NullableTime `json:"end_actual,omitempty"`
	// Included when includes contains 'live_info'
	LiveViewers NullableInt32 `json:"live_viewers,omitempty"`
	// Included when includes contains 'description'
	Description *string `json:"description,omitempty"`
	// Number of tagged songs for this video
	Songcount *float32 `json:"songcount,omitempty"`
	ChannelId *string `json:"channel_id,omitempty"`
	Channel *Channel `json:"channel,omitempty"`
}

// NewVideo instantiates a new Video object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVideo() *Video {
	this := Video{}
	return &this
}

// NewVideoWithDefaults instantiates a new Video object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVideoWithDefaults() *Video {
	this := Video{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Video) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Video) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Video) SetId(v string) {
	o.Id = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *Video) GetTitle() string {
	if o == nil || o.Title == nil {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetTitleOk() (*string, bool) {
	if o == nil || o.Title == nil {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *Video) HasTitle() bool {
	if o != nil && o.Title != nil {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *Video) SetTitle(v string) {
	o.Title = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Video) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Video) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Video) SetType(v string) {
	o.Type = &v
}

// GetTopicId returns the TopicId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetTopicId() string {
	if o == nil || o.TopicId.Get() == nil {
		var ret string
		return ret
	}
	return *o.TopicId.Get()
}

// GetTopicIdOk returns a tuple with the TopicId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetTopicIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.TopicId.Get(), o.TopicId.IsSet()
}

// HasTopicId returns a boolean if a field has been set.
func (o *Video) HasTopicId() bool {
	if o != nil && o.TopicId.IsSet() {
		return true
	}

	return false
}

// SetTopicId gets a reference to the given NullableString and assigns it to the TopicId field.
func (o *Video) SetTopicId(v string) {
	o.TopicId.Set(&v)
}
// SetTopicIdNil sets the value for TopicId to be an explicit nil
func (o *Video) SetTopicIdNil() {
	o.TopicId.Set(nil)
}

// UnsetTopicId ensures that no value is present for TopicId, not even an explicit nil
func (o *Video) UnsetTopicId() {
	o.TopicId.Unset()
}

// GetPublishedAt returns the PublishedAt field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetPublishedAt() time.Time {
	if o == nil || o.PublishedAt.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.PublishedAt.Get()
}

// GetPublishedAtOk returns a tuple with the PublishedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetPublishedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.PublishedAt.Get(), o.PublishedAt.IsSet()
}

// HasPublishedAt returns a boolean if a field has been set.
func (o *Video) HasPublishedAt() bool {
	if o != nil && o.PublishedAt.IsSet() {
		return true
	}

	return false
}

// SetPublishedAt gets a reference to the given NullableTime and assigns it to the PublishedAt field.
func (o *Video) SetPublishedAt(v time.Time) {
	o.PublishedAt.Set(&v)
}
// SetPublishedAtNil sets the value for PublishedAt to be an explicit nil
func (o *Video) SetPublishedAtNil() {
	o.PublishedAt.Set(nil)
}

// UnsetPublishedAt ensures that no value is present for PublishedAt, not even an explicit nil
func (o *Video) UnsetPublishedAt() {
	o.PublishedAt.Unset()
}

// GetAvailableAt returns the AvailableAt field value if set, zero value otherwise.
func (o *Video) GetAvailableAt() time.Time {
	if o == nil || o.AvailableAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AvailableAt
}

// GetAvailableAtOk returns a tuple with the AvailableAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetAvailableAtOk() (*time.Time, bool) {
	if o == nil || o.AvailableAt == nil {
		return nil, false
	}
	return o.AvailableAt, true
}

// HasAvailableAt returns a boolean if a field has been set.
func (o *Video) HasAvailableAt() bool {
	if o != nil && o.AvailableAt != nil {
		return true
	}

	return false
}

// SetAvailableAt gets a reference to the given time.Time and assigns it to the AvailableAt field.
func (o *Video) SetAvailableAt(v time.Time) {
	o.AvailableAt = &v
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *Video) GetDuration() int32 {
	if o == nil || o.Duration == nil {
		var ret int32
		return ret
	}
	return *o.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetDurationOk() (*int32, bool) {
	if o == nil || o.Duration == nil {
		return nil, false
	}
	return o.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *Video) HasDuration() bool {
	if o != nil && o.Duration != nil {
		return true
	}

	return false
}

// SetDuration gets a reference to the given int32 and assigns it to the Duration field.
func (o *Video) SetDuration(v int32) {
	o.Duration = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Video) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Video) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Video) SetStatus(v string) {
	o.Status = &v
}

// GetStartScheduled returns the StartScheduled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetStartScheduled() time.Time {
	if o == nil || o.StartScheduled.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartScheduled.Get()
}

// GetStartScheduledOk returns a tuple with the StartScheduled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetStartScheduledOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartScheduled.Get(), o.StartScheduled.IsSet()
}

// HasStartScheduled returns a boolean if a field has been set.
func (o *Video) HasStartScheduled() bool {
	if o != nil && o.StartScheduled.IsSet() {
		return true
	}

	return false
}

// SetStartScheduled gets a reference to the given NullableTime and assigns it to the StartScheduled field.
func (o *Video) SetStartScheduled(v time.Time) {
	o.StartScheduled.Set(&v)
}
// SetStartScheduledNil sets the value for StartScheduled to be an explicit nil
func (o *Video) SetStartScheduledNil() {
	o.StartScheduled.Set(nil)
}

// UnsetStartScheduled ensures that no value is present for StartScheduled, not even an explicit nil
func (o *Video) UnsetStartScheduled() {
	o.StartScheduled.Unset()
}

// GetStartActual returns the StartActual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetStartActual() time.Time {
	if o == nil || o.StartActual.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.StartActual.Get()
}

// GetStartActualOk returns a tuple with the StartActual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetStartActualOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.StartActual.Get(), o.StartActual.IsSet()
}

// HasStartActual returns a boolean if a field has been set.
func (o *Video) HasStartActual() bool {
	if o != nil && o.StartActual.IsSet() {
		return true
	}

	return false
}

// SetStartActual gets a reference to the given NullableTime and assigns it to the StartActual field.
func (o *Video) SetStartActual(v time.Time) {
	o.StartActual.Set(&v)
}
// SetStartActualNil sets the value for StartActual to be an explicit nil
func (o *Video) SetStartActualNil() {
	o.StartActual.Set(nil)
}

// UnsetStartActual ensures that no value is present for StartActual, not even an explicit nil
func (o *Video) UnsetStartActual() {
	o.StartActual.Unset()
}

// GetEndActual returns the EndActual field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetEndActual() time.Time {
	if o == nil || o.EndActual.Get() == nil {
		var ret time.Time
		return ret
	}
	return *o.EndActual.Get()
}

// GetEndActualOk returns a tuple with the EndActual field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetEndActualOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.EndActual.Get(), o.EndActual.IsSet()
}

// HasEndActual returns a boolean if a field has been set.
func (o *Video) HasEndActual() bool {
	if o != nil && o.EndActual.IsSet() {
		return true
	}

	return false
}

// SetEndActual gets a reference to the given NullableTime and assigns it to the EndActual field.
func (o *Video) SetEndActual(v time.Time) {
	o.EndActual.Set(&v)
}
// SetEndActualNil sets the value for EndActual to be an explicit nil
func (o *Video) SetEndActualNil() {
	o.EndActual.Set(nil)
}

// UnsetEndActual ensures that no value is present for EndActual, not even an explicit nil
func (o *Video) UnsetEndActual() {
	o.EndActual.Unset()
}

// GetLiveViewers returns the LiveViewers field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Video) GetLiveViewers() int32 {
	if o == nil || o.LiveViewers.Get() == nil {
		var ret int32
		return ret
	}
	return *o.LiveViewers.Get()
}

// GetLiveViewersOk returns a tuple with the LiveViewers field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Video) GetLiveViewersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.LiveViewers.Get(), o.LiveViewers.IsSet()
}

// HasLiveViewers returns a boolean if a field has been set.
func (o *Video) HasLiveViewers() bool {
	if o != nil && o.LiveViewers.IsSet() {
		return true
	}

	return false
}

// SetLiveViewers gets a reference to the given NullableInt32 and assigns it to the LiveViewers field.
func (o *Video) SetLiveViewers(v int32) {
	o.LiveViewers.Set(&v)
}
// SetLiveViewersNil sets the value for LiveViewers to be an explicit nil
func (o *Video) SetLiveViewersNil() {
	o.LiveViewers.Set(nil)
}

// UnsetLiveViewers ensures that no value is present for LiveViewers, not even an explicit nil
func (o *Video) UnsetLiveViewers() {
	o.LiveViewers.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Video) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Video) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Video) SetDescription(v string) {
	o.Description = &v
}

// GetSongcount returns the Songcount field value if set, zero value otherwise.
func (o *Video) GetSongcount() float32 {
	if o == nil || o.Songcount == nil {
		var ret float32
		return ret
	}
	return *o.Songcount
}

// GetSongcountOk returns a tuple with the Songcount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetSongcountOk() (*float32, bool) {
	if o == nil || o.Songcount == nil {
		return nil, false
	}
	return o.Songcount, true
}

// HasSongcount returns a boolean if a field has been set.
func (o *Video) HasSongcount() bool {
	if o != nil && o.Songcount != nil {
		return true
	}

	return false
}

// SetSongcount gets a reference to the given float32 and assigns it to the Songcount field.
func (o *Video) SetSongcount(v float32) {
	o.Songcount = &v
}

// GetChannelId returns the ChannelId field value if set, zero value otherwise.
func (o *Video) GetChannelId() string {
	if o == nil || o.ChannelId == nil {
		var ret string
		return ret
	}
	return *o.ChannelId
}

// GetChannelIdOk returns a tuple with the ChannelId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Video) GetChannelIdOk() (*string, bool) {
	if o == nil || o.ChannelId == nil {
		return nil, false
	}
	return o.ChannelId, true
}

// HasChannelId returns a boolean if a field has been set.
func (o *Video) HasChannelId() bool {
	if o != nil && o.ChannelId != nil {
		return true
	}

	return false
}

// SetChannelId gets a reference to the given string and assigns it to the ChannelId field.
func (o *Video) SetChannelId(v string) {
	o.ChannelId = &v
}

func (o Video) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Title != nil {
		toSerialize["title"] = o.Title
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.TopicId.IsSet() {
		toSerialize["topic_id"] = o.TopicId.Get()
	}
	if o.PublishedAt.IsSet() {
		toSerialize["published_at"] = o.PublishedAt.Get()
	}
	if o.AvailableAt != nil {
		toSerialize["available_at"] = o.AvailableAt
	}
	if o.Duration != nil {
		toSerialize["duration"] = o.Duration
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.StartScheduled.IsSet() {
		toSerialize["start_scheduled"] = o.StartScheduled.Get()
	}
	if o.StartActual.IsSet() {
		toSerialize["start_actual"] = o.StartActual.Get()
	}
	if o.EndActual.IsSet() {
		toSerialize["end_actual"] = o.EndActual.Get()
	}
	if o.LiveViewers.IsSet() {
		toSerialize["live_viewers"] = o.LiveViewers.Get()
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Songcount != nil {
		toSerialize["songcount"] = o.Songcount
	}
	if o.ChannelId != nil {
		toSerialize["channel_id"] = o.ChannelId
	}
	return json.Marshal(toSerialize)
}

type NullableVideo struct {
	value *Video
	isSet bool
}

func (v NullableVideo) Get() *Video {
	return v.value
}

func (v *NullableVideo) Set(val *Video) {
	v.value = val
	v.isSet = true
}

func (v NullableVideo) IsSet() bool {
	return v.isSet
}

func (v *NullableVideo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVideo(val *Video) *NullableVideo {
	return &NullableVideo{value: val, isSet: true}
}

func (v NullableVideo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVideo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


