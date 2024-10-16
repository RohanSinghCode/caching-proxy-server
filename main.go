package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var inMemoryCache = make(map[string]string)
var origin string

func main() {
	port := flag.Int("port", 80, "Port for proxy server")
	originFlag := flag.String("origin", "", "Origin value")
	clearCache := flag.Bool("clear-cache", false, "Clear the cache")
	flag.Parse()
	if *clearCache {
		for k := range inMemoryCache {
			delete(inMemoryCache, k)
		}
		fmt.Print("Cache Cleared!")
		return
	}
	if *originFlag == "" {
		fmt.Print("Origin is required")
		return
	}
	fmt.Printf("Starting proxy server on port: %v and origin: %v", *port, *originFlag)

	origin = *originFlag

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hey! Proxy Server Up !",
		})
	})

	// Ignore the favicon request by returning a 204 (No Content) response
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(204) // No Content
	})

	router.GET("/:path", handleRequest)

	router.Run(fmt.Sprintf(":%v", *port))
}

func handleRequest(c *gin.Context) {
	type Request struct {
		Path string `uri:"path" binding:"required"`
	}
	var request Request
	if err := c.ShouldBindUri(&request); err != nil {
		c.JSON(400, gin.H{"msg": "Path required"})
		return
	}

	response, isCacheHit, err := processRequest(request.Path)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if isCacheHit {
		c.Header("X-CACHE", "HIT")
	} else {
		c.Header("X-CACHE", "MISS")
	}
	c.JSON(200, gin.H{"response": response})
}

func processRequest(path string) (string, bool, error) {
	cachedResponse, ok := inMemoryCache[path]
	if !ok {
		requestUrl := fmt.Sprintf("%v/%v", origin, path)
		fmt.Print(requestUrl)
		resp, err := http.Get(requestUrl)
		if err != nil {
			return "", false, err
		}

		defer resp.Body.Close()

		//Read the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return "", false, err
		}

		// Cache the response
		inMemoryCache[path] = string(body)

		return string(body), false, nil
	} else {
		return cachedResponse, true, nil
	}
}
