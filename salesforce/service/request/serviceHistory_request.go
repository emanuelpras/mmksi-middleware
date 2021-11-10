package request

import (
	"middleware-mmksi/salesforce/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type HeaderRequest struct {
	Authorization string
}

type Misal struct {
	Code int
}
type ServiceHistoryRequest struct {
	MSP_No__c             string
	Dnet_ID__c            string
	Dealer_code__c        string
	Service_Start_Date__c string
	Service_Start_Time__c string
	Service_End_Date__c   string
	Service_End_Time__c   string
	Mechanic_Name__c      string
	Work_Order_Number__c  string
	No_Rangka__c          string
	Service_Kind__c       string
	Odometer__c           string
	Service_Type__c       string
	Stall_Code__c         string
	Booking_Code__c       string
	Status__c             string
	Mechanic_Notes__c     string
}

type SalesRequestAuthorization struct {
	AccessToken string
	TokenType   string
	InstanceURL string
}

func (f *ServiceHistoryRequest) Validate() error {

	if err := validation.Validate(f.Dnet_ID__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "DNET_ID not found",
				"id": "DNET_ID tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Dealer_code__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "DEALER_CODE not found",
				"id": "DEALER_CODE tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Service_Start_Date__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "SERVICE_START_DATE not found",
				"id": "SERVICE_START_DATE tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Service_Start_Time__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "SERVICE_START_TIME not found",
				"id": "SERVICE_START_TIME tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Service_End_Date__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "SERVICE_END_DATE not found",
				"id": "SERVICE_END_DATE tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Service_End_Time__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "SERVICE_END_TIME not found",
				"id": "SERVICE_END_TIME tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.No_Rangka__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "NO_RANGKA not found",
				"id": "NO_RANGKA tidak ditemukan",
			},
		}
	}
	return nil
}

func (f *HeaderRequest) Validate() error {
	if err := validation.Validate(f.Authorization, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Authorization not found",
				"id": "Authorization tidak ditemukan",
			},
		}
	}

	return nil
}
