./memjudgectl stop
sleep 2

echo "Building memjudgectl" && go build github.com/RemmargorP/memjudge/cmd/memjudgectl
echo "Building createadmin" && go build github.com/RemmargorP/memjudge/cmd/createadmin

./memjudgectl start