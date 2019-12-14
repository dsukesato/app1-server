package database

import (
	"context"
	"database/sql"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"log"
)

type RecognizeRepository struct {
	DBConn
}

func (repo *RecognizeRepository) GetSelect(ctx context.Context, identifier int) (recognize model.Recognize, err error) {
	row, err := repo.QueryContext(ctx, "select * from recognize where id = ?", identifier)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	var (
		id           int
		restaurantId int
		userId        int
		createdAt    sql.NullTime
		updatedAt    sql.NullTime
		deletedAt    sql.NullTime
	)

	row.Next()
	if err = row.Scan(&id, &restaurantId, &userId, &createdAt, &updatedAt, &deletedAt);
		err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	// sql.NullTimeからtime.Timeに変換するといいかも
	recognize = model.Recognize {
		Id:            id,
		RestaurantId:  restaurantId,
		UserId:        userId,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}

	return
}

func (repo *RecognizeRepository) GetSelectUID(ctx context.Context, uid int) (rids []int, err error) {
	row, err := repo.QueryContext(ctx, "select restaurant_id from recognize where user_id = ?", uid)
	if err != nil {
		log.Printf("Could not scan result with GetSelectUID(recognize-table): %v", err)
		return
	}
	defer row.Close()

	for row.Next() {
		var restaurantId int

		if err = row.Scan(&restaurantId); err != nil {
			log.Printf("row.Scan()でerror: %v\n", err)
			return
		}
		rids = append(rids, restaurantId)
	}
	return
}

func (repo *RecognizeRepository) GetSelectRID(ctx context.Context, rid int) (rr model.RecognizeResponse, err error) {
	rows, err := repo.QueryContext(ctx, "select id, name, image, created_at from restaurant where id = ?", rid)
	if err != nil {
		log.Printf("Could not scan result with GetSelectUID(restaurant-table): %v", err)
		return
	}
	defer rows.Close()

	var (
		id int
		name string
		image string
		createdAt sql.NullTime
	)
	rows.Next()
	if err = rows.Scan(&id, &name, &image, &createdAt); err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	rr = model.RecognizeResponse{
		Id: id,
		Name: name,
		Image: image,
		CreatedAt: createdAt,
	}
	return
}

func (repo *RecognizeRepository) GetAll(ctx context.Context) (rec model.Rec, err error){
	rows, err := repo.QueryContext(ctx, "select * from recognize")
	if err != nil {
		log.Printf("Could not scan result with GetAll: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id            int
			restaurantId  int
			userId        int
			createdAt     sql.NullTime
			updatedAt     sql.NullTime
			deletedAt     sql.NullTime
		)
		if err := rows.Scan(&id, &restaurantId, &userId, &createdAt, &updatedAt, &deletedAt);
			err != nil {
			log.Printf("row.Scan()でerror: %v\n", err)
			continue
		}
		recognize := model.Recognize {
			Id:            id,
			RestaurantId:  restaurantId,
			UserId:        userId,
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
			DeletedAt:     deletedAt,
		}
		rec = append(rec, recognize)
	}
	return
}

func (repo *RecognizeRepository) Store(ctx context.Context, reRegistry model.PostRecognizeRequest) (id int, err error) {
	result, err := repo.ExecContext(ctx,
		"insert into pbl_app1.recognize (restaurant_id, user_id, created_at) values (?, ?, now())",
		reRegistry.RestaurantId, reRegistry.UserId)
	if err != nil {
		return
	}
	log.Printf("restaurant_id: %d, user_id: %d\n",
		reRegistry.RestaurantId, reRegistry.UserId)

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	log.Printf("id: %d\n", id)

	return
}
