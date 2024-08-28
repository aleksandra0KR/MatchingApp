// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/addPlaylist": {
            "get": {
                "description": "Render the template to add a playlist.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "playlists"
                ],
                "summary": "redirect to post request with html page",
                "responses": {}
            }
        },
        "/createPlaylist": {
            "post": {
                "description": "Create a new playlist for the authenticated user in their profile based on which program will find match.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "playlists"
                ],
                "summary": "Add a new playlist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Playlist ID",
                        "name": "playlistID",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Username",
                        "name": "username",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Playlist created successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/match": {
            "get": {
                "description": "Match a playlist and render the template.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/html"
                ],
                "tags": [
                    "playlists"
                ],
                "summary": "Match a playlist",
                "responses": {}
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8080",
	BasePath:         "/MatchingApp",
	Schemes:          []string{},
	Title:            "Matching App",
	Description:      "finding your spotify match",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
