package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Products struct {
	ID    string
	Name  string
	Price float64
}

func NewProducts(name string, price float64) *Products {
	return &Products{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/mydb")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	product := NewProducts("Macbook", 6890.00)

	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}

	product.Price = 5990.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	// p, err := selectOneProduct(db, product.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("O product %+v, possui o preço de: %.2f", p.Name, p.Price)

	products, err := selectAllProducts(db)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nLista de produtos:")
	for _, prod := range products {
		fmt.Printf("ID: %s, Name: %s, Price: %.2f\n", prod.ID, prod.Name, prod.Price)
	}

	err = deleteProduct(db, product.ID)
	if err != nil {
		panic(err)
	}
}

func insertProduct(db *sql.DB, product *Products) error {
	stmt, err := db.Prepare("INSERT INTO products (id, name, price) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Products) error {
	stmt, err := db.Prepare("UPDATE products SET name = ?, price = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Products, error) {
	stmt, err := db.Prepare("SELECT id, name, price FROM products WHERE id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var product Products
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func selectAllProducts(db *sql.DB) ([]*Products, error) {
	rows, err := db.Query("SELECT id, name, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*Products
	for rows.Next() {
		var product Products
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, &product)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}
