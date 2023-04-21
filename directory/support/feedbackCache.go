package support

type FedBackCache struct {
}

func (r *FedBackCache) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (r *FedBackCache) Write(p []byte) (n int, err error) {
	return 0, nil
}
