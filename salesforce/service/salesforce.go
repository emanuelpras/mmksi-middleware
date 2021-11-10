package service

import (
	"middleware-mmksi/salesforce/repo"
	"middleware-mmksi/salesforce/response"
	"middleware-mmksi/salesforce/service/request"
)

type SalesforceService interface {
	GetTokenSales() (*response.TokenOauthResponse, error)
	GetServiceHistory(params request.ServiceHistoryRequest, authorizationSalesforce request.SalesRequestAuthorization) (*response.ServiceHistoryResponse, error)
	GetSparepartSalesHistory(params request.SparepartSalesHistoryRequest, authorizationSalesforce request.SalesRequestAuthorization) (*response.ServiceHistoryResponse, error)
}

type salesforceService struct {
	salesforceRepo repo.SalesforceRepo
}

func NewSalesforceService(
	salesforceRepo repo.SalesforceRepo,
) SalesforceService {
	return &salesforceService{
		salesforceRepo: salesforceRepo,
	}
}

func (s *salesforceService) GetTokenSales() (*response.TokenOauthResponse, error) {

	result, err := s.salesforceRepo.GetTokenSales()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *salesforceService) GetServiceHistory(params request.ServiceHistoryRequest, authorizationSalesforce request.SalesRequestAuthorization) (*response.ServiceHistoryResponse, error) {
	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.salesforceRepo.GetServiceHistory(request.ServiceHistoryRequest{
		MSP_No__c:             params.MSP_No__c,
		Dnet_ID__c:            params.Dnet_ID__c,
		Dealer_code__c:        params.Dealer_code__c,
		Service_Start_Date__c: params.Service_Start_Date__c,
		Service_Start_Time__c: params.Service_Start_Time__c,
		Service_End_Date__c:   params.Service_End_Date__c,
		Service_End_Time__c:   params.Service_End_Time__c,
		Mechanic_Name__c:      params.Mechanic_Name__c,
		Work_Order_Number__c:  params.Work_Order_Number__c,
		No_Rangka__c:          params.No_Rangka__c,
		Service_Kind__c:       params.Service_Kind__c,
		Odometer__c:           params.Odometer__c,
		Service_Type__c:       params.Service_Type__c,
		Stall_Code__c:         params.Stall_Code__c,
		Booking_Code__c:       params.Booking_Code__c,
		Status__c:             params.Status__c,
		Mechanic_Notes__c:     params.Mechanic_Notes__c,
	}, request.SalesRequestAuthorization{
		AccessToken: authorizationSalesforce.AccessToken,
		TokenType:   authorizationSalesforce.TokenType,
		InstanceURL: authorizationSalesforce.InstanceURL,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *salesforceService) GetSparepartSalesHistory(params request.SparepartSalesHistoryRequest, authorizationSalesforce request.SalesRequestAuthorization) (*response.ServiceHistoryResponse, error) {

	if err := params.Validate(); err != nil {
		return nil, err
	}

	result, err := s.salesforceRepo.GetSparepartSalesHistory(request.SparepartSalesHistoryRequest{
		SalesforceID:            params.SalesforceID,
		Dnet_ID__c:              params.Dnet_ID__c,
		Transaction_Date__c:     params.Transaction_Date__c,
		Parts_Code__c:           params.Parts_Code__c,
		Parts_Name__c:           params.Parts_Name__c,
		Quantity__c:             params.Quantity__c,
		Is_Campaign__c:          params.Is_Campaign__c,
		Campaign_No__c:          params.Campaign_No__c,
		Campaign_Description__c: params.Campaign_Description__c,
		Status__c:               params.Status__c,
		Sales_Price__c:          params.Sales_Price__c,
	}, request.SalesRequestAuthorization{
		AccessToken: authorizationSalesforce.AccessToken,
		TokenType:   authorizationSalesforce.TokenType,
		InstanceURL: authorizationSalesforce.InstanceURL,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}
