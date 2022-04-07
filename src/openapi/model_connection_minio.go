/*
 * MassBank Minio API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type ConnectionMinio struct {
	Endpoint string `json:"endpoint,omitempty"`

	Username string `json:"username,omitempty"`

	Password string `json:"password,omitempty"`

	UseSsl bool `json:"useSsl,omitempty"`
}

// AssertConnectionMinioRequired checks if the required fields are not zero-ed
func AssertConnectionMinioRequired(obj ConnectionMinio) error {
	return nil
}

// AssertRecurseConnectionMinioRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of ConnectionMinio (e.g. [][]ConnectionMinio), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseConnectionMinioRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aConnectionMinio, ok := obj.(ConnectionMinio)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertConnectionMinioRequired(aConnectionMinio)
	})
}
