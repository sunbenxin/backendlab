// Package rest provides ...
package rest

type BaseResp struct {
	Meta Meta
	Data interface{}
}

type Meta struct {
	Code    uint64
	Message string
}
