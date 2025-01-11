# Swiper

**Swiper** is a prototype backend service for swipe-based dating applications similar to Tinder or Bumble. It provides essential functionalities such as user registration, login, and swipe mechanics, including daily swipe limits and match-making logic.

## Features

- **User Registration**: Allows new users to create an account with a name, username, password, age, gender, and location.

- **User Login**: Authenticates existing users using their username and password.

- **Swipe Functionality**:
  - Users can swipe right (like) or left (dislike) on other user profiles.
  - Each user is limited to swiping up to 10 profiles per day.
  - Ensures that no profile is shown to the same user that swiped them more than once per day.

- **Match-Making**: Detects mutual likes between users and records matches.

- **Enable Premium**: Enable premium to bypass swipe limit and add `verified` status to the profile.

## Getting Started

### Prerequisites

- **Go**: Ensure that Go is installed on your machine. You can download it from the [official website](https://golang.org/dl/).

- **Redis**: The application uses Redis for caching swipe data. Install Redis by following the instructions on the [official website](https://redis.io/download).

- **PostgreSQL**: A PostgreSQL database is used to store user information and swipe history. Download and install it from the [official website](https://www.postgresql.org/download/).

- **Migrate**: go-migrate/migrate for database migration. Download and install it from the repo [migrate repository](https://github.com/golang-migrate/migrate).

### Installation

1. **Clone the Repository**:
    ```bash
    git clone https://github.com/rahadianir/swiper.git
    cd swiper
    ```
2. **Setup Environment Variable**:
    
    Copy example environment file with
    ```bash
    cp .env.example .env
    ```
    and update the values accordingly.
3. **Install Dependencies**
    ```bash
    go mod tidy
    ```
4. **Run Database Migration**
    ```bash
    migrate -path migration -database $DB_URI up
    ```
5. **Start The Service**
    ```bash
    go run main.go
    ```
## API
API documentations and collections can be viewed [here](https://github.com/rahadianir/swiper/docs).