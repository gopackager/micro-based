package config

import "github.com/gin-gonic/gin"

type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	DBName   string
	DBDriver string
}

type Routing struct {
	Method  string
	Path    string
	Handler gin.HandlerFunc
	Mw      func(*gin.Context)
}

type Application struct {
	Mode    string
	Version string
	Port    string
}
