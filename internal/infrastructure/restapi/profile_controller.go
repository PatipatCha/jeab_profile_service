package restapi

import (
	"fmt"

	"github.com/PatipatCha/sgoc_account_service/internal/domain/services"
)

type ProfileController struct {
	services.ProfileService
}

func NewProfileController() *ProfileController {
	return &ProfileController{}
}

func (h *ProfileController) GetProfile() {
	fmt.Println(h.ProfileService.GetProfileAll())
}
