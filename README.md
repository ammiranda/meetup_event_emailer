## Meetup Emailer

### Overview

This application is designed to read JSON data for events (originally from Meetup.com) and inject it into a prompt to OpenAI to compose an email body. The email body is then sent to the configured recipients.

---

## Features

- Reads event data from a JSON file (e.g. Meetup.com export)
- Uses OpenAI's GPT model to generate a professional, friendly HTML email summarizing relevant events
- Focuses on technology and networking events for software engineers
- Sends the generated email to a configurable list of recipients via SMTP
- Fully configurable via environment variables
- Docker support for easy deployment

---

## How It Works

1. **Configuration**: Loads environment variables for API keys, SMTP credentials, and file paths.
2. **Data Loading**: Reads event data from a JSON file (see `data/events.json` for an example).
3. **Email Generation**: Passes the event data to OpenAI's API to generate an HTML email body tailored for networking and tech events.
4. **Email Sending**: Sends the generated email to the specified recipients using SMTP.

---

## Configuration

Set the following environment variables (e.g., in a `.env` file):

```
JSON_FILE_PATH=/app/data/events.json
OPENAI_API_KEY=your_openai_api_key
SMTP_HOST=smtp.example.com
SMTP_PORT=587
SMTP_USER=your_email@example.com
SMTP_PASSWORD=your_email_password
SMTP_RECEIPENTS=recipient1@example.com,recipient2@example.com
```

- `JSON_FILE_PATH`: Path to the events JSON file
- `OPENAI_API_KEY`: Your OpenAI API key
- `SMTP_HOST`, `SMTP_PORT`, `SMTP_USER`, `SMTP_PASSWORD`: SMTP server details
- `SMTP_RECEIPENTS`: Comma-separated list of recipient email addresses

---

## Usage

### Docker

Build and run with Docker:

```sh
docker build --no-cache -t amiranda/meetup-emailer .
docker run --rm -v "$(pwd)/data:/app/data" --env-file .env amiranda/meetup-emailer
```

### Locally (Go)

1. Install Go 1.24+
2. Set up your `.env` file as above
3. Run:

```sh
go run ./cmd/main.go
```

---

## Example Event Data

See `data/events.json` for a sample. Example entry:

```json
{
  "event_id": "309051093",
  "title": "Celebrating 1000 Members! - JULY MEETUP",
  "url": "https://www.meetup.com/austin-film-social/events/309051093/",
  "date": "2025-07-25T00:00:00-05:00",
  "date_display": "Thu, Jul 24 · 7:00 PM CDT",
  "group_name": "Austin Film Social",
  "rating": "4.7",
  "attendees": 0,
  "image_url": "https://secure.meetupstatic.com/photos/event/a/e/a/9/highres_529004713.jpeg"
}
```

---

## Dependencies

- [Go 1.24+](https://golang.org/)
- [github.com/openai/openai-go](https://github.com/openai/openai-go)
- [github.com/joho/godotenv](https://github.com/joho/godotenv)
- [github.com/tidwall/gjson](https://github.com/tidwall/gjson) (and related tidwall packages)

---

## Contributing

Contributions are welcome! Please open issues or submit pull requests for improvements, bug fixes, or new features.


