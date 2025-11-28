package catalog

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Save(product Product) error
	FindByID(id string) (Product, error)
	FindAll() ([]Product, error)
	Update(product Product) error
	Delete(id string) error
}

type pgRepository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &pgRepository{db: db}
}

func (r *pgRepository) Save(product Product) error {
	query := `
		INSERT INTO products (
			id, name, description, price, quantity_in_inventory, category_id, created_at, updated_at
		) VALUES (
			:id, :name, :description, :price, :quantity_in_inventory, :category_id, :created_at, :updated_at
		)
	`

	_, err := r.db.NamedExec(query, product)

	if err != nil {
		return fmt.Errorf("falha ao salvar produto %w", err)
	}
	return nil
}

func (r *pgRepository) FindByID(id string) (Product, error) {
	query := `
		SELECT * FROM products WHERE id = $1
	`

	var product Product

	err := r.db.Get(&product, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Product{}, fmt.Errorf("produto não encontrado: %w", err)
		}

		return Product{}, fmt.Errorf("falha ao buscar produto: %w", err)
	}
	return product, nil
}

func (r *pgRepository) FindAll() ([]Product, error) {
	query := `SELECT * FROM products`

	var products []Product

	err := r.db.Select(&products, query)

	if err != nil {
		return nil, fmt.Errorf("falha ao buscar todos os produtos: %w", err)
	}
	return products, nil
}

func (r *pgRepository) Update(product Product) error {
	query := `
		UPDATE products SET
			name = :name,
			description = :description,
			price = :price,
			quantity_in_inventory = :quantity_in_inventory,
			updated_at = NOW()
		WHERE id = :id
	`

	result, err := r.db.NamedExec(query, product)
	if err != nil {
		return fmt.Errorf("falha ao executar update: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("falha ao verificar linhas afetadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produto com ID %s não encontrado para atualização", product.ID)
	}

	return nil
}

func (r *pgRepository) Delete(id string) error {
	query := `DELETE FROM products WHERE id = :id`

	args := map[string]interface{}{
		"id": id,
	}

	result, err := r.db.NamedExec(query, args)
	if err != nil {
		return fmt.Errorf("falha ao executar delete: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("falha ao verificar linhas afetadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("produto com ID %s não encontrado para deleção", id)
	}

	return nil
}
