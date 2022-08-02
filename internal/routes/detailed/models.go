package detailed

type request struct {
	UserID      string         `json:"user_id" binding:"required"`
	Platform    string         `json:"platform" binding:"required"`
	Event       string         `json:"event" binding:"required"`
	Category    string         `json:"category"`
	Value       string         `json:"value"`
	Additionals map[string]any `json:"additionals"`
}

type response struct {
}
