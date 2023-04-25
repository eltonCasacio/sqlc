package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eltoncasacio/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/fullcycle")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// _, err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "DB",
	// 	Description: sql.NullString{String: "Database com SQLC", Valid: true},
	// })
	// if err != nil {
	// 	panic(err)
	// }

	// queries.DeleteCategory(ctx, "ec8f6ac6-8979-4bfc-84fe-5f554a7913f3")

	err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
		Name:        "Repository",
		Description: sql.NullString{String: "SQLC enjoy", Valid: true},
		ID:          "75d8a0b6-e1b7-4cc2-ae7d-14b042f0b2d6",
	})
	if err != nil {
		panic(err)
	}

	categories, err := queries.ListCategories(ctx)
	if err != nil {
		panic(err)
	}

	for _, v := range categories {
		fmt.Println(v)
	}

}
