/*
 * api
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"errors"
	"log"

	"gorm.io/gorm"
)

// MentorApiService is a service that implents the logic for the MentorApiServicer
// This service should implement the business logic for every endpoint for the MentorApi API.
// Include any external packages or services that will be required by this service.
type MentorApiService struct {
}

// for preloading only id and name (to exclude things like createdAt and updatedAt)
func getOnlyIdAndName(d *gorm.DB) *gorm.DB {
	return d.Select("id, name")
}

// NewMentorApiService creates a default api service
func NewMentorApiService() MentorApiServicer {
	return &MentorApiService{}
}

// GetMentorByMentorId - Get a mentor
func (s *MentorApiService) GetMentorByMentorId(mentorId string) (interface{}, error) {
	// TODO - update GetMentorByMentorId with the required logic for this service method.
	// Add api_mentor_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'GetMentorByMentorId' not implemented")
}

// GetMentors - Get all mentors
func (s *MentorApiService) GetMentors() (interface{}, error) {
	// TODO - update GetMentors with the required logic for this service method.
	// Add api_mentor_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	var mentors []Mentor
	if err := Db.Preload("Term", getOnlyIdAndName).Preload("Types", getOnlyIdAndName).Find(&mentors).Error; err != nil {
		log.Fatal(err)
		return nil, err
	}
	return mentors, nil
	// return nil, errors.New("service method 'GetMentors' not implemented")
}

// PatchMentorMentorId - Update mentor&#39;s information
func (s *MentorApiService) PatchMentorMentorId(mentorId string, inlineObject2 InlineObject2) (interface{}, error) {
	// TODO - update PatchMentorMentorId with the required logic for this service method.
	// Add api_mentor_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PatchMentorMentorId' not implemented")
}

// PostMentor - Add a mentor
func (s *MentorApiService) PostMentor(mentor Mentor) (interface{}, error) {
	// TODO - update PostMentor with the required logic for this service method.
	// Add api_mentor_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	return nil, errors.New("service method 'PostMentor' not implemented")
}
