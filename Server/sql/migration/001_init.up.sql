CREATE TABLE IF NOT EXISTS users(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  firstname TEXT NOT NULL DEFAULT '',
  lastname TEXT NOT NULL DEFAULT '',
  govt_id TEXT NOT NULL,
  email TEXT NOT NULL,
  wallet_address TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON users(email);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_wallet ON users(wallet_address);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_govt_id ON users(govt_id);

CREATE TABLE IF NOT EXISTS titles(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title TEXT NOT NULL,
  verification TEXT NOT NULL DEFAULT 'ONBOARDING',
  email TEXT NOT NULL,
  wallet_address TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_title_email ON titles(email);

CREATE TABLE IF NOT EXISTS support_docs(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  govt_id TEXT NOT NULL,
  verification TEXT NOT NULL DEFAULT 'ONBOARDING',
  email TEXT NOT NULL,
  wallet_address TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_support_docs_email ON support_docs(email);

CREATE TABLE IF NOT EXISTS payments(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email TEXT NOT NULL,
  amount INTEGER NOT NULL,
  currency TEXT NOT NULL,
  reason TEXT NOT NULL,
  status TEXT NOT NULL,
  reference_id TEXT NOT NULL,
  title_id UUID REFERENCES titles(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_reference_id ON payments(reference_id);
CREATE INDEX IF NOT EXISTS idx_title_id ON payments(title_id);
