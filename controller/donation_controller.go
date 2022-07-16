package controller

import (
	"encoding/json"
	"fmt"
	"simpedonationapps/model"
	"simpedonationapps/service"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DonationController struct {
	donationService service.DonationService
}

func NewDonationController(donationService *service.DonationService) DonationController {
	return DonationController{donationService: *donationService}
}

func (controller *DonationController) Route(app *fiber.App) {
	app.Post("/api/create-donation", controller.CreateDonation)
	app.Put("/update-donation/:id", controller.UpdateDonation)
	app.Delete("/delete-donation/:id", controller.DeleteDonation)
}

func (controller *DonationController) CreateDonation(c *fiber.Ctx) error {
	donation := new(model.Donation)
	if err := c.BodyParser(donation); err != nil {
		return err
	}
	donation.Id = uuid.New().String()
	donation.DueDate = time.Now().Unix()
	responseServer := controller.donationService.CreateDonation(*donation)
	return c.JSON(responseServer)
}

func (controller *DonationController) UpdateDonation(c *fiber.Ctx) error {
	id := c.Params("id")
	donation := new(model.Donation)
	if err := c.BodyParser(donation); err != nil {
		return err
	}

	controller.donationService.UpdateDonation(id, *donation)
	_, err := rdb.Get(id).Result()
	if err != nil {
		panic(err)
	}

	donation := new(model.Donation)
	if err := c.BodyParser(donation); err != nil {
		return err
	}

	valueMarshal, err := json.Marshal(donation)
	if err != nil {
		panic(err)
	}

	err1 := rdb.Set(id, valueMarshal, 0).Err()
	if err1 != nil {
		panic(err1)
	}

	message := Message{}
	message.Message = "Donation successfully update"
	return c.JSON(message)

	return fmt.Errorf("Testing 1")
}

func (controller *DonationController) DeleteDonation(c *fiber.Ctx) error {
	return fmt.Errorf("Testing 2")
}
