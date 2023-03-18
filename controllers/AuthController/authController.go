package authcontroller

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	database "ohas-api.com/v2/databases"
	"ohas-api.com/v2/helpers"
	"ohas-api.com/v2/models"
)

type ProfileJSON struct {
	ID       uint   `json:"-"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func Login(c echo.Context) error {
	// var userInput models.User
	// decoder := json.NewDecoder(r.Body)

	// if err := decoder.Decode(&userInput); err != nil {
	// 	helpers.ResponseFailure(w, "Gagal mengambil parameter")
	// 	return
	// }

	// defer r.Body.Close()

	var user models.User

	if err := database.Instance.Where("username = ?", c.FormValue("username")).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return helpers.ResponseFailure(c, "Username atau password salah")
		default:
			return helpers.ResponseServerError(c, err.Error())
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.FormValue("password"))); err != nil {
		return helpers.ResponseFailure(c, "Password tidak cocok")
	}

	//making jwt
	expired := time.Now().Add(time.Hour * 300)
	claims := &helpers.JWTClaim{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expired),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenAlgo.SignedString(helpers.JWTKey)
	if err != nil {
		return helpers.ResponseFailure(c, err.Error())
	}

	response := map[string]string{
		"token":        token,
		"expired_time": expired.String(),
		"username":     claims.Username,
	}
	return helpers.ResponseSuccess(c, response)
}

func Register(c echo.Context) error {
	inputUser := models.User{
		UniqueID: uuid.New().String(),
		Name:     c.FormValue("firstName") + " " + c.FormValue("lastName"),
		Username: strings.ToLower(c.FormValue("firstName") + c.FormValue("lastName")),
		Password: c.FormValue("password"),
		Email:    c.FormValue("email"),
		Phone:    c.FormValue("phone"),
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(inputUser.Password), bcrypt.DefaultCost)
	inputUser.Password = string(hashPass)

	if err := database.Instance.Create(&inputUser).Error; err != nil {
		log.Fatal("Gagal Simpan Data")
		return helpers.ResponseFailure(c, "Gagal Simpan Data")
	}

	return helpers.ResponseCreated(c, "User")
}

func Logout(c echo.Context) error {
	http.SetCookie(c.Response(), &http.Cookie{
		Name:     "token",
		Path:     "/",
		Value:    "",
		HttpOnly: true,
		MaxAge:   -1,
	})

	return helpers.ResponseMessage(c, true, http.StatusOK, "Successfully Logged out")
}

func AuthUser(r *http.Request) (response *ProfileJSON) {

	var user models.User

	var profile ProfileJSON

	c := strings.Split(r.Header.Get("authorization"), "Bearer ")

	if len(c) != 2 {
		return nil
	}

	tokenString := c[1]
	claims := &helpers.JWTClaim{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return helpers.JWTKey, nil
	})

	if err != nil {
		return nil
	}

	if !token.Valid {
		return nil
	}
	database.Instance.Where("username = ?", claims.Username).Model(&user).Find(&profile)
	return &profile
}
