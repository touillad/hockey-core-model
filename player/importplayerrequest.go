package player

// ImportPlayerRequest r
type ImportPlayerRequest struct {
	Name     string       `json:"name"`
	Age      int          `json:"age"`
	Position PositionEnum `json:"position"`
}
