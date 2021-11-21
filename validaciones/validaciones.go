package validaciones

import (
	"errors"
	"fmt"

	"gopkg.in/go-playground/validator.v9"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate = validator.New()

func ValidateStruct(s interface{}) error { //c Cuenta)

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(s)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return errors.New("ERROR")
		}
		//--------

		var msgEror string
		fmt.Println("ERROR: ")
		for _, err := range err.(validator.ValidationErrors) {

			msgEror = err.Namespace() + " " + err.Tag() + "=" + err.Param()

			fmt.Println("err.Namespace(): " + err.Namespace())
			fmt.Println("err.Field(): " + err.Field())
			fmt.Println("err.StructNamespace(): " + err.StructNamespace())
			fmt.Println("err.StructField(): " + err.StructField())
			fmt.Println("err.Tag(): " + err.Tag())
			fmt.Println("err.ActualTag(): " + err.ActualTag())
			fmt.Println("err.Kind(): ")
			fmt.Println(err.Kind())
			fmt.Println()
			fmt.Println("err.Type(): ")
			fmt.Println(err.Type())
			fmt.Println()
			fmt.Println("err.Value(): ")
			fmt.Println(err.Value())
			fmt.Println()
			fmt.Println("err.Param(): ")
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		//return "ERROR DE VALIDACION EN CAMPO " + msgEror
		return errors.New("ERROR DE VALIDACION EN CAMPO " + msgEror)
	}
	fmt.Println("OK - VALIDADO Cuenta")
	//return "OK"
	return nil
}
