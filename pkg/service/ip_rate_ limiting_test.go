package service

import (
	"context"
	pkg "dcard/pkg"
	"dcard/pkg/model"
	"testing"
	"time"
)

func Test_service_RateLimitIsAllow(t *testing.T) {
	type fields struct {
		repo pkg.IRepository
	}
	type args struct {
		ctx context.Context
		in  model.IPRateLimit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				repo: suite.repo,
			},
			args: args{
				ctx: context.TODO(),
				in: model.IPRateLimit{
					IP:        "192.168.1.1",
					RateCount: 10,
					RateSec:   10 * time.Second,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			if err := s.RateLimitIsAllow(tt.args.ctx, tt.args.in); (err != nil) != tt.wantErr {
				t.Errorf("service.RateLimitIsAllow() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_service_GetRateLimitCount(t *testing.T) {
	type fields struct {
		repo pkg.IRepository
	}
	type args struct {
		ctx context.Context
		in  model.IPRateLimit
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "1",
			fields: fields{
				repo: suite.repo,
			},
			args: args{
				ctx: context.TODO(),
				in: model.IPRateLimit{
					IP: "192.168.1.1.",
				},
			},
			want:    0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				repo: tt.fields.repo,
			}
			got, err := s.GetRateLimitCount(tt.args.ctx, tt.args.in)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.GetRateLimitCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("service.GetRateLimitCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
