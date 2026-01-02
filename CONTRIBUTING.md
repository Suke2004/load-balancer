# Contributing to Load Balancer

First off, thank you for considering contributing to this project! It's people like you that make the open-source community such an amazing place to learn, inspire, and create.

## 🤝 How Can I Contribute?

### 1. Reporting Bugs
This section guides you through submitting a bug report.
* **Check existing issues** to see if the bug has already been reported.
* **Open a new Issue** using a descriptive title.
* Include a **reproducible test case** or steps to reproduce the bug.
* Mention your **Go version** and **OS environment**.

### 2. Suggesting Enhancements
* Open an Issue with the tag `enhancement`.
* Explain **why** this enhancement would be useful to most users.

### 3. Pull Requests (Code Contributions)
We follow the standard GitHub "Fork & Pull" workflow.

1.  **Fork** the repository on GitHub.
2.  **Clone** your fork locally:
    ```bash
    git clone [https://github.com/YOUR-USERNAME/load-balancer.git](https://github.com/YOUR-USERNAME/load-balancer.git)
    ```
3.  **Create a branch** for your feature or fix:
    ```bash
    git checkout -b feature/amazing-feature
    # or
    git checkout -b fix/critical-bug
    ```
4.  **Make your changes**. Ensure you write clean, commented code.
5.  **Run Tests** to ensure you haven't broken existing functionality:
    ```bash
    go test ./...
    ```
6.  **Commit** your changes using descriptive commit messages:
    ```bash
    git commit -m "feat: add weighted round-robin algorithm"
    ```
7.  **Push** to your fork:
    ```bash
    git push origin feature/amazing-feature
    ```
8.  **Open a Pull Request** (PR) from your fork's branch to the `main` branch of this repository.

## 🎨 Coding Guidelines

### Go Code
* **Formatting:** Always run `go fmt ./...` before committing.
* **Linting:** We recommend using `golangci-lint` to catch common errors.
* **Naming:** Follow standard Go naming conventions (PascalCase for exported, camelCase for internal).
* **Error Handling:** Don't ignore errors. Handle them or wrap them with context.

### Dashboard (Frontend)
If you are contributing to the `dashboard/` folder:
* Ensure the UI is responsive.
* Keep the design consistent with the existing layout.

## 📋 Code of Conduct

Please note that this project is released with a Contributor Code of Conduct. By participating in this project, you agree to abide by its terms. We expect all contributors to treat others with respect.

## ❓ Questions?

Feel free to open an issue with the tag `question` if you have any doubts about the codebase or how to implement a specific feature.

Happy Coding! 🚀
