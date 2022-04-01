package rest

import (
	"github.com/RogerDurdn/users/domain"
	"github.com/RogerDurdn/users/model"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type Rest struct {
	service domain.Service
	address string
}

func NewRest(service domain.Service, address string) *Rest {
	return &Rest{service: service, address: ":" + address}
}

func (rs *Rest) Start() {
	r := gin.Default()
	rg := r.Group("/api")
	rg.GET("/user/:id", rs.getUserById)
	rg.GET("/user", rs.getUserByName)
	rg.POST("/user", rs.createUser)
	rg.PUT("/user", rs.updateUser)
	rg.DELETE("/user/:id", rs.deleteUserById)
	rg.POST("/user/auth", rs.authUser)
	log.Panic(r.Run(rs.address))
}

func (rs *Rest) getUserByName(c *gin.Context) {
	name := c.Query("name")
	if shouldAbort(c, name == "") {
		return
	}
	user, err := rs.service.FindUserByName(name)
	responseOkOrErrorIfPresent(c, user, err)
}

func (rs *Rest) getUserById(c *gin.Context) {
	id := c.Param("id")
	if idInt, abort := shouldAbortInt(c, id); abort {
		return
	} else {
		user, err := rs.service.FindUserById(idInt)
		responseOkOrErrorIfPresent(c, user, err)
	}
}

func (rs *Rest) authUser(c *gin.Context) {
	userName := c.GetHeader("userName")
	pwd := c.GetHeader("pwd")
	if shouldAbort(c, userName == "", pwd == "") {
		return
	}
	ok, err := rs.service.AuthUser(userName, pwd)
	responseOkOrErrorIfPresent(c, ok, err)
}

func (rs *Rest) deleteUserById(c *gin.Context) {
	id := c.Param("id")
	if idInt, abort := shouldAbortInt(c, id); abort {
		return
	} else {
		ok, err := rs.service.DeleteUserById(idInt)
		responseOkOrErrorIfPresent(c, ok, err)
	}
}

func (rs *Rest) updateUser(c *gin.Context) {
	user := &model.User{}
	if shouldAbortBind(c, user) {
		return
	}
	user, err := rs.service.CreateOrUpdateUser(user)
	responseOkOrErrorIfPresent(c, user, err)
}

func (rs *Rest) createUser(c *gin.Context) {
	user := &model.User{}
	if shouldAbortBind(c, user) {
		return
	}
	user, err := rs.service.CreateOrUpdateUser(user)
	responseOkOrErrorIfPresent(c, user, err)
}

func shouldAbortBind(c *gin.Context, user *model.User) bool {
	return shouldAbort(c, c.ShouldBindJSON(&user) != nil)
}

func shouldAbortInt(c *gin.Context, param string) (int, bool) {
	if abort := shouldAbort(c, param == ""); !abort {
		paramInt, ok := strconv.Atoi(param)
		return paramInt, shouldAbort(c, ok != nil)
	} else {
		return 0, abort
	}
}

func shouldAbort(c *gin.Context, validations ...bool) bool {
	for _, validation := range validations {
		if validation {
			c.AbortWithStatusJSON(400, "Bad request")
			break
		}
	}
	return c.IsAborted()
}

type Backs interface {
	*model.User | bool
}

func responseOkOrErrorIfPresent[B Backs](c *gin.Context, back B, err *model.ErrorWrap) {
	if err != nil {
		c.JSON(err.Code, gin.H{"msg": err.ErrorMsg()})
	} else {
		c.JSON(200, gin.H{"data": back})
	}
}
