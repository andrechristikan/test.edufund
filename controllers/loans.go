package controllers

import (
	"strconv"

	"test.edufund/models"

	"net/http"

	"github.com/gin-gonic/gin"
)

//ArticleController ...
type LoanController struct{}

var LoanModel = new(models.LoanModel)

//Create ...
func (ctrl LoanController) Create(c *gin.Context) {
	
		number := c.PostForm("number")
		amount := c.PostForm("amount")
		tenor := c.PostForm("tenor")
		status := c.PostForm("status")
		borrower := c.PostForm("borrower")

		iNumber, err := strconv.ParseInt(number, 0, 16)
		iAmount, err := strconv.ParseInt(amount, 0, 16)
		iTenor, err := strconv.ParseInt(tenor, 0, 16)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Number, Amount, or Tenor not in number"})
			return
		}

		loanId, err := LoanModel.Create(uint16(iNumber), uint32(iAmount), uint8(iTenor), status, borrower)

		if loanId == 0 && err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something error with your input"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Loan Created", "id": loanId})

}

//All ...
func (ctrl LoanController) All(c *gin.Context) {
		loans, err := LoanModel.All()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Something error with your input"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"loans": loans})
	
}

//Update ...
func (ctrl LoanController) Update(c *gin.Context) {
	
	id := c.Param("id")
	status := c.PostForm("status")

	iId, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Id must in number"})
		return
	}

	loanId, err := LoanModel.Update(iId, status)
	if loanId == 0 && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something error with your input"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan Updated", "id": loanId})

}


//Approve ...
func (ctrl LoanController) Approve(c *gin.Context) {
	
	id := c.Param("id")

	iId, err := strconv.ParseInt(id, 0, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Id must in number"})
		return
	}

	loanId, err := LoanModel.Approve(iId)
	if loanId == 0 && err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Something error with your input"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Loan Approved", "id": loanId})

}
