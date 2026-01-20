# goderpad

sce conducts a cutting-edge industry pipeline through its summer internship program. this program draws upwards of one hundred applicants each year, requiring a hefty amount of interviews. goderpad simplifies the interview process by centralization, allowing our hardworking internship mentors to focus on evaluation. 

## setup
1. copy `/frontend/.env.example` and create a `.env.development` file and enter your API and websocket URLs
2. copy `/backend/config/config.example.yml` and create a `config.yml` file and enter the port that you want your backend to run at

## running this
**with docker**: `docker-compose -f docker-compose.dev.yml up --build`  

**without docker**:  
1. `cd frontend && npm run dev`
2. in a separate terminal: `cd backend && go run main.go`  

the frontend runs at `http://localhost:7777` and the backend runs at `http://localhost:7778`.

**note**: if you have a windows machine, frick you and this app won't run for you
