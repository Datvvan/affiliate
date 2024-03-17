package db

import (
	"context"
	"fmt"
	"time"

	"github.com/datvvan/affiliate/config"
	"github.com/go-pg/pg/v10"
)

type DB struct {
	*pg.DB
}

type Instance struct {
	DB *pg.DB
}

type dbLogger struct{}

var instance *Instance = nil

func New() (*Instance, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbURL := fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable",
		config.Default.DB_USER,
		config.Default.DB_PASSWORD,
		config.Default.DB_HOST,
		config.Default.DB_NAME)
	opt, err := pg.ParseURL(dbURL)
	if err != nil {
		return nil, err
	}

	db := pg.Connect(opt)
	db.AddQueryHook(dbLogger{})

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	instance = &Instance{
		DB: db,
	}

	return instance, nil
}

func GetInstance() *Instance {
	return instance
}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fq, _ := q.FormattedQuery()
	fmt.Println(string(fq))
	return nil

}
