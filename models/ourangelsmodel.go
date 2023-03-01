package models

import (
	"database/sql"

	"github.com/husnain3184/creative-studio/config"
	"github.com/husnain3184/creative-studio/entities"
)

type AngelsModel struct {
	db *sql.DB
}

func NewAngels() *AngelsModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &AngelsModel{db: db}
}

func (m *AngelsModel) FindAll(angels *[]entities.Angels) error {

	rows, err := m.db.Query("select * from 	angels")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Angels
		rows.Scan(&data.Id, &data.Designation, &data.Name, &data.Detail, &data.Facebook, &data.Github, &data.Dropbox, &data.Twitter, &data.Image)
		*angels = append(*angels, data)
	}
	return nil
}

func (m *AngelsModel) Find(id int64, angels *entities.Angels) error {

	return m.db.QueryRow("select * from angels where id = ?", id).Scan(

		&angels.Id,
		&angels.Designation,
		&angels.Name,
		&angels.Detail,
		&angels.Facebook,
		&angels.Github,
		&angels.Dropbox,
		&angels.Twitter,
		&angels.Image)
}

func (m *AngelsModel) Create(angels *entities.Angels) error {

	result, err := m.db.Exec("insert into angels (designation,name,detail,facebook,github,dropbox,Twitter,image)values(?,?,?,?,?,?,?,?)",
		angels.Designation, angels.Name, angels.Detail, angels.Facebook, angels.Github, angels.Dropbox, angels.Twitter, angels.Image)
	if err != nil {

		return err
	}

	lastInsertId, _ := result.LastInsertId()
	angels.Id = lastInsertId
	return nil

}

func (m *AngelsModel) Update(angels entities.Angels) error {

	_, err := m.db.Exec("update angels set designation = ?, name = ?, detail = ?, facebook = ?, github = ? , dropbox = ? , Twitter = ? , image = ? where id = ?", angels.Designation, angels.Name, angels.Detail, angels.Facebook, angels.Github, angels.Dropbox, angels.Twitter, angels.Image, angels.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *AngelsModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from angels where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
