package models

import (
	"database/sql"

	"github.com/husnain3184/creative-studio/config"
	"github.com/husnain3184/creative-studio/entities"
)

type ServiceModel struct {
	db *sql.DB
}

func NewService() *ServiceModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &ServiceModel{db: db}
}

func (m *ServiceModel) FindAll(service *[]entities.Service) error {

	rows, err := m.db.Query("select * from 	service")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Service
		rows.Scan(&data.Id, &data.Title, &data.Detail, &data.Image)
		*service = append(*service, data)
	}
	return nil
}

func (m *ServiceModel) Find(id int64, service *entities.Service) error {

	return m.db.QueryRow("select * from service where id = ?", id).Scan(

		&service.Id,
		&service.Title,
		&service.Detail,
		&service.Image)
}

func (m *ServiceModel) Create(service *entities.Service) error {

	result, err := m.db.Exec("insert into service (title,details,image)values(?,?,?)",
		service.Title, service.Detail, service.Image)
	if err != nil {

		return err
	}

	lastInsertId, _ := result.LastInsertId()
	service.Id = lastInsertId
	return nil

}

func (m *ServiceModel) Update(service entities.Service) error {

	_, err := m.db.Exec("update service set title = ?, title = ?, details = ?, image = ? where id = ?", service.Title, service.Detail, service.Image, service.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *ServiceModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from service where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
