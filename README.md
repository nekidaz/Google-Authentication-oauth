# Google-Authentication-oauth

This repository contains a Go application that implements Google Authentication using OAuth, wrapped in a Docker container using Docker Compose.

## Overview

This application allows users to log in with their Google accounts using OAuth authentication. It provides two main functionalities:

1. **Login**: When a user accesses the login endpoint, they are redirected to Google's authentication page, where they can authorize the application to access their Google account information.

2. **Callback**: After the user authorizes the application, Google redirects them back to the callback endpoint. The application then retrieves the user's access token and uses it to fetch the user's information from the Google API. The user's email, name, and email verification status are extracted from the response and stored in the application's database. The user's access token is also stored in a secure HTTP-only cookie named "Authorization".

## Prerequisites

Before running this application, make sure you have the following installed on your machine:

- Docker: [Installation Guide](https://docs.docker.com/get-docker/)
- Docker Compose: [Installation Guide](https://docs.docker.com/compose/install/)

## Getting Started

1. Clone the repository:
   ```shell
   git clone https://github.com/nekidaz/Google-Authentication-oauth.git
2. Navigate to the project directory:
   ```shell
   cd Google-Authentication-oauth
3. Configure the application:
  - Replace the placeholder values in .env.app and .env.postgres file with your own Google OAuth client credentials and database configuration.
4. Build and run the Docker containers:
   ```shell
   docker-compose up --build
5.Access the application in your browser at http://localhost:8080 or the specified port.
