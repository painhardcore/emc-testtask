# emc-testtask

## GET /car/:serialnumber
Get car information by serialnumber

## POST /car
* Content-Type: "application/json"
  
Add car information to the system
* Example
```json
{
    "serialNumber": 1001,
    "ownerName": "Valya Peshehodka",
    "modelYear": 1999,
    "code": "code",
    "vehicleCode": "vcode",
    "engine": {
        "capacity": 2000,
        "numCylinders": 12,
        "maxRpm": 100,
        "manufacturerCode": "mcode"
    },
    "fuelFigures": {
        "speed": 100,
        "mpg": 100,
        "usageDescription": "desc"
    },
    "performanceFigures": {
        "octaneRating": 99,
        "acceleration": {
            "mph": 23,
            "seconds": 1
        }
    },
    "manufacturer": "Subaru",
    "model": "Rx100",
    "activationCode": "lucky"
}
```