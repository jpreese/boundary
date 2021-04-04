package job

type Status string

const (
	// Running represents that the job run is actively running on a server
	Running Status = "running"

	// Completed represents that the job run has successfully finished
	Completed Status = "completed"

	// Failed represent that the job run had an error during execution
	Failed Status = "failed"

	// Interrupted represents that the job run was interrupted by a server
	// other than the server running the job.
	Interrupted Status = "interrupted"
)

func (s Status) isValidRunStatus() bool {
	switch s {
	case Running, Completed, Failed, Interrupted:
		return true
	}
	return false
}

func (s Status) isFinalRunStatus() bool {
	switch s {
	case Completed, Failed, Interrupted:
		return true
	}
	return false
}

func (s Status) String() string {
	return string(s)
}
