# \PubsubApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PubsubList**](PubsubApi.md#PubsubList) | **Get** /pubsub/topics | Returns the list of topics the peer is subscribed to
[**PubsubSubscribe**](PubsubApi.md#PubsubSubscribe) | **Post** /pubsub/topics | Returns the list of topics the peer is subscribed to


# **PubsubList**
> string PubsubList(ctx, )
Returns the list of topics the peer is subscribed to

### Required Parameters
This endpoint does not need any parameter.

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PubsubSubscribe**
> string PubsubSubscribe(ctx, optional)
Returns the list of topics the peer is subscribed to

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***PubsubSubscribeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a PubsubSubscribeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscribeToPubsubTopicPayload** | [**optional.Interface of SubscribeToPubsubTopicPayload**](SubscribeToPubsubTopicPayload.md)|  | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json, */*

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

