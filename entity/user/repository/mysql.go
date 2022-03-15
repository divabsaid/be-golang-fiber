package repository

import (
	"be-golang-fiber/entity/user"
	"database/sql"
)

type UserRepository interface {
	UserRegister(u *user.UserModel) (*user.UserModel, error)
	UserLogin(u *user.UserLoginModel) (*user.UserLoginModel, error)
	GetByID(id int) (*user.UserModel, error)
}

type mysqlUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) UserRepository {
	return &mysqlUserRepository{
		db: db,
	}
}

func (m *mysqlUserRepository) UserRegister(u *user.UserModel) (*user.UserModel, error) {
	query := "INSERT user SET first_name=?, last_name=?, password=?, email=?, role_id=?, created_at=?"
	stmt, err := m.db.Prepare(query)
	if err != nil {
		return u, err
	}
	res, err := stmt.Exec(u.Firstname, u.Lastname, u.Password, u.Email, u.RoleID, u.CreatedAt)
	if err != nil {
		return u, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return u, err
	}
	u.ID = int(id)
	return u, nil

}

func (m *mysqlUserRepository) UserLogin(u *user.UserLoginModel) (*user.UserLoginModel, error) {
	userObj := new(user.UserLoginModel)
	row := m.db.QueryRow("SELECT id, password, role_id, active FROM user WHERE email=?", u.Email)
	err := row.Scan(&userObj.ID, &userObj.Password, &userObj.RoleID, &userObj.Active)
	if err != nil {
		return userObj, err
	}
	return userObj, nil

}

func (m *mysqlUserRepository) GetByID(id int) (*user.UserModel, error) {
	userObj := new(user.UserModel)
	row := m.db.QueryRow("SELECT id, first_name, last_name, email, role_id, image_name FROM user WHERE id=?", id)
	err := row.Scan(&userObj.ID, &userObj.Firstname, &userObj.Lastname, &userObj.Email, &userObj.RoleID, &userObj.ImageName)
	if err != nil {
		return userObj, err
	}
	return userObj, nil
}
