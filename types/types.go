package types

type Permissions struct {
	USER     bool `json:"user"`
	LEARNING bool `json:"learning"`
	ORGS     bool `json:"orgs"`
}

type UserToken struct {
	ID        string      `json:"id"`
	UserType  string      `json:"userType"`
	AccountID string      `json:"accountID"`
	Username  string      `json:"username"`
	FirstName string      `json:"firstname"`
	LastName  string      `json:"lastname"`
	Email     string      `json:"email"`
	Avatar    string      `json:"avatar"`
	Create    Permissions `json:"create"`
	Update    Permissions `json:"update"`
	Read      Permissions `json:"read"`
	Delete    Permissions `json:"delete"`
}
