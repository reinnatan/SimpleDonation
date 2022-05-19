package repository

import (
	"encoding/json"
	"fmt"
	"simpedonationapps/entity"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2/internal/uuid"
)

type donationRepositoryImpl struct {
	Client *redis.Client
}

func NewDonationRepository() donationRepositoryImpl {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	return donationRepositoryImpl{
		Client: rdb,
	}

}

func (repository *donationRepositoryImpl) CreateDonation(donation entity.Donation) {
	randomString := uuid.New().String()
	donationJSON, err := json.Marshal(donation)

	if err != nil {
		panic(err)
	}
	err1 := repository.Client.Set(randomString, donationJSON, 0).Err()
	if err1 != nil {
		panic(err1)
	}

}

func (repository *donationRepositoryImpl) UpdateDonation(key string, donation entity.Donation) {
	_, err := repository.Client.Get(key).Result()
	if err != nil {
		panic(err)
	}

	donationJSON, err := json.Marshal(donation)
	if err != nil {
		panic(err)
	}

	err1 := repository.Client.Set(key, donationJSON, 0).Err()
	if err1 != nil {
		panic(err)
	}

}

func (repository *donationRepositoryImpl) DeleteDonation(key string) {
	_, err := repository.Client.Get(key).Result()
	if err != nil {
		panic(err)
	}

	val := repository.Client.Del(key)
	if val.Val() == 0 {
		panic(fmt.Errorf("Key %s couldn't be delete"))
	}

}
