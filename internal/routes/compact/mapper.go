package compact

import (
	"encoding/json"
	domain "go-analytics/internal/domain/compact"
)

func mapRequestToModel(req *request) (*domain.Compact, error) {
	json, err := json.Marshal(req.Data)

	if err != nil {
		return nil, err
	}

	return &domain.Compact{
		UserID:   req.UserID,
		Platform: req.Platform,
		Data:     string(json),
	}, err
}
