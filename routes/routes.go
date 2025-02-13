package routes

import (
	"computerextra/elaerning-go/db"
	"computerextra/elaerning-go/handler"
	"computerextra/elaerning-go/templates"
	"fmt"
	"net/http"

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

	// Session Routes
	router.HandleFunc("/login", loginHandler).Methods(http.MethodGet, http.MethodPost)
	router.Handle("/dashboard", authMiddleware(http.HandlerFunc(homeHandler))).Methods(http.MethodGet)

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

func GetApiRoutes(router *mux.Router, client *db.PrismaClient) {}
