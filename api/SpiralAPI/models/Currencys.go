package models

type CurrencyData struct {
	Id                int64   `json:"id"`
	Code              string  `json:"code"`
	Name              string  `json:"name"`
	Precision         int64   `json:"precision"`
	CanDeposit        bool    `json:"can_deposit"`
	CanWithdrawal     bool    `json:"can_withdrawal"`
	MinConfirms       int64   `json:"min_confirms"`
	WithdrawalFee     float64 `json:"withdrawal_fee,string"`
	WithdrawMinAmount float64 `json:"withdraw_min_amount,string"`
}

type CurrencysReturn struct {
	Data []CurrencyData `json:"data"` // 系统支持的所有币种
	errorResponse
}
