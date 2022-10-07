package main

import (
	"context"
	"fmt"
	v1 "staff/api/staff/v1"
)

func TestCreateSubTeam() {
	rsp, err := staffClient.CreateSubTeam(context.Background(), &v1.CreateSubTeamRequest{
		Tid: 2,
		Pid: 1,
	})
	if err != nil {
		panic("grpc 创建teamTeam失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}

func TestUpdateSubTeam() {
	rsp, err := staffClient.UpdateSubTeam(context.Background(), &v1.UpdateSubTeamRequest{
		Tid:  2,
		OPid: 1,
		Pid:  3,
	})
	if err != nil {
		panic("grpc 更新teamTeam失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}

func TestDeleteSubTeam() {
	rsp, err := staffClient.DeleteSubTeam(context.Background(), &v1.DeleteSubTeamRequest{
		Tid: 2,
	})
	if err != nil {
		panic("grpc 删除teamTeam失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}

func TestListSubTeam() {
	rsp, err := staffClient.ListTeamSubTeam(context.Background(), &v1.ListTeamSubTeamRequest{
		Pid: 3,
	})
	if err != nil {
		panic("grpc 获取teamTeam列表失败" + err.Error())
	}
	fmt.Println(rsp.Data)
}
