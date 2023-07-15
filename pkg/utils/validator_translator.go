package utils

import (
	"bytes"
	"errors"
	"fmt"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations_en "github.com/go-playground/validator/v10/translations/en"
	"log"
	"reflect"
	"strings"
)

func ValidateRequest(payload any) error {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	en := en2.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")
	err := translations_en.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Println("new validator: ", err)
		return nil
	}
	errs := make(map[string][]string)
	err = validate.Struct(payload)
	if err != nil {
		var errVals validator.ValidationErrors
		if errors.As(err, &errVals) {
			for i, _ := range errVals {
				errs[errVals[i].Field()] = []string{errVals[i].Translate(trans)}
			}
		}
		b := new(bytes.Buffer)
		var j int
		for key, value := range errs {
			fmt.Fprintf(b, "%s: ", key)
			for _, item := range value {
				if j == len(errs)-1 {
					fmt.Fprintf(b, "%s", item)
				} else {
					fmt.Fprintf(b, "%s, ", item)
				}
			}
			j = j + 1
		}
		return errors.New(b.String())
	}
	return nil
}
