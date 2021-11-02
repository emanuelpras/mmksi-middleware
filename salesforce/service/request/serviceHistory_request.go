package request

import (
	"middleware-mmksi/salesforce/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ServiceHistoryRequest struct {
	MSP_No__c             string
	Dnet_ID__c            string
	Dealer_code__c        string
	Service_Start_Date__c string
	Service_Start_Time__c string
	Service_End_Date__c   string
	Service_End_Time__c   string
	Work_Order_Number__c  string
	No_Rangka__c          string
	Service_Kind__c       string
	Odometer__c           string
	Service_Type__c       string
	Stall_Code__c         string
	Booking_Code__c       string
	Status__c             string
}

type SalesRequestAuthorization struct {
	AccessToken string `form:"AccessToken"`
	TokenType   string `form:"TokenType"`
}

type HeaderAuthorizationRequest struct {
	Authorization string `form:"Authorization"`
}

func (f *ServiceHistoryRequest) Validate() error {
	if err := validation.Validate(f.MSP_No__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "MSP_NO not found",
				"id": "MSP_NO tidak ditemukan",
			},
		}
	}
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
	if err := validation.Validate(f.Work_Order_Number__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "WORK_ORDER not found",
				"id": "WORK_ORDER tidak ditemukan",
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
	if err := validation.Validate(f.Service_Kind__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "SERVICE_KIND not found",
				"id": "SERVICE_KIND tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Odometer__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "ODOMETER not found",
				"id": "ODOMETER tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Service_Type__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "SERVICE_TYPE not found",
				"id": "SERVICE_TYPE tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Stall_Code__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "STALL_Code not found",
				"id": "STALL_Code tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Booking_Code__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Booking_Code not found",
				"id": "Booking_Code tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Status__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "STATUS not found",
				"id": "STATUS tidak ditemukan",
			},
		}
	}
	return nil
}
