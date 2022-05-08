package auths

type User struct {
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	FullName  string   `json:"fullName"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
}
