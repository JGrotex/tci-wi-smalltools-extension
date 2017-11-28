/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package createHTML

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivLogoURL       = "LogoURL"
	ivHeadline      = "Headline"
	ivBody          = "Body"
	ivDirectLinkURL = "DirectLinkURL"
	ivFooter        = "Footer"
	ovHTML          = "html"
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
	if context.GetInput(ivLogoURL) == nil {
		return false, activity.NewError("LogoURL string is not configured", "createHTML-4001", nil)
	}
	LogoURL := context.GetInput(ivLogoURL).(string)
	if context.GetInput(ivHeadline) == nil {
		return false, activity.NewError("Headline string is not configured", "createHTML-4002", nil)
	}
	Headline := context.GetInput(ivHeadline).(string)
	if context.GetInput(ivBody) == nil {
		return false, activity.NewError("Body string is not configured", "createHTML-4003", nil)
	}
	Body := context.GetInput(ivBody).(string)
	if context.GetInput(ivDirectLinkURL) == nil {
		return false, activity.NewError("DirectLinkURL string is not configured", "createHTML-4004", nil)
	}
	DirectLinkURL := context.GetInput(ivDirectLinkURL).(string)
	if context.GetInput(ivFooter) == nil {
		return false, activity.NewError("Footer string is not configured", "createHTML-4005", nil)
	}
	Footer := context.GetInput(ivFooter).(string)

	//createHTML - Start
	activityLog.Info("Executing createHTML with: " + Headline + " :: " + Body + " :: " + Footer)

	if LogoURL == "" {
		LogoURL = "" // no special Logo in EMail
	} else {
		LogoURL = strings.Replace(HTMLlogo, "{logoURL}", LogoURL, 1)
	}

	if DirectLinkURL == "" {
		DirectLinkURL = "" // no direct Link in EMail
	} else {
		DirectLinkURL = strings.Replace(HTMLdirectLink, "{linkURL}", DirectLinkURL, 1)
	}

	//mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	fullHTML := mime + HTMLheader + HTMLhead1 + LogoURL + HTMLhead2 + Headline + HTMLcontent1 + Body + HTMLcontent2 + DirectLinkURL + HTMLcontent3 + Footer + HTMLfoot

	context.SetOutput(ovHTML, strings.Replace(fullHTML, "'", "\"", -1))
	//createHTML - End

	return true, nil
}
