# ComponentHub Service

**ComponentHub Service** is a backend API for managing and distributing UI components (React/Vue/JSX, etc.) in a modular, scalable, and easy-to-integrate way for modern frontend marketplaces.

## 🚀 Main Features

- CRUD for Components, Categories, and Tags
- Advanced Filtering & Search (by tag, category, status, approval, keyword)
- Approval workflow (admin/reviewer)
- Pagination & Sorting
- Tagging system (many-to-many relationships)
- Standardized API responses (success/error)
- Healthcheck endpoint for monitoring
- Automatic API documentation (Swagger/OpenAPI)
- Modular, scalable project structure

---

## 📚 Technologies Used

- **Golang (Go) v1.20+**
- **Gin Web Framework** – Routing, middleware, JSON API
- **GORM** – ORM for PostgreSQL
- **PostgreSQL** – Main database
- **swaggo/gin-swagger** – Automatic Swagger documentation
- **uuid** – For primary keys and entity relationships

---

## 🗂️ Project Structure

```
service_components/
├── cmd/
│   └── main.go
├── internal/
│   ├── config/           # Configuration (DB, env, etc.)
│   ├── database/         # DB initialization & seeder
│   ├── handler/          # HTTP handlers (component, category, tag, etc.)
│   ├── model/            # GORM models (Component, Category, Tag)
│   ├── utils/            # API response helpers, error handling, etc.
├── docs/                 # Auto-generated Swagger documentation
├── go.mod
├── go.sum
└── README.md
```

---

## ⚡ Quick Start (Local Setup)

1. **Clone the repository**
   ```bash
   git clone <repo-url>
   cd service_components
   ```

2. **Setup PostgreSQL**
   - Create a new database (e.g., `componenthub_dev`)
   - Ensure user, password, and port match your config/env file

3. **Install dependencies**
   ```bash
   go mod tidy
   ```

4. **Run the application**
   ```bash
   go run cmd/main.go
   ```
   The server will run on `localhost:8080`

5. **Access Swagger documentation**
   - Open: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 🌐 Main API Endpoints

### Component

- **Create Component**
  - `POST /api/v1/components`
  - Body:
    ```json
    {
      "name": "Button",
      "description": "Reusable button",
      "category_id": "UUID",
      "code_jsx": "<button>...",
      "code_css": ".btn {...}",
      "props_definition": { ... }
    }
    ```

- **Get All Components (with filter/search)**
  - `GET /api/v1/components?tag=React,UI&category=ui-kit&status=published&q=button&page=1&limit=20&sort=created_at&order=desc`

- **Get Component by Slug**
  - `GET /api/v1/components/{slug}`

- **Update Component**
  - `PATCH /api/v1/components/{slug}`
  - Body:
    ```json
    {
      "name": "New Button",
      "description": "Updated description"
    }
    ```

- **Delete Component**
  - `DELETE /api/v1/components/{slug}`

- **Add Tag to Component**
  - `POST /api/v1/components/{slug}/tags`
  - Body: `{ "tag_id": "UUID" }`

- **Update Component Status**
  - `PATCH /api/v1/components/{slug}/status`
  - Body: `{ "status": "published" }`

- **Update Component Approval**
  - `PATCH /api/v1/components/{slug}/approval`
  - Body: `{ "approval_status": "approved", "reviewer_id": "UUID" }`

---

### Category & Tag

- **Create/Get Category**
  - `POST /api/v1/categories` `{ "name": "UI Kit" }`
  - `GET /api/v1/categories`

- **Create/Get Tag**
  - `POST /api/v1/tags` `{ "name": "React" }`
  - `GET /api/v1/tags`

---

### Healthcheck

- `GET /health`  
  Response: `{ "status": "ok", "service": "component-service" }`

---

## 📑 API Response & Error Format

All responses follow these standards:
- **Success:**  
  ```json
  {
    "status": "success",
    "data": { ... }
  }
  ```
- **Error:**  
  ```json
  {
    "status": "error",
    "message": "Detailed error message"
  }
  ```

---

## 🔒 Best Practices & Highlights

- Clean separation between handler, model, utils, and config (clean architecture)
- Consistent API response format for easy frontend integration
- Flexible filtering, search, pagination, and sorting
- Healthcheck endpoint for easy monitoring in production
- Auto-generated Swagger docs for fast onboarding and integration
- Ready for JWT authentication and role-based access (can be added easily)
- Designed for easy scaling and microservice expansion

---

## 🛠️ Development & Testing

- **Manual Testing:** Use Postman for all endpoint combinations (see example requests above).
- **Swagger:** Always up-to-date, auto-generated via `swag init`.
- **Unit Test (optional):** Add unit tests in `internal/handler/` for main request handlers.

---

## 📩 Contact & Contribution

- **Author:** [Ghalib Assaidy](https://github.com/ghalibassaidy)
- For questions or contributions, please open an issue or pull request in this repository.

---

## 📋 Next Development Notes

- Add JWT token & role-based authorization (optional, for production)
- Add unit/integration tests
- Integrate with notification/event service (e.g., for approval events)
- Database query and indexing optimization for large scale

---

**Thank you for checking out ComponentHub Service!**
