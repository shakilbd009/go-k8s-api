package data

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/shakilbd009/go-k8s-api/src/client/cassandra"
	"github.com/shakilbd009/go-k8s-api/src/domain/k8s"
	"github.com/shakilbd009/go-k8s-api/src/utils/token"
	"github.com/shakilbd009/go-k8s-api/src/utils/utility"
	"github.com/shakilbd009/go-utils-lib/rest_errors"
)

const (
	queryGetAccessToken    = "SELECT access_token,namespace FROM k8s_access_tokens WHERE access_token=?;"
	queryCreateAccessToken = "INSERT INTO k8s_access_tokens(access_token,namespace,user_id,email) VALUES (?,?,?,?);"
)

var (
	Database dataInterface = &dataService{}
)

type dataService struct{}

type dataInterface interface {
	CreateUser(*k8s.K8sUser) (*k8s.K8sUser, rest_errors.RestErr)
	GetAccessToken(string) (*k8s.K8sUser, rest_errors.RestErr)
}

func (*dataService) CreateUser(u *k8s.K8sUser) (*k8s.K8sUser, rest_errors.RestErr) {

	if err := u.Validate(); err != nil {
		return nil, err
	}
	access_token := token.GetMD5HashToken(u.UserName, u.Email)
	if err := cassandra.GetSession().Query(queryCreateAccessToken,
		access_token,
		u.Namespace,
		u.UserName,
		u.Email).Exec(); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to create an access_token", utility.ErrDatabase)
	}
	u.Token = access_token
	return u, nil
}

func (*dataService) GetAccessToken(token string) (*k8s.K8sUser, rest_errors.RestErr) {

	var user k8s.K8sUser
	if err := cassandra.GetSession().Query(queryGetAccessToken, token).Scan(
		&user.Token,
		&user.Namespace,
	); err != nil {
		fmt.Println(err.Error())
		if err == gocql.ErrNotFound {
			return nil, rest_errors.NewNotFoundError(fmt.Sprintf("access_token: %s is not valid", token))
		}
		return nil, rest_errors.NewInternalServerError("error when trying to get access_token", utility.ErrDatabase)
	}
	return &user, nil
}
