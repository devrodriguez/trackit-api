package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/devrodriguez/first-class-api-go/models"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

func GetChecks(gCtx *gin.Context) {
	var resModel models.Response
	var checks []models.Check
	var check models.Check

	ctx := context.Background()

	// Firestore client
	client, err := getFirestoreClient(gCtx)

	if err != nil {
		log.Println(err)
		resModel.Error = err.Error()
		gCtx.JSON(http.StatusOK, resModel)
		return
	}

	defer client.Close()

	// Doccument reference
	docsRef := client.Collection("checks").Documents(ctx)
	defer docsRef.Stop()

	for {
		doc, err := docsRef.Next()

		if err == iterator.Done {
			break
		}

		if err != nil {
			log.Println(err)
			resModel.Message = "Fail on read data"
			resModel.Error = err.Error()
			gCtx.JSON(http.StatusOK, resModel)
		}

		if err := doc.DataTo(&check); err != nil {
			log.Println(err)
			resModel.Message = "Fail on set data"
			resModel.Error = err.Error()
			gCtx.JSON(http.StatusOK, resModel)
			break
		}

		checks = append(checks, check)
	}

	resModel.Message = "Success"
	resModel.Data = gin.H{"checks": checks}

	// Build json response
	gCtx.JSON(http.StatusOK, resModel)
}

func CreateCheck(gCtx *gin.Context) {
	var resModel models.Response
	var check models.Check
	ctx := context.Background()

	// Get data from request
	if err := gCtx.BindJSON(&check); err != nil {
		resModel.Error = err.Error()
		resModel.Message = "Fail binding JSON data"
		gCtx.JSON(http.StatusInternalServerError, resModel)
		return
	}

	// Firestore client
	client, err := getFirestoreClient(gCtx)

	if err != nil {
		log.Println(err)
		resModel.Error = err.Error()
		gCtx.JSON(http.StatusInternalServerError, resModel)
		return
	}

	defer client.Close()

	// Add document
	doc, _, errAdd := client.Collection("checks").Add(ctx, check)

	if errAdd != nil {
		log.Println(errAdd)
		resModel.Error = errAdd.Error()
		resModel.Message = "Error saving data"
		gCtx.JSON(http.StatusInternalServerError, resModel)
		return
	}

	resModel.Message = "Document created"
	resModel.Data = gin.H{"docId": doc.ID}

	gCtx.JSON(http.StatusOK, resModel)
}

func UpdateCheck(gCtx *gin.Context) {
	var resModel models.Response

	data, _ := getBodyData(gCtx)
	ctx := context.Background()

	// Firestore client
	client, err := getFirestoreClient(gCtx)

	if err != nil {
		log.Println(err)
		resModel.Error = err.Error()
		gCtx.JSON(http.StatusInternalServerError, resModel)
		return
	}

	defer client.Close()
	// Doccument reference
	docRef := client.Doc("checks/" + data["id"])

	_, errUpdate := docRef.Update(ctx, []firestore.Update{{Path: "Hour", Value: data["hour"]}})

	if errUpdate != nil {
		log.Println(errUpdate)
		resModel.Error = errUpdate.Error()
		resModel.Message = "Error updating data"
		gCtx.JSON(http.StatusInternalServerError, resModel)
		return
	}

	resModel.Message = "Document updated"

	gCtx.JSON(http.StatusOK, resModel)
}

func getFirestoreClient(gCtx *gin.Context) (*firestore.Client, error) {
	ctx := context.Background()
	sa := option.WithCredentialsFile("./starqsoft-taskapp-firebase-adminsdk-yf7i4-5c6d8c909a.json")
	app, err := firebase.NewApp(ctx, nil, sa)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	return client, nil
}

func getBodyData(gCtx *gin.Context) (map[string]string, error) {
	var raw map[string]string

	body, err := ioutil.ReadAll(gCtx.Request.Body)

	if err != nil {
		return nil, err
	}

	bodyStr := []byte(string(body))

	if err := json.Unmarshal(bodyStr, &raw); err != nil {
		return nil, err
	}

	return raw, nil
}
