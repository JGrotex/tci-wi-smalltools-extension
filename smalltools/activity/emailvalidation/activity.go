package emailvalidation

import (
	"errors"
	"regexp"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivEmail = "email"
	ovValid = "valid"
)

var (
	//emailRegexp restrictive Regexp
	emailRegexp = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	//ErrBadFormat Email Format Error
	ErrBadFormat = errors.New("invalid format")

	activityLog = logger.GetLogger("tools-activity-emailvalidation")
)

type EmailValidationActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &EmailValidationActivity{metadata: metadata}
}

func (a *EmailValidationActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *EmailValidationActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing emailvalidation activity")
	//Read Inputs
	if context.GetInput(ivEmail) == nil {
		// Email string is not configured
		// return error to the engine
		return false, activity.NewError("Email string is not configured", "emailvalidation-4001", nil)
	}
	email := context.GetInput(ivEmail).(string)

	//Validation - Start
	//activityLog.Info("Executing emailvalidation for Email: " + email)
	EmailErr := ValidateFormat(email)
	if EmailErr != nil {
		context.SetOutput(ovValid, false)
	} else {
		context.SetOutput(ovValid, true)
	}
	//Validation - End

	return true, nil
}

func ValidateFormat(email string) error {
	if !emailRegexp.MatchString(email) {
		return ErrBadFormat
	}
	return nil
}

func split(email string) (account, host string) {
	i := strings.LastIndexByte(email, '@')
	account = email[:i]
	host = email[i+1:]
	return
}
