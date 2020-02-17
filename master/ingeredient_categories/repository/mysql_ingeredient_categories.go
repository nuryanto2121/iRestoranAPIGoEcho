package repoingeredientcategories

import (
	"context"
	"encoding/base64"
	"fmt"
	"time"

	// "database/sql"

	ingeredientcategories "yanto/irestora/master/ingeredient_categories/domain"
	"yanto/irestora/models"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

const (
	timeFormat = "2006-01-02T15:04:05.999Z07:00" // reduce precision from RFC3339Nano as date format
)

type mysqlIngeredientCategoriesRepository struct {
	Conn *sqlx.DB
	// Conn *sql.DB
}

// NewMysqlIngeredientCategoriesRepository :
// func NewMysqlIngeredientCategoriesRepository(Conn *sql.DB) ingeredientcategories.Repository {
// 	// func NewMysqlIngeredientCategoriesRepository(Conn *sqlx.DB) ingeredientcategories.Repository {
// 	return &mysqlIngeredientCategoriesRepository{Conn}
// }

// NewMysqlIngeredientCategoriesRepository :
func NewMysqlIngeredientCategoriesRepository(Conn *sqlx.DB) ingeredientcategories.Repository {
	return &mysqlIngeredientCategoriesRepository{Conn}
}

func (m *mysqlIngeredientCategoriesRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.IngredientCategories, error) {
	fmt.Printf("%s\n", query)
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
		// return []models.IngredientCategories{}, err
	}
	defer func() {
		err := rows.Close()
		if err != nil {
			logrus.Error(err)
		}
	}()

	result := make([]*models.IngredientCategories, 0)
	for rows.Next() {
		t := new(models.IngredientCategories)
		err = rows.Scan(
			&t.ID,
			&t.CategoryName,
			&t.Description,
			&t.UserID,
			&t.CompanyID,
			&t.DelStatus,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}
	return result, nil
}

func (m *mysqlIngeredientCategoriesRepository) GetIngeredientCategories(ctx context.Context) (result []models.IngredientCategories, err error) {
	Squery := `SELECT id,category_name,description,user_id,company_id,del_status FROM tbl_ingredient_categories WHERE del_status ='LIVE'`
	// 	var cursor = ""
	// 	decodedCursor, err := DecodeCursor(cursor)
	// 	if err != nil && cursor != "" {
	// 		return nil, models.ErrBadParamInput
	// 	}

	// 	res, err := m.fetch(ctx, Squery, decodedCursor, 10)
	err = m.Conn.SelectContext(ctx, &result, Squery)
	if err != nil {
		return nil, err
	}
	return result, nil
	// 	// nextCursor :=""

	// 	return res, nil
	// }

	// // DecodeCursor will decode cursor from user for mysql
	// func DecodeCursor(encodedTime string) (time.Time, error) {
	// 	byt, err := base64.StdEncoding.DecodeString(encodedTime)
	// 	if err != nil {
	// 		return time.Time{}, err
	// 	}

	// 	timeString := string(byt)
	// 	t, err := time.Parse(timeFormat, timeString)

	// 	return t, err

}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeFormat)

	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
