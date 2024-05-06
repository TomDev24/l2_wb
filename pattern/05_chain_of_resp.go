package main

import (
	"fmt"
	"strings"
)

type Handler interface {
	SetNext(handler Handler)
	HandleRequest(request string)
}

type BaseHandler struct {
	next Handler
}

func (b *BaseHandler) SetNext(handler Handler) {
	b.next = handler
}

func (b *BaseHandler) HandleRequest(request string) {
	if b.next != nil {
		b.next.HandleRequest(request)
	}
}

type AuthenticationHandler struct {
	BaseHandler
}

func (a *AuthenticationHandler) HandleRequest(request string) {
	if strings.HasPrefix(request, "auth ") {
		fmt.Println("AuthenticationHandler: обработка запроса", request)
	} else {
		a.BaseHandler.HandleRequest(request)
	}
}

type AuthorizationHandler struct {
	BaseHandler
}

func (a *AuthorizationHandler) HandleRequest(request string) {
	if strings.HasPrefix(request, "authz ") {
		fmt.Println("AuthorizationHandler: обработка запроса", request)
	} else {
		a.BaseHandler.HandleRequest(request)
	}
}

type ValidationHandler struct {
	BaseHandler
}

func (v *ValidationHandler) HandleRequest(request string) {
	if strings.HasPrefix(request, "validate ") {
		fmt.Println("ValidationHandler: обработка запроса", request)
	} else {
		v.BaseHandler.HandleRequest(request)
	}
}

func main() {
	// Handlers
	authenticationHandler := &AuthenticationHandler{}
	authorizationHandler := &AuthorizationHandler{}
	validationHandler := &ValidationHandler{}

	// Chain of calls
	authenticationHandler.SetNext(authorizationHandler)
	authorizationHandler.SetNext(validationHandler)

	authenticationHandler.HandleRequest("auth user:password")
	authenticationHandler.HandleRequest("authz user=admin")
	authenticationHandler.HandleRequest("validate data")
	authenticationHandler.HandleRequest("other request")
}
