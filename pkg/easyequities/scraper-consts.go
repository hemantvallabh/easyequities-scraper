package easyequities

const labelKey string = "Label"
const valueKey string = "Value"
const fundSummaryAvailable string = "Your Funds to Invest"
const fundSummaryWithdrawal string = "Withdrawable Funds"
const fundSummaryUnsettledCash string = "Unsettled Cash"
const fundSummaryLockedFunds string = "Locked Funds"

type DocumentType string

const (
	Statement DocumentType = "statement"
	TaxCertificate = "taxCertificate"
)
