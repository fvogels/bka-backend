# Table

## `BKPF`

| Field | Type | Length | Description | Example Values |
| ----- | ---- | ------ | ----------- | -------------- |
| `BUKRS` | `CHAR` | 4 | Bedrijfsnummer | 1000, 2000, 3000, ... |
| `BELNR` | `CHAR` | 10 | Documentnummer boekhoudingsdocument | 1900000001, 1900000002, ... |
| `GJAHR` | `CHAR` | 4 | Boekjaar | 2024 |
| `BLART` | `CHAR` | 2 | Documentsoort | SA, KR, DR, DZ, KZ, AB |
| `BLDAT` | `DATUM` | 8 | Documentdatum in document `YYYYMMDD` | 20240115 |
| `BUDAT` | `DATUM` | 8 | Boekingsdatum in document `YYYYMMDD` | 20240115 |
| `MONAT` | `CHAR` | 2 | Boekmaand (01-12, special periods 13-16) | 01 |
| `CPUDT` | `DATUM` | 8 | Datum waarop boekhoudingsdocument is ingevoerd | 20240115  |
| `CPUTM` | `TIJD` | 6 | Tijd waarop gegevens zijn ingevoerd `HHMMSS` | 101523 |

Typical constraints:

* `BLDAT <= BUDAT <= CPUDT`
* `MONAT = month(BUDAT)`

## `BSEG`

| Field | Type | Length | Description | Example Values |
| ----- | ---- | ------ | ----------- | -------------- |
| `BUKRS` | `CHAR` | 4 | Bedrijfsnummer | 1000, 2000, 3000, ... |
| `BELNR` | `CHAR` | 10 | Documentnummer boekhoudingsdocument | 1900000001, 1900000002, ... |
| `GJAHR` | `CHAR` | 4 | Boekjaar | 2024 |
| `BUZEI` | `CHAR` | 3 | Nummer van boekingsregel in boekhoudingsdoc | 001, 002, 003, ... |
| `BUZID` | `CHAR` | 1 | Identificatie van boekingsregel | Often blank |
| `AUGDT` | `DATUM` | 8 | Datum van vereffening | 00000000, 20240220 |
| `AUGCP` | `DATUM` | 8 | Invoerdatum van de vereffening | 00000000, 20240220 |
| `AUGBL` | `CHAR` | 10 | Documentnummer van vereffeningsdocument | 0000000000, 1900000010 |
| `BSCHL` | `CHAR` | 2 | Boekingssleutel | 01, 11, 21, 31, 40, 50 |

* Each document has at least two segments
  * One with `BSCHL=40`, one with `BSCHL=50`
* If not `0`, `AUGBL` must refer to valid `BKPF` row.