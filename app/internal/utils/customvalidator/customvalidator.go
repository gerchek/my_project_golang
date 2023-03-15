package customvalidator

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateStruct(data interface{}) error {
	validate := validator.New()
	err := validate.Struct(data)
	return err
}

func CheckExtension(ext string) bool {
	extensions := []string{"png", "svg", "mp4", "mkv", "pdf"}
	for _, v := range extensions {
		if v == ext {
			return true
		}
	}
	return false
}

func ValidatePhone(phone string) bool {
	r := regexp.MustCompile("[+9936]{5}[1-5][0-9]{6}$")
	return r.MatchString(phone)
}
