package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

type Hero struct {
	Id            int    `json:"id"`
	Name          string `json:"name"`
	SpecialPowers string `json:"special_powers"`
	Publisher     string `json:"publisher"`
}

var con *sql.DB

func main() {
	host := os.Getenv("MYSQL_HOST")
	if host == "" {
		host = "localhost"
	}

	datasource := fmt.Sprintf("root:superheroes@tcp(%s:3306)/heroes", host)
	db, err := sql.Open("mysql", datasource)
	db.SetMaxIdleConns(100)
	if err != nil {
		panic(err)
	}
	con = db
	r := gin.Default()
	r.GET("/heroes", heroes)
	fmt.Println("server up on 8080")
	r.Run()
}

func heroes(context *gin.Context) {
	heroes := getHeroes()
	context.SecureJSON(http.StatusOK, heroes)
}

func getHeroes() []Hero {
	rows, err := con.Query("SELECT * FROM heroes")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	heroes := []Hero{}
	for rows.Next() {
		var h Hero
		rows.Scan(&h.Id, &h.Name, &h.SpecialPowers, &h.Publisher)
		heroes = append(heroes, h)
	}
	return heroes
}
