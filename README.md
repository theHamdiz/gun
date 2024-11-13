# Gun CLI Tool

**Gun** is a powerful and intuitive Command-Line Interface (CLI) tool for the Go programming language. It streamlines the creation of boilerplate code for resource-centric applications, enabling developers to focus on building features rather than setting up repetitive structures. With Gun, you can effortlessly generate projects, models, handlers, routes, middleware, and views—all adhering to Go's best practices and idiomatic conventions.

---

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Getting Started](#getting-started)
    - [Creating a New Project](#creating-a-new-project)
    - [Generating Components](#generating-components)
- [Usage Examples](#usage-examples)
    - [Model Generation](#model-generation)
    - [Handler and Route Generation](#handler-and-route-generation)
    - [Middleware Creation](#middleware-creation)
    - [View Generation](#view-generation)
- [Project Structure](#project-structure)
- [Customization](#customization)
    - [Styling Frameworks](#styling-frameworks)
    - [Including Channels and Signals](#including-channels-and-signals)
- [Contributing](#contributing)
- [License](#license)

---

## Features

- **Project Initialization**: Quickly scaffold a new Go project with a structured directory layout, including separate command packages for the server and API versions.
- **Resource Generation**: Create models, handlers, routes, views, and middleware for resources using simple commands.
- **Standard Library**: Relies solely on Go's standard library—no external templating engines or dependencies.
- **Best Practices**: Generates idiomatic Go code, following industry best practices and conventions.
- **Customization**: Offers flags and options to tailor the generated code to your project's needs.
- **Lightweight App Wrapper**: Provides an `App` struct to encapsulate application initialization and execution logic.
- **Project Model Binding**: Uses a `Project` model to bind CLI input throughout the code generation process for consistency.
- **Styling Options**: Supports styling with Tailwind CSS and shadcn/ui components.
- **Concurrency Utilities**: Optionally include channel utilities and signal handling for advanced concurrency patterns.

---

## Installation

Ensure you have Go installed on your system. Gun requires Go 1.16 or higher.

To install Gun, run:

```bash
go install github.com/theHamdiz/gun@latest
```

This command fetches the latest version of Gun from GitHub and installs it into your `$GOPATH/bin` directory.

---

## Getting Started

### Creating a New Project

To create a new Go project with Gun, use the `new project` command:

```bash
gun new project MyApp --module-name github.com/theHamdiz/MyApp
```

**Options:**

- `--module-name`: Specify the module name for your project (e.g., your GitHub repository path).
- `--style <tailwind|shadcn|both>`: Choose the styling framework(s) to use.
- `--with-channels`: Include channel utilities for concurrency patterns.
- `--with-signals`: Include signal handling utilities for graceful shutdown.

**Example with Options:**

```bash
gun new project MyApp --module-name github.com/theHamdiz/MyApp --style both --with-channels --with-signals
```

This command initializes a new Go project named `MyApp`, sets up the module with the specified module name, includes both Tailwind CSS and shadcn/ui for styling, and creates the structured directory layout, including optional utilities.

### Generating Components

Gun allows you to generate various components for your project:

- **Models**: Data structures representing your application's entities.
- **Handlers**: Functions that handle HTTP requests.
- **Routes**: Definitions that map URLs to handlers.
- **Middleware**: Functions that execute before or after handlers.
- **Views**: HTML templates for rendering web pages.

---

## Usage Examples

### Model Generation

Create a model with specified fields:

```bash
gun generate model User --fields "ID:int Name:string Email:string"
```

**This command:**

- Generates a `User` model in `internal/models/user.go`.
- Includes fields `ID`, `Name`, and `Email` with their respective types.

### Handler and Route Generation

Generate handlers and routes for a resource:

```bash
gun generate handler User
gun generate route User
```

**These commands:**

- Create handler functions for the `User` resource in `internal/handlers/user_handler.go`.
- Set up RESTful routes in `internal/routes/user_routes.go`, mapping HTTP methods and URLs to handlers.

### Middleware Creation

Create custom middleware:

```bash
gun generate middleware auth
```

**This command:**

- Generates an `AuthMiddleware` in `internal/middleware/auth_middleware.go`.
- Provides a template for adding authentication logic.

### View Generation

Generate HTML views for a resource:

```bash
gun generate view User --fields "ID:int Name:string Email:string"
```

**This command:**

- Creates default HTML templates (`index.html`, `show.html`, `edit.html`, `new.html`) in `internal/views/user/`.
- Populates templates with the specified fields.

---

## Project Structure

Gun organizes your project into a clean and modular structure:

```
MyApp/
├── cmd/
│   └── http/
│       ├── api/
│       │   └── v1/
│       │       └── apiv1.go
│       └── server/
│           └── main.go
├── internal/
│   ├── app/
│   │   └── app.go
│   ├── models/
│   │   └── user.go
│   ├── handlers/
│   │   └── user_handler.go
│   ├── routes/
│   │   ├── routes.go
│   │   └── user_routes.go
│   ├── middleware/
│   │   └── auth_middleware.go
│   ├── views/
│   │   └── user/
│   │       ├── index.html
│   │       ├── show.html
│   │       ├── edit.html
│   │       └── new.html
│   └── utils/
│       ├── channels.go   // If --with-channels is used
│       └── signals.go    // If --with-signals is used
├── assets/
│   └── css/
│       └── input.css     // Tailwind CSS input file
├── go.mod
├── package.json          // If styling frameworks are used
```

**Key Directories:**

- **`cmd/`**: Contains the main application entry points.
    - **`http/server/main.go`**: The server's main function.
    - **`http/api/v1/apiv1.go`**: API version 1 setup and route registration.
- **`internal/`**: Holds the core application code.
    - **`app/`**: Contains the `App` wrapper for application setup.
    - **`models/`**: Holds data models.
    - **`handlers/`**: Contains HTTP request handlers.
    - **`routes/`**: Defines routes mapping URLs to handlers.
    - **`middleware/`**: Includes middleware functions.
    - **`views/`**: Stores HTML templates.
    - **`utils/`**: Provides utility functions (e.g., channels, signals).
- **`assets/`**: Contains static assets like CSS files (if styling is used).

---

## Customization

### Styling Frameworks

Gun supports integrating styling frameworks into your project to enhance the visual appearance of your web application.

#### Tailwind CSS

To include Tailwind CSS:

```bash
gun new project MyApp --style tailwind
```

**Features:**

- Utility-first CSS framework for rapid UI development.
- Automatically sets up Tailwind CSS configuration.
- Includes an input CSS file (`assets/css/input.css`) with Tailwind directives.

#### shadcn/ui

To include shadcn/ui components:

```bash
gun new project MyApp --style shadcn
```

**Features:**

- Pre-built, accessible UI components built on top of Tailwind CSS.
- Automatically installs shadcn/ui components.
- Ensures Tailwind CSS is set up as a prerequisite.

#### Both Tailwind CSS and shadcn/ui

To include both:

```bash
gun new project MyApp --style both
```

### Including Channels and Signals

Gun allows you to include utilities for concurrency and graceful shutdown.

#### Channels

Include channel utilities:

```bash
gun new project MyApp --with-channels
```

**Features:**

- Provides templates and examples for concurrency patterns using Go channels.
- Facilitates building concurrent applications.

#### Signals

Include signal handling utilities:

```bash
gun new project MyApp --with-signals
```

**Features:**

- Incorporates signal handling for graceful shutdowns.
- Ensures your application can handle interrupts and terminate cleanly.

---

## Contributing

We welcome contributions to enhance Gun's functionality and improve its usability. Here's how you can help:

1. **Fork the Repository**: Create a personal copy of the repository on your GitHub account.

2. **Clone the Fork**: Clone your forked repository to your local machine.

   ```bash
   git clone https://github.com/theHamdiz/gun.git
   ```

3. **Create a Branch**: Create a new branch for your feature or bug fix.

   ```bash
   git checkout -b feature/your-feature-name
   ```

4. **Make Changes**: Implement your feature or fix.

5. **Commit Changes**: Commit your modifications with a descriptive message.

   ```bash
   git commit -am "Add new feature for X"
   ```

6. **Push to Fork**: Push your changes to your forked repository.

   ```bash
   git push origin feature/your-feature-name
   ```

7. **Open a Pull Request**: Submit a pull request to the main repository.

**Guidelines:**

- Follow Go's best practices and code conventions.
- Write clear and concise commit messages.
- Include documentation and examples where appropriate.
- Ensure your code passes all tests.

---

## License

Gun is released under the [MIT License](LICENSE). This means you can use, modify, and distribute the software freely. See the [LICENSE](LICENSE) file for more details.

---

**Happy coding with Gun! If you have any questions or need assistance, feel free to open an issue or reach out to the maintainers. Let's build amazing Go applications together!**

---

### Notes

- **Requirements**: Ensure you have Node.js and npm installed if you plan to use the styling options, as Tailwind CSS and shadcn/ui require them.
- **Module Name**: When creating a new project, specifying the `--module-name` is crucial for correct import paths, especially if you plan to publish your project or use version control systems like GitHub.
- **Templates**: All templates used by Gun are embedded within the tool and use Go's `text/template` package, ensuring no external dependencies are required.

---

**Example Workflow:**

1. **Create a New Project:**

   ```bash
   gun new project MyApp --module-name github.com/theHamdiz/MyApp --style tailwind --with-channels --with-signals
   ```

2. **Generate a Model:**

   ```bash
   gun generate model User --fields "ID:int Name:string Email:string"
   ```

3. **Generate Handlers and Routes:**

   ```bash
   gun generate handler User
   gun generate route User
   ```

4. **Generate Views:**

   ```bash
   gun generate view User --fields "ID:int Name:string Email:string"
   ```

5. **Run the Application:**

   Navigate to the server directory and run your application:

   ```bash
   cd cmd/http/server
   go run main.go
   ```

---

**Thank you for choosing Gun to accelerate your Go development!**