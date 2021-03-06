package sendHTMLEmail

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/TIBCOSoftware/flogo-contrib/action/flow/test"
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/stretchr/testify/assert"
)

var activityMetadata *activity.Metadata

func getActivityMetadata() *activity.Metadata {
	if activityMetadata == nil {
		jsonMetadataBytes, err := ioutil.ReadFile("activity.json")
		if err != nil {
			panic("No Json Metadata found for activity.json path")
		}
		activityMetadata = activity.NewMetadata(string(jsonMetadataBytes))
	}

	props, err := ReadPropertiesFile("c:\\GODev\\smalltoolsApp.properties")
	gprops = props
	if err != nil {
		panic("Error while reading properties file")
	}

	return activityMetadata
}

func TestActivityRegistration(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestEval(t *testing.T) {
	act := NewActivity(getActivityMetadata())
	tc := test.NewTestActivityContext(act.Metadata())

	//setup input attrs - Sample is using GMail Settings
	// relace <...> with your connection details for testing

	tc.SetInput("Server", "smtp.gmail.com")
	tc.SetInput("Port", "587")
	tc.SetInput("Sender", gprops["Sender"])
	tc.SetInput("Pass", gprops["Pass"])
	tc.SetInput("To", gprops["To"])
	tc.SetInput("Subject", "TCI WI Test Email")
	tc.SetInput("HTML", "<!DOCTYPE html PUBLIC '-//W3C//DTD XHTML 1.0 Transitional//EN' 'http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd'><html xmlns:v='urn:schemas-microsoft-com:vml'><head><meta http-equiv='Content-Type' content='text/html; charset=UTF-8' /><meta name='viewport' content='width=device-width; initial-scale=1.0; maximum-scale=1.0;' /><meta name='viewport' content='width=600,initial-scale = 2.3,user-scalable=no'><!--[if !mso]><!-- --><link href='https://fonts.googleapis.com/css?family=Work+Sans:300,400,500,600,700' rel='stylesheet'><link href='https://fonts.googleapis.com/css?family=Quicksand:300,400,700' rel='stylesheet'><link href='https://fonts.googleapis.com/icon?family=Material+Icons' rel='stylesheet'><!-- <![endif]--><title>Email Template Sample</title><style type='text/css'>body {width: 100%;background-color: #ffffff;margin: 0;padding: 0; -webkit-font-smoothing: antialiased;mso-margin-top-alt: 0px;mso-margin-bottom-alt: 0px;mso-padding-alt: 0px 0px 0px 0px;}p,h1,h2,h3,h4 {margin-top: 0;margin-bottom: 0;padding-top: 0;padding-bottom: 0;}span.preheader {display: none;font-size: 1px;} html {width: 100%;} table {font-size: 14px;border: 0;}@media only screen and (max-width: 640px) {.goContainer {width: 445px !important;}.section-img img {width: 325px !important;height: auto !important;}}@media only screen and (max-width: 480px) {.goContainer {width: 280px !important;}.section-img img {width: 200px !important;height: auto !important;}}</style><!-- [if gte mso 9]><style type=”text/css”>body {font-family: arial, sans-serif!important;}</style><![endif]--></head><body class='respond' leftmargin='0' topmargin='0' marginwidth='0' marginheight='0'><table border='0' width='100%' cellpadding='0' cellspacing='0' bgcolor='ffffff'><tr><td align='center'><table border='0' align='center' width='590' cellpadding='0' cellspacing='0' class='goContainer'><tr><td height='25' style='font-size: 25px; line-height: 25px;'>&nbsp;</td></tr><tr><td align='center'><table border='0' align='center' width='590' cellpadding='0' cellspacing='0' class='goContainer'><tr><td align='center' height='70' style='height:70px;'><img width='100' border='0' style='display: block; width: 100px;' src='http://www.godev.de/img/tibco-logo.png' alt='' /></td></tr></table></td></tr><tr><td height='25' style='font-size: 25px; line-height: 25px;'>&nbsp;</td></tr></table></td></tr></table><table border='0' width='100%' cellpadding='0' cellspacing='0' bgcolor='ffffff' class='bg_color'><tr><td align='center'><table border='0' align='center' width='590' cellpadding='0' cellspacing='0' class='goContainer'><tr><td align='center' class='section-img'><img src='http://www.godev.de/img/godev.png' style='display: block; width: 285px;' width='590' border='0' alt='' /></td></tr><tr><td height='20' style='font-size: 20px; line-height: 20px;'>&nbsp;</td></tr><tr><td align='center' style='color: #343434; font-size: 24px; font-family: quicksand, Calibri, sans-serif; font-weight:700;letter-spacing: 3px; line-height: 35px;' class='main-header'><div style='line-height: 35px'>TCI WI Test Email</div></td></tr><tr><td height='10' style='font-size: 10px; line-height: 10px;'>&nbsp;</td></tr><tr><td align='center'><table border='0' width='40' align='center' cellpadding='0' cellspacing='0' bgcolor='eeeeee'><tr><td height='2' style='font-size: 2px; line-height: 2px;'>&nbsp;</td></tr></table></td></tr><tr><td height='20' style='font-size: 20px; line-height: 20px;'>&nbsp;</td></tr><tr><td align='center'><table border='0' width='400' align='center' cellpadding='0' cellspacing='0' class='goContainer'><tr><td align='center' style='color: #888888; font-size: 16px; font-family: quicksand, Calibri, sans-serif; line-height: 24px;'><div style='line-height: 24px'>This a GO Test Email coming for GO Testscript execution.</div></td></tr></table></td></tr><tr><td height='25' style='font-size: 25px; line-height: 25px;'>&nbsp;</td></tr><tr><td align='center'><table border='0' align='center' width='160' cellpadding='0' cellspacing='0' bgcolor='3e8ddd' style=''><tr><td height='10' style='font-size: 10px; line-height: 10px;'>&nbsp;</td></tr><tr><td align='center' style='color: #ffffff; font-size: 14px; font-family: quicksand, calibri, sans-serif; line-height: 26px;'><div style='line-height: 26px;'><a href='http://cloud.tibco.com' style='color: #ffffff; text-decoration: none;'>Direct Link</a></div></td></tr><tr><td height='10' style='font-size: 10px; line-height: 10px;'>&nbsp;</td></tr></table></td></tr></table><div align='center' style='color: #888888; font-size: 16px; font-family: quicksand, Calibri, sans-serif; line-height: 24px;' ><br/>your Operation's Team</div></td></tr></table><table border='0' width='100%' cellpadding='0' cellspacing='0' bgcolor='ffffff' class='bg_color'><tr class='hide'><td height='25' style='font-size: 25px; line-height: 25px;'>&nbsp;</td></tr><tr><td height='60' style='border-top: 1px solid #e0e0e0;font-size: 30px; line-height: 30px;'>&nbsp;</td></tr><tr><td align='center'><table border='0' align='center' width='590' cellpadding='0' cellspacing='0' class='goContainer bg_color'><tr><td><table border='0' width='300' align='left' cellpadding='0' cellspacing='0' style='border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt;' class='goContainer'><tr><td align='left'><a href='http://tibco.com' style='display: block; border-style: none !important; border: 0 !important;'><img width='80' border='0' style='display: block; width: 80px;' src='http://www.godev.de/img/tibco-logo.png' alt='' /></a></td></tr><tr><td height='25' style='font-size: 25px; line-height: 25px;'>&nbsp;</td></tr><tr><td align='left' style='color: #888888; font-size: 14px; font-family: calibri, sans-serif; line-height: 23px;' class='text_color'><div style='color: #333333; font-size: 14px; font-family: calibri, sans-serif; font-weight: 600; mso-line-height-rule: exactly; line-height: 23px;'>Contact us: <br/> <a href='https://www.tibco.com' style='color: #888888; display: block; border-style: none !important; border: 0 !important;' ><i class='material-icons'>link</i></a></div></td></tr><tr><td align='left' style='color: #888888; font-size: 14px; font-family: calibri, sans-serif; line-height: 23px;' class='text_color'><div style='color: #333333; font-size: 14px; font-family: calibri, sans-serif; font-weight: 600; mso-line-height-rule: exactly; line-height: 23px;'>Email us: <br/> <a href='mailto:' style='color: #888888; font-size: 14px; font-family: calibri, Sans-serif; font-weight: 400;'>info@tibco.com</a></div></td></tr></table></td></tr></table></td></tr><tr><td height='60' style='font-size: 60px; line-height: 60px;'>&nbsp;</td></tr></table><table border='0' width='100%' cellpadding='0' cellspacing='0' bgcolor='f4f4f4'><tr><td height='25' style='font-size: 25px; line-height: 25px;'>&nbsp;</td></tr><tr><td align='center'><table border='0' align='center' width='590' cellpadding='0' cellspacing='0' class='goContainer'><tr><td><table border='0' align='left' cellpadding='0' cellspacing='0' style='border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt;' class='goContainer'><tr><td align='left' style='color: #aaaaaa; font-size: 14px; font-family: calibri, sans-serif; line-height: 24px;'><div style='line-height: 24px;'><span style='color: #333333;'>TIBCO Cloud Integration</span></div></td></tr></table><table border='0' align='left' width='5' cellpadding='0' cellspacing='0' style='border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt;' class='goContainer'><tr><td height='20' width='5' style='font-size: 20px; line-height: 20px;'>&nbsp;</td></tr></table><table border='0' align='right' cellpadding='0' cellspacing='0' style='border-collapse:collapse; mso-table-lspace:0pt; mso-table-rspace:0pt;' class='goContainer'><tr><td align='center'><table align='center' border='0' cellpadding='0' cellspacing='0'><tr><td align='center'><a style='font-size: 14px; font-family: calibri, sans-serif; line-height: 24px;color: #0082bf; text-decoration: none;font-weight:bold;' href='http://cloud.tibco.com'>TIBCO Cloud</a></td></tr></table></td></tr></table></td></tr></table></td></tr><tr><td height='25' style='font-size: 25px; line-height: 25px;'>&nbsp;</td></tr></table></body></html>")

	_, err := act.Eval(tc)
	assert.Nil(t, err)

	assert.Equal(t, "done", tc.GetOutput("feedback"))

	t.Log(tc.GetOutput("feedback"))
}

//Helper Functions
// read Security Settings from external Propery File
//

type ConfigProperties map[string]string

var gprops ConfigProperties

func ReadPropertiesFile(filepath string) (ConfigProperties, error) {
	config := ConfigProperties{}

	if len(filepath) == 0 {
		return config, nil
	}
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	for scan.Scan() {
		line := scan.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				config[key] = value
			}
		}
	}

	if err := scan.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return config, nil
}
