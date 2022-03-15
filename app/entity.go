package app

import (
	userDelivery "be-golang-fiber/entity/user/delivery"
	userRepo "be-golang-fiber/entity/user/repository"
	userUsecase "be-golang-fiber/entity/user/usecase"

	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func InitEntity(f *fiber.App, dbConn *sql.DB) {
	userRepo := userRepo.NewMySQLUserRepository(dbConn)
	userUsecase := userUsecase.NewUserUseCase(userRepo)
	userDelivery.NewHttpDelivery(f, userUsecase)
}
