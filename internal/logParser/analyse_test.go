package logParser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findTopN(t *testing.T) {
	type args struct {
		input []string
		n     int
	}
	tests := []struct {
		name string
		args args
		want []uniqueWithCount
	}{
		{
			name: "Success - top 3 values that are not equal placement",
			args: args{
				input: []string{
					"/docs",
					"/docs",
					"/docs",
					"/docs",
					"/temp-redirect",
					"/temp-redirect",
					"/temp-redirect",
					"/newsletter",
					"/newsletter",
					"/nada",
				},
				n: 3,
			},
			want: []uniqueWithCount{
				{
					value: "/docs",
					count: 4,
				}, {
					value: "/temp-redirect",
					count: 3,
				}, {
					value: "/newsletter",
					count: 2,
				}},
		},
		{
			name: "Success - top 3 values that are equal placement",
			args: args{
				input: []string{
					"/docs",
					"/temp-redirect",
					"/newsletter",
					"/nada",
				},
				n: 3,
			},
			want: []uniqueWithCount{
				{
					value: "/docs",
					count: 1,
				}, {
					value: "/temp-redirect",
					count: 1,
				}, {
					value: "/newsletter",
					count: 1,
				}},
		},
		{
			name: "Success - no top value provided",
			args: args{
				input: []string{
					"/docs",
					"/temp-redirect",
					"/newsletter",
					"/nada",
				},
				n: 0,
			},
			want: []uniqueWithCount{
				{
					value: "/docs",
					count: 1,
				}, {
					value: "/temp-redirect",
					count: 1,
				}, {
					value: "/newsletter",
					count: 1,
				},
				{
					value: "/nada",
					count: 1,
				}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findTopN(tt.args.input, tt.args.n), "findTopN(%v, %v)", tt.args.input, tt.args.n)
		})
	}
}

func Test_findUnique(t *testing.T) {
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "success - get uniques",
			args: args{
				input: []string{
					"/docs",
					"/docs",
					"/translations",
				},
			},
			want: []string{
				"/docs",
				"/translations"},
		},
		{
			name: "success - empty list",
			args: args{},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.ElementsMatch(t, tt.want, findUnique(tt.args.input), "findUnique(%v)", tt.args.input)
		})
	}
}
