package cli

import (
	"reflect"
	"testing"
)

func s(s string) *string {
	return &s
}

func Test_loadFlagFromArgs(t *testing.T) {
	type args struct {
		name  string
		alias string
		args  []string
	}
	tests := []struct {
		name      string
		args      args
		wantFound bool
		wantValue *string
		wantErr   bool
	}{
		{
			name: "-v",
			args: args{
				alias: "v",
				args:  []string{"foo", "--key", "bar", "-v"},
			},
			wantFound: true,
			wantValue: nil,
			wantErr:   false,
		},
		{
			name: "--key",
			args: args{
				name: "key",
				args: []string{"foo", "--key", "bar"},
			},
			wantFound: true,
			wantValue: nil,
			wantErr:   false,
		},
		{
			name: "--key=val",
			args: args{
				name: "key",
				args: []string{"foo", "--key=value", "bar"},
			},
			wantFound: true,
			wantValue: s("value"),
			wantErr:   false,
		},
		{
			name: "--key='val'",
			args: args{
				name: "key",
				args: []string{"foo", "--key='value'", "bar"},
			},
			wantFound: true,
			wantValue: s("value"),
			wantErr:   false,
		},
		{
			name: `--key="val"`,
			args: args{
				name: "key",
				args: []string{"foo", `--key="value"`, "bar"},
			},
			wantFound: true,
			wantValue: s("value"),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotFound, gotValue, err := LoadFlagFromArgs(tt.args.name, tt.args.alias, tt.args.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadFlagFromArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFound != tt.wantFound {
				t.Errorf("LoadFlagFromArgs() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(gotValue, tt.wantValue) {
				t.Errorf("LoadFlagFromArgs() gotValue = %v, want %v", gotValue, tt.wantValue)
			}
		})
	}
}
