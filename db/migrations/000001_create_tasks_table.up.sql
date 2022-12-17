CREATE TABLE IF NOT EXISTS tasks(
   id serial PRIMARY KEY,
   name VARCHAR (50),
   description VARCHAR (300),
   end_date DATE,
   status VARCHAR (50), 
   priority VARCHAR (50),
   created_at TIMESTAMP,
   updated_at TIMESTAMP
);