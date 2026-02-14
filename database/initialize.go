package database

import (
	"bass-backend/database/meta"
	"database/sql"
	"fmt"

	"github.com/hoisie/mustache"
)

func InitializeDatabase(database *sql.DB) error {
	if err := createHeaderTable(database); err != nil {
		return err
	}

	if err := createSegmentTable(database); err != nil {
		return err
	}

	return nil
}

func createHeaderTable(db *sql.DB) error {
	statement := mustache.Render(`
		CREATE TABLE {{tableName}} (
			{{bedrijfsnummer}} CHAR(4)  NOT NULL CHECK(length({{bedrijfsnummer}}) = 4),
			{{documentnummer}} CHAR(10) NOT NULL CHECK(length({{documentnummer}}) = 10),
			{{boekjaar}}       CHAR(4)  NOT NULL CHECK(length({{boekjaar}}) = 4),
			{{documentsoort}}  CHAR(2)           CHECK(length({{documentsoort}}) = 2),
			{{documentdatum}}  CHAR(8)           CHECK(length({{documentdatum}}) = 8),
			{{boekingdatum}}   CHAR(8)           CHECK(length({{boekingdatum}}) = 8),
			{{boekmaand}}      CHAR(2)           CHECK(length({{boekmaand}}) = 2),
			{{invoerdatum}}    CHAR(8)           CHECK(length({{invoerdatum}}) = 8),
			{{invoertijd}}     CHAR(6)           CHECK(length({{invoertijd}}) = 6),

			PRIMARY KEY ({{bedrijfsnummer}}, {{documentnummer}}, {{boekjaar}})
		)`,
		map[string]string{
			"tableName":      meta.DocumentKop.Table,
			"bedrijfsnummer": meta.DocumentKop.BedrijfsNummer,
			"documentnummer": meta.DocumentKop.DocumentNummer,
			"boekjaar":       meta.DocumentKop.BoekJaar,
			"documentsoort":  meta.DocumentKop.DocumentSoort,
			"documentdatum":  meta.DocumentKop.DocumentDatum,
			"boekingdatum":   meta.DocumentKop.BoekingDatum,
			"boekmaand":      meta.DocumentKop.BoekMaand,
			"invoerdatum":    meta.DocumentKop.InvoerDatum,
			"invoertijd":     meta.DocumentKop.InvoerTijd,
		},
	)

	if _, err := db.Exec(statement); err != nil {
		return fmt.Errorf("failed to create table %s: %w", meta.DocumentKop.Table, err)
	}

	return nil
}

func createSegmentTable(db *sql.DB) error {
	statement := mustache.Render(`
		CREATE TABLE {{tableName}} (
			{{bedrijfsnummer}}            CHAR(4)  NOT NULL CHECK(length({{bedrijfsnummer}}) = 4),
			{{documentnummer}}            CHAR(10) NOT NULL CHECK(length({{documentnummer}}) = 10),
			{{boekJaar}}                  CHAR(4)  NOT NULL CHECK(length({{boekJaar}}) = 4),
			{{regelNummer}}               CHAR(3)  NOT NULL CHECK(length({{regelNummer}}) = 3),
			{{regelId}}                   CHAR(1)           CHECK(length({{regelId}}) = 1),
			{{vereffeningDatum}}          CHAR(8)           CHECK(length({{vereffeningDatum}}) = 8),
			{{vereffeningInvoerDatum}}    CHAR(8)           CHECK(length({{vereffeningInvoerDatum}}) = 8),
			{{vereffeningDocumentnummer}} CHAR(10)          CHECK(length({{vereffeningDocumentnummer}}) = 10),
			{{boekingSleutel}}            CHAR(2)           CHECK(length({{boekingSleutel}}) = 2),

			PRIMARY KEY ({{bedrijfsnummer}}, {{documentnummer}}, {{boekJaar}}, {{regelNummer}}),
			CONSTRAINT BSEG_FK FOREIGN KEY ({{bedrijfsnummer}}, {{documentnummer}}, {{boekJaar}}) REFERENCES {{documentKop}} ({{bedrijfsnummer}}, {{documentnummer}}, {{boekJaar}})
		)`,
		map[string]string{
			"tableName":                 meta.DocumentSegment.Table,
			"bedrijfsnummer":            meta.DocumentSegment.BedrijfsNummer,
			"documentnummer":            meta.DocumentSegment.DocumentNummer,
			"boekJaar":                  meta.DocumentSegment.BoekJaar,
			"regelNummer":               meta.DocumentSegment.BoekingsregelNummer,
			"regelId":                   meta.DocumentSegment.BoekingRegelID,
			"vereffeningDatum":          meta.DocumentSegment.VereffeningDatum,
			"vereffeningInvoerDatum":    meta.DocumentSegment.VereffeningInvoerDatum,
			"vereffeningDocumentnummer": meta.DocumentSegment.VereffeningsDocumentNummer,
			"boekingSleutel":            meta.DocumentSegment.Boekingssleutel,
			"documentKop":               meta.DocumentKop.Table,
		},
	)

	if _, err := db.Exec(statement); err != nil {
		return fmt.Errorf("failed to create table %s: %w", meta.DocumentSegment.Table, err)
	}

	return nil
}
