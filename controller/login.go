package controller 
import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"example.com/m/v2/database"
	"example.com/m/v2/model"
	"time"
)

func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	var input LoginInput

	if err := c.BodyParser(&input); err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}
	username := input.Username
	
	database := database.DB
	var user model.User
	database.Where("username = ?" , username).First(&user)

	if !CheckPasswordHash(input.Password, user.Password) {
		return c.Status(404).JSON(fiber.Map{"status" : "error" , "message": "Wrong password"})
	}
	if user.Username == ""{
		return c.Status(404).JSON(fiber.Map{"status" : "error" , "message": "Can not found user"})
	}
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"data" : user , "token": t})
}
// Decode token and send it as response
func Restricted(c *fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	name := claims["name"].(string)
	return c.SendString("Welcome " + name)
}