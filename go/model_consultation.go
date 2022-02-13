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

type Consultation struct {
	*gorm.Model

	UserID uint `json:"userId"`

	User User

	MentorID uint `json:"mentorId"`

	Mentor Mentor

	DesirableDate string `json:"desirableDate"`

	Message string `json:"message"`
}
