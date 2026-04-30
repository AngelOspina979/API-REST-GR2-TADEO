package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"todo-api/internal/auth"
	"todo-api/internal/handlers"
	"todo-api/internal/middleware"
	"todo-api/internal/repository"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// ✅ CREAR CARPETA DB PRIMERO
	if err := os.MkdirAll("db", os.ModePerm); err != nil {
		log.Fatalf("Error creando carpeta db: %v", err)
	}

	// ✅ INICIALIZAR BASE DE DATOS UNA SOLA VEZ
	repo, err := repository.NewSQLiteRepository("db/todo.db")
	if err != nil {
		log.Fatalf("Error al inicializar la base de datos: %v", err)
	}
	defer repo.Close()

	jwtService := auth.NewJWTService("mi-secreto-super-seguro-cambiar-en-produccion")

	taskHandler := handlers.NewTaskHandler(repo)
	authHandler := handlers.NewAuthHandler(repo, jwtService)

	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/auth/login", authHandler.Login)
	mux.HandleFunc("GET /api/health", healthCheck)

	protected := middleware.Chain(
		middleware.Logger,
		middleware.Auth(jwtService),
	)

	mux.Handle("GET /api/tasks", protected(http.HandlerFunc(taskHandler.GetAll)))
	mux.Handle("POST /api/tasks", protected(http.HandlerFunc(taskHandler.Create)))
	mux.Handle("GET /api/tasks/{id}", protected(http.HandlerFunc(taskHandler.GetByID)))
	mux.Handle("PUT /api/tasks/{id}", protected(http.HandlerFunc(taskHandler.Update)))
	mux.Handle("DELETE /api/tasks/{id}", protected(http.HandlerFunc(taskHandler.Delete)))
	mux.Handle("PATCH /api/tasks/{id}/complete", protected(http.HandlerFunc(taskHandler.MarkComplete)))

	fmt.Printf("Servidor corriendo en http://localhost:%s\n", port)

	if err := http.ListenAndServe(":"+port, mux); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, `{"status":"ok"}`)
}
