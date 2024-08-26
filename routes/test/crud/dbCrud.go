package crud

import (
	"context"
	"database/sql"
	"log"
	"net/http"

	"github.com/google/uuid"

	"github.com/siddhant-vij/PokeChat-Universe/config"
	"github.com/siddhant-vij/PokeChat-Universe/database"
)

var (
	testUserId   = uuid.New()
	testAuthID   = "testAuthID"
	testUsername = "testUser"
	testEmail    = "test@email.com"
	testPicUrl   = "testPictureUrl"
)

func DbCreateHandler(w http.ResponseWriter, r *http.Request, cfg *config.AppConfig) {
	testUser := database.InsertUserParams{
		ID:       testUserId,
		AuthID:   testAuthID,
		Username: testUsername,
		Email:    testEmail,
		PictureUrl: sql.NullString{
			String: testPicUrl,
			Valid:  true,
		},
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
