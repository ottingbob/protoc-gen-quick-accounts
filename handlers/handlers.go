package handlers

import (
	"context"
	"fmt"
	bb "grpc-gw/buffb"
	pb "grpc-gw/example"

	"github.com/golang/protobuf/ptypes/empty"
)

type RpcServer struct {
	Name string
}

func (srv *RpcServer) Echo(ctx context.Context, msg *pb.StringMessage) (*pb.StringMessage, error) {
	fmt.Println("GRPC Received: ", msg.Value)
	res := &pb.StringMessage{
		Value: "Hello from RPC Server!",
	}
	return res, nil
}

type QuickRPCServer struct {
	Name         string
	QAController *quickAccountController
}

func (srv *QuickRPCServer) GetAccounts(ctx context.Context, e *empty.Empty) (*bb.QuickAccounts, error) {
	accounts, err := srv.QAController.getAllAccounts()
	if err != nil {
		return nil, err
	}

	var accountsArr = make([]*bb.QuickAccount, 0)
	var accountDefaultRecovery = &bb.Recovery{
		Recovery: "default-recovery-phrase",
	}

	for _, s := range accounts {
		fmt.Printf("%v\n", s)
		accountsArr = append(accountsArr,
			&bb.QuickAccount{
				Uuid:     s.UUID,
				Username: s.Username,
				Pswdhash: s.Pswdhash,
				Email:    s.Email,
				Recovery: accountDefaultRecovery,
			})
	}

	var accountsResp = &bb.QuickAccounts{
		Accounts: accountsArr,
	}
	return accountsResp, nil
}

func (srv *QuickRPCServer) GetEntries(ctx context.Context, e *empty.Empty) (*bb.QuickEntries, error) {
	entries, err := srv.QAController.getAllEntries()
	if err != nil {
		return nil, err
	}

	var entriesArr = make([]*bb.QuickEntry, 0)
	for _, s := range entries {
		entriesArr = append(entriesArr,
			&bb.QuickEntry{
				QuickID:   s.UUID,
				UserID:    s.UserID,
				QuickTime: s.Time,
			})
	}

	var entriesResp = &bb.QuickEntries{
		Entries: entriesArr,
	}
	return entriesResp, nil
}

func (srv *QuickRPCServer) CreateAccount(ctx context.Context, qa *bb.QuickAccount) (*bb.QuickAccount, error) {
	fmt.Println("CreateAccount RPC Handler received: ", qa)
	dbres, err := srv.QAController.createAccount(qa)
	if err != nil {
		return nil, err
	}

	res := &bb.QuickAccount{
		Uuid:     dbres.UUID,
		Username: dbres.Username,
		Pswdhash: dbres.Pswdhash,
		Email:    dbres.Email,
	}
	return res, nil
}

func (srv *QuickRPCServer) GetAccountByUUID(ctx context.Context, qa *bb.QuickAccount) (*bb.QuickAccount, error) {
	fmt.Println("GetAccountByUUID RPC Handler received: ", qa)

	account, err := srv.QAController.getAccount(qa.Uuid)
	if err != nil {
		return nil, err
	}

	res := &bb.QuickAccount{
		Uuid:     account.UUID,
		Username: account.Username,
		Pswdhash: account.Pswdhash,
		Email:    account.Email,
	}
	return res, nil
}

func (srv *QuickRPCServer) CreateEntry(ctx context.Context, qe *bb.QuickEntry) (*bb.QuickEntry, error) {
	fmt.Println("CreateEntry RPC Handler received: ", qe)

	entry, err := srv.QAController.createEntry(qe)
	if err != nil {
		return nil, err
	}

	res := &bb.QuickEntry{
		QuickID:   entry.UUID,
		UserID:    entry.UserID,
		QuickTime: entry.Time,
	}
	return res, nil
}
