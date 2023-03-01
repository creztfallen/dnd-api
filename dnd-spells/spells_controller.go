package spells

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

var spellCollection *mongo.Collection = configs.GetCollection(configs.DB, "spells")
var validate = validator.New()

func CreateSpell(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var spell Spell
	defer cancel()

	if err := c.BodyParser(&spell); err != nil {
		return c.Status(http.StatusBadRequest).JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(&spell); validationErr != nil {
		return c.Status(http.StatusBadRequest).JSON(dtos.Response{
			Status:  http.StatusBadRequest,
			Message: "error",
			Data:    &fiber.Map{"data": validationErr.Error()}})
	}

	newSpell := Spell{
		Id:              primitive.NewObjectID(),
		Name:            spell.Name,
		Level:           spell.Level,
		School:          spell.School,
		ConjurationTime: spell.ConjurationTime,
		SpellRange:      spell.SpellRange,
		Components:      spell.Components,
		Duration:        spell.Duration,
		Description:     spell.Description,
		Source:          spell.Source,
	}

	result, err := spellCollection.InsertOne(ctx, newSpell)
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
