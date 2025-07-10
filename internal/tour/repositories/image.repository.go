package repositories

import (
	"errors"
	"fmt"

	"github.com/go-jet/jet/v2/qrm"
	. "github.com/go-jet/jet/v2/sqlite"
	. "github.com/oliwallywonka/alpaca_backend/db/table"

	"github.com/oliwallywonka/alpaca_backend/internal/shared"
	commonerrors "github.com/oliwallywonka/alpaca_backend/internal/shared/errors"
)

func (r *tourRepository) GetTourImages(id string) ([]string, error) {
	var tourStruct struct {
		Images shared.ImageField
	}
	stm := SELECT(Tour.Images).
		FROM(Tour).
		WHERE(Tour.ID.EQ(String(id))).
		LIMIT(1)
	err := stm.Query(r.db, &tourStruct)
	if errors.Is(err, qrm.ErrNoRows) {
		return nil, commonerrors.NotFoundError
	}
	if err != nil {
		return nil, fmt.Errorf("unmarshal: %w", err)
	}
	return tourStruct.Images, nil
}

func (r *tourRepository) UpdateImageTour(tourID string, images shared.ImageField) error {
	err := r.gorm.Table("tour").
		Where("id = ?", tourID).
		Updates(map[string]interface{}{
			"images": images,
		}).
		Error
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}
	return nil
}
