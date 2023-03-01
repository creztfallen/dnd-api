package dndclasses

import (
	"context"
	"net/http"
	"time"

	dtos "github.com/dnd-api/DTOs"
	"github.com/dnd-api/configs"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var classCollection *mongo.Collection = configs.GetCollection(configs.DB, "classes")
var validate = validator.New()

func CreateClass(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var class Classes
	defer cancel()

	if err := c.BodyParser(&class); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&class); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": validationErr.Error()}})
	}

	newClass := Classes{
		Id:   primitive.NewObjectID(),
		Name: class.Name,
		Key:  class.Key,
		URL:  class.URL,
	}

	result, err := classCollection.InsertOne(ctx, newClass)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(dtos.Response{
			Status:  http.StatusInternalServerError,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	return c.Status(http.StatusCreated).JSON(dtos.Response{
		Status:  http.StatusCreated,
		Message: "success",
		Data:    &fiber.Map{"data": result}})
}
