# Tweety

## Introduction
Welcome to the **Tweety**! This is a modern, feature-rich social networking platform that enables users to connect, share, and interact seamlessly. The platform provides a user-friendly interface, real-time messaging, media sharing, and more.

## Tech Stack
Below is a list of technologies used in this project:

| Package Name | Link                                     | Use Case                                    |
|--------------|------------------------------------------|---------------------------------------------|
| Go        | [Go](https://go.dev/)            | Golang          |
| pgx       | [pgx](https://github.com/jackc/pgx/wiki/Getting-started-with-pgx)             | Postgres driver          |
| Branca       | [branca](https://github.com/hako/branca)             | Secure alternative to JWT          |
| CockroachDB       | [cockroachDb](https://www.cockroachlabs.com/)             | Distributed SQL database          |
| Way router       | [way](https://github.com/matryer/way)             | HTTP router for Go          |

## Database
**Install** -> `brew install cockroachdb/tap/cockroach`

**Verify** -> `cockroach version`

**Usage**
1. Start a single node cluster -> `cockroach start-single-node --insecure --listen-addr=localhost`

2. Stop cockroach server -> `pkill cockroach`

3. Apply schema to cockroach instance -> `cockroach sql --insecure --database=<mydatabase> < <path/to/schema.sql>`

## Images
### Screenshots
![Home Page](./images/homepage.png)
![User Profile](./images/userprofile.png)
![Chat System](./images/chat.png)

## Contribute
We welcome contributions! To contribute:
1. Fork the repository.
2. Clone your forked repository:
   ```sh
   git clone https://github.com/yourusername/social-network.git
   ```
3. Create a new branch for your feature:
   ```sh
   git checkout -b feature-name
   ```
4. Make your changes and commit:
   ```sh
   git commit -m "Added new feature"
   ```
5. Push your changes:
   ```sh
   git push origin feature-name
   ```
6. Open a Pull Request (PR) to the main branch.

## License
This project is licensed under the MIT License.

---

Happy coding! 🚀