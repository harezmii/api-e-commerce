package validate

import (
	"github.com/go-playground/locales/en"
	_ "github.com/go-playground/locales/tr"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	tr2 "github.com/go-playground/validator/v10/translations/tr"
)

func ValidateStructToTurkish(i interface{}) []string {
	var s []string
	validate := validator.New()

	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("tr")
	_ = tr2.RegisterDefaultTranslations(validate, trans)
	errr := validate.Struct(i)

	if errr != nil {
		for _, err := range errr.(validator.ValidationErrors) {
			a := err.Translate(trans)
			s = append(s, a)
		}
	}
	return s
}
