package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://wjjunior:Go123@db:5432/clientdb?sslmode=disable")
	if err != nil {
		panic(err)
	}

	// Create the "clients" table if it doesn't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS clients (
		id SERIAL PRIMARY KEY,
		name VARCHAR(50),
		age INT
	)`)
	if err != nil {
		panic(err)
	}
}

type Client struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var clients = map[int]*Client{}

func GetClients() map[int]*Client {
	return clients
}

func GetClient(id int) (*Client, bool) {
	client, ok := clients[id]
	return client, ok
}

func CreateClient(client *Client) {
	client.ID = len(clients) + 1
	clients[client.ID] = client
}

func DeleteClient(id int) {
	delete(clients, id)
}
