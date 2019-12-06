package minimizers

import (
	"bytes"
	"encoding/json"
	"log"
)

type MinimizationLevel int

const (
	// No data minimization
	MinimizationNone MinimizationLevel = iota

	// Coarse-grained data minimization
	MinimizationCoarse

	// Fine-grained data minimization
	MinimizationFine

	// Data anonymization
	MinimizationAnonymize
)

var levelStringMap = map[MinimizationLevel]string {
	MinimizationNone:	"none",
	MinimizationCoarse:	"coarse",
	MinimizationFine:	"fine",
	MinimizationAnonymize:	"anonymize",
}

var stringLevelMap = map[string]MinimizationLevel {
	"none":		MinimizationNone,
	"coarse":	MinimizationCoarse,
	"fine":		MinimizationFine,
	"anonymize":	MinimizationAnonymize,
}

func levelFromString(levelString string) MinimizationLevel {
	if level, ok := stringLevelMap[levelString]; ok {
		return level
	}

	return MinimizationNone
}

func (l MinimizationLevel) MarshalJSON() ([]byte, error) {
	buf := bytes.NewBufferString(`"`)
	buf.WriteString(l.String())
	buf.WriteString(`"`)
	return buf.Bytes(), nil
}

func (l *MinimizationLevel) UnmarshalJSON(b []byte) error {
	var levelString string

	err := json.Unmarshal(b, &levelString)
	if err != nil {
		return err
	}

	*l = levelFromString(levelString)
	return nil
}

func (l MinimizationLevel) String() string {
	if name, ok := levelStringMap[l]; ok {
		return name
	}

	log.Printf("Unknown minimization level specified (%d)\n", l)
	return ""
}
