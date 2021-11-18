package bankcustomer
import "testing"
var accForTestsender *customer=New(1,"neha",334,5000,"savings",12)
var accForTestreciver *customer=New(2,"pooja",384,2000,"savings",12)
func TestMoneyTranfer(t *testing.T){
	expected:=nil

	actual:=accForTestsender.MoneyTransfer(accForTestreciver, 334, 384, 12, 100, "savings")
	if expected!=actual{
		t.Errorf("expected %v,actual %v",expected,actual)
	}
}
