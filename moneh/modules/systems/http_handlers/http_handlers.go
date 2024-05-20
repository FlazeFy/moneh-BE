package httphandlers

import (
	"moneh/modules/systems/models"
	"moneh/modules/systems/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func GetDictionaryByType(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	dctType := c.Param("type")
	result, err := repositories.GetDictionaryByType(page, 10, "api/v1/dct/"+dctType, dctType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelDictionaryById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelDictionaryById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostDictionary(c echo.Context) error {
	var obj models.PostDictionaryByType

	// Data
	obj.DctName = c.FormValue("dictionaries_name")
	obj.DctType = c.FormValue("dictionaries_type")

	result, err := repositories.PostDictionary(obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostFeedback(c echo.Context) error {
	var obj models.PostFeedback
	fdbRateInt, _ := strconv.Atoi(c.FormValue("feedbacks_rate"))

	// Data
	obj.FdbRate = fdbRateInt
	obj.FdbDesc = c.FormValue("feedbacks_desc")

	result, err := repositories.PostFeedback(obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllFeedback(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	ord := c.Param("ord")
	ord_obj := c.Param("ord_obj")
	result, err := repositories.GetAllFeedback(page, 10, "api/v1/feedback/"+ord_obj+"/"+ord, ord_obj, ord)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTags(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	result, err := repositories.GetAllTags(page, 10, "api/v1/tag")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetAllTagsFirebase(c echo.Context) error {
	result, err := repositories.GetAllTagsFirebase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func GetDictionaryByTypeFirebase(c echo.Context) error {
	result, err := repositories.GetDictionaryByTypeFirebase()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func HardDelTagById(c echo.Context) error {
	id := c.Param("id")
	result, err := repositories.HardDelTagById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}

func PostTag(c echo.Context) error {
	var obj models.GetTags

	// Data
	obj.TagSlug = c.FormValue("tags_slug")
	obj.TagName = c.FormValue("tags_name")

	result, err := repositories.PostTag(obj)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, result)
}
