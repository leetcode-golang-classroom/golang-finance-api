CREATE TABLE IF NOT EXISTS categories (
  id TEXT NOT NULL PRIMARY KEY DEFAULT nanoid(),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  user_id TEXT REFERENCES users(id),
  title TEXT NOT NULL
);
CREATE OR REPLACE TRIGGER categories_updated_at
  BEFORE UPDATE ON categories
  FOR EACH ROW
  EXECUTE procedure moddatetime (updated_at);

CREATE TYPE transaction_type AS ENUM ('income', 'expense');

CREATE TABLE IF NOT EXISTS transactions (
  id TEXT NOT NULL PRIMARY KEY DEFAULT nanoid(),
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  updated_at TIMESTAMP,
  user_id TEXT NOT NULL REFERENCES users(id),
  categorie_id TEXT NOT NULL REFERENCES categories(id),
  title TEXT NOT NULL,
  currency TEXT NOT NULL,
  amount NUMERIC(12,2) NOT NULL,
  tx_type transaction_type NOT NULL
);

CREATE OR REPLACE TRIGGER transactions_updated_at
  BEFORE UPDATE ON transactions
  FOR EACH ROW
  EXECUTE procedure moddatetime (updated_at);