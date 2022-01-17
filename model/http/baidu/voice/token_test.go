package voice

import "testing"

func Test_GetToken(t *testing.T) {
	result, err := GetToken()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", result)
}
