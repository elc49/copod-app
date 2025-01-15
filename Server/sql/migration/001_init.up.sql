CREATE TABLE IF NOT EXISTS users(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  firstname TEXT NOT NULL DEFAULT '',
  lastname TEXT NOT NULL DEFAULT '',
  email TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_user_email ON users(email);

CREATE TABLE IF NOT EXISTS support_docs(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  url TEXT NOT NULL,
  verification TEXT NOT NULL DEFAULT 'ONBOARDING',
  email TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_support_docs_email ON support_docs(email);

CREATE TABLE IF NOT EXISTS title_deeds(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  url TEXT NOT NULL,
  title TEXT,
  verification TEXT NOT NULL DEFAULT 'ONBOARDING',
  email TEXT NOT NULL,
  support_doc_id UUID NOT NULL REFERENCES support_docs(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_title_email ON title_deeds(email);

CREATE TABLE IF NOT EXISTS display_pictures(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  url TEXT NOT NULL,
  verification TEXT NOT NULL DEFAULT 'ONBOARDING',
  email TEXT NOT NULL,
  support_doc_id UUID NOT NULL REFERENCES support_docs(id) ON DELETE CASCADE,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS onboardings(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  title_id UUID NOT NULL REFERENCES title_deeds(id) ON DELETE CASCADE,
  support_doc_id UUID NOT NULL REFERENCES support_docs(id) ON DELETE CASCADE,
  display_picture_id UUID NOT NULL REFERENCES display_pictures(id) ON DELETE CASCADE,
  email TEXT NOT NULL,
  verification TEXT NOT NULL DEFAULT 'ONBOARDING',
  payment_status TEXT,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_onboarding_email ON onboardings(email);

CREATE TABLE IF NOT EXISTS payments(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email TEXT NOT NULL,
  amount INTEGER NOT NULL,
  currency TEXT NOT NULL,
  reason TEXT NOT NULL,
  status TEXT NOT NULL,
  reference_id TEXT NOT NULL,
  onboarding_id UUID REFERENCES onboardings(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_reference_id ON payments(reference_id);
CREATE INDEX IF NOT EXISTS idx_onboarding_id ON payments(onboarding_id);

CREATE TABLE IF NOT EXISTS early_signups(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  email TEXT NOT NULL,
  onboarded TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_early_signup_email ON early_signups(email);
