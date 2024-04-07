package chapter1

import (
	"testing"
)

func TestAddition(t *testing.T) {
	// 正常系のテストパターン
	success := map[string]struct {
		numA int
		numB int
		want int
	}{
		"case1": {numA: 0, numB: 10, want: 10},
		"case2": {numA: 100, numB: 200, want: 300},
	}
	// エラー系のテストパターン
	fail := map[string]struct {
		numA       int
		numB       int
		wantErrStr string
	}{
		"case1": {numA: -1, numB: 10, wantErrStr: "numAは0以上の数値を指定してください。"},
		"case2": {numA: 101, numB: 10, wantErrStr: "numAは100以下の数値を指定してください。"},
		"case3": {numA: 0, numB: 0, wantErrStr: "numBは10以上の数値を指定してください。"},
		"case4": {numA: 0, numB: 201, wantErrStr: "numBは200以下の数値を指定してください。"},
	}

	for tt, tc := range success {
		t.Run(tt, func(t *testing.T) {
			got, err := addition(tc.numA, tc.numB)
			if err != nil {
				t.Errorf("err is not nil: %s", err)
			}
			if tc.want != got {
				t.Errorf("unexpected return. want:%d actual:%d", tc.want, got)
			}
		})
	}
	for tt, tc := range fail {
		t.Run(tt, func(t *testing.T) {
			got, err := addition(tc.numA, tc.numB)
			if got != 0 {
				t.Errorf("unexpected return. want:0 actual:%d", got)
			}
			if tc.wantErrStr != err.Error() {
				t.Errorf("unexpected err. want:%s actual:%s", tc.wantErrStr, err)
			}
		})
	}
}
