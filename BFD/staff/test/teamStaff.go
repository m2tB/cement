package main

import (
	"context"
	"fmt"
	v1 "staff/api/staff/v1"
)

func TestInviteStaff() {
	rsp, err := staffClient.InviteStaff(context.Background(), &v1.InviteStaffRequest{
		Tid:  1,
		Sids: []int64{1},
	})
	if err != nil {
		panic("grpc 创建teamStaff失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}

func TestExpelStaff() {
	rsp, err := staffClient.ExpelStaff(context.Background(), &v1.ExpelStaffRequest{
		Tid: 1,
		Sid: 1,
	})
	if err != nil {
		panic("grpc 删除teamStaff失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}

func TestListTeamStaff() {
	rsp, err := staffClient.ListTeamStaff(context.Background(), &v1.ListTeamStaffRequest{
		Tid: 1,
	})
	if err != nil {
		panic("grpc 获取teamStaff列表失败" + err.Error())
	}
	fmt.Println(rsp.Total)
}

func TestListStaffTeam() {
	rsp, err := staffClient.ListStaffTeam(context.Background(), &v1.ListStaffTeamRequest{
		Sid: 1,
	})
	if err != nil {
		panic("grpc 获取teamStaff列表失败" + err.Error())
	}
	fmt.Println(rsp.Total)
}
