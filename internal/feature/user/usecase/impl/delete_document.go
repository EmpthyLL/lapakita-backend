package impl

import (
	"context"

	"lapakita-backend/internal/feature/user/usecase"

	"github.com/google/uuid"
)

func DeleteDocument(ctx context.Context, u *usecase.UserUsecase, userID uuid.UUID, documentID uuid.UUID) error {
	return u.Repo.DeleteIdentityProfile(ctx, userID, documentID)
}
