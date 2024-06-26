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
        "/admin/make": {
            "post": {
                "description": "This endpoint allows an admin to promote a user to an admin role.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Make user admin",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The email of the user to be promoted to admin.",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.EmailInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"message\\\":\\\"User example@example.com is now an admin\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\":\\\"error validating email: invalid format\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\":\\\"error making admin:\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/admin/profile": {
            "get": {
                "description": "This endpoint retrieves the profile information for all users.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get all user profiles",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A list of all user profiles.",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/profile.Profile"
                            }
                        }
                    },
                    "404": {
                        "description": "{\"error\":\"No profiles found\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\"error\":\"Failed to retrieve profiles\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/admin/quote/approve/{id}": {
            "post": {
                "description": "Approve a quote by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Approve quote",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Quote ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Quote approved successfully",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "403": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Quote not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to approve quote",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/quote/unapproved": {
            "get": {
                "description": "Get all unapproved quotes",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Get unapproved quotes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of unapproved quotes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/quote.Quote"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "No unapproved quotes found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to get unapproved quotes",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/admin/remove": {
            "delete": {
                "description": "This endpoint allows an admin to revoke admin rights from a user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Remove user admin rights",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The email of the user to have admin rights revoked.",
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/admin.EmailInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"message\\\":\\\"User example@example.com admin rights have been revoked\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\":\\\"error validating email: invalid format\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\":\\\"error removing admin rights:\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/profile/create": {
            "post": {
                "description": "This endpoint allows a user to create their profile.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Create user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "{\\\"message\\\":\\\"Profile created successfully\\\", \\\"user_email\\\":\\\"example@example.com\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\":\\\"Profile already exists\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "{\\\"error\\\":\\\"Unauthorized\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\":\\\"Failed to create profile\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/profile/delete/{id}": {
            "delete": {
                "description": "This endpoint allows a user to delete their profile. Admins can delete any profile.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Delete user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The ID of the user whose profile is to be deleted.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"message\\\":\\\"Profile deleted successfully\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "{\\\"error\\\":\\\"Unauthorized\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "{\\\"error\\\":\\\"not authorized\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "{\\\"error\\\":\\\"Profile not found\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\":\\\"Failed to delete profile\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/profile/update": {
            "put": {
                "description": "This endpoint allows a user to update their profile information.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Update user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The new bio for the user.",
                        "name": "bio",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    },
                    {
                        "description": "The new username for the user.",
                        "name": "username",
                        "in": "body",
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"message\\\":\\\"Profile updated successfully\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\":\\\"Invalid request body\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "{\\\"error\\\":\\\"Unauthorized\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "{\\\"error\\\":\\\"Profile not found\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\":\\\"Failed to update profile\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/profile/{id}": {
            "get": {
                "description": "This endpoint retrieves the profile information for a specific user.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profile"
                ],
                "summary": "Get user profile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The ID of the user whose profile is to be retrieved.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The user's profile information.",
                        "schema": {
                            "$ref": "#/definitions/profile.Profile"
                        }
                    },
                    "404": {
                        "description": "{\"error\":\"Profile not found\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\"error\":\"Failed to retrieve profile\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/quote": {
            "get": {
                "description": "This endpoint retrieves quotes. Admins can retrieve all quotes, while other users can only retrieve approved quotes.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quote"
                ],
                "summary": "Get quotes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "A list of quotes.",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "array",
                                "items": {
                                    "$ref": "#/definitions/quote.Quote"
                                }
                            }
                        }
                    },
                    "401": {
                        "description": "{\\\"error\\\":\\\"Unauthorized\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "{\\\"error\\\":\\\"Quote not found\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\":\\\"Failed to get quote\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/quote/create": {
            "post": {
                "description": "This endpoint allows a user to submit a new quote.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quote"
                ],
                "summary": "Create a new quote",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "The quote to be submitted.",
                        "name": "quote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/quote.QuoteRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "{\\\"message\\\":\\\"Quote created successfully\\\", \\\"quote\\\":\\\"Example quote text\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\":\\\"Invalid request body\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "{\\\"error\\\":\\\"Unauthorized\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\":\\\"Failed to create quote\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/quote/delete/{id}": {
            "delete": {
                "description": "This endpoint allows a user to delete an existing quote. Admins can delete any quote.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quote"
                ],
                "summary": "Delete an existing quote",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The ID of the quote to be deleted.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"message\\\":\\\"Quote deleted successfully\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\":\\\"Failed to delete quote\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "{\\\"error\\\":\\\"Unauthorized access\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "{\\\"error\\\":\\\"not authorized\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "{\\\"error\\\":\\\"Quote not found\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/quote/update": {
            "put": {
                "description": "This endpoint allows a user to update an existing quote. Admins can update any quote.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Quote"
                ],
                "summary": "Update an existing quote",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The ID of the quote to be updated.",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "The updated quote text.",
                        "name": "quote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/quote.QuoteUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "{\\\"message\\\":\\\"Quote updated successfully\\\", \\\"quote\\\":\\\"Updated quote text\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "{\\\"error\\\":\\\"Invalid request body\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "401": {
                        "description": "{\\\"error\\\":\\\"Unauthorized access\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "403": {
                        "description": "{\\\"error\\\":\\\"not authorized\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "{\\\"error\\\":\\\"Quote not found\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "{\\\"error\\\":\\\"Failed to update quote\\\"}",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/quote/{profile-id}": {
            "get": {
                "description": "Retrieve quotes by user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quote"
                ],
                "summary": "Get quotes by user ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "JWT token",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "profile-id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of quotes",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/quote.Quote"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Quote not found",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Failed to get quote",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "admin.EmailInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "profile.Profile": {
            "type": "object",
            "properties": {
                "bio": {
                    "type": "string"
                },
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "quote.Quote": {
            "type": "object",
            "properties": {
                "approved": {
                    "type": "boolean"
                },
                "created_at": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "quote": {
                    "type": "string"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "quote.QuoteRequest": {
            "type": "object",
            "properties": {
                "quote": {
                    "type": "string"
                }
            }
        },
        "quote.QuoteUpdateRequest": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "quote": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "localhost:8082",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "FireGo",
	Description:      "A server for a simple Go application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
