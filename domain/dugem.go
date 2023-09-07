package domain

type DetailShippingOrderPDFRequest struct {
	AirwaybillNumber    string
	Price               int
	ServiceType         string
	InsurancePrice      int
	Weight              int
	Quantity            int
	Recipient           Recipient
	Sender              Sender
	AdditionalInfo      string
	QRCode              string
	Barcode             string
	CompanyLogo         string
	DeliveryCompanyLogo string
}

type Sender struct {
	Name        string
	Address     string
	PhoneNumber string
}
type Recipient struct {
	Name        string
	Address     string
	PhoneNumber string
}
