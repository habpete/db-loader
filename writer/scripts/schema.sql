CREATE DATABASE my_services

CREATE TABLE profiles (
    id serial PRIMARY KEY,
    user_id bigint NOT NULL,
    first_name nvarchar(40) NOT NULL,
    second_name nvarchar(50) NOT NULL,
    phone_number integer,
    email nvarchar(100)
)

CREATE TABLE reviews_bind (
    user_id bigint PRIMARY KEY,
    reviews_ids serial[]
)

CREATE TABLE reviews (
    id serial PRIMARY KEY,
    comment text
)

CREATE TABLE scores_bind (
    product_id bigint,
    user_id bigint,
    scores_id serial[]
)

CREATE TABLE scores (
    id serial PRIMARY KEY,
    score bigint
)

CREATE TABLE products (
    id serial PRIMARY KEY,
    product_name nvarchar(100) NOT NULL
)