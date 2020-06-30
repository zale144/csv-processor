CREATE TABLE IF NOT EXISTS public."user"(
     id INTEGER UNIQUE NOT NULL,
     first_name VARCHAR(255) NOT NULL,
     last_name VARCHAR(255) NOT NULL,
     email VARCHAR(255) NOT NULL,
     phone VARCHAR(44) NOT NULL
);
