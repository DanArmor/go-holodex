# \DefaultApi

All URIs are relative to *https://holodex.net/api/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCachedLive**](DefaultApi.md#GetCachedLive) | **Get** /users/live | Quickly Access Live / Upcoming for a set of Channels
[**GetChannels**](DefaultApi.md#GetChannels) | **Get** /channels | List Channels
[**GetV2ChannelsChannelId**](DefaultApi.md#GetV2ChannelsChannelId) | **Get** /channels/{channelId} | Get Channel Information
[**GetV2ChannelsChannelIdClips**](DefaultApi.md#GetV2ChannelsChannelIdClips) | **Get** /channels/{channelId}/{type} | Query Videos Related to Channel
[**GetVideosVideoId**](DefaultApi.md#GetVideosVideoId) | **Get** /videos/{videoId} | Get a single Video&#39;s metadata
[**LiveGet**](DefaultApi.md#LiveGet) | **Get** /live | Query Live and Upcoming Videos
[**PostSearchCommentSearch**](DefaultApi.md#PostSearchCommentSearch) | **Post** /search/commentSearch | 
[**PostSearchVideoSearch**](DefaultApi.md#PostSearchVideoSearch) | **Post** /search/videoSearch | 
[**VideosGet**](DefaultApi.md#VideosGet) | **Get** /videos | Query Videos



## GetCachedLive

> []Video GetCachedLive(ctx).Channels(channels).Execute()

Quickly Access Live / Upcoming for a set of Channels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    channels := "channels_example" // string | comma separated Youtube Channel IDs (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetCachedLive(context.Background()).Channels(channels).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetCachedLive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCachedLive`: []Video
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetCachedLive`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCachedLiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channels** | **string** | comma separated Youtube Channel IDs | 

### Return type

[**[]Video**](Video.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChannels

> []Channel GetChannels(ctx).Type_(type_).Offset(offset).Limit(limit).Org(org).Lang(lang).Sort(sort).Order(order).Execute()

List Channels



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    type_ := "type__example" // string | Type of Channel, whether it's a vtuber or a subber. Leave unset to query all.  (optional)
    offset := int32(56) // int32 | Offset results (optional) (default to 0)
    limit := int32(56) // int32 | Results limit (optional) (default to 25)
    org := "Hololive" // string | If set, filter for Vtuber belonging to a specific org (optional)
    lang := "en,ja,zh" // string | Comma separated list of languages. Language is a property of Channel, so only Channels satisfying the language will be returned. Leave empty to search for Vtubers and/or all clippers. (optional)
    sort := "sort_example" // string | Column to sort on, leave default to use 'org' as sort. Any first level property of channel should work here. (optional)
    order := "order_example" // string | ASC or DESC order, default asc. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetChannels(context.Background()).Type_(type_).Offset(offset).Limit(limit).Org(org).Lang(lang).Sort(sort).Order(order).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetChannels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannels`: []Channel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetChannels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **type_** | **string** | Type of Channel, whether it&#39;s a vtuber or a subber. Leave unset to query all.  | 
 **offset** | **int32** | Offset results | [default to 0]
 **limit** | **int32** | Results limit | [default to 25]
 **org** | **string** | If set, filter for Vtuber belonging to a specific org | 
 **lang** | **string** | Comma separated list of languages. Language is a property of Channel, so only Channels satisfying the language will be returned. Leave empty to search for Vtubers and/or all clippers. | 
 **sort** | **string** | Column to sort on, leave default to use &#39;org&#39; as sort. Any first level property of channel should work here. | 
 **order** | **string** | ASC or DESC order, default asc. | 

### Return type

[**[]Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV2ChannelsChannelId

> Channel GetV2ChannelsChannelId(ctx, channelId).Execute()

Get Channel Information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    channelId := "channelId_example" // string | ID of the Youtube Channel that is being queried

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetV2ChannelsChannelId(context.Background(), channelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetV2ChannelsChannelId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetV2ChannelsChannelId`: Channel
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetV2ChannelsChannelId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | ID of the Youtube Channel that is being queried | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2ChannelsChannelIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Channel**](Channel.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetV2ChannelsChannelIdClips

> GetV2ChannelsChannelIdClips200Response GetV2ChannelsChannelIdClips(ctx, channelId, type_).Lang(lang).Include(include).Limit(limit).Offset(offset).Paginated(paginated).Execute()

Query Videos Related to Channel



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    channelId := "channelId_example" // string | ID of the Youtube Channel that is being queried
    type_ := "type__example" // string | The type of video resource to fetch. Clips finds clip videos of a `vtuber` channel, Video finds the `channelId` channel's uploads, and collabs finds videos uploaded by other channels that mention this `channelId`
    lang := "en,ja" // string | A comma separated list of language codes to filter channels/clips, official streams do not follow this parameter (optional) (default to "all")
    include := []string{"Include_example"} // []string | Comma separated string of extra info for video. Should be a string instead of an array. (optional)
    limit := int32(56) // int32 | Results limit (optional) (default to 25)
    offset := int32(56) // int32 | Offset results (optional) (default to 0)
    paginated := "paginated_example" // string | If paginated is set to any non-empty value, return an object with total, otherwise returns an array. (optional) (default to "<empty>")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetV2ChannelsChannelIdClips(context.Background(), channelId, type_).Lang(lang).Include(include).Limit(limit).Offset(offset).Paginated(paginated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetV2ChannelsChannelIdClips``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetV2ChannelsChannelIdClips`: GetV2ChannelsChannelIdClips200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetV2ChannelsChannelIdClips`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**channelId** | **string** | ID of the Youtube Channel that is being queried | 
**type_** | **string** | The type of video resource to fetch. Clips finds clip videos of a &#x60;vtuber&#x60; channel, Video finds the &#x60;channelId&#x60; channel&#39;s uploads, and collabs finds videos uploaded by other channels that mention this &#x60;channelId&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetV2ChannelsChannelIdClipsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **lang** | **string** | A comma separated list of language codes to filter channels/clips, official streams do not follow this parameter | [default to &quot;all&quot;]
 **include** | **[]string** | Comma separated string of extra info for video. Should be a string instead of an array. | 
 **limit** | **int32** | Results limit | [default to 25]
 **offset** | **int32** | Offset results | [default to 0]
 **paginated** | **string** | If paginated is set to any non-empty value, return an object with total, otherwise returns an array. | [default to &quot;&lt;empty&gt;&quot;]

### Return type

[**GetV2ChannelsChannelIdClips200Response**](GetV2ChannelsChannelIdClips200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVideosVideoId

> GetVideosVideoId200Response GetVideosVideoId(ctx, videoId).Lang(lang).C(c).Execute()

Get a single Video's metadata



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    videoId := "videoId_example" // string | ID of a Youtube Video
    lang := "en,ja" // string | A comma separated list of language codes to filter channels/clips, official streams do not follow this parameter (optional) (default to "all")
    c := "1" // string | if `1` then will reply with timestamp comments for this video (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.GetVideosVideoId(context.Background(), videoId).Lang(lang).C(c).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.GetVideosVideoId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVideosVideoId`: GetVideosVideoId200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.GetVideosVideoId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**videoId** | **string** | ID of a Youtube Video | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVideosVideoIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **lang** | **string** | A comma separated list of language codes to filter channels/clips, official streams do not follow this parameter | [default to &quot;all&quot;]
 **c** | **string** | if &#x60;1&#x60; then will reply with timestamp comments for this video | 

### Return type

[**GetVideosVideoId200Response**](GetVideosVideoId200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiveGet

> LiveGet200Response LiveGet(ctx).ChannelId(channelId).Status(status).Lang(lang).Type_(type_).Topic(topic).Include(include).Org(org).MentionedChannelId(mentionedChannelId).Sort(sort).Order(order).Limit(limit).Offset(offset).Paginated(paginated).MaxUpcomingHours(maxUpcomingHours).Id(id).Execute()

Query Live and Upcoming Videos



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    channelId := "UCl_gCybOJRIgOXw6Qb4qJzQ" // string | Filter by video uploader channel id (optional)
    status := "past" // string | Filter by video status (optional)
    lang := "en,ja" // string | A comma separated list of language codes to filter channels/clips, official streams do not follow this parameter (optional) (default to "all")
    type_ := "type__example" // string | Filter by type of video (optional)
    topic := "singing" // string | Filter by video topic id (optional)
    include := []string{"Include_example"} // []string | Comma separated string of extra info for video. Should be a string instead of an array. (optional)
    org := "Hololive" // string | Filter by clips that feature the org's talent or videos posted by the org's talent (optional)
    mentionedChannelId := "mentionedChannelId_example" // string | Filter by mentioned channel id, excludes itself. Generally used to find collabs/clips that include the requested channel (optional)
    sort := "start_scheduled" // string | Sort by any returned video field (optional) (default to "available_at")
    order := "order_example" // string | Order by ascending or descending (optional) (default to "desc")
    limit := int32(56) // int32 | Results limit (optional) (default to 25)
    offset := int32(56) // int32 | Offset results (optional) (default to 0)
    paginated := "paginated_example" // string | If paginated is set to any non-empty value, return an object with total, otherwise returns an array. (optional) (default to "<empty>")
    maxUpcomingHours := float32(24) // float32 | Number of maximum hours upcoming to get upcoming videos by (for rejecting waiting rooms that are two years out) (optional)
    id := "id_example" // string | A single Youtube Video ID. If Specified, only this video can be returned (may be filtered out by other conditions though) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.LiveGet(context.Background()).ChannelId(channelId).Status(status).Lang(lang).Type_(type_).Topic(topic).Include(include).Org(org).MentionedChannelId(mentionedChannelId).Sort(sort).Order(order).Limit(limit).Offset(offset).Paginated(paginated).MaxUpcomingHours(maxUpcomingHours).Id(id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.LiveGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LiveGet`: LiveGet200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.LiveGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLiveGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Filter by video uploader channel id | 
 **status** | **string** | Filter by video status | 
 **lang** | **string** | A comma separated list of language codes to filter channels/clips, official streams do not follow this parameter | [default to &quot;all&quot;]
 **type_** | **string** | Filter by type of video | 
 **topic** | **string** | Filter by video topic id | 
 **include** | **[]string** | Comma separated string of extra info for video. Should be a string instead of an array. | 
 **org** | **string** | Filter by clips that feature the org&#39;s talent or videos posted by the org&#39;s talent | 
 **mentionedChannelId** | **string** | Filter by mentioned channel id, excludes itself. Generally used to find collabs/clips that include the requested channel | 
 **sort** | **string** | Sort by any returned video field | [default to &quot;available_at&quot;]
 **order** | **string** | Order by ascending or descending | [default to &quot;desc&quot;]
 **limit** | **int32** | Results limit | [default to 25]
 **offset** | **int32** | Offset results | [default to 0]
 **paginated** | **string** | If paginated is set to any non-empty value, return an object with total, otherwise returns an array. | [default to &quot;&lt;empty&gt;&quot;]
 **maxUpcomingHours** | **float32** | Number of maximum hours upcoming to get upcoming videos by (for rejecting waiting rooms that are two years out) | 
 **id** | **string** | A single Youtube Video ID. If Specified, only this video can be returned (may be filtered out by other conditions though) | 

### Return type

[**LiveGet200Response**](LiveGet200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSearchCommentSearch

> PostSearchCommentSearch200Response PostSearchCommentSearch(ctx).PostSearchCommentSearchRequest(postSearchCommentSearchRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    postSearchCommentSearchRequest := *openapiclient.NewPostSearchCommentSearchRequest("Sort_example", "Lemon", float32(123), float32(123)) // PostSearchCommentSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostSearchCommentSearch(context.Background()).PostSearchCommentSearchRequest(postSearchCommentSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSearchCommentSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSearchCommentSearch`: PostSearchCommentSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSearchCommentSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSearchCommentSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSearchCommentSearchRequest** | [**PostSearchCommentSearchRequest**](PostSearchCommentSearchRequest.md) |  | 

### Return type

[**PostSearchCommentSearch200Response**](PostSearchCommentSearch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostSearchVideoSearch

> PostSearchVideoSearch200Response PostSearchVideoSearch(ctx).PostSearchVideoSearchRequest(postSearchVideoSearchRequest).Execute()





### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    postSearchVideoSearchRequest := *openapiclient.NewPostSearchVideoSearchRequest("Sort_example", float32(123), float32(123)) // PostSearchVideoSearchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.PostSearchVideoSearch(context.Background()).PostSearchVideoSearchRequest(postSearchVideoSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.PostSearchVideoSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostSearchVideoSearch`: PostSearchVideoSearch200Response
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.PostSearchVideoSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostSearchVideoSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postSearchVideoSearchRequest** | [**PostSearchVideoSearchRequest**](PostSearchVideoSearchRequest.md) |  | 

### Return type

[**PostSearchVideoSearch200Response**](PostSearchVideoSearch200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VideosGet

> VideoFull VideosGet(ctx).ChannelId(channelId).Status(status).Lang(lang).Type_(type_).Topic(topic).Include(include).Org(org).MentionedChannelId(mentionedChannelId).Sort(sort).Order(order).Limit(limit).Offset(offset).Paginated(paginated).MaxUpcomingHours(maxUpcomingHours).Id(id).From(from).To(to).Execute()

Query Videos



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    channelId := "UCl_gCybOJRIgOXw6Qb4qJzQ" // string | Filter by video uploader channel id (optional)
    status := "past" // string | Filter by video status (optional)
    lang := "en,ja" // string | A comma separated list of language codes to filter channels/clips, official streams do not follow this parameter (optional) (default to "all")
    type_ := "type__example" // string | Filter by type of video (optional)
    topic := "singing" // string | Filter by video topic id (optional)
    include := []string{"Include_example"} // []string | Comma separated string of extra info for video. Should be a string instead of an array. (optional)
    org := "Hololive" // string | Filter by clips that feature the org's talent or videos posted by the org's talent (optional)
    mentionedChannelId := "mentionedChannelId_example" // string | Filter by mentioned channel id, excludes itself. Generally used to find collabs/clips that include the requested channel (optional)
    sort := "start_scheduled" // string | Sort by any returned video field (optional) (default to "available_at")
    order := "order_example" // string | Order by ascending or descending (optional) (default to "desc")
    limit := int32(56) // int32 | Results limit (optional) (default to 25)
    offset := int32(56) // int32 | Offset results (optional) (default to 0)
    paginated := "paginated_example" // string | If paginated is set to any non-empty value, return an object with total, otherwise returns an array. (optional) (default to "<empty>")
    maxUpcomingHours := float32(24) // float32 | Number of maximum hours upcoming to get upcoming videos by (for rejecting waiting rooms that are two years out) (optional)
    id := "id_example" // string | A single Youtube Video ID. If Specified, only this video can be returned (may be filtered out by other conditions though) (optional)
    from := "from_example" // string | ISO8601 Date String for minimum `available_at`. (`available_at` is the most accurate timestamp of when a video has gone live or became viewable - it is the first non null value of the `start_actual`, `start_scheduled` or `published_at` fields) (optional)
    to := "to_example" // string | ISO8601 Date String for maximum `available_at` (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DefaultApi.VideosGet(context.Background()).ChannelId(channelId).Status(status).Lang(lang).Type_(type_).Topic(topic).Include(include).Org(org).MentionedChannelId(mentionedChannelId).Sort(sort).Order(order).Limit(limit).Offset(offset).Paginated(paginated).MaxUpcomingHours(maxUpcomingHours).Id(id).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.VideosGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VideosGet`: VideoFull
    fmt.Fprintf(os.Stdout, "Response from `DefaultApi.VideosGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVideosGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **channelId** | **string** | Filter by video uploader channel id | 
 **status** | **string** | Filter by video status | 
 **lang** | **string** | A comma separated list of language codes to filter channels/clips, official streams do not follow this parameter | [default to &quot;all&quot;]
 **type_** | **string** | Filter by type of video | 
 **topic** | **string** | Filter by video topic id | 
 **include** | **[]string** | Comma separated string of extra info for video. Should be a string instead of an array. | 
 **org** | **string** | Filter by clips that feature the org&#39;s talent or videos posted by the org&#39;s talent | 
 **mentionedChannelId** | **string** | Filter by mentioned channel id, excludes itself. Generally used to find collabs/clips that include the requested channel | 
 **sort** | **string** | Sort by any returned video field | [default to &quot;available_at&quot;]
 **order** | **string** | Order by ascending or descending | [default to &quot;desc&quot;]
 **limit** | **int32** | Results limit | [default to 25]
 **offset** | **int32** | Offset results | [default to 0]
 **paginated** | **string** | If paginated is set to any non-empty value, return an object with total, otherwise returns an array. | [default to &quot;&lt;empty&gt;&quot;]
 **maxUpcomingHours** | **float32** | Number of maximum hours upcoming to get upcoming videos by (for rejecting waiting rooms that are two years out) | 
 **id** | **string** | A single Youtube Video ID. If Specified, only this video can be returned (may be filtered out by other conditions though) | 
 **from** | **string** | ISO8601 Date String for minimum &#x60;available_at&#x60;. (&#x60;available_at&#x60; is the most accurate timestamp of when a video has gone live or became viewable - it is the first non null value of the &#x60;start_actual&#x60;, &#x60;start_scheduled&#x60; or &#x60;published_at&#x60; fields) | 
 **to** | **string** | ISO8601 Date String for maximum &#x60;available_at&#x60; | 

### Return type

[**VideoFull**](VideoFull.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

