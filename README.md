# Plantera

Plantera is a plant care tracking application designed to help you keep track of watering, fertilizing, and repotting your plants. This application is built entirely in Go, utilizing `templ` for templating, `htmx` for handling dynamic updates, `Tailwind CSS` for styling, and a touch of `hyperscript` for additional interactivity.

## Features

- **Track Plant Care**: Easily keep track of when each plant was last watered, fertilized, and repotted.
- **Add New Plants**: Add new plants with information on watering, fertilizing, and repotting schedules.
- **Dynamic Updates**: Enjoy a dynamic user experience with `htmx`, allowing for seamless content updates without full page reloads.
- **[Firebase](https://firebase.google.com/) Storage Integration**: Upload and store images of your plants in Firebase Storage.

## Tech Stack

- **[Go](https://golang.org/)**: The core programming language used for the entire application.
- **[templ](https://github.com/a-h/templ)**: Used for server-side templating.
- **[htmx](https://htmx.org/)**: Handles dynamic updates and interactions.
- **[Tailwind CSS](https://tailwindcss.com/)**: Provides a utility-first CSS framework for styling the application.
- **[Hyperscript](https://hyperscript.org/)**: Adds additional interactivity to the application.

## Getting Started

### Prerequisites

- Go 1.18+
- Node.js (for Tailwind CSS)
- Firebase account and a project setup with Firebase Storage

### Installation

1. **Clone the repository:**

    ```bash
    git clone https://github.com/davidtemelkov/plantera-go.git
    cd plantera-go
    ```

2. **Install Go dependencies:**

    ```bash
    go mod tidy
    ```

3. **Set up Firebase:**

    - Place your `serviceAccountKey.json` file in the root of the project.
    - Set up your environment variables in a `.env` file:

      ```env
      FIREBASE_BUCKET_NAME=your-firebase-bucket-name
      FIREBASE_URL=your-firebase-url
      ```

### Running the Application

1. **Run Makefile:**

    ```bash
    make
    ```
    
3. **Access the application:**

    Open your browser and navigate to `http://localhost:8080`.

## Project Structure

- `/components`: Contains reusable UI components.
- `/data`: Handles data models and database interactions.
- `/pages`: Contains the main application pages.
- `/static`: Contains static files like CSS and JavaScript.
- `main.go`: The entry point for the application.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
