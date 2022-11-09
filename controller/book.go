package controller

import (
	"workshop/project/database"
	"workshop/project/model"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
	db := database.DBConn

	var books []model.Book
	db.Find(&books)

	return c.JSON(books)
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")
	db := database.DBConn

	var book model.Book
	db.Find(&book, id)

	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn

	book := new(model.Book)

	if err := c.BodyParser(book); err != nil {
		return err
	}

	db.Create(&book)

	return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
	db := database.DBConn

	id := c.Params("id")

	var book model.Book
	db.First(&book, id)

	if book.Title == "" {
		return c.Status(500).SendString("No book found with given ID")
	}

	db.Delete(&book)

	return c.SendString("Book successfully deleted")
}

func UpdateBook(c *fiber.Ctx) error {
	db := database.DBConn

	id := c.Params("id")

	var book model.Book
	db.First(&book, id)

	if book.Title == "" {
		return c.Status(500).SendString("No book found with given ID")
	}

	if err := c.BodyParser(&book); err != nil {
		return err
	}

	db.Save(&book)

	return c.JSON(book)
}
