package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	e                  *casbin.Enforcer
	introspectEndPoint = flag.String("introspect-endpoint", "http://zitadel:8080/oauth/v2/introspect", "The introspect endpoint")
	issuer             = flag.String("issuer", "http://zitadel:8080", "issuer of your ZITADEL instance (in the form: https://<instance>.zitadel.cloud or https://<yourdomain>)")
	api                = flag.String("api", "zitadel", "gRPC endpoint of your ZITADEL instance (in the form: <instance>.zitadel.cloud:443 or <yourdomain>:443)")
	key                = flag.String("key", "./key.json", "path to the key file")
	model              = flag.String("model", "/app/model.conf", "path to the casbin model file")
	policy             = flag.String("policy", "/app/policy.csv", "path to the casbin policy file")
	userPath           = flag.String("users", "/app/users.json", "path to the users file")
)

func InitEnforcer() {
	var err error
	e, err = casbin.NewEnforcer(*model, *policy)
	if err != nil {
		log.Fatalf("failed to create enforcer: %v", err)
	}
}

func main() {
	flag.Parse()
	InitEnforcer()

	users := ReadUsers()

	r := gin.Default()
	r.Use(TokenAuthMiddleware())

	r.POST("/check", func(c *gin.Context) {
		sub, exists := c.Get("sub")
		if !exists {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "sub not found"})
			c.Abort()
			return
		}

		user := users[fmt.Sprintf("%v", sub)]
		if user.UID == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "user not found"})
			c.Abort()
			return
		}
		log.Printf("User: %v", user)

		target := c.Request.Header.Get("X-RESOURCE-TARGET")
		method := c.Request.Header.Get("X-RESOURCE-METHOD")
		log.Printf("target: %s method: %s", target, method)
		f := Enforce(user.Roles, user.Tid, target, method)
		if !f {
			c.JSON(http.StatusForbidden, gin.H{"message": "forbidden"})
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "ok"})
	})

	err := r.Run("0.0.0.0:17010")
	if err != nil {
		panic(err)
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")

		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API token required"})
			c.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")

		form := fmt.Sprintf("token_type_hint=access_token&scope=openid&token=%s", token)
		payload := strings.NewReader(form)

		req, err := http.NewRequest("POST", *introspectEndPoint, payload)
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.SetBasicAuth(os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating introspection request"})
			c.Abort()
			return
		}

		client := &http.Client{}
		resp, err := client.Do(req)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending introspection request"})
			c.Abort()
			return
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading introspection response"})
			c.Abort()
			return
		}

		var introspectionResponse IntrospectionResponse
		err = json.Unmarshal(body, &introspectionResponse)

		if err != nil || !introspectionResponse.Active {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API token"})
			c.Abort()
			return
		}
		c.Set("user", introspectionResponse.UserName)
		c.Set("sub", introspectionResponse.Sub)
		c.Next()
	}
}

type IntrospectionResponse struct {
	Active   bool   `json:"active"`
	UserName string `json:"username"`
	Sub      string `json:"sub"`
}

type User struct {
	UID   string   `json:"uid"`
	Roles []string `json:"roles"`
	Tid   string   `json:"tid"`
}

func ReadUsers() map[string]User {
	file, err := os.Open(*userPath)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	var users []User
	_ = json.Unmarshal(byteValue, &users)
	userMap := make(map[string]User)
	for _, user := range users {
		userMap[user.UID] = user
	}
	return userMap
}

func Enforce(roles []string, dom string, obj string, act string) bool {
	reqs := GetEnforceRequests(roles, dom, obj, act)
	results, err := e.BatchEnforce(reqs)
	if err != nil {
		return false
	}
	for _, value := range results {
		if value {
			return true
		}
	}
	return false
}

func GetEnforceRequests(roles []string, dom string, obj string, act string) [][]interface{} {
	var reqs [][]interface{}
	for _, role := range roles {
		reqs = append(reqs, []interface{}{role, dom, obj, act})
	}
	return reqs
}
