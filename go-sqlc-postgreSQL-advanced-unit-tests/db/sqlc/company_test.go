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
