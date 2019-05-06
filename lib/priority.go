package task

// Priority struct for diffrent priorities available
type Priority int

// consts of all Priorities
const (
	HIGH Priority = iota
	MEDIUM
	LOW
)

var priorities = [...]string{"HIGH", "MEDIUM", "LOW"}

func (p Priority) String() string {
	return priorities[p]
}
