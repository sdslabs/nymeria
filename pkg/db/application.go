package db

import (
	"github.com/sdslabs/nymeria/helper"
)

func CreateApplication(name string, redirectURL string, allowedDomains string, organization string, clientKey string, clientSecret string) error {
	sqlStatement := `INSERT INTO application (name, redirect_url, allowed_domains, organization, created_at, client_key, client_secret) VALUES ($1, $2, $3, $4, now(), $5,$6);`
	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement, name, redirectURL, allowedDomains, organization, clientKey, clientSecret)

	if err != nil {
		return err
	}

	return nil

}

func UpdateApplication(id int, name string, redirectURL string, allowedDomains string, organization string) error {
	sqlStatement := `UPDATE application SET name=$1, redirect_url=$2, allowed_domains=$3, organization=$4 WHERE id=$5;`
	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement, name, redirectURL, allowedDomains, organization, id)

	if err != nil {
		return err
	}

	return nil

}

func DeleteApplication(id int) error {
	sqlStatement := `DELETE FROM application WHERE id=$1;`
	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement, id)

	if err != nil {
		return err
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

		err = rows.Scan(&t.ID, &t.Name, &t.RedirectURL, &t.AllowedDomains, &t.UpdatedAt, &t.Organization, &t.CreatedAt, &t.ClientKey, &t.ClientSecret)
		if err != nil {
			return nil, err
		}

		application = append(application, t)

	}

	return application, nil
}

func UpdateClientSecret(id int) error {
	sqlStatement := `UPDATE application SET client_secret=$1 WHERE id=$2;`
	db, err := Connection()

	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec(sqlStatement, helper.RandomString(30), id)

	if err != nil {
		return err
	}

	return nil

}

func GetApplication(client_key string, client_secret string) (Application, error) {
	sqlStatement := `SELECT * FROM application WHERE client_key=$1 AND client_secret=$2;`
	db, err := Connection()

	if err != nil {
		return Application{}, err
	}
	defer db.Close()

	var t Application

	err = db.QueryRow(sqlStatement, client_key, client_secret).Scan(&t.ID, &t.Name, &t.RedirectURL, &t.AllowedDomains, &t.UpdatedAt, &t.Organization, &t.CreatedAt, &t.ClientKey, &t.ClientSecret)

	if err != nil {
		return Application{}, err
	}

	return t, nil
}
