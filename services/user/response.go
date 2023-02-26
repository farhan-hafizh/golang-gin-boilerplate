package user

type RegisterResponse struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
}

type LoginResponse struct {
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatRegisterResponse(user User) RegisterResponse {
	return RegisterResponse{
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
	}
}

func FormatLoginResponse(user User, token string) LoginResponse {
	return LoginResponse{
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      token,
	}
}
