package handlers

import (
	"errors"
	"fmt"
	"time"

	bb "grpc-gw/buffb"

	uuid "github.com/satori/go.uuid"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	DB       *mgo.Database
	Accounts *mgo.Collection
	Entries  *mgo.Collection
)

func init() {
	s, err := mgo.Dial("mongodb://root:rooty@localhost:27017")
	if err != nil {
		panic(err)
	}

	if err = s.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to the db...")

	DB = s.DB("quick-accounts")
	Accounts = DB.C("accounts")
	Entries = DB.C("entries")
}

type quickAccountController struct{}

func CreateQAController() *quickAccountController {
	return &quickAccountController{}
}

type entry struct {
	UUID   string `json:"uuid" bson:"_id"`
	UserID string `json:"userid" bson:"userid"`
	Time   int64  `json:"time" bson:"time"`
}

type account struct {
	UUID     string   `json:"uuid" bson:"_id"`
	Username string   `json:"username" bson:"username"`
	Pswdhash string   `json:"pswdhash" bson:"pswdhash"`
	Email    string   `json:"email" bson:"email"`
	Recovery recovery `json:"recovery" bson:"recovery"`
}

type recovery struct {
	Recovery string `json:"recovery" bson:"recovery"`
}

func (quickAccountController) createAccount(qa *bb.QuickAccount) (account, error) {
	id := uuid.NewV4()
	account := account{
		UUID:     id.String(),
		Username: qa.Username,
		Pswdhash: qa.Pswdhash,
		Email:    qa.Email,
	}

	err := Accounts.Insert(account)
	if err != nil {
		return account, errors.New("500. Internal Server Error." + err.Error())
	}
	return account, nil
}

func (quickAccountController) getAccount(uuid string) (account, error) {
	account := account{}
	if uuid == "" {
		errMsg := "Empty accountID provided for query"
		fmt.Println(errMsg)
		return account, errors.New(errMsg)
	}

	err := Accounts.Find(bson.M{"_id": uuid}).One(&account)
	if err != nil {
		return account, err
	}
	return account, nil
}

func (quickAccountController) getAllAccounts() ([]account, error) {
	accounts := []account{}
	err := Accounts.Find(bson.M{}).All(&accounts)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return accounts, nil
}

func (quickAccountController) createEntry(qe *bb.QuickEntry) (entry, error) {
	id := uuid.NewV4()
	now := time.Now().Unix()
	entry := entry{
		UUID:   id.String(),
		UserID: qe.UserID,
		Time:   now,
	}

	err := Entries.Insert(entry)
	if err != nil {
		fmt.Println(err)
		return entry, errors.New("500. Internal Server Error." + err.Error())
	}
	return entry, nil
}

func (quickAccountController) getAllEntries() ([]entry, error) {
	entries := []entry{}
	err := Entries.Find(bson.M{}).All(&entries)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return entries, nil
}
