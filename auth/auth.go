package auth

type User struct {
	UID             int64  `json:"uid"`
	Username        string `json:"username"`
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	ProfilePicture  string `json:"profile_picture"`
	PermissionLevel int    `json:"permission_level"`
}
