package shared

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/pocketbase/pocketbase/tools/router"
)

var v = validator.New()

// This function helps to get the json field name from the struct tag
func init() {
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func ValidateRequest(dto interface{}) []string {

	if err := validator.New().Struct(dto); err != nil {
		errors := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, generateErrorString(err))
		}
		return errors
	}
	return nil
}

func generateErrorString(err validator.FieldError) string {
	filed, tag, param := err.Field(), err.Tag(), err.Param()
	switch tag {
	case "required":
		return filed + " value is required"
	case "min":
		return filed + " value must be greater than " + param
	case "max":
		return filed + " value must be less than " + param
	case "oneof":
		return filed + " value is must be one of asc or desc"
	default:
		return filed + " value is invalid"
	}
}

type safeErrorItem struct {
	Tag string
	Msg string
}

func (s safeErrorItem) Code() string  { return s.Tag }
func (s safeErrorItem) Error() string { return s.Msg }

/* It is mandatory to return map[string]router.SafeErrorItem because pb response only */
/* will parse it if it is a map of router.SafeErrorItem */
func ValidateRequestV2(dto interface{}) map[string]router.SafeErrorItem {

	if err := v.Struct(dto); err != nil {
		errors := make(map[string]router.SafeErrorItem)
		for _, ve := range err.(validator.ValidationErrors) {
			field := ve.Field()
			errors[field] = generateSafeErrorItem(ve)
		}
		return errors
	}
	return nil
}

func generateSafeErrorItem(err validator.FieldError) safeErrorItem {
	field, tag, param := err.Field(), err.Tag(), err.Param()

	switch tag {
	case "required":
		return safeErrorItem{
			Tag: "validation_required",
			Msg: field + " value is required",
		}
	case "min":
		return safeErrorItem{
			Tag: "validation_min",
			Msg: field + " value must be greater than " + param,
		}
	case "max":
		return safeErrorItem{
			Tag: "validation_max",
			Msg: field + " value must be less than " + param,
		}
	case "oneof":
		return safeErrorItem{
			Tag: "validation_oneof",
			Msg: field + " value must be one of asc or desc",
		}
	default:
		return safeErrorItem{
			Tag: "validation_invalid_value",
			Msg: field + " value is invalid",
		}
	}
}
