package repository

import (
	"encoding/json"
	"simpedonationapps/entity"

	"github.com/go-redis/redis"
	"github.com/google/uuid"
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

func (repository *donationRepositoryImpl) CreateDonation(donation entity.Donation) (bool, string) {
	randomString := uuid.New().String()
	donationJSON, err := json.Marshal(donation)

	if err != nil {
		return false, err.Error()
	}
	err1 := repository.Client.Set(randomString, donationJSON, 0).Err()
	if err1 != nil {
		return false, err.Error()
	}
	return true, ""
}

func (repository *donationRepositoryImpl) UpdateDonation(key string, donation entity.Donation) (bool, string) {
	_, err := repository.Client.Get(key).Result()
	if err != nil {
		return false, err.Error()
	}

	donationJSON, err := json.Marshal(donation)
	if err != nil {
		return false, err.Error()
	}

	err1 := repository.Client.Set(key, donationJSON, 0).Err()
	if err1 != nil {
		return false, err.Error()
	}
	return true, ""
}

func (repository *donationRepositoryImpl) DeleteDonation(key string) (bool, string) {
	_, err := repository.Client.Get(key).Result()
	if err != nil {
		return false, err.Error()
	}

	val := repository.Client.Del(key)
	if val.Val() == 0 {
		return false, err.Error()
	}
	return true, ""
}
