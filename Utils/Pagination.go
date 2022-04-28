package utils

import (
	models "first/Model"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GeneratePaginationFromRequest(c *gin.Context) models.Pagination {

	limit := 2
	page := 1
	sort := "created_at asc"
	query := c.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = queryValue
			break
		}
	}
	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}

}
