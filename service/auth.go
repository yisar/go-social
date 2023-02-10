package service

import (
	"errors"
	"fmt"

	"github.com/cliclitv/htwxc/helper"
	"github.com/cliclitv/htwxc/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Auth(aid string, token string) error {
	userClaims, err := helper.AnalyseToken(token)
	if err != nil {
		return err
	}
	// 查找当前用户
	objectID, _ := primitive.ObjectIDFromHex(aid)
	user, err := model.GetAuthorByIdentity(objectID)

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
