package task

// TaskStatus represents the status of a task.
type TaskStatus string

const (
	StatusTODO       TaskStatus = "TODO"
	StatusInProgress TaskStatus = "IN_PROGRESS"
	StatusDone       TaskStatus = "DONE"
	StatusCancelled  TaskStatus = "CANCELLED"
)

// IsValid checks whether the task status is a known value.
func (s TaskStatus) IsValid() bool {
	switch s {
	case StatusTODO, StatusInProgress, StatusDone, StatusCancelled:
		return true
	default:
		return false
	}
}

// validTransitions defines the allowed status transitions.
var validTransitions = map[TaskStatus][]TaskStatus{
	StatusTODO:       {StatusInProgress, StatusCancelled},
	StatusInProgress: {StatusDone, StatusCancelled},
	StatusDone:       {},
	StatusCancelled:  {},
}

// IsValidTransition checks if a transition from one status to another is allowed.
func IsValidTransition(from, to TaskStatus) bool {
	if from == to {
		return true
	}
	allowed, ok := validTransitions[from]
	if !ok {
		return false
	}
	for _, s := range allowed {
		if s == to {
			return true
		}
	}
	return false
}

// TaskPriority represents the priority level of a task.
type TaskPriority string

const (
	PriorityLow    TaskPriority = "LOW"
	PriorityMedium TaskPriority = "MEDIUM"
	PriorityHigh   TaskPriority = "HIGH"
	PriorityUrgent TaskPriority = "URGENT"
)

// IsValid checks whether the task priority is a known value.
func (p TaskPriority) IsValid() bool {
	switch p {
	case PriorityLow, PriorityMedium, PriorityHigh, PriorityUrgent:
		return true
	default:
		return false
	}
}
