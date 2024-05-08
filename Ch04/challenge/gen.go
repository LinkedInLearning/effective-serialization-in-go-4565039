package rides

/* Tools

# protoc

Install via package manager (apt, brew, choco ...)

# Go Protocol Buffers Plugin

go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
*/

//go:generate mkdir -p pb
//go:generate protoc --go_out=pb --go_opt=paths=source_relative rides.proto
