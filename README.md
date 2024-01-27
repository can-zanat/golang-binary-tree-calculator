# golang-technical-assignment
Golang Web Server for binary tree calculation

# You can use this curl command to test the API
```
curl --location 'localhost:83/postBinaryTreeSum' \
--header 'Content-Type: application/json' \
--data '{
    "tree": {
        "nodes": [
            {
                "id": "1",
                "left": "2",
                "right": "3",
                "value": 1
            },
            {
                "id": "3",
                "left": "6",
                "right": "7",
                "value": 3
            },
            {
                "id": "7",
                "left": null,
                "right": null,
                "value": 7
            },
            {
                "id": "6",
                "left": null,
                "right": null,
                "value": 6
            },
            {
                "id": "2",
                "left": "4",
                "right": "5",
                "value": 2
            },
            {
                "id": "5",
                "left": null,
                "right": null,
                "value": 5
            },
            {
                "id": "4",
                "left": null,
                "right": null,
                "value": 4
            }
        ],
        "root": "1"
    }
}'
```
