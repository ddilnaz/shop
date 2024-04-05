### Project Overview:

**Project Name:** Tour_Shop  
**Programming Language:** Golang  
**Database:** PostgreSQL  
**Educational Program:** Information Systems  
**Student:** Zeynolla Dilnaz Bauyrzhanovna  
**Year of Study:** 2nd

### Description:

Tour_Shop is a web application developed in Golang, serving as an online shopping platform. It facilitates various functionalities including user, order, and product item management.

### Key Features:

#### User Management:
- Create, retrieve, update, and delete users.
- Access user information by ID.

#### Order Management:
- Create, retrieve, update, and delete orders.
- Retrieve order details by ID.

#### Product Item Management:
- Create, retrieve, update, and delete product items.
- Access product item information by ID.

### Technologies Used:

- **Programming Language:** Golang
- **Database Library:** database/sql
- **Logging:** log package
- **Context Handling:** context package

### Project Structure:

The project follows the Standard Project Layout for Golang applications:
- `cmd`: Contains main applications.
- `pkg`: Holds reusable library code.
- `model`: Stores data models.

### API Endpoints:

#### User Management:
- `POST /users`: Create a new user.
- `GET /users/:id`: Retrieve user information by ID.
- `PUT /users/:id`: Update user information by ID.
- `DELETE /users/:id`: Delete a user by ID.

#### Order Management:
- `POST /orders`: Create a new order.
- `GET /orders/:id`: Retrieve order information by ID.
- `PUT /orders/:id`: Update order information by ID.
- `DELETE /orders/:id`: Delete an order by ID.

#### Product Item Management:
- `POST /product-items`: Create a new product item.
- `GET /product-items/:id`: Retrieve product item information by ID.
- `PUT /product-items/:id`: Update product item information by ID.
- `DELETE /product-items/:id`: Delete a product item by ID.

### Database Structure:

#### User Model:
```sql
CREATE TABLE IF NOT EXISTS users (
    id          bigserial PRIMARY KEY,
    name        text NOT NULL,
    email       text NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
```

#### Order Model:
```sql
CREATE TABLE IF NOT EXISTS orders (
    order_id    bigserial PRIMARY KEY,
    item_id     bigserial REFERENCES product_item (id),
    user_id     bigserial REFERENCES users (id),
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    status      text NOT NULL DEFAULT 'Pending'
);
```

#### Product Item Model:
```sql
CREATE TABLE IF NOT EXISTS product_item (
    id          bigserial PRIMARY KEY,
    title       text NOT NULL,
    description text,
    price       int NOT NULL,
    created_at  timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated_at  timestamp(0) with time zone NOT NULL DEFAULT NOW()
);
```

This project adheres to REST principles and utilizes PostgreSQL for efficient data storage.
