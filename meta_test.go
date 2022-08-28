package application_test

import (
	"reflect"
	"testing"

	"kkn.fi/application"
)

func TestNewMeta(t *testing.T) {
	type args struct {
		buildTime string
		version   string
		revision  string
	}
	tests := []struct {
		name    string
		args    args
		want    *application.Meta
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				buildTime: "2022-08-28T16:13:19+0000",
				version:   "v0.0.8",
				revision:  "87833c220402a5838a73dc347433700b68c6333e",
			},
			want: &application.Meta{
				BuildTime: "2022-08-28T16:13:19+0000",
				Version:   "v0.0.8",
				Revision:  "87833c220402a5838a73dc347433700b68c6333e",
			},
			wantErr: false,
		},
		{
			name: "build time is required",
			args: args{
				buildTime: "",
				version:   "v0.0.8",
				revision:  "87833c220402a5838a73dc347433700b68c6333e",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "version is required",
			args: args{
				buildTime: "2022-08-28T16:13:19+0000",
				version:   "",
				revision:  "87833c220402a5838a73dc347433700b68c6333e",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "revision is required",
			args: args{
				buildTime: "2022-08-28T16:13:19+0000",
				version:   "v0.0.8",
				revision:  "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := application.NewMeta(tt.args.buildTime, tt.args.version, tt.args.revision)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMeta() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMeta() = %v, want %v", got, tt.want)
			}
		})
	}
}
