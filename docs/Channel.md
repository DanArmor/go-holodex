# Channel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**EnglishName** | Pointer to **NullableString** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Org** | Pointer to **NullableString** |  | [optional] 
**Suborg** | Pointer to **NullableString** |  | [optional] 
**Photo** | Pointer to **NullableString** |  | [optional] 
**Banner** | Pointer to **NullableString** |  | [optional] 
**Twitter** | Pointer to **NullableString** |  | [optional] 
**VideoCount** | Pointer to **NullableString** |  | [optional] 
**SubscriberCount** | Pointer to **NullableString** |  | [optional] 
**ViewCount** | Pointer to **NullableString** |  | [optional] 
**ClipCount** | Pointer to **NullableString** |  | [optional] 
**Lang** | Pointer to **NullableString** |  | [optional] 
**PublishedAt** | Pointer to **time.Time** |  | [optional] 
**Inactive** | Pointer to **bool** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewChannel

`func NewChannel() *Channel`

NewChannel instantiates a new Channel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChannelWithDefaults

`func NewChannelWithDefaults() *Channel`

NewChannelWithDefaults instantiates a new Channel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Channel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Channel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Channel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Channel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Channel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Channel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Channel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Channel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnglishName

`func (o *Channel) GetEnglishName() string`

GetEnglishName returns the EnglishName field if non-nil, zero value otherwise.

### GetEnglishNameOk

`func (o *Channel) GetEnglishNameOk() (*string, bool)`

GetEnglishNameOk returns a tuple with the EnglishName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnglishName

`func (o *Channel) SetEnglishName(v string)`

SetEnglishName sets EnglishName field to given value.

### HasEnglishName

`func (o *Channel) HasEnglishName() bool`

HasEnglishName returns a boolean if a field has been set.

### SetEnglishNameNil

`func (o *Channel) SetEnglishNameNil(b bool)`

 SetEnglishNameNil sets the value for EnglishName to be an explicit nil

### UnsetEnglishName
`func (o *Channel) UnsetEnglishName()`

UnsetEnglishName ensures that no value is present for EnglishName, not even an explicit nil
### GetType

`func (o *Channel) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Channel) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Channel) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Channel) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrg

`func (o *Channel) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *Channel) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *Channel) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *Channel) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### SetOrgNil

`func (o *Channel) SetOrgNil(b bool)`

 SetOrgNil sets the value for Org to be an explicit nil

### UnsetOrg
`func (o *Channel) UnsetOrg()`

UnsetOrg ensures that no value is present for Org, not even an explicit nil
### GetSuborg

`func (o *Channel) GetSuborg() string`

GetSuborg returns the Suborg field if non-nil, zero value otherwise.

### GetSuborgOk

`func (o *Channel) GetSuborgOk() (*string, bool)`

GetSuborgOk returns a tuple with the Suborg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuborg

`func (o *Channel) SetSuborg(v string)`

SetSuborg sets Suborg field to given value.

### HasSuborg

`func (o *Channel) HasSuborg() bool`

HasSuborg returns a boolean if a field has been set.

### SetSuborgNil

`func (o *Channel) SetSuborgNil(b bool)`

 SetSuborgNil sets the value for Suborg to be an explicit nil

### UnsetSuborg
`func (o *Channel) UnsetSuborg()`

UnsetSuborg ensures that no value is present for Suborg, not even an explicit nil
### GetPhoto

`func (o *Channel) GetPhoto() string`

GetPhoto returns the Photo field if non-nil, zero value otherwise.

### GetPhotoOk

`func (o *Channel) GetPhotoOk() (*string, bool)`

GetPhotoOk returns a tuple with the Photo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoto

`func (o *Channel) SetPhoto(v string)`

SetPhoto sets Photo field to given value.

### HasPhoto

`func (o *Channel) HasPhoto() bool`

HasPhoto returns a boolean if a field has been set.

### SetPhotoNil

`func (o *Channel) SetPhotoNil(b bool)`

 SetPhotoNil sets the value for Photo to be an explicit nil

### UnsetPhoto
`func (o *Channel) UnsetPhoto()`

UnsetPhoto ensures that no value is present for Photo, not even an explicit nil
### GetBanner

`func (o *Channel) GetBanner() string`

GetBanner returns the Banner field if non-nil, zero value otherwise.

### GetBannerOk

`func (o *Channel) GetBannerOk() (*string, bool)`

GetBannerOk returns a tuple with the Banner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBanner

`func (o *Channel) SetBanner(v string)`

SetBanner sets Banner field to given value.

### HasBanner

`func (o *Channel) HasBanner() bool`

HasBanner returns a boolean if a field has been set.

### SetBannerNil

`func (o *Channel) SetBannerNil(b bool)`

 SetBannerNil sets the value for Banner to be an explicit nil

### UnsetBanner
`func (o *Channel) UnsetBanner()`

UnsetBanner ensures that no value is present for Banner, not even an explicit nil
### GetTwitter

`func (o *Channel) GetTwitter() string`

GetTwitter returns the Twitter field if non-nil, zero value otherwise.

### GetTwitterOk

`func (o *Channel) GetTwitterOk() (*string, bool)`

GetTwitterOk returns a tuple with the Twitter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwitter

`func (o *Channel) SetTwitter(v string)`

SetTwitter sets Twitter field to given value.

### HasTwitter

`func (o *Channel) HasTwitter() bool`

HasTwitter returns a boolean if a field has been set.

### SetTwitterNil

`func (o *Channel) SetTwitterNil(b bool)`

 SetTwitterNil sets the value for Twitter to be an explicit nil

### UnsetTwitter
`func (o *Channel) UnsetTwitter()`

UnsetTwitter ensures that no value is present for Twitter, not even an explicit nil
### GetVideoCount

`func (o *Channel) GetVideoCount() string`

GetVideoCount returns the VideoCount field if non-nil, zero value otherwise.

### GetVideoCountOk

`func (o *Channel) GetVideoCountOk() (*string, bool)`

GetVideoCountOk returns a tuple with the VideoCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCount

`func (o *Channel) SetVideoCount(v string)`

SetVideoCount sets VideoCount field to given value.

### HasVideoCount

`func (o *Channel) HasVideoCount() bool`

HasVideoCount returns a boolean if a field has been set.

### SetVideoCountNil

`func (o *Channel) SetVideoCountNil(b bool)`

 SetVideoCountNil sets the value for VideoCount to be an explicit nil

### UnsetVideoCount
`func (o *Channel) UnsetVideoCount()`

UnsetVideoCount ensures that no value is present for VideoCount, not even an explicit nil
### GetSubscriberCount

`func (o *Channel) GetSubscriberCount() string`

GetSubscriberCount returns the SubscriberCount field if non-nil, zero value otherwise.

### GetSubscriberCountOk

`func (o *Channel) GetSubscriberCountOk() (*string, bool)`

GetSubscriberCountOk returns a tuple with the SubscriberCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberCount

`func (o *Channel) SetSubscriberCount(v string)`

SetSubscriberCount sets SubscriberCount field to given value.

### HasSubscriberCount

`func (o *Channel) HasSubscriberCount() bool`

HasSubscriberCount returns a boolean if a field has been set.

### SetSubscriberCountNil

`func (o *Channel) SetSubscriberCountNil(b bool)`

 SetSubscriberCountNil sets the value for SubscriberCount to be an explicit nil

### UnsetSubscriberCount
`func (o *Channel) UnsetSubscriberCount()`

UnsetSubscriberCount ensures that no value is present for SubscriberCount, not even an explicit nil
### GetViewCount

`func (o *Channel) GetViewCount() string`

GetViewCount returns the ViewCount field if non-nil, zero value otherwise.

### GetViewCountOk

`func (o *Channel) GetViewCountOk() (*string, bool)`

GetViewCountOk returns a tuple with the ViewCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewCount

`func (o *Channel) SetViewCount(v string)`

SetViewCount sets ViewCount field to given value.

### HasViewCount

`func (o *Channel) HasViewCount() bool`

HasViewCount returns a boolean if a field has been set.

### SetViewCountNil

`func (o *Channel) SetViewCountNil(b bool)`

 SetViewCountNil sets the value for ViewCount to be an explicit nil

### UnsetViewCount
`func (o *Channel) UnsetViewCount()`

UnsetViewCount ensures that no value is present for ViewCount, not even an explicit nil
### GetClipCount

`func (o *Channel) GetClipCount() string`

GetClipCount returns the ClipCount field if non-nil, zero value otherwise.

### GetClipCountOk

`func (o *Channel) GetClipCountOk() (*string, bool)`

GetClipCountOk returns a tuple with the ClipCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClipCount

`func (o *Channel) SetClipCount(v string)`

SetClipCount sets ClipCount field to given value.

### HasClipCount

`func (o *Channel) HasClipCount() bool`

HasClipCount returns a boolean if a field has been set.

### SetClipCountNil

`func (o *Channel) SetClipCountNil(b bool)`

 SetClipCountNil sets the value for ClipCount to be an explicit nil

### UnsetClipCount
`func (o *Channel) UnsetClipCount()`

UnsetClipCount ensures that no value is present for ClipCount, not even an explicit nil
### GetLang

`func (o *Channel) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Channel) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Channel) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *Channel) HasLang() bool`

HasLang returns a boolean if a field has been set.

### SetLangNil

`func (o *Channel) SetLangNil(b bool)`

 SetLangNil sets the value for Lang to be an explicit nil

### UnsetLang
`func (o *Channel) UnsetLang()`

UnsetLang ensures that no value is present for Lang, not even an explicit nil
### GetPublishedAt

`func (o *Channel) GetPublishedAt() time.Time`

GetPublishedAt returns the PublishedAt field if non-nil, zero value otherwise.

### GetPublishedAtOk

`func (o *Channel) GetPublishedAtOk() (*time.Time, bool)`

GetPublishedAtOk returns a tuple with the PublishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedAt

`func (o *Channel) SetPublishedAt(v time.Time)`

SetPublishedAt sets PublishedAt field to given value.

### HasPublishedAt

`func (o *Channel) HasPublishedAt() bool`

HasPublishedAt returns a boolean if a field has been set.

### GetInactive

`func (o *Channel) GetInactive() bool`

GetInactive returns the Inactive field if non-nil, zero value otherwise.

### GetInactiveOk

`func (o *Channel) GetInactiveOk() (*bool, bool)`

GetInactiveOk returns a tuple with the Inactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInactive

`func (o *Channel) SetInactive(v bool)`

SetInactive sets Inactive field to given value.

### HasInactive

`func (o *Channel) HasInactive() bool`

HasInactive returns a boolean if a field has been set.

### GetDescription

`func (o *Channel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Channel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Channel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Channel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


