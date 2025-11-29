package auth

type User struct {
	Username string
	Password string
}

// in-memory storage
var users = []User{
	{Username: "admin", Password: "1234"},
	{Username: "moussa", Password: "232"},
}

func AuthCore(username, password string) bool {
	for _, u := range users {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}
