package db

import (
	"errors"
	"fmt"

	"github.com/sdslabs/nymeria/helper"
)

func CreateApplication(name, redirectURL, allowedDomains, organization, clientKey, clientSecret string) error {
	if name == "" {
		return errors.New("400 application name is required")
	}
	if redirectURL == "" {
		return errors.New("400 redirect URL is required")
	}
	if allowedDomains == "" {
		return errors.New("400 allowed domains is required")
	}
	if organization == "" {
		return errors.New("400 organization is required")
	}
	if clientKey == "" {
		return errors.New("400 client key is required")
	}
	if clientSecret == "" {
		return errors.New("400 client secret is required")
	}

	sqlStatement := `INSERT INTO application (name, redirect_url, allowed_domains, organization, created_at, client_key, client_secret) 
					 VALUES ($1, $2, $3, $4, now(), $5, $6);`

	db, err := Connection()
	if err != nil {
		return fmt.Errorf("500 database connection error: %w", err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement, name, redirectURL, allowedDomains, organization, clientKey, clientSecret)
	if err != nil {
		return fmt.Errorf("500 failed to create application: %w", err)
	}

	return nil
}

func UpdateApplication(id int, name, redirectURL, allowedDomains, organization string) error {
	if name == "" {
		return errors.New("400 application name is required")
	}
	if redirectURL == "" {
		return errors.New("400 redirect URL is required")
	}
	if allowedDomains == "" {
		return errors.New("400 allowed domains is required")
	}
	if organization == "" {
		return errors.New("400 organization is required")
	}

	sqlStatement := `UPDATE application SET name=$1, redirect_url=$2, allowed_domains=$3, organization=$4 WHERE id=$5;`

	db, err := Connection()
	if err != nil {
		return fmt.Errorf("500 database connection error: %w", err)
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement, name, redirectURL, allowedDomains, organization, id)
	if err != nil {
		return fmt.Errorf("500 failed to update application: %w", err)
	}

	return nil
}

func DeleteApplication(id int) error {
	sqlStatement := `DELETE FROM application WHERE id=$1;`

	db, err := Connection()
	if err != nil {
		return fmt.Errorf("500 database connection error: %w", err)
	}
	defer db.Close()

	result, err := db.Exec(sqlStatement, id)
	if err != nil {
		return fmt.Errorf("500 failed to delete application: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("500 error checking affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return errors.New("404 application not found")
	}

	return nil
}

func GetAllApplication() ([]Application, error) {
	sqlStatement := `SELECT * FROM application;`
	db, err := Connection()

	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(sqlStatement)

	if err != nil {
		return nil, err
	}

	var application []Application

	for rows.Next() {
		var t Application

		err = rows.Scan(&t.ID, &t.Name, &t.RedirectURL, &t.AllowedDomains, &t.Organization, &t.CreatedAt, &t.ClientKey, &t.ClientSecret)
		if err != nil {
			return nil, err
		}

		application = append(application, t)

	}

	return application, nil
}

func UpdateClientSecret(id int) (string, error) {

	if id == 0 {
		return "", errors.New("400 id is required")
	}

	sqlStatement := `UPDATE application SET client_secret=$1 WHERE id=$2;`
	db, err := Connection()

	if err != nil {
		return "", err
	}
	defer db.Close()

	newSecret := helper.RandomString(30)

	_, err = db.Exec(sqlStatement, newSecret, id)

	if err != nil {
		return "", err
	}

	return newSecret, nil

}

func UpdateClientKey(id int) (string, error) {

	if id == 0 {
		return "", errors.New("400 id is required")
	}

	sqlStatement := `UPDATE application SET client_key=$1 WHERE id=$2;`
	db, err := Connection()

	if err != nil {
		return "", err
	}
	defer db.Close()

	newKey := helper.RandomString(30)

	_, err = db.Exec(sqlStatement, newKey, id)

	if err != nil {
		return "", err
	}

	return newKey, nil
}

func GetApplication(client_key string, client_secret string) (Application, error) {
	sqlStatement := `SELECT * FROM application WHERE client_key=$1 AND client_secret=$2;`
	db, err := Connection()

	if err != nil {
		return Application{}, err
	}
	defer db.Close()

	var t Application

	err = db.QueryRow(sqlStatement, client_key, client_secret).Scan(&t.ID, &t.Name, &t.RedirectURL, &t.AllowedDomains, &t.Organization, &t.CreatedAt, &t.ClientKey, &t.ClientSecret)

	if err != nil {
		return Application{}, err
	}

	return t, nil
}

// GetApplicationByKey retrieves an application using only the client key
func GetApplicationByKey(client_key string) (Application, error) {
	sqlStatement := `SELECT * FROM application WHERE client_key=$1;`
	db, err := Connection()

	if err != nil {
		return Application{}, err
	}
	defer db.Close()

	var t Application

	err = db.QueryRow(sqlStatement, client_key).Scan(&t.ID, &t.Name, &t.RedirectURL, &t.AllowedDomains, &t.Organization, &t.CreatedAt, &t.ClientKey, &t.ClientSecret)

	if err != nil {
		return Application{}, err
	}

	return t, nil
}
