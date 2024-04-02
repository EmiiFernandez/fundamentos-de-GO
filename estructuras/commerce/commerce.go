package commerce

/*
Estructuras que simularan un carrito de compras
*/

//Estructuras

type Product struct {
	ID    uint8    `json: "id"`
	Name  string   `json: name"`
	Type  Type     `json: "type"`
	Count uint16   `json: "count"`
	Price float64  `json: "price"`
	Tags  []string `json: "tags"`
}

type Type struct {
	ID          uint8  `json: "id"`
	Code        string `json: code"`
	Description string `json: "description"`
}

type Car struct {
	ID       uint      `json: "id"`
	UserID   uint      `json: user_id"`
	Products []Product `json: products"`
}

//Metodos
func (p Product) TotalPrice() float64 {
	return float64 /*Casteo a flotante*/ (p.Count) * p.Price
}

//Agrego * po r que modifico la estructura original
func (c *Car) AddProducts(products ...Product) {
	c.Products = append(c.Products, products...) //products... transforma el slice (p1, p2,p3)
}

func (c Car) Total() float64 {
	var total float64
	for _, p := range c.Products {
		total += p.TotalPrice()
	}
	return total
}

/*
es un constructor que crea y devuelve un nuevo objeto de tipo Car con el UserID especificado

1. Recibe un parámetro userID de tipo uint, que representa el identificador del usuario al que pertenece el carrito de compras.
2. Crea una nueva instancia de la estructura Car con el UserID establecido en el valor proporcionado.
3. Devuelve esta nueva instancia de Car con el UserID especificado.
*/
func NewCar(userID uint) Car {
	return Car{UserID: userID}
}
