package delivery

import (
	"be-golang-fiber/entity/user"
	"be-golang-fiber/entity/user/usecase"
	"be-golang-fiber/utils"
	"be-golang-fiber/utils/config_variable"
	"be-golang-fiber/utils/jwt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

type UserHttpDelivery struct {
	UserUsecase usecase.UserUsecase
}

func NewHttpDelivery(f *fiber.App, u usecase.UserUsecase) {
	handler := &UserHttpDelivery{
		UserUsecase: u,
	}

	api := f.Group("api/v1/user")
	api.Post("/register", handler.UserRegister)
	api.Post("/login", handler.UserLogin)
	// JWT Middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(config_variable.Secret),
	}))
	api.Get("/profile", handler.GetProfile)

}

func (u *UserHttpDelivery) UserRegister(c *fiber.Ctx) error {
	resp := new(user.ResponseModel)
	userObj := new(user.UserModel)
	err := c.BodyParser(userObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}
	res, err := u.UserUsecase.UserRegister(userObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.REGISTER_SUCCESS
	resp.Data = res

	return c.Status(http.StatusOK).JSON(resp)
}

func (u *UserHttpDelivery) UserLogin(c *fiber.Ctx) error {
	userObj := new(user.UserLoginModel)
	err := c.BodyParser(userObj)
	resp := new(user.ResponseModel)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}
	token, err := u.UserUsecase.UserLogin(userObj)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.LOGIN_SUCCESS
	resp.Data = token
	return c.Status(http.StatusOK).JSON(resp)
}

func (u *UserHttpDelivery) GetProfile(c *fiber.Ctx) error {
	resp := new(user.ResponseModel)
	reqToken := c.Get(fiber.HeaderAuthorization)
	id, err := jwt.GetIDfromToken(reqToken)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}

	art, err := u.UserUsecase.GetProfile(id)
	if err != nil {
		resp.Status = utils.FAILED
		resp.Message = err.Error()
		return c.Status(http.StatusInternalServerError).JSON(resp)
	}
	resp.Status = utils.SUCCESS
	resp.Message = utils.OK
	resp.Data = art
	return c.Status(http.StatusOK).JSON(resp)
}
