package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/rizkyfazri23/dripay/model/entity"
)

type PaymentTypeRepo interface {
	CreateType(newType *entity.TransactionTypeInput) (entity.TransactionType, error)
	ReadAllType() ([]entity.TransactionType, error)
	ReadTypeById(typeID int) (entity.TransactionType, error)
	UpdateType(typeID int, typeEdit *entity.TransactionTypeInput) (entity.TransactionType, error)
	DeleteType(typeID int) error
}

type paymentTypeRepo struct {
	db *sql.DB
}

func (r *paymentTypeRepo) CreateType(newType *entity.TransactionTypeInput) (entity.TransactionType, error) {
	var typeOutput entity.TransactionType
	query := "INSERT INTO m_transaction_type(type_name, description) VALUES($1, $2) RETURNING type_id, type_name, description"
	err := r.db.QueryRow(query, newType.TypeName, newType.Description).Scan(&typeOutput.TypeId, &typeOutput.TypeName, &typeOutput.Description)
	if err != nil {
		return entity.TransactionType{}, err
	}
	return typeOutput, nil
}

func (r *paymentTypeRepo) ReadAllType() ([]entity.TransactionType, error) {
	var transactionTypeList []entity.TransactionType
	query := "SELECT type_id, type_name, desription FROM m_transaction_type"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var transactionType entity.TransactionType
		if err = rows.Scan(&transactionType.TypeId, &transactionType.TypeName, &transactionType.Description); err != nil {
			log.Println(err)

			return nil, err
		}
		transactionTypeList = append(transactionTypeList, transactionType)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)

		return nil, err
	}
	return transactionTypeList, nil
}

func (r *paymentTypeRepo) ReadTypeById(typeID int) (entity.TransactionType, error) {
	var getType entity.TransactionType
	query := "SELECT type_id, type_name, desription FROM m_transaction_type WHERE type_id = $1 RETURNING type_id, type_name, desription"
	row := r.db.QueryRow(query, typeID)
	err := row.Scan(&getType.TypeId, &getType.TypeName, &getType.Description)
	if err == sql.ErrNoRows {
		log.Println(err)
		return entity.TransactionType{}, fmt.Errorf("unidentified transaction type")
	} else if err != nil {
		log.Println(err)
	}
	return getType, nil
}

func (r *paymentTypeRepo) UpdateType(typeID int, typeEdit *entity.TransactionTypeInput) (entity.TransactionType, error) {
	var typeInformation entity.TransactionType
	query := "UPDATE m_transaction_type SET type_name = $1, description = &2 WHERE type_id = $3 "
	row := r.db.QueryRow(query, typeEdit.TypeName, typeEdit.Description, typeID)
	err := row.Scan(&typeInformation.TypeId, &typeInformation.TypeName, &typeInformation.Description)
	if err != nil {
		log.Println(err)
		return entity.TransactionType{}, err
	}
	return typeInformation, nil
}

func (r *paymentTypeRepo) DeleteType(typeID int) error {
	query := "DELETE FROM n_transaction_type WHERE type_id = $1"
	result, err := r.db.Exec(query, typeID)
	if err != nil {
		log.Println(err)
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println(err)
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("unidentified type")
	}
	log.Println("Type deleted")
	return nil
}

func newTransactionTypeRepo(db *sql.DB) PaymentTypeRepo {
	repo := new(paymentTypeRepo)
	repo.db = db
	return repo
}
