package database

import (
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
			{{bedrijfsNummer}} CHAR(4)  NOT NULL CHECK(length({{bedrijfsNummer}}) = 4),
			{{documentNummer}} CHAR(10) NOT NULL CHECK(length({{documentNummer}}) = 10),
			{{boekJaar}}       CHAR(4)  NOT NULL CHECK(length({{boekJaar}}) = 4),
			{{documentSoort}}  CHAR(2)           CHECK(length({{documentSoort}}) = 2),
			{{documentDatum}}  CHAR(8)           CHECK(length({{documentDatum}}) = 8),
			{{boekingDatum}}   CHAR(8)           CHECK(length({{boekingDatum}}) = 8),
			{{boekMaand}}      CHAR(2)           CHECK(length({{boekMaand}}) = 2),
			{{invoerDatum}}    CHAR(8)           CHECK(length({{invoerDatum}}) = 8),
			{{invoerTijd}}     CHAR(6)           CHECK(length({{invoerTijd}}) = 6),

			PRIMARY KEY ({{bedrijfsNummer}}, {{documentNummer}}, {{boekJaar}})
		)`,
		map[string]string{
			"tableName":      TableDocumentKop,
			"bedrijfsNummer": ColumnBedrijfsNummer,
			"documentNummer": ColumnDocumentNummer,
			"boekJaar":       ColumnBoekJaar,
			"documentSoort":  ColumnDocumentSoort,
			"documentDatum":  ColumnDocumentDatum,
			"boekingDatum":   ColumnBoekingDatum,
			"boekMaand":      ColumnBoekMaand,
			"invoerDatum":    ColumnInvoerDatum,
			"invoerTijd":     ColumnInvoerTijd,
		},
	)

	if _, err := db.Exec(statement); err != nil {
		return fmt.Errorf("failed to create table %s: %w", TableDocumentKop, err)
	}

	return nil
}

func createSegmentTable(db *sql.DB) error {
	statement := mustache.Render(`
		CREATE TABLE {{tableName}} (
			{{bedrijfsNummer}}            CHAR(4)  NOT NULL CHECK(length({{bedrijfsNummer}}) = 4),
			{{documentNummer}}            CHAR(10) NOT NULL CHECK(length({{documentNummer}}) = 10),
			{{boekJaar}}                  CHAR(4)  NOT NULL CHECK(length({{boekJaar}}) = 4),
			{{regelNummer}}               CHAR(3)  NOT NULL CHECK(length({{regelNummer}}) = 3),
			{{regelId}}                   CHAR(1)           CHECK(length({{regelId}}) = 1),
			{{vereffeningDatum}}          CHAR(8)           CHECK(length({{vereffeningDatum}}) = 8),
			{{vereffeningInvoerDatum}}    CHAR(8)           CHECK(length({{vereffeningInvoerDatum}}) = 8),
			{{vereffeningDocumentNummer}} CHAR(10)          CHECK(length({{vereffeningDocumentNummer}}) = 10),
			{{boekingSleutel}}            CHAR(2)           CHECK(length({{boekingSleutel}}) = 2),

			PRIMARY KEY ({{bedrijfsNummer}}, {{documentNummer}}, {{boekJaar}}, {{regelNummer}}),
			CONSTRAINT BSEG_FK FOREIGN KEY ({{bedrijfsNummer}}, {{documentNummer}}, {{boekJaar}}) REFERENCES {{documentKop}} ({{bedrijfsNummer}}, {{documentNummer}}, {{boekJaar}})
		)`,
		map[string]string{
			"tableName":                 TableDocumentSegment,
			"bedrijfsNummer":            ColumnBedrijfsNummer,
			"documentNummer":            ColumnDocumentNummer,
			"boekJaar":                  ColumnBoekJaar,
			"regelNummer":               ColumnBoekingsregelNummer,
			"regelId":                   ColumnBoekingRegelID,
			"vereffeningDatum":          ColumnVereffeningDatum,
			"vereffeningInvoerDatum":    ColumnVereffeningInvoerDatum,
			"vereffeningDocumentNummer": ColumnVereffeningsDocumentNummer,
			"boekingSleutel":            ColumnBoekingssleutel,
			"documentKop":               TableDocumentKop,
		},
	)

	if _, err := db.Exec(statement); err != nil {
		return fmt.Errorf("failed to create table %s: %w", TableDocumentSegment, err)
	}

	return nil
}
