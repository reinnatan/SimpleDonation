package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Donation struct {
	Id          string
	Description string `json:"description" xml:"description" form:"description"`
	Total       int64  `json:"total" xml:"total" form:"total"`
	DueDate     int64  `json:"dueDate" xml:"dueDate" form:"dueDate"`
}

type Donatur struct {
	Id         string `json:"id" xml:"id" form:"id"`
	Name       string `json:"name" xml:"name" form:"name"`
	DueDate    int64  `json:"dueDate" xml:"dueDate" form:"dueDate"`
	IdDonation string `json:"idDonation" xml:"idDonation" form:"idDonation"`
}

type Message struct {
	Message string `json:"message" xml:"message" form:"message"`
}

func main() {
	app := fiber.New()
	//ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	app.Get("/", func(c *fiber.Ctx) error {
		msg := fmt.Sprintf("✋ %s", "Hello World")
		return c.SendString(msg)
	})

	app.Post("/create-donation", func(c *fiber.Ctx) error {
		donation := new(Donation)
		if err := c.BodyParser(donation); err != nil {
			return err
		}
		donation.Id = uuid.New().String()
		//donation.DueDate = time.Now().Unix()
		valueMarshal, err := json.Marshal(donation)
		if err != nil {
			panic(err)
		}
		err1 := rdb.Set(donation.Id, valueMarshal, 0).Err()
		if err1 != nil {
			panic(err1)
		}

		message := Message{}
		message.Message = fmt.Sprintf("Donation successfully created %s", donation.Id)
		return c.JSON(message)
	})

	app.Put("/update-donation/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		_, err := rdb.Get(id).Result()
		if err != nil {
			panic(err)
		}

		donation := new(Donation)
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
	})

	app.Delete("/delete-donation/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		_, err := rdb.Get(id).Result()
		if err != nil {
			message := Message{}
			message.Message = fmt.Sprintf("Couldn't found for keys %s", id)
			return c.JSON(message)
		}

		value := rdb.Del(id)
		if value.Val() == 0 {
			message := Message{}
			message.Message = "There are some error for delete"
			return c.JSON(message)
		}

		message := Message{}
		message.Message = "Donation successfully delete"
		return c.JSON(message)
	})

	/*
		err := rdb.Set(ctx, "key", "value", 0).Err()
		if err != nil {
			panic(err)
		}

		val, err := rdb.Get(ctx, "key").Result()
		if err != nil {
			panic(err)
		}
		fmt.Println("key", val)

		val2, err := rdb.Get(ctx, "key2").Result()
		if err == redis.Nil {
			fmt.Println("key2 does not exist")
		} else if err != nil {
			panic(err)
		} else {
			fmt.Println("key2", val2)
		}
	*/

	log.Fatal(app.Listen(":3000"))
}
