package generator

type Generator interface {
	// GenerateText generate a text based on a strategy
	GenerateBaconText() string
}