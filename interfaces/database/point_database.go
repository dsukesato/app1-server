package database

import (
	"context"
	"database/sql"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"log"
)

type PointRepository struct {
	DBConn
}

func (repo *PointRepository) GetSelect(ctx context.Context, identifier int) (point model.Point, err error) {
	row, err := repo.QueryContext(ctx, "select * from point where id = ?", identifier)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	var (
		id           int
		restaurantId int
		userId       int
		transaction  string
		createdAt    sql.NullTime
	)

	row.Next()
	if err = row.Scan(&id, &restaurantId, &userId, &transaction, &createdAt);
		err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	// sql.NullTimeからtime.Timeに変換するといいかも
	point = model.Point {
		Id:            id,
		RestaurantId:  restaurantId,
		UserId:        userId,
		Transaction:   transaction,
		CreatedAt:     createdAt,
	}

	return
}

func (repo *PointRepository) GetAll(ctx context.Context) (points model.Points, err error){
	rows, err := repo.QueryContext(ctx, "select * from point")
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
			transaction   string
			createdAt     sql.NullTime
		)
		if err := rows.Scan(&id, &restaurantId, &userId, &transaction, &createdAt);
			err != nil {
			log.Printf("row.Scan()でerror: %v\n", err)
			continue
		}
		point := model.Point {
			Id:            id,
			RestaurantId:  restaurantId,
			UserId:        userId,
			Transaction:   transaction,
			CreatedAt:     createdAt,
		}
		points = append(points, point)
	}
	return
}

func (repo *PointRepository) Store(ctx context.Context, poRegistry model.PostPointRequest) (id int, err error) {
	result, err := repo.ExecContext(ctx,
		"insert into pbl_app1.point (restaurant_id, user_id, transaction, created_at) values (?, ?, ?, now())",
		poRegistry.RestaurantId, poRegistry.UserId, poRegistry.Transaction)
	if err != nil {
		return
	}
	log.Printf("restaurant_id: %d, user_id: %d, transaction: %s\n",
		poRegistry.RestaurantId, poRegistry.UserId, poRegistry.Transaction)

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	log.Printf("id: %d\n", id)

	return
}
