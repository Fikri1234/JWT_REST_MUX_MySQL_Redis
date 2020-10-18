package model

import "net/http"

// Route struct
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
	Protected   bool
}

// RoutePrefix struct
type RoutePrefix struct {
	Prefix    string
	SubRoutes []Route
}

// AppRoutes array of RoutePrefix
var AppRoutes []RoutePrefix
