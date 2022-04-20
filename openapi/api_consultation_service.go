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

	"gorm.io/gorm"
)

// ConsultationApiService is a service that implents the logic for the ConsultationApiServicer
// This service should implement the business logic for every endpoint for the ConsultationApi API.
// Include any external packages or services that will be required by this service.
type ConsultationApiService struct {
}

// NewConsultationApiService creates a default api service
func NewConsultationApiService() ConsultationApiServicer {
	return &ConsultationApiService{}
}

// GetConsultations - Get all consultations
func (s *ConsultationApiService) GetConsultations(userId int, mentorId int) (interface{}, error) {
	// TODO - update GetConsultations with the required logic for this service method.
	// Add api_consultation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	var consultations []Consultation
	// Preload("Mentors").Preload("Users").Omit("created_at, updated_at, deleted_at")

	chain := Db.Where("")
	if userId != 0 {
		chain.Where("user_id = ?", userId)
	}
	if mentorId != 0 {
		chain.Where("mentor_id = ?", mentorId)
	}

	if err := chain.Preload("User", func(d *gorm.DB) *gorm.DB {
		return d.Select("id, username")
	}).Preload("Mentor", func(d *gorm.DB) *gorm.DB {
		return d.Select("id, username, price")
	}).Find(&consultations).Error; err != nil {
		return nil, errors.New("Not found")
	}
	return consultations, nil
}

// PostConsultation - Create a consultation
func (s *ConsultationApiService) PostConsultation(consultation Consultation) (interface{}, error) {
	// TODO - update PostConsultation with the required logic for this service method.
	// Add api_consultation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	if err := Db.Create(&consultation).Error; err != nil {
		return nil, errors.New("Create failed!")
	}
	return consultation, nil
}

func (s *ConsultationApiService) PatchConsultationByUid(uid string, consultationRequest ConsultationRequest) (interface{}, error) {
	// TODO - update PostConsultation with the required logic for this service method.
	// Add api_consultation_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.
	if err := Db.Model(&consultation).Where("uid = ?", uid).Update(map[string]ConsultationRequest{}).Error; err != nil {
		return nil, errors.New("Update failed!")
	}

	return consultation, nil
}
