package compact

type request struct {
	UserID   string         `json:"user_id" binding:"required,min=1,max=40"`
	Platform string         `json:"platform" binding:"required,min=1,max=40"`
	Data     map[string]any `json:"data" binding:"required,min=1"`
}

type response struct {
}
