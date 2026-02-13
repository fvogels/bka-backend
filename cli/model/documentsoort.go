package model

type DocumentSoort string

func (soort DocumentSoort) String() string {
	return string(soort)
}

const (
	VendorInvoice   DocumentSoort = "KR"
	VendorPayment   DocumentSoort = "KZ"
	CustomerInvoice DocumentSoort = "DR"
	CustomerPayment DocumentSoort = "DZ"
)
