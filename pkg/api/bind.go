package api

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type Binder interface {
	BindAndValidate(c *gin.Context, dto any) error
}

type DefaultBinder struct {
	validate *validator.Validate
}

func NewDefaultBinder(validate *validator.Validate) *DefaultBinder {
	binding.Validator = nil
	if validate == nil {
		validate = validator.New()
	}
	validate.SetTagName("binding")
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		for _, tagName := range []string{"json", "form", "uri"} {
			name := strings.TrimSpace(strings.Split(field.Tag.Get(tagName), ",")[0])
			if name != "" && name != "-" {
				return name
			}
		}
		return field.Name
	})
	return &DefaultBinder{validate: validate}
}

var defaultBinder Binder = NewDefaultBinder(validator.New())

func BindAndValidate(c *gin.Context, dto any) error {
	return defaultBinder.BindAndValidate(c, dto)
}

func (b *DefaultBinder) BindAndValidate(c *gin.Context, dto any) error {
	var bindErr error
	if c.Request.Method == http.MethodGet || c.Request.Method == http.MethodDelete {
		bindErr = c.ShouldBindQuery(dto)
	} else {
		bindErr = c.ShouldBind(dto)
	}
	if bindErr != nil {
		return newBindError(dto, c, bindErr)
	}

	if len(c.Params) > 0 && hasURITag(dto) {
		if err := c.ShouldBindUri(dto); err != nil {
			return newBindError(dto, c, err)
		}
	}

	applySetDefaults(dto)
	if err := b.validate.Struct(dto); err != nil {
		return newBindError(dto, c, err)
	}
	return nil
}

type bindError struct {
	fields []FieldError
}

func (e *bindError) Error() string {
	return "request binding or validation failed"
}

func BindErrorFields(err error) []FieldError {
	var bindingErr *bindError
	if errors.As(err, &bindingErr) {
		return bindingErr.fields
	}
	return nil
}

func newBindError(dto any, c *gin.Context, err error) error {
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		fields := make([]FieldError, 0, len(validationErrors))
		for _, validationErr := range validationErrors {
			field := validationErr.Field()
			message := fmt.Sprintf("%s is invalid", field)
			switch validationErr.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required", field)
			case "min":
				message = fmt.Sprintf("%s must be at least %s", field, validationErr.Param())
			case "max":
				message = fmt.Sprintf("%s must be at most %s", field, validationErr.Param())
			case "oneof":
				message = fmt.Sprintf("%s must be one of: %s", field, validationErr.Param())
			}
			fields = append(fields, FieldError{Field: field, Message: message})
		}
		return &bindError{fields: fields}
	}

	if fields := numericQueryTypeMismatch(c, dto); len(fields) > 0 {
		return &bindError{fields: fields}
	}
	return &bindError{fields: []FieldError{{Field: "request", Message: "invalid request format"}}}
}

func applySetDefaults(dto any) {
	if dto == nil {
		return
	}
	if defaultSetter, ok := dto.(interface{ SetDefaults() }); ok {
		defaultSetter.SetDefaults()
		return
	}

	value := reflect.ValueOf(dto)
	if value.Kind() == reflect.Pointer {
		if value.IsNil() {
			return
		}
		value = value.Elem()
	}
	if value.Kind() == reflect.Struct && value.CanAddr() {
		if defaultSetter, ok := value.Addr().Interface().(interface{ SetDefaults() }); ok {
			defaultSetter.SetDefaults()
		}
	}
}

func hasURITag(dto any) bool {
	typeOfDTO := reflect.TypeOf(dto)
	if typeOfDTO == nil {
		return false
	}
	if typeOfDTO.Kind() == reflect.Pointer {
		typeOfDTO = typeOfDTO.Elem()
	}
	if typeOfDTO.Kind() != reflect.Struct {
		return false
	}
	for index := 0; index < typeOfDTO.NumField(); index++ {
		field := typeOfDTO.Field(index)
		if field.Anonymous {
			if hasURITag(reflect.New(field.Type).Elem().Interface()) {
				return true
			}
			continue
		}
		uriTag := strings.TrimSpace(strings.Split(field.Tag.Get("uri"), ",")[0])
		if uriTag != "" && uriTag != "-" {
			return true
		}
	}
	return false
}

func numericQueryTypeMismatch(c *gin.Context, dto any) []FieldError {
	fields := make([]FieldError, 0)
	collectNumericQueryTypeMismatch(c, reflect.TypeOf(dto), &fields)
	return fields
}

func collectNumericQueryTypeMismatch(c *gin.Context, typeOfDTO reflect.Type, fields *[]FieldError) {
	if typeOfDTO == nil {
		return
	}
	if typeOfDTO.Kind() == reflect.Pointer {
		typeOfDTO = typeOfDTO.Elem()
	}
	if typeOfDTO.Kind() != reflect.Struct {
		return
	}
	for index := 0; index < typeOfDTO.NumField(); index++ {
		field := typeOfDTO.Field(index)
		if field.Anonymous {
			collectNumericQueryTypeMismatch(c, field.Type, fields)
			continue
		}
		if !isIntegerKind(field.Type.Kind()) {
			continue
		}
		formTag := strings.TrimSpace(strings.Split(field.Tag.Get("form"), ",")[0])
		if formTag == "" || formTag == "-" {
			continue
		}
		value := c.Query(formTag)
		if value == "" {
			continue
		}
		if _, err := strconv.ParseInt(value, 10, 64); err != nil {
			*fields = append(*fields, FieldError{Field: formTag, Message: formTag + " must be a number"})
		}
	}
}

func isIntegerKind(kind reflect.Kind) bool {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return true
	default:
		return false
	}
}
