package usecase

import (
	"fmt"

	"wot-statistics-server/domain"
)

type profileUseCase struct {
	profileRepository domain.ProfileRepository
}

// NewProfileUseCase creates a new profile use case instance
func NewProfileUseCase(profileRepository domain.ProfileRepository) domain.ProfileUseCase {
	return &profileUseCase{
		profileRepository: profileRepository,
	}
}

// GetProfileByTankID retrieves a profile by tank_id
func (u *profileUseCase) GetProfileByTankID(tankID int) (*domain.Profile, error) {
	if tankID <= 0 {
		return nil, fmt.Errorf("invalid tank ID: %d", tankID)
	}
	return u.profileRepository.GetProfileByTankID(tankID)
}

// CreateProfile creates a new profile
func (u *profileUseCase) CreateProfile(profile *domain.Profile) error {
	// Validate that either TankID or ModuleID is set
	if (profile.TankID == nil || *profile.TankID <= 0) && (profile.ModuleID == nil || *profile.ModuleID <= 0) {
		return fmt.Errorf("invalid profile: either TankID or ModuleID must be set")
	}
	// Validate that both are not set
	if profile.TankID != nil && profile.ModuleID != nil {
		return fmt.Errorf("invalid profile: cannot set both TankID and ModuleID")
	}
	return u.profileRepository.CreateProfile(profile)
}

// UpdateProfile updates an existing profile
func (u *profileUseCase) UpdateProfile(profile *domain.Profile) error {
	// Validate that either TankID or ModuleID is set
	if (profile.TankID == nil || *profile.TankID <= 0) && (profile.ModuleID == nil || *profile.ModuleID <= 0) {
		return fmt.Errorf("invalid profile: either TankID or ModuleID must be set")
	}
	// Validate that both are not set
	if profile.TankID != nil && profile.ModuleID != nil {
		return fmt.Errorf("invalid profile: cannot set both TankID and ModuleID")
	}
	return u.profileRepository.UpdateProfile(profile)
}

// DeleteProfile deletes a profile by tank_id
func (u *profileUseCase) DeleteProfile(tankID int) error {
	if tankID <= 0 {
		return fmt.Errorf("invalid tank ID: %d", tankID)
	}
	return u.profileRepository.DeleteProfile(tankID)
}

// GetAllProfiles retrieves all profiles
func (u *profileUseCase) GetAllProfiles() ([]domain.Profile, error) {
	return u.profileRepository.GetAllProfiles()
}
