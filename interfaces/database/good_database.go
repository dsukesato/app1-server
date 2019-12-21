package database

import (
	"context"
	"github.com/dsukesato/go13/pbl/app1-server/entity/model"
	"log"
)

type GoodRepository struct {
	DBConn
}

func (repo *GoodRepository) GetSelect(ctx context.Context, identifier int) (good model.Good, err error) {
	row, err := repo.QueryContext(ctx, "select * from good where id = ?", identifier)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	var (
		id int
		postId int
		userId int
		state bool
	)

	row.Next()
	if err = row.Scan(&id, &postId, &userId, &state);
		err != nil {
		log.Printf("row.Scan()でerror: %v with GetSelect\n", err)
		return
	}
	// sql.NullTimeからtime.Timeに変換するといいかも
	good = model.Good {
		Id: id,
		PostId: postId,
		UserId: userId,
		State:  state,
	}

	return
}

func (repo *GoodRepository) GetId(ctx context.Context, pid int, uid int) (id int, err error) {
	row, err := repo.QueryContext(ctx, "select id from good where post_id = ? and user_id = ?", pid, uid)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	row.Next()
	if err = row.Scan(&id);
		err != nil {
		log.Printf("row.Scan()でerror: %v with GetIdt\n", err)
		return
	}
	return
}

func (repo *GoodRepository) GetGood(ctx context.Context, identifier int) (id int, err error) {
	row, err := repo.QueryContext(ctx, "select post_id from good where id = ?", identifier)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	var pId int
	row.Next()
	if err = row.Scan(&pId); err != nil {
		log.Printf("row.Scan()でerror: %v with GetGood\n", err)
		return
	}
	row, err = repo.QueryContext(ctx, "select good from post where id = ?", pId)
	if err != nil {
		log.Printf("Could not scan result with GetSelect: %v", err)
		return
	}
	defer row.Close()

	row.Next()
	if err = row.Scan(&id); err != nil {
		log.Printf("row.Scan()でerror: %v with GetGood\n", err)
		return
	}
	return
}

func (repo *GoodRepository) GetSelectPUId(ctx context.Context, pid int , uid int) (b bool) {
	row, err := repo.QueryContext(ctx, "select count(*) from good where post_id = ? and user_id = ?", pid, uid)
	if err != nil {
		log.Printf("Could not scan result with GetSelectPUId: %v", err)
		return
	}
	defer row.Close()

	var count int

	for row.Next() {
		if err := row.Scan(&count); err != nil {
			log.Fatal(err)
		}
	}
	log.Printf("count: %d\n", count)
	if count == 0 {
		b = true
	} else {
		b = false
	}
	return
}

func (repo *GoodRepository) GetAll(ctx context.Context) (goods model.Goods, err error){
	rows, err := repo.QueryContext(ctx, "select * from good")
	if err != nil {
		log.Printf("Could not scan result with GetAll: %v", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			id int
			postId int
			userId int
			state bool
		)
		if err := rows.Scan(&id, &postId, &userId, &state);
			err != nil {
			log.Printf("row.Scan()でerror: %v with GetAll\n", err)
			continue
		}
		good := model.Good {
			Id: id,
			PostId: postId,
			UserId: userId,
			State:  state,
		}
		goods = append(goods, good)
	}
	return
}

func (repo *GoodRepository) Store(ctx context.Context, good model.PostGoodRequest) (id int, err error) {
	result, err := repo.ExecContext(ctx,
		"insert into pbl_app1.good (post_id, user_id) values (?, ?)",
		good.PostId, good.UserId)
	if err != nil {
		return
	}
	log.Printf("post_id: %d, user_id: %d\n", good.PostId, good.UserId)

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}
	id = int(id64)
	log.Printf("id: %d\n", id)

	return
}

func (repo *GoodRepository) Change(ctx context.Context, good model.PutGoodRequest) (state bool, err error) {
	_, err = repo.ExecContext(ctx,
		"update pbl_app1.good set state = ? where post_id = ? and user_id = ?",
		good.State, good.PostId, good.UserId)
	if err != nil {
		return
	}
	log.Printf("post_id: %d, user_id: %d, state: %t\n", good.PostId, good.UserId, good.State)
	state = good.State
	return
}

func (repo *GoodRepository) CountIncrease(ctx context.Context, pid int) (identifier int, nGood int, err error) {
	_, err = repo.ExecContext(ctx, "update pbl_app1.post set good = good + 1 where id = ?", pid)
	if err != nil {
		return
	}
	row, err := repo.QueryContext(ctx, "select id, good from post where id = ?", pid)
	if err != nil {
		log.Printf("Could not scan result with CountIncrease: %v", err)
		return
	}
	defer row.Close()

	var (
		id int
		good int
	)

	row.Next()
	if err = row.Scan(&id, &good);
		err != nil {
		log.Printf("row.Scan()でerror: %v with CountIncrease\n", err)
		return
	} else {
		identifier = id
		nGood = good
	}
	return
}

func (repo *GoodRepository) CountDecrease(ctx context.Context, pid int) (identifier int, nGood int, err error) {
	_, err = repo.ExecContext(ctx, "update pbl_app1.post set good = good - 1 where id = ?", pid)
	if err != nil {
		return
	}
	row, err := repo.QueryContext(ctx, "select id, good from post where id = ?", pid)
	if err != nil {
		log.Printf("Could not scan result with CountDecrease: %v", err)
		return
	}
	defer row.Close()

	var (
		id int
		good int
	)

	row.Next()
	if err = row.Scan(&id, &good);
		err != nil {
		log.Printf("row.Scan()でerror: %v with CountDecrease\n", err)
		return
	} else {
		identifier = id
		nGood = good
	}
	return
}
