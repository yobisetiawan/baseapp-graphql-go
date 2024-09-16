package utils

import (
	"baseapp/app/database"
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func BaseCrudSetParams(c echo.Context, db *gorm.DB, searchFields []string, defSortOrder string) (int, int, string, *gorm.DB) {
	// Pagination parameters
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil || limit < 1 {
		limit = 0
	}

	// Sorting parameter
	sort := c.QueryParam("sort")
	orderBy := c.QueryParam("order_by")
	sortOrder := ""
	if sort != "" && orderBy != "" {
		sortOrder = orderBy + " " + sort
	} else {
		if defSortOrder != "" {
			sortOrder = defSortOrder
		} else {
			sortOrder = "created_at desc"
		}
	}

	// Search parameter
	search := c.QueryParam("search")

	if search != "" {
		searchQuery := ""
		searchArgs := make([]interface{}, 0)
		for _, field := range searchFields {
			if searchQuery != "" {
				searchQuery += " OR "
			}
			searchQuery += field + " ILIKE ?"
			searchArgs = append(searchArgs, "%"+search+"%")
		}

		db = db.Where(searchQuery, searchArgs...)
	}

	return page, limit, sortOrder, db
}

func BaseCrudPagination(c echo.Context, db *gorm.DB, model interface{}, page int, limit int, sortOrder string) error {
	modelType := reflect.TypeOf(model)
	sliceType := reflect.SliceOf(modelType)
	listPtr := reflect.New(sliceType).Interface()

	//modelName := modelType.Name()

	var total int64
	db.Model(model).Count(&total)

	// Apply pagination and sorting
	offset := (page - 1) * limit

	db = db.Order(sortOrder)
	if limit > 0 {
		db = db.Limit(limit).Offset(offset)
	}
	db.Find(listPtr)

	// Calculate total pages
	totalPages := int64(1)
	if limit > 0 {
		totalPages = (total + int64(limit) - 1) / int64(limit)
	}

	listData := reflect.ValueOf(listPtr).Elem().Interface()

	if c.QueryParam("type") == "collection" {
		return c.JSON(http.StatusOK, echo.Map{"data": listData})
	} else {
		return c.JSON(http.StatusOK, echo.Map{
			"data": listData,
			"meta": echo.Map{
				"total":        total,
				"current_page": page,
				"total_pages":  totalPages,
				"per_page":     limit,
			},
		})
	}
}

func BaseCrudBindAndValidate(c echo.Context, req, dst interface{}) error {
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, NewValidationErrorResponse(err))
	}

	if err := copier.Copy(dst, req); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func BaseCrudFindDataByField(c echo.Context, model interface{}, field string, value interface{}) (interface{}, error) {
	if result := database.DB.First(model, fmt.Sprintf("%s = ?", field), value); result.Error != nil {
		return nil, c.JSON(http.StatusNotFound, echo.Map{"error": "Data is not found"})
	}
	return model, nil
}
