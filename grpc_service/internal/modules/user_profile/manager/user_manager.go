package manager

import (
	"time"

	"github.com/cathatino/topic-tree/grpc_service/internal/modules/user_profile/manager/dbmodel"
	"github.com/go-xorm/xorm"
)

// UserManager provides the user related CRUD functionality
type UserManager struct {
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
}

// CreateUser creates user based on the inputs
// TODO: write test cases
func (um UserManager) CreateUser(
	authMethods string,
	profile string,
	metaData string,
) (
	*dbmodel.UserModel,
	error,
) {
	currentTimeStampInSecs := uint32(time.Now().Unix())
	userModel := dbmodel.UserModel{
		AuthMethods: authMethods,
		Profile:     profile,
		MetaData:    metaData,
		Ctime:       currentTimeStampInSecs,
		Mtime:       currentTimeStampInSecs,
	}
	affected, err := um.masterEngine.Table(userModel.GetTableName()).Insert(&userModel)
	if err != nil {
		return nil, err
	}
	userModel.UserId = affected
	return &userModel, err
}

// NewUserManager returns a new UserManager to manipulate the user data
func NewUserManager(
	masterEngine *xorm.Engine,
	slaveEngine *xorm.Engine,
) UserManager {
	return UserManager{
		masterEngine: masterEngine,
		slaveEngine:  slaveEngine,
	}
}
