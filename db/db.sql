DROP TABLE IF EXISTS Customer CASCADE;
DROP TABLE IF EXISTS Item CASCADE;
DROP TABLE IF EXISTS Purchase_Order CASCADE;

CREATE TABLE Customer
(
    cust_id SERIAL PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    address VARCHAR(50) NOT NULL
);

CREATE TABLE Item
(
    item_id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    stock INTEGER NOT NULL,
    price INTEGER NOT NULL
);

CREATE TABLE Purchase_Order 
(
    purchase_order_id SERIAL PRIMARY KEY,
    cust_id INTEGER NOT NULL,
    item_id INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    dispatched BOOLEAN NOT NULL DEFAULT FALSE
);
-- All cart items are from the PurchaseOrder table mapped to a customer id with his/her item
-- SELECT * FROM PurchaseOrder WHERE customer.cust_id = PurchaseOrder.cust_id

-- Adding foreign keys after table generation
ALTER TABLE Purchase_Order ADD FOREIGN KEY (cust_id) REFERENCES Customer (cust_id) ON DELETE CASCADE;
ALTER TABLE Purchase_Order ADD FOREIGN KEY (item_id) REFERENCES Item (item_id) ON DELETE CASCADE;

-- Adding test data
INSERT INTO Item (name, stock, price) VALUES ('Bag', 10, 20);
INSERT INTO Item (name, stock, price) VALUES ('Shirts', 5, 60);
INSERT INTO Item (name, stock, price) VALUES ('Suits', 2, 300);
INSERT INTO Item (name, stock, price) VALUES ('Clock', 4, 70);
INSERT INTO Item (name, stock, price) VALUES ('Shoes', 1, 120);
INSERT INTO Item (name, stock, price) VALUES ('Ties', 8, 20);

INSERT INTO Customer (first_name, last_name, address) VALUES ('John', 'Doe', 'Jakarta');
INSERT INTO Customer (first_name, last_name, address) VALUES ('Bob', 'Williams', 'Medan');

-- Adding Index
-- Note: better to load all data and then create the index. Use Explain Analyse to detect bottlenecks in query.
-- Add index on fields/columns that are commonly used in the 'WHERE' or 'Group By' clauses
-- If indexing joins, index the field on the left hand side of the assignment
-- CREATE INDEX item_idx ON Item (item_id);