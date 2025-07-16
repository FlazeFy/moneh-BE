package stats

import (
	"context"
	"encoding/json"
	"fmt"
	"moneh/config"
	"moneh/models"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// Stats Interface
type StatsCache interface {
	// General
	SetStatsCache(redis *redis.Client, key string, data []byte) error

	// Context
	StatsKeyMostUsedContext(table, field string, userID uuid.UUID) string
	RefreshStatsMostUsedContext(redis *redis.Client, table string, fields []string, userID uuid.UUID)
	GetStatsMostUsedContext(redis *redis.Client, key string) ([]models.StatsContextTotal, error)
	StatsKeyMonthlyFlow(year int, userID uuid.UUID) string
	RefreshStatsMonthlyFlow(redis *redis.Client, year int, userID uuid.UUID)
	GetStatsMonthlyFlow(redis *redis.Client, key string) ([]models.StatsFlowMonthly, error)
}

// Stats Struct
type statsCache struct {
	redisClient *redis.Client
}

// Stats Constructor
func NewStatsCache(redisClient *redis.Client) StatsCache {
	return &statsCache{
		redisClient: redisClient,
	}
}

// General Stats
func (ch *statsCache) SetStatsCache(redis *redis.Client, key string, data []byte) error {
	return redis.Set(context.Background(), key, data, config.RedisTime).Err()
}

// For API GET : Most Used Context
func (ch *statsCache) StatsKeyMostUsedContext(table, field string, userID uuid.UUID) string {
	return fmt.Sprintf("stats:%s:%s:%s", table, field, userID.String())
}

func (ch *statsCache) RefreshStatsMostUsedContext(redis *redis.Client, table string, fields []string, userID uuid.UUID) {
	for _, field := range fields {
		key := ch.StatsKeyMostUsedContext(table, field, userID)
		redis.Del(context.Background(), key)
	}
}

func (ch *statsCache) GetStatsMostUsedContext(redis *redis.Client, key string) ([]models.StatsContextTotal, error) {
	var stats []models.StatsContextTotal

	cached, err := redis.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(cached), &stats); err != nil {
		return nil, err
	}

	return stats, nil
}

// For API GET : Monthly Flow
func (ch *statsCache) StatsKeyMonthlyFlow(year int, userID uuid.UUID) string {
	return fmt.Sprintf("stats:flow_monthly:%s:%s", year, userID.String())
}

func (ch *statsCache) RefreshStatsMonthlyFlow(redis *redis.Client, year int, userID uuid.UUID) {
	key := ch.StatsKeyMonthlyFlow(year, userID)
	redis.Del(context.Background(), key)
}

func (ch *statsCache) GetStatsMonthlyFlow(redis *redis.Client, key string) ([]models.StatsFlowMonthly, error) {
	var stats []models.StatsFlowMonthly

	cached, err := redis.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(cached), &stats); err != nil {
		return nil, err
	}

	return stats, nil
}
