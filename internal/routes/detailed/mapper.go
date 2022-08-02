package detailed

import (
	"encoding/json"
	domain "go-analytics/internal/domain/detailed"
)

func mapRequestToModel(req *request) (*domain.Detailed, error) {
	json, err := json.Marshal(req.Additionals)

	if err != nil {
		return nil, err
	}

	return &domain.Detailed{
		UserID:      req.UserID,
		Platform:    req.Platform,
		Event:       req.Event,
		Category:    req.Category,
		Value:       req.Value,
		Additionals: string(json),
	}, nil
}
