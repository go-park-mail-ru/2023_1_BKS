package repository

import (
	"context"
	"post/domain"

	"github.com/go-redis/redis"
)

type CartRedisRepository struct {
	posts *redis.Client
}

func (c *CartRedisRepository) Add(ctx context.Context, cart domain.Cart) error {
	err := c.posts.Set(cart.Title.String(), cart.IdPost, 0).Err()
	return err
}

func (c *CartRedisRepository) Remove(ctx context.Context, cart domain.Cart) error {
	err := c.posts.Del(cart.Title.String()).Err()
	return err
}

func (c CartRedisRepository) Get(ctx context.Context) ([]domain.Cart, error) {
	err := c.posts.Sort(rg, redis.)
	return err
}
