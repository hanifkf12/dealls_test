package presentation

type Payment struct {
	UserID      int     `json:"user_id,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	PackageName string  `json:"package_name,omitempty"`
}
