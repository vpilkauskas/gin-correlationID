# Gin-CorrelationID
[![Go Report Card](https://goreportcard.com/badge/github.com/grayMou5e/gin-correlationID)](https://goreportcard.com/report/github.com/grayMou5e/gin-correlationID)

CorrelationID middleware for Golang [Gin Web framework](https://github.com/gin-gonic/gin) 

# Use case
Gin-CorraltaionID middleware is used to check if there is an corralation id added into http request header. If that header is not found middleware adds it to http request.

Correlation ID enriches logs and reduces your headache when it comes to tracing up an actions produced by application.

# How to use it in Gin web-api

```golang
func main(){
    router := gin.New()
	router.Use(ginCorrelationID.CorrelationIDMiddleware())

    router.GET("api/test", func(c *gin.Context) {
		c.String(http.StatusOK, c.Request.Header.Get("CorrelationID"))
	})
}
```

# Dependencies
* [Gin Web framework](https://github.com/gin-gonic/gin) - "gopkg.in/gin-gonic/gin.v1"
* [UUID](https://github.com/nu7hatch/gouuid) - "github.com/nu7hatch/gouuid"
