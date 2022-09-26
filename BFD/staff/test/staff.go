package main

import (
	"context"
	"fmt"
	v1 "staff/api/staff/v1"
	"time"
)

func TestCreateStaff() {
	rsp, err := staffClient.CreateStaff(context.Background(), &v1.CreateStaffRequest{
		Mobile: fmt.Sprintf("1388888%d%d", time.Now().Minute(), time.Now().Second()),
		Name:   fmt.Sprintf("YWW%d", 1),
	})
	if err != nil {
		panic("grpc 创建用户失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}

func TestListStaff() *v1.StaffReply {
	rsp, err := staffClient.ListStaff(context.Background(), &v1.ListStaffRequest{})
	if err != nil {
		panic("grpc 获取用户列表失败" + err.Error())
	}
	fmt.Println(rsp.Data[0])
	return rsp.Data[0]
}

func TestUpdateStaff(s *v1.StaffReply) {
	fmt.Println(s.Id)
	rsp, err := staffClient.UpdateStaff(context.Background(), &v1.UpdateStaffRequest{
		Id:   s.Id,
		Name: fmt.Sprintf("YW%d", 22),
	})
	if err != nil {
		panic("grpc 更新用户失败" + err.Error())
	}
	fmt.Println(rsp)
}

func TestDeleteStaff(s *v1.StaffReply) {
	fmt.Println(s.Id)
	rsp, err := staffClient.DeleteStaff(context.Background(), &v1.DeleteStaffRequest{
		Id: s.Id,
	})
	if err != nil {
		panic("grpc 删除用户失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}

func TestRecoveryStaff(s *v1.StaffReply) {
	fmt.Println(s.Id)
	rsp, err := staffClient.RecoveryStaff(context.Background(), &v1.RecoveryStaffRequest{
		Id: s.Id,
	})
	if err != nil {
		panic("grpc 恢复用户失败" + err.Error())
	}
	fmt.Println(rsp.Exec)
}
