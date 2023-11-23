SET statement_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = ON;
SET check_function_bodies = FALSE;
SET client_min_messages = WARNING;
SET search_path = public, extensions;
SET default_tablespace = '';
SET default_with_oids = FALSE;

CREATE TABLE stores (
  id SERIAL PRIMARY KEY, 
  name TEXT, 
  is_available BOOLEAN
);

CREATE TABLE products (
  id SERIAL PRIMARY KEY, 
  name TEXT, 
  size TEXT, 
  code TEXT UNIQUE, 
  quantity INTEGER NOT NULL, 
  store_id INTEGER REFERENCES stores(id)
);

CREATE INDEX idx_products_code ON products (code);
CREATE INDEX idx_products_quantity ON products (quantity);
CREATE INDEX idx_stores_name ON stores (name);
