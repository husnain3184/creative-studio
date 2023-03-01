package models

import (
	"database/sql"

	"github.com/husnain3184/creative-studio/config"
	"github.com/husnain3184/creative-studio/entities"
)

type UsersModel struct {
	db *sql.DB
}

func NewUsers() *UsersModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &UsersModel{db: db}
}

func (m *UsersModel) FindAll(users *[]entities.Users) error {

	rows, err := m.db.Query("select * from 	users")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Users
		rows.Scan(
			&data.Id,
			&data.Name,
			&data.Username,
			&data.Email,
			&data.Password,
			&data.Roll)
		*users = append(*users, data)
	}
	return nil
}

func (m *UsersModel) Find(id int64, users *entities.Users) error {

	return m.db.QueryRow("select * from users where id = ?", id).Scan(

		&users.Id,
		&users.Name,
		&users.Username,
		&users.Email,
		&users.Password,
		&users.Roll)
}

func (m *UsersModel) Create(users *entities.Users) error {

	result, err := m.db.Exec("insert into users (name, email, username, password, roll )values(?,?,?,?,?)",
		users.Name,
		users.Email,
		users.Username,
		users.Password,
		users.Roll)

	if err != nil {

		return err
	}

	lastInsertId, _ := result.LastInsertId()
	users.Id = lastInsertId
	return nil

}

func (m *UsersModel) Update(users entities.Users) error {

	_, err := m.db.Exec("update users set name = ?, email = ? , username = ? , password = ? , roll = ?  where id = ?",
		users.Name,
		users.Username,
		users.Email,
		users.Password,
		users.Roll)

	if err != nil {
		return err
	}

	return nil
}

func (m *UsersModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from users where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}

func (u UsersModel) Where(admin *entities.Admin, fieldName, fieldValue string) error {
	row, err := u.db.Query("select * from admin where "+fieldName+"= ? limit 1", fieldValue)
	// log.Println(row)

	if err != nil {
		return err
	}
	defer row.Close()
	for row.Next() {
		row.Scan(&admin.Id, &admin.Name, &admin.Email, &admin.Username, &admin.Password)
	}
	return nil
}
