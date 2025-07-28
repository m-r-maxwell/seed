package models

import "time"

type Reviews struct {
	ReviewerName string    `json:"reviewer_name" parquet:"name=reviewer_name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Rating       int       `json:"rating" parquet:"name=rating, type=INT32"`
	ProductID    string    `json:"product_id" parquet:"name=product_id, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Date         time.Time `json:"date" parquet:"name=date, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
}
type Orders struct {
	OrderID    string  `json:"order_id" parquet:"name=order_id, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	ProductID  string  `json:"product_id" parquet:"name=product_id, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Quantity   int     `json:"quantity" parquet:"name=quantity, type=INT32"`
	Price      float64 `json:"price" parquet:"name=price, type=DOUBLE"`
	CustomerID string  `json:"customer_id" parquet:"name=customer_id, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Status     string  `json:"status" parquet:"name=status, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
}
type Companies struct {
	Name        string  `json:"name" parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Address     string  `json:"address" parquet:"name=address, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Phone       string  `json:"phone" parquet:"name=phone, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Website     string  `json:"website" parquet:"name=website, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Employees   int     `json:"employees" parquet:"name=employees, type=INT32"`
	Revenue     float64 `json:"revenue" parquet:"name=revenue, type=DOUBLE"`
	Established int     `json:"established" parquet:"name=established, type=INT32"`
}

type Loan struct {
	Amount   float64 `json:"amount" parquet:"name=amount, type=DOUBLE"`
	Interest float64 `json:"interest" parquet:"name=interest, type=DOUBLE"`
	Term     int     `json:"term" parquet:"name=term, type=INT32"` // Loan term in months
	Default  string  `json:"default" parquet:"name=default, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Gender   string  `json:"gender" parquet:"name=gender, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Age      int     `json:"age" parquet:"name=age, type=INT32"`        // Age of the borrower
	Income   float64 `json:"income" parquet:"name=income, type=DOUBLE"` // Monthly income of the borrower
}

type Products struct {
	Name        string  `json:"name" parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Description string  `json:"description" parquet:"name=description, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Price       float64 `json:"price" parquet:"name=price, type=DOUBLE"`
	Category    string  `json:"category" parquet:"name=category, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
}

type People struct {
	Name    string `parquet:"name=name, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Email   string `parquet:"name=email, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Address string `parquet:"name=address, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
	Phone   string `parquet:"name=phone, type=BYTE_ARRAY, convertedtype=UTF8, encoding=PLAIN"`
}
