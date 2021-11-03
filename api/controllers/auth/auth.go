package auth

import (
	"net/http"
	"rumah_resep/models"

	echo "github.com/labstack/echo/v4"
)

// ------------------------------------------------------------------
// Start Request
// ------------------------------------------------------------------

type RegisterUserRequest struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Gender   string `json:"gender" form:"gender"`
	Address  string `json:"address" form:"address"`
	Role     string `json:"role" form:"role"`
}

type LoginUserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// ------------------------------------------------------------------
// End Request
// ------------------------------------------------------------------

type AuthController struct {
	userModel models.UserModel
}

func NewAuthController(userModel models.UserModel) *AuthController {
	return &AuthController{
		userModel,
	}
}

func (controller *AuthController) RegisterUserController(c echo.Context) error {
	var userRequest RegisterUserRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	user := models.User{
		Name:     userRequest.Name,
		Email:    userRequest.Email,
		Gender:   userRequest.Gender,
		Password: userRequest.Password,
		Address:  userRequest.Address,
		Role:     userRequest.Role,
	}

	_, err := controller.userModel.Register(user)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"success": false,
			"code":    500,
			"message": "Internal Server Error",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Register Account",
	})
}

func (controller *AuthController) LoginUserController(c echo.Context) error {
	var userRequest LoginUserRequest

	if err := c.Bind(&userRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	user, err := controller.userModel.Login(userRequest.Email, userRequest.Password)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"success": false,
			"code":    400,
			"message": "Bad Request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"success": true,
		"code":    200,
		"message": "Success Login",
		"token":   user.Token,
	})
}
