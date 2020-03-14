package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	firebase "firebase.google.com/go"
	"github.com/devrodriguez/first-class-api-go/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var checks []models.Check

func GetChecks(res http.ResponseWriter, req *http.Request) {
	var check models.Check

	ctx := context.Background()
	sa := option.WithCredentialsFile("./starqsoft-taskapp-firebase-adminsdk-yf7i4-5c6d8c909a.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Println(err)
		err := models.Error{Description: err.Error()}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(err)
		return
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Println(err)
		err := models.Error{Description: err.Error()}
		res.Header().Set("Content-Type", "application/json")
		json.NewEncoder(res).Encode(err)
		return
	}

	docsRef := client.Collection("checks").Documents(ctx)

	for {
		doc, err := docsRef.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println(err)
			errRes := models.Error{Description: err.Error()}
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(errRes)
		}

		parseError := doc.DataTo(&check)

		if parseError != nil {
			log.Println(parseError)
			errRes := models.Error{Description: parseError.Error()}
			res.Header().Set("Content-Type", "application/json")
			json.NewEncoder(res).Encode(errRes)
			break
		}

		checks = append(checks, check)
	}

	// Build json response
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(checks)
}

func CreateCheck(res http.ResponseWriter, req *http.Request) {
	var check models.Check
	json.NewDecoder(req.Body).Decode(&check)
	checks := append(checks, check)

	json.NewEncoder(res).Encode(checks)
}
