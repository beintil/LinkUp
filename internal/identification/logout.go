package identification

import (
	"LinkUp_Update/internal/identification/services"
	"github.com/gin-gonic/gin"
)

func Logout(c *gin.Context) {
	services.Get(c, nil, nil).Logout()
}
