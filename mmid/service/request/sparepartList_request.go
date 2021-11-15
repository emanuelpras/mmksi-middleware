package request

type SparepartList struct {
	Transaction_Date__c     string
	Parts_Code__c           string
	Parts_Name__c           string
	Quantity__c             int
	Sales_Price__c          int
	Is_Campaign__c          string
	Campaign_No__c          string
	Campaign_Description__c string
	Status__c               string
	Dnet_ID__c              string
	Dnet_Sparepart_ID__c    string
}

type SparepartListRequest []SparepartList
