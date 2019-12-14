package minimizers

import (
	"fmt"
	"github.com/adaptant-labs/go-minimizer/datagen"
	"github.com/brianvoe/gofakeit"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"reflect"
	"strings"
	"time"
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
		var restriction MinimizationLevel

		level := l
		fieldVal := val.Field(i)
		fieldType := val.Type().Field(i)

		tagValue := fieldType.Tag.Get(tagName)
		tagSpecs := strings.Split(tagValue, ",")

		// The default case: `minimizer:"tag"`
		tag := tagSpecs[0]
		if len(tagSpecs) > 2 {
			err := fmt.Errorf("invalid tag specification for field '%s', at most 1 restriction can be specified", fieldType.Name)
			errs = append(errs, err)
			break
		} else if len(tagSpecs) == 2 {
			// Restriction specified: `minimizer:"-,mask"`
			restrictTo := tagSpecs[1]
			restriction = LevelFromString(restrictTo)
		}

		if ((fieldVal.Kind() == reflect.Struct) || (fieldVal.Kind() == reflect.Ptr && fieldVal.Elem().Kind() == reflect.Struct)) && tag != "-" {
			// Recurse over nested structs
			err := minimizer(l, fieldVal.Interface(), &fieldVal)
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			if tag != "" {
				var rv reflect.Value

				if restriction != MinimizationNone {
					log.Debugf("Restricting level to %s based on tag restriction\n", restriction.String())
					level = restriction
				}

				log.Debugf("Applying '%s' minimizer for field '%s' using level '%s'\n",
					tag, fieldType.Name, level)

				if level == MinimizationTokenize {
					token, err := datagen.GenerateRandomToken()
					if err != nil {
						errs = append(errs, err)
					} else {
						rv = reflect.ValueOf(token)
					}
				} else if level == MinimizationMask {
					result, err := Mask(fieldVal.String())
					if err != nil {
						errs = append(errs, err)
					} else {
						rv = reflect.ValueOf(result)
					}
				} else if minFunc, ok := TagMap[tag]; ok {
					result := minFunc(level, fieldVal.Interface())
					rv = reflect.ValueOf(result)
				} else {
					errStr := fmt.Errorf("missing minimizer for type '%s'", tag)
					errs = append(errs, errStr)
				}

				if rv.IsValid() {
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

func EnableDebugLogging() {
	log.SetLevel(log.DebugLevel)
}

func init() {
	rand.Seed(time.Now().UnixNano())
	gofakeit.Seed(time.Now().UnixNano())
}
