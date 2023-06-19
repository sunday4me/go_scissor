package server

import (
	"fmt"
	"scissor/model"
	"scissor/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func redirect(c *fiber.Ctx) error {
	scissorUrl := c.Params("redirect")
	scissor, err := model.FindByScissorUrl(scissorUrl)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not find scissor in DB " + err.Error(),
		})
	}
	// grab any stats you want...
	scissor.Clicked += 1
	err = model.UpdateScissor(scissor)
	if err != nil {
		fmt.Printf("error updating: %v\n", err)
	}


	return c.Redirect(scissor.Redirect, fiber.StatusTemporaryRedirect)
}

func getAllScissors(c *fiber.Ctx) error {
	scissors, err := model.GetAllScissors()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error getting all scissor links " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(scissors)
}

func getScissor(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not parse id " + err.Error(),
		})
	}

	scissor, err := model.GetScissor(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error could not retreive scissor from db " + err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(scissor)
}

func createScissor(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var scissor model.Scissor
	err := c.BodyParser(&scissor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "error parsing JSON " + err.Error(),
		})
	}

	if scissor.Random {
		scissor.Scissor = utils.RandomURL(8)
	}

	err = model.CreateScissor(scissor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not create scissor in db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(scissor)

}

func updateScissor(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var scissor model.Scissor

	err := c.BodyParser(&scissor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse json " + err.Error(),
		})
	}

	err = model.UpdateScissor(scissor)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not update scissor link in DB " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(scissor)
}

func deleteScissor(c *fiber.Ctx) error {
	id, err := strconv.ParseUint( c.Params("id"), 10, 64)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not parse id from url " + err.Error(),
		})
	}

	err = model.DeleteScissor(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map {
			"message": "could not delete from db " + err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map {
		"message": "scissor deleted.",
	})
}


func SetupAndListen() {

	router := fiber.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Get("/r/:redirect", redirect)

	router.Get("/scissor", getAllScissors)
	router.Get("/scissor/:id", getScissor)
	router.Post("/scissor", createScissor)
	router.Patch("/scissor", updateScissor)
	router.Delete("/scissor/:id", deleteScissor)

	router.Listen(":3000")
	
}
