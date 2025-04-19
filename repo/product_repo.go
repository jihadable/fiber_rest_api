package repo

import (
	"context"
	"database/sql"
	"fiber_rest_api/model"
	"log"
)

type ProductRepo interface {
	FindById(ctx context.Context, productId int) (model.Product, error)
	FindAll(ctx context.Context) ([]model.Product, error)
	Save(ctx context.Context, product model.Product) (model.Product, error)
	Update(ctx context.Context, product model.Product) (model.Product, error)
	Delete(ctx context.Context, productId int) error
}

type productRepoImpl struct {
	DB *sql.DB
}

func NewProductRepo(db *sql.DB) ProductRepo {
	return &productRepoImpl{DB: db}
}

func (repo *productRepoImpl) FindById(ctx context.Context, productId int) (model.Product, error) {
	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	sqlQuery := "SELECT id, name, quantity FROM products WHERE id = ?"
	stmt, err := tx.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	row := stmt.QueryRowContext(ctx, productId)

	product := model.Product{}

	err = row.Scan(&product.Id, &product.Name, &product.Quantity)
	if err != nil {
		return product, err
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return product, nil
}

func (repo *productRepoImpl) FindAll(ctx context.Context) ([]model.Product, error) {
	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	sqlQuery := "SELECT id, name, quantity FROM products"
	stmt, err := tx.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []model.Product{}
	for rows.Next() {
		product := model.Product{}

		err := rows.Scan(&product.Id, &product.Name, &product.Quantity)
		if err != nil {
			log.Println("Scan error:", err)
			continue
		}

		products = append(products, product)
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return products, nil
}

func (repo *productRepoImpl) Save(ctx context.Context, product model.Product) (model.Product, error) {
	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	sqlQuery := "INSERT INTO products (name, quantity) VALUES (?, ?)"
	stmt, err := tx.PrepareContext(ctx, sqlQuery)
	if err != nil {
		return product, err
	}
	defer stmt.Close()

	result, err := stmt.ExecContext(ctx, product.Name, product.Quantity)
	if err != nil {
		return product, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return product, err
	}
	product.Id = int(id)

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return product, nil
}

func (repo *productRepoImpl) Update(ctx context.Context, product model.Product) (model.Product, error) {
	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	sqlQuery := "UPDATE products SET name = ?, quantity = ? WHERE id = ?"
	stmt, err := tx.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product.Name, product.Quantity, product.Id)
	if err != nil {
		return product, err
	}

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return product, nil
}

func (repo *productRepoImpl) Delete(ctx context.Context, productId int) error {
	tx, err := repo.DB.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	sqlQuery := "DELETE FROM products WHERE id = ?"
	stmt, err := tx.PrepareContext(ctx, sqlQuery)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, productId)

	if err := tx.Commit(); err != nil {
		panic(err)
	}

	return err
}
