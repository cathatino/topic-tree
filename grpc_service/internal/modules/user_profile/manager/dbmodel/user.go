package dbmodel

// Design based on the suggestion of:
// https://cloud.google.com/blog/products/identity-security/account-authentication-and-password-management-best-practices
// https://stackoverflow.com/questions/11938044/what-are-the-best-practices-for-using-a-guid-as-a-primary-key-specifically-rega/11938495
// https://www.sqlskills.com/blogs/kimberly/guids-as-primary-keys-andor-the-clustering-key/

// UserModel defines the user account content,
// while the login method will be handle in other tables to facilidate the different login mechanism
// Note: let's start with a non-sharded version of the user table, assuming the user is with small size
type UserModel struct {
	UserId      int64  `xorm:"'user_id' autoincr BIGINT(20)"`
	Status      uint32 `xorm:"'status' not null INT(11)"`
	AuthMethods string `xorm:"auth_methods"`
	Profile     string `xorm:"profile"`
	MetaData    string `xorm:"meta_data"`
	Ctime       uint32 `xorm:"'ctime' not null INT(11)"`
	Mtime       uint32 `xorm:"'mtime' not null INT(11)"`
}

// GetTableName returns the table name of UserModel
func (um UserModel) GetTableName() string {
	return UserTabName
}
