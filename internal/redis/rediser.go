package redis

import (
	"context"
	"time"
)

type Redisr interface {
	// Interface for setting single key value in redis
	HSet(ctx context.Context, hkey string, keyValueMap map[string]interface{}, expiration time.Duration) error
	// Interface for setting multiple keys values in redis
	HMSet(ctx context.Context, hkey string, keyValueMap []map[string]interface{}, expiration time.Duration) error
}
