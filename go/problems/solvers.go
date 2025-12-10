package problems

// Solvers is a map of problem numbers to their solver functions.
// It's populated automatically via init() in each problem file.
var Solvers = map[int]func() int64{}
