package routes

import (
	"computerextra/elaerning-go/db"
	"computerextra/elaerning-go/handler"
	"computerextra/elaerning-go/templates"

	"github.com/gorilla/mux"
)

func GetRoutes(router *mux.Router, client *db.PrismaClient) {
	// ctx := context.Background()
	// create, err := client.User.CreateOne(
	// 	db.User.Password.Set("Test"),
	// 	db.User.ID.Set(cuid.New()),
	// 	db.User.Name.Set("Test User"),
	// ).Exec(ctx)
	// if err != nil {
	// 	panic(err)
	// }
	// result, _ := json.MarshalIndent(create, "", "  ")
	// fmt.Printf("created post: %s\n", result)

	// Main Route
	router.Handle("/", handler.Component(templates.Index()))

}

func GetApiRoutes(router *mux.Router, client *db.PrismaClient) {}
