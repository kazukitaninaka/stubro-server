/*
 * api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type InlineResponse200 struct {
	UserId int `json:"userId"`

	MentorId int `json:"mentorId"`

	DesirableDate string `json:"desirableDate"`

	Message string `json:"message"`

	Id int `json:"id,omitempty"`
}
