package mysql

import (
	"testing"
)

func TestConfig_GetDsn(t *testing.T) {
	type fields struct {
		Host            string
		Port            string
		User            string
		Password        string
		DbName          string
		MaxOpenConns    int
		MaxIdleConns    int
		ConnMaxLifetime duration
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test_for_correct_config_name",
			fields: fields{
				Host:     "host",
				Port:     "port",
				User:     "user",
				Password: "password",
				DbName:   "db_name",
			},
			want: "user:password@tcp(host:port)/db_name?charset=utf8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				Host:            tt.fields.Host,
				Port:            tt.fields.Port,
				User:            tt.fields.User,
				Password:        tt.fields.Password,
				DbName:          tt.fields.DbName,
				MaxOpenConns:    tt.fields.MaxOpenConns,
				MaxIdleConns:    tt.fields.MaxIdleConns,
				ConnMaxLifetime: tt.fields.ConnMaxLifetime,
			}
			if got := c.GetDsn(); got != tt.want {
				t.Errorf("Config.GetDsn() = %v, want %v", got, tt.want)
			}
		})
	}
}
