//go:build !integration

package application_test

import (
	"reflect"
	"testing"

	"kkn.fi/application"
)

func TestNewBuild(t *testing.T) {
	type args struct {
		buildTime string
		version   string
		revision  string
		goVersion string
	}
	tests := []struct {
		name    string
		args    args
		want    *application.Build
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				buildTime: "2022-08-28T16:13:19+0000",
				version:   "v0.0.8",
				revision:  "87833c220402a5838a73dc347433700b68c6333e",
				goVersion: "1.19",
			},
			want: &application.Build{
				Time:      "2022-08-28T16:13:19+0000",
				Version:   "v0.0.8",
				Revision:  "87833c220402a5838a73dc347433700b68c6333e",
				GoVersion: "1.19",
			},
			wantErr: false,
		},
		{
			name: "build time is required",
			args: args{
				buildTime: "",
				version:   "v0.0.8",
				revision:  "87833c220402a5838a73dc347433700b68c6333e",
				goVersion: "1.19",
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
				goVersion: "1.19",
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
				goVersion: "1.19",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "go version is required",
			args: args{
				buildTime: "2022-08-28T16:13:19+0000",
				version:   "v0.0.8",
				revision:  "87833c220402a5838a73dc347433700b68c6333e",
				goVersion: "",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := application.NewBuild(tt.args.buildTime, tt.args.version, tt.args.revision, tt.args.goVersion)
			if (err != nil) != tt.wantErr {
				t.Errorf("application.NewBuild() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("application.NewBuild() = %v, want %v", got, tt.want)
			}
		})
	}
}
