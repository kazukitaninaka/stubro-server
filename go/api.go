/*
 * openapi
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"net/http"
)

// ConsultationApiRouter defines the required methods for binding the api requests to a responses for the ConsultationApi
// The ConsultationApiRouter implementation should parse necessary information from the http request,
// pass the data to a ConsultationApiServicer to perform the required actions, then write the service results to the http response.
type ConsultationApiRouter interface {
	GetConsultations(http.ResponseWriter, *http.Request)
	PostConsultation(http.ResponseWriter, *http.Request)
}

// MentorApiRouter defines the required methods for binding the api requests to a responses for the MentorApi
// The MentorApiRouter implementation should parse necessary information from the http request,
// pass the data to a MentorApiServicer to perform the required actions, then write the service results to the http response.
type MentorApiRouter interface {
	GetMentorByMentorId(http.ResponseWriter, *http.Request)
	GetMentors(http.ResponseWriter, *http.Request)
	PatchMentor(http.ResponseWriter, *http.Request)
	PostMentor(http.ResponseWriter, *http.Request)
}

// TermApiRouter defines the required methods for binding the api requests to a responses for the TermApi
// The TermApiRouter implementation should parse necessary information from the http request,
// pass the data to a TermApiServicer to perform the required actions, then write the service results to the http response.
type TermApiRouter interface {
	GetTerms(http.ResponseWriter, *http.Request)
}

// TypeApiRouter defines the required methods for binding the api requests to a responses for the TypeApi
// The TypeApiRouter implementation should parse necessary information from the http request,
// pass the data to a TypeApiServicer to perform the required actions, then write the service results to the http response.
type TypeApiRouter interface {
	GetTypes(http.ResponseWriter, *http.Request)
}

// UserApiRouter defines the required methods for binding the api requests to a responses for the UserApi
// The UserApiRouter implementation should parse necessary information from the http request,
// pass the data to a UserApiServicer to perform the required actions, then write the service results to the http response.
type UserApiRouter interface {
	GetUserByUid(http.ResponseWriter, *http.Request)
	GetUserByUserId(http.ResponseWriter, *http.Request)
	GetUsers(http.ResponseWriter, *http.Request)
	PatchUsersUserId(http.ResponseWriter, *http.Request)
	PostUser(http.ResponseWriter, *http.Request)
	UpdateUserByUserId(http.ResponseWriter, *http.Request)
}

// ConsultationApiServicer defines the api actions for the ConsultationApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type ConsultationApiServicer interface {
	GetConsultations(int, int) (interface{}, error)
	PostConsultation(Consultation) (interface{}, error)
}

// MentorApiServicer defines the api actions for the MentorApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type MentorApiServicer interface {
	GetMentorByMentorId(string) (interface{}, error)
	GetMentors() (interface{}, error)
	PatchMentor(string, InlineObject2) (interface{}, error)
	PostMentor(Mentor) (interface{}, error)
}

// TermApiServicer defines the api actions for the TermApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TermApiServicer interface {
	GetTerms() (interface{}, error)
}

// TypeApiServicer defines the api actions for the TypeApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type TypeApiServicer interface {
	GetTypes() (interface{}, error)
}

// UserApiServicer defines the api actions for the UserApi service
// This interface intended to stay up to date with the openapi yaml used to generate it,
// while the service implementation can be ignored with the .openapi-generator-ignore file
// and updated with the logic required for the API.
type UserApiServicer interface {
	GetUserByUid(string) (interface{}, error)
	GetUserByUserId(int) (interface{}, error)
	GetUsers() (interface{}, error)
	PatchUser(int, UserRequest) (interface{}, error)
	PostUser(User) (interface{}, error)
}
