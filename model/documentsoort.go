package model

type DocumentSoort string

func NewDocumentSoort(str string) DocumentSoort {
	return DocumentSoort(str)
}

func (soort DocumentSoort) String() string {
	return string(soort)
}

const (
	VendorInvoice   DocumentSoort = "KR"
	VendorPayment   DocumentSoort = "KZ"
	CustomerInvoice DocumentSoort = "DR"
	CustomerPayment DocumentSoort = "DZ"
)
