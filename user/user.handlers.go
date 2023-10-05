package user

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
}

func (h *UserHandler) Index(c *gin.Context) {
	var provinces []Provinces

	provinces, err := h.userService.List(c)
	if err != nil {

		temp, err := template.ParseFiles("views/index.html")
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)

			return

		}
		fmt.Println(`Hasil`)
		if err := temp.Execute(c.Writer, provinces); err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
		}
	}
}
func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h UserHandler) Delete(c *gin.Context) {
	var u Provinces
	if err := c.ShouldBindUri(&u); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	err := h.userService.Delete(c, u.Id)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(200, gin.H{"msg": "DELETED"})
}

func (h UserHandler) Create(c *gin.Context) {
	var u Provinces
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	resp, err := h.userService.Create(c, u)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(200, resp)
}

func (h UserHandler) Update(c *gin.Context) {
	var u Provinces
	if err := c.ShouldBind(&u); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	resp, err := h.userService.Update(c, u)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(200, resp)
}

func (h UserHandler) List(c *gin.Context) {
	var paging Paging
	if err := c.ShouldBindUri(&paging); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	resp, err := h.userService.List(c)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(200, resp)
}
func (h UserHandler) Detail(c *gin.Context) {
	var u Provinces
	if err := c.ShouldBindUri(&u); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	resp, err := h.userService.Detail(c, u.Id)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(200, resp)
}
