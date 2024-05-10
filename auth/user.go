package auth

type User struct {
	ID           string `json:"id"`
	Username     string `json:"username"`
	SessionToken string `json:"session_token"`
	ExpiresAt    string `json:"expires_at"`
}
