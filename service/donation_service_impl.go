package service

import (
	"simpedonationapps/entity"
	"simpedonationapps/model"
	"simpedonationapps/repository"
)

func NewDomainService(donationRepository *repository.DonationRespository) DonationService {
	return &domainServiceImpl{
		DonationRepository: *donationRepository,
	}
}

type domainServiceImpl struct {
	DonationRepository repository.DonationRespository
}

func (service *domainServiceImpl) CreateDonation(donation model.Donation) model.GeneralResponse {
	donationEntity := entity.Donation{
		Id:          donation.Id,
		Description: donation.Description,
		Total:       donation.Total,
		DueDate:     donation.DueDate,
	}
	isSuccess, errMessage := service.DonationRepository.CreateDonation(donationEntity)
	response := model.GeneralResponse{}
	if isSuccess {
		response.ResponseCode = 200
		response.Message = "Successfully create donation"
	} else {
		response.ResponseCode = 400
		response.Message = errMessage
	}
	return response
}
func (service *domainServiceImpl) UpdateDonation(key string, donation model.Donation) model.GeneralResponse {
	donationEntity := entity.Donation{
		Id:          donation.Id,
		Description: donation.Description,
		Total:       donation.Total,
		DueDate:     donation.DueDate,
	}
	isSuccess, errMessage := service.DonationRepository.UpdateDonation(key, donationEntity)
	response := model.GeneralResponse{}
	if isSuccess {
		response.ResponseCode = 200
		response.Message = "Successfully update donation"
	} else {
		response.ResponseCode = 400
		response.Message = errMessage
	}

	return response
}

func (service *domainServiceImpl) DeleteDonation(key string) model.GeneralResponse {
	isSuccess, errMessage := service.DonationRepository.DeleteDonation(key)
	response := model.GeneralResponse{}
	if isSuccess {
		response.ResponseCode = 200
		response.Message = "Successfully delete donation"
	} else {
		response.ResponseCode = 400
		response.Message = errMessage
	}
	return response
}
