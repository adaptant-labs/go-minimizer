package minimizers

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
)

var (
	tagName = "minimize"
)

func minimizer(l MinimizationLevel, s interface{}, parent *reflect.Value) error {
	val := reflect.Indirect(reflect.ValueOf(s))
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	kind := val.Kind()
	if kind != reflect.Struct {
		return fmt.Errorf("function only accepts structs; got %s", kind)
	}

	var errs Errors
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		fieldType := val.Type().Field(i)

		tag := fieldType.Tag.Get(tagName)

		if ((fieldVal.Kind() == reflect.Struct) || (fieldVal.Kind() == reflect.Ptr && fieldVal.Elem().Kind() == reflect.Struct)) && tag != "-" {
			// Recurse over nested structs
			err := minimizer(l, fieldVal.Interface(), &fieldVal)
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			if tag != "" {
				log.Debug("Applying '%s' minimizer for field '%s' using level %d", tag, fieldType.Name, l)

				if minFunc, ok := TagMap[tag]; ok {
					result := minFunc(l, fieldVal.Interface())
					rv := reflect.ValueOf(result)

					if fieldVal.CanSet() == true {
						fieldVal.Set(rv)
					} else {
						if parent != nil && parent.CanSet() {
							parent.FieldByName(fieldType.Name).Set(rv)
						} else {
							errStr := fmt.Errorf("unable to set field value for field '%s'", fieldType.Name)
							errs = append(errs, errStr)
						}
					}
				} else {
					errStr := fmt.Errorf("missing minimizer for type '%s'", tag)
					errs = append(errs, errStr)
				}
			}
		}
	}

	if len(errs) > 0 {
		return errs
	}

	return nil
}

func MinimizeStruct(l MinimizationLevel, s interface{}) error {
	return minimizer(l, s, nil)
}
