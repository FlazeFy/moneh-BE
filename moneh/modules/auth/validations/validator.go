package validations

import (
	"moneh/modules/auth/models"
	"moneh/packages/helpers/converter"
	"moneh/packages/helpers/generator"
	"moneh/packages/utils/validator"
	"strconv"
)

func GetValidateRegister(body models.UserRegister) (bool, string) {
	var msg = ""
	var status = true

	// Rules
	minUname, maxUname := validator.GetValidationLength("username")
	minPass, maxPass := validator.GetValidationLength("password")
	minEmail, maxEmail := validator.GetValidationLength("email")
	minFName, maxFName := validator.GetValidationLength("first_name")
	_, maxLName := validator.GetValidationLength("last_name")
	minValidUntil, maxValidUntil := validator.GetValidationLength("valid_until")

	// Value
	uname := converter.TotalChar(body.Username)
	pass := converter.TotalChar(body.Password)
	email := converter.TotalChar(body.Email)
	fname := converter.TotalChar(body.FirstName)
	lname := converter.TotalChar(body.LastName)
	valid, _ := strconv.Atoi(body.ValidUntil)

	// Validate
	if uname <= minUname || uname >= maxUname {
		status = false
		msg += generator.GenerateValidatorMsg("Username", minUname, maxUname)
	}
	if pass <= minPass || pass >= maxPass {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Password", minPass, maxPass)
	}
	if email <= minEmail || email >= maxEmail {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Email", minEmail, maxEmail)
	}
	if fname <= minFName || fname >= maxFName {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("First name", minFName, maxFName)
	}
	if lname >= maxLName {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Last name", 0, maxFName)
	}
	if valid <= minValidUntil || valid >= maxValidUntil {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Valid until", minValidUntil, maxValidUntil)
	}

	if status {
		return status, "Validation success"
	} else {
		return status, msg
	}
}

func GetValidateLogin(username, password string) (bool, string) {
	var msg = ""
	var status = true

	// Rules
	minUname, maxUname := validator.GetValidationLength("username")
	minPass, maxPass := validator.GetValidationLength("password")

	// Value
	uname := converter.TotalChar(username)
	pass := converter.TotalChar(password)

	// Validate
	if uname <= minUname || uname >= maxUname {
		status = false
		msg += generator.GenerateValidatorMsg("Username", minUname, maxUname)
	}
	if pass <= minPass || pass >= maxPass {
		status = false
		if msg != "" {
			msg += ", "
		}
		msg += generator.GenerateValidatorMsg("Password", minPass, maxPass)
	}

	if status {
		return status, "Validation success"
	} else {
		return status, msg
	}
}
