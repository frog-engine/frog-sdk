module github.com/frog-engine/frog-sdk

replace github.com/frog-engine/frog-sdk => .

go 1.23

require github.com/sirupsen/logrus v1.9.3

require (
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
)
