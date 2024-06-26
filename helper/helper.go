package helper

import (
	"crypto/md5"
	"crypto/tls"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/jordan-wright/email"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/smtp"
)

type UserClaims struct {
	Identity primitive.ObjectID `json:"identity"`
	Name     string             `json:"name"`
	Level    int                `json:"level"`
	jwt.StandardClaims
}

func GetMd5(s string) string {
	tmp := md5.Sum([]byte(s))
	nextPwd := fmt.Sprintf("%x", tmp)
	res := md5.Sum([]byte(nextPwd + "%132yse@htwxc.com+changhao2333?"))
	newPwd := fmt.Sprintf("%x", res)
	return newPwd
}

var myKey = []byte("cuipiya")

func GenerateToken(identity, name string, level int) (string, error) {
	objectID, err := primitive.ObjectIDFromHex(identity)
	if err != nil {
		return "", err
	}
	UserClaim := &UserClaims{
		Identity:       objectID,
		Name:           name,
		Level:          level,
		StandardClaims: jwt.StandardClaims{},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim)
	tokenString, err := token.SignedString(myKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func AnalyseToken(tokenString string) (*UserClaims, error) {
	userClaim := new(UserClaims)
	claims, err := jwt.ParseWithClaims(tokenString, userClaim, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})
	if err != nil {
		return nil, err
	}
	if !claims.Valid {
		return nil, fmt.Errorf("analyse Token Error:%v", err)
	}
	return userClaim, nil
}

// SendCode
// 发送验证码
func SendCode(toUserEmail, code string) error {
	e := email.NewEmail()
	e.From = "Get <nugei@foxmail.com>"
	e.To = []string{toUserEmail}
	e.Subject = "验证码已发送，请查收"
	e.HTML = []byte("您的验证码：<b>" + code + "</b>")
	return e.SendWithTLS("smtp.qq.com:465",
		smtp.PlainAuth("", "nugei@foxmail.com", "changhao", "smtp.qq.com"),
		&tls.Config{InsecureSkipVerify: true, ServerName: "smtp.qq.com"})
}
