/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package createHTML

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivHeadline = "Headline"
	ivBody     = "Body"
	ivFooter   = "Footer"
	ovHTML     = "html"
)

var (
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
	activityLog.Info("Executing createHTML activity")
	//Read Inputs
	if context.GetInput(ivHeadline) == nil {
		return false, activity.NewError("Headline string is not configured", "createHTML-4001", nil)
	}
	Headline := context.GetInput(ivHeadline).(string)
	if context.GetInput(ivBody) == nil {
		return false, activity.NewError("Body string is not configured", "createHTML-4002", nil)
	}
	Body := context.GetInput(ivBody).(string)
	if context.GetInput(ivFooter) == nil {
		return false, activity.NewError("Footer string is not configured", "createHTML-4003", nil)
	}
	Footer := context.GetInput(ivFooter).(string)

	//createHTML - Start
	activityLog.Info("Executing createHTML with: " + Headline + " :: " + Body + " :: " + Footer)

	context.SetOutput(ovHTML, Headline+" - "+Body+" - "+Footer)

	//createHTML - End

	return true, nil
}
