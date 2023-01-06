package storage

import (
	"crypto/md5"
	"encoding/hex"
	"sync"

	"github.com/paveltyukin/practicum-go-service-devops/internal"
)

type Storage struct {
	mx     sync.Mutex
	values map[string]internal.UpdateParams
}

func hashKey(mType, mName string) string {
	hash := md5.Sum([]byte(mType + "_" + mName))
	return hex.EncodeToString(hash[:])
}

func (s *Storage) Set(mType, mName, mValue string) {
	s.mx.Lock()
	h := hashKey(mType, mName)
	s.values[h] = internal.UpdateParams{
		MType:  mType,
		MName:  mName,
		MValue: mValue,
	}

	s.mx.Unlock()
}

func New() *Storage {
	return &Storage{
		values: make(map[string]internal.UpdateParams),
	}
}
