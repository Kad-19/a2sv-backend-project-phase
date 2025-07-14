# Library Management System (Go Console Application)

This is a simple console-based Library Management System written in Go. The application allows users to manage library members and books, as well as handle borrowing and returning of books. It is designed for educational purposes and demonstrates basic CRUD operations and business logic in Go.

## Features

- **Add Members:** Register new members to the library.
- **Delete Members:** Remove existing members from the library.
- **List All Members:** Display a list of all registered members.
- **Add Books:** Add new books to the library's collection.
- **Delete Books:** Remove books from the library.
- **Borrow Books:** Allow members to borrow available books.
- **Return Books:** Allow members to return borrowed books.
- **List All Available Books:** Show all books currently available for borrowing.
- **List All Borrowed Books by a Member:** Display all books currently borrowed by a specific member.

## Project Structure

```
Task3/
  library_management/
    main.go                # Entry point of the application
    controllers/           # Handles user input and output
      library_controller.go
    models/                # Data models for books and members
      book.go
      member.go
    services/              # Business logic for library operations
      library_service.go
    docs/
      documentation.md     # This documentation
    go.mod, go.sum         # Go module files
```

## How to Run

1. **Install Go:** Ensure you have Go installed on your system. You can download it from [https://golang.org/dl/](https://golang.org/dl/).
2. **Navigate to the Project Directory:**
   ```sh
   cd Task3/library_management
   ```
3. **Run the Application:**
   ```sh
   go run main.go
   ```
4. **Follow the On-Screen Menu:** Use the console menu to perform library operations.

## Usage

When you run the application, you will be presented with a menu to choose from the available functionalities. Enter the corresponding number to perform an action, such as adding a member or borrowing a book. The application will prompt you for any required information (e.g., member name, book title).

## Example Operations

- **Add a Member:** Enter the member's details when prompted.
- **Add a Book:** Enter the book's details (title, author, etc.).
- **Borrow a Book:** Select a member and an available book to borrow.
- **Return a Book:** Select the member and the book to return.
- **List All Members/Books:** View all registered members or available books.

## Notes

- The application stores data in memory; all data will be lost when the program exits.
- Input validation and error handling are implemented for a smooth user experience.

## License

This project is for educational purposes.
