package crypt

import (
	"strings"

	"github.com/google/uuid"
)

func RandomID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	correlationID := strings.ToUpper(strings.Replace(id.String(), "-", "", -1))[:24]
	//lg.Infof(correlationID, "Generated correlationId: %s", correlationID)
	return correlationID, nil
}
