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
		uuid      string
		name      string
		password  string
		gender    string
		birthday  string
		point     int
		createdAt sql.NullTime
		updatedAt sql.NullTime
		deletedAt sql.NullTime
	)

	row.Next()
	if err = row.Scan(&id, &uuid,  &name, &password, &gender, &birthday, &point, &createdAt, &updatedAt, &deletedAt);
	err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	// sql.NullTimeからtime.Timeに変換するといいかも
	user = model.User {
		Id:            id,
		Uuid:          uuid,
		Name:          name,
		Password:      password,
		Gender:        gender,
		BirthDay:      birthday,
		Point:         point,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}

	return
}

func (repo *UsersRepository) GetSelectUuid(ctx context.Context, uuid string) (id int, err error) {
	row, err := repo.QueryContext(ctx, "select id from user where uuid = ?", uuid)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	row.Next()
	if err = row.Scan(&id); err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	return
}

func (repo *UsersRepository) CheckUuid(ctx context.Context, uuid string) (b bool, err error) {
	row, err := repo.QueryContext(ctx, "select count(*) from user where uuid = ?", uuid)
	if err != nil {
		log.Printf("Could not scan result with GetPass: %v", err)
		return
	}
	defer row.Close()

	var count int
	row.Next()
	if err = row.Scan(&count); err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	log.Printf("count: %d\n", count)
	if count == 0 {
		b = true
	} else {
		b = false
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
			uuid      string
			name      string
			password  string
			gender    string
			birthday  string
			point     int
			createdAt sql.NullTime
			updatedAt sql.NullTime
			deletedAt sql.NullTime
		)
		if err := rows.Scan(&id, &uuid, &name, &password, &gender, &birthday, &point, &createdAt, &updatedAt, &deletedAt);
			err != nil {
			log.Printf("row.Scan()でerror: %v\n", err)
			continue
		}
		user := model.User {
			Id:            id,
			Uuid:          uuid,
			Name:          name,
			Password:      password,
			Gender:        gender,
			BirthDay:      birthday,
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
		"insert into pbl_app1.user (uuid, name, password, gender, birthday, created_at) values (?, ?, ?, ?, ?, now())",
		uRegistry.Uuid, uRegistry.Name, uRegistry.Password, uRegistry.Gender, uRegistry.BirthDay)
	if err != nil {
		return
	}
	log.Printf("uuid: %s, name: %s, password: %s, gender: %s, birthday: %s\n",
		uRegistry.Uuid ,uRegistry.Name, uRegistry.Password, uRegistry.Gender, uRegistry.BirthDay)

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
		"update pbl_app1.user set uuid=?, name=?, password=?, gender=?, birthday=cast(? as date), updated_at=now() where id=?",
		request.Uuid ,request.Name, request.Password, request.Gender, request.BirthDay, request.Id)
	if err != nil {
		return
	}
	log.Printf("uuid: %s, name: %s, password: %s, gender: %s, birthday: %s\n",
		request.Uuid ,request.Name, request.Password, request.Gender, request.BirthDay)

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
