package entropy

type entropy struct {
	frequencies [256]int
	total       int
}

// NewEntropyCalculator allocates new entropy calculator and returns a pointer to
// created object.
func NewEntropyCalculator() *entropy { return new(entropy) }

func (e *entropy) Write(data []byte) (n int, err error) {
	for _, d := range data {
		e.frequencies[int(d)] += 1
		e.total += 1
	}

	return e.total, nil
}

func (e *entropy) Calculate() float64 {

	return 0.0
}
