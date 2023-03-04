package firestore

import (
	"context"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/pincheng0101/go-linebot-server-template/config"
	"google.golang.org/api/option"
)

type User struct {
	UserID   string `firestore:"user_id"`
	Name     string `firestore:"name"`
	Phone    string `firestore:"phone"`
	Region   string `firestore:"region"`
	Birthday string `firestore:"birthday"`
	CarType  string `firestore:"cartype"`
	Points   int    `firestore:"points"`
}

type Firestore struct {
	client *firestore.Client
	ctx    context.Context
}

func NewFirestore() (*Firestore, error) {
	cfg, _ := config.LoadConfig()

	ctx := context.Background()
	opt := option.WithCredentialsFile(cfg.Firebase.ServiceAccountFile)
	app, err := firebase.NewApp(ctx, nil, opt)
	if err != nil {
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, err
	}
	return &Firestore{client, ctx}, nil
}

func (f *Firestore) AddUser(user User) error {
	_, err := f.client.Collection("users").Doc(user.UserID).Set(f.ctx, user)
	return err
}

func (f *Firestore) GetUser(userId string) (*User, error) {
	doc, err := f.client.Collection("users").Doc(userId).Get(f.ctx)
	if err != nil {
		return nil, err
	}

	var user User
	if err := doc.DataTo(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (f *Firestore) UpdateUser(userID string, userUpdates map[string]interface{}) error {
	var updates []firestore.Update
	for key, val := range userUpdates {
		updates = append(updates, firestore.Update{
			Path:  key,
			Value: val,
		})
	}
	_, err := f.client.Collection("users").Doc(userID).Update(f.ctx, updates)
	return err
}

func (f *Firestore) AddPoints(userID string, pointsToAdd int) error {
	_, err := f.client.Collection("users").Doc(userID).Update(f.ctx, []firestore.Update{
		{
			Path:  "points",
			Value: firestore.Increment(pointsToAdd),
		},
	})
	return err
}
