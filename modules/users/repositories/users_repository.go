package repositories

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rayato159/clean-arch-medium/modules/entities"
)

type usersRepo struct {
	Db *sqlx.DB
}

func NewUsersRepository(db *sqlx.DB) entities.UsersRepository {
	return &usersRepo{
		Db: db,
	}
}

func (r *usersRepo) Register(req *entities.UsersRegisterReq) (*entities.UsersRegisterRes, error) {
	query := `
	INSERT INTO "users"(
		"username",
		"password"
	)
	VALUES ($1, $2)
	RETURNING "id", "username";
	`

	// Initail a user object
	user := new(entities.UsersRegisterRes)

	// Query part
	rows, err := r.Db.Queryx(query, req.Username, req.Password)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	for rows.Next() {
		if err := rows.StructScan(user); err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
	}
	defer r.Db.Close()

	return user, nil
}

func (r *usersRepo) FindOneUser(username string) (*entities.UsersPassport, error) {
	query := `
	SELECT
	"id",
	"username",
	"password"
	FROM "users"
	WHERE "username" = $1;
	`

	res := new(entities.UsersPassport)
	if err := r.Db.Get(res, query, username); err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("error, user not found")
	}
	return res, nil
}
