-- MySQL dump 10.13  Distrib 8.0.35, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: proxyapi
-- ------------------------------------------------------
-- Server version	8.0.35

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `members`
--

DROP TABLE IF EXISTS `members`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `members` (
  `id` int NOT NULL AUTO_INCREMENT,
  `uid` varchar(20) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `type` tinyint DEFAULT '0',
  `mobile` varchar(15) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `email` varchar(50) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `is_delete` tinyint NOT NULL DEFAULT '0',
  `created_at` int NOT NULL,
  `updated_at` int NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `members`
--

LOCK TABLES `members` WRITE;
/*!40000 ALTER TABLE `members` DISABLE KEYS */;
INSERT INTO `members` VALUES (1,'200',1,'156188','Email@aefasf.com',0,1705112088,1705112088);
/*!40000 ALTER TABLE `members` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `request_logs`
--

DROP TABLE IF EXISTS `request_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `request_logs` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `client_id` int DEFAULT NULL,
  `path` varchar(100) COLLATE utf8mb4_unicode_520_ci DEFAULT NULL,
  `params` json DEFAULT NULL,
  `app_id` int DEFAULT NULL,
  `created_at` int NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `request_logs_client_id_app_id_created_at_index` (`client_id`,`app_id`,`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_520_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `request_logs`
--

LOCK TABLES `request_logs` WRITE;
/*!40000 ALTER TABLE `request_logs` DISABLE KEYS */;
INSERT INTO `request_logs` VALUES (1,1,'/post','{\"password\": \"dsgsdg\", \"username\": \"testuser\"}',0,1705146504),(2,1,'/post','{\"password\": \"dsgsdg\", \"username\": \"testuser\"}',0,1705146594),(3,1,'/post','{\"password\": \"dsgsdg\", \"username\": \"testuser\"}',4627,1705147390),(4,1,'/post','{\"password\": \"dsgsdg\", \"username\": \"testuser\"}',4627,1705147572),(5,1,'/post','{\"password\": \"dsgsdg\", \"username\": \"testuser\"}',4627,1705147585),(6,1,'/post','{\"password\": \"dsgsdg\", \"username\": \"testuser\"}',4627,1705147607);
/*!40000 ALTER TABLE `request_logs` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-01-13 20:08:30