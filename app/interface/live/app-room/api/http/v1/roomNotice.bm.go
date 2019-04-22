// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: roomNotice.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	roomNotice.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathRoomNoticeBuyGuard = "/live.approom.v1.RoomNotice/buy_guard"

// ====================
// RoomNotice Interface
// ====================

// 房间提示 相关服务
type RoomNoticeBMServer interface {
	// 是否弹出大航海购买提示
	BuyGuard(ctx context.Context, req *RoomNoticeBuyGuardReq) (resp *RoomNoticeBuyGuardResp, err error)
}

var v1RoomNoticeSvc RoomNoticeBMServer

func roomNoticeBuyGuard(c *bm.Context) {
	p := new(RoomNoticeBuyGuardReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1RoomNoticeSvc.BuyGuard(c, p)
	c.JSON(resp, err)
}

// RegisterV1RoomNoticeService Register the blademaster route with middleware map
// midMap is the middleware map, the key is defined in proto
func RegisterV1RoomNoticeService(e *bm.Engine, svc RoomNoticeBMServer, midMap map[string]bm.HandlerFunc) {
	v1RoomNoticeSvc = svc
	e.GET("/xlive/app-room/v1/roomNotice/buy_guard", roomNoticeBuyGuard)
}

// RegisterRoomNoticeBMServer Register the blademaster route
func RegisterRoomNoticeBMServer(e *bm.Engine, server RoomNoticeBMServer) {
	v1RoomNoticeSvc = server
	e.GET("/live.approom.v1.RoomNotice/buy_guard", roomNoticeBuyGuard)
}