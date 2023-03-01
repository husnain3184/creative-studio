package models

import (
	"database/sql"

	"github.com/husnain3184/creative-studio/config"
	"github.com/husnain3184/creative-studio/entities"
)

type BlogsModel struct {
	db *sql.DB
}

func NewBlogs() *BlogsModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &BlogsModel{db: db}
}

func (m *BlogsModel) FindAll(blogs *[]entities.Blogs) error {

	rows, err := m.db.Query("select * from 	blogs")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Blogs
		rows.Scan(&data.Id, &data.Detail, &data.Image, &data.Title)
		*blogs = append(*blogs, data)
	}
	return nil
}

func (m *BlogsModel) Find(id int64, blogs *entities.Blogs) error {

	return m.db.QueryRow("select * from blogs where id = ?", id).Scan(

		&blogs.Id,
		&blogs.Detail,
		&blogs.Image,
		&blogs.Title)
}

func (m *BlogsModel) Create(blogs *entities.Blogs) error {

	result, err := m.db.Exec("insert into blogs (details,image,title)values(?,?,?)",
		blogs.Detail, blogs.Image, blogs.Title)
	if err != nil {

		return err
	}

	lastInsertId, _ := result.LastInsertId()
	blogs.Id = lastInsertId
	return nil

}

func (m *BlogsModel) Update(blogs entities.Blogs) error {

	_, err := m.db.Exec("update blogs set details = ?, image = ? , title = ? where id = ?", blogs.Detail, blogs.Image, blogs.Title, blogs.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *BlogsModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from blogs where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
