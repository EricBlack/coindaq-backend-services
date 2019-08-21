package test

import (
	"context"
	"testing"

	"fmt"

	"bx.com/user-service/bxgo"
	"bx.com/user-service/config"
	"bx.com/user-service/model"
	ctr "bx.com/user-service/controller"
	log "github.com/sirupsen/logrus"
	"bx.com/user-service/proto"
)

func init() {
	/*
	bxgo.CreateOrmEngin(config.DataSource{
		DriverName: "mysql",
		URI:        "postgres://postgres:postgres@localhost:5432/userinfo?sslmode=disable",
		MaxIdle:    10,
		MaxOpen:    5,
		ShowSQL:    true,
	})
	*/
	bxgo.CreateOrmEngin(config.DataSource{
		DriverName: "mysql",
		URI:        "root:123456@tcp(192.168.1.145:3306)/test?charset=utf8mb4",
		MaxIdle:    10,
		MaxOpen:    5,
		ShowSQL:    true,
	})
	bxgo.OrmEngin.Sync2(new(model.User))
}

func TestAuth(t *testing.T) {
	s := ctr.UserController{}
	// set up test cases
	testCase := struct {
		email       	string
		password    	string
		displayName 	string
		enableInvite 	bool
		inviteCode		string
	}{
		email:       	"blackeye@gmail.com",
		password:    	"123456",
		displayName: 	"Blackeye_Account",
	}
	registReq := &proto.RegisterReq{
		Email:       testCase.email,
		Password:    testCase.password,
		DisplayName: testCase.displayName,
		EnableCode:	 false,
	}
	resp, err := s.Signup(context.Background(), registReq)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignup got unexpected error")
	}
	if resp.Email != testCase.email {
		log.Errorf(err.Error())
		t.Errorf("TestSignup email=%v, wanted %v", resp.Email, testCase.email)
	}
}

func TestLoginPass (t *testing.T){
	authReq := &proto.AuthReq{
		Email:    "blackeye@gmail.com",
		Password: "123456",
		LoginIp: "192.168.0.2",
		DeviceId: "sfddghjllkjtgfe",
	}
	s := ctr.UserController{}
	resp, err := s.Signin(context.Background(), authReq)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignin got unexpected error")
	}
	if resp.Email != authReq.Email {
		log.Errorf(err.Error())
		t.Errorf("TestSignin email=%v, wanted %v", resp.Email, "blackeye@gmail.com")
	}
}

func TestLoginFail (t *testing.T){
	authReq := &proto.AuthReq{
		Email:    "blackeye@gmail.com",
		Password: "12345",
		LoginIp: "192.168.0.2",
		DeviceId: "sfddghjllkjtgfe",
	}
	s := ctr.UserController{}
	resp, err := s.Signin(context.Background(), authReq)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignin got unexpected error")
	}
	if resp.Email != authReq.Email {
		log.Errorf(err.Error())
		t.Errorf("TestSignin email=%v, wanted %v", resp.Email, "blackeye@gmail.com")
	}
}

func TestAuthWithInviteCode(t *testing.T){
	s := ctr.UserController{}
	// set up test cases
	testCase := struct {
		email       	string
		password    	string
		displayName 	string
		enableInvite 	bool
		inviteCode		string
	}{
		email:       	"blackeye1024@gmail.com",
		password:    	"123456",
		displayName: 	"Test 1024",
		enableInvite:	true,
		inviteCode:		"0UK4M8E3DA6CWY1UX",
	}
	registReq := &proto.RegisterReq{
		Email:       testCase.email,
		Password:    testCase.password,
		DisplayName: testCase.displayName,
		EnableCode:	 testCase.enableInvite,
		InviteCode:  testCase.inviteCode,
	}
	resp, err := s.Signup(context.Background(), registReq)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignup got unexpected error")
	}
	if resp.Email != testCase.email {
		log.Errorf(err.Error())
		t.Errorf("TestSignup email=%v, wanted %v", resp.Email, testCase.email)
	}

	authReq := &proto.AuthReq{
		Email:    testCase.email,
		Password: testCase.password,
	}
	resp, err = s.Signin(context.Background(), authReq)
	if err != nil {
		log.Errorf(err.Error())
		t.Errorf("TestSignin got unexpected error")
	}
	if resp.Email != testCase.email {
		log.Errorf(err.Error())
		t.Errorf("TestSignin email=%v, wanted %v", resp.Email, testCase.email)
	}
}

func TestForgetPasswordViaEmail(t *testing.T) {
	s := ctr.UserController{}
	// set up test cases
	testCase := struct {
		email string
	}{
		email: "blackeye124@gmail.com",
	}
	emailReq := &proto.ForgetEmailReq{
		Email: testCase.email,
	}
	resp, err := s.ForgetPasswordViaEmail(context.Background(), emailReq)
	if err != nil {
		t.Errorf("TestForgetPasswordViaEmail got unexpected error")
	}
	if resp.Information != testCase.email {
		t.Errorf("TestForgetPasswordViaEmail email=%v, wanted %v", resp.Information, testCase.email)
	}
}

func TestForgetPasswordViaPhone(t *testing.T) {
	s := ctr.UserController{}
	// set up test cases
	testCase := struct {
		phone string
	}{
		phone: "18500797779",
	}
	phoneReq := &proto.ForgetPhoneReq{
		Phone: testCase.phone,
	}
	resp, err := s.ForgetPasswordViaPhone(context.Background(), phoneReq)
	if err != nil {
		t.Errorf("TestForgetPasswordViaPhone got unexpected error")
	}
	if resp.Information != testCase.phone {
		t.Errorf("TestForgetPasswordViaPhone phone=%v, wanted %v", resp.Information, testCase.phone)
	}
}

func TestRecordMessageInfo(t *testing.T) {
	s := ctr.UserController{}
	// set up test cases
	msgReq := &proto.RecordMessageReq{
		Destination: "18500797779",
		Message:	 "[Coindaq]:您在币信科技请求的验证码为:8230",
		SendStatus:  proto.MsgSendStatus_Failed,
		ReturnMessage:"网络错误",
	}
	_, err := s.RecordMessageInfo(context.Background(), msgReq)
	if err != nil {
		t.Errorf("TestRecordMessageInfo got unexpected error: " +err.Error())
	}
}

func TestVerifyUserToken(t *testing.T){
	s := ctr.UserController{}
	// set up test cases
	userReq := &proto.UserReq{
		Token: "fae12faf-7c8b-4eec-be59-7bae3bbbf347",
	}
	_, err := s.VerifyUserToken(context.Background(), userReq)
	if err != nil {
		t.Errorf("TestVerifyUserToken got unexpected error: " +err.Error())
	}
}

func TestResetPassword(t *testing.T) {
	s := ctr.UserController{}
	// set up test cases
	testCase := struct {
		id     int64
		newPwd string
		code   string
	}{
		id:     6,
		newPwd: "test123",
		code:   "8111",
	}
	modifyReq := &proto.ModifyPasswordReq{
		Password: testCase.newPwd,
		Code:     testCase.code,
	}
	_, err := s.ResetPassword(context.Background(), modifyReq)
	if err != nil {
		t.Errorf("TestResetPassword got unexpected error")
	}
}

func TestUpdatePassword(t *testing.T){
	s := ctr.UserController{}
	// set up test cases
	testCase := struct {
		id     int64
		oldPwd string
		newPwd string
	}{
		id:     5,
		oldPwd:	"123456",
		newPwd: "test123",
	}
	modifyReq := &proto.UpdatePasswordReq{
		Id:       	 testCase.id,
		OldPassword: testCase.oldPwd,
		NewPassword: testCase.newPwd,
	}
	_, err := s.UpdatePassword(context.Background(), modifyReq)
	if err != nil {
		t.Errorf("TestUpdatePassword got unexpected error: " + err.Error())
	}
}

func TestUpdateUserInfo(t *testing.T) {
	s := ctr.UserController{}
	userId := int64(2)
	displayName := "Jacky"
	modify := &proto.ModifyUserReq{
		Id:          userId,
		DisplayName: displayName,
	}

	_, err := s.UpdateUserInfo(context.Background(), modify)
	if err != nil {
		t.Errorf("TestUpdateUserInfo got unexpected error")
	}
}

func TestGenPwd(t *testing.T) {
	genCode := model.GenPwd("123", "123")
	t.Logf("%s", genCode)
}

func TestDisableUser(t *testing.T){
	id := int64(6)
	resp, err :=model.DisableUser(id)
	if err != nil{
		log.Errorf(err.Error())
		t.Errorf("TestDisableUser got unexpected error")
	}

	if resp != id{
		t.Errorf("TestDisableUser id=%v, wanted %v", resp, id)
	}
}

func TestQueryUsers1(t *testing.T) {
	s := ctr.UserController{}

	displayName := "Test QQ"

	queryReq := &proto.QueryUserReq{
		DisplayName: displayName,
	}

	resp, err := s.QueryUsers(context.Background(), queryReq)
	if err != nil {
		t.Errorf("TestQueryUsers got unexpected error")
	}
	if len(resp.Users) != 0 {
		fmt.Printf("%s", resp.Users)
	} else {
	t.Errorf("TestQueryUsers code=%v, wanted %v", len(resp.Users), 2)
	fmt.Printf("No Users find.")
	}
}

func TestQueryUsers2(t *testing.T) {
	s := ctr.UserController{}

	queryReq := &proto.QueryUserReq{
		Activated: 1,
	}

	resp, err := s.QueryUsers(context.Background(), queryReq)
	if err != nil {
		t.Errorf("TestQueryUsers got unexpected error")
	}
	if len(resp.Users) != 0 {
		fmt.Printf("%s", resp.Users)
	} else {
		t.Errorf("TestQueryUsers code=%v, wanted %v", len(resp.Users), 2)
		fmt.Printf("No Users find.")
	}
}

func TestBindGoogleFactor(t *testing.T) {
	s := ctr.UserController{}
	idReq := &proto.IdReq{ Id:15 }
	resp, err := s.BindUserGoogleFactor(context.Background(), idReq)
	if err != nil {
		t.Errorf("%s", err)
	}else {
		t.Logf("%s", resp)
	}
}
