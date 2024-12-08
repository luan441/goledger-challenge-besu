CREATE TABLE IF NOT EXISTS public.blockchain_value(
  id serial PRIMARY KEY,
  value bigint NOT NULL,
  created_at timestamp DEFAULT current_timestamp
);
