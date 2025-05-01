# UniDirectory
## A list of Bangladeshi all university (UGC Approved) list.
UniDirectory is a simple and powerful REST API providing an up-to-date list of all UGC-approved universities in Bangladesh — including public, private, international, and CBHE (Cross-Border Higher Education) institutions.

[![Status](https://img.shields.io/endpoint?url=https://unidirectory.fly.dev/health&style=flat-square)](https://unidirectory.fly.dev/health)
[![Language](https://img.shields.io/badge/Built%20with-Go-00ADD8?logo=go&style=flat-square)](https://go.dev)
[![Hosted on Fly.io](https://img.shields.io/badge/Hosted%20on-Fly.io-1a1a1a?logo=fly.io&style=flat-square)](https://fly.io)
[![License](https://img.shields.io/badge/license-MIT-green?style=flat-square)](LICENSE)

Recently, I worked on a project that required displaying the names of all universities in Bangladesh. Since there was no available API providing this information, we had to manually list them in a JSON file. While learning Go and exploring web scraping, I thought—why not build an API myself?
So here we are. A directory of universities of Bangladesh. 

###  Live API Endpoint
Base URL: https://unidirectory.fly.dev

### GET `/universities`

Get a list of **all** universities (across all categories):

```
GET https://unidirectory.fly.dev/universities
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "University of Dhaka",
    "category": "Public",
    "website": "https://du.ac.bd",
    ...
  },
  ...
]
```
### GET `/universities?category={category}`

Filter universities by category:

```
GET https://unidirectory.fly.dev/universities?category=Private
```

**Available categories:**
- `Public`
- `Private`
- `International`
- `CBHE`

### GET `/universities/:id`

Get details of a specific university by its unique ID:

```
GET https://unidirectory.fly.dev/universities/5
```

## 🛠 Built With

- **Go+Gin** – Backend and API
- **GoColly** – For web scrapping
- **GORM + PostgreSQL** – Data storage
---


If you find this helpful, please consider leaving a star on the GitHub repo!
Your support motivates me to keep building useful tools!

