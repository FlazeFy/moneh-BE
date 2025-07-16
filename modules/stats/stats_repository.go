package stats

import (
	"fmt"
	"moneh/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Stats Interface
type StatsRepository interface {
	// Flow & Pocket
	FindMostUsedContext(tableName, targetCol string, userId uuid.UUID) ([]models.StatsContextTotal, error)
	// Flow
	FindMonthlyFlow(year int, userId uuid.UUID) ([]models.StatsFlowMonthly, error)
}

// Stats Struct
type statsRepository struct {
	db *gorm.DB
}

// Stats Constructor
func NewStatsRepository(db *gorm.DB) StatsRepository {
	return &statsRepository{db: db}
}

func (r *statsRepository) FindMostUsedContext(tableName, targetCol string, userId uuid.UUID) ([]models.StatsContextTotal, error) {
	// Models
	var stats []models.StatsContextTotal

	// Query
	result := r.db.Table(tableName).
		Select(fmt.Sprintf("COUNT(%s) as total, %s as context", targetCol, targetCol)).
		Where("created_by", userId).
		Group(targetCol).
		Order("total DESC").
		Limit(7).
		Find(&stats)

	if result.Error != nil {
		return nil, result.Error
	}
	if len(stats) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return stats, nil
}

func (r *statsRepository) FindMonthlyFlow(year int, userId uuid.UUID) ([]models.StatsFlowMonthly, error) {
	// Models
	var stats []models.StatsFlowMonthly

	// Query
	result := r.db.Table("flows f").
		Select(`
			SUM(CASE WHEN f.flow_type = 'Income' THEN fr.ammount ELSE 0 END) AS total_income,
			SUM(CASE WHEN f.flow_type = 'Spending' THEN fr.ammount ELSE 0 END) AS total_spending,
			TRIM(TO_CHAR(fr.created_at, 'Month')) AS context,
			TO_CHAR(fr.created_at, 'MM') AS month_num
		`).
		Joins("JOIN flow_relations fr ON fr.flow_id = f.id").
		Where("f.created_by", userId).
		Where("TO_CHAR(fr.created_at, 'YYYY') = ?", fmt.Sprintf("%04d", year)).
		Group("context, month_num").
		Order("month_num ASC").
		Find(&stats)

	if result.Error != nil {
		return nil, result.Error
	}
	if len(stats) == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return stats, nil
}
