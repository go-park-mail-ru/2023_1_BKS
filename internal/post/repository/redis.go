package repository

import (
	"context"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type CartRedisRepository struct {
	cart redis.Conn
}

func (r *CartRedisRepository) Add(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error {
	_, err := r.cart.Do("sadd", userId.String(), postId.String())
	if err != nil {
		return err
	}
	return nil
}

func (r *CartRedisRepository) Remove(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error {
	_, err := r.cart.Do("srem", userId.String(), postId.String())
	if err != nil {
		return err
	}
	return nil
}

func (r *CartRedisRepository) Get(ctx context.Context, userId uuid.UUID) ([]string, error) {
	values, err := redis.Strings(r.cart.Do("smembers", userId.String()))
	if err != nil {
		return []string{}, err
	}

	return values, nil
}
