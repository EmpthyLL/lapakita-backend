package usecase

import (
	"context"

	"github.com/google/uuid"
)

func (u *UserUsecase) DeleteDocument(ctx context.Context, userID uuid.UUID, documentID uuid.UUID) error {
	return u.Repo.DeleteIdentityProfile(ctx, userID, documentID)
}
