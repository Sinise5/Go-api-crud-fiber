package controllers

import (
	"myapp/config"
	"myapp/models"

	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	err := config.DB.Select(&users, "SELECT id, username FROM users")
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to get users"})
	}
	return c.JSON(users)
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	err := config.DB.Get(&user, "SELECT id, username FROM users WHERE id=$1", id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid input"})
	}

	_, err := config.DB.Exec("UPDATE users SET username=$1 WHERE id=$2", user.Username, id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to update user"})
	}
	return c.JSON(fiber.Map{"message": "User updated"})
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	_, err := config.DB.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to delete user"})
	}
	return c.JSON(fiber.Map{"message": "User deleted"})
}
