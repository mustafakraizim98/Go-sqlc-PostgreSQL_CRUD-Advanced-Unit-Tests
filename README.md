# Go-sqlc-PostgreSQL_CRUD-Advanced-Unit-Tests
![maxresdefault](https://user-images.githubusercontent.com/113289516/207132003-cb96d714-04c4-4e0b-a98c-05964feb7ad6.jpg)

# Prerequisites
- [Go-PostgreSQL-Database-Migration](https://github.com/mustafakraizim98/Go-PostgreSQL-Database-Migration) // An Important Pre-Step Should be done.
- Go - [latest version](https://go.dev/dl/)
- Docker
- PostgreSQL - // Pull a Docker Image Preferred Option.
- sqlc - [Installing sqlc](https://docs.sqlc.dev/en/latest/overview/install.html#installing-sqlc) - /* Important */
- MinGW Makefile - // Optional
- TablePlus - // Optional
- [dbdiagram.io](https://dbdiagram.io/) - // Optional >>> Simple tool to draw ER diagrams by just writing code.

![dbdiagram_go_sqlc_postgresql_unit_tests](https://user-images.githubusercontent.com/113289516/207153433-2abe3ccd-89e1-412c-a043-5dfea6c3d1fb.png)

## What is sqlc?
### sqlc: A SQL Compiler
sqlc generates type-safe code from SQL. Here's how it works:
1. You write queries in SQL.
2. You run sqlc to generate code with type-safe interfaces to those queries.
3. You write application code that calls the generated code.

[sqlc - Getting started with PostgreSQL](https://docs.sqlc.dev/en/latest/tutorials/getting-started-postgresql.html)

### Things to Consider About sqlc:
- Very fast & easy to use.
- Automatic code generation.
- Catch SQL query errors before generating codes.
- Full support of PostgreSQL.

# Unit Tests in Go
A unit test is a function that tests a specific piece of code from a program or package. The job of unit tests is to check the correctness of an application, and they are a crucial part of the Go programming language.

NOTE: Go testing files are always located in the same folder, or package, where the code they are testing resides.

And as with everything in Go, the language is opinionated about testing. The Go language provides a minimal yet complete package called testing that developers use alongside the go test command. The testing package provides some useful conventions, such as coverage tests and benchmarks.

## Write Unit Test
[Testify - Thou Shalt Write Tests](https://github.com/stretchr/testify#testify---thou-shalt-write-tests)
: Go code (golang) set of packages that provide many tools for testifying that our code will behave as we intend.

In company_test.go file for our project, here are our CRUD unit tests implemented.
```
package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/go-sqlc-postgresql-advanced-unit-tests/db/util"
	"github.com/stretchr/testify/require"
)

func createRandomCompany(t *testing.T) Company {
	arg := CreateCompanyParams{
		Owner:        util.RandomOwner(),
		Headquarters: util.RandomHeadquarters(),
		Founded:      util.RandomFoundationYear(),
	}

	company, err := testQueries.CreateCompany(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, company)

	require.Equal(t, arg.Owner, company.Owner)
	require.Equal(t, arg.Headquarters, company.Headquarters)
	require.Equal(t, arg.Founded, company.Founded)

	require.NotZero(t, company.ID)
	require.NotZero(t, company.CreatedAt)

	return company
}

func TestCreateCompany(t *testing.T) {
	createRandomCompany(t)
}

func TestGetCompany(t *testing.T) {
	company1 := createRandomCompany(t)
	company2, err := testQueries.GetCompany(context.Background(), company1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, company2)

	require.Equal(t, company1.ID, company2.ID)
	require.Equal(t, company1.Owner, company2.Owner)
	require.Equal(t, company1.Headquarters, company2.Headquarters)
	require.Equal(t, company1.Founded, company2.Founded)
	require.WithinDuration(t, company1.CreatedAt, company2.CreatedAt, time.Second)
}

func TestUpdateCompany(t *testing.T) {
	company1 := createRandomCompany(t)

	arg := UpdateCompanyParams{
		ID:           company1.ID,
		Headquarters: util.RandomHeadquarters(),
	}

	company2, err := testQueries.UpdateCompany(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, company2)

	require.Equal(t, company1.ID, company2.ID)
	require.Equal(t, company1.Owner, company2.Owner)
	require.Equal(t, arg.Headquarters, company2.Headquarters)
	require.Equal(t, company1.Founded, company2.Founded)
	require.WithinDuration(t, company1.CreatedAt, company2.CreatedAt, time.Second)
}

func TestDeleteCompany(t *testing.T) {
	company1 := createRandomCompany(t)
	err := testQueries.DeleteCompany(context.Background(), company1.ID)
	require.NoError(t, err)

	company2, err := testQueries.GetCompany(context.Background(), company1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, company2)
}

func TestListCompanies(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomCompany(t)
	}

	// with these parameters we expects to get 3 records
	arg := ListCompaniesParams{
		Limit:  3, // returns 3 records based-on the offset value or without offset clause it will returns 3 rows from the first row returned by the SELECT clause.
		Offset: 7, // skip the first 7 records
	}

	companies, err := testQueries.ListCompanies(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, companies, 3)

	for _, company := range companies {
		require.NotEmpty(t, company)
	}
}
```

## Writing Coverage Tests in Go

```go test``` command based on our project where files are located inside ```./db/sdlc/``` directory:
```
go test -v -cover ./...
```

In this step, we will test our code. ```go test``` is a powerful subcommand that helps us automate our tests. ```go test``` accepts different flags that can configure the tests we wish to run, how much verbosity ``` -v``` the tests return, and more. When writing tests, it is often important to know how much of our actual code the tests cover. This is generally referred to as coverage ``` -cover```.

We will receive the following output:
![coverage_go_sqlc_postgresql_unit_tests_result](https://user-images.githubusercontent.com/113289516/207152691-d62a403c-4507-440d-9902-8f1b21e26624.png)

```PASS``` means the code is working as expected. When a test fails, we will see ```FAIL```.

# Sources to follow for detailed information: 
- [Go Official Documentation - Add a test](https://go.dev/doc/tutorial/add-a-test)
- [Medium - How to Write Unit Test in Go](https://medium.com/yemeksepeti-teknoloji/how-to-write-unit-test-in-go-1df2b98ad510)
- [DigitalOcean - How To Write Unit Tests in Go](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)
- [LogRocket - A deep dive into unit testing in Go](https://blog.logrocket.com/a-deep-dive-into-unit-testing-in-go/)
