package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)


type Leads struct {
	Partner string `json:"partner" pg:"partner"`
	Users int `json:"users" pg:"users"`
	Active_Leads  int  `json:"active_leads" pg:"active_leads"`
	Unassigned_Leads int `json:"unassigned_leads" pg:"unassigned_leads"`
	Active_Prospect int `json:"active_prospects" pg:"active_prospects"`
	Active_Opportunity int `json:"active_opportunity" pg:"active_opportunity"`
}

func main() {
	db := loadDatabase()

	router := gin.Default()
	router.Use(cors.Default())

	router.GET("/leads",func(ctx *gin.Context) {
		var leads []Leads
		err := db.Model(&leads).Select()
		if err != nil {
			ctx.JSON(500, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, leads)
	})

	router.Run(":8080")
}


func loadDatabase() *pg.DB {

	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "shouryagautam",
		Password: "shourya",
		Database: "postgres",
	})

	err := db.Model((*Leads)(nil)).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})

	if err != nil {
		panic(err)
	}

	return db
}