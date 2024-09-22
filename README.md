# pls find my... 🕵🙏

A simple web application to search for movies and find their streaming availability.
Data is mocked and the app is for learning purposes of GO and HTMX. For really useful app I recommend using justwatch.com or similar.

## Tech Stack

- Backend: Go with Gin web framework
- Database: SQLite
- Frontend: HTML, CSS, and HTMX for dynamic content loading
- Templating: Templ for Go HTML templating

## Setup and Running

1. Ensure you have Go installed on your system.
2. Clone this repository.
3. Install dependencies:
   ```
   go mod tidy
   ```
4. Create SQLite database file named `movies.sqlite` in the project root and run commands from `create-movies.sql` to create the tables.
5. Run the application:
   ```
   go run main.go
   ```
   Alternatively, you can use Air for hot reloading during development:
   ```
   air
   ```
6. Open your browser and navigate to `http://localhost:8080`

## Project Structure

- `main.go`: Main application file containing route handlers and database initialization
- `models/`: Data models for the application
- `views/`: Templ templates for rendering HTML
- `static/`: Static assets like CSS and images

## API Endpoints

- `GET /`: Renders the main search page
- `GET /search`: Handles movie search requests
- `GET /streamings/:id`: Retrieves streaming information for a specific movie

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
