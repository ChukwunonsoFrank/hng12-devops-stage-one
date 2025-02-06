# HNG12 DevOps Task(Stage 1)

👩‍💻 HNG12 DevOps Stage 1 - An API that takes a number and returns interesting mathematical properties about it, along with a fun fact.

## Getting Started

#### Clone the repo:

```bash
git clone https://github.com/ChukwunonsoFrank/hng12-devops-stage-one.git
cd hng12-devops-stage-one
```

#### Set environment variables:

```bash
cp .env.example .env
```

####  Start the server:

```bash
go build && ./hng12-devops-stage-one
```

## Request/Response Formats

#### JSON Response Format (200 OK)

```bash
GET /api/classify-number?number=123 HTTP/1.1
```
```json
    {
        "number": 123,
        "is_prime": false,
        "is_perfect": false,
        "properties": ["odd"],
        "digit_sum": 6,
        "fun_fact": "123 is the 10th Lucas number."
    }
```

#### JSON Response Format (400 Bad Request)

```bash
GET /api/classify-number?number=abc HTTP/1.1
```
```json
    {
        "number": "alphabet",
        "error": true
    }
```

## Sample Usage

As shown below, we are making a GET request with cURL with jq(for pretty printing JSON responses) but you can make use of any API testing tool available to you.

#### Sample GET request(200 OK)
![Sample GET request 200 OK](https://github.com/ChukwunonsoFrank/hng12-devops-stage-one/blob/main/assets/sample-200-request.png)

#### Sample GET request(400 Bad Request)
![Sample GET request 400 Bad Request](https://github.com/ChukwunonsoFrank/hng12-devops-stage-one/blob/main/assets/sample-400-request.png)
