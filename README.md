<div style="text-align: center">
  <h1>Seed</h1>
 </div>

 <div style="text-align: center">
<img src="seed.png" height="250" width="250">
 </div>

# Overview

Seed is a CLI tool for generating realistic test data for people, products, loans, companies, orders, and reviews. The generated data can be printed to the console or saved to CSV or Parquet files for use in development, testing, or analytics.


# Example Usage

## Print generated data to console
```sh
seed data --people 5 --products 3
```

## Save generated data to CSV
```sh
seed data --people 10 --file people --format csv
```
This will create `people.csv` in the current directory.

## Save generated data to Parquet
```sh
seed data --people 10 --file people --format parquet
```
This will create `people.parquet` in the current directory.

## Generate multiple types of data
```sh
seed data --people 5 --products 5 --loan 2 --companies 2 --file all_data --format csv
```

# Supported Flags
- `--people, -p` Number of people records to generate
- `--products, -P` Number of products records to generate
- `--loan, -l` Number of loan records to generate
- `--companies, -c` Number of companies records to generate
- `--orders, -o` Number of orders records to generate
- `--reviews, -r` Number of reviews records to generate
- `--file, -f` Filename to save the generated data (optional)
- `--format, -F` File format to save the data: `csv` or `parquet` (optional)

# Requirements
- Go 1.18+
- [parquet-go](https://github.com/xitongsys/parquet-go) for Parquet support
