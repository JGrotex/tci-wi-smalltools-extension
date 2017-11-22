# TCI WI SmallTools Extension
first Version with just an own version of a Concat Activity, more to come ...
e.g. will add as soon as possible a Email Validation Activity.

Attached ZIP contains the first release v.1.2 and can just uploaded under 
TIBCO Cloud Integration Extensions

This is just the start.

## Activities
available Activities so far
### Concat
This activity is just using GO, and no UI customization using TypeScript, etc.
Just to show how simple a Implemenation could be.

Input
- string1 (String)
- string2 (String)
- Seperator (String, one of ";","-","+","_","|" Default is "-")

Output
- result (String) as full String

### EMail Addr Validation
Just planed, ... not Implemented yet, ... but it will look like follows

Input
- Email Addr (String)

Output
- valid (Boolean) overall Status Summary
- validFormat (Boolean)
- validHost (Boolean)
- validUser (Boolean)
 
