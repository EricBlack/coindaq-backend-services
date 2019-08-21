package model

import (
	"time"

	"bx.com/user-service/bxgo"
	"github.com/pquerna/otp/totp"
	qr "github.com/qpliu/qrencode-go/qrencode"
	"hash"
	"crypto/sha1"
	"net/url"
	"fmt"
	"bytes"
	"strings"
	"strconv"
	"image/png"
)

const (
	CoinDaqIssuer = "coindaq.com"
)

func (tf *TwoFactor) VerifyOTP(otp string) bool {
	valid := totp.Validate(otp, tf.OtpSecret)
	if valid {
		if _, err := bxgo.OrmEngin.Id(tf.Id).
			Cols("last_verify_at", "activated").
			Update(&TwoFactor{LastVerifyAt: time.Now(), Activated: True}); err != nil {
			return false
		}
	}

	return valid
}

func (tf *TwoFactor) GenOTPSecret(refresh bool) error {
	user, err := GetUserByID(tf.UserId)
	if err != nil {
		return err
	}
	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      CoinDaqIssuer,
		AccountName: user.Email,
	})
	if err != nil {
		return err
	}
	tf.OtpSecret = key.Secret()
	if refresh {
		_, err = bxgo.OrmEngin.Where("user_id=? ", tf.UserId).
			Where("verify_type=? ", GoogleAuthType).
			Cols("otp_secret", "activated", "last_verify_at", "refreshed_at").
			Update(&TwoFactor{OtpSecret: tf.OtpSecret, Activated: False, LastVerifyAt: time.Now(), RefreshedAt: time.Now()})
	}else{
		_, err = bxgo.OrmEngin.Where("user_id=? ", tf.UserId).
			Where("verify_type=? ", GoogleAuthType).
			Cols("otp_secret", "activated", "last_verify_at").
			Update(&TwoFactor{OtpSecret: tf.OtpSecret, Activated: False, LastVerifyAt: time.Now()})
	}

	return err
}

func (tf *TwoFactor) Refresh() error {
	if tf.Activated == 1 {
		return nil
	}
	return tf.GenOTPSecret(true)
}

func (tf *TwoFactor) BarcodeImage(opt *Options) ([]byte, error) {
	if opt == nil {
		opt = DefaultOptions
	}

	user, err := GetUserByID(tf.UserId)
	if err != nil {
		return nil, err
	}

	u := &url.URL{
		Scheme: 	"otpauth",
		Host:   	"totp",
		Path:   	fmt.Sprintf("/%s", user.Email),
	}
	params := url.Values{
		"secret": {strings.TrimRight(tf.OtpSecret, "=")},
		"issuer": {"Coindaq"},
		"digits": {strconv.Itoa(int(opt.Digits))},
		"period": {strconv.Itoa(int(opt.TimeStep / time.Second))},
	}

	u.RawQuery = params.Encode()

	c, err := qr.Encode(u.String(), qr.ECLevelM)

	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer

	err = png.Encode(&buf, c.Image(8))

	return buf.Bytes(), err
}

// Options contains the different configurable values for a given TOTP
// invocation.
type Options struct {
	Time     func() time.Time
	Tries    []int64
	TimeStep time.Duration
	Digits   uint8
	Hash     func() hash.Hash
}

// NewOptions constructs a pre-configured Options. The returned Options' uses
// time.Now to get the current time, has a window size of 30 seconds, and
// tries the currently active window, and the previous one. It expects 6 digits,
// and uses sha1 for its hash algorithm. These settings were chosen to be
// compatible with Google Authenticator.
func NewOptions() *Options {
	return &Options{
		Time:     time.Now,
		Tries:    []int64{0, -1},
		TimeStep: 30 * time.Second,
		Digits:   6,
		Hash:     sha1.New,
	}
}

var DefaultOptions = NewOptions()

var digit_power = []int64{
	1,          // 0
	10,         // 1
	100,        // 2
	1000,       // 3
	10000,      // 4
	100000,     // 5
	1000000,    // 6
	10000000,   // 7
	100000000,  // 8
	1000000000, // 9
}
