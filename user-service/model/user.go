package model

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"time"

	"bx.com/user-service/bxgo"
	"bx.com/user-service/config"
	"bx.com/user-service/utils"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"regexp"
)

type User struct {
	Id              int64
	Email           string    `xorm:"email text notnull"`
	Password        string    `xorm:"password text notnull"`
	Salt            string    `xorm:"salt text notnull"`
	DisplayName     string    `xorm:"display_name text notnull"`
	PhoneNumber     string    `xorm:"phone_number text"`
	Kind            int32     `xorm:"kind int notnull"`
	Activated       int32     `xorm:"activated int notnull"`
	CountryCode     string    `xorm:"country_code text notnull"`
	IdentityType    int32     `xorm:"identity_type text"`
	IdentityId      string    `xorm:"identity_id text"`
	RealName        string    `xorm:"real_name text"`
	Disabled        int32     `xorm:"disabled integer default 0"`
	PhotoFront      string    `xorm:"photo_front text"`
	PhotoBack       string    `xorm:"photo_back text"`
	PhotoHand       string    `xorm:"photo_hand text"`
	InviteCode      string    `xorm:"invite_code text"`
	RegisterIp      string    `xorm:"register_ip text"`
	DeviceId        string    `xorm:"device_id text"`
	AccessToken     string    `xorm:"access_token text"`
	CreatedAt       time.Time `xorm:"created_at datetime"`
	UpdatedAt       time.Time `xorm:"updated_at datetime"`
	ActivatedAt     time.Time `xorm:"activated_at datetime"`
	DisabledAt      time.Time `xorm:"disabled_at datetime"`
	PaymentPassword string    `xorm:"payment_password text"`
	GoogleQRImage   string    `xorm:"google_qrimage text"`
	GoogleIsBind    int32     `xorm:"google_is_bind int"`
	LockTime        time.Time `xorm:"lock_time datetime"`
}

type UserFilter struct {
	Email       string
	DisplayName string
	PhoneNumber string
	RealName    string
	Kind        int32
	Activated   int32
	CountryCode string
	IdentityId  string
	Disabled    int32
}

func (u User) TableName() string {
	return "users"
}

func RandSalt() string {
	baseStr := "0123456789"
	bytes := []byte(baseStr)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 6; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	log.Infof("Salt: %s", result)
	return string(result)
}

func GenPwd(code string, salt string) string {
	basePwd := code + "@" + salt
	baseSha := sha256.Sum256([]byte(basePwd))
	shaStr := string(baseSha[:])
	base64Str := base64.StdEncoding.EncodeToString([]byte(shaStr))
	log.Infof("Password: %s", base64Str)

	return base64Str
}

func GetUserByID(id int64) (User, error) {
	user := User{}
	_, err := bxgo.OrmEngin.Id(id).Get(&user)
	if err != nil {
		return user, err
	}
	return user, err
}

func GetUserByEmail(email string) (User, error) {
	user := User{}
	_, err := bxgo.OrmEngin.Where("email = ?", email).Get(&user)
	if err != nil {
		return user, err
	}

	return user, err
}

func GetUserByPhone(phone string) (User, error) {
	user := User{}
	_, err := bxgo.OrmEngin.Where("phone_number = ?", phone).Get(&user)
	if err != nil {
		return user, err
	}
	return user, err
}

func CreateUser(u User) (int64, error) {
	user := User{
		Email:        u.Email,
		Password:     u.Password,
		Salt:         u.Salt,
		RealName:     u.RealName,
		DisplayName:  u.DisplayName,
		CountryCode:  u.CountryCode,
		RegisterIp:   u.RegisterIp,
		DeviceId:     u.DeviceId,
		Kind:         UserKindPerson,
		CreatedAt:    time.Now(),
		Activated:    False,
		Disabled:     False,
		GoogleIsBind: False,
	}
	_, err := bxgo.OrmEngin.Insert(&user)
	if err != nil {
		return user.Id, err
	}
	return user.Id, nil
}

func ActivateUser(id int64, secret string) (int64, error) {
	factor, err := GetVaildTwoFactor(id, secret)
	if err != nil {
		return -1, status.Error(codes.Internal, "Find factor error")
	}
	if factor.Id == 0 {
		return 0, status.Error(codes.InvalidArgument, "Cannot find related verify factor or timeout.")
	}

	user, err := GetUserByID(id)
	if err != nil {
		return -1, status.Error(codes.Internal, "Find user error")
	}

	//添加事物处理更新user状态和factor状态
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	// add Begin() before any action
	err = session.Begin()
	user.Activated = True
	user.ActivatedAt = time.Now()
	user.InviteCode = GenInviteCode()
	affected, err := session.Cols("activated", "activated_at", "invite_code").Update(&user, &User{Id: user.Id})
	if err != nil {
		session.Rollback()
		return affected, err
	}

	affected = 0
	factor.Activated = True
	affected, err = session.Cols("activated").Update(&factor, &TwoFactor{Id: factor.Id})

	if err != nil {
		session.Rollback()
		return affected, err
	}

	// add Commit() after all actions
	if err = session.Commit(); err != nil {
		return affected, err
	}

	return affected, nil
}

func DisableUser(id int64) (int64, error) {
	user, err := GetUserByID(id)
	if err != nil {
		return user.Id, err
	}

	user.Disabled = True
	user.DisabledAt = time.Now()
	_, err = bxgo.OrmEngin.Cols("disabled", "disabled_at").Update(&user, &User{Id: user.Id})
	if err != nil {
		return user.Id, err
	}

	return user.Id, nil
}

func ResetPassword(password, code string) (int64, error) {
	user, err := QueryUserByToken(code)
	if err != nil {
		return -1, err
	}

	if user.Id == 0  {
		return -1, errors.New("Reset code is not correct or timeout.")
	}

	encryPwd := GenPwd(password, user.Salt)

	user.Password = encryPwd
	user.AccessToken = ""
	user.UpdatedAt = time.Now()
	user.LockTime = time.Now().Add(24 * time.Hour)

	//添加事物处理保存代码和code状态更新
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	// add Begin() before any action
	err = session.Begin()
	affected, err := session.Where("id=? ", user.Id).Cols("password", "updated_at", "lock_time", "access_token").Update(&user)
	if err != nil {
		session.Rollback()
		return affected, err
	}

	if err != nil {
		session.Rollback()
		return affected, err
	}

	// add Commit() after all actions
	if err = session.Commit(); err != nil {
		return affected, err
	}

	return affected, nil
}

func BindPhone(id int64, phone, code string) (int64, error) {
	factor, err := GetVaildTwoFactor(id, code)
	if err != nil {
		return -1, err
	}
	if factor.Id == 0 {
		return factor.Id, status.Error(codes.Internal, "Code is not correct or timeout.")
	}

	user, err := QueryUserById(id)
	if err != nil {
		return -1, err
	}
	if user.Email == "" {

		return -1, errors.New("No user found.")
	}

	user.PhoneNumber = phone
	user.UpdatedAt = time.Now()
	user.LockTime = time.Now().Add(24 * time.Hour) //添加锁定时间

	//添加事物处理保存代码和code状态更新
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	// add Begin() before any action
	err = session.Begin()
	affected, err := session.Where("id=?", user.Id).Cols("phone_number", "updated_at", "lock_time").Update(&user)
	if err != nil {
		session.Rollback()
		return affected, err
	}

	affected = 0
	factor.Activated = 1
	factor.LastVerifyAt = time.Now()
	affected, err = session.Cols("activated", "last_verify_at").Update(&factor, &TwoFactor{Id: factor.Id})

	if err != nil {
		session.Rollback()
		return affected, err
	}

	// add Commit() after all actions
	if err = session.Commit(); err != nil {
		return affected, err
	}

	return affected, nil
}

func UpdatePassword(id int64, oldPwd, newPwd string) (int64, error) {
	user, err := QueryUserById(id)
	if err != nil {
		return 0, err
	}
	if user.Password == "" {
		return 0, errors.New("No such user, user id not correct.")
	}

	encryPwd := GenPwd(oldPwd, user.Salt)
	if encryPwd != user.Password {
		return 0, errors.New("User old password not correct.")
	}

	user.Password = GenPwd(newPwd, user.Salt)
	user.UpdatedAt = time.Now()
	user.LockTime = time.Now().Add(24 * time.Hour) //添加锁定时间
	affted, err := bxgo.OrmEngin.Id(user.Id).Cols("password", "updated_at", "lock_time").Update(&user)
	return affted, err
}

func UpdatePaymentPassword(userId int64, payPassword string) error {
	m, _ := regexp.MatchString("^[0-9]{6}$", payPassword)
	if !m {
		return errors.New("Password should be length of six numbers")
	}

	user, err := GetUserByID(userId)
	if err != nil {
		return err
	}
	//密码加密
	encryPwd := GenPwd(payPassword, user.Salt)

	_, err = bxgo.OrmEngin.Id(userId).Cols("payment_password", "updated_at", "lock_time").Update(&User{
		PaymentPassword: encryPwd,
		UpdatedAt:       time.Now(),
		LockTime:        time.Now().Add(24 * time.Hour),
	})

	return err
}

func VerifyPaymentPassword(userId int64, payPassword string) (bool, error) {
	m, _ := regexp.MatchString("^[0-9]{6}$", payPassword)
	if !m {
		return false, errors.New("Password should be length of six")
	}

	user, err := GetUserByID(userId)
	if err != nil {
		return false, err
	}
	//密码加密
	encryPwd := GenPwd(payPassword, user.Salt)

	userInfo := User{}
	_, err = bxgo.OrmEngin.Where("id=? ", userId).
		Where("payment_password=? ", encryPwd).Get(&userInfo)
	if err != nil {
		return false, err
	}
	if userInfo.Email == "" {
		return false, errors.New("Payment password not correct.")
	}
	return true, nil
}

//查询是否允许资金操作
func VerifyFinancialOperation(userId int64) (bool, error) {
	user := User{}
	_, err := bxgo.OrmEngin.Where("id=? ", userId). //id判断
							Where("activated=? ", True).                 //激活判断
							Where("disabled=? ", False).                 //禁止判断
							Where("lock_time<? ", time.Now()).Get(&user) //锁定时间判断
	if err != nil {
		return false, err
	}
	if user.Email == "" {
		return false, errors.New("User not available for financial operation right now")
	}
	return true, nil
}

func UpdateUserInfo(u User) (int64, error) {
	user := User{
		DisplayName: u.DisplayName,
		UpdatedAt:   time.Now(),
	}
	_, err := bxgo.OrmEngin.Update(&user, &User{Id: u.Id})
	if err != nil {
		return user.Id, err
	}
	return u.Id, nil
}

func QueryUserByToken(token string) (User, error) {
	user := User{}
	if token == "" {
		return user, errors.New("User token cannot be null")
	}
	_, err := bxgo.OrmEngin.Where("access_token=? ", token).Get(&user)
	if err != nil {
		return user, err
	}
	if user.Id == 0 {
		return user, errors.New("Invalid user token")
	}
	return user, nil
}

func UpdateUserToken(id int64, token string) error {
	user, err := QueryUserById(id)
	if err != nil {
		return err
	}

	user.AccessToken = token
	user.UpdatedAt = time.Now()
	_, err = bxgo.OrmEngin.Id(user.Id).Cols("access_token", "updated_at").Update(&user)

	return err
}

func QueryUserById(id int64) (User, error) {
	user := User{}
	_, err := bxgo.OrmEngin.Id(id).Get(&user)

	return user, err
}

func QueryUsersByFilter(filter UserFilter) ([]User, error) {
	users := []User{}
	session := bxgo.OrmEngin.NewSession()
	defer session.Close()

	if filter.Email != "" {
		session.Where("email = ?", filter.Email)
	}
	if filter.PhoneNumber != "" {
		session.Where("phone_number = ?", filter.PhoneNumber)
	}
	if filter.DisplayName != "" {
		session.Where("display_name = ?", filter.DisplayName)
	}
	if filter.RealName != "" {
		session.Where("real_name = ?", filter.RealName)
	}
	if filter.Kind != 0 {
		session.Where("kind = ?", filter.Kind)
	}
	if filter.Activated == True || filter.Activated == False {
		session.Where("activated = ?", filter.Activated)
	}
	if filter.Disabled == True || filter.Disabled == False {
		session.Where("disabled = ?", filter.Disabled)
	}

	if err := session.Find(&users); err != nil {
		return users, err
	}

	return users, nil
}

func QueryInvitCode(code string) (int64, error) {
	user := User{}
	_, err := bxgo.OrmEngin.Where("invite_code = ?", code).Get(&user)
	if err != nil {
		return user.Id, err
	}
	return user.Id, err
}

func QueryInvitedUser(id int64) ([]User, error) {
	users := []User{}

	inviterInfo, err := QueryInvitUsersInfo(id)
	if err != nil {
		return users, err
	}

	if len(inviterInfo) == 0 {
		return users, nil
	} else {
		ids := []int64{}
		for i, info := range inviterInfo {
			ids[i] = info.InviterId
		}

		err := bxgo.OrmEngin.Where("activated=? ", True).
			In("id", ids).Find(&users)
		if err != nil {
			return users, err
		}

		return users, nil
	}
}

func BindUserGoogleFactor(id int64) (string, string, error) {
	user, err := QueryUserById(id)
	if err != nil {
		return "", "", err
	}
	if user.GoogleIsBind == True {
		factor, err := GetTwoFactor(id, GoogleAuthType)
		if err != nil || factor.Id == 0 {
			return user.GoogleQRImage, "", err
		}

		return user.GoogleQRImage, factor.OtpSecret, nil
	}

	//绑定操作
	factor := TwoFactor{
		UserId:     id,
		VerifyType: GoogleAuthType,
	}
	//创建Factor
	err = factor.GenOTPSecret(false)
	if err != nil {
		return "", "", err
	}
	//生成Image bytes
	byteInfo, err := factor.BarcodeImage(nil)
	if err != nil {
		return "", factor.OtpSecret, err
	}

	//生成图片
	cfg := config.Parse("../config/app.yaml") //获取配置文件信息
	imageName, err := utils.GenerateImage(cfg.ImageSource.Path, byteInfo)
	if err != nil {
		return "", "factor.OtpSecret", errors.New("Generate image got error: " + err.Error())
	}
	user.GoogleQRImage = imageName

	//保存结果
	user.GoogleIsBind = True
	user.UpdatedAt = time.Now()

	_, err = bxgo.OrmEngin.Id(id).
		Cols("google_qrimage", "google_is_bind", "updated_at").
		Update(&user)

	return user.GoogleQRImage, factor.OtpSecret, err
}

func DeleteBindGoogleFactor(id int64) error {
	_, err := bxgo.OrmEngin.Id(id).
		Cols("google_qrimage", "google_is_bind", "updated_at").
		Update(&User{
			GoogleQRImage: "",
			GoogleIsBind:  False,
			UpdatedAt:     time.Now(),
		})
	return err
}
