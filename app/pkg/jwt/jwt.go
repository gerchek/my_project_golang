package jwt

import (
	"fmt"
	"my_project/internal/model"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type JWTAdminService interface {
	// GenerateAdminToken(userID, username string) (*model.AdminTokenDetails, error)
	ValidateAdminAccessToken(token string) (*jwt.Token, error)
	ValidateAdminRefreshToken(token string) (*jwt.Token, error)
	CreateToken(userid string, username string) (*model.TokenDetails, error)
}

type jwtAdminService struct {
	// issuer                string
	secretKeyAccessToken  string
	secretKeyRefreshToken string
}

func NewJWTAdminService() JWTAdminService {
	return &jwtAdminService{
		// issuer:                "e.gov.tm/api/portal/panel",
		secretKeyAccessToken:  "my_project_access_secret",
		secretKeyRefreshToken: "my_project_access_secret",
	}
}

func (j *jwtAdminService) CreateToken(userid string, username string) (*model.TokenDetails, error) {
	td := &model.TokenDetails{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	// td.AccessUuid = uuid.NewV4().String()
	access_uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	td.AccessUuid = access_uuid.String()

	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	// td.RefreshUuid = uuid.NewV4().String()
	refresh_uuid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}
	td.RefreshUuid = refresh_uuid.String()
	// fmt.Println(userid)

	// var err error
	//Creating Access Token
	os.Setenv("ACCESS_SECRET", "my_project_access_secret") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["user_id"] = userid
	atClaims["username"] = username
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return nil, err
	}
	//Creating Refresh Token
	os.Setenv("REFRESH_SECRET", "my_project_access_secret") //this should be in an env file
	rtClaims := jwt.MapClaims{}
	rtClaims["refresh_uuid"] = td.RefreshUuid
	rtClaims["user_id"] = userid
	rtClaims["username"] = username
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("REFRESH_SECRET")))
	if err != nil {
		return nil, err
	}
	return td, nil

}

func (j *jwtAdminService) ValidateAdminAccessToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKeyAccessToken), nil
	})
}

func (j *jwtAdminService) ValidateAdminRefreshToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t_ *jwt.Token) (interface{}, error) {
		if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
		}
		return []byte(j.secretKeyRefreshToken), nil
	})
}
