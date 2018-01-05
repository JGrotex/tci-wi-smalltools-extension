# TCI WI SmallTools Extension

[![Go Report Card](https://goreportcard.com/badge/github.com/JGrotex/tci-wi-smalltools-extension)](https://goreportcard.com/report/github.com/JGrotex/tci-wi-smalltools-extension) [![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

first Version with just an own version of a Concat Activity, more to come ... e.g. will add as soon as possible a Email Validation Activity. Attached ZIP contains the first release v.1.2 and can just uploaded under TIBCO Cloud Integration Extensions

This is just the start.

## Activities
available Activities so far
### Concat
This activity is just using GO, and no UI customization using TypeScript, etc.
Just to show how simple a Implemenation could be.

Input
- string1 (String)
- string2 (String)
- Seperator (String, one of ";","-","+","_","|" Default is "-") part of Configuration Dialog 

Output
- result (String) as full String

### EMail Addr Validation
validates an EMail Addr

Input
- Email Addr (String)

Output
- valid (Boolean) just Format check

### create HTML
Tool creates a pretty HTML page to store e.g. into an Email body of an HTML SMTP EMail Sender.

Input
- LogoURL (String)
- Headline (String)
- Body (String)
- DirectLinkURL (String)
- Footer (String)

Output
- valid (Boolean) just Format check

Example HTML string content as Screenshot

![Pretty Email image](screenshots/prettyHTMLMail.png?raw=true "TCI WI Pretty Email Screenshot")

### send HTML Email
send HTML Emails with the Content of of the create HTML Activity, fully tested with Google SMTP Mail.
Implementation is limited to what is realy needed to send a single Notification EMail. 

Input
- Server (String) part of Configuration Dialog : default "smtp.gmail.com"
- Port (String) part of Configuration Dialog : default "587"
- Sender (String) part of Configuration Dialog 
- Pass (Password) part of Configuration Dialog 
- To (String)
- Subject (String)
- HTML (String)

Output
- feedback (String) on success always "done" so far

<hr>
<sub><b>Note:</b> more TCI Extensions can be found here: https://tibcosoftware.github.io/tci-awesome/ </sub>


