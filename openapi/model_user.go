/*
 * api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import "gorm.io/gorm"

// User -
type User struct {
	*gorm.Model
	ID       uint   `json:"id"`
	Uid      string `json:"uid"`
	Username string `json:"username"`

	Email string `json:"email"`
}
