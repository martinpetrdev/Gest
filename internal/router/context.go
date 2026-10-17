package router

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type RouteContext struct {
	req *http.Request
}

func NewRouteContext(req *http.Request) *RouteContext {
	return &RouteContext{req: req}
}

func (r *RouteContext) GetParam(key string) (string, error) {
	val := r.req.PathValue(key)
	if val == "" {
		return "", errors.New(fmt.Sprintf("Parameter %s not found", key))
	}

	return val, nil
}

func (r *RouteContext) GetParamInt(key string) (int, error) {
	i, err := strconv.Atoi(r.req.PathValue(key))
	if err != nil {
		return 0, errors.New(fmt.Sprintf("Parameter %s not found or is not an int", key))
	}

	return i, nil
}
