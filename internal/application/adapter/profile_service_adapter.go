package adapters

import (
	"github.com/PatipatCha/sgoc_account_service/internal/domain/repositories"
	"github.com/PatipatCha/sgoc_account_service/internal/domain/services"
)

type ProfileAdapter struct {
	profileRepo repositories.ProfileRepository
}

func NewProfileAdapter(profileRepo repositories.ProfileRepository) services.ProfileService {
	return &ProfileAdapter{}
}

func (s ProfileAdapter) GetProfileAll() (string, error) {
	return "GetProfile", nil
}
