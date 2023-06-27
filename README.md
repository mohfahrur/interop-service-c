
# Intero Service C

Kode interoperabilitas service c yang digunakan untuk update spreadsheet google




## Installation

Clone from github

```bash
  git clone https://github.com/mohfahrur/interop-service-c.git
  cd my-project
```
Install library with Golang

```bash
  export GOPATH=/path/to/your/gopath
  go mod download
```
    
## Environment Variables

Membutuhkan environtment variable di dalam file .env untuk menjalankan projek ini

`spreadsheetID`



## Documentation

[Google Spreadsheet API](https://developers.google.com/sheets/api/quickstart/go)

[Tutorial Credential JSON Spreadsheet](https://dickyaryakesuma.medium.com/integrasi-google-sheet-dengan-golang-3369011a632)


## API Reference

#### Test API

```http
  GET /ping
```

#### Update Spreadsheet

```http
  POST /update-data
```

| Body | Type     | Description                       |
| :----| :------- | :-------------------------------- |
| `user` | `json` | **Required**. Nama user |
| `email` | `json` | **Required**. Email user |
| `hp` | `json` | **Required**. Nomor telepon user |
| `item` | `json` | **Required**. Item user |


