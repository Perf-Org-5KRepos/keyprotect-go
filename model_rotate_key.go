/*
 * IBM Key Protect API
 *
 * IBM Key Protect helps you provision encrypted keys for apps across IBM Cloud. As you manage the lifecycle of your keys, you can benefit from knowing that your keys are secured by cloud-based FIPS 140-2 Level 3 hardware security modules (HSMs) that protect against theft of information. You can use the Key Protect API to store, generate, and retrieve your key material. Keys within the service can protect any type of data in your symmetric key based encryption solution.
 *
 * API version: 2.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package keyprotect
// RotateKey Properties that are associated with wrap actions.
type RotateKey struct {
	// The key material that you want to import into the service for rotating an existing root key. This value is  required for a `rotate` action if you initially imported the key material when you created the key.     To rotate an imported root key, provide a base64 encoded payload in the request entity-body. To rotate a root key that was initially generated by Key Protect, omit the `payload` property and pass in an empty request entity-body.
	Payload string `json:"payload,omitempty"`
}