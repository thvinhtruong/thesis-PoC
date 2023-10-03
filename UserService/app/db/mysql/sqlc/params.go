package db

type RegisterUserParams struct {
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

type RegisterUserResult struct {
	ID int32
}

type LoginUserParams struct {
	Email    string `json:"phone"`
	Password string `json:"password"`
}

type LoginUserResult struct {
	ID int32
}

type UpdateUserParams struct {
	ID       int32  `json:"id"`
	Fullname string `json:"fullname"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}
