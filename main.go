package main

//   **** DECORATOR ****

// type Pizza interface {
// 	Cost() int
// 	Description() string
// }

// type PlainPizza struct{}

// func (p *PlainPizza) Cost() int {
// 	return 10
// }

// func (p *PlainPizza) Description() string {
// 	return "Plain pizza"
// }

// type PizzaDecorator struct {
// 	pizza       Pizza
// 	extraCost   int
// 	extraTopping string
// }

// func NewPizzaDecorator(pizza Pizza, extraCost int, extraTopping string) *PizzaDecorator {
// 	return &PizzaDecorator{
// 		pizza:       pizza,
// 		extraCost:   extraCost,
// 		extraTopping: extraTopping,
// 	}
// }

// func (d *PizzaDecorator) Cost() int {
// 	return d.pizza.Cost() + d.extraCost
// }

// func (d *PizzaDecorator) Description() string {
// 	return d.pizza.Description() + ", " + d.extraTopping
// }

// func main() {
// 	pizza := &PlainPizza{}
// 	fmt.Println("Cost:", pizza.Cost(), "Description:", pizza.Description())

// 	pizzaWithCheese := NewPizzaDecorator(pizza, 2, "Extra Cheese")
// 	fmt.Println("Cost:", pizzaWithCheese.Cost(), "Description:", pizzaWithCheese.Description())

// 	pizzaWithPepperoni := NewPizzaDecorator(pizza, 3, "Pepperoni")
// 	fmt.Println("Cost:", pizzaWithPepperoni.Cost(), "Description:", pizzaWithPepperoni.Description())
// }

// **** FACTORY PATTERN ****
// type Shape interface {
// 	Draw() string
// }

// type Circle struct{}

// func (c *Circle) Draw() string {
// 	return "Circle"
// }

// type Square struct{}

// func (s *Square) Draw() string {
// 	return "Square"
// }

// type ShapeFactory interface {
// 	CreateShape() Shape
// }

// type CircleFactory struct{}

// func (cf *CircleFactory) CreateShape() Shape {
// 	return &Circle{}
// }

// type SquareFactory struct{}

// func (sf *SquareFactory) CreateShape() Shape {
// 	return &Square{}
// }

// func main() {
// 	circleFactory := &CircleFactory{}
// 	circle := circleFactory.CreateShape()
// 	fmt.Println("Shape created:", circle.Draw())

// 	squareFactory := &SquareFactory{}
// 	square := squareFactory.CreateShape()
// 	fmt.Println("Shape created:", square.Draw())
// }
