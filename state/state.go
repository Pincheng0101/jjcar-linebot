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
	Points   int
}

type UserState struct {
	UserID           string
	UserInfo         UserInfo
	BeforeAskCommand command.AskCommand
	BeforeReplay     string
	IsRegistered     bool
}

func (us *UserState) UpdateBeforeAsk(userID string, askCommand command.AskCommand) {
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

func (us *UserState) AddPoints(points int) {
	us.UserInfo.Points += points
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
	Data map[string]*UserState
}

func NewUserStates() UserStates {
	return UserStates{
		Data: make(map[string]*UserState),
	}
}

func (us UserStates) GetUserState(userID string) *UserState {
	if userState, ok := us.Data[userID]; !ok {
		return nil
	} else {
		return userState
	}
}

func (us UserStates) CreateUserStateByUser(userID string, userInfo UserInfo) *UserState {
	userState := UserState{
		UserID:       userID,
		UserInfo:     userInfo,
		IsRegistered: true,
	}
	us.Data[userID] = &userState
	return us.Data[userID]
}

func (us UserStates) CreateUserStateIfNotExist(userID string) *UserState {
	if _, ok := us.Data[userID]; !ok {
		us.Data[userID] = &UserState{
			UserID: userID,
		}
	}
	return us.Data[userID]
}

func (us UserStates) UserExists(userID string) bool {
	if _, ok := us.Data[userID]; !ok {
		return false
	}
	return true
}
