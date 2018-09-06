package fetch

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DbObject struct {
}

func (o *DbObject) All() ([]interface{}, error) {
	db, err := sqlx.Connect("mysql", "cmfive:cmfive@/cmfive")
	if err != nil {
		log.Fatalln(err)
	}

	serves := []FoodServes{}
	db.Select(&serves, "SELECT id, user_id, food_group_id, serves, d_date, notes FROM food_serves")

}
