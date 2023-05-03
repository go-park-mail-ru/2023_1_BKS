package delivery

import (
	app "auth/usecase"
	context "context"
	"fmt"
)

type GrpcServer struct {
	UnimplementedAuthServer

	command app.Commands
	query   app.Queries
}

func (g *GrpcServer) GenerateAccessToken(ctx context.Context, in *Id) (*UuidAuth, error) {
	claims := make(map[string]string)
	claims["id"] = in.GetId()

	resultByte, err := g.command.CreateToken.CreateJWSWithClaims(claims, "appUniqFront", "auth")
	if err != nil {
		return nil, err
	}
	result := UuidAuth{Value: string(resultByte)}
	fmt.Println(result.GetValue())
	return &result, nil
}
