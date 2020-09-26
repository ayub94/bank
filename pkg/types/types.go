package types

// Money представляет ссобой денежную сумму в минимальних единицах
type Money int64

// Category представляет с собой котегори, в которо был совершён платеж (авто, аптеки , ресторани и т д)
type Category string

// Status представляет с собой статус платежей
type Status string

// Предопределенные статусы платежей
const (
    StatusOk    Status = "Ok"
    StatusFail    Status = "FAIL"
    StatusInProgress   Status = "INPROGRESS"
)


// Payment
type Payment struct{
	ID                 int
    Amount             Money
    Category           Category
    Status             Status
}
