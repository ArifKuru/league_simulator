# ⚽ League Simulator

This project is a backend case study developed as part of the **Insider Backend Development Intern Case Study**.

It simulates a simple football league where teams play weekly matches, standings are updated, and championship predictions are calculated dynamically based on match outcomes and team strength.

---

## 🧠 Features

- Weekly match simulation with random but strength-based scoring
- Morale updates based on match results
- League standings table (PTS, W/D/L, GD, etc.)
- Match results viewable week by week
- Championship predictions using:
  - Default algorithm
  - Monte Carlo simulation
- API-first architecture
- Minimal frontend with pure HTML/JS for visualizing data

---

## 🚀 Live Demo

Test at [here!](https://arifkuru.com/insider):<br>
<small>Note that initial request for endpoint may take time up to 30 seconds since project backend using render for deployment in free plan</small><br>
Backend is deployed on [Render](https://render.com/):

> 🔗 [https://league-simulator-qvbo.onrender.com](https://league-simulator-qvbo.onrender.com)

---

## 📡 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | [**/teams**](https://league-simulator-qvbo.onrender.com/teams) | Returns all teams |
| POST   | [**/simulate**](https://league-simulator-qvbo.onrender.com/simulate) | Simulates the next week |
| GET    | [**/matches**](https://league-simulator-qvbo.onrender.com/matches) | Returns all matches |
| GET    | [**/standings**](https://league-simulator-qvbo.onrender.com/standings) | Returns current league standings |
| GET    | [**/predict**](https://league-simulator-qvbo.onrender.com/predict) | Predicts champion (default method) |
| GET    | [**/predict?method=montecarlo**](https://league-simulator-qvbo.onrender.com/predict?method=montecarlo) | Predicts champion using Monte Carlo |
| GET    | [**/matches/last**](https://league-simulator-qvbo.onrender.com/matches/last) | Returns last played week matches |
| POST   | [**/simulate/all**](https://league-simulator-qvbo.onrender.com/simulate/all) | Simulates all remaining weeks |
| POST   | [**/reset**](https://league-simulator-qvbo.onrender.com/reset) | Resets all match and season data |
| POST   | [**/matches/:id/edit**](https://league-simulator-qvbo.onrender.com/matches/1/edit) | Edits a specific match score send 'home_score' and 'away_score' in json |

---

## 🛠️ Technologies Used

- **Go** (Golang)
- **Fiber** framework
- **GORM** ORM
- **SQLite** database
## 🧱 Architecture Overview

> A high-level view of the layered backend system.

<img src="https://github.com/user-attachments/assets/a185b55e-f46e-4872-a33a-19cb9173a00e" alt="Architecture" width="700"/>

---

## 🗃️ Database Schema (UML)

> Entity relationships for Team, Match, Standing, SeasonState.

<img src="https://github.com/user-attachments/assets/3e9d04c2-2f23-4729-8da7-ce6a8acd2372" alt="DB UML" width="650"/>

---

## 🔁 Simulation & Prediction Workflow

> Sequence of how simulate & predict endpoints work through the system.

<img src="https://github.com/user-attachments/assets/72344746-d132-4ebe-9048-43113bd2a018" alt="Workflow Diagram" width="750"/>

---

## 🖥️ Final System Screenshot

> Visual representation of the running UI (standings, matches, predictions).

<img src="https://github.com/user-attachments/assets/9478ccfe-a212-4e8a-8f64-6e939ce1a4d1" alt="Final UI Screenshot" width="800"/>

---

### TO SETUP THIS PROJECT ON LOCAL 
## ✅ Prerequisites

To run this project locally, make sure you have the following installed:

- **Go** (v1.24 or higher): [Install Go](https://go.dev/doc/install)
- **Git**: Used to clone the repository. [Download Git](https://git-scm.com/downloads)


---

## 🚀 Installation & Local Setup

Follow these steps to clone and run the project on your machine:

### 1. Clone the repository

```bash
git clone https://github.com/ArifKuru/league_simulator.git
cd league_simulator
```
```bash
go run cmd/main.go
```
Then You May Test endpoint via Postman 
here endpoint collection <a href='https://web.postman.co/workspace/44d159a8-9d5a-4c08-8db4-3a7669347591'>Postman Link</a>
