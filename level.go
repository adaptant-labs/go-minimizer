package minimizers

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
