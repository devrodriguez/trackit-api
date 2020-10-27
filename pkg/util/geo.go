package util

import (
	"encoding/json"
	"net/http"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
)

func GetAddresPredictions(term string) (entity.GoogleApi, error) {
	// TO DO: set on environment
	var key = "AIzaSyA9UqT-ykyMUf2MJbc7EBsIQj6D69DHSJo"
	var apiRes entity.GoogleApi

	// TO DO: set URL on environment
	res, err := http.Get("https://maps.googleapis.com/maps/api/place/autocomplete/json?input=" + term + "&key=" + key)

	if err != nil {
		return entity.GoogleApi{}, err
	}

	json.NewDecoder(res.Body).Decode(&apiRes)

	return apiRes, nil
}
