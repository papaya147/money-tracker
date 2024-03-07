CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS profile(
    id TEXT NOT NULL PRIMARY KEY,
    email TEXT NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
CREATE TABLE IF NOT EXISTS expenditureCategory(
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    profileId TEXT NOT NULL,
    name TEXT NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_name UNIQUE (profileId, name)
);
ALTER TABLE expenditureCategory
ADD FOREIGN KEY (profileId) REFERENCES profile (id);
CREATE TABLE IF NOT EXISTS expenditure(
    id UUID NOT NULL DEFAULT uuid_generate_v4() PRIMARY KEY,
    profileId TEXT NOT NULL,
    paisa INT NOT NULL,
    categoryId UUID NOT NULL,
    createdAt TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updatedAt TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
ALTER TABLE expenditure
ADD FOREIGN KEY (profileId) REFERENCES profile (id);
ALTER TABLE expenditure
ADD FOREIGN KEY (categoryId) REFERENCES expenditureCategory (id);