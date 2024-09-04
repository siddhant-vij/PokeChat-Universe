# PokeChat Universe

A full-stack web application built using the GoTTH stack (Golang, Templ, TailwindCSS, HTMX). The app allows users to explore a universe of Pokémon, collect their favorites, and engage in AI-driven conversations with them.

- Why? I mean why not?

Leveraging a local instance of the Llama3 model via Ollama, the app provides real-time chat features powered by SSE (Server-Sent Events) for a seamless experience.

The web app also includes OAuth2/OIDC-based login via Auth0, dynamic frontend features (search, sort, pagination), and a fully scalable backend with persistent chat histories stored in PostgreSQL. Designed with performance and scalability in mind, the app is capable of handling high volumes of concurrent users while maintaining a smooth, responsive interface.

<br>

## Table of Contents

1. [Product Features](#product-features)
1. [Technical Scope](#technical-scope)
1. [Future Improvements](#future-improvements)
1. [Contributing](#contributing)
1. [License](#license)

<br>

## Product Features

- **Explore Pokémon Universe**: Browse and filter a comprehensive list of Pokémon.
- **OAuth2/OIDC Authentication**: Secure user login with Auth0.
- **Collect Pokémon**: Add favorite Pokémon to your personal collection.
- **User-Specific Pokedex**: View all collected Pokémon in a dedicated tab.
- **Search, Sort, and Load More**: Dynamic frontend with real-time search, sorting options, and pagination for large lists.
- **Dynamic Frontend Updates**: Seamless user interactions using HTMX without full-page reloads.
- **Real-Time AI Chat**: Engage in conversations with collected Pokémon powered by Llama3 via Ollama.
- **Persistent Chat Histories**: Conversations are stored in a database for later retrieval.

<br>

## Technical Scope

- **Frontend**: TailwindCSS for styling, HTMX for dynamic interactions, and Templ for rendering.
- **Backend**: Golang-based API, handling user sessions, chat logic, and serving data via REST.
- **Database**: PostgreSQL for relational data (users, Pokémon, chat histories) with Redis for session storage.
- **Authentication**: OAuth2/OIDC-based login with Auth0.
- **AI Integration**: Local Llama3 instance running via Ollama for generating contextual responses in real-time.
- **Real-Time Communication**: Server-Sent Events for streaming AI responses to the frontend.

<br>

## Future Improvements

- **Microservices**: Split the app into microservices to handle specific tasks like chat, authentication, and AI responses independently.
- **Database Partitioning and Indexing**: Implement sharding, table partitioning, and advanced indexing strategies as chat histories grow.
- **Horizontal Scaling**: Use container orchestration (like Kubernetes) to handle increased load across multiple instances.
- **Message Queues for AI Processing**: Offload AI requests to message queues (e.g., RabbitMQ) for more efficient handling of heavy processing tasks.
- **Data Archival**: Implement strategies for archiving older chat histories to keep the primary database lean.

<br>

## Contributing

Contributions are what make the open-source community such an amazing place to learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. **Fork the Project**
2. **Create your Feature Branch**:
   ```bash
   git checkout -b feature/AmazingFeature
   ```
3. **Commit your Changes**:
   ```bash
   git commit -m 'Add some AmazingFeature'
   ```
4. **Push to the Branch**:
   ```bash
   git push origin feature/AmazingFeature
   ```
5. **Open a Pull Request**

<br>

## License

Distributed under the MIT License. See [`LICENSE`](https://github.com/siddhant-vij/PokeChat-Universe/blob/main/LICENSE) for more information.
