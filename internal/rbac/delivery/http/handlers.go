package http

import (
	"github.com/casbin/casbin/v2"
	"github.com/csc13010-student-management/internal/rbac"
	"github.com/gin-gonic/gin"
)

type rbacHandlers struct {
	e *casbin.Enforcer
}

func NewRbacHandlers(
	e *casbin.Enforcer,
) rbac.IRbacHandlers {
	return &rbacHandlers{
		e: e,
	}
}

type Role struct {
	User   string `json:"user"`
	Role   string `json:"role"`
	Domain string `json:"domain"`
}

type RoleData struct {
	Role   string   `json:"role"`
	API    []string `json:"api"`
	Method []string `json:"method"`
}

// AddRole implements rbac.IRbacHandlers.
func (rh *rbacHandlers) AddRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RoleData

		// read data and assign to data
		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"error": "binding data failed"})
			return
		}

		// if no endpoint is provided, we will add the default endpoint
		if data.API == nil {
			data.API = append(data.API, "http://localhost:8080")
		}

		// add policy rules
		for _, api := range data.API {
			for _, method := range data.Method {
				_, err := rh.e.AddPolicy(data.Role, api, method)
				if err != nil {
					c.JSON(400, gin.H{"error": "adding policy failed"})
					return
				}
			}
		}

		c.JSON(200, gin.H{"message": "role added successfully"})
	}
}

// DeleteRole implements rbac.IRbacHandlers.
func (rh *rbacHandlers) DeleteRole() gin.HandlerFunc {
	return func(c *gin.Context) {
	}
}

type UserRole struct {
	UserID []string `json:"user_id"`
	Role   string   `json:"role"`
}

// AddRoleForUser implements rbac.IRbacHandlers.
func (rh *rbacHandlers) AddRoleForUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data UserRole

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"error": "binding data failed"})
			return
		}

		for _, id := range data.UserID {
			_, err := rh.e.AddGroupingPolicy(id, data.Role)
			if err != nil {
				c.JSON(400, gin.H{"error": "adding policy failed"})
				return
			}
		}

		c.JSON(200, gin.H{"message": "role added successfully"})
	}
}

// DeleteRoleForUser implements rbac.IRbacHandlers.
func (rh *rbacHandlers) DeleteRoleForUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data UserRole

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"error": "binding data failed"})
			return
		}

		for _, id := range data.UserID {
			_, err := rh.e.RemoveGroupingPolicy(id, data.Role)
			if err != nil {
				c.JSON(400, gin.H{"error": "removing policy failed"})
				return
			}
		}

		c.JSON(200, gin.H{"message": "role removed successfully"})
	}
}

// AddAPIForRole implements rbac.IRbacHandlers.
func (rh *rbacHandlers) AddAPIForRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RoleData

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"error": "binding data failed"})
			return
		}
		// if no method is provided, add the default method GET
		if data.Method == nil {
			data.Method = append(data.Method, "GET")
		}

		for _, api := range data.API {
			for _, method := range data.Method {
				_, err := rh.e.AddPolicy(data.Role, api, method)
				if err != nil {
					c.JSON(400, gin.H{"error": "adding policy failed"})
					return
				}
			}
		}

		c.JSON(200, gin.H{"message": "api added successfully"})
	}
}

type RoleAPI struct {
	Role string   `json:"role"`
	API  []string `json:"api"`
}

// DeleteAPIForRole implements rbac.IRbacHandlers.
func (rh *rbacHandlers) DeleteAPIForRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data RoleAPI

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"error": "binding data failed"})
			return
		}

		allAction, err := rh.e.GetAllActions()
		if err != nil {
			c.JSON(400, gin.H{"error": "getting all actions failed"})
			return
		}

		for _, api := range data.API {
			for _, action := range allAction {
				rh.e.RemovePolicy(data.Role, api, action)
			}
		}
		// if no endpoint is left, add http://localhost:8080
		filteredPolicy, err := rh.e.GetFilteredPolicy(0, data.Role)
		if err != nil {
			c.JSON(400, gin.H{"error": "getting filtered policy failed"})
			return
		}

		if (len(filteredPolicy)) == 0 {
			_, err := rh.e.AddPolicy(data.Role, "http://localhost:8080", "GET")
			if err != nil {
				c.JSON(400, gin.H{"error": "adding policy failed"})
				return
			}
		}

		c.JSON(200, gin.H{"message": "api removed successfully"})
	}
}

type APIData struct {
	API    string   `json:"api"`
	Role   []string `json:"role"`
	Method []string `json:"method"`
}

// AddRoleForAPI implements rbac.IRbacHandlers.
func (rh *rbacHandlers) AddRoleForAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data APIData

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"error": "binding data failed"})
			return
		}

		for _, role := range data.Role {
			for _, method := range data.Method {
				_, err := rh.e.AddPolicy(role, data.API, method)
				if err != nil {
					c.JSON(400, gin.H{"error": "adding policy failed"})
					return
				}
			}
		}

		c.JSON(200, gin.H{"message": "role added successfully"})
	}
}

type APIRole struct {
	API  string   `json:"api"`
	Role []string `json:"role"`
}

// DeleteRoleForAPI implements rbac.IRbacHandlers.
func (rh *rbacHandlers) DeleteRoleForAPI() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data APIRole

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"error": "binding data failed"})
			return
		}

		allAction, err := rh.e.GetAllActions()
		if err != nil {
			c.JSON(400, gin.H{"error": "getting all actions failed"})
			return
		}

		for _, role := range data.Role {
			for _, action := range allAction {
				_, err := rh.e.RemovePolicy(role, data.API, action)
				if err != nil {
					c.JSON(400, gin.H{"error": "removing policy failed"})
					return
				}
			}
		}

		c.JSON(200, gin.H{"message": "role removed successfully"})
	}
}

type User struct {
	UserID string `json:"user_id"`
}

// CheckAuth implements rbac.IRbacHandlers.
func (rh *rbacHandlers) CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data User

		if err := c.ShouldBind(&data); err != nil {
			c.JSON(400, gin.H{"error": "binding data failed"})
			return
		}

		path := "http://localhost:8080" + c.Request.URL.Path

		ok, err := rh.e.Enforce(data.UserID, path, "GET")
		if !ok || err != nil {
			c.JSON(403, gin.H{"error": "forbidden"})
			return
		}

		c.Next()
	}
}

// Notification implements rbac.IRbacHandlers.
func (rh *rbacHandlers) Notification() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "notification"})
	}
}
