package database

import (
	"context"
	"database/sql"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
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
		createdAt    sql.NullTime
		updatedAt    sql.NullTime
		deletedAt    sql.NullTime
	)

	row.Next()
	if err = row.Scan(&id, &restaurantId, &createdAt, &updatedAt, &deletedAt);
		err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	// sql.NullTimeからtime.Timeに変換するといいかも
	recognize = model.Recognize {
		Id:            id,
		RestaurantId:  restaurantId,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
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
			createdAt     sql.NullTime
			updatedAt     sql.NullTime
			deletedAt     sql.NullTime
		)
		if err := rows.Scan(&id, &restaurantId, &createdAt, &updatedAt, &deletedAt);
			err != nil {
			log.Printf("row.Scan()でerror: %v\n", err)
			continue
		}
		recognize := model.Recognize {
			Id:            id,
			RestaurantId:  restaurantId,
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
		"insert into pbl_app1.recognize (restaurant_id, created_at) values (?, now())",
		reRegistry.RestaurantId)
	if err != nil {
		return
	}
	log.Printf("restaurant_id: %d\n",
		reRegistry.RestaurantId)

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	log.Printf("id: %d\n", id)

	return
}
