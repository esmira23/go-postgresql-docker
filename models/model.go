package models

type Data struct {
	TransactionId      int     `json:"transaction_id" db:"transaction_id"`
	RequestId          int     `json:"request_id" db:"request_id"`
	TerminalId         int     `json:"terminal_id" db:"terminal_id"`
	PartnerObjectId    int     `json:"partner_object_id" db:"partner_object_id"`
	AmountTotal        float64 `json:"amount_total" db:"amount_total"`
	AmountOriginal     float64 `json:"amount_original" db:"amount_original"`
	CommissionPS       float64 `json:"commission_ps" db:"commission_ps"`
	CommissionClient   float64 `json:"commission_client" db:"commission_client"`
	CommissionProvider float64 `json:"commission_provider" db:"commission_provider"`
	DateInput          string  `json:"date_input" db:"date_input"`
	DatePost           string  `json:"date_post" db:"date_post"`
	Status             string  `json:"status" db:"status"`
	PaymentType        string  `json:"payment_type" db:"payment_type"`
	PaymentNumber      string  `json:"payment_number" db:"payment_number"`
	ServiceId          int     `json:"service_id" db:"service_id"`
	Service            string  `json:"service" db:"service"`
	PayeeId            int     `json:"payee_id" db:"payee_id"`
	PayeeName          string  `json:"payee_name" db:"payee_name"`
	PayeeBankMfo       int     `json:"payee_bank_mfo" db:"payee_bank_mfo"`
	PayeeBankAccount   string  `json:"payee_bank_account" db:"payee_bank_account"`
	PaymentNarrative   string  `json:"payment_narrative" db:"payment_narrative"`
}

type ErrorMsg struct {
	Message string `json:"message"`
}
