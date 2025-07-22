package helpers

import (
	"encoding/csv"
	"fmt"
	"os"
	models "seed/models"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

// SavePeopleToCSV saves a slice of People to a CSV file
func SavePeopleToCSV(filename string, people []models.People) error {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Name", "Email", "Address", "Phone"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write records
	for _, p := range people {
		record := []string{p.Name, p.Email, p.Address, p.Phone}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

// SavePeopleToParquet saves a slice of People to a Parquet file
func SavePeopleToParquet(filename string, people []models.People) error {
	fmt.Println("Length of people slice:", len(people))
	fw, err := local.NewLocalFileWriter(filename + ".parquet")
	if err != nil {
		return err
	}
	pw, err := writer.NewParquetWriter(fw, &models.People{}, 4)
	if err != nil {
		return err
	}
	for i := range people {
		if err := pw.Write(&people[i]); err != nil {
			fmt.Println("Write error:", err)
			return err
		}
	}
	if err := pw.WriteStop(); err != nil {
		fmt.Println("WriteStop error:", err)
		return err
	}
	if err := fw.Close(); err != nil {
		return err
	}
	return nil
}

func SaveProductsToCSV(filename string, products []models.Products) error {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Name", "Description", "Price", "Category"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write records
	for _, p := range products {
		record := []string{p.Name, p.Description, fmt.Sprintf("%.2f", p.Price), p.Category}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func SaveProductsToParquet(filename string, products []models.Products) error {
	fw, err := local.NewLocalFileWriter(filename + ".parquet")
	if err != nil {
		return err
	}
	pw, err := writer.NewParquetWriter(fw, &models.Products{}, 4)
	if err != nil {
		return err
	}
	for i := range products {
		if err := pw.Write(&products[i]); err != nil {
			fmt.Println("Write error:", err)
			return err
		}
	}
	if err := pw.WriteStop(); err != nil {
		fmt.Println("WriteStop error:", err)
		return err
	}
	if err := fw.Close(); err != nil {
		return err
	}
	return nil
}

func SaveCompaniesToCSV(filename string, companies []models.Companies) error {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write header
	header := []string{"Name", "Address", "Phone", "Website", "Employees",
		"Revenue", "Established"}
	if err := writer.Write(header); err != nil {
		return err
	}
	// Write records
	for _, c := range companies {
		record := []string{c.Name, c.Address, c.Phone, c.Website, fmt.Sprintf("%d", c.Employees), fmt.Sprintf("%f", c.Revenue), fmt.Sprintf("%d", c.Established)}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
func SaveCompaniesToParquet(filename string, companies []models.Companies) error {
	fw, err := local.NewLocalFileWriter(filename + ".parquet")
	if err != nil {
		return err
	}
	pw, err := writer.NewParquetWriter(fw, &models.Companies{}, 4)
	if err != nil {
		return err
	}
	for i := range companies {
		if err := pw.Write(&companies[i]); err != nil {
			fmt.Println("Write error:", err)
			return err
		}
	}
	if err := pw.WriteStop(); err != nil {
		fmt.Println("WriteStop error:", err)
		return err
	}
	if err := fw.Close(); err != nil {
		return err
	}
	return nil
}

func SaveLoanToCSV(filename string, loans []models.Loan) error {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"ID", "Amount", "Date", "Status"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write records
	for _, l := range loans {
		record := []string{
			fmt.Sprintf("%f", l.Amount),
			fmt.Sprintf("%f", l.Interest),
			fmt.Sprintf("%d", l.Term),
			l.Default}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
func SaveLoanToParquet(filename string, loans []models.Loan) error {
	fw, err := local.NewLocalFileWriter(filename + ".parquet")
	if err != nil {
		return err
	}
	pw, err := writer.NewParquetWriter(fw, &models.Loan{}, 4)
	if err != nil {
		return err
	}
	for i := range loans {
		if err := pw.Write(&loans[i]); err != nil {
			fmt.Println("Write error:", err)
			return err
		}
	}
	if err := pw.WriteStop(); err != nil {
		fmt.Println("WriteStop error:", err)
		return err
	}
	if err := fw.Close(); err != nil {
		return err
	}
	return nil
}
func SaveOrdersToCSV(filename string, orders []models.Orders) error {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"OrderID", "ProductID", "Quantity", "Price", "Status"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write records
	for _, o := range orders {
		record := []string{
			o.OrderID,
			o.ProductID,
			fmt.Sprintf("%d", o.Quantity),
			fmt.Sprintf("%f", o.Price),
			o.CustomerID,
			o.Status,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
func SaveOrdersToParquet(filename string, orders []models.Orders) error {
	fw, err := local.NewLocalFileWriter(filename + ".parquet")
	if err != nil {
		return err
	}
	pw, err := writer.NewParquetWriter(fw, &models.Orders{}, 4)
	if err != nil {
		return err
	}
	for i := range orders {
		if err := pw.Write(&orders[i]); err != nil {
			fmt.Println("Write error:", err)
			return err
		}
	}
	if err := pw.WriteStop(); err != nil {
		fmt.Println("WriteStop error:", err)
		return err
	}
	if err := fw.Close(); err != nil {
		return err
	}
	return nil
}

func SaveReviewsToCSV(filename string, reviews []models.Reviews) error {
	file, err := os.Create(filename + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"ID", "ProductID", "CustomerID", "Rating", "Comment"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write records
	for _, r := range reviews {
		record := []string{
			r.ReviewerName,
			r.ProductID,
			fmt.Sprintf("%d", r.Rating),
			r.Date.Format("2006-01-02"),
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}

func SaveReviewsToParquet(filename string, reviews []models.Reviews) error {
	fw, err := local.NewLocalFileWriter(filename + ".parquet")
	if err != nil {
		return err
	}
	pw, err := writer.NewParquetWriter(fw, &models.Reviews{}, 4)
	if err != nil {
		return err
	}
	for i := range reviews {
		if err := pw.Write(&reviews[i]); err != nil {
			fmt.Println("Write error:", err)
			return err
		}
	}
	if err := pw.WriteStop(); err != nil {
		fmt.Println("WriteStop error:", err)
		return err
	}
	if err := fw.Close(); err != nil {
		return err
	}
	return nil
}
