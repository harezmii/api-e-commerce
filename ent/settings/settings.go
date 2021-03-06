// Code generated by entc, DO NOT EDIT.

package settings

const (
	// Label holds the string label denoting the settings type in the database.
	Label = "settings"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldKeywords holds the string denoting the keywords field in the database.
	FieldKeywords = "keywords"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldCompany holds the string denoting the company field in the database.
	FieldCompany = "company"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldFax holds the string denoting the fax field in the database.
	FieldFax = "fax"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldMailServerAddress holds the string denoting the mailserveraddress field in the database.
	FieldMailServerAddress = "mail_server_address"
	// FieldMailServerEmail holds the string denoting the mailserveremail field in the database.
	FieldMailServerEmail = "mail_server_email"
	// FieldMailServerPassword holds the string denoting the mailserverpassword field in the database.
	FieldMailServerPassword = "mail_server_password"
	// FieldMailServerPort holds the string denoting the mailserverport field in the database.
	FieldMailServerPort = "mail_server_port"
	// FieldFacebook holds the string denoting the facebook field in the database.
	FieldFacebook = "facebook"
	// FieldInstagram holds the string denoting the instagram field in the database.
	FieldInstagram = "instagram"
	// FieldTwitter holds the string denoting the twitter field in the database.
	FieldTwitter = "twitter"
	// FieldAbout holds the string denoting the about field in the database.
	FieldAbout = "about"
	// FieldContact holds the string denoting the contact field in the database.
	FieldContact = "contact"
	// FieldReferences holds the string denoting the references field in the database.
	FieldReferences = "references"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// Table holds the table name of the settings in the database.
	Table = "settings"
)

// Columns holds all SQL columns for settings fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldKeywords,
	FieldDescription,
	FieldCompany,
	FieldAddress,
	FieldPhone,
	FieldFax,
	FieldEmail,
	FieldMailServerAddress,
	FieldMailServerEmail,
	FieldMailServerPassword,
	FieldMailServerPort,
	FieldFacebook,
	FieldInstagram,
	FieldTwitter,
	FieldAbout,
	FieldContact,
	FieldReferences,
	FieldStatus,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
