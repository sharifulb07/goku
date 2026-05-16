package repository

import (
	"database/sql"
	"fmt"

	"github.com/sharifulb07/goku/internal/db"
	"github.com/sharifulb07/goku/internal/models"
)

type ResourceRepo struct {
}

func NewResourceRepo() *ResourceRepo {
	return &ResourceRepo{}
}

func (r *ResourceRepo) Create(resource *models.Resource)error{
	
	query:= `INSERT INTO resources (name, role, stack, email, status) VALUES ($1, $2, $3, $4, $5) RETURNING id, created_at, updated_at`


	err:=db.DB.QueryRow(query, resource.Name, resource.Role, resource.Stack, resource.Email, resource.Status).Scan(&resource.ID, &resource.CreatedAt, &resource.UpdatedAt)

	return err 
}


func (r *ResourceRepo) List() ([]models.Resource, error)  {

	row, err:=db.DB.Query(`SELECT id, name, role, stack, email, status FROM resources ORDER BY created_at DESC`)

	if err !=nil{
		return nil, err 
	}

	defer row.Close()


	var resources []models.Resource

	for row.Next() {
		var res models.Resource
		err:=row.Scan(&res.ID, &res.Name, &res.Role, &res.Stack, &res.Email, &res.Status)

		if err!=nil{
			return nil, err 
		}

		resources=append(resources, res)

	}

	return resources, nil 
	

}

func (r *ResourceRepo)GetByID(id int) (*models.Resource, error){

	var res models.Resource
	query:=`SELECT id, name, role, stack, email, status , created_at, update_at FROM resources WHERE id=$1`

	err:=db.DB.QueryRow(query, id).Scan(&res.ID, &res.Name, &res.Role, &res.Stack, &res.Status, &res.CreatedAt, &res.UpdatedAt)

	if err!=sql.ErrNoRows{
		return nil, fmt.Errorf("Resource not found ")
	}

	return &res, nil 


}




// Update updates a resource by ID
func (r *ResourceRepo) Update(id int, resource *models.Resource) error {
	query := `UPDATE resources 
	          SET name = $2, role = $3, stack = $4, email = $5, status = $6, updated_at = CURRENT_TIMESTAMP
	          WHERE id = $1`

	result, err := db.DB.Exec(query, id, resource.Name, resource.Role, 
		resource.Stack, resource.Email, resource.Status)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("resource with id %d not found", id)
	}

	return nil
}

// Delete deletes a resource by ID (Soft Delete - Recommended)
func (r *ResourceRepo) Delete(id int) error {
	query := `UPDATE resources SET deleted_at = CURRENT_TIMESTAMP, status = 'archived' WHERE id = $1`
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("resource with id %d not found", id)
	}

	fmt.Printf("✅ Resource with ID %d has been archived (soft deleted)\n", id)
	return nil
}

// HardDelete permanently deletes a resource
func (r *ResourceRepo) HardDelete(id int) error {
	query := `DELETE FROM resources WHERE id = $1`
	result, err := db.DB.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("resource with id %d not found", id)
	}

	fmt.Printf("✅ Resource with ID %d permanently deleted\n", id)
	return nil
}