package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiConfig) handlerUsersRetrieve(w http.ResponseWriter, r *http.Request) {
	dbUsers, err := cfg.DB.GetUsers()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve users")
		return
	}

	users := []User{}

	for _, dbUser := range dbUsers {
		users = append(users, User{
			ID:          dbUser.ID,
			Email:       dbUser.Email,
			IsChirpyRed: dbUser.IsChirpyRed,
		})
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].ID < users[j].ID
	})

	respondWithJSON(w, http.StatusOK, users)
}

func (cfg *apiConfig) handlerUsersGet(w http.ResponseWriter, r *http.Request) {
	userIDString := chi.URLParam(r, "userID")
	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	dbUser, err := cfg.DB.GetUser(userID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get user")
		return
	}

	respondWithJSON(w, http.StatusOK, User{
		ID:          dbUser.ID,
		Email:       dbUser.Email,
		IsChirpyRed: dbUser.IsChirpyRed,
	})
}
