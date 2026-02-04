package infrastructure

import (
	"api-gateway/pkg/domain"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Repository struct {
	TableName string
	db        *sql.DB
}

func NewPostgresRepository(tableName string, db *sql.DB) *Repository {
	return &Repository{TableName: tableName, db: db}
}
func (r *Repository) Insert(user *domain.User) (err error) {
	_, err = r.db.Exec("INSERT INTO "+r.TableName+" (id,name,age) VALUES ($1,$2,$3)", user.Id, user.Name, user.Age)
	if err != nil {
		fmt.Println("Error inserting data:", err)
		return err
	}
	return nil
}
