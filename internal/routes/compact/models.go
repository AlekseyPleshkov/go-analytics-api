package compact

type request struct {
	UserID   string         `json:"user_id" binding:"required"`
	Platform string         `json:"platform" binding:"required"`
	Data     map[string]any `json:"data" binding:"required"`
}

type response struct {
}
