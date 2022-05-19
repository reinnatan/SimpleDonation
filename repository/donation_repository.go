package repository

import "simpedonationapps/entity"

type DonationRespository interface {
	CreateDonation(donation entity.Donation)
	UpdateDonation(key string, donation entity.Donation)
	DeleteDonation(key string)
}
