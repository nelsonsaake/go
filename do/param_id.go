package do

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParamID: extracts :id value in the request
func ParamID(c *gin.Context) (id int, err error) {

	idStr := c.Param("id")
	if len(idStr) == 0 {
		err = fmt.Errorf("error, no id provided")
		return
	}

	id, err = strconv.Atoi(idStr)
	if err != nil {
		err = fmt.Errorf("error parsing id to int: %v", err)
		return
	}

	return
}
