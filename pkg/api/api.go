package api

// The Api needs a metrics endpoint that can return all defined metrics within a namespace
// defined what are namespaces in the context of the API

// by definition this is incorrect.
type namespace struct {
	metric metrics
}

// temporary struct definition for metrics
type metrics struct {
	// gauges are numbers that are expected to fluctuate over time
	// a gauge is a snapshot in time
	gauge uint
	// a counter is a number that increases over time but, unlike gauges, doesn't decrease unless you reset it to zero.
	counter int
	// timers are ment to track how much something took
	timers int
	// deltas are the difference between two metric values in time.
	delta delta
}

type delta struct {
	initialValue int
	finalValue   int
}
