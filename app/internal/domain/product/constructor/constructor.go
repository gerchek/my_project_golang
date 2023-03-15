package constructor

import (
	"my_project/internal/domain/product/controller"
	"my_project/internal/domain/product/service"
	"my_project/internal/domain/product/storage"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	ProductRepository storage.ProductStorage
	ProductService    service.ProductService
	ProductController controller.ProductController
)

func ProductRequirementsCreator(client *gorm.DB, redis *redis.Client) {
	ProductRepository = storage.NewProductStorage(client)
	ProductService = service.NewProductService(ProductRepository, redis)
	ProductController = controller.NewProductController(ProductService)
}
