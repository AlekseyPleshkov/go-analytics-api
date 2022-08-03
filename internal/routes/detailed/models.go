package detailed

type request struct {
	UserID      string         `json:"user_id" binding:"required,min=1,max=40"`
	Platform    string         `json:"platform" binding:"required,min=1,max=40"`
	Event       string         `json:"event" binding:"required,min=1,max=100"`
	Category    string         `json:"category" binding:"max=100"`
	Value       string         `json:"value" binding:"max=200"`
	Additionals map[string]any `json:"additionals"`
}

type response struct {
}
