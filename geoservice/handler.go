package geoservice

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary      Get nearby locations
// @Description  Returns all locations within the given radius (in kilometers) of the provided latitude/longitude.
// @Tags         Location
// @Accept       json
// @Produce      json
// @Param        lat     query   number  true  "Latitude, in decimal degrees"   example(28.7041)
// @Param        lon     query   number  true  "Longitude, in decimal degrees"  example(77.1025)
// @Param        radius  query   number  true  "Search radius in kilometers"    example(5)
// @Success      200     {array} Location  "List of nearby locations"
// @Failure      400     {object} ErrorResponse    "Missing or invalid parameters"
// @Router       /nearby [get]
func NearbyHandler(c *gin.Context, root *KDNode) {
	latStr := c.Query("lat")
	lonStr := c.Query("lon")
	radiusStr := c.Query("radius")

	if latStr == "" || lonStr == "" || radiusStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing lat/lon/radius"})
		return
	}

	lat, _ := strconv.ParseFloat(latStr, 64)
	lon, _ := strconv.ParseFloat(lonStr, 64)
	radius, _ := strconv.ParseFloat(radiusStr, 64)

	var result []Location
	RadiusSearch(root, lat, lon, radius, 0, &result)
	c.JSON(http.StatusOK, result)

}

// for swag doc
// ErrorResponse represents the error message returned by the API
type ErrorResponse struct {
	Error string `json:"error" example:"Missing lat/lon/radius"`
}
