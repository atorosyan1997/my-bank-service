CREATE DATABASE  IF NOT EXISTS `BankServer`;
USE `BankServer`;
-- MySQL dump 10.13  Distrib 5.7.35, for Linux (x86_64)
--
-- Host: localhost    Database: BankServer
-- ------------------------------------------------------
-- Server version	5.7.35-0ubuntu0.18.04.1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `auth`
--

DROP TABLE IF EXISTS `auth`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `auth` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(36) NOT NULL,
  `auth_uuid` varchar(36) NOT NULL,
  UNIQUE KEY `access_auth_id_uindex` (`id`),
  UNIQUE KEY `access_auth_auth_uuid_uindex` (`auth_uuid`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `balance`
--

DROP TABLE IF EXISTS `balance`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `balance` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` varchar(36) NOT NULL,
  `integer_part` double NOT NULL DEFAULT '0',
  `fraction_part` double NOT NULL DEFAULT '0',
  `currency` varchar(20) NOT NULL,
  UNIQUE KEY `balance_id_uindex` (`id`),
  UNIQUE KEY `balance_user_id_uindex` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `balance`
--

LOCK TABLES `balance` WRITE;
/*!40000 ALTER TABLE `balance` DISABLE KEYS */;
INSERT INTO `balance` VALUES (1,'ad987de7-a5cd-400c-ae4a-9eeb52fdf3a8',8,0,'USD');
/*!40000 ALTER TABLE `balance` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `payment_history`
--

DROP TABLE IF EXISTS `payment_history`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `payment_history` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `balance_id` int(11) NOT NULL,
  `createdat` varchar(100) NOT NULL,
  `initial_balance` double NOT NULL,
  `final_balance` double NOT NULL,
  `difference_balance` double NOT NULL,
  UNIQUE KEY `payment_history_id_uindex` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Table structure for table `users`
--

DROP TABLE IF EXISTS `users`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `users` (
  `id` varchar(36) DEFAULT NULL,
  `userName` varchar(250) NOT NULL,
  `email` varchar(100) DEFAULT NULL,
  `password` varchar(255) NOT NULL,
  `tokenhash` varchar(45) DEFAULT NULL,
  `updatedat` varchar(100) DEFAULT NULL,
  `createdat` varchar(100) NOT NULL,
  UNIQUE KEY `users_userName_uindex` (`userName`),
  UNIQUE KEY `users_email_uindex` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `users`
--

LOCK TABLES `users` WRITE;
/*!40000 ALTER TABLE `users` DISABLE KEYS */;
INSERT INTO `users` VALUES ('ad987de7-a5cd-400c-ae4a-9eeb52fdf3a8','admin','admin@gmail.com','$2a$10$Ulut9JhzFjsCj57k/H57juM8dr25.k61xLNpPxLoOaL36EavOFgH.','2r56EcjTEmkWzHr','2021-09-10 23:33:14','2021-09-10 22:17:29');
/*!40000 ALTER TABLE `users` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-09-12 14:04:57
