package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	v1 "staff/api/staff/v1"
	"time"
)

var staffClient v1.StaffClient
var conn *grpc.ClientConn

func main() {
	Init()
	TestCreateStaff()      // 创建用户
	rsp := TestListStaff() // 获取用户列表
	TestUpdateStaff(rsp)   // 更新用户
	TestDeleteStaff(rsp)   // 删除用户
	TestRecoveryStaff(rsp) // 恢复用户
	conn.Close()
}

// Init 初始化 grpc 链接 注意这里链接的 端口
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		panic("grpc link err" + err.Error())
	}
	staffClient = v1.NewStaffClient(conn)
}

func TestCreateStaff() {
	rsp, err := staffClient.CreateStaff(context.Background(), &v1.CreateStaffRequest{
		Mobile: fmt.Sprintf("1388888%d%d", time.Now().Minute(), time.Now().Second()),
		Name:   fmt.Sprintf("YWWW%d", 1),
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
		Name: fmt.Sprintf("YWWW%d", 22344),
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
