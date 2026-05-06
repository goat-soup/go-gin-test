package models

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
	ID       int    `json:"id" gorm:"primary_key"`
}

func CheckAuth(username, password string) bool {
	var auth Auth
	db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth)
	return auth.ID > 0
}
