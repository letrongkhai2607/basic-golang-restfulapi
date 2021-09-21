package controller 
import(
	"example.com/m/v2/model"
	"example.com/m/v2/database"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"strconv"
)
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
func CheckPasswordHash(password, passwordHashed string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(password)) 
	return err == nil
}

func CheckUserExists(id string) bool {
	database := database.DB
	var user model.User
	database.First(&user, id)
	if user.Username == "" {
		return false
	}
	return true
}

func GetAllUsers(c *fiber.Ctx) error {
	limitQuery := c.Query("limit")
	pageQuery := c.Query("page")
	
	database := database.DB
	var users []model.User

	if limitQuery == "" || pageQuery == "" {
		database.Find(&users)
		return c.JSON(fiber.Map{"status": "success" , "data": users})
	}
	limit , _ := strconv.Atoi(limitQuery)
	page , _ := strconv.Atoi(pageQuery)
	offset := (page - 1) * limit
	database.Limit(limit).Offset(offset).Find(&users)
	return c.JSON(fiber.Map{"status": "success" , "data": users , "limit": limit , "page": page})
}

func GetUserById(c *fiber.Ctx) error {
	userId := c.Params("id")
	users := new(model.User)

	database := database.DB
	database.First(users , userId)
	
	if users.Username == "" {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No user found with ID", "data": nil})
	}
	return c.JSON(fiber.Map{"status": "success" , "data": users})
}

func CreateNewUser(c *fiber.Ctx) error {
	// NewUser is esponse result when success 
	type NewUser struct {
		Username string `json:"username"`
		Password    string `json:"password"`
	}
	database := database.DB
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	

	hash, err := HashPassword(user.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't hash password", "data": err})
	}

	user.Password = hash

	if err := database.Create(&user).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Couldn't create user", "data": err})
	}

	response := NewUser{
		Username: user.Username,
		Password: user.Password,
	}
	return c.JSON(fiber.Map{"status": "success", "message": "Created user", "data": response})
}


func DeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	// Check if user exists , if exists then delete user
	if !CheckUserExists(userId) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})
	}

	database := database.DB
	var users model.User

	database.First(&users, userId) // Find the fisrt user with that userId

	database.Delete(&users)
	return c.JSON(fiber.Map{"status": "success", "message": "User successfully deleted", "data": nil})
}

func UpdateUser(c *fiber.Ctx) error {
	// a new struct for new user
	type UpdateUserInput struct {
		Names string `json:"name"`
	}
	var updateUserInput UpdateUserInput
	if err := c.BodyParser(&updateUserInput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}

	userId := c.Params("id")
	// Check if user exists , if exists then delete user
	if !CheckUserExists(userId) {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Not valid user", "data": nil})
	}

	database := database.DB
	var users model.User
	database.First(&users , userId)
	users.Names = updateUserInput.Names
	database.Save(&users)

	return c.JSON(fiber.Map{"status": "success", "message": "User successfully updated", "data": users}) 
}
