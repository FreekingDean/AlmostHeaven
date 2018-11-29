package auth

type User struct {
	ID           string
	Username     string
	FirstName    string
	LastName     string
	EmailAddress string
}

func BasicAuth(username, password string) *User {
	if username == password {
		return &User{
			ID:           "2edd3c0e-db9d-4256-83bf-7fd064122948",
			Username:     username,
			FirstName:    "User",
			LastName:     "Name",
			EmailAddress: "none@ya.biz",
		}
	} else {
		return nil
	}
}
