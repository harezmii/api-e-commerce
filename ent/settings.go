// Code generated by entc, DO NOT EDIT.

package ent

import (
	"api/ent/settings"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Settings is the model entity for the Settings schema.
type Settings struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Keywords holds the value of the "keywords" field.
	Keywords string `json:"keywords,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Company holds the value of the "company" field.
	Company string `json:"company,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Fax holds the value of the "fax" field.
	Fax string `json:"fax,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// MailServerAddress holds the value of the "mailServerAddress" field.
	MailServerAddress string `json:"mailServerAddress,omitempty"`
	// MailServerEmail holds the value of the "mailServerEmail" field.
	MailServerEmail string `json:"mailServerEmail,omitempty"`
	// MailServerPassword holds the value of the "mailServerPassword" field.
	MailServerPassword string `json:"mailServerPassword,omitempty"`
	// MailServerPort holds the value of the "mailServerPort" field.
	MailServerPort string `json:"mailServerPort,omitempty"`
	// Facebook holds the value of the "facebook" field.
	Facebook string `json:"facebook,omitempty"`
	// Instagram holds the value of the "Instagram" field.
	Instagram string `json:"Instagram,omitempty"`
	// Twitter holds the value of the "twitter" field.
	Twitter string `json:"twitter,omitempty"`
	// About holds the value of the "about" field.
	About string `json:"about,omitempty"`
	// Contact holds the value of the "contact" field.
	Contact string `json:"contact,omitempty"`
	// References holds the value of the "references" field.
	References string `json:"references,omitempty"`
	// Status holds the value of the "status" field.
	Status bool `json:"status,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Settings) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case settings.FieldStatus:
			values[i] = new(sql.NullBool)
		case settings.FieldID:
			values[i] = new(sql.NullInt64)
		case settings.FieldTitle, settings.FieldKeywords, settings.FieldDescription, settings.FieldCompany, settings.FieldAddress, settings.FieldPhone, settings.FieldFax, settings.FieldEmail, settings.FieldMailServerAddress, settings.FieldMailServerEmail, settings.FieldMailServerPassword, settings.FieldMailServerPort, settings.FieldFacebook, settings.FieldInstagram, settings.FieldTwitter, settings.FieldAbout, settings.FieldContact, settings.FieldReferences:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Settings", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Settings fields.
func (s *Settings) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case settings.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case settings.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				s.Title = value.String
			}
		case settings.FieldKeywords:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field keywords", values[i])
			} else if value.Valid {
				s.Keywords = value.String
			}
		case settings.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				s.Description = value.String
			}
		case settings.FieldCompany:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field company", values[i])
			} else if value.Valid {
				s.Company = value.String
			}
		case settings.FieldAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field address", values[i])
			} else if value.Valid {
				s.Address = value.String
			}
		case settings.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				s.Phone = value.String
			}
		case settings.FieldFax:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field fax", values[i])
			} else if value.Valid {
				s.Fax = value.String
			}
		case settings.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				s.Email = value.String
			}
		case settings.FieldMailServerAddress:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mailServerAddress", values[i])
			} else if value.Valid {
				s.MailServerAddress = value.String
			}
		case settings.FieldMailServerEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mailServerEmail", values[i])
			} else if value.Valid {
				s.MailServerEmail = value.String
			}
		case settings.FieldMailServerPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mailServerPassword", values[i])
			} else if value.Valid {
				s.MailServerPassword = value.String
			}
		case settings.FieldMailServerPort:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mailServerPort", values[i])
			} else if value.Valid {
				s.MailServerPort = value.String
			}
		case settings.FieldFacebook:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field facebook", values[i])
			} else if value.Valid {
				s.Facebook = value.String
			}
		case settings.FieldInstagram:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field Instagram", values[i])
			} else if value.Valid {
				s.Instagram = value.String
			}
		case settings.FieldTwitter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field twitter", values[i])
			} else if value.Valid {
				s.Twitter = value.String
			}
		case settings.FieldAbout:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field about", values[i])
			} else if value.Valid {
				s.About = value.String
			}
		case settings.FieldContact:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field contact", values[i])
			} else if value.Valid {
				s.Contact = value.String
			}
		case settings.FieldReferences:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field references", values[i])
			} else if value.Valid {
				s.References = value.String
			}
		case settings.FieldStatus:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				s.Status = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Settings.
// Note that you need to call Settings.Unwrap() before calling this method if this Settings
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Settings) Update() *SettingsUpdateOne {
	return (&SettingsClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Settings entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Settings) Unwrap() *Settings {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Settings is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Settings) String() string {
	var builder strings.Builder
	builder.WriteString("Settings(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", title=")
	builder.WriteString(s.Title)
	builder.WriteString(", keywords=")
	builder.WriteString(s.Keywords)
	builder.WriteString(", description=")
	builder.WriteString(s.Description)
	builder.WriteString(", company=")
	builder.WriteString(s.Company)
	builder.WriteString(", address=")
	builder.WriteString(s.Address)
	builder.WriteString(", phone=")
	builder.WriteString(s.Phone)
	builder.WriteString(", fax=")
	builder.WriteString(s.Fax)
	builder.WriteString(", email=")
	builder.WriteString(s.Email)
	builder.WriteString(", mailServerAddress=")
	builder.WriteString(s.MailServerAddress)
	builder.WriteString(", mailServerEmail=")
	builder.WriteString(s.MailServerEmail)
	builder.WriteString(", mailServerPassword=")
	builder.WriteString(s.MailServerPassword)
	builder.WriteString(", mailServerPort=")
	builder.WriteString(s.MailServerPort)
	builder.WriteString(", facebook=")
	builder.WriteString(s.Facebook)
	builder.WriteString(", Instagram=")
	builder.WriteString(s.Instagram)
	builder.WriteString(", twitter=")
	builder.WriteString(s.Twitter)
	builder.WriteString(", about=")
	builder.WriteString(s.About)
	builder.WriteString(", contact=")
	builder.WriteString(s.Contact)
	builder.WriteString(", references=")
	builder.WriteString(s.References)
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", s.Status))
	builder.WriteByte(')')
	return builder.String()
}

// SettingsSlice is a parsable slice of Settings.
type SettingsSlice []*Settings

func (s SettingsSlice) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
