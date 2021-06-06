// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Admin struct {
	ID       string `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Category struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	ImageURL string     `json:"image_url"`
	Products []*Product `json:"products"`
}

type Customer struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Phone    string   `json:"phone"`
	Address  string   `json:"address"`
	Region   string   `json:"region"`
	CcNumber string   `json:"cc_number"`
	Orders   []*Order `json:"orders"`
}

type Image struct {
	ID  string `json:"id"`
	URL string `json:"url"`
}

type Order struct {
	ID              string            `json:"id"`
	Amount          int               `json:"amount"`
	CreatedAt       string            `json:"created_at"`
	OrderedProducts []*OrderedProduct `json:"ordered_products"`
}

type OrderedProduct struct {
	ID         string      `json:"id"`
	SubProduct *SubProduct `json:"sub_product"`
}

type Product struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Fabric       string        `json:"fabric"`
	Manufacturer string        `json:"manufacturer"`
	Images       []*Image      `json:"images"`
	SubProducts  []*SubProduct `json:"sub_products"`
}

type SubProduct struct {
	ID     string  `json:"id"`
	Price  float64 `json:"price"`
	Size   string  `json:"size"`
	Color  string  `json:"color"`
	Amount int     `json:"amount"`
}