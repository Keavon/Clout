package api

import "github.com/garyburd/redigo/redis"

// API is the parent type for all methods and stores the redis client
type API struct {
	Pool *redis.Pool
}

// New creates a new API object
func New(rp *redis.Pool) *API {
	return &API{Pool: rp}
}
