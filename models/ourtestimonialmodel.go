package models

import (
	"database/sql"

	"github.com/husnain3184/creative-studio/config"
	"github.com/husnain3184/creative-studio/entities"
)

type TestimonialModel struct {
	db *sql.DB
}

func NewTestimonial() *TestimonialModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &TestimonialModel{db: db}
}

func (m *TestimonialModel) FindAll(testimonial *[]entities.Testimonial) error {

	rows, err := m.db.Query("select * from 	testimonial")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Testimonial
		rows.Scan(
			&data.Id,
			&data.Name,
			&data.Details,
			&data.Image)
		*testimonial = append(*testimonial, data)
	}
	return nil
}

func (m *TestimonialModel) Find(id int64, testimonial *entities.Testimonial) error {

	return m.db.QueryRow("select * from testimonial where id = ?", id).Scan(

		&testimonial.Id,
		&testimonial.Name,
		&testimonial.Details,
		&testimonial.Image)
}

func (m *TestimonialModel) Create(testimonial *entities.Testimonial) error {

	result, err := m.db.Exec("insert into testimonial (name,details,image)values(?,?,?)",
		testimonial.Name, testimonial.Details, testimonial.Image)
	if err != nil {

		return err
	}

	lastInsertId, _ := result.LastInsertId()
	testimonial.Id = lastInsertId
	return nil

}

func (m *TestimonialModel) Update(testimonial entities.Testimonial) error {

	_, err := m.db.Exec("update testimonial set name = ?, details = ? , image = ? where id = ?", testimonial.Name, testimonial.Details, testimonial.Image, testimonial.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *TestimonialModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from testimonial where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
