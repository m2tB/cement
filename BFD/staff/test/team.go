package main

import (
	"context"
	"fmt"
	v1 "staff/api/staff/v1"
)

func TestCreateTeam() {
	rsp, err := staffClient.CreateTeam(context.Background(), &v1.CreateTeamRequest{
		CName:       "YWWW",
		PreTeamCode: "YWWW",
	})
	if err != nil {
		panic("grpc 创建team失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}

func TestListTeam() *v1.TeamReply {
	rsp, err := staffClient.ListTeam(context.Background(), &v1.ListTeamRequest{})
	if err != nil {
		panic("grpc 获取team列表失败" + err.Error())
	}
	fmt.Println(rsp.Data[0])
	return rsp.Data[0]
}

func TestUpdateTeam(s *v1.TeamReply) {
	fmt.Println(s.Id)
	rsp, err := staffClient.UpdateTeam(context.Background(), &v1.UpdateTeamRequest{
		Id:    s.Id,
		CName: fmt.Sprintf("YW%d", 22),
	})
	if err != nil {
		panic("grpc 更新team失败" + err.Error())
	}
	fmt.Println(rsp)
}

func TestDeleteTeam(s *v1.TeamReply) {
	fmt.Println(s.Id)
	rsp, err := staffClient.DeleteTeam(context.Background(), &v1.DeleteTeamRequest{
		Id: s.Id,
	})
	if err != nil {
		panic("grpc 删除team失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}
