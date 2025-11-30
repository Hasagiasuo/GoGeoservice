package postgres

import (
	"geoserv/internal/domain/models"
	"log"
)

type PostgresDriverRepository struct {
	client *PostgresClient
}

func NewPostgresDriverRepository(client *PostgresClient) *PostgresDriverRepository {
	return &PostgresDriverRepository{client}
}

func (pdr *PostgresDriverRepository) Add(name string) error {
	query := "insert into drivers (name, is_busy) values ($1, $2)"
	if _, err := pdr.client.DB.Exec(query, name, false); err != nil {
		return err
	}
	return nil
}
func (pdr *PostgresDriverRepository) Get(id int) (models.Driver, error) {
	query := "select * from drivers where id = $1"
	rows, err := pdr.client.DB.Query(query, id)
	if err != nil {
		log.Printf("driver with id %d undefined: %s", id, err.Error())
		return models.Driver{}, err
	}
	defer rows.Close()
	var res models.Driver
	for rows.Next() {
		if err := rows.Scan(&res.Id, &res.Name, &res.IsBusy); err != nil {
			log.Printf("error scan driver: %s", err.Error())
			return models.Driver{}, err
		}
	}
	return res, nil
}
func (pdr *PostgresDriverRepository) Del(id int) error {
	return nil
}
func (pdr *PostgresDriverRepository) GetAll() ([]models.Driver, error) {
	query := "select * from drivers"
	rows, err := pdr.client.DB.Query(query)
	if err != nil {
		log.Printf("cannot get all drivers: %s", err.Error())
		return nil, err
	}
	defer rows.Close()
	var res []models.Driver
	for rows.Next() {
		var sub models.Driver
		if err := rows.Scan(&sub.Id, &sub.Name, &sub.IsBusy); err != nil {
			log.Printf("cannot parse driver: %s", err.Error())
			return nil, err
		}
		res = append(res, sub)
	}
	return res, nil
}

func (pdr *PostgresDriverRepository) SetBusy(name string, is_busy bool) error {
	query := `update drivers set is_busy=$1 where name=$2`
	_, err := pdr.client.DB.Exec(query, is_busy, name)
	if err != nil {
		log.Printf("cannot update param is_busy for driver: %v", err)
		return err
	}
	return nil
}

func (pdr *PostgresDriverRepository) GetByName(name string) (*models.Driver, error) {
	query := `select * from drivers where name=$1`
	rows, err := pdr.client.DB.Query(query, name)
	if err != nil {
		log.Printf("cannot get driver by name: %v", err)
		return nil, err
	}
	defer rows.Close()
	var res models.Driver
	for rows.Next() {
		if err := rows.Scan(&res.Id, &res.Name, &res.IsBusy); err != nil {
			log.Printf("cannot parse driver: %v", err)
			return nil, err
		}
	}
	return &res, nil
}
