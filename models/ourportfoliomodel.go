package models

import (
	"database/sql"

	"github.com/husnain3184/creative-studio/config"
	"github.com/husnain3184/creative-studio/entities"
)

type PortfolioModel struct {
	db *sql.DB
}

func NewPortfolio() *PortfolioModel {

	db, err := config.DBConnection()
	if err != nil {

		panic(err)
	}
	return &PortfolioModel{db: db}
}

func (m *PortfolioModel) FindAll(portfolio *[]entities.Portfolio) error {

	rows, err := m.db.Query("select * from 	portfolio")
	if err != nil {

		return err
	}

	defer rows.Close()

	for rows.Next() {

		var data entities.Portfolio
		rows.Scan(
			&data.Id,
			&data.Title,
			&data.Url,
			&data.Image)
		*portfolio = append(*portfolio, data)
	}
	return nil
}

func (m *PortfolioModel) Find(id int64, portfolio *entities.Portfolio) error {

	return m.db.QueryRow("select * from portfolio where id = ?", id).Scan(

		&portfolio.Id,
		&portfolio.Title,
		&portfolio.Url,
		&portfolio.Image)
}

func (m *PortfolioModel) Create(portfolio *entities.Portfolio) error {

	result, err := m.db.Exec("insert into portfolio (title,url,image)values(?,?,?)",
		portfolio.Title, portfolio.Url, portfolio.Image)
	if err != nil {

		return err
	}

	lastInsertId, _ := result.LastInsertId()
	portfolio.Id = lastInsertId
	return nil

}

func (m *PortfolioModel) Update(portfolio entities.Portfolio) error {

	_, err := m.db.Exec("update portfolio set title = ?, url = ? , image = ? where id = ?", portfolio.Title, portfolio.Url, portfolio.Image, portfolio.Id)

	if err != nil {
		return err
	}

	return nil
}

func (m *PortfolioModel) Delete(id int64) error {

	_, err := m.db.Exec("delete from portfolio where id = ?", id)
	if err != nil {
		return err
	}

	return nil
}
