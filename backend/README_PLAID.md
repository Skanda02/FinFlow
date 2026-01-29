# Plaid Integration Setup Guide

## Overview
FinFlow uses Plaid to automatically fetch bank transactions. The application will run without Plaid configured, but bank integration features will be unavailable.

## Getting Plaid Credentials

1. Sign up for a Plaid account at https://dashboard.plaid.com/signup
2. After verification, you'll get access to your dashboard
3. Navigate to **Team Settings** > **Keys**
4. Copy your credentials:
   - `client_id`
   - `secret` (for sandbox environment)

## Environment Variables

Create a `.env` file in the `backend` directory with the following variables:

```bash
# Plaid Configuration
PLAID_CLIENT_ID=your_client_id_here
PLAID_SECRET=your_secret_here
PLAID_ENVIRONMENT=sandbox  # Options: sandbox, development, production

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=finflow
DB_PASSWORD=your_password
DB_NAME=finflowdb

# JWT Configuration
JWT_SECRET=your_jwt_secret_here
```

## Plaid Environments

- **sandbox**: For testing, uses mock data, no real bank credentials needed
- **development**: For integration testing with real (non-production) credentials
- **production**: For live application with real bank data

Start with `sandbox` for development.

## Running Without Plaid

If Plaid credentials are not configured:
- The application will start normally
- A warning will be logged: "Plaid credentials not configured. Bank integration features will be unavailable."
- Bank endpoints (`/bank/*`) will return 503 Service Unavailable
- Manual transaction entry (`/add_income`, `/add_expense`) will work normally

## Testing Plaid Integration

### 1. Link a Bank Account (Frontend)

In sandbox mode, use these test credentials when linking:
- **Username**: `user_good`
- **Password**: `pass_good`
- **Institution**: Any bank from the list

### 2. API Flow

```bash
# 1. Get link token (implement this endpoint to call Plaid's LinkTokenCreate)
# Frontend uses Plaid Link with this token

# 2. Exchange public token (after user completes Plaid Link)
POST /bank/link
{
  "public_token": "public-sandbox-xxx"
}

# 3. Sync transactions
POST /bank/sync

# 4. View connected banks
GET /bank/connections

# 5. Unlink a bank
POST /bank/unlink
{
  "connection_id": 1
}
```

### 3. Sandbox Test Scenarios

Plaid provides special test accounts in sandbox:

- `user_good` / `pass_good`: Successful account with transactions
- `user_bad` / `pass_bad`: Invalid credentials error
- Custom accounts: See https://plaid.com/docs/sandbox/test-credentials/

## Implementation Details

### PlaidClient Interface
The system uses a `PlaidClient` interface, making it easy to:
- Swap between real and mock implementations
- Test without real Plaid credentials
- Extend functionality

### Transaction Sync
- Fetches last 30 days of transactions by default
- Prevents duplicate imports using `bank_transaction_id`
- Updates `last_synced_at` timestamp
- Handles debit (expenses) and credit (income) transactions

### Error Handling
- Network errors: Returns 500 with message
- Invalid tokens: Returns 400 with error details
- Service unavailable: Returns 503 when Plaid not configured

## Security Considerations

1. **Never commit `.env` file** - Add it to `.gitignore`
2. **Access tokens are encrypted** - Stored securely in database
3. **Use HTTPS in production** - Plaid requires secure connections
4. **Rotate secrets regularly** - Especially when moving to production
5. **Limit webhook IPs** - If implementing Plaid webhooks

## Migration to Production

Before going live:

1. Apply for Production access in Plaid Dashboard
2. Complete Plaid's compliance requirements
3. Update `PLAID_ENVIRONMENT=production`
4. Use production secret keys
5. Test thoroughly with real bank accounts
6. Implement proper error monitoring

## Troubleshooting

### "Bank integration is not configured"
- Check environment variables are set
- Verify `.env` file is in `backend/` directory
- Restart the application after setting env vars

### "Invalid credentials" when linking
- In sandbox, use `user_good` / `pass_good`
- Check PLAID_ENVIRONMENT matches your use case

### Transactions not syncing
- Check `last_synced_at` in `bank_connections` table
- Verify access token is valid
- Check application logs for API errors

### Rate Limiting
- Sandbox: 100 requests/minute
- Development: Higher limits (varies)
- Production: Enterprise limits (varies)

## Useful Links

- Plaid Quickstart: https://plaid.com/docs/quickstart/
- API Reference: https://plaid.com/docs/api/
- Go SDK: https://github.com/plaid/plaid-go
- Dashboard: https://dashboard.plaid.com/
