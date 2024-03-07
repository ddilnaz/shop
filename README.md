Project description:
Tour_Shop is a Golang-based web application that serves as an online shopping platform. It encompasses various features for managing users, orders, and product items. Key functionalities include:

**User Management:**
- Create new users.
- Retrieve user information by ID.
- Update user information.
- Delete users.

**Order Management:**
- Create new orders.
- Retrieve order information by ID.
- Update order information.
- Delete orders.

**Product Item Management:**
- Create new product items.
- Retrieve product item information by ID.
- Update product item information.
- Delete product items.

**Technologies and Libraries Used:**
- Programming language: Golang.
- Database library: database/sql.
- Logging: log package.
- Context: context package.

**Project Structure:**
The project follows the Standard Project Layout for Golang applications. Key directories include `cmd` for main applications, `pkg` for reusable library code, and `model` for data models.

**API Endpoints:**
- User Management:
  - `POST /users`: Create a new user.
  - `GET /users/:id`: Get user information by ID.
  - `PUT /users/:id`: Update user information by ID.
  - `DELETE /users/:id`: Delete a user by ID.

- Order Management:
  - `POST /orders`: Create a new order.
  - `GET /orders/:id`: Get order information by ID.
  - `PUT /orders/:id`: Update order information by ID.
  - `DELETE /orders/:id`: Delete an order by ID.

- Product Item Management:
  - `POST /product-items`: Create a new product item.
  - `GET /product-items/:id`: Get product item information by ID.
  - `PUT /product-items/:id`: Update product item information by ID.
  - `DELETE /product-items/:id`: Delete a product item by ID.

**Database Structure:**
- User Model:
  ```
  Table users {
    id bigserial [primary key]
    created_at timestamp
    updated_at timestamp
    name text
    email text
  }
  ```

- Order Model:
  ```
  Table orders {
    id bigserial [primary key]
    created_at timestamp
    updated_at timestamp
    title text
    description text
    status text
  }
  ```

- Product Item Model:
  ```
  Table product_item {
    id bigserial [primary key]
    created_at timestamp
    updated_at timestamp
    title text
    description text
    price int
  }
  ```

This project adheres to REST principles and uses PostgreSQL as the underlying database for efficient data storage.
