package chapter2

import (
	"testing"

	"github.com/dip-dev/go-test-tutorial/chapters/chapter2/communication"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestExecWithMock(t *testing.T) {
	success := map[string]struct {
		want string
	}{
		"mock化テスト": {
			want: "Nice to meet you!!",
		},
	}

	for tt, tc := range success {
		t.Run(tt, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mock := communication.NewMockInterfaceCommunication(ctrl)

			mock.EXPECT().Greeting().Return(tc.want)

			chapter2 := New(mock)
			greeting := chapter2.exec()

			assert.Equal(t, tc.want, greeting)
		})
	}
}
