package service

import (
	"middleware-mmksi/mmid/repo"
	"middleware-mmksi/mmid/response"
	"middleware-mmksi/mmid/service/request"
)

type MmidService interface {
	GetServiceHistory(params request.ServiceHistoryRequest) (*response.ServiceHistoryResponse, error)
	GetServiceHistoryBatch(params request.Batch) (*response.ServiceHistoryBatchResponse, error)
	GetSparepartList(params request.SparepartListRequest) (*response.SparepartListResponse, error)
}

type mmidService struct {
	mmidRepo repo.MmidRepo
}

func NewMmidService(
	mmidRepo repo.MmidRepo,
) MmidService {
	return &mmidService{
		mmidRepo: mmidRepo,
	}
}

func (s *mmidService) GetServiceHistory(params request.ServiceHistoryRequest) (*response.ServiceHistoryResponse, error) {

	result, err := s.mmidRepo.GetServiceHistory(request.ServiceHistoryRequest{
		Dnet_ID__c:            params.Dnet_ID__c,
		Dealer_code__c:        params.Dealer_code__c,
		Service_Start_Date__c: params.Service_Start_Date__c,
		Service_Start_Time__c: params.Service_Start_Time__c,
		Service_End_Date__c:   params.Service_End_Date__c,
		Service_End_Time__c:   params.Service_End_Time__c,
		Mechanic_Name__c:      params.Mechanic_Name__c,
		Work_Order_Number__c:  params.Work_Order_Number__c,
		No_Rangka__c:          params.No_Rangka__c,
		Mechanic_Notes__c:     params.Mechanic_Notes__c,
		Service_Kind__c:       params.Service_Kind__c,
		Odometer__c:           params.Odometer__c,
		MSP_No__c:             params.MSP_No__c,
		Service_Type__c:       params.Service_Type__c,
		Stall_Code__c:         params.Stall_Code__c,
		Booking_Code__c:       params.Booking_Code__c,
		Status__c:             params.Status__c,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *mmidService) GetServiceHistoryBatch(params request.Batch) (*response.ServiceHistoryBatchResponse, error) {

	result, err := s.mmidRepo.GetServiceHistoryBatch(params)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *mmidService) GetSparepartList(params request.SparepartListRequest) (*response.SparepartListResponse, error) {

	result, err := s.mmidRepo.GetSparepartList(params)

	if err != nil {
		return nil, err
	}

	return result, nil
}
