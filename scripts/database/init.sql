CREATE TABLE IF NOT EXISTS public.blockchain_value(
  id serial PRIMARY KEY,
  value TEXT NOT NULL,
  created_at timestamp DEFAULT current_timestamp
);
