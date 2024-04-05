info about me :
-Образовательная программа	Информационные системы

-ID	22B031177

-Студент	Зейнолла Дильназ Бауыржанқызы

-Год обучения	2

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
   TABLE  users (
    id          bigserial PRIMARY KEY,
    name        text NOT NULL,
    email       text NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
  ```

- Order Model:
  ```
  Table orders {
      order_id    bigserial PRIMARY KEY,
    item_id     bigserial REFERENCES product_item (id),
    user_id     bigserial REFERENCES users (id),
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    status      text NOT NULL DEFAULT 'Pending'
);
  }
  ```

- Product Item Model:
  ```
  Table product_item {
    id          bigserial PRIMARY KEY,
    title       text NOT NULL,
    description text,
    price       int NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
  }
  ```

This project adheres to REST principles and uses PostgreSQL as the underlying database for efficient data storage.
