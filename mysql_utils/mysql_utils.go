package mysql_utils

import (
	"strings"

	"github.com/rampo0/go-utils/rest_error"

	"github.com/go-sql-driver/mysql"
)

const (
	ErrorNoRows = "no rows in result set"
)

func ParseError(err error) *rest_error.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), ErrorNoRows) {
			return rest_error.NewNotFoundError("no record match given id")
		}
		return rest_error.NewInternalServerError("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return rest_error.NewBadRequestError("duplicated records")
	}
	return rest_error.NewInternalServerError("error processing request")
}
