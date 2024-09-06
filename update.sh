cd ~/solvmaster_backend
pkill uvicorn
git pull
/home/ubuntu/env/bin/python3 /home/ubuntu/env/bin/uvicorn backend.main:app --reload > log.txt & disown