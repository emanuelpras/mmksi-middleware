package request

import (
	"middleware-mmksi/salesforce/response"

	validation "github.com/go-ozzo/ozzo-validation"
)

type SparepartSalesHistoryRequest struct {
	SalesforceID            string
	Dnet_ID__c              string
	Transaction_Date__c     string
	Parts_Code__c           string
	Parts_Name__c           string
	Quantity__c             string
	Is_Campaign__c          string
	Campaign_No__c          string
	Campaign_Description__c string
	Status__c               string
	Sales_Price__c          string
}

func (f *SparepartSalesHistoryRequest) Validate() error {
	if err := validation.Validate(f.SalesforceID, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "SalesforceID not found",
				"id": "SalesforceID tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Dnet_ID__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Dnet_ID__c not found",
				"id": "Dnet_ID__c tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Transaction_Date__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Transaction_Date__c not found",
				"id": "Transaction_Date__c tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Parts_Code__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Parts_Code__c not found",
				"id": "Parts_Code__c tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Parts_Name__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Parts_Name__c not found",
				"id": "Parts_Name__c tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Quantity__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Quantity__c not found",
				"id": "Quantity__c tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Is_Campaign__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Is_Campaign__c not found",
				"id": "Is_Campaign__c tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Campaign_No__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Campaign_No__c not found",
				"id": "Campaign_No__c tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Campaign_Description__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Campaign_Description__c not found",
				"id": "Campaign_Description__c tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Status__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Status__c not found",
				"id": "Status__c tidak ditemukan",
			},
		}
	}
	if err := validation.Validate(f.Sales_Price__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"en": "Sales_Price__c not found",
				"id": "Sales_Price__c tidak ditemukan",
			},
		}
	}
	return nil
}
