package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func GenerateUUID() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		return "", fmt.Errorf("error creating uuid: %w", err)
	}
	return uid.String(), nil
}
