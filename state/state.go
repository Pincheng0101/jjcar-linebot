package state

import (
	"github.com/pincheng0101/go-linebot-server-template/command"
)

type UserID string

type UserInfo struct {
	Name     string
	Phone    string
	Region   string
	Birthday string
	CarType  string
	Point    uint
}

type UserState struct {
	UserID           UserID
	UserInfo         UserInfo
	BeforeAskCommand command.AskCommand
	BeforeReplay     string
	IsRegistered     bool
}

func (us *UserState) UpdateBeforeAsk(userID UserID, askCommand command.AskCommand) {
	us.BeforeAskCommand = askCommand
}

func (us *UserState) UpdateBeforeReplay(replay string) {
	us.BeforeReplay = replay
}

func (us *UserState) UpdateName(name string) {
	us.UserInfo.Name = name
}

func (us *UserState) UpdatePhone(phone string) {
	us.UserInfo.Phone = phone
}

func (us *UserState) UpdateRegion(region string) {
	us.UserInfo.Region = region
}

func (us *UserState) UpdateBirthday(birthday string) {
	us.UserInfo.Birthday = birthday
}

func (us *UserState) UpdateCarType(cartype string) {
	us.UserInfo.CarType = cartype
}

func (us *UserState) Registered() {
	us.IsRegistered = true
}

func (us *UserState) ResetUserState() {
	us = &UserState{
		UserID: us.UserID,
	}
}

type UserStates struct {
	Data map[UserID]*UserState
}

func NewUserStates() UserStates {
	return UserStates{
		Data: make(map[UserID]*UserState),
	}
}

func (us UserStates) CreateUserStateIfNotExist(userID UserID) *UserState {
	if _, ok := us.Data[userID]; !ok {
		us.Data[userID] = &UserState{
			UserID: userID,
		}
	}
	return us.Data[userID]
}
