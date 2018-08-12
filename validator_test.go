package go_vo_validator

import "testing"

type VOExample struct {
	Name string `mandatory:"true"`
	Rating int  `validateMax:"5"`
}

func TestValidateMandatory(t *testing.T) {
	vo := VOExample{}
	errs := Validate(&vo)
	if len(errs) > 0 {
		t.Log(errs[0].Error())
	} else {
		t.Error("validator mandatory unsuccessful")
	}
}


func TestValidateMaxValue(t *testing.T) {
	vo := VOExample{Name:"Example Name", Rating:10}
	errs := Validate(&vo)
	if len(errs) > 0 {
		t.Log(errs[0].Error())
	} else {
		t.Error("validator max value unsuccessful")
	}

}

