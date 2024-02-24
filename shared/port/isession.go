package port

import "time"

type ISession interface {
	Start(id string, data map[string]interface{}, expiry ...time.Duration) error
	GetData(id string) (map[string]interface{}, error)
	Destroy(id string) error
	Set(id, key string, value interface{}, expiry ...time.Duration) error
	Unset(id, key string, expiry ...time.Duration) error
	RemoveKey(id, key string, expiry ...time.Duration) error
}
