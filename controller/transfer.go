package controller 
import(
	"example.com/m/v2/model"
	"example.com/m/v2/database"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func Transfer(c *fiber.Ctx) error {
	database := database.DB

	var to_user model.User
	var from_user model.User
	var transferInput model.Transfer

	if err := c.BodyParser(&transferInput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	if transferInput.To_User_ID == transferInput.From_User_ID{
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": nil})
	}

	
	database.First(&from_user , transferInput.From_User_ID)
	if transferInput.Amount > from_user.Balance{
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Insufficient balance", "data": nil})
	}
	from_user.Balance = from_user.Balance - transferInput.Amount
	database.Save(&from_user)
	

	database.First(&to_user , transferInput.To_User_ID)
	to_user.Balance = to_user.Balance + transferInput.Amount
	database.Save(&to_user)

	database.Create(&transferInput)

	return c.JSON(fiber.Map{"status": "success" , "message": "Transaction is success"})
}

func Recharge(c *fiber.Ctx) error {

	database := database.DB

	var rechargeInput model.Recharge
	var user model.User
	if err := c.BodyParser(&rechargeInput); err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	
	database.First(&user , rechargeInput.To_User_ID)
	user.Balance = user.Balance + rechargeInput.Amount
	database.Save(&user)
	database.Create(&rechargeInput)
	return c.JSON(fiber.Map{"status": "success" , "message": "Recharge is success" , "data" : user})
}

func GetAllTransferHistory(c *fiber.Ctx) error {
	limitQuery := c.Query("limit")
	pageQuery := c.Query("page")
	
	database := database.DB
	var transfers []model.Transfer

	if limitQuery == "" || pageQuery == "" {
		database.Find(&transfers)
		return c.JSON(fiber.Map{"status": "success" , "data": transfers})
	}
	limit , _ := strconv.Atoi(limitQuery)
	page , _ := strconv.Atoi(pageQuery)
	offset := (page - 1) * limit
	database.Limit(limit).Offset(offset).Find(&transfers)
	return c.JSON(fiber.Map{"status": "success" , "data": transfers , "limit": limit , "page": page})
}

func GetAllRechargeHistory(c *fiber.Ctx) error {
	limitQuery := c.Query("limit")
	pageQuery := c.Query("page")
	
	database := database.DB
	var recharges []model.Recharge

	if limitQuery == "" || pageQuery == "" {
		database.Find(&recharges)
		return c.JSON(fiber.Map{"status": "success" , "data": recharges})
	}
	limit , _ := strconv.Atoi(limitQuery)
	page , _ := strconv.Atoi(pageQuery)
	offset := (page - 1) * limit
	database.Limit(limit).Offset(offset).Find(&recharges)
	return c.JSON(fiber.Map{"status": "success" , "data": recharges , "limit": limit , "page": page})
}


func GetRechargeHistoryByUserId(c *fiber.Ctx) error {
	userId := c.Params("id")
	var recharges []model.Recharge

	database := database.DB
	database.Where("to_user_id = ?", userId).Find(&recharges)

	return c.JSON(fiber.Map{"status": "success" , "data": recharges})
}

func GetTransferHistoryByUserId(c *fiber.Ctx) error {
	userId := c.Params("id")
	var transfers []model.Transfer

	database := database.DB
	database.Where("from_user_id = ?", userId).Find(&transfers)

	return c.JSON(fiber.Map{"status": "success" , "data": transfers})
}