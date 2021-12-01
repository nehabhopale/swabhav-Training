package pizzamenu

type TomatoTopping struct{
	Pizza Pizza
}

func (t *TomatoTopping) GetPrize() uint32 {
	pizzaPrice := t.Pizza.GetPrize()
	return pizzaPrice + 20
}