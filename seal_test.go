// Package libseal
//
// @author: xwc1125
package libseal

import (
	"os"
	"testing"
)

func TestSeal(t *testing.T) {
	dc, err := BuildCompanySeal("北京xxx科技有限公司", "110xxxxxxxx55", "Songti.ttc")
	if err != nil {
		t.Fatal(err)
	}
	os.MkdirAll("./img", os.ModePerm)
	err = dc.SavePNG("./img/out-company.png")
	if err != nil {
		t.Fatal(err)
	}
}

func TestPersonSeal(t *testing.T) {
	dc, err := BuildPersonalSeal("小明明", "Songti.ttc")
	if err != nil {
		t.Fatal(err)
	}
	os.MkdirAll("./img", os.ModePerm)
	err = dc.SavePNG("./img/out-personal2.png")
	if err != nil {
		t.Fatal(err)
	}
}
