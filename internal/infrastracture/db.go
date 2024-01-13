package infrastracture

import "errors"

var ErrEntityNotFound = errors.New("entity not found")

type KeyValueStore struct {
	data map[string]interface{}
}

var constUser = map[string]interface{}{
	"user1@example.com": "11111",
	"user2@example.com": "22222",
}

func NewKeyValueStore() *KeyValueStore {
	return &KeyValueStore{
		data: constUser,
	}
}

func (kvs *KeyValueStore) Get(key string) (interface{}, error) {
	v, ok := kvs.data[key]
	if !ok {
		return nil, ErrEntityNotFound
	}
	return v, nil
}
