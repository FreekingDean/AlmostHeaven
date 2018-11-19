package auth

func BasicAuth(username, password string) *AuthUser {
	if username == password {
		return &User{
			Username:     username,
			FirstName:    "User",
			LastName:     "Name",
			EmailAddress: "none@ya.biz",
		}
	} else {
		return nil
	}
}
