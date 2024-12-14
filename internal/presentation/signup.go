package presentation

type Signup struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Age      int    `json:"age,omitempty"`
	Bio      string `json:"bio,omitempty"`
	Avatar   string `json:"avatar,omitempty"`
}
