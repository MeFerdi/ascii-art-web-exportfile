# ASCII Art Web-Exportfile

**ASCII Art Web Exportfile** is a simple web application that generates ASCII art from user input. The application allows users to create ASCII art in various styles and download the result as a text file. It is built using Go for the backend, handling server-side rendering and file management.

## Features

- **Generate ASCII Art**: Convert user-provided text into ASCII art with different styles.
- **Download ASCII Art**: Save the generated ASCII art as a `.txt` file.
- **Responsive Design**: Works on various devices and screen sizes.

## Technologies Used

- **Go**: Backend server and template rendering.
- **HTML/CSS**: Frontend design and user interface.
- **ASCII Art Libraries**: Custom Go functions for generating ASCII art.

## Getting Started

To get started with this project, follow these instructions:

### Prerequisites

Ensure you have the following installed:
- [Go](https://golang.org/dl/) (version 1.18 or later)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)

### Installation

1. **Clone the Repository:**

```bash
git clone https://learn.zone01kisumu.ke/git/jowala/ascii-art-web-stylize.git
cd ascii-art-web
```

### Running the Application

3. **Run the Application:**

    Start the server with:

    ```bash
    go run .
    ```

    By default, the server will run on `http://localhost:8080`.

4. **Access the Application:**

    Open your web browser and navigate to `http://localhost:8080` to access the ASCII Art Web application.

## Usage

1. **Generate ASCII Art:**

- Enter the text you want to convert into ASCII art in the provided text area.
- Select a banner style (Standard, Shadow, Thinkertoy).
- Click the "Generate" button to display the ASCII art.

2. **Download ASCII Art:**

- After generating the ASCII art, a "Download ASCII Art" button will appear.
- Click the button to download the ASCII art as a `.txt` file.

## Project Structure

- `main.go`: Entry point of the application.
- `handler/`: Contains handlers for different routes (`HomeHandler`, `AsciiArtHandler`, `DownloadAsciiArtHandler`).
- `templates/`: HTML templates and CSS files.
- `ascii/`: Package with functions for generating ASCII art.

## Code Documentation

- **`HomeHandler`**: Serves the home page (index.html) for GET requests.
- **`AsciiArtHandler`**: Handles ASCII art generation based on user input (POST requests).
- **`DownloadAsciiArtHandler`**: Allows users to download the generated ASCII art as a `.txt` file.

## Contributing

Contributions are welcome! If you have suggestions or improvements, please fork the repository and submit a pull request.

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Make your changes and commit (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Create a new Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
