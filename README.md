
# Data Penduduk

Data penduduk adalah API untuk mengelola data penduduk termasuk provinsi, kabupaten/kota, kecamatan, dan informasi pribadi penduduk.

## Endpoints API

### Province

#### 1. Menambahkan Provinsi

- **Method:** POST
- **Endpoint:** `/province`
- **Payload Contoh:**
  ```json
  {
      "id": "01",
      "name": "DKI Jakarta",
  }

 #### 2.Mendapatkan Semua Provinsi

- **Method:** GET
- **Endpoint:** `/province`

 #### 3. Mengupdate Provinsi
- **Method:** PUT
- **Endpoint:** `/province/:id`
- **Payload Contoh:
```json
  {
    "name": "DKI Jakarta Baru",
  }
```
#### 4. Menghapus Provinsi
- **Method:** DELETE
- **Endpoint:** `/province/:id`

### Regency

#### 1. Menambahkan Regency

- **Method:** POST
- **Endpoint:** `/regency`
- **Payload Contoh:**
  ```json
  {
      "id": "01",
      "name": "Jakarta Timur",
  }

 #### 2.Mendapatkan Semua Regency

- **Method:** GET
- **Endpoint:** `/regency`

 #### 3. Mengupdate Regency
- **Method:** PUT
- **Endpoint:** `/regency/:id`
- **Payload Contoh:
```json
  {
    "name": "Jakarta Timur Baru",
  }
```
#### 4. Menghapus Regency
- **Method:** DELETE
- **Endpoint:** `/regency/:id`

 ### District

#### 1. Menambahkan District

- **Method:** POST
- **Endpoint:** `/district`
- **Payload Contoh:**
  ```json
  {
      "id": "01",
      "name": "Gedong",
  }

 #### 2.Mendapatkan Semua District

- **Method:** GET
- **Endpoint:** `/district`

 #### 3. Mengupdate District
- **Method:** PUT
- **Endpoint:** `/district/:id`
- **Payload Contoh:
```json
  {
    "name": "Gedong Baru",
  }
```
#### 4. Menghapus District
- **Method:** DELETE
- **Endpoint:** `/district/:id`

### People

#### 1. Menambahkan People

- **Method:** POST
- **Endpoint:** `/people`
- **Payload Contoh:**
  ```json
  {
    "name": "Taufik",
    "gender": "laki-laki",
    "dob": "25-Jul-1998",
    "pob": "Jakarta",
    "province_id": "02",
    "regency_id": "02",
    "district_id": "02"
    }


 #### 2.Mendapatkan Semua People

- **Method:** GET
- **Endpoint:** `/people`

 #### 3.Mendapatkan People by NIK

- **Method:** GET
- **Endpoint:** `/people/:nik`

 #### 4. Mengupdate People
- **Method:** PUT
- **Endpoint:** `/people/:id`
- **Payload Contoh:
```json
 {
    "name": "Taufik Updated",
    "province_id": "02",
    "regency_id": "0202",
    "district_id": "01"
}
```
#### 5. Menghapus People
- **Method:** DELETE
- **Endpoint:** `/regency/:id`
