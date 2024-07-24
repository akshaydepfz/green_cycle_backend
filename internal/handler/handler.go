package handler

import (
	"database/sql"

	"example.com/main/models"
	"github.com/lib/pq"
)

var DB *sql.DB

func PostUser(name string, orders *pq.StringArray) (uint, error) {
	var id uint

	err := DB.QueryRow("INSERT INTO users (name, orders) VALUES ($1,$2) RETURNING id", name, orders).Scan(&id)
	if err != nil {
		return 0, nil
	}

	return id, nil
}

func GetUser() ([]models.User, error) {
	rows, err := DB.Query("SELECT id, name, orders FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Orders)
		if err != nil {

		}

		users = append(users, user)
	}

	return users, nil
}
