package createHTML

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivtemplateURL   = "templateURL"
	ivLogoURL       = "LogoURL"
	ivHeadline      = "Headline"
	ivBody          = "Body"
	ivDirectLinkURL = "DirectLinkURL"
	ivFooter        = "Footer"
	ovHTML          = "html"
)

var (
	activityLog = logger.GetLogger("tools-activity-createHTML")
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
	var templateURL = ""
	if context.GetInput(ivtemplateURL) == nil {
		templateURL = ""
	} else {
		templateURL = context.GetInput(ivtemplateURL).(string)
	}
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

	// *** createHTML - Start

	var fullHTML = ""
	if templateURL == "" {
		// *** concat HTML from embedded 'prettyemail.go' building blocks
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
		// backup --> mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		fullHTML = HTMLheader + HTMLhead1 + LogoURL + HTMLhead2 + Headline + HTMLcontent1 + Body + HTMLcontent2 + DirectLinkURL + HTMLcontent3 + Footer + HTMLfoot
	} else {
		// *** create HTML from URL loaded template
		activityLog.Info("Executing createHTML with: " + Headline + " :: " + Body + " :: " + Footer + " :: " + templateURL)
		// load template File
		res, err := http.Get(templateURL)
		if err != nil {
			return false, activity.NewError("template URL not found", "createHTML-4006", nil)
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		var HTMLtemplate = string(body)
		//replace content

		//placeholder {logoURL} {linkURL} {headline} {body} {footer}
		HTMLtemplate = strings.Replace(HTMLtemplate, "{logoURL}", LogoURL, 1)
		HTMLtemplate = strings.Replace(HTMLtemplate, "{linkURL}", DirectLinkURL, 1)
		HTMLtemplate = strings.Replace(HTMLtemplate, "{headline}", Headline, 1)
		HTMLtemplate = strings.Replace(HTMLtemplate, "{body}", Body, 1)
		HTMLtemplate = strings.Replace(HTMLtemplate, "{footer}", Footer, 1)

		//result
		fullHTML = HTMLtemplate
	}
	//context.SetOutput(ovHTML, strings.Replace(fullHTML, "'", "\"", -1))
	context.SetOutput(ovHTML, fullHTML)
	//createHTML - End
	activityLog.Info("done with createHTML")

	return true, nil
}
