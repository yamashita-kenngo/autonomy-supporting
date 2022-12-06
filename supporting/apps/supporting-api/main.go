package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"supporting/apps/supporting-api/infra/postgres/models"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func Hello(name string) string {
	result := "Hello " + name
	return result
}

func fetchOneGovernmentByID(db *sql.DB, ID int) (*models.Government, error) {
	g, err := models.FindGovernment(context.Background(), db, 1)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return g, nil
}

func createGovernment(db *sql.DB) {
	gov := models.Government{
		ID:   2,
		Name: null.StringFrom("竜宮市"),
	}
	err := gov.Insert(context.Background(), db, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}
}

func fetchAllGovernments(db *sql.DB) ([]*models.Government, error) {
	gs, err := models.Governments().All(context.Background(), db)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return gs, nil
}

func deleteAllGovernments(db *sql.DB) {
	gs, err := models.Governments().All(context.Background(), db)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := gs.DeleteAll(context.Background(), db); err != nil {
		log.Fatal(err)
	}
}

func updateOneGovenmentNameByID(db *sql.DB, ID int) {
	g, err := models.FindGovernment(context.Background(), db, ID)
	if err != nil {
		log.Fatal(err)
	}
	g.Name = null.StringFrom("廿日市市")
	if _, err := g.Update(context.Background(), db, boil.Infer()); err != nil {
		log.Fatal(err)
	}
}

func dbConnect() *sql.DB {
	fmt.Println("connect database")
	uri := "postgres://admin:password@localhost:5432/ausu?sslmode=disable"
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func main() {
	fmt.Println(Hello("supporting-api"))
	db := dbConnect()
	boil.SetDB(db)
	updateOneGovenmentNameByID(db, 1)
	g1, err := fetchOneGovernmentByID(db, 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", g1)
}
