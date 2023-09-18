package services

import (
	"github.com/PatipatCha/sgoc_account_service/internal/domain/repositories"
)

type ProfileService struct {
	profileRepo repositories.ProfileRepository
}

func NewProfileService(profileRepo repositories.ProfileRepository) usecase.ProfileService {
	return &ProfileService{}
}

func (s ProfileService) GetProfileAll() (string, error) {
	return "GetProfile", nil
}
