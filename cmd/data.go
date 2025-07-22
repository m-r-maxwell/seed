package cmd

import (
	"fmt"
	helpers "seed/helpers"
	dpkg "seed/pkg"

	"github.com/spf13/cobra"
)

var dataCmd = &cobra.Command{
	Use:   "data",
	Short: "Data commands for generating data",
	Long:  "Data commands for generating data of various types.",
	Args:  cobra.MinimumNArgs(0),
	Example: `seed data --people 10 --products 5 --loan 3 --companies 2 --orders 4 --reviews 6 --file output --format csv
seed data -p 10 -P 5 -l 3 -c 2 -o 4 -r 6 -f output -F parquet`,
	Run: func(cmd *cobra.Command, args []string) {
		filename, _ := cmd.Flags().GetString("file")
		format, _ := cmd.Flags().GetString("format")

		if cmd.Flags().Changed("people") {
			peopleCount, _ := cmd.Flags().GetInt("people")
			ppl := dpkg.GeneratePeopleData(peopleCount)
			fmt.Print("Generated People Data:\n")
			if filename != "" && format != "" {
				switch format {
				case "csv":
					if err := helpers.SavePeopleToCSV(filename, ppl); err != nil {
						fmt.Println("Error saving CSV:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				case "parquet":
					if err := helpers.SavePeopleToParquet(filename, ppl); err != nil {
						fmt.Println("Error saving Parquet:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				default:
					fmt.Println("Unsupported file format. Please use 'csv' or 'parquet'.")
				}
			} else {
				for _, p := range ppl {
					fmt.Printf("Name: %s, Email: %s, Address: %s, Phone: %s\n", p.Name, p.Email, p.Address, p.Phone)
				}
			}
		}

		if cmd.Flags().Changed("products") {
			productsCount, _ := cmd.Flags().GetInt("products")
			prod := dpkg.GenerateProductsData(productsCount)
			fmt.Print("Generated Products Data:\n")
			if filename != "" && format != "" {
				switch format {
				case "csv":
					if err := helpers.SaveProductsToCSV(filename, prod); err != nil {
						fmt.Println("Error saving CSV:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				case "parquet":
					if err := helpers.SaveProductsToParquet(filename, prod); err != nil {
						fmt.Println("Error saving Parquet:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				default:
					fmt.Println("Unsupported file format. Please use 'csv' or 'parquet'.")
				}
			} else {
				for _, p := range prod {
					fmt.Printf("Name: %s, Description: %s, Price: %f, Category: %s\n", p.Name, p.Description, p.Price, p.Category)
				}
			}
		}

		if cmd.Flags().Changed("loan") {
			loanCount, _ := cmd.Flags().GetInt("loan")
			loan := dpkg.GenerateLoanData(loanCount)
			fmt.Print("Generated Loan Data:\n")

			if filename != "" && format != "" {
				switch format {
				case "csv":
					if err := helpers.SaveLoanToCSV(filename, loan); err != nil {
						fmt.Println("Error saving CSV:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				case "parquet":
					if err := helpers.SaveLoanToParquet(filename, loan); err != nil {
						fmt.Println("Error saving Parquet:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				default:
					fmt.Println("Unsupported file format. Please use 'csv' or 'parquet'.")

				}
			} else {
				for _, l := range loan {
					fmt.Printf("Amount: %f, Interest: %f, Term: %d, Default: %s, Gender: %s, Age: %d\n",
						l.Amount, l.Interest, l.Term, l.Default, l.Gender, l.Age)
				}
			}
		}

		if cmd.Flags().Changed("companies") {
			companiesCount, _ := cmd.Flags().GetInt("companies")
			companies := dpkg.GenerateCompaniesData(companiesCount)
			fmt.Print("Generated Companies Data:\n")
			for _, c := range companies {
				fmt.Printf("Name: %s, Address: %s, Phone: %s, Website: %s, Employees: %d, Revenue: %f, Established: %d\n",
					c.Name, c.Address, c.Phone, c.Website, c.Employees, c.Revenue, c.Established)
			}
			if filename != "" && format != "" {
				switch format {
				case "csv":
					if err := helpers.SaveCompaniesToCSV(filename, companies); err != nil {
						fmt.Println("Error saving CSV:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				case "parquet":
					if err := helpers.SaveCompaniesToParquet(filename, companies); err != nil {
						fmt.Println("Error saving Parquet:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				default:
					fmt.Println("Unsupported file format. Please use 'csv' or 'parquet'.")
				}
			} else {
				for _, c := range companies {
					fmt.Printf("Name: %s, Address: %s, Phone: %s, Website: %s, Employees: %d, Revenue: %f, Established: %d\n",
						c.Name, c.Address, c.Phone, c.Website, c.Employees, c.Revenue, c.Established)
				}
			}
		}

		if cmd.Flags().Changed("orders") {
			ordersCount, _ := cmd.Flags().GetInt("orders")
			orders := dpkg.GenerateOrdersData(ordersCount)
			fmt.Print("Generated Orders Data:\n")
			if filename != "" && format != "" {
				switch format {
				case "csv":
					if err := helpers.SaveOrdersToCSV(filename, orders); err != nil {
						fmt.Println("Error saving CSV:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				case "parquet":
					if err := helpers.SaveOrdersToParquet(filename, orders); err != nil {
						fmt.Println("Error saving Parquet:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				default:
					fmt.Println("Unsupported file format. Please use 'csv' or 'parquet'.")

				}
			} else {
				for _, o := range orders {
					fmt.Printf("Order ID: %s, Product: %s, Quantity: %d, Price: %f\n",
						o.OrderID, o.ProductID, o.Quantity, o.Price)
				}
			}
		}

		if cmd.Flags().Changed("reviews") {
			reviewsCount, _ := cmd.Flags().GetInt("reviews")
			reviews := dpkg.GenerateReviewsData(reviewsCount)
			fmt.Print("Generated Reviews Data:\n")
			if filename != "" && format != "" {
				switch format {
				case "csv":
					if err := helpers.SaveReviewsToCSV(filename, reviews); err != nil {
						fmt.Println("Error saving CSV:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				case "parquet":
					if err := helpers.SaveReviewsToParquet(filename, reviews); err != nil {
						fmt.Println("Error saving Parquet:", err)
					} else {
						fmt.Printf("Data saved to %s successfully!\n", filename)
					}
				default:
					fmt.Println("Unsupported file format. Please use 'csv' or 'parquet'.")
				}
			} else {
				for _, r := range reviews {
					fmt.Printf("Reviewer Name: %s, Rating: %d, Product ID: %s, Date: %s\n",
						r.ReviewerName, r.Rating, r.ProductID, r.Date.Format("2006-01-02"))
				}
			}
		}

		// File saving is now handled by flags
	},
}

func init() {
	rootCmd.AddCommand(dataCmd)

	// Add flags for data commands
	dataCmd.Flags().IntP("people", "p", 10, "Number of people records to generate")
	dataCmd.Flags().IntP("products", "P", 10, "Number of products records to generate")
	dataCmd.Flags().IntP("loan", "l", 10, "Number of loan records to generate")
	dataCmd.Flags().IntP("companies", "c", 10, "Number of companies records to generate")
	dataCmd.Flags().IntP("orders", "o", 10, "Number of orders records to generate")
	dataCmd.Flags().IntP("reviews", "r", 10, "Number of reviews records to generate")
	dataCmd.Flags().StringP("file", "f", "", "Filename to save the generated data (optional)")
	dataCmd.Flags().StringP("format", "F", "", "File format to save the data: csv or parquet (optional)")

	// Add help flag
	dataCmd.Flags().BoolP("help", "h", false, "Help for data command")
}
