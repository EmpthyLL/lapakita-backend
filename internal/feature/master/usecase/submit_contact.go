package usecase

import (
	"context"
	"time"

	"lapakita-backend/internal/entity"
	"lapakita-backend/internal/feature/master/dto"
	"lapakita-backend/pkg/api"
)

func (u *MasterUsecase) SubmitContact(ctx context.Context, req dto.SubmitContactRequest) (*dto.ContactResponse, error) {
	inquiry := &entity.ContactInquiry{
		Name:        req.Name,
		Email:       req.Email,
		Whatsapp:    req.Whatsapp,
		Persona:     req.Persona,
		InquiryType: req.InquiryType,
		Message:     req.Message,
		Status:      "unread",
	}

	if err := u.Repo.CreateContactInquiry(ctx, inquiry); err != nil {
		return nil, err
	}

	go func(email, name, persona, inqType, msg string) {
		_ = u.Mailer.SendContactAutoReply(email, name, persona, inqType, msg)
	}(inquiry.Email, inquiry.Name, inquiry.Persona, inquiry.InquiryType, inquiry.Message)

	return &dto.ContactResponse{
		ID:        api.EncodeResponse(inquiry.ID),
		Name:      inquiry.Name,
		Email:     inquiry.Email,
		Status:    inquiry.Status,
		CreatedAt: inquiry.CreatedAt.Format(time.RFC3339),
	}, nil
}
