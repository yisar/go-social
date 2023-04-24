package service

import (
	"errors"
	"fmt"

	"github.com/yisar/footsie/helper"
	"github.com/yisar/footsie/db"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Auth(uid string, token string) error {
	userClaims, err := helper.AnalyseToken(token)
	if err != nil {
		return err
	}
	// 查找当前用户
	objectID, _ := primitive.ObjectIDFromHex(uid)
	user, err := model.GetUserByIdentity(objectID)

	fmt.Println(userClaims.Name)

	if err != nil {
		return err
	}

	if user.Name == userClaims.Name {
		// 本人编辑，ok
		return nil
	}

	if userClaims.Level > user.Level {
		// 编辑者权限 > 作者权限
		return nil
	}

	return errors.New("权限不足")
}
