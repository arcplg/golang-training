module example.com/hello

go 1.23.2

// run command: go mod edit -replace example.com/greetings=../greetings
// đảm bảo xuất bản module greetings => example.com/hello để nó có thể tìm thấy mã example.com/greetings trên hệ thống
replace example.com/greetings => ../greetings

// run command: go mod tidy
// để đồng bộ hóa các phụ thuộc của mô-đun example.com/hello, thêm những phụ thuộc mà mã yêu cầu nhưng chưa được theo dõi trong mô-đun.
require example.com/greetings v0.0.0-00010101000000-000000000000
