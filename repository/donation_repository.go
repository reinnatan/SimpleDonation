package repository

import "simpedonationapps/entity"

type DonationRespository interface {
	CreateDonation(donation entity.Donation) (bool, string)
	UpdateDonation(key string, donation entity.Donation) (bool, string)
	DeleteDonation(key string) (bool, string)
}
