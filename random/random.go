package random

import (
	"github.com/google/uuid"
	"github.com/syzhang42/go-fire/random/snow"
)

func GetInt64NextID() int64 {
	if snow.DefaultSf == nil {
		snow.NewSnowflake()
	}
	return snow.DefaultSf.NextID()
}

func GetUUid() string {
	return uuid.NewString()
}
