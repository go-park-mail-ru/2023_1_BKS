package repository

import (
	"context"
	"fmt"

	"github.com/gomodule/redigo/redis"
	"github.com/google/uuid"
)

type CartRedisRepository struct {
	cart redis.Conn
}

func (r *CartRedisRepository) Add(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error {
	result, err := r.cart.Do("SET", userId, postId)
	if err != nil {
		return err
	}
	if result != "OK" {
		return fmt.Errorf("result not OK")
	}
	return nil
}

func (r *CartRedisRepository) Remove(ctx context.Context, userId uuid.UUID, postId uuid.UUID) error {
	result, err := r.cart.Do("SET", "Favorite Movie", "Repo Man")
	if err != nil {
		return err
	}
	if result != "OK" {
		return fmt.Errorf("result not OK")
	}
	return nil
}

func (r *CartRedisRepository) Get(ctx context.Context, userId uuid.UUID) ([]string, error) {
	values, err := redis.Strings(r.cart.Do("Smembers", userId.String()))
	if err != nil {
		return []string{}, err
	}

	return values, nil
}
