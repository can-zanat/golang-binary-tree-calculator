# golang-binay-tree-calculator
Golang API for binary tree calculation

![explanationOfHowToCalculate](https://github.com/can-zanat/golang-binary-tree-calculator/assets/65563976/a9569fec-08b6-4925-bfb6-ded040df47b6)
![postmanRequestAndResponse](https://github.com/can-zanat/golang-binary-tree-calculator/assets/65563976/dbe8bd63-f866-43a4-bc66-4f10162d6021)

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
