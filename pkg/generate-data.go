package pkg

import (
	models "seed/models"
	"time"

	faker "github.com/brianvoe/gofakeit/v6"
)

// anything that returns a float....we should limit to 2 decimal places

// GeneratePeopleData generates fake people data
func GeneratePeopleData(count int) []models.People {
	people := make([]models.People, count)
	for i := range count {
		person := models.People{
			Name:    faker.Name(),
			Email:   faker.Email(),
			Address: faker.Address().Address,
			Phone:   faker.Phone(),
		}
		people[i] = person
	}
	return people
}

// GenerateProductsData generates fake products data
func GenerateProductsData(count int) []models.Products {
	products := make([]models.Products, count)
	for i := range count {
		product := models.Products{
			Name:        faker.ProductName(),
			Description: faker.ProductDescription(),
			Price:       faker.Price(10, 10000),
			Category:    faker.ProductCategory(),
		}
		products[i] = product
	}
	return products
}

// GenerateLoanData generates fake loan data
func GenerateLoanData(count int) []models.Loan {
	loans := make([]models.Loan, count)
	for i := range count {
		loan := models.Loan{
			Amount:   faker.Price(1000, 50000),
			Interest: faker.Price(1, 20),
			Term:     faker.Number(6, 60), // Loan term in months
			Default:  faker.RandomString([]string{"yes", "no"}),
			Gender:   faker.RandomString([]string{"male", "female"}),
			Age:      faker.Number(18, 65), // Age of the borrower
		}
		loans[i] = loan
	}
	return loans
}

// GenerateCompaniesData generates fake companies data
func GenerateCompaniesData(count int) []models.Companies {
	companies := make([]models.Companies, count)
	for i := range count {
		company := models.Companies{
			Name:        faker.Company(),
			Address:     faker.Address().Address,
			Phone:       faker.Phone(),
			Website:     faker.URL(),
			Employees:   faker.Number(1, 500),
			Revenue:     faker.Price(100000, 10000000),
			Established: faker.Year(),
		}
		companies[i] = company
	}
	return companies
}

func GenerateOrdersData(count int) []models.Orders {
	orders := make([]models.Orders, count)
	for i := range count {
		order := models.Orders{
			OrderID:    faker.UUID(),
			ProductID:  faker.UUID(),
			Quantity:   faker.Number(1, 100),
			Price:      faker.Price(10, 1000),
			CustomerID: faker.UUID(),
			Status:     faker.RandomString([]string{"pending", "shipped", "delivered", "cancelled"}),
		}
		orders[i] = order
	}
	return orders
}

func GenerateReviewsData(count int) []models.Reviews {
	reviews := make([]models.Reviews, count)
	startDate, _ := time.Parse("2006-01-02", "2020-01-01")
	endDate, _ := time.Parse("2006-01-02", "2025-12-31")
	for i := range count {
		review := models.Reviews{
			ReviewerName: faker.Name(),
			Rating:       faker.Number(1, 5),
			ProductID:    faker.UUID(),
			Date:         faker.DateRange(startDate, endDate),
		}
		reviews[i] = review
	}
	// order by date descending
	for i := 0; i < len(reviews)-1; i++ {
		for j := i + 1; j < len(reviews); j++ {
			if reviews[i].Date.Before(reviews[j].Date) {
				reviews[i], reviews[j] = reviews[j], reviews[i]
			}
		}
	}
	return reviews
}
