package controllers

import (
	"encoding/json"
	"github.com/joshua-ather/sv_be/app/models"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func GetPosts(c echo.Context) error {

	post := models.Post{}

	posts, err := post.FindPostWLimitOffset(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, posts)
}

func GetPostById(c echo.Context) error {

	post := models.Post{}

	posts, err := post.FindPostByID(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, posts)
}

func CreatePost(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	post := models.Post{}
	err = json.Unmarshal(body, &post)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	post.Prepare()
	err = post.Validate()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	_, err = post.SavePost()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, []string{})
}

func UpdatePost(c echo.Context) error {

	body, err := ioutil.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	// Start processing the request data
	postUpdate := models.Post{}
	err = json.Unmarshal(body, &postUpdate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	postUpdate.Prepare()
	err = postUpdate.Validate()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	_, err = postUpdate.UpdatePost(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, []string{})
}

func DeletePost(c echo.Context) error {

	// Check if the post exist
	post := models.Post{}

	_, err := post.DeletePost(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	return c.JSON(http.StatusOK, []string{})
}
