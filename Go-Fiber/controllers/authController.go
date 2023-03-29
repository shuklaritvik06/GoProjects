package controllers

import (
	"context"
	"time"

	"github.com/gofiber/fiber"
	"github.com/shuklaritvik06/GoProjects/fiber/database"
	"github.com/shuklaritvik06/GoProjects/fiber/models"
	"github.com/shuklaritvik06/GoProjects/fiber/utils"
)

func Login(c *fiber.Ctx) {
	var user models.User
	var found models.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}
	database.GetDB().Database("users").Collection("user").FindOne(context.Background(), fiber.Map{"email": user.Email}).Decode(&found)
	if utils.CheckPasswordHash(user.Password, found.Password) {
		c.Status(400).JSON(fiber.Map{"message": "Invalid credentials"})
	}
	token, refreshToken, _ := utils.GetTokens(found.Email, found.First_Name, found.Last_Name)
	c.Cookie(&fiber.Cookie{
		Name:    "jwt",
		Value:   token,
		Path:    "/",
		Expires: time.Now().Local().Add(time.Hour * 1),
	})
	c.Cookie(&fiber.Cookie{
		Name:    "refresh",
		Value:   refreshToken,
		Path:    "/",
		Expires: time.Now().Local().Add(time.Hour * 24),
	})
	c.Status(200).JSON(fiber.Map{"token": token, "refreshToken": refreshToken})
}

func Register(c *fiber.Ctx) {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}
	user.Password, _ = utils.HashPassword(user.Password)
	database.GetDB().Database("users").Collection("user").InsertOne(context.Background(), user)
	c.Status(200).JSON(fiber.Map{"message": "User created"})
}

func User(c *fiber.Ctx) {
	if c.Cookies("jwt") == "" {
		c.Status(400).JSON(fiber.Map{"message": "Please give token"})
	}
	claims, _ := utils.ValidateToken(c.Cookies("jwt"))
	c.Status(200).JSON(fiber.Map{"message": "User found", "user": claims})
}

func RefreshToken(c *fiber.Ctx) {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		c.Status(400).JSON(fiber.Map{"message": "Invalid data"})
	}
	claims, _ := utils.ValidateToken(c.Cookies("refresh"))
	token, refreshToken, _ := utils.GetTokens(claims.Email, claims.First_Name, claims.Last_Name)
	c.Status(200).JSON(fiber.Map{"token": token, "refreshToken": refreshToken})
}
