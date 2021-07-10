package mysql

import (
	"testing"
	"time"
)

func TestConfig_GetDsn(t *testing.T) {
	type fields struct {
		host            string
		port            string
		user            string
		password        string
		dbName          string
		maxOpenConns    int
		maxIdleConns    int
		connMaxLifetime time.Duration
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test_for_correct_config_name",
			fields: fields{
				host:     "host",
				port:     "port",
				user:     "user",
				password: "password",
				dbName:   "db_name",
			},
			want: "user:password@tcp(host:port)/db_name?charset=utf8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				host:            tt.fields.host,
				port:            tt.fields.port,
				user:            tt.fields.user,
				password:        tt.fields.password,
				dbName:          tt.fields.dbName,
				maxOpenConns:    tt.fields.maxOpenConns,
				maxIdleConns:    tt.fields.maxIdleConns,
				connMaxLifetime: tt.fields.connMaxLifetime,
			}
			if got := c.GetDsn(); got != tt.want {
				t.Errorf("Config.GetDsn() = %v, want %v", got, tt.want)
			}
		})
	}
}
