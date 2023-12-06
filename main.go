package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	opt "github.com/jaimesoad/go-optional"
	_ "github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
	"sqlc.test/src"
)

//go:embed sql/schema.sql
var ddl string

func NewDbConn() (*src.Queries, context.Context) {
	/*opt.Fatal(godotenv.Load())

	user := os.Getenv("USER")
	passwd := os.Getenv("PASSWD")
	dbName := os.Getenv("DBNAME")
	host := os.Getenv("HOST")

	cfg := mysql.Config{
		User:   user,
		Passwd: passwd,
		DBName: dbName,
		Addr:   host,
	}*/

	ctx := context.Background()

	db := opt.Get(sql.Open("sqlite3", "database.db")).Unwrap()

	db.ExecContext(ctx, ddl)

	return src.New(db), ctx
}

type number interface {
	int | int64
}

func stoi[T number](val string) T {
	return T(opt.Get(strconv.Atoi(val)).Default(0))
}

func main() {
	q, ctx := NewDbConn()
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept" + strings.Join([]string{
			fiber.MethodGet,
			fiber.MethodPost,
			fiber.MethodHead,
			fiber.MethodPut,
			fiber.MethodDelete,
			fiber.MethodPatch,
		}, ","),
	}))

	app.Get("/todo/id/:id", func(c *fiber.Ctx) error {
		id := stoi[int64](c.Params("id"))

		todo := opt.Get(q.TodoById(ctx, id)).Default(src.Todo{})

		return c.JSON(todo)
	})

	app.Get("/todo/last", func(c *fiber.Ctx) error {
		todo := opt.Get(q.LastTenTodos(ctx)).Default([]src.Todo{})

		return c.JSON(todo)
	})

	app.Patch("/todo/id/:id", func(c *fiber.Ctx) error {
		id := stoi[int64](c.Params("id"))

		if opt.Get(q.ToggleTodo(ctx, id)).Ok() {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	})

	app.Post("/todo", func(c *fiber.Ctx) error {
		var newTodo src.Todo

		if !opt.Ok(json.Unmarshal(c.Body(), &newTodo)) {
			return c.SendStatus(fiber.StatusBadRequest)
		}

		if opt.Get(q.CreateTodo(ctx, newTodo.Content)).Ok() {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.SendStatus(fiber.StatusInternalServerError)
	})

	app.Put("/todo/id/:id", func(c *fiber.Ctx) error {
		var chName src.ChangeNameParams

		if !opt.Ok(json.Unmarshal(c.Body(), &chName)) {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		chName.ID = stoi[int64](c.Params("id"))

		if opt.Get(q.ChangeName(ctx, chName)).Ok() {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	})

	app.Delete("/todo/id/:id", func(c *fiber.Ctx) error {
		id := stoi[int64](c.Params("id"))

		if opt.Get(q.DeleteTodo(ctx, id)).Ok() {
			return c.SendStatus(fiber.StatusOK)
		}
		return c.SendStatus(fiber.StatusInternalServerError)
	})
	//lui etubo aki
	opt.Fatal(app.Listen(":3000"))
}
