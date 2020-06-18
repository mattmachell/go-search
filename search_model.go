package main

//SearchDocument search document format
type SearchDocument struct {
	Index   string `json:"_index"`
	Type    string `json:"_type"`
	ID      string `json:"_id"`
	Version int    `json:"_version"`
	Found   bool   `json:"found"`
	Source  struct {
		UID             string `json:"uId"`
		NormalizedUID   int64  `json:"normalizedUid"`
		CaseRecNumber   string `json:"caseRecNumber"`
		WorkPhoneNumber struct {
			ID          int    `json:"id"`
			PhoneNumber string `json:"phoneNumber"`
			Type        string `json:"type"`
			Default     bool   `json:"default"`
			ClassName   string `json:"className"`
		} `json:"workPhoneNumber"`
		HomePhoneNumber   interface{} `json:"homePhoneNumber"`
		MobilePhoneNumber interface{} `json:"mobilePhoneNumber"`
		Email             string      `json:"email"`
		Dob               string      `json:"dob"`
		Firstname         string      `json:"firstname"`
		Middlenames       string      `json:"middlenames"`
		Surname           string      `json:"surname"`
		AddressLine1      string      `json:"addressLine1"`
		AddressLine2      string      `json:"addressLine2"`
		AddressLine3      string      `json:"addressLine3"`
		Town              string      `json:"town"`
		County            string      `json:"county"`
		Postcode          string      `json:"postcode"`
		Country           string      `json:"country"`
		IsAirmailRequired bool        `json:"isAirmailRequired"`
		Addresses         []struct {
			AddressLines []string `json:"addressLines"`
			Postcode     string   `json:"postcode"`
			ClassName    string   `json:"className"`
		} `json:"addresses"`
		PhoneNumber  string `json:"phoneNumber"`
		PhoneNumbers []struct {
			ID          int    `json:"id"`
			PhoneNumber string `json:"phoneNumber"`
			Type        string `json:"type"`
			Default     bool   `json:"default"`
			ClassName   string `json:"className"`
		} `json:"phoneNumbers"`
		PersonType string `json:"personType"`
		Cases      []struct {
			UID           string      `json:"uId"`
			NormalizedUID int64       `json:"normalizedUid"`
			CaseRecNumber string      `json:"caseRecNumber"`
			BatchID       interface{} `json:"batchId"`
			ClassName     string      `json:"className"`
		} `json:"cases"`
		CompanyName interface{} `json:"companyName"`
		ClassName   string      `json:"className"`
	} `json:"_source"`
}

// Search JSON payload search
type Search struct {
	Search string
}

// Results search results
type Results struct {
	Count int
}
