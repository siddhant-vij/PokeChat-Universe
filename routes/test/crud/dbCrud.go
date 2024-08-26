package crud

import (
	"context"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/database"
)

var testUserId = uuid.New()

func DbCreateHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	testUser := database.InsertUserParams{
		ID:    testUserId,
		Email: "test@email.com",
	}

	user, err := cfg.DBQueries.InsertUser(context.Background(), testUser)
	if err != nil {
		log.Fatalf("error inserting user. Err: %v", err)
	}

	w.Write([]byte(user.Email + " is inserted in the database!"))
}

func DbReadHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	user, err := cfg.DBQueries.GetUserByID(context.Background(), testUserId)
	if err != nil {
		log.Fatalf("error getting user. Err: %v", err)
	}

	w.Write([]byte(user.Email))
}

func DbUpdateHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	emailUpdateParams := database.UpdateUserEmailByIDParams{
		ID:    testUserId,
		Email: "test@gmail.com",
	}

	err := cfg.DBQueries.UpdateUserEmailByID(context.Background(), emailUpdateParams)
	if err != nil {
		log.Fatalf("error updating user. Err: %v", err)
	}

	w.Write([]byte(emailUpdateParams.Email + " is updated in the database!"))
}

func DbDeleteHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	err := cfg.DBQueries.DeleteUserByID(context.Background(), testUserId)
	if err != nil {
		log.Fatalf("error deleting user. Err: %v", err)
	}

	w.Write([]byte("user deleted from the database!"))
}
