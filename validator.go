package go_vo_validator


import (
	"reflect"
	"strconv"
	"errors"
)

// Description: validate a given struct and check if the values of its fields
//              are corresponding with the tags values mapped.
// Evaluated tags:
//           'validateMin'
//           'validateMax'
//           'mandatory'
// Returns:
//  - error: Error type

func Validate(obj interface{}) []error {
	var errorsRes []error
	objValue := reflect.ValueOf(obj).Elem()

	for i := 0; i < objValue.NumField(); i++ {

		if len(objValue.Type().Field(i).Tag.Get("validateMin")) > 0{
			errorsRes = append( errorsRes, ValidateMinMaxField(objValue.Type().Field(i), objValue, "validateMin", ValidateMin)...)
		}

		if len(objValue.Type().Field(i).Tag.Get("validateMax")) > 0  {

			errorsRes = append( errorsRes, ValidateMinMaxField(objValue.Type().Field(i), objValue, "validateMax", ValidateMax)...)
		}

		if len(objValue.Type().Field(i).Tag.Get("mandatory")) > 0  {

			errorsRes =  append(errorsRes, ValidateMandatory(objValue.Type().Field(i), objValue)...)
		}
	}
	return errorsRes
}


// Description: validate tags 'validateMin', 'validateMax'
// Params:
// 	- field: field to be evaluated
//  - value: of the struct that is been evaluated
//  - tagName: tag name for evaluation
//  - condition: function that will be run with the conditions code
// Returns:
//  - []error
func ValidateMinMaxField(field reflect.StructField, value reflect.Value, tagName string, condition func(int64, int64, string) error ) []error {
	var errorsRes []error

	min := extractIntValue(tagName, field)

	fieldVal := extractIntFieldValue(field.Name, value)

	conditionErrorRes := condition(fieldVal, min, field.Name)

	if conditionErrorRes != nil {
		errorsRes = append(errorsRes, conditionErrorRes)
	}

	return errorsRes
}

func extractIntFieldValue( fieldName string, value reflect.Value) int64 {
	fieldValue := value.FieldByName(fieldName).Int()
	return fieldValue
}

func extractIntValue(tagName string, field reflect.StructField) int64 {
	val, _ := field.Tag.Lookup(tagName)
	min, _ := strconv.ParseInt(val, 10, 64)
	return min
}


func ValidateMin(fieldValue int64, minCondition  int64, fieldName string) error {
	if fieldValue < minCondition {
		return errors.New("The value for " + fieldName + " can't be less than " + strconv.FormatInt(minCondition, 10))
	}
	return nil
}

func ValidateMax(fieldValue int64, maxCondition  int64, fieldName string) error {
	if fieldValue > maxCondition {
		return errors.New("The value for " + fieldName + " can't be more than " + strconv.FormatInt(maxCondition, 10))
	}
	return nil
}

func ValidateMandatory(value reflect.StructField, val reflect.Value) []error {
	var res []error
	if len(GetValue(val.FieldByName(value.Name))) == 0 {
		res = []error{errors.New("The field "+ value.Name + " is mandatory")}
	}
	return res
}


//Return value string of a given file
func GetValue(fieldValue reflect.Value) string {
	res := ""
	val := fieldValue.Type().Name()

	switch val {
	case "int":
		res = strconv.FormatInt(fieldValue.Int(), 10)
	case "uint":
		res = strconv.FormatUint(fieldValue.Uint(), 10)
	case "string":
		res =  fieldValue.String()
	case "bool":
		res =  strconv.FormatBool(fieldValue.Bool())
	}
	return res
}
