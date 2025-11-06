#!/bin/bash

# API Testing Script
BASE_URL="http://localhost:8080/api/v1"

echo "🚀 Starting API Testing..."

# Login dan dapatkan token
echo "📝 Logging in..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "email": "admin@company.com",
    "password": "admin123"
  }')

TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"token":"[^"]*' | cut -d'"' -f4)

if [ -z "$TOKEN" ]; then
  echo "❌ Login failed!"
  echo "Response: $LOGIN_RESPONSE"
  exit 1
fi

echo "✅ Login successful! Token obtained."

# Health check
echo "🏥 Testing health check..."
curl -s -X GET "$BASE_URL/health" | jq .

# Get user profile
echo "👤 Getting user profile..."
curl -s -X GET "$BASE_URL/users/me" \
  -H "Authorization: Bearer $TOKEN" | jq .

# List assets
echo "📦 Listing assets..."
curl -s -X GET "$BASE_URL/assets?limit=5" \
  -H "Authorization: Bearer $TOKEN" | jq .

# List tickets
echo "🎫 Listing tickets..."
curl -s -X GET "$BASE_URL/tickets?limit=5" \
  -H "Authorization: Bearer $TOKEN" | jq .

# List locations
echo "📍 Listing locations..."
curl -s -X GET "$BASE_URL/locations?limit=5" \
  -H "Authorization: Bearer $TOKEN" | jq .

echo "✅ API Testing completed!"