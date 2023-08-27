package user

import (
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h UserHandler) Delete(c *gin.Context) {
	var u UserModel
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
	var u UserModel
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
	var u UserModel
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

	resp, err := h.userService.List(c, paging.Page, paging.Limit)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}

	c.JSON(200, resp)
}
func (h UserHandler) Detail(c *gin.Context) {
	var u UserModel
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
