package handler

import (
	"net/http"
	"strconv"

	"github.com/AsaHero/todolist/db/models"
	"github.com/gin-gonic/gin"
)

func (h* Handler) getAllLists(c *gin.Context) {
	userId, err := getUserID(c)

	if err != nil {
		return 
	}
	lists, err := h.service.ToDoList.GetAll(userId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)

}

func (h* Handler) getListById(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id privided")
		return
	}

	list, err := h.service.ToDoList.GetById(userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return 
	}

	c.JSON(http.StatusOK, list)
}

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	var input models.ToDoList
	if err = c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.ToDoList.Create(userId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
	
}
	

func (h *Handler) upgradeList(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var input models.UpdateToDoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
	} 
	err = h.service.ToDoList.Update(userId, listId, input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserID(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "id value is invalid")
		return
	}
	err = h.service.ToDoList.DeleteById(userId, listId)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
	
}