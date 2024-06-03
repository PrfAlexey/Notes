package entity

// SignUpRequest структура HTTP-запроса на регистрацию пользователя
type SignUpRequest struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
