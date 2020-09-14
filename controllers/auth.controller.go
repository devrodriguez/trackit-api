package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/devrodriguez/first-class-api-go/models"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

// SignIn retorna un token de autenticacion
func SignIn(gCtx *gin.Context) {
	var resModel models.Response

	// == VALIDATE USER AND PASSWORD ==
	// data, _ := getBodyData(gCtx)
	user := gCtx.Query("user")
	password := gCtx.Query("password")

	if !ValidateUserAuth(user, password) {
		resModel.Message = "Wrong user or password"

		gCtx.JSON(http.StatusOK, resModel)
		return
	}

	// == CREATE JWT TOKEN ==
	token, err := CreateToken(gCtx.Request)

	log.Println(string(token))

	if err != nil {
		resModel.Message = "Autenticacion fallida"
		resModel.Error = err.Error()

		gCtx.JSON(http.StatusOK, resModel)
		return
	}

	resModel.Data = gin.H{"token": string(token)}
	gCtx.JSON(http.StatusOK, resModel)
}

// Login valida el token de Authorization
func Login(gCtx *gin.Context) {
	var resModel models.Response
	req := gCtx.Request

	err := VerifyToken(req)

	if err != nil {
		resModel.Message = err.Error() + " | ¡Usuario no autorizado!"

		gCtx.JSON(http.StatusOK, resModel)
		return
	}

	resModel.Message = "Welcome"
	gCtx.JSON(http.StatusOK, resModel)
}

func CreateToken(r *http.Request) (string, error) {
	var hs = jwt.NewHS256([]byte("dev1986"))
	now := time.Now()

	payload := models.JwtPayload{
		Payload: jwt.Payload{
			Issuer:         "devrodriguez",
			Subject:        "dev",
			Audience:       jwt.Audience{"http://localhost:3000"},
			ExpirationTime: jwt.NumericDate(now.Add(300 * time.Second)),
			NotBefore:      jwt.NumericDate(now),
			IssuedAt:       jwt.NumericDate(now),
			JWTID:          "foobar",
		},
	}

	token, err := jwt.Sign(payload, hs)

	log.Println(err)

	if err != nil {
		return "", err
	}

	return string(token), nil
}

func VerifyToken(r *http.Request) error {
	var secret = jwt.NewHS256([]byte("dev1986"))
	var payload models.JwtPayload
	now := time.Now()

	token := []byte(r.Header.Get("Authorization"))

	expValidator := jwt.ExpirationTimeValidator(now)
	nbfValidator := jwt.NotBeforeValidator(now)
	validatePayload := jwt.ValidatePayload(&payload.Payload, expValidator, nbfValidator)

	hd, err := jwt.Verify(token, secret, &payload, validatePayload)

	log.Println(hd)

	if err != nil {
		return err
	}

	return nil
}

func ValidateUserAuth(user, password string) bool {
	log.Println(user, password)
	if user == "john" && password == "12345" {
		return true
	}

	return false
}
