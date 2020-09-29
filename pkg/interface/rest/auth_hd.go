package rest

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/service"
	"github.com/gbrlsnchs/jwt/v3"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	empSrv service.EmployeeService
}

func NewAuthHandler(empSrv service.EmployeeService) *authHandler {
	return &authHandler{
		empSrv,
	}
}

func (ah *authHandler) SignIn(c *gin.Context) {
	var resModel APIResponse

	// == VALIDATE USER AND PASSWORD ==
	user := c.Query("user")
	password := c.Query("password")

	if !validateUserAuth(ah, user, password) {
		resModel.Message = "Wrong user or password"

		c.JSON(http.StatusOK, resModel)
		return
	}

	// == CREATE JWT TOKEN ==
	token, err := CreateToken(c.Request)

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

	token := extractToken(r)

	expValidator := jwt.ExpirationTimeValidator(now)
	nbfValidator := jwt.NotBeforeValidator(now)
	validatePayload := jwt.ValidatePayload(&payload.Payload, expValidator, nbfValidator)

	_, err := jwt.Verify([]byte(token), secret, &payload, validatePayload)

	if err != nil {
		return err
	}

	return nil
}

func extractToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if bearToken := strings.Split(token, " "); len(bearToken) == 2 {
		return bearToken[1]
	}

	return ""
}

func validateUserAuth(ah *authHandler, email, password string) bool {
	valid, err := ah.empSrv.ValidateCredentials(email, password)

	if err != nil {
		return false
	}

	return valid
}
