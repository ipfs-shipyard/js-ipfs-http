/*
 * IPFS API Documentation
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type DagLink struct {
	Name string `json:"name,omitempty"`
	Size int32 `json:"size,omitempty"`
	Cid string `json:"cid"`
}
