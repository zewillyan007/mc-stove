package port

import "time"

type IDocument interface {
	Create(id, bucket string, data map[string]interface{}, expiry ...time.Duration) error
	GetData(id, bucket string) (map[string]interface{}, error)
	Destroy(id, bucket string) error
	Set(id, bucket string, data map[string]interface{}, expiry ...time.Duration) error
	RemoveKey(id, bucket string) error
}
