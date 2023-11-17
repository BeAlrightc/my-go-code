-- MySQL dump 10.13  Distrib 5.7.19, for Win64 (x86_64)
--
-- Host: localhost    Database: testt
-- ------------------------------------------------------
-- Server version	5.7.19

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
-- Table structure for table `course`
--

DROP TABLE IF EXISTS `course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `course` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(10) DEFAULT NULL COMMENT '课程名称',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='课程表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `course`
--

LOCK TABLES `course` WRITE;
/*!40000 ALTER TABLE `course` DISABLE KEYS */;
INSERT INTO `course` VALUES (1,'java'),(2,'php'),(3,'mysql'),(4,'hadoop');
/*!40000 ALTER TABLE `course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student`
--

DROP TABLE IF EXISTS `student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `student` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(10) DEFAULT NULL COMMENT '姓名',
  `no` varchar(10) DEFAULT NULL COMMENT '学号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='学生表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student`
--

LOCK TABLES `student` WRITE;
/*!40000 ALTER TABLE `student` DISABLE KEYS */;
INSERT INTO `student` VALUES (1,'张无忌','200250221'),(2,'赵敏','200250222'),(3,'刘发财','200250223'),(4,'李云龙','200250224');
/*!40000 ALTER TABLE `student` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student_course`
--

DROP TABLE IF EXISTS `student_course`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `student_course` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `student_id` int(11) NOT NULL COMMENT '学生ID',
  `course_id` int(11) NOT NULL COMMENT '课程ID',
  PRIMARY KEY (`id`),
  KEY `fk_courseid` (`course_id`),
  KEY `fk_studentid` (`student_id`),
  CONSTRAINT `fk_courseid` FOREIGN KEY (`course_id`) REFERENCES `course` (`id`),
  CONSTRAINT `fk_studentid` FOREIGN KEY (`student_id`) REFERENCES `student` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='学生课程中间表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student_course`
--

LOCK TABLES `student_course` WRITE;
/*!40000 ALTER TABLE `student_course` DISABLE KEYS */;
INSERT INTO `student_course` VALUES (1,1,1),(2,1,2),(3,1,3),(4,2,2),(5,2,3),(6,3,4);
/*!40000 ALTER TABLE `student_course` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_user`
--

DROP TABLE IF EXISTS `tb_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `name` varchar(10) DEFAULT NULL COMMENT '姓名',
  `age` int(11) DEFAULT NULL COMMENT '年龄',
  `gender` char(1) DEFAULT NULL COMMENT '1:男，2:女',
  `phone` char(11) DEFAULT NULL COMMENT '手机号',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='用户基本信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_user`
--

LOCK TABLES `tb_user` WRITE;
/*!40000 ALTER TABLE `tb_user` DISABLE KEYS */;
INSERT INTO `tb_user` VALUES (1,'吕布',23,'2','24463849323'),(2,'张飞',45,'2','63697899323'),(3,'关羽',31,'1','63967843223'),(4,'赵云',21,'1','93367899323');
/*!40000 ALTER TABLE `tb_user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_user_edu`
--

DROP TABLE IF EXISTS `tb_user_edu`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_user_edu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `degree` varchar(20) DEFAULT NULL COMMENT '学历',
  `major` varchar(50) DEFAULT NULL COMMENT '专业',
  `primaryschool` varchar(50) DEFAULT NULL COMMENT '小学',
  `middleschool` varchar(50) DEFAULT NULL COMMENT '中学',
  `university` varchar(50) DEFAULT NULL COMMENT '大学',
  `userid` int(11) DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  UNIQUE KEY `userid` (`userid`),
  CONSTRAINT `fk_userid` FOREIGN KEY (`userid`) REFERENCES `tb_user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='用户信息表';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_user_edu`
--

LOCK TABLES `tb_user_edu` WRITE;
/*!40000 ALTER TABLE `tb_user_edu` DISABLE KEYS */;
INSERT INTO `tb_user_edu` VALUES (1,'本科','舞蹈','厦门一中','双十中学','北京舞蹈学院\r\n',1),(2,'硕士','表演','厦门一中','双十中学','北京电影学院',2),(3,'本科','英语','杭州市第一小学','杭州市第一中控学','杭州师范大学',3),(4,'本科','应用数学','阳泉第一小学','阳泉第一中学','清华大学',4);
/*!40000 ALTER TABLE `tb_user_edu` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `tb_user_pro`
--

DROP TABLE IF EXISTS `tb_user_pro`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `tb_user_pro` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `profession` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_user_pro`
--

LOCK TABLES `tb_user_pro` WRITE;
/*!40000 ALTER TABLE `tb_user_pro` DISABLE KEYS */;
INSERT INTO `tb_user_pro` VALUES (1,'项羽','金属材料'),(2,'刘备','工商管理'),(3,'张飞','机械工程');
/*!40000 ALTER TABLE `tb_user_pro` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `name` varchar(10) DEFAULT NULL COMMENT '姓名',
  `phone` char(11) DEFAULT NULL COMMENT '手机号',
  `email` varchar(20) DEFAULT NULL COMMENT '邮箱',
  `profession` varchar(20) DEFAULT NULL COMMENT '专业',
  `age` int(11) DEFAULT NULL COMMENT '年龄',
  `gender` int(11) DEFAULT NULL COMMENT '级别',
  `status` int(11) DEFAULT NULL COMMENT '状态',
  `createtime` datetime DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'项羽','11979765491','113@qq.com','土木工程',9,0,1,'2001-04-22 00:00:00'),(2,'刘备','12979765492','123@qq.com','土木工程',15,1,1,'2002-02-22 00:00:00'),(3,'张飞','13979765493','133@qq.com','土木工程',17,2,1,'2003-03-22 00:00:00'),(4,'关羽','14979765494','143@qq.com','土木工程',24,3,1,'2004-04-22 00:00:00'),(5,'赵云','15979765495','153@qq.com','土木工程',25,4,1,'2005-05-22 00:00:00'),(6,'吕布','16979765496','163@qq.com','化学装备',42,6,1,'2006-06-22 00:00:00'),(7,'吕蒙','17979765497','173@qq.com','体育锻炼',22,7,1,'2007-07-22 00:00:00');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 trigger user_insert_trigger
    after insert on user for each row
begin
    insert into user_logs(id, operation, operate_time, operate_id, operate_params)values
    (null,'insert',now(),new.id,concat('插入的数据内容为：id=',new.id,',name=',new.name,',phone=',new.phone,',email=',new.email,',profession=',new.profession));
end */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 trigger user_update_trigger
    after update on user for each row
begin
    insert into user_logs(id, operation, operate_time, operate_id, operate_params)values
    (null,'update',now(),new.id,
     concat('更新之前的数据内容为：id=',old.id,',name=',old.name,',phone=',old.phone,',email=',old.email,',profession=',old.profession,
    '   更新之后的数据内容为：id=',new.id,',name=',new.name,',phone=',new.phone,',email=',new.email,',profession=',new.profession));
end */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8mb4 */ ;
/*!50003 SET character_set_results = utf8mb4 */ ;
/*!50003 SET collation_connection  = utf8mb4_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
/*!50003 CREATE*/ /*!50017 DEFINER=`root`@`localhost`*/ /*!50003 trigger user_delete_trigger
    after delete on user for each row
begin
    insert into user_logs(id, operation, operate_time, operate_id, operate_params)values
    (null,'delete',now(),old.id,
     concat('删除之前的数据内容为：id=',old.id,',name=',old.name,',phone=',old.phone,',email=',old.email,',profession=',old.profession));
end */;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;

--
-- Table structure for table `user_logs`
--

DROP TABLE IF EXISTS `user_logs`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user_logs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `operation` varchar(20) NOT NULL COMMENT '操作类型 insert into/update/delete',
  `operate_time` datetime NOT NULL COMMENT '操作时间',
  `operate_id` int(11) NOT NULL COMMENT '操作id',
  `operate_params` varchar(500) DEFAULT NULL COMMENT '操作参数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user_logs`
--

LOCK TABLES `user_logs` WRITE;
/*!40000 ALTER TABLE `user_logs` DISABLE KEYS */;
INSERT INTO `user_logs` VALUES (1,'insert','2023-08-12 10:39:44',8,'插入的数据内容为：id=8,name=王老五,phone=12979765491,email=119@qq.com,profession=计算与科学'),(2,'update','2023-08-12 11:13:02',8,'更新之前的数据内容为：id=8,name=王老五,phone=12979765491,email=119@qq.com,profession=计算与科学   更新之后的数据内容为：id=8,name=王老五,phone=12979765491,email=119@qq.com,profession=计算与科学'),(3,'update','2023-08-12 11:16:00',8,'更新之前的数据内容为：id=8,name=王老五,phone=12979765491,email=119@qq.com,profession=计算与科学   更新之后的数据内容为：id=8,name=王老五,phone=12979765491,email=119@qq.com,profession=土木工程'),(4,'update','2023-08-12 11:20:43',1,'更新之前的数据内容为：id=1,name=项羽,phone=11979765491,email=113@qq.com,profession=金属材料   更新之后的数据内容为：id=1,name=项羽,phone=11979765491,email=113@qq.com,profession=土木工程'),(5,'update','2023-08-12 11:20:43',2,'更新之前的数据内容为：id=2,name=刘备,phone=12979765492,email=123@qq.com,profession=工商管理   更新之后的数据内容为：id=2,name=刘备,phone=12979765492,email=123@qq.com,profession=土木工程'),(6,'update','2023-08-12 11:20:43',3,'更新之前的数据内容为：id=3,name=张飞,phone=13979765493,email=133@qq.com,profession=机械工程   更新之后的数据内容为：id=3,name=张飞,phone=13979765493,email=133@qq.com,profession=土木工程'),(7,'update','2023-08-12 11:20:43',4,'更新之前的数据内容为：id=4,name=关羽,phone=14979765494,email=143@qq.com,profession=软件工程   更新之后的数据内容为：id=4,name=关羽,phone=14979765494,email=143@qq.com,profession=土木工程'),(8,'update','2023-08-12 11:20:43',5,'更新之前的数据内容为：id=5,name=赵云,phone=15979765495,email=153@qq.com,profession=化学装备   更新之后的数据内容为：id=5,name=赵云,phone=15979765495,email=153@qq.com,profession=土木工程'),(9,'delete','2023-08-12 11:24:49',8,'删除之前的数据内容为：id=8,name=王老五,phone=12979765491,email=119@qq.com,profession=土木工程');
/*!40000 ALTER TABLE `user_logs` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-15 11:59:32
