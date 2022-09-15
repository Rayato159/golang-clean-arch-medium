package entities

type UsersUsecase interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
}

type UsersRepository interface {
	Register(req *UsersRegisterReq) (*UsersRegisterRes, error)
}

type UsersRegisterReq struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type UsersRegisterRes struct {
	Id       uint64 `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}
