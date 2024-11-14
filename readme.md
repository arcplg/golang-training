                +-------------------------+
                |       Client Side       |
                |   (Web/Mobile App)      |
                +-----------+-------------+
                            |
                            |
                    (gRPC API & WebSocket)
                            |
            +---------------+----------------+
            |       Chat Application Server  |
            |               (Go)             |
            +--------+---------------+-------+
                     |               |
                     |               |
         +-----------+               +------------+
         |                                        |
+--------v--------+                     +---------v--------+
|   gRPC Services |                     |   WebSocket      |
|                 |                     |   Service        |
+--------+--------+                     +--------+---------+
         |                                        |
         |                                        |
+--------v--------+                       +-------v--------+
| Authentication  |                       | Room &         |
| Service         |                       | Message        |
+--------+--------+                       | Broadcaster    |
         |                                +----------------+
         |
+--------v--------+
|  PostgreSQL     |
|  Database       |
+-----------------+



chat-app/
├── cmd/
│   └── server/          # Chứa mã khởi chạy ứng dụng
│       └── main.go
├── internal/
│   ├── handlers/        # Các hàm xử lý gRPC và WebSocket
│   ├── models/          # Định nghĩa các model Go
│   ├── services/        # gRPC services
│   └── storage/         # Database và caching
├── proto/               # Định nghĩa các file .proto cho gRPC
├── go.mod
└── go.sum


