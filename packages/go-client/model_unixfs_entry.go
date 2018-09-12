/*
 * IPFS API Documentation
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type UnixfsEntry struct {
	Cid string `json:"cid"`
	Type string `json:"type"`
	CumulativeSize int32 `json:"cumulativeSize"`
	Blocks int32 `json:"blocks"`
	Size int32 `json:"size,omitempty"`
	WithLocality bool `json:"withLocality,omitempty"`
	Local bool `json:"local,omitempty"`
	SizeLocal int32 `json:"sizeLocal,omitempty"`
	Contents Contents `json:"contents,omitempty"`
}
