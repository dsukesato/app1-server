package database

import (
	"context"
	"database/sql"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
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
		gender    string
		birthday  string
		state     bool
		point     int
		createdAt sql.NullTime
		updatedAt sql.NullTime
		deletedAt sql.NullTime
	)

	row.Next()
	if err = row.Scan(&id, &name, &password, &gender, &birthday, &state, &point, &createdAt, &updatedAt, &deletedAt);
	err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	// sql.NullTimeからtime.Timeに変換するといいかも
	user = model.User {
		Id:            id,
		Name:          name,
		Password:      password,
		Gender:        gender,
		BirthDay:      birthday,
		State:         state,
		Point:         point,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}

	return
}

func (repo *UsersRepository) GetPass(ctx context.Context, identifier int) (pass string, err error) {
	row, err := repo.QueryContext(ctx, "select password from user where id = ?", identifier)
	if err != nil {
		log.Printf("Could not scan result with GetPass: %v", err)
		return
	}
	defer row.Close()

	row.Next()
	if err = row.Scan(&pass);
		err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
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
			id        int
			name      string
			password  string
			gender    string
			birthday  string
			state     bool
			point     int
			createdAt sql.NullTime
			updatedAt sql.NullTime
			deletedAt sql.NullTime
		)
		if err := rows.Scan(&id, &name, &password, &gender, &birthday, &state, &point, &createdAt, &updatedAt, &deletedAt);
			err != nil {
			log.Printf("row.Scan()でerror: %v\n", err)
			continue
		}
		user := model.User {
			Id:            id,
			Name:          name,
			Password:      password,
			Gender:        gender,
			BirthDay:      birthday,
			State:         state,
			Point:         point,
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
		"insert into pbl_app1.user (name, password, gender, birthday, created_at) values (?, ?, ?, ?, now())",
		uRegistry.Name, uRegistry.Password, uRegistry.Gender, uRegistry.BirthDay)
	if err != nil {
		return
	}
	log.Printf("name: %s, password: %s, gender: %s, birthday: %s\n",
		uRegistry.Name, uRegistry.Password, uRegistry.Gender, uRegistry.BirthDay)

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	log.Printf("id: %d\n", id)

	return
}

func (repo *UsersRepository) Change(ctx context.Context, request model.PutUserRequest) (id int, err error) {
	_, err = repo.ExecContext(ctx,
		"update pbl_app1.user set name=?, password=?, gender=?, birthday=cast(? as date), updated_at=now() where id=?",
		request.Name, request.Password, request.Gender, request.BirthDay, request.Id)
	if err != nil {
		return
	}
	log.Printf("name: %s, password: %s, gender: %s, birthday: %s\n",
		request.Name, request.Password, request.Gender, request.BirthDay)

	row, err := repo.QueryContext(ctx, "select id from user where id = ?", request.Id)
	if err != nil {
		log.Printf("Could not scan result with GetAll: %v", err)
		return
	}
	defer row.Close()

	row.Next()
	if err = row.Scan(&id);
	err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
	}
	//id = int(id64)
	log.Printf("id: %d\n", id)

	return
}
