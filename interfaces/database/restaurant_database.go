package database

import (
	"context"
	"database/sql"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
	"log"
)

type RestaurantsRepository struct {
	DBConn
}

func (repo *RestaurantsRepository) GetSelectById(ctx context.Context, identifier int) (restaurant model.Restaurant, err error) {
	row, err := repo.QueryContext(ctx, "select * from restaurant where id = ?", identifier)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	var (
		id            int
		name          string
		businessHours string
		image         string
		createdAt     sql.NullTime
		updatedAt     sql.NullTime
		deletedAt     sql.NullTime
	)

	row.Next()
	if err = row.Scan(&id, &name, &businessHours, &image, &createdAt, &updatedAt, &deletedAt);
		err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	// sql.NullTimeからtime.Timeに変換するといいかも
	restaurant = model.Restaurant {
		Id:            id,
		Name:          name,
		BusinessHours: businessHours,
		Image:         image,
		CreatedAt:     createdAt,
		UpdatedAt:     updatedAt,
		DeletedAt:     deletedAt,
	}

	return
}

func (repo *RestaurantsRepository) GetAll(ctx context.Context) (restaurants model.Restaurants, err error){
	rows, err := repo.QueryContext(ctx, "select * from restaurant")
	if err != nil {
		log.Printf("Could not scan result with GetAll: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id            int
			name          string
			businessHours string
			image         string
			createdAt     sql.NullTime
			updatedAt     sql.NullTime
			deletedAt     sql.NullTime
		)
		if err := rows.Scan(&id, &name, &businessHours, &image, &createdAt, &updatedAt, &deletedAt);
			err != nil {
			log.Printf("row.Scan()でerror: %v\n", err)
			continue
		}
		restaurant := model.Restaurant {
			Id:            id,
			Name:          name,
			BusinessHours: businessHours,
			Image:         image,
			CreatedAt:     createdAt,
			UpdatedAt:     updatedAt,
			DeletedAt:     deletedAt,
		}
		restaurants = append(restaurants, restaurant)
	}
	return
}

func (repo *RestaurantsRepository) Store(ctx context.Context, rRegistry model.PostRestaurantRequest) (id int, err error) {
	result, err := repo.ExecContext(ctx,
		"insert into pbl_app1.restaurant (name, business_hours, image, created_at) values (?, ?, ?, now())",
		rRegistry.Name, rRegistry.BusinessHours, rRegistry.Image)
	if err != nil {
		return
	}
	log.Printf("name: %s, business_hours: %s, image: %s\n",
		rRegistry.Name, rRegistry.BusinessHours, rRegistry.Image)

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	log.Printf("id: %d\n", id)

	return
}