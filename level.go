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

var LevelStringMap = map[MinimizationLevel]string {
	MinimizationNone:	"none",
	MinimizationCoarse:	"coarse",
	MinimizationFine:	"fine",
	MinimizationAnonymize:	"anonymize",
}

var StringLevelMap = map[string]MinimizationLevel {
	"none":		MinimizationNone,
	"coarse":	MinimizationCoarse,
	"fine":		MinimizationFine,
	"anonymize":	MinimizationAnonymize,
}

func LevelFromString(levelString string) MinimizationLevel {
	if level, ok := StringLevelMap[levelString]; ok {
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

	*l = LevelFromString(levelString)
	return nil
}

func (l MinimizationLevel) String() string {
	if name, ok := LevelStringMap[l]; ok {
		return name
	}

	log.Printf("Unknown minimization level specified (%d)\n", l)
	return ""
}
