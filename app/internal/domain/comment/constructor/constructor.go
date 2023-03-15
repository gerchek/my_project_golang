package constructor

import (
	"my_project/internal/domain/comment/controller"
	"my_project/internal/domain/comment/service"
	"my_project/internal/domain/comment/storage"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	CommentRepository storage.CommentStorage
	CommentService    service.CommentService
	CommentController controller.CommentController
)

func CommentRequirementsCreator(client *gorm.DB, redis *redis.Client) {
	CommentRepository = storage.NewCommentStorage(client)
	CommentService = service.NewCommentService(CommentRepository, redis)
	CommentController = controller.NewCommentController(CommentService)
}
