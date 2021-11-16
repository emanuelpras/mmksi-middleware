package request

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

type BatchServiceHistoryRequest []ServiceHistoryRequest
