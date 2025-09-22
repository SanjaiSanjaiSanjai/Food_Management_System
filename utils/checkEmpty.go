package utils

import (
	customlogger "Food_Delivery_Management/HandleCustomLogger"
	"fmt"
)

// Error
// common function check value is not nil return error
func IsNotNilError(value error, coming string, message string) {
	fmt.Println("token error: ", value)
	if value != nil {
		error_Msg := fmt.Sprintf("[%s]: %s", coming, message)
		customlogger.Log.Error(error_Msg)
		return
	}
}

// Success
// common function check value is not nil return success
func IsNotNilSuccess(value any, coming string, message string) {
	if value != nil {
		success_Msg := fmt.Sprintf("[%s]: %s", coming, message)
		customlogger.Log.Info(success_Msg)
		return
	}
}

// Success
// common function check value is nil return success
func IsNillSuccess(value any, coming string, message string) {
	if value == nil {
		success_Msg := fmt.Sprintf("[%s]: %s", coming, message)
		customlogger.Log.Info(success_Msg)
	}
}
