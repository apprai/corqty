package worker

// TerraformWorker provide terraform job as worker
type TerraformWorker struct {
}

// GetName return name of worker
func (worker TerraformWorker) GetName() string {
	return "terraform"
}
