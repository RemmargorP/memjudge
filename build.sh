mkdir -p bin > /dev/null 2> /dev/null
cd bin && \
echo "Building mjudgectl" && go build github.com/RemmargorP/mjudge/cmd/mjudgectl && \
echo && \
echo "Building createadmin" && go build github.com/RemmargorP/mjudge/cmd/createadmin && \
cd ..