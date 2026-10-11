package api

import (
	"encoding/base64"
	"errors"

	"github.com/google/uuid"
)

var ErrInvalidID = errors.New("invalid UUID identifier")

// EncodeResponse converts a UUID to its 22-character URL-safe representation.
func EncodeResponse(value uuid.UUID) string {
	return base64.RawURLEncoding.EncodeToString(value[:])
}

// DecodeRequest accepts compact IDs and legacy canonical UUID strings.
func DecodeRequest(value string) (uuid.UUID, error) {
	if parsed, err := uuid.Parse(value); err == nil {
		return parsed, nil
	}

	decoded, err := base64.RawURLEncoding.DecodeString(value)
	if err != nil || len(decoded) != 16 {
		return uuid.Nil, ErrInvalidID
	}

	parsed, err := uuid.FromBytes(decoded)
	if err != nil {
		return uuid.Nil, ErrInvalidID
	}
	return parsed, nil
}

// Normalize returns a canonical UUID string for compact or legacy IDs.
func Normalize(value string) (string, error) {
	parsed, err := DecodeRequest(value)
	if err != nil {
		return "", err
	}
	return parsed.String(), nil
}
