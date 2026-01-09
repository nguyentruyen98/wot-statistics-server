# 🚀 Hướng dẫn Import Tank Data vào Database

Hướng dẫn từng bước để import dữ liệu tank từ file `tanks_X.json` vào PostgreSQL database.

---

## 📋 Yêu cầu

- Docker và Docker Compose đã cài đặt
- Go 1.21+ đã cài đặt
- File `tanks_X.json` ở thư mục root của project

---

## 🎯 Các Cách Import Data

### **Cách 1: Tự động hoàn toàn (Khuyến nghị)** ⭐

Script này sẽ tự động:
- Khởi động PostgreSQL container
- Tạo database
- Chạy migrations
- Import dữ liệu

```bash
make seed
```

**Hoặc:**

```bash
chmod +x scripts/seed.sh
./scripts/seed.sh
```

---

### **Cách 2: Import nhanh (Database đã sẵn sàng)** ⚡

Nếu database đã được setup và migrations đã chạy:

```bash
make seed-quick
```

**Hoặc:**

```bash
chmod +x scripts/seed-quick.sh
./scripts/seed-quick.sh
```

---

### **Cách 3: Chạy trực tiếp với Go** 🔧

```bash
make seed-go
```

**Hoặc:**

```bash
DATABASE_URL="postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable" \
TANKS_JSON_FILE="tanks_X.json" \
go run cmd/seed/main.go
```

---

## 📝 Setup Thủ công Từng Bước

Nếu bạn muốn chạy từng bước riêng biệt:

### **Bước 1: Khởi động PostgreSQL**

```bash
docker-compose up -d postgres
```

Đợi 5 giây để PostgreSQL khởi động hoàn toàn.

### **Bước 2: Tạo Database**

```bash
docker exec wot-statistics-postgres psql -U postgres -c "CREATE DATABASE \"wot-statistics\";"
```

Nếu database đã tồn tại, bạn sẽ thấy message: `database "wot-statistics" already exists`

### **Bước 3: Chạy Migration**

```bash
docker exec -i wot-statistics-postgres psql -U postgres -d wot-statistics < migrations/001_create_tanks_table.up.sql
```

### **Bước 4: Import Data**

```bash
DATABASE_URL="postgres://postgres:postgres@localhost:5432/wot-statistics?sslmode=disable" \
TANKS_JSON_FILE="tanks_X.json" \
go run cmd/seed/main.go
```

---

## ✅ Kết quả mong đợi

Sau khi chạy thành công, bạn sẽ thấy:

```
🚀 Starting tank data seeder...
📡 Connecting to database...
✅ Database connected successfully
📄 Reading tanks data from: /path/to/tanks_X.json

============================================================
✅ Successfully imported 110 tanks!
⏱️  Time taken: 2.5s
============================================================

📊 Database Statistics:
   Total tanks: 110
   Premium tanks: 35

   By Nation:
      ussr: 25
      usa: 20
      germany: 18
      france: 12
      uk: 10
      china: 10
      japan: 8
      czech: 7

   By Tier:
      Tier 10: 110

   By Type:
      heavyTank: 45
      mediumTank: 30
      lightTank: 15
      AT-SPG: 15
      SPG: 5
```

---

## 🔍 Kiểm tra Data đã Import

### **Cách 1: Dùng psql**

```bash
make db-psql
```

Sau đó chạy SQL:

```sql
-- Đếm tổng số tanks
SELECT COUNT(*) FROM tanks;

-- Xem 5 tanks đầu tiên
SELECT tank_id, name, nation, tier, type FROM tanks LIMIT 5;

-- Đếm theo nation
SELECT nation, COUNT(*) FROM tanks GROUP BY nation ORDER BY COUNT(*) DESC;

-- Đếm theo tier
SELECT tier, COUNT(*) FROM tanks GROUP BY tier ORDER BY tier;

-- Xem premium tanks
SELECT name, nation, tier FROM tanks WHERE is_premium = true;
```

### **Cách 2: Dùng API**

Nếu bạn đã chạy server:

```bash
# Khởi động server
make run
```

Mở browser hoặc dùng curl:

```bash
# Lấy tất cả tanks
curl http://localhost:8080/api/tanks

# Lấy tanks theo filter
curl "http://localhost:8080/api/tanks?nation=ussr&tier=10"

# Lấy 1 tank cụ thể
curl http://localhost:8080/api/tanks/6145
```

---

## 🛠️ Troubleshooting

### **Lỗi: PostgreSQL container không chạy**

```bash
docker-compose up -d postgres
docker ps | grep postgres
```

### **Lỗi: Database connection failed**

Kiểm tra PostgreSQL đã sẵn sàng:

```bash
docker exec wot-statistics-postgres pg_isready -U postgres
```

### **Lỗi: tanks_X.json not found**

Đảm bảo file `tanks_X.json` ở thư mục root:

```bash
ls -la tanks_X.json
```

### **Lỗi: Migration failed - table already exists**

Bỏ qua lỗi này, table đã tồn tại là OK.

Hoặc reset database:

```bash
make db-reset
```

### **Lỗi: Permission denied on scripts**

```bash
chmod +x scripts/seed.sh scripts/seed-quick.sh
```

### **Lỗi: Wire generation failed**

```bash
cd wire
go generate
```

---

## 🔄 Reset và Import lại

Nếu muốn xóa hết và import lại từ đầu:

```bash
# Reset database (xóa hết)
make db-reset

# Import lại
make seed-quick
```

---

## 📊 Import Data qua API

Bạn cũng có thể import data qua API endpoint:

```bash
# Start server
make run

# Import via API
curl -X POST http://localhost:8080/api/tanks/import \
  -H "Content-Type: application/json" \
  -d '{"file_path": "tanks_X.json"}'
```

---

## 🎓 Tips

1. **Lần đầu tiên:** Dùng `make seed` (full setup)
2. **Import lại:** Dùng `make seed-quick` (nhanh hơn)
3. **Development:** Dùng `make db-reset` để reset và test lại
4. **Production:** Nên dùng migration tools như golang-migrate

---

## 📚 Các lệnh hữu ích

```bash
# Database setup
make db-setup          # Setup database + migrations
make db-reset          # Reset database
make db-psql           # Open PostgreSQL shell

# Seeding
make seed              # Full auto seed
make seed-quick        # Quick seed
make seed-go           # Direct go run

# Application
make run               # Start API server
make help              # Show all commands
```

---

## ❓ FAQ

**Q: Tôi có thể import nhiều lần không?**  
A: Có! Seed script dùng UPSERT (ON CONFLICT DO UPDATE), nên data sẽ được update nếu đã tồn tại.

**Q: Import mất bao lâu?**  
A: Khoảng 2-5 giây cho 110 tanks.

**Q: Tôi có thể thay đổi file JSON khác không?**  
A: Có! Set biến môi trường:
```bash
TANKS_JSON_FILE="my_tanks.json" go run cmd/seed/main.go
```

**Q: Làm sao kiểm tra database đã có data?**  
A: 
```bash
make db-psql
SELECT COUNT(*) FROM tanks;
```

---

## 🎉 Hoàn thành!

Sau khi import thành công, bạn có thể:

1. ✅ Query data qua API: `http://localhost:8080/api/tanks`
2. ✅ Filter tanks: `?nation=ussr&tier=10&premium=true`
3. ✅ Get single tank: `http://localhost:8080/api/tanks/6145`
4. ✅ CRUD operations qua API

Chúc bạn code vui vẻ! 🚀
