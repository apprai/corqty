// Package worker provide interface and common function for worker
package worker

// IWorker is base of worker interface
type IWorker interface {
	GetName() string
}

// GetWorker return worker instance that match with name of worker
func GetWorker(workerName string) IWorker {
	switch workerName {
	case "terraform":
		return TerraformWorker{}
	case "something":
		return TerraformWorker{}
	}
	return nil
}
