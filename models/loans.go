package models

import (
	"test.edufund/db"
	"time"
)



type LoanModel struct{}

//Create ...
func (m LoanModel) Create(number uint16, amount uint32, tenor uint8, status string, borrower string) (loanId uint, err error) {
	loan := db.Loan{
		Number: number,
		Amount: amount,
		Tenor: tenor,
		Status: status,
		Borrower: borrower,
	}


	result := db.GetDB().Create(&loan);

	return loan.ID, result.Error
}

// All
func (m LoanModel) All() (results []db.Loan, err error) {
	var loans []db.Loan
	result := db.GetDB().Find(&loans)
	return loans, result.Error
}

//Update ...
func (m LoanModel) Update(id int64, status string) (loanId uint, err error) {
	var query db.Loan
	result := db.GetDB().First(&query, id)

	query.Status = status;
	db.GetDB().Save(&query)
	return query.ID, result.Error
}

//Approve ...
func (m LoanModel) Approve(id int64) (loanId uint, err error) {
	var query db.Loan
	result := db.GetDB().First(&query, id)

	query.ApproveDate = time.Now();
	db.GetDB().Save(&query)
	return query.ID, result.Error
}