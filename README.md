# Hiểu gì về kafka ?
- Kafka là event log phân tán, nơi mọi sự kiện được publish–subscribe, lưu trữ bền vững, và xử lý song song, thường dùng
  để kết nối microservices và phân tích dữ liệu real-time, throughput cao và khả năng scale ngang.

Bên trong kafka gồm các thành phần chính:
- Broker: Máy chủ lưu trữ và quản lý các topic.
- Kafka cluster: Tập hợp nhiều broker làm việc cùng nhau để cung cấp tính sẵn sàng cao và khả năng mở rộng.
- Producer: Gửi sự kiện (message) đến các topic trong Kafka.
- Consumer: Đọc sự kiện từ các topic.
- Topic: Kênh để tổ chức và phân loại các sự kiện.
- Partition: Mỗi topic được chia thành các phân vùng để tăng khả năng xử lý song song.
- Zookeeper: Quản lý cấu hình và điều phối các broker trong cụm Kafka.
- Offset: Vị trí của một message trong partition, giúp consumer theo dõi việc đọc message.
- Replication: Sao chép dữ liệu giữa các broker để đảm bảo tính sẵn sàng và độ bền.
- Retention: Cấu hình thời gian lưu trữ message trong topic trước khi bị xóa.
- Consumer Group: Tập hợp các consumer làm việc cùng nhau để chia sẻ việc đọc message từ các partition.

# Hiểu gì về RabbitMQ ?
- RabbitMQ là message broker mã nguồn mở, sử dụng giao thức AMQP để gửi và nhận message giữa các ứng dụng.
- RabbitMQ hỗ trợ nhiều mô hình messaging như point-to-point, publish/subscribe, và request/reply.
- RabbitMQ cung cấp các tính năng như routing linh hoạt, message acknowledgment, và hỗ trợ nhiều ngôn ngữ lập trình.

Bên trong RabbitMQ gồm các thành phần chính:
- Broker: Máy chủ RabbitMQ quản lý các hàng đợi và xử lý message.
- Producer: Ứng dụng gửi message đến RabbitMQ.
- Consumer: Ứng dụng nhận message từ RabbitMQ.
- Queue: Hàng đợi lưu trữ message cho đến khi chúng được tiêu thụ.
- Exchange: Thành phần định tuyến message đến các hàng đợi dựa trên các quy tắc định tuyến.
- Binding: Liên kết giữa exchange và queue để xác định cách message được định tuyến.
- Channel: Kênh giao tiếp giữa producer/consumer và RabbitMQ broker.
- Virtual Host: Không gian tên để phân tách các tài nguyên trong RabbitMQ.
- Acknowledgment: Cơ chế xác nhận rằng message đã được tiêu thụ thành công.
- Dead Letter Exchange: Exchange đặc biệt để xử lý message không thể tiêu thụ.
- TTL (Time-To-Live): Cấu hình thời gian sống của message trong hàng đợi trước khi bị xóa.
- Prefetch Count: Giới hạn số lượng message mà consumer có thể nhận trước khi gửi acknowledgment.
- Clustering: Kết nối nhiều broker RabbitMQ để cung cấp tính sẵn sàng cao và khả năng mở rộng.
- Mirroring: Sao chép hàng đợi giữa các broker trong cụm để đảm bảo tính sẵn sàng và độ bền.

# Kafka khác gì RabbitMQ
RabbitMQ là message broker truyền thống theo mô hình queue-based, ưu tiên routing linh hoạt và ack per message.
Kafka thì là distributed log được thiết kế cho throughput cao, lưu trữ lâu dài, và khả năng xử lý event streaming real-time.