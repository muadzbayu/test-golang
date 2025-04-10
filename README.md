# test-golang

# 1. Setelah git clone, jalankan go mod tidy untuk mendownload modulenya (pastikan go sudah diinstal)
# 2. Create Manual DB article 
     CREATE DATABASE `article`

# 3. Create Table Manual (opsional). karena sudah auto migrate     
     CREATE TABLE `posts` (
     `id` bigint unsigned NOT NULL AUTO_INCREMENT,
     `title` varchar(200) NOT NULL,
     `content` text,
     `category` varchar(100) DEFAULT NULL,
     `created_date` datetime(3) DEFAULT NULL,
     `updated_date` datetime(3) DEFAULT NULL,
     `status` varchar(100) DEFAULT NULL,
     PRIMARY KEY (`id`)
     ) ENGINE=MyISAM AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
    
# 4. Jangan lupa buat file .env bisa ambil dari .env-example
     
# 4. Jalankan GO, bisa ketik air (auto reload) atau "go run main.go"
# 5. Selesai
