package entity

type UpdateSheetRequest struct {
	User  string `json:"user"`
	Email string `json:"email"`
	Hp    string `json:"hp"`
	Item  string `json:"item"`
}
