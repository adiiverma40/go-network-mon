<h1 align="center">Go Network Mon</h1>


A Network monitor, That monitors your network speed, Ping, And other stats and Much more. Made in Go-Lang.
```text 
Go Network Mon, Use Ping, It`s Super effective.
```

> [!note]
> The name is inspired by Pokémon.


## Features
- **Telegram notification**: Notify you using Telegram if a bad internet speed is found.
- **Dashboard**: A dashboard for the stats and visualize the result.
- **Docker**: Easily depolyable by Docker.
- **Re-connect Network**: Some router has a timeout in which we have to Re-Authenticate to use network. Automate it.


## Ambitious Feature:
- **Machine Learning**: A ML model that will be trainned on the data to predict the correct network time frame for the next day or today. 


## Roadmap & Architecture

This project is built in two distinct phases: separating the data collection engine from the predictive analytics pipeline.

### Phase 1: Data Collection & Automation (Golang)

A Docker/Background worker that re-connectes, re-loggin if the router needs auth. And logs the data for ML model

* **Automated Re-Authentication:** Detects captive portal timeouts and automatically injects login payloads to keep the internet connection alive.
* **Loggin:** Logs ping, packet loss, bandwidth drops, and network state to a Database at a X-minute intervals.
* **Telegram Notifications:** Pushes instant alerts when a bad internet speed is detected or the connection drops.
* **Dashboard:** A visual interface to track current network stats and uptime.
* **Dockerized:** Easily deployable in a lightweight, self-contained environment.

### Phase 2: Predictive Machine Learning (Python)

A forecasting model to forecast the network status of the next day, it takes, Day, Holiday, Time, and other stuff in considration to predict. 

* **Progressive Modeling:** Starts with a lightweight baseline model (e.g., Prophet or Random Forest) for short-term predictions, scaling to advanced models (XGBoost) as the dataset grows across months.
* **Holiday & Calendar Integration:** Injects calendar features to accurately forecast bandwidth shifts during weekends, exam weeks, and national holidays.
* **Garbage data** Allows users to manually flag and skip "garbage data" (e.g., campus-wide power outages or router hardware failures) so the model does not train on false drops.
