package rest

import (
	"log"
	"net/http"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

type authHandler struct{}

func NewAuthHandler() *authHandler {
	return &authHandler{}
}

func (ah *authHandler) SignIn(c *gin.Context) {
	var resModel APIResponse

	// == VALIDATE USER AND PASSWORD ==
	// data, _ := getBodyData(gCtx)
	user := c.Query("user")
	password := c.Query("password")

	if !validateUserAuth(user, password) {
		resModel.Message = "Wrong user or password"

		c.JSON(http.StatusOK, resModel)
		return
	}

	// == CREATE JWT TOKEN ==
	token, err := CreateToken(c.Request)

	log.Println(string(token))

	if err != nil {
		resModel.Message = "Autenticacion fallida"
		resModel.Errors = []APIError{
			{
				Status:      http.StatusUnauthorized,
				Title:       http.StatusText(http.StatusUnauthorized),
				Description: err.Error(),
			},
		}

		c.JSON(http.StatusOK, resModel)
		return
	}

	resModel.Data = gin.H{"token": string(token)}
	c.JSON(http.StatusOK, resModel)
}

// Login valida el token de Authorization
func (ah *authHandler) Login(gCtx *gin.Context) {
	var resModel APIResponse
	req := gCtx.Request

	err := VerifyToken(req)

	if err != nil {
		resModel.Message = "usuario no autorizado"
		resModel.Errors = []APIError{
			{
				Status:      http.StatusUnauthorized,
				Title:       http.StatusText(http.StatusUnauthorized),
				Description: err.Error(),
			},
		}

		gCtx.JSON(http.StatusOK, resModel)
		return
	}

	resModel.Message = "success authentication"
	gCtx.JSON(http.StatusOK, resModel)
}

func CreateToken(r *http.Request) (string, error) {
	var hs = jwt.NewHS256([]byte("dev1986"))
	now := time.Now()

	payload := JwtPayload{
		Payload: jwt.Payload{
			Issuer:         "devrodriguez",
			Subject:        "dev",
			Audience:       jwt.Audience{"http://localhost:3000"},
			ExpirationTime: jwt.NumericDate(now.Add(3000 * time.Second)),
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
	var payload JwtPayload
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

func validateUserAuth(user, password string) bool {
	if user == "john" && password == "12345" {
		return true
	}

	return false
}
