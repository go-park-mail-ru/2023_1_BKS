package delivery

import (
	context "context"
	app "user/usecase"
)

type GrpcServer struct {
	UnimplementedUserServer

	command app.Commands
	query   app.Queries
}

func (g *GrpcServer) CheckAccount(ctx context.Context, in *UserCheck) (*Uuid, error) {
	result := Uuid{Value: g.query.CheckUser.Handle(ctx, in.Login, in.Password)}
	return &result, nil
}
