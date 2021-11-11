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
	if err := validation.Validate(f.Dnet_ID__c, validation.Required); err != nil {
		return &response.ErrorResponse{
			ErrorID: 422,
			Msg: map[string]string{
				"id": "Format data tidak sesuai",
				"en": "Invalid data format",
			},
		}
	}
	if err := validation.Validate(f.Transaction_Date__c, validation.Required); err != nil {
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
