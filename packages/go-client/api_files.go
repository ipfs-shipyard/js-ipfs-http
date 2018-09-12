/*
 * IPFS API Documentation
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"fmt"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type FilesApiService service

/*
FilesApiService Remove an MFS path
Removes a directory or file from your MFS
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param path The MFS path you wish to remove
 * @param optional nil or *DeleteFilesPathOpts - Optional Parameters:
 * @param "Recursive" (optional.Bool) -  Remove directories recursively
*/

type DeleteFilesPathOpts struct {
    Recursive optional.Bool
}

func (a *FilesApiService) DeleteFilesPath(ctx context.Context, path string, localVarOptionals *DeleteFilesPathOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Delete")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/files/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Recursive.IsSet() {
		localVarQueryParams.Add("recursive", parameterToString(localVarOptionals.Recursive.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
FilesApiService Get a file or directory from your MFS
Returns a file or directory from your MFS&lt;br/&gt;&lt;br/&gt;Specify &#x60;accept: application/json&#x60; for file/directory metadata or &#x60;accept: application/octet-stream&#x60; to download file data.&lt;br/&gt;&lt;br/&gt;If the path resolves to a file and &#x60;accept: application/octet-stream&#x60; has been specified, you may pass the &#x60;offset&#x60; and &#x60;length&#x60; parameters.
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param path The MFS path you wish to retrieve
 * @param optional nil or *GetFilesPathOpts - Optional Parameters:
 * @param "Offset" (optional.Int32) -  Return file data starting at this offset
 * @param "Length" (optional.Int32) -  Return only this many bytes of file data
 * @param "CidBase" (optional.String) -  Which number base to use when returning a CID
 * @param "CidVersion" (optional.Int32) -  Which CID version to use
@return UnixfsEntry
*/

type GetFilesPathOpts struct {
    Offset optional.Int32
    Length optional.Int32
    CidBase optional.String
    CidVersion optional.Int32
}

func (a *FilesApiService) GetFilesPath(ctx context.Context, path string, localVarOptionals *GetFilesPathOpts) (UnixfsEntry, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue UnixfsEntry
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/files/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Length.IsSet() {
		localVarQueryParams.Add("length", parameterToString(localVarOptionals.Length.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CidBase.IsSet() {
		localVarQueryParams.Add("cidBase", parameterToString(localVarOptionals.CidBase.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CidVersion.IsSet() {
		localVarQueryParams.Add("cidVersion", parameterToString(localVarOptionals.CidVersion.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json", "application/octet-stream"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode < 300 {
		// If we succeed, return the data, otherwise pass on to decode error.
		err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
		if err == nil {
			return localVarReturnValue, localVarHttpResponse, err
		}
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v UnixfsEntry
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"));
				if err != nil {
					newErr.error = err.Error()
					return localVarReturnValue, localVarHttpResponse, newErr
				}
				newErr.model = v
				return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
FilesApiService Update an MFS path
Updates a file or directory at the passed MFS path
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param path The MFS path you wish to update
 * @param optional nil or *PatchFilesPathOpts - Optional Parameters:
 * @param "Offset" (optional.Int32) -  Byte offset to begin writing at
 * @param "Parents" (optional.Bool) -  Make parent directories as needed
 * @param "Truncate" (optional.Bool) -  Truncate the file to size zero before writing
 * @param "Length" (optional.Int32) -  Maximum number of bytes to write
 * @param "RawLeaves" (optional.Bool) -  Use raw blocks for newly created leaf nodes
 * @param "HashAlg" (optional.String) - 
 * @param "CidVersion" (optional.Int32) -  Which CID version to use
 * @param "Flush" (optional.Bool) -  Flush the changes to disk immediately
 * @param "UNKNOWNBASETYPE" (optional.Interface of UNKNOWN_BASE_TYPE) - 
*/

type PatchFilesPathOpts struct {
    Offset optional.Int32
    Parents optional.Bool
    Truncate optional.Bool
    Length optional.Int32
    RawLeaves optional.Bool
    HashAlg optional.String
    CidVersion optional.Int32
    Flush optional.Bool
    UNKNOWNBASETYPE optional.Interface
}

func (a *FilesApiService) PatchFilesPath(ctx context.Context, path string, localVarOptionals *PatchFilesPathOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Patch")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/files/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Parents.IsSet() {
		localVarQueryParams.Add("parents", parameterToString(localVarOptionals.Parents.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Truncate.IsSet() {
		localVarQueryParams.Add("truncate", parameterToString(localVarOptionals.Truncate.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Length.IsSet() {
		localVarQueryParams.Add("length", parameterToString(localVarOptionals.Length.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RawLeaves.IsSet() {
		localVarQueryParams.Add("rawLeaves", parameterToString(localVarOptionals.RawLeaves.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HashAlg.IsSet() {
		localVarQueryParams.Add("hashAlg", parameterToString(localVarOptionals.HashAlg.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CidVersion.IsSet() {
		localVarQueryParams.Add("cidVersion", parameterToString(localVarOptionals.CidVersion.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Flush.IsSet() {
		localVarQueryParams.Add("flush", parameterToString(localVarOptionals.Flush.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/octet-stream", "multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.UNKNOWNBASETYPE.IsSet() {
		localVarOptionalUNKNOWNBASETYPE, localVarOptionalUNKNOWNBASETYPEok := localVarOptionals.UNKNOWNBASETYPE.Value().(UNKNOWN_BASE_TYPE)
		if !localVarOptionalUNKNOWNBASETYPEok {
			return nil, reportError("uNKNOWNBASETYPE should be UNKNOWN_BASE_TYPE")
		}
		localVarPostBody = &localVarOptionalUNKNOWNBASETYPE
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
FilesApiService Create an MFS path
Create a new file or directory at the passed MFS path
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param path The MFS path you wish to create
 * @param optional nil or *PostFilesPathOpts - Optional Parameters:
 * @param "Offset" (optional.Int32) -  Byte offset to begin writing at
 * @param "Parents" (optional.Bool) -  Make parent directories as needed
 * @param "Length" (optional.Int32) -  Maximum number of bytes to write
 * @param "RawLeaves" (optional.Bool) -  Use raw blocks for newly created leaf nodes
 * @param "HashAlg" (optional.String) - 
 * @param "CidVersion" (optional.Int32) -  Which CID version to use
 * @param "Flush" (optional.Bool) -  Flush the changes to disk immediately
 * @param "UNKNOWNBASETYPE" (optional.Interface of UNKNOWN_BASE_TYPE) - 
*/

type PostFilesPathOpts struct {
    Offset optional.Int32
    Parents optional.Bool
    Length optional.Int32
    RawLeaves optional.Bool
    HashAlg optional.String
    CidVersion optional.Int32
    Flush optional.Bool
    UNKNOWNBASETYPE optional.Interface
}

func (a *FilesApiService) PostFilesPath(ctx context.Context, path string, localVarOptionals *PostFilesPathOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Post")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/files/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Offset.IsSet() {
		localVarQueryParams.Add("offset", parameterToString(localVarOptionals.Offset.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Parents.IsSet() {
		localVarQueryParams.Add("parents", parameterToString(localVarOptionals.Parents.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Length.IsSet() {
		localVarQueryParams.Add("length", parameterToString(localVarOptionals.Length.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RawLeaves.IsSet() {
		localVarQueryParams.Add("rawLeaves", parameterToString(localVarOptionals.RawLeaves.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HashAlg.IsSet() {
		localVarQueryParams.Add("hashAlg", parameterToString(localVarOptionals.HashAlg.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CidVersion.IsSet() {
		localVarQueryParams.Add("cidVersion", parameterToString(localVarOptionals.CidVersion.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Flush.IsSet() {
		localVarQueryParams.Add("flush", parameterToString(localVarOptionals.Flush.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/octet-stream", "multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.UNKNOWNBASETYPE.IsSet() {
		localVarOptionalUNKNOWNBASETYPE, localVarOptionalUNKNOWNBASETYPEok := localVarOptionals.UNKNOWNBASETYPE.Value().(UNKNOWN_BASE_TYPE)
		if !localVarOptionalUNKNOWNBASETYPEok {
			return nil, reportError("uNKNOWNBASETYPE should be UNKNOWN_BASE_TYPE")
		}
		localVarPostBody = &localVarOptionalUNKNOWNBASETYPE
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}

/*
FilesApiService Update an MFS path
Updates a file or directory at the passed MFS path
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param path The MFS path you wish to update
 * @param optional nil or *PutFilesPathOpts - Optional Parameters:
 * @param "Parents" (optional.Bool) -  Make parent directories as needed
 * @param "RawLeaves" (optional.Bool) -  Use raw blocks for newly created leaf nodes
 * @param "HashAlg" (optional.String) - 
 * @param "CidVersion" (optional.Int32) -  Which CID version to use
 * @param "Flush" (optional.Bool) -  Flush the changes to disk immediately
 * @param "UNKNOWNBASETYPE" (optional.Interface of UNKNOWN_BASE_TYPE) - 
*/

type PutFilesPathOpts struct {
    Parents optional.Bool
    RawLeaves optional.Bool
    HashAlg optional.String
    CidVersion optional.Int32
    Flush optional.Bool
    UNKNOWNBASETYPE optional.Interface
}

func (a *FilesApiService) PutFilesPath(ctx context.Context, path string, localVarOptionals *PutFilesPathOpts) (*http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Put")
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/files/{path}"
	localVarPath = strings.Replace(localVarPath, "{"+"path"+"}", fmt.Sprintf("%v", path), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if localVarOptionals != nil && localVarOptionals.Parents.IsSet() {
		localVarQueryParams.Add("parents", parameterToString(localVarOptionals.Parents.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.RawLeaves.IsSet() {
		localVarQueryParams.Add("rawLeaves", parameterToString(localVarOptionals.RawLeaves.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.HashAlg.IsSet() {
		localVarQueryParams.Add("hashAlg", parameterToString(localVarOptionals.HashAlg.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.CidVersion.IsSet() {
		localVarQueryParams.Add("cidVersion", parameterToString(localVarOptionals.CidVersion.Value(), ""))
	}
	if localVarOptionals != nil && localVarOptionals.Flush.IsSet() {
		localVarQueryParams.Add("flush", parameterToString(localVarOptionals.Flush.Value(), ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/octet-stream", "multipart/form-data"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.UNKNOWNBASETYPE.IsSet() {
		localVarOptionalUNKNOWNBASETYPE, localVarOptionalUNKNOWNBASETYPEok := localVarOptionals.UNKNOWNBASETYPE.Value().(UNKNOWN_BASE_TYPE)
		if !localVarOptionalUNKNOWNBASETYPEok {
			return nil, reportError("uNKNOWNBASETYPE should be UNKNOWN_BASE_TYPE")
		}
		localVarPostBody = &localVarOptionalUNKNOWNBASETYPE
	}

	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body: localVarBody,
			error: localVarHttpResponse.Status,
		}
		return localVarHttpResponse, newErr
	}

	return localVarHttpResponse, nil
}
