﻿# Go Twitter API

## Introduction

This project is designed to help users understand how to authenticate with the Twitter API using OAuth and perform two key operations: posting a tweet and deleting a tweet programmatically. It demonstrates how to work with REST APIs, secure credentials using environment variables, and handle API responses and errors.

## Table of Contents

- Introduction
- Project Structure
- Features
- Setup Instructions
  - Create a Twitter Developer Account
  - Generate API Keys
  - Store Environment Variables
  - Run the Program
- API Implementation
  - Post a Tweet
  - Delete a Tweet
- Error Handling
- Testing API Endpoints
  - Example Responses
- Technologies Used

## Project Structure

```bash
.
├── main.go          # Main Go file that interacts with Twitter API
├── .env             # Stores environment variables (should not be committed)
├── go.mod           # Go module file
├── go.sum           # Dependency management file
└── README.md        # Project documentation
```

## Features

- Post a tweet using Twitter's POST `statuses/update` endpoint.
- Delete a tweet using Twitter's POST `statuses/destroy` endpoint.
- Secure OAuth 1.0a authentication to interact with the Twitter API.
- Comprehensive error handling for various API issues.

## Setup Instructions

### 1. Create a Twitter Developer Account

1. Visit the [Twitter Developer Platform](https://developer.twitter.com/).
2. Sign up for a developer account or log in with your existing Twitter account.
3. Create a new project and app after your account is approved.
4. You will get access to your API keys and tokens.

### 2. Generate API Keys

After creating an app, navigate to `Projects & Apps` → `Your App` → `Keys and Tokens`.
Generate the following keys:

- API Key (also known as `TWITTER_CONSUMER_KEY`)
- API Secret Key (also known as `TWITTER_CONSUMER_SECRET`)
- `Access Token`
- `Access Token Secret`

These credentials are needed to authenticate your Go application with Twitter's API.

### 3. Store Environment Variables

#### Locally

To securely store your API keys and tokens for local development, use a `.env` file.

Create a `.env` file in the root of your project directory:

```bash
touch .env
```

Add the following lines to the `.env` file, replacing the values with your actual keys and tokens:

```bash
TWITTER_CONSUMER_KEY=your-consumer-key
TWITTER_CONSUMER_SECRET=your-consumer-secret
TWITTER_ACCESS_TOKEN=your-access-token
TWITTER_ACCESS_SECRET=your-access-secret
```

Install the Go package `godotenv` to load these variables:

```bash
go get github.com/joho/godotenv
```

Then, load the environment variables in your Go code with the following:

```go
err := godotenv.Load()
if err != nil {
    log.Fatal("Error loading .env file")
}
```

#### Storing Secrets in GitHub

When deploying the project or running it in production, you can store your API keys and tokens securely in GitHub Secrets:

1. Navigate to your repository on GitHub.
2. Go to `Settings` → `Secrets and variables` → `Actions` → `New repository secret`.
3. Add the following secrets:
    - TWITTER_CONSUMER_KEY
    - TWITTER_CONSUMER_SECRET
    - TWITTER_ACCESS_TOKEN
    - TWITTER_ACCESS_SECRET

These secrets can be accessed securely in your GitHub Actions workflow for CI/CD.

### 4. Run the Program

1. Clone the repository:

    ```bash
    git clone https://github.com/your-username/twitter-api-go.git
    cd twitter-api-go
    ```

2. Install dependencies: Run this command to download the required Go modules:

    ```bash
    go mod tidy
    ```

3. Set up environment variables: If not using `.env`, ensure your environment variables are set properly:

    ```bash
    export TWITTER_CONSUMER_KEY=your-consumer-key
    export TWITTER_CONSUMER_SECRET=your-consumer-secret
    export TWITTER_ACCESS_TOKEN=your-access-token
    export TWITTER_ACCESS_SECRET=your-access-secret
    ```

4. Run the application: To post a tweet and then delete it, run:

    ```bash
    go run main.go
    ```

The program will post a tweet with `"Hello from Twitter API using Go!"` and delete it afterward.

## API Implementation

### Post a Tweet

- **Endpoint**: POST `statuses/update`
- **Description**: Posts a tweet to your Twitter account.
- **Implementation**: The tweet content is sent as a form parameter in the HTTP POST request. OAuth 1.0a authentication is used for secure access.
- **Example request**:

```go
postTweet("Hello from Twitter API using Go!")
```

### Delete a Tweet

- **Endpoint**: POST `statuses/destroy/:id`
- **Description**: Deletes a tweet based on the provided tweet ID.
- **Implementation**: The tweet ID is passed as a parameter in the URL to delete the corresponding tweet.
- **Example request**:

```go
deleteTweet(tweetID)
```

## Error Handling

The program has robust error handling in place. It captures the following errors:

- **Invalid credentials**: If API keys or tokens are incorrect, an error is logged.
- **Rate limiting**: If you hit Twitter's rate limit, an appropriate error message is displayed.
- **Invalid tweet ID**: Trying to delete a non-existent tweet will trigger an error message.
- **Network issues**: All network errors (e.g., failed requests) are handled gracefully and logged.

## Testing API Endpoints

To manually test the Twitter API, you can run the program in two parts:

- **Post a Tweet**: When you run the program, the tweet will be posted using the `postTweet` function.
- **Delete a Tweet**: After posting a tweet, its ID is passed to the `deleteTweet` function to delete it.

### Example Responses

**Post Tweet Response**:

```json
{
  "created_at": "Fri Oct 13 16:29:29 +0000 2024",
  "id": 145334567890123456,
  "text": "Hello from Twitter API using Go!"
}
```

**Delete Tweet Response**:

```json
{
  "created_at": "Fri Oct 13 16:29:29 +0000 2024",
  "id": 145334567890123456,
  "text": "Hello from Twitter API using Go!",
  "deleted": true
}
```

## Technologies Used

- **Go**: The main programming language used.
- **OAuth 1.0a**: For authenticating API requests.
- **Twitter API**: Used for posting and deleting tweets.
- **Environment Variables**: API keys and tokens are stored securely in a `.env` file.
- **GitHub Secrets**: Used for secure storage of credentials in production.
