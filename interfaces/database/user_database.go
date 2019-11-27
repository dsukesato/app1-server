package database

import (
	"context"
	"database/sql"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
	"log"
)

type UsersRepository struct {
	DBConn
}

func (repo *UsersRepository) GetSelect(ctx context.Context, identifier int) (user model.User, err error) {
	row, err := repo.QueryContext(ctx, "select * from user where id = ?", identifier)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	var (
		id        int
		name      string
		password  string
		createdAt sql.NullTime
		updatedAt sql.NullTime
		deletedAt sql.NullTime
	)

	row.Next()
	if err = row.Scan(&id, &name, &password, &createdAt, &updatedAt, &deletedAt);
		err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	// sql.NullTimeからtime.Timeに変換するといいかも
	user = model.User {
		Id:            id,
		Name:          name,
		Password:      password,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}

	return
}

func (repo *UsersRepository) GetAll(ctx context.Context) (users model.Users, err error){
	rows, err := repo.QueryContext(ctx, "select * from user")
	if err != nil {
		log.Printf("Could not scan result with GetAll: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id            int
			name          string
			password      string
			createdAt     sql.NullTime
			updatedAt     sql.NullTime
			deletedAt     sql.NullTime
		)
		if err := rows.Scan(&id, &name, &password, &createdAt, &updatedAt, &deletedAt);
			err != nil {
			log.Printf("row.Scan()でerror: %v\n", err)
			continue
		}
		user := model.User {
			Id:            id,
			Name:          name,
			Password:      password,
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
			DeletedAt:     deletedAt,
		}
		users = append(users, user)
	}
	return
}

func (repo *UsersRepository) Store(ctx context.Context, uRegistry model.PostUserRequest) (id int, err error) {
	result, err := repo.ExecContext(ctx,
		"insert into pbl_app1.user (name, password, created_at) values (?, ?, now())",
		uRegistry.Name, uRegistry.Password)
	if err != nil {
		return
	}
	log.Printf("name: %s, password: %s\n",
		uRegistry.Name, uRegistry.Password)

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	log.Printf("id: %d\n", id)

	return
}
