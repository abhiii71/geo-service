# Geo-Service Prototype

A lightweight geolocation-based service finder API built with Golang, Gin, and a custom KD-Tree.

### Features
- Find nearby places within a radius
- Uses in-memory KD-Tree for fast geo search
- Implements the Haversine formula for accurate distance

### API Usage

**GET** `http://localhost:8080/get-nearby?lat=30.733&lon=76.78&radius=2`
**curl** `http://localhost:8080/get-nearby\?lat\=30.733\&lon\=76.78\&radius\=2 `
### Example Response
```json
[
  {
    "name": "City Hospital",
    "category": "hospital",
    "latitude": 30.7333,
    "longitude": 76.7794
  }
]
R