# API Contract - Inventory & Ticketing Management System
 
## 1. Overview
This document defines the REST API contract for the Inventory & Ticketing Management System. The API provides endpoints for managing assets, tickets, locations, and user authentication with role-based access control.
 
## 2. Base URL
```
https://api.inventory-system.com/v1
```
 
## 3. Authentication
All endpoints (except authentication) require a JWT token in the Authorization header:
```
Authorization: Bearer <jwt_token>
```
 
## 4. Common Response Format
All API responses follow this structure:
```json
{
  "success": true,
  "data": { /* response data */ },
  "message": "Success message",
  "timestamp": "2023-10-01T10:00:00Z"
}
```
 
Error responses:
```json
{
  "success": false,
  "error": {
    "code": "ERROR_CODE",
    "message": "Detailed error message"
  },
  "timestamp": "2023-10-01T10:00:00Z"
}
```
 
## 5. Authentication Endpoints
 
### POST /auth/login
Authenticate user and return JWT token.
 
**Request:**
```json
{
  "email": "admin@company.com",
  "password": "secret"
}
```
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": "usr-001",
      "name": "Admin User",
      "email": "admin@company.com",
      "role": "admin"
    }
  },
  "message": "Login successful"
}
```
 
## 6. Assets Endpoints
 
### GET /assets
List all assets with pagination and filtering.
 
**Query Parameters:**
- `page` (int): Page number (default: 1)
- `limit` (int): Items per page (default: 20)
- `location` (string): Filter by location ID
- `status` (string): Filter by status (available, booked, broken, repair)
- `type` (string): Filter by type (it, non_it)
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "items": [
      {
        "id": "ast-001",
        "uniqueId": "COMP-001",
        "name": "Dell Monitor",
        "comment": "Primary monitor",
        "detail": "24-inch LED monitor",
        "qty": 5,
        "brand": "Dell",
        "type": "it",
        "status": "available",
        "category": "Display",
        "locationId": "loc-001",
        "locationLabel": "Integrity — 2nd Floor",
        "createdAt": "2023-10-01T10:00:00Z",
        "updatedAt": "2023-10-01T10:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 20,
      "total": 45,
      "totalPages": 3
    }
  },
  "message": "Assets retrieved successfully"
}
```
 
### POST /assets
Create a new asset (Admin only).
 
**Request:**
```json
{
  "uniqueId": "COMP-002",
  "name": "HP Laptop",
  "comment": "Standard company laptop",
  "detail": "15-inch business laptop",
  "qty": 10,
  "brand": "HP",
  "type": "it",
  "status": "available",
  "category": "Computer",
  "locationId": "loc-002"
}
```
 
**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": "ast-002",
    "uniqueId": "COMP-002",
    "name": "HP Laptop",
    "comment": "Standard company laptop",
    "detail": "15-inch business laptop",
    "qty": 10,
    "brand": "HP",
    "type": "it",
    "status": "available",
    "category": "Computer",
    "locationId": "loc-002",
    "locationLabel": "Technology — 1st Floor",
    "createdAt": "2023-10-01T10:00:00Z",
    "updatedAt": "2023-10-01T10:00:00Z"
  },
  "message": "Asset created successfully"
}
```
 
### GET /assets/{id}
Get asset details by ID.
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "ast-001",
    "uniqueId": "COMP-001",
    "name": "Dell Monitor",
    "comment": "Primary monitor",
    "detail": "24-inch LED monitor",
    "qty": 5,
    "brand": "Dell",
    "type": "it",
    "status": "available",
    "category": "Display",
    "locationId": "loc-001",
    "locationLabel": "Integrity — 2nd Floor",
    "createdAt": "2023-10-01T10:00:00Z",
    "updatedAt": "2023-10-01T10:00:00Z"
  },
  "message": "Asset retrieved successfully"
}
```
 
### PUT /assets/{id}
Update asset details (Admin only).
 
**Request:**
```json
{
  "qty": 3,
  "status": "booked"
}
```
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "ast-001",
    "uniqueId": "COMP-001",
    "name": "Dell Monitor",
    "comment": "Primary monitor",
    "detail": "24-inch LED monitor",
    "qty": 3,
    "brand": "Dell",
    "type": "it",
    "status": "booked",
    "category": "Display",
    "locationId": "loc-001",
    "locationLabel": "Integrity — 2nd Floor",
    "createdAt": "2023-10-01T10:00:00Z",
    "updatedAt": "2023-10-01T10:30:00Z"
  },
  "message": "Asset updated successfully"
}
```
 
### DELETE /assets/{id}
Delete an asset (Admin only).
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": null,
  "message": "Asset deleted successfully"
}
```
 
## 7. Tickets Endpoints
 
### GET /tickets
List all tickets with filtering and pagination.
 
**Query Parameters:**
- `page` (int): Page number (default: 1)
- `limit` (int): Items per page (default: 20)
- `status` (string): Filter by status (open, in_progress, resolved, closed)
- `severity` (string): Filter by severity (low, medium, high, critical)
- `assignedTo` (string): Filter by assigned user ID
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "items": [
      {
        "id": "tck-001",
        "assetId": "ast-001",
        "category": "Maintenance",
        "severity": "high",
        "duration": "2 hours",
        "dueDate": "2023-10-15",
        "reporting": "usr-001",
        "assignedTo": "usr-002",
        "comment": "Monitor not working properly",
        "status": "open",
        "createdAt": "2023-10-01T10:00:00Z",
        "updatedAt": "2023-10-01T10:00:00Z"
      }
    ],
    "pagination": {
      "page": 1,
      "limit": 20,
      "total": 15,
      "totalPages": 1
    }
  },
  "message": "Tickets retrieved successfully"
}
```
 
### POST /tickets
Create a new ticket.
 
**Request:**
```json
{
  "assetId": "ast-001",
  "category": "Maintenance",
  "severity": "high",
  "duration": "2 hours",
  "dueDate": "2023-10-15",
  "comment": "Monitor not working properly"
}
```
 
**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": "tck-002",
    "assetId": "ast-001",
    "category": "Maintenance",
    "severity": "high",
    "duration": "2 hours",
    "dueDate": "2023-10-15",
    "reporting": "usr-001",
    "assignedTo": null,
    "comment": "Monitor not working properly",
    "status": "open",
    "createdAt": "2023-10-01T10:00:00Z",
    "updatedAt": "2023-10-01T10:00:00Z"
  },
  "message": "Ticket created successfully"
}
```
 
### GET /tickets/{id}
Get ticket details by ID.
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "tck-001",
    "assetId": "ast-001",
    "category": "Maintenance",
    "severity": "high",
    "duration": "2 hours",
    "dueDate": "2023-10-15",
    "reporting": "usr-001",
    "assignedTo": "usr-002",
    "comment": "Monitor not working properly",
    "status": "open",
    "createdAt": "2023-10-01T10:00:00Z",
    "updatedAt": "2023-10-01T10:00:00Z"
  },
  "message": "Ticket retrieved successfully"
}
```
 
### PUT /tickets/{id}
Update ticket details (Admin only).
 
**Request:**
```json
{
  "status": "in_progress",
  "assignedTo": "usr-003"
}
```
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "tck-001",
    "assetId": "ast-001",
    "category": "Maintenance",
    "severity": "high",
    "duration": "2 hours",
    "dueDate": "2023-10-15",
    "reporting": "usr-001",
    "assignedTo": "usr-003",
    "comment": "Monitor not working properly",
    "status": "in_progress",
    "createdAt": "2023-10-01T10:00:00Z",
    "updatedAt": "2023-10-01T10:30:00Z"
  },
  "message": "Ticket updated successfully"
}
```
 
### DELETE /tickets/{id}
Delete a ticket (Admin only).
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": null,
  "message": "Ticket deleted successfully"
}
```
 
## 8. Locations Endpoints
 
### GET /locations
List all locations.
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": [
    {
      "id": "loc-001",
      "name": "Integrity",
      "area": "2nd Floor",
      "description": "Located on the 2nd Floor",
      "capacity": 6,
      "createdAt": "2023-10-01T10:00:00Z",
      "updatedAt": "2023-10-01T10:00:00Z"
    },
    {
      "id": "loc-010",
      "name": "General",
      "area": "Outside Floor",
      "description": "General area",
      "capacity": 0,
      "createdAt": "2023-10-01T10:00:00Z",
      "updatedAt": "2023-10-01T10:00:00Z"
    }
  ],
  "message": "Locations retrieved successfully"
}
```
 
### GET /locations/{id}
Get location details by ID.
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "loc-001",
    "name": "Integrity",
    "area": "2nd Floor",
    "description": "Located on the 2nd Floor",
    "capacity": 6,
    "createdAt": "2023-10-01T10:00:00Z",
    "updatedAt": "2023-10-01T10:00:00Z"
  },
  "message": "Location retrieved successfully"
}
```
 
### POST /locations
Create a new location (Admin only).
 
**Request:**
```json
{
  "name": "New Conference Room",
  "area": "4th Floor",
  "description": "Large conference room with AV equipment",
  "capacity": 20
}
```
 
**Response (201 Created):**
```json
{
  "success": true,
  "data": {
    "id": "loc-011",
    "name": "New Conference Room",
    "area": "4th Floor",
    "description": "Large conference room with AV equipment",
    "capacity": 20,
    "createdAt": "2023-10-01T10:00:00Z",
    "updatedAt": "2023-10-01T10:00:00Z"
  },
  "message": "Location created successfully"
}
```
 
### PUT /locations/{id}
Update location details (Admin only).
 
**Request:**
```json
{
  "capacity": 25
}
```
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": {
    "id": "loc-011",
    "name": "New Conference Room",
    "area": "4th Floor",
    "description": "Large conference room with AV equipment",
    "capacity": 25,
    "createdAt": "2023-10-01T10:00:00Z",
    "updatedAt": "2023-10-01T10:30:00Z"
  },
  "message": "Location updated successfully"
}
```
 
### DELETE /locations/{id}
Delete a location (Admin only).
 
**Response (200 OK):**
```json
{
  "success": true,
  "data": null,
  "message": "Location deleted successfully"
}
```
 
## 9. Error Codes
 
| Code | Description |
|------|-------------|
| AUTH_REQUIRED | Authentication required |
| INVALID_TOKEN | Invalid or expired token |
| RESOURCE_NOT_FOUND | Requested resource not found |
| VALIDATION_ERROR | Request data validation failed |
| DUPLICATE_RESOURCE | Resource already exists |
 
 
## 10. HTTP Status Codes
 
| Code | Description |
|------|-------------|
| 200 | OK - Request successful |
| 201 | Created - Resource created successfully |
| 400 | Bad Request - Invalid request data |
| 401 | Unauthorized - Authentication required |