package wire

import (
	"wot-statistics-server/domain"
	"wot-statistics-server/internal/infrastructure/postgres"
	"wot-statistics-server/repository"
	"wot-statistics-server/usecase"
)

// ProvideModuleRepository provides ModuleRepository
func ProvideModuleRepository(db *postgres.DB) domain.ModuleRepository {
	return repository.NewModuleRepository(db)
}

// ProvideModuleUseCase provides ModuleUseCase
func ProvideModuleUseCase(moduleRepo domain.ModuleRepository, profileRepo domain.ProfileRepository) domain.ModuleUseCase {
	return usecase.NewModuleUseCase(moduleRepo, profileRepo)
}
