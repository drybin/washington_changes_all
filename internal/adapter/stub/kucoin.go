package stub

type KucoinStub struct {
}

func NewKucoinStub() *KucoinStub {
	return &KucoinStub{}
}

func (*KucoinStub) GetBalance() (float64, error) {
	return 12.232, nil
}
