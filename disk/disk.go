package disk

// Memory holds memory metadata about the host
type Disk struct {
}

const name = "disk"

func (d *Disk) Name() string {
	return name
}

func (d *Disk) Collect() (result interface{}, err error) {
	result, err = getDiskInfo()
	return
}
