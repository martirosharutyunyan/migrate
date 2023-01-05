git add .
read -p "Enter commit message: " MESSAGE
git commit -m "${MESSAGE}"
git pull
git push
