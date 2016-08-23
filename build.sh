./mjudgectl stop
sleep 2

echo "Building mjudgectl" && go build github.com/RemmargorP/mjudge/cmd/mjudgectl
echo "Building createadmin" && go build github.com/RemmargorP/mjudge/cmd/createadmin

./mjudgectl start