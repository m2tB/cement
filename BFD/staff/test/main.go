package main

import (
	"google.golang.org/grpc"
	v1 "staff/api/staff/v1"
)

var staffClient v1.StaffClient
var conn *grpc.ClientConn

// Init 初始化 grpc 链接 注意这里链接的 端口
func Init() {
	var err error
	conn, err = grpc.Dial("127.0.0.1:9000", grpc.WithInsecure())
	if err != nil {
		panic("grpc link err" + err.Error())
	}
	staffClient = v1.NewStaffClient(conn)
}

func main() {
	Init()
	teamStaff()
	conn.Close()
}

func staff() {
	TestCreateStaff()      // 创建用户
	rsp := TestListStaff() // 获取用户列表
	TestUpdateStaff(rsp)   // 更新用户
	TestDeleteStaff(rsp)   // 删除用户
	TestRecoveryStaff(rsp) // 恢复用户
}

func team() {
	TestCreateTeam()      // 创建team
	rsp := TestListTeam() // 获取team列表
	TestUpdateTeam(rsp)   // 更新team
	TestDeleteTeam(rsp)   // 删除team
}

func teamStaff() {
	TestInviteStaff()
	TestListTeamStaff()
	TestListStaffTeam()
	TestExpelStaff()
}
