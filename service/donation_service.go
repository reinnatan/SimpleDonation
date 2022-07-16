package service

import (
	"simpedonationapps/model"
)

type DonationService interface {
	CreateDonation(donation model.Donation) model.GeneralResponse
	UpdateDonation(key string, donation model.Donation) model.GeneralResponse
	DeleteDonation(key string) model.GeneralResponse
}
