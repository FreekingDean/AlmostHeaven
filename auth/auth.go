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
			ID:           "TEST_ID",
			Username:     username,
			FirstName:    "User",
			LastName:     "Name",
			EmailAddress: "none@ya.biz",
		}
	} else {
		return nil
	}
}
