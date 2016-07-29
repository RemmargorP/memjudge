control/control stop
cd control
go build
cd ..

sudo sh copyhtmls.sh

sleep 2
control/control start
