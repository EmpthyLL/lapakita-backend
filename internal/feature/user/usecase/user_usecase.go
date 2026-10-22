package usecase

import (
	"strings"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/user/repository"
	"lapakita-backend/pkg/storage"
)

type UserUsecase struct {
	Repo    *repository.UserRepository
	Seaweed *storage.SeaweedFSService
}

func NewUserUsecase(repo *repository.UserRepository, seaweed *storage.SeaweedFSService) *UserUsecase {
	return &UserUsecase{Repo: repo, Seaweed: seaweed}
}

func TransferRolesToTarget(phoneNumbers entity.PhoneNumbers, targetIndex int, newRoles []string) entity.PhoneNumbers {
	if len(newRoles) == 0 {
		return phoneNumbers
	}
	targetRoleMap := make(map[string]bool)
	for _, role := range newRoles {
		clean := strings.TrimSpace(role)
		if clean != "" {
			targetRoleMap[clean] = true
		}
	}
	for i := range phoneNumbers {
		if targetIndex >= 0 && i == targetIndex {
			continue
		}
		updatedRoles := make([]string, 0, len(phoneNumbers[i].Roles))
		for _, existingRole := range phoneNumbers[i].Roles {
			if !targetRoleMap[existingRole] {
				updatedRoles = append(updatedRoles, existingRole)
			}
		}
		phoneNumbers[i].Roles = updatedRoles
	}
	return phoneNumbers
}
