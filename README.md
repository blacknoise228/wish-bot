# Wish Bot
=================

Wish Bot is a Telegram bot designed to help users create, manage and share their wishes with friends.

## Features

* Create and manage wishes with descriptions, links, and status updates
* Share wishes with friends and view their wishes in return
* User-friendly interface for easy navigation and wish management
* Add stores, admins, and products to add to wishes

## Technologies Used

* Go programming language
* Telegram Bot API
* PostgreSQL database

## Getting Started

1. Clone the repository: `git clone https://github.com/your-username/wish-bot.git`
2. Install dependencies: `go get -u ./...`
3. Create a new PostgreSQL database and update the `config.yaml` file with your database credentials
4. Run the bot: `go run main.go`

## Configuration

The bot uses the `config.yaml` file to store configuration settings. You can update the file to customize the bot's behavior.

## API Documentation

API documentation is available in the `internal/api/telegram` directory.

## License

This project is licensed under the MIT license.
