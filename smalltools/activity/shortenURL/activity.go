/*
 * Copyright © 2017. TIBCO Software Inc.
 * This file is subject to the license terms contained
 * in the license file that is distributed with this file.
 */
package shortenURL

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)

const (
	ivlongURL  = "longURL"
	ovshortURL = "shortURL"
)

var activityLog = logger.GetLogger("tools-activity-shortenURL")

//shortenURLActivity Metadata
type shortenURLActivity struct {
	metadata *activity.Metadata
}

func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &shortenURLActivity{metadata: metadata}
}

func (a *shortenURLActivity) Metadata() *activity.Metadata {
	return a.metadata
}
func (a *shortenURLActivity) Eval(context activity.Context) (done bool, err error) {
	activityLog.Info("Executing shortenURL activity")
	//Read Inputs
	if context.GetInput(ivlongURL) == nil {
		// First string is not configured
		// return error to the engine
		return false, activity.NewError("long URL string is not configured", "shortenURL-4001", nil)
	}
	longURL := context.GetInput(ivlongURL).(string)

	//execution

	//developerKey := "<your api key>"

	// **** Sample post on Google API, ... to be implemented later!
	/*
	   	POST https://www.googleapis.com/urlshortener/v1/url?fields=id%2ClongUrl&key={YOUR_API_KEY}
	     {
	   	"longUrl": "http://www.GODev.de"
	      }
	*/

	context.SetOutput(ovshortURL, longURL)
	//context.SetOutput(ovshortURL, url.Id)

	return true, nil
}
