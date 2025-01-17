package tUtil

import (
	"context"
	"time"

	"github.com/ZiRunHua/LeapLedger/global/constant"
	"github.com/ZiRunHua/LeapLedger/global/cus"
	"github.com/ZiRunHua/LeapLedger/global/db"
	accountModel "github.com/ZiRunHua/LeapLedger/model/account"
	categoryModel "github.com/ZiRunHua/LeapLedger/model/category"
	transactionModel "github.com/ZiRunHua/LeapLedger/model/transaction"
	userModel "github.com/ZiRunHua/LeapLedger/model/user"
	"github.com/ZiRunHua/LeapLedger/util/rand"
)

type Build struct {
}

var build = &Build{}

func (n *Build) User() (user userModel.User, err error) {
	user, err = userService.CreateTourist(context.TODO())
	if err != nil {
		return
	}
	tour, err := userModel.NewDao().SelectTour(user.ID)
	if err != nil {
		return
	}
	err = db.Transaction(
		context.TODO(), func(ctx *cus.TxContext) error {
			return tour.Use(ctx.GetDb())
		},
	)
	return user, err
}

func (n *Build) Account(user userModel.User, t accountModel.Type) (
	account accountModel.Account, aUser accountModel.User,
	err error,
) {
	accountTmpl := templateService.NewAccountTmpl()
	err = accountTmpl.ReadFromJson(constant.ExampleAccountJsonPath)
	if err != nil {
		return
	}
	accountTmpl.Type = t
	return templateService.CreateAccountByTemplate(accountTmpl, user, context.TODO())
}

func (n *Build) TransInfo(user userModel.User, category categoryModel.Category) transactionModel.Info {
	return transactionModel.Info{
		UserId:        user.ID,
		AccountId:     category.AccountId,
		CategoryId:    category.ID,
		IncomeExpense: category.IncomeExpense,
		Amount:        rand.Int(1000),
		Remark:        "test",
		TradeTime:     time.Now(),
	}
}
