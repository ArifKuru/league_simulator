# ⚽ League Simulator

This project is a backend case study developed as part of the **Insider Backend Development Hiring Process**.

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

Test at [here!](https://arifkuru.com/):

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
| POST   | [**/matches/:id/edit**](https://league-simulator-qvbo.onrender.com/matches/1/edit) | Edits a specific match score |

---

## 🛠️ Technologies Used

- **Go** (Golang)
- **Fiber** framework
- **GORM** ORM
- **SQLite** database
