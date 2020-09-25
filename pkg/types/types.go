package types

// Money представляет ссобой денежную сумму в минимальних единицах
type Money int64

// Category представляет с собой котегори, в которо был совершён платеж (авто, аптеки , ресторани и т д)
type Category string

// 
type Payment struct{
	ID                 int
    Amount             Money
    Category           Category
}
