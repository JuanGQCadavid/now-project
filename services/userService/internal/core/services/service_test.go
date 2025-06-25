package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimpleTestCheck(t *testing.T) {
	var (
		svc   = Service{}
		tests = []struct {
			title, body string
			want        bool
		}{
			{
				"A valid example",
				"Juan Gonzalo",
				true,
			},
			{
				"Not space at the begining",
				" Juan Gonzalo",
				false,
			},
			{
				"Not space at the end",
				"Juan Gonzalo ",
				false,
			},
			{
				"Not numbers",
				"Juan Gonzalo 21 ",
				false,
			},
			{
				"Not special characters",
				"Juan 🐊 Gonzalo ",
				false,
			},
			{
				"Not special characters 2",
				"Juan ? Gonzalo ",
				false,
			},
			{
				"Not special characters 3",
				"Juan @ Gonzalo ",
				false,
			},
			{
				"Not special characters 3",
				"Juan ǎ Gonzalo ",
				false,
			},
			// {
			// 	"Should allow tildes",
			// 	"Juán  Gonzalo ",
			// 	true,
			// },
		}
	)

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			assert.Equal(t, svc.isSimpleString(tt.body), tt.want)
		})
	}

}
