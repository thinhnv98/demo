package repositories

import "database/sql"

type IUserRepo interface {
	RegisterAccount() error
}


type UserRepo struct {
	Db *sql.DB
}

func (_self UserRepo) RegisterAccount() error {
	return nil
}
