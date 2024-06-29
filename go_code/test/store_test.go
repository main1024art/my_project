package aaa
import (
	"testing"
)
func TestStore(t *testing.T) {
	aa := monster{
		Name : "qqqq",
		Age : 12,
		Skill : "FHAU",
	}
	res := aa.Store()
	if !res {
		t.Fatalf("失败应该为%v", true)
	}
	t.Logf("成功")
	
}
func TestRestore(t *testing.T) {
	var ss monster
	res := ss.Restore()
	if !res {
		t.Fatalf("失败应该为%v", true)
	}
	t.Logf("成功")

}