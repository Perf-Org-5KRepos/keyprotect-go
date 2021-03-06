/*
 * IBM Key Protect API
 *
 * IBM Key Protect helps you provision encrypted keys for apps across IBM Cloud. As you manage the lifecycle of your keys, you can benefit from knowing that your keys are secured by cloud-based FIPS 140-2 Level 3 hardware security modules (HSMs) that protect against theft of information. You can use the Key Protect API to store, generate, and retrieve your key material. Keys within the service can protect any type of data in your symmetric key based encryption solution.
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package keyprotect
import (
	"time"
)
// KeyMetadata Additional properties that describe a key.
type KeyMetadata struct {
	AlgorithmMetadata KeyMetadataAlgorithmMetadata `json:"algorithmMetadata,omitempty"`
	// The algorithm type used to generate the key. Currently, AES is supported.
	AlgorithmType string `json:"algorithmType,omitempty"`
	// The unique identifier for the resource that created the key.
	CreatedBy string `json:"createdBy,omitempty"`
	// The date the key material was created. The date format follows RFC 3339.
	CreationDate time.Time `json:"creationDate,omitempty"`
	// Updates to show when the key was last rotated. The date format follows RFC 3339.
	LastRotateDate time.Time `json:"lastRotateDate,omitempty"`
	// Updates when any part of the key metadata is modified. The date format follows RFC 3339.
	LastUpdateDate time.Time `json:"lastUpdateDate,omitempty"`
	// A code indicating the reason the key is not in the activation state.
	NonactiveStateReason int32 `json:"nonactiveStateReason,omitempty"`
}
