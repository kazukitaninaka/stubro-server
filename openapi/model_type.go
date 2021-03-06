/*
 * openapi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import "gorm.io/gorm"

type Type struct {
	*gorm.Model
	ID   uint   `json:"id" gorm:"primarykey"`
	Name string `json:"name" gorm:"unique"`
}

// AssertTypeRequired checks if the required fields are not zero-ed
func AssertTypeRequired(obj Type) error {
	elements := map[string]interface{}{
		"name": obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRecurseTypeRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of Type (e.g. [][]Type), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseTypeRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aType, ok := obj.(Type)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertTypeRequired(aType)
	})
}
