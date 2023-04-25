package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/eltoncasacio/sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

type CourseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

func (c *CourseDB) callTx(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	q := db.New(tx)
	err = fn(q)
	if err != nil {
		if errRb := tx.Rollback(); errRb != nil {
			return fmt.Errorf(" rollback error: %v\n original error: %v", errRb, err)
		}
		return err
	}
	return tx.Commit()
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CourseParams) error {
	err := c.callTx(ctx, func(q *db.Queries) error {
		err := q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})
		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
			CategoryID:  argsCategory.ID,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/fullcycle")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	categories, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}

	for _, c := range categories {
		fmt.Printf("Category: %v\n", c.Description)
	}

	// courseDB := NewCourseDB(dbConn)
	// err = courseDB.CreateCourseAndCategory(ctx,
	// 	CategoryParams{
	// 		ID:          uuid.New().String(),
	// 		Name:        "Desenvolvimento",
	// 		Description: sql.NullString{String: "Desenvolvimento de sistemas performaticos e confiaveis", Valid: true},
	// 	},
	// 	CourseParams{
	// 		ID:          uuid.New().String(),
	// 		Name:        "Kotlin",
	// 		Description: sql.NullString{String: "Apps nativos", Valid: true},
	// 		Price:       350.99,
	// 	},
	// )
	// if err != nil {
	// 	panic(err)
	// }
}
