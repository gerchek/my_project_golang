package constructor

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func SetConstructor(client *gorm.DB, logger *logrus.Logger) {
	fmt.Println("SetConstructor")
}
