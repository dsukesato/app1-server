package database

import (
	"context"
	"database/sql"
	"github.com/dsukesato/go13/pbl/app1-server/domain/model"
	"log"
)

type PostsRepository struct {
	DBConn
}

func (repo *PostsRepository) GetLastId(ctx context.Context) (identifier int, err error) {
	row, err := repo.QueryContext(ctx, "select id from post order by id desc limit 1")
	if err != nil {
		log.Printf("Could not query result with GetLastId: %v", err)
	}
	defer row.Close()

	row.Next()
	if err = row.Scan(&identifier); err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
	}
	return
}

func (repo *PostsRepository) GetSelect(ctx context.Context, identifier int) (post model.Post, err error) {
	row, err := repo.QueryContext(ctx, "select * from post where id = ?", identifier)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	var (
		id int
		userId int
		restaurantId int
		image string
		good int
		genre string
		comment string
		createdAt sql.NullTime
		updatedAt sql.NullTime
		deletedAt sql.NullTime
	)

	row.Next()
	if err = row.Scan(&id, &userId, &restaurantId, &image, &good, &genre, &comment, &createdAt, &updatedAt, &deletedAt);
	err != nil {
		log.Printf("row.Scan()でerror: %v\n", err)
		return
	}
	// sql.NullTimeからtime.Timeに変換するといいかも
	post = model.Post{
		Id: id,
		UserId: userId,
		RestaurantId: restaurantId,
		Image: image,
		Good: good,
		Genre: genre,
		Comment: comment,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}

	return
}

func (repo *PostsRepository) GetAll(ctx context.Context) (posts model.Posts, err error){
	rows, err := repo.QueryContext(ctx, "select * from post")
	if err != nil {
		log.Printf("Could not scan result with GetAll: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id int
			userId int
			restaurantId int
			image string
			good int
			genre string
			comment string
			createdAt sql.NullTime
			updatedAt sql.NullTime
			deletedAt sql.NullTime
		)
		if err := rows.Scan(&id, &userId, &restaurantId, &image, &good, &genre, &comment, &createdAt, &updatedAt, &deletedAt);
		err != nil {
			log.Printf("row.Scan()でerror: %v\n", err)
			continue
		}
		post := model.Post{
			Id: id,
			UserId: userId,
			RestaurantId: restaurantId,
			Image: image,
			Good: good,
			Genre: genre,
			Comment: comment,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			DeletedAt: deletedAt,
		}
		posts = append(posts, post)
	}
	return
}

func (repo *PostsRepository) Store(ctx context.Context, posting model.PostsRequest) (id int, err error) {
	result, err := repo.ExecContext(ctx,
		"insert into pbl_app1.post (user_id, restaurant_id, image, genre, comment, created_at) values (?, ?, ?, ?, ?, now())",
		posting.UserId, posting.RestaurantId, posting.Image, posting.Genre, posting.Comment)
	if err != nil {
		return
	}
	log.Printf("user_id: %d, restaurant_id: %d, image: %s, genre: %s, comment: %s\n",
		posting.UserId, posting.RestaurantId, posting.Image, posting.Genre, posting.Comment)

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	log.Printf("id: %d\n", id)

	return
}
