/*
 * IPFS API Documentation
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UnixfsEntrySummary struct {

	Name string `json:"name"`

	Type string `json:"type"`

	Cid string `json:"cid"`

	Size int32 `json:"size,omitempty"`
}
