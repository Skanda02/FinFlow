# Database Setup Instructions

## Quick Setup (Recommended)
Run this as PostgreSQL superuser first to create user and database:
```sh
psql -U postgres
```

```sql
-- Create user
CREATE USER finflow WITH PASSWORD 'finflow';

-- Create database
CREATE DATABASE finflowdb OWNER finflow;

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE finflowdb TO finflow;

-- Connect to the new database
\c finflowdb

-- Grant schema privileges (required for PostgreSQL 15+)
GRANT ALL ON SCHEMA public TO finflow;
GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO finflow;
GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO finflow;

-- To quit use \q
```

To setup the database

```bash
psql -U finflow -d finflowdb -h localhost -f setup_db.sql
```

## Manual Setup (Alternative)

If you prefer to run commands manually:

1. Create user and database:
```bash
sudo -u postgres psql -c "CREATE USER finflow WITH PASSWORD 'finflow';"
sudo -u postgres psql -c "CREATE DATABASE finflowdb OWNER finflow;"
sudo -u postgres psql -c "GRANT ALL PRIVILEGES ON DATABASE finflowdb TO finflow;"
```

2. Create tables:
```bash
psql -U finflow -d finflowdb -h localhost -f setup_db.sql
```

## Quick Reset (Drop and Recreate Tables)

If you need to reset the database:

```bash
psql -U finflow -d finflowdb -h localhost
```

```sql
DROP TABLE IF EXISTS transactions CASCADE;
DROP TABLE IF EXISTS users CASCADE;
```

Then run the table creation commands from `setup_db.sql` again.

## Environment Variables

Make sure your `.env` file has:

```
DB_URL=postgres://finflow:finflow@localhost:5432/finflowdb?sslmode=disable
```
