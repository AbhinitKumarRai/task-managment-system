package routes

import (
	"net/http"

	"github.com/AbhinitKumarRai/user-service/internal/handler"
	"github.com/AbhinitKumarRai/user-service/internal/service"
	"github.com/AbhinitKumarRai/user-service/pkg/grpcclient"
	"github.com/AbhinitKumarRai/user-service/pkg/middleware"
	"github.com/gorilla/mux"
)

func RegisterRoutes(userService *service.UserService, taskClient *grpcclient.TaskGRPCClient) *mux.Router {
	router := mux.NewRouter()

	// --- Health Check ---
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	}).Methods("GET")

	// --- User Handlers ---
	userHandler := handler.NewUserHandler(userService, taskClient)

	// Public User Routes
	router.HandleFunc("/user/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/user/login", userHandler.Login).Methods("POST")

	// Protected User Routes
	protectedUser := router.PathPrefix("/user").Subrouter()
	protectedUser.Use(middleware.AuthMiddleware)
	protectedUser.HandleFunc("", userHandler.GetByID).Methods("GET")
	protectedUser.HandleFunc("/update", userHandler.Update).Methods("POST")
	protectedUser.HandleFunc("", userHandler.Delete).Methods("DELETE")

	// --- Task Handlers ---
	taskHandler := handler.NewTaskHandler(taskClient)

	// Protected Task Routes
	protectedTask := router.PathPrefix("/task").Subrouter()
	protectedTask.Use(middleware.AuthMiddleware)
	protectedTask.HandleFunc("", taskHandler.List).Methods("GET")      // /task?status=...
	protectedTask.HandleFunc("", taskHandler.Create).Methods("POST")   // /task
	protectedTask.HandleFunc("", taskHandler.GetByID).Methods("GET")   // /task?id=...
	protectedTask.HandleFunc("", taskHandler.Update).Methods("PUT")    // /task?id=...
	protectedTask.HandleFunc("", taskHandler.Delete).Methods("DELETE") // /task?id=...

	return router
}
