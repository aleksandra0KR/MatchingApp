package kafka

type Message struct {
	Username   string `json:"user_name"`
	UserId     string `json:"user_id"`
	PlaylistID string `json:"playlist_id"`
}
