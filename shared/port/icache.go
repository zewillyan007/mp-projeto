package port

type ICache interface {
	Delete(key string) error
	Get(key string) ([]byte, error)
	Set(key string, entry []byte) error
}
