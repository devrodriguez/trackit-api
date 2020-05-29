package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/devrodriguez/first-class-api-go/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func GetCompanies(gCtx *gin.Context) {
	var resModel models.Response
	var company models.Company
	var companies []models.Company
	ctx := context.Background()
	client, err := getFirestoreClient(gCtx)

	if err != nil {
		log.Println(err)
		resModel.Error = err.Error()
		gCtx.JSON(http.StatusOK, resModel)
		return
	}

	defer client.Close()

	docsRef := client.Collection("companies").Documents(ctx)
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

		if err := doc.DataTo(&company); err != nil {
			log.Println(err)
			resModel.Message = "Fail on set data"
			resModel.Error = err.Error()
			gCtx.JSON(http.StatusOK, resModel)
			break
		}

		companies = append(companies, company)
	}

	resModel.Data = gin.H{"companies": companies}

	gCtx.JSON(http.StatusOK, resModel)
}
