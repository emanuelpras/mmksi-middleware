package request

import (
	"middleware-mmksi/salesforce/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type ServiceHistoryRequest struct {
	Dnet_ID__c            string
	Dealer_code__c        string
	Service_Start_Date__c string
	Service_Start_Time__c string
	Service_End_Date__c   string
	Service_End_Time__c   string
	Mechanic_Name__c      string
	Work_Order_Number__c  string
	No_Rangka__c          string
	Mechanic_Notes__c     string
	Service_Kind__c       string
	Odometer__c           string
	MSP_No__c             string
	Service_Type__c       string
	Stall_Code__c         string
	Booking_Code__c       string
	Status__c             string
}

type Batch []ServiceHistoryRequest

func (f *ServiceHistoryRequest) Validate() error {

	if err := validation.Validate(f.Dnet_ID__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Dealer_code__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Service_Start_Date__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Service_Start_Time__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Service_End_Date__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Service_End_Time__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Mechanic_Name__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Work_Order_Number__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.No_Rangka__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Service_Kind__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Odometer__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.MSP_No__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Service_Type__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Status__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	return nil
}
