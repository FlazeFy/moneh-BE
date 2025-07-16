package stats

import (
	"encoding/json"
	"moneh/models"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// Stats Interface
type StatsService interface {
	GetMostUsedContext(tableName, targetCol string, userId uuid.UUID) ([]models.StatsContextTotal, error)
	GetMonthlyFlow(year int, userId uuid.UUID) ([]models.StatsFlowMonthly, error)
}

// Stats Struct
type statsService struct {
	statsRepo   StatsRepository
	redisClient *redis.Client
	statsCache  StatsCache
}

// Stats Constructor
func NewStatsService(statsRepo StatsRepository, redisClient *redis.Client, statsCache StatsCache) StatsService {
	return &statsService{
		statsRepo:   statsRepo,
		redisClient: redisClient,
		statsCache:  statsCache,
	}
}

func (s *statsService) GetMostUsedContext(tableName, targetCol string, userId uuid.UUID) ([]models.StatsContextTotal, error) {
	// Cache : Get Key
	cacheKey := s.statsCache.StatsKeyMostUsedContext("flows", targetCol, userId)
	// Cache : Temp Stats
	stats, err := s.statsCache.GetStatsMostUsedContext(s.redisClient, cacheKey)
	if err == nil {
		return stats, nil
	}

	// Repo : Find Most Used Context
	stats, err = s.statsRepo.FindMostUsedContext(tableName, targetCol, userId)
	if err != nil {
		return nil, err
	}

	// Cache : Store Redis
	jsonData, _ := json.Marshal(stats)
	s.statsCache.SetStatsCache(s.redisClient, cacheKey, jsonData)

	return stats, nil
}

func (s *statsService) GetMonthlyFlow(year int, userId uuid.UUID) ([]models.StatsFlowMonthly, error) {
	// Cache : Get Key
	cacheKey := s.statsCache.StatsKeyMonthlyFlow(year, userId)
	// Cache : Temp Stats
	stats, err := s.statsCache.GetStatsMonthlyFlow(s.redisClient, cacheKey)
	if err == nil {
		return stats, nil
	}

	// Repo : Find Most Used Context
	stats, err = s.statsRepo.FindMonthlyFlow(year, userId)
	if err != nil {
		return nil, err
	}

	// Cache : Store Redis
	jsonData, _ := json.Marshal(stats)
	s.statsCache.SetStatsCache(s.redisClient, cacheKey, jsonData)

	return stats, nil
}
