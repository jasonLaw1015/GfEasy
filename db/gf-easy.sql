# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.31)
# Database: gf-easy
# Generation Time: 2021-09-16 04:28:00 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table app_goods_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `app_goods_info`;

CREATE TABLE `app_goods_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `subTitle` varchar(255) DEFAULT NULL COMMENT '副标题',
  `pic` varchar(255) NOT NULL COMMENT '商品主图',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `types` tinyint(4) NOT NULL COMMENT '类型#1:上架,2:下架',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态#1:启用,2:禁用',
  PRIMARY KEY (`id`),
  KEY `IDX_d8d0d86a0adf1001ce12aaac41` (`createTime`),
  KEY `IDX_9fe44721d1e517fa24383db56b` (`updateTime`),
  FULLTEXT KEY `IDX_91f50fa9907d5ac2c864f175bb` (`title`,`subTitle`) /*!50100 WITH PARSER `ngram` */ 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

LOCK TABLES `app_goods_info` WRITE;
/*!40000 ALTER TABLE `app_goods_info` DISABLE KEYS */;

INSERT INTO `app_goods_info` (`id`, `createTime`, `updateTime`, `title`, `subTitle`, `pic`, `sort`, `types`, `status`)
VALUES
	(3,'2021-08-05 17:54:57.000000','2021-09-02 16:29:41.000000','16546','13','http://127.0.0.1:8199/cdz9b9z5hjpkbfw5nf.png',12,1,2),
	(4,'2021-08-06 15:38:20.737766','2021-08-06 15:38:20.737766','hello','hello','http://127.0.0.1:8001/uploads/20210806/3caaa1a0-f689-11eb-b8a8-73ea0888018e.svg',10,1,1);

/*!40000 ALTER TABLE `app_goods_info` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_app_space_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_app_space_info`;

CREATE TABLE `base_app_space_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `url` varchar(255) NOT NULL COMMENT '地址',
  `type` varchar(255) NOT NULL COMMENT '类型',
  `classifyId` bigint(20) DEFAULT NULL COMMENT '分类ID',
  PRIMARY KEY (`id`),
  KEY `IDX_4aed04cbfa2ecdc01485b86e51` (`createTime`),
  KEY `IDX_abd5de4a4895eb253a5cabb20f` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='图片空间信息';

LOCK TABLES `base_app_space_info` WRITE;
/*!40000 ALTER TABLE `base_app_space_info` DISABLE KEYS */;

INSERT INTO `base_app_space_info` (`id`, `createTime`, `updateTime`, `url`, `type`, `classifyId`)
VALUES
	(1,'2021-08-03 22:35:55.900457','2021-08-03 22:35:55.900457','https://admin.cool-js.cool/uploads/20210803/1c9bf0d0-f468-11eb-92d0-bba305da9aaf.svg','image/svg+xml',NULL),
	(2,'2021-08-05 16:02:50.630508','2021-08-05 16:02:50.630508','http://127.0.0.1:8001/uploads/20210805/8782b490-f5c3-11eb-b751-19abbc7c44a6.svg','image/svg+xml',NULL),
	(3,'2021-08-05 21:31:26.364898','2021-08-05 21:31:26.364898','http://127.0.0.1:8001/uploads/20210805/6f01e250-f5f1-11eb-b197-135addbfa1fe.svg','image/svg+xml',NULL),
	(4,'2021-08-05 22:00:54.485070','2021-08-05 22:00:54.485070','http://127.0.0.1:8001/uploads/20210805/8ce3bce0-f5f5-11eb-b197-135addbfa1fe.svg','image/svg+xml',NULL),
	(5,'2021-08-06 15:38:05.403516','2021-08-06 15:38:05.403516','http://127.0.0.1:8001/uploads/20210806/3caaa1a0-f689-11eb-b8a8-73ea0888018e.svg','image/svg+xml',NULL),
	(6,'2021-09-02 16:29:38.000000','2021-09-02 16:29:38.000000','http://127.0.0.1:8199/cdz9b9z5hjpkbfw5nf.png','image/png',NULL),
	(7,'2021-09-06 11:26:25.000000','2021-09-06 11:26:25.000000','http://127.0.0.1:8199/ce2hday4j3js9qnemk.png','image/png',NULL);

/*!40000 ALTER TABLE `base_app_space_info` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_app_space_type
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_app_space_type`;

CREATE TABLE `base_app_space_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `name` varchar(255) NOT NULL COMMENT '类别名称',
  `parentId` tinyint(4) DEFAULT NULL COMMENT '父分类ID',
  PRIMARY KEY (`id`),
  KEY `IDX_5e8376603f89fdf3e7bb05103a` (`createTime`),
  KEY `IDX_500ea9e8b2c5c08c9b86a0667e` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='图片空间类型';

LOCK TABLES `base_app_space_type` WRITE;
/*!40000 ALTER TABLE `base_app_space_type` DISABLE KEYS */;

INSERT INTO `base_app_space_type` (`id`, `createTime`, `updateTime`, `name`, `parentId`)
VALUES
	(1,'2021-02-26 14:07:48.867045','2021-02-26 14:07:48.867045','a',NULL),
	(2,'2021-02-26 14:07:52.285531','2021-02-26 14:07:52.285531','b',NULL);

/*!40000 ALTER TABLE `base_app_space_type` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_comm
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_comm`;

CREATE TABLE `base_comm` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `IDX_905208f206a3ff9fd513421971` (`createTime`),
  KEY `IDX_4c6f27f6ecefe51a5a196a047a` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='常用接口-图片上传、用户信息及修改、权限信息';



# Dump of table base_open
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_open`;

CREATE TABLE `base_open` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `IDX_905208f206a3ff9fd513421971` (`createTime`),
  KEY `IDX_4c6f27f6ecefe51a5a196a047a` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='无需token的接口';



# Dump of table base_sys_conf
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_conf`;

CREATE TABLE `base_sys_conf` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `cKey` varchar(255) NOT NULL COMMENT '配置键',
  `cValue` varchar(255) NOT NULL COMMENT '配置值',
  PRIMARY KEY (`id`),
  UNIQUE KEY `IDX_9be195d27767b4485417869c3a` (`cKey`),
  KEY `IDX_905208f206a3ff9fd513421971` (`createTime`),
  KEY `IDX_4c6f27f6ecefe51a5a196a047a` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统配置';

LOCK TABLES `base_sys_conf` WRITE;
/*!40000 ALTER TABLE `base_sys_conf` DISABLE KEYS */;

INSERT INTO `base_sys_conf` (`id`, `createTime`, `updateTime`, `cKey`, `cValue`)
VALUES
	(1,'2021-02-25 14:23:26.810981','2021-08-24 11:54:04.000000','logKeep','1');

/*!40000 ALTER TABLE `base_sys_conf` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_sys_department
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_department`;

CREATE TABLE `base_sys_department` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `name` varchar(255) NOT NULL COMMENT '部门名称',
  `parentId` bigint(20) DEFAULT NULL COMMENT '上级部门ID',
  `orderNum` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `IDX_be4c53cd671384fa588ca9470a` (`createTime`),
  KEY `IDX_ca1473a793961ec55bc0c8d268` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='admin部门';

LOCK TABLES `base_sys_department` WRITE;
/*!40000 ALTER TABLE `base_sys_department` DISABLE KEYS */;

INSERT INTO `base_sys_department` (`id`, `createTime`, `updateTime`, `name`, `parentId`, `orderNum`)
VALUES
	(1,'2021-02-24 21:17:11.971397','2021-09-08 22:47:14.605936','GoEasyTeam',NULL,0),
	(11,'2021-02-26 14:17:06.690613','2021-02-26 14:17:06.690613','开发',1,0),
	(12,'2021-02-26 14:17:11.576369','2021-02-26 14:17:11.576369','测试',1,0),
	(13,'2021-02-26 14:28:59.685177','2021-02-26 14:28:59.685177','游客',1,0);

/*!40000 ALTER TABLE `base_sys_department` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_sys_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_log`;

CREATE TABLE `base_sys_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `adminUserId` bigint(20) DEFAULT NULL COMMENT '用户ID',
  `action` varchar(100) NOT NULL COMMENT '行为',
  `ip` varchar(50) DEFAULT NULL COMMENT 'ip',
  `ipAddr` varchar(50) DEFAULT NULL COMMENT 'ip地址',
  `params` text COMMENT '参数',
  PRIMARY KEY (`id`),
  KEY `IDX_51a2caeb5713efdfcb343a8772` (`adminUserId`),
  KEY `IDX_938f886fb40e163db174b7f6c3` (`action`),
  KEY `IDX_24e18767659f8c7142580893f2` (`ip`),
  KEY `IDX_a03a27f75cf8d502b3060823e1` (`ipAddr`),
  KEY `IDX_c9382b76219a1011f7b8e7bcd1` (`createTime`),
  KEY `IDX_bfd44e885b470da43bcc39aaa7` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='admin日志';

LOCK TABLES `base_sys_log` WRITE;
/*!40000 ALTER TABLE `base_sys_log` DISABLE KEYS */;

INSERT INTO `base_sys_log` (`id`, `createTime`, `updateTime`, `adminUserId`, `action`, `ip`, `ipAddr`, `params`)
VALUES
	(9107,'2021-09-09 21:22:10.000000','2021-09-09 21:22:10.000000',1,'/admin/baseSysLog/page','127.0.0.1','127.0.0.1:61284','{\"page\":1,\"size\":20,\"sort\":\"desc\",\"order\":\"createTime\"}'),
	(9108,'2021-09-09 21:22:15.000000','2021-09-09 21:22:15.000000',1,'/admin/appGoodsInfo/page','127.0.0.1','127.0.0.1:61307','{\"page\":1,\"size\":20}'),
	(9109,'2021-09-09 21:22:16.000000','2021-09-09 21:22:16.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:61316','{\"page\":1,\"size\":20}'),
	(9110,'2021-09-09 21:23:12.000000','2021-09-09 21:23:12.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:61575',NULL),
	(9111,'2021-09-09 21:23:12.000000','2021-09-09 21:23:12.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:61582',NULL),
	(9112,'2021-09-09 21:23:13.000000','2021-09-09 21:23:13.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:61583','{\"page\":1,\"size\":20}'),
	(9113,'2021-09-09 21:23:31.000000','2021-09-09 21:23:31.000000',1,'/admin/genCodeConfig/firstInfo','127.0.0.1','127.0.0.1:61658',NULL),
	(9114,'2021-09-09 23:29:55.000000','2021-09-09 23:29:55.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:59980',NULL),
	(9115,'2021-09-09 23:29:55.000000','2021-09-09 23:29:55.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:59983',NULL),
	(9116,'2021-09-09 23:29:56.000000','2021-09-09 23:29:56.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:59985','{\"page\":1,\"size\":20}'),
	(9117,'2021-09-12 11:56:43.000000','2021-09-12 11:56:43.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:53665',NULL),
	(9118,'2021-09-12 11:56:43.000000','2021-09-12 11:56:43.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:53668',NULL),
	(9119,'2021-09-12 11:57:19.000000','2021-09-12 11:57:19.000000',1,'/admin/baseComm/uploadMode','127.0.0.1','127.0.0.1:53686',NULL),
	(9120,'2021-09-12 11:57:19.000000','2021-09-12 11:57:19.000000',1,'/admin/baseComm/upload','127.0.0.1','127.0.0.1:53691',NULL),
	(9121,'2021-09-12 11:57:20.000000','2021-09-12 11:57:20.000000',1,'/admin/baseComm/personUpdate','127.0.0.1','127.0.0.1:53694','{\"headImg\":\"http://127.0.0.1:8199/ce7ls8387c6ohcrmdk.png\",\"nickName\":\"管理员\",\"nowPassword\":\"\",\"newPassword\":\"\"}'),
	(9122,'2021-09-12 11:57:29.000000','2021-09-12 11:57:29.000000',1,'/admin/baseComm/personUpdate','127.0.0.1','127.0.0.1:53697','{\"headImg\":\"http://127.0.0.1:8199/ce7ls8387c6ohcrmdk.png\",\"nickName\":\"管理员\",\"nowPassword\":\"123456\",\"newPassword\":\"123456\"}'),
	(9123,'2021-09-12 11:57:29.000000','2021-09-12 11:57:29.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:53700',NULL),
	(9124,'2021-09-12 11:57:37.000000','2021-09-12 11:57:37.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:53719',NULL),
	(9125,'2021-09-12 11:57:37.000000','2021-09-12 11:57:37.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:53720',NULL),
	(9126,'2021-09-12 11:58:09.000000','2021-09-12 11:58:09.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:53863',NULL),
	(9127,'2021-09-12 11:58:09.000000','2021-09-12 11:58:09.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:53866',NULL),
	(9128,'2021-09-12 11:58:16.000000','2021-09-12 11:58:16.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:53909',NULL),
	(9129,'2021-09-12 11:58:16.000000','2021-09-12 11:58:16.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:53910',NULL),
	(9130,'2021-09-12 11:58:33.000000','2021-09-12 11:58:33.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:53995',NULL),
	(9131,'2021-09-12 11:58:33.000000','2021-09-12 11:58:33.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:53997',NULL),
	(9132,'2021-09-12 11:58:41.000000','2021-09-12 11:58:41.000000',1,'/admin/baseSysUser/page','127.0.0.1','127.0.0.1:54035','{\"page\":1,\"size\":20,\"sort\":\"desc\",\"order\":\"createTime\"}'),
	(9133,'2021-09-12 11:58:41.000000','2021-09-12 11:58:41.000000',1,'/admin/baseSysDepartment/list','127.0.0.1','127.0.0.1:54034','{}'),
	(9134,'2021-09-12 11:58:45.000000','2021-09-12 11:58:45.000000',1,'/admin/baseSysMenu/list','127.0.0.1','127.0.0.1:54053','{}'),
	(9135,'2021-09-12 12:02:27.000000','2021-09-12 12:02:27.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:55122',NULL),
	(9136,'2021-09-12 12:02:27.000000','2021-09-12 12:02:27.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:55125',NULL),
	(9137,'2021-09-12 12:03:19.000000','2021-09-12 12:03:19.000000',1,'/admin/taskInfo/page','127.0.0.1','127.0.0.1:55348','{\"types\":2,\"status\":1,\"size\":10,\"page\":1,\"total\":0}'),
	(9138,'2021-09-12 12:03:19.000000','2021-09-12 12:03:19.000000',1,'/admin/taskInfo/page','127.0.0.1','127.0.0.1:55350','{\"types\":null,\"status\":2,\"size\":10,\"page\":1,\"total\":0}'),
	(9139,'2021-09-12 12:03:19.000000','2021-09-12 12:03:19.000000',1,'/admin/taskInfo/getTaskScheduleMethods','127.0.0.1','127.0.0.1:55351',NULL),
	(9140,'2021-09-12 12:03:19.000000','2021-09-12 12:03:19.000000',1,'/admin/taskInfo/log','127.0.0.1','127.0.0.1:55353','{\"page\":\"1\",\"size\":\"10\"}'),
	(9141,'2021-09-12 12:03:19.000000','2021-09-12 12:03:19.000000',1,'/admin/taskInfo/page','127.0.0.1','127.0.0.1:55349','{\"types\":1,\"status\":1,\"size\":10,\"page\":1,\"total\":0}'),
	(9142,'2021-09-12 12:03:23.000000','2021-09-12 12:03:23.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:55382','{\"page\":1,\"size\":20}'),
	(9143,'2021-09-12 12:03:34.000000','2021-09-12 12:03:34.000000',1,'/admin/baseSysParam/page','127.0.0.1','127.0.0.1:55429','{\"page\":1,\"size\":20}'),
	(9144,'2021-09-12 12:03:37.000000','2021-09-12 12:03:37.000000',1,'/admin/baseSysLog/page','127.0.0.1','127.0.0.1:55447','{\"page\":1,\"size\":20,\"sort\":\"desc\",\"order\":\"createTime\"}'),
	(9145,'2021-09-12 12:03:37.000000','2021-09-12 12:03:37.000000',1,'/admin/baseSysLog/getKeep','127.0.0.1','127.0.0.1:55448',NULL),
	(9146,'2021-09-12 12:29:23.000000','2021-09-12 12:29:23.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:61772','{\"page\":1,\"size\":20}'),
	(9147,'2021-09-12 12:29:26.000000','2021-09-12 12:29:26.000000',1,'/admin/genCodeConfig/firstInfo','127.0.0.1','127.0.0.1:61789',NULL),
	(9148,'2021-09-12 12:29:33.000000','2021-09-12 12:29:33.000000',1,'/admin/genCodeConfig/firstInfo','127.0.0.1','127.0.0.1:61820',NULL),
	(9149,'2021-09-12 12:29:36.000000','2021-09-12 12:29:36.000000',1,'/admin/genCodeConfig/firstInfo','127.0.0.1','127.0.0.1:61840',NULL),
	(9150,'2021-09-12 12:29:46.000000','2021-09-12 12:29:46.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:61906',NULL),
	(9151,'2021-09-12 12:29:46.000000','2021-09-12 12:29:46.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:61908',NULL),
	(9152,'2021-09-12 12:29:46.000000','2021-09-12 12:29:46.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:61909',NULL),
	(9153,'2021-09-12 12:29:46.000000','2021-09-12 12:29:46.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:61910',NULL),
	(9154,'2021-09-12 12:29:47.000000','2021-09-12 12:29:47.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:61918','{\"page\":1,\"size\":20}'),
	(9155,'2021-09-12 12:29:58.000000','2021-09-12 12:29:58.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:61986',NULL),
	(9156,'2021-09-12 12:29:58.000000','2021-09-12 12:29:58.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:61989',NULL),
	(9157,'2021-09-12 12:29:59.000000','2021-09-12 12:29:59.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:61994','{\"page\":1,\"size\":20}'),
	(9158,'2021-09-12 12:29:59.000000','2021-09-12 12:29:59.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:61995',NULL),
	(9159,'2021-09-12 12:29:59.000000','2021-09-12 12:29:59.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:61996',NULL),
	(9160,'2021-09-12 14:49:57.000000','2021-09-12 14:49:57.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:64634','{\"page\":1,\"size\":20}'),
	(9161,'2021-09-12 14:49:59.000000','2021-09-12 14:49:59.000000',1,'/admin/genCodeConfig/firstInfo','127.0.0.1','127.0.0.1:64651',NULL),
	(9162,'2021-09-12 14:50:09.000000','2021-09-12 14:50:09.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:64724',NULL),
	(9163,'2021-09-12 14:50:09.000000','2021-09-12 14:50:09.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:64726',NULL),
	(9164,'2021-09-12 14:50:10.000000','2021-09-12 14:50:10.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:64727',NULL),
	(9165,'2021-09-12 14:50:10.000000','2021-09-12 14:50:10.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:64728',NULL),
	(9166,'2021-09-12 14:50:10.000000','2021-09-12 14:50:10.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:64736','{\"page\":1,\"size\":20}'),
	(9167,'2021-09-12 14:50:10.000000','2021-09-12 14:50:10.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:64737','{\"page\":1,\"size\":20}'),
	(9168,'2021-09-12 14:50:21.000000','2021-09-12 14:50:21.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:64808',NULL),
	(9169,'2021-09-12 14:50:21.000000','2021-09-12 14:50:21.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:64809',NULL),
	(9170,'2021-09-12 14:50:22.000000','2021-09-12 14:50:22.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:64813',NULL),
	(9171,'2021-09-12 14:50:22.000000','2021-09-12 14:50:22.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:64814',NULL),
	(9172,'2021-09-12 14:50:22.000000','2021-09-12 14:50:22.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:64817','{\"page\":1,\"size\":20}'),
	(9173,'2021-09-12 14:50:22.000000','2021-09-12 14:50:22.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:64822','{\"page\":1,\"size\":20}'),
	(9174,'2021-09-16 11:38:16.000000','2021-09-16 11:38:16.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:51554',NULL),
	(9175,'2021-09-16 11:38:16.000000','2021-09-16 11:38:16.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:51557',NULL),
	(9176,'2021-09-16 11:39:25.000000','2021-09-16 11:39:25.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:51884',NULL),
	(9177,'2021-09-16 11:39:25.000000','2021-09-16 11:39:25.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:51887',NULL),
	(9178,'2021-09-16 11:42:30.000000','2021-09-16 11:42:30.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:52688',NULL),
	(9179,'2021-09-16 11:42:30.000000','2021-09-16 11:42:30.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:52693',NULL),
	(9180,'2021-09-16 11:42:40.000000','2021-09-16 11:42:40.000000',1,'/admin/baseSysDepartment/list','127.0.0.1','127.0.0.1:52751','{}'),
	(9181,'2021-09-16 11:42:40.000000','2021-09-16 11:42:40.000000',1,'/admin/baseSysUser/page','127.0.0.1','127.0.0.1:52752','{\"page\":1,\"size\":20,\"sort\":\"desc\",\"order\":\"createTime\"}'),
	(9182,'2021-09-16 11:42:41.000000','2021-09-16 11:42:41.000000',1,'/admin/baseSysMenu/list','127.0.0.1','127.0.0.1:52761','{}'),
	(9183,'2021-09-16 11:42:42.000000','2021-09-16 11:42:42.000000',1,'/admin/baseSysRole/page','127.0.0.1','127.0.0.1:52768','{\"page\":1,\"size\":20,\"sort\":\"desc\",\"order\":\"createTime\"}'),
	(9184,'2021-09-16 11:42:44.000000','2021-09-16 11:42:44.000000',1,'/admin/baseSysLog/getKeep','127.0.0.1','127.0.0.1:52782',NULL),
	(9185,'2021-09-16 11:42:44.000000','2021-09-16 11:42:44.000000',1,'/admin/baseSysLog/page','127.0.0.1','127.0.0.1:52783','{\"page\":1,\"size\":20,\"sort\":\"desc\",\"order\":\"createTime\"}'),
	(9186,'2021-09-16 11:42:45.000000','2021-09-16 11:42:45.000000',1,'/admin/taskInfo/page','127.0.0.1','127.0.0.1:52799','{\"types\":2,\"status\":1,\"size\":10,\"page\":1,\"total\":0}'),
	(9187,'2021-09-16 11:42:45.000000','2021-09-16 11:42:45.000000',1,'/admin/taskInfo/getTaskScheduleMethods','127.0.0.1','127.0.0.1:52796',NULL),
	(9188,'2021-09-16 11:42:45.000000','2021-09-16 11:42:45.000000',1,'/admin/taskInfo/page','127.0.0.1','127.0.0.1:52801','{\"types\":null,\"status\":2,\"size\":10,\"page\":1,\"total\":0}'),
	(9189,'2021-09-16 11:42:45.000000','2021-09-16 11:42:45.000000',1,'/admin/taskInfo/page','127.0.0.1','127.0.0.1:52800','{\"types\":1,\"status\":1,\"size\":10,\"page\":1,\"total\":0}'),
	(9190,'2021-09-16 11:42:45.000000','2021-09-16 11:42:45.000000',1,'/admin/taskInfo/log','127.0.0.1','127.0.0.1:52802','{\"page\":\"1\",\"size\":\"10\"}'),
	(9191,'2021-09-16 11:42:45.000000','2021-09-16 11:42:45.000000',1,'/admin/taskInfo/log','127.0.0.1','127.0.0.1:52815','{\"page\":\"2\",\"size\":\"10\",\"total\":\"737\"}'),
	(9192,'2021-09-16 11:51:05.000000','2021-09-16 11:51:05.000000',1,'/admin/baseComm/logout','127.0.0.1','127.0.0.1:54949',NULL),
	(9193,'2021-09-16 11:56:51.000000','2021-09-16 11:56:51.000000',1,'/admin/baseComm/person','127.0.0.1','127.0.0.1:56649',NULL),
	(9194,'2021-09-16 11:56:51.000000','2021-09-16 11:56:51.000000',1,'/admin/baseComm/permmenu','127.0.0.1','127.0.0.1:56652',NULL),
	(9195,'2021-09-16 11:57:00.000000','2021-09-16 11:57:00.000000',1,'/admin/appGoodsInfo/page','127.0.0.1','127.0.0.1:56699','{\"page\":1,\"size\":20}'),
	(9196,'2021-09-16 11:57:00.000000','2021-09-16 11:57:00.000000',1,'/admin/demoGo/page','127.0.0.1','127.0.0.1:56710','{\"page\":1,\"size\":20}'),
	(9197,'2021-09-16 11:57:00.000000','2021-09-16 11:57:00.000000',1,'/admin/appGoodsInfo/list','127.0.0.1','127.0.0.1:56709','{}'),
	(9198,'2021-09-16 11:57:02.000000','2021-09-16 11:57:02.000000',1,'/admin/demoGo/info','127.0.0.1','127.0.0.1:56725','{\"id\":\"2\"}'),
	(9199,'2021-09-16 11:57:42.000000','2021-09-16 11:57:42.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:56892','{\"page\":1,\"size\":20}'),
	(9200,'2021-09-16 12:02:21.000000','2021-09-16 12:02:21.000000',1,'/admin/appGoodsInfo/list','127.0.0.1','127.0.0.1:58065','{}'),
	(9201,'2021-09-16 12:02:21.000000','2021-09-16 12:02:21.000000',1,'/admin/demoGo/page','127.0.0.1','127.0.0.1:58066','{\"page\":1,\"size\":20}'),
	(9202,'2021-09-16 12:02:55.000000','2021-09-16 12:02:55.000000',1,'/admin/baseSysUser/page','127.0.0.1','127.0.0.1:58216','{\"page\":1,\"size\":20,\"sort\":\"desc\",\"order\":\"createTime\"}'),
	(9203,'2021-09-16 12:02:55.000000','2021-09-16 12:02:55.000000',1,'/admin/baseSysDepartment/list','127.0.0.1','127.0.0.1:58215','{}'),
	(9204,'2021-09-16 12:02:55.000000','2021-09-16 12:02:55.000000',1,'/admin/baseSysMenu/list','127.0.0.1','127.0.0.1:58223','{}'),
	(9205,'2021-09-16 12:19:21.000000','2021-09-16 12:19:21.000000',1,'/admin/genCodeLog/page','127.0.0.1','127.0.0.1:62370','{\"page\":1,\"size\":20}'),
	(9206,'2021-09-16 12:19:21.000000','2021-09-16 12:19:21.000000',1,'/admin/genCodeConfig/firstInfo','127.0.0.1','127.0.0.1:62375',NULL),
	(9207,'2021-09-16 12:21:15.000000','2021-09-16 12:21:15.000000',1,'/admin/genCodeConfig/save','127.0.0.1','127.0.0.1:62844','{\"dbHost\":\"127.0.0.1\",\"dbPort\":\"3306\",\"dbUser\":\"root\",\"dbPass\":\"123456\",\"dbName\":\"go-easy\",\"dbType\":\"mysql\",\"adminGenPath\":\"/Users/yons/Documents/develop/go/src/gitee.com/jasonLaw1015/GfEasyAdmin/\",\"serverGenPath\":\"/Users/yons/Documents/develop/go/src/gitee.com/jasonLaw1015/GfEasy/\",\"defaultCreatedAtLabel\":\"createTime\",\"defaultUpdatedAtLabel\":\"updateTime\",\"ignoreTablePrefix\":\"u_,base_\",\"activeCode\":\"\",\"id\":1}'),
	(9208,'2021-09-16 12:21:16.000000','2021-09-16 12:21:16.000000',1,'/admin/genCodeConfig/firstInfo','127.0.0.1','127.0.0.1:62858',NULL),
	(9209,'2021-09-16 12:27:46.000000','2021-09-16 12:27:46.000000',1,'/admin/genCodeConfig/firstInfo','127.0.0.1','127.0.0.1:64463',NULL);

/*!40000 ALTER TABLE `base_sys_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_sys_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_menu`;

CREATE TABLE `base_sys_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `parentId` bigint(20) DEFAULT NULL COMMENT '父菜单ID',
  `name` varchar(255) NOT NULL COMMENT '菜单名称',
  `router` varchar(255) DEFAULT NULL COMMENT '菜单地址',
  `perms` text COMMENT '权限标识',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型 0：目录 1：菜单 2：按钮',
  `icon` varchar(255) DEFAULT NULL COMMENT '图标',
  `orderNum` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `viewPath` varchar(255) DEFAULT NULL COMMENT '视图地址',
  `keepAlive` tinyint(4) NOT NULL DEFAULT '1' COMMENT '路由缓存',
  `isShow` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否显示#0:不显示,1:显示',
  PRIMARY KEY (`id`),
  KEY `IDX_05e3d6a56604771a6da47ebf8e` (`createTime`),
  KEY `IDX_d5203f18daaf7c3fe0ab34497f` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='admin菜单';

LOCK TABLES `base_sys_menu` WRITE;
/*!40000 ALTER TABLE `base_sys_menu` DISABLE KEYS */;

INSERT INTO `base_sys_menu` (`id`, `createTime`, `updateTime`, `parentId`, `name`, `router`, `perms`, `type`, `icon`, `orderNum`, `viewPath`, `keepAlive`, `isShow`)
VALUES
	(1,'2019-09-11 11:14:44.000000','2019-11-18 15:56:36.000000',NULL,'工作台','/',NULL,0,'icon-workbench',1,NULL,1,1),
	(2,'2019-09-11 11:14:47.000000','2021-02-27 17:16:05.000000',NULL,'系统管理','/sys',NULL,0,'icon-system',2,NULL,1,1),
	(8,'1900-01-20 23:19:57.000000','2021-03-08 22:59:12.000000',27,'菜单列表','/sys/menu',NULL,1,'icon-menu',2,'cool/modules/base/views/menu.vue',1,1),
	(10,'1900-01-20 00:19:27.325000','2021-08-21 16:26:55.000000',8,'新增',NULL,'baseSysMenu:add',2,NULL,1,NULL,0,1),
	(11,'1900-01-20 00:19:51.101000','2021-08-21 16:26:55.000000',8,'删除',NULL,'baseSysMenu:delete',2,NULL,2,NULL,0,1),
	(12,'1900-01-20 00:20:05.150000','2021-08-21 16:26:55.000000',8,'修改',NULL,'baseSysMenu:update',2,NULL,3,NULL,0,1),
	(13,'1900-01-20 00:20:19.341000','2021-08-21 16:26:55.000000',8,'查询',NULL,'baseSysMenu:page,baseSysMenu:list,baseSysMenu:info',2,NULL,4,NULL,0,1),
	(22,'2019-09-12 00:34:01.000000','2021-03-08 22:59:23.000000',27,'角色列表','/sys/role',NULL,1,'icon-common',3,'cool/modules/base/views/role.vue',1,1),
	(23,'1900-01-20 00:34:23.459000','2021-08-21 16:26:55.000000',22,'新增',NULL,'baseSysRole:add',2,NULL,1,NULL,0,1),
	(24,'1900-01-20 00:34:40.523000','2021-08-21 16:26:55.000000',22,'删除',NULL,'baseSysRole:delete',2,NULL,2,NULL,0,1),
	(25,'1900-01-20 00:34:53.306000','2021-08-21 16:26:55.000000',22,'修改',NULL,'baseSysRole:update',2,NULL,3,NULL,0,1),
	(26,'1900-01-20 00:35:05.024000','2021-08-21 16:26:55.000000',22,'查询',NULL,'baseSysRole:page,baseSysRole:list,baseSysRole:info',2,NULL,4,NULL,0,1),
	(27,'2019-09-12 15:52:44.000000','2019-09-15 22:11:56.000000',2,'权限管理',NULL,NULL,0,'icon-auth',1,NULL,0,1),
	(29,'2019-09-12 17:35:51.000000','2021-03-08 23:01:39.000000',105,'请求日志','/sys/log',NULL,1,'icon-log',1,'cool/modules/base/views/log.vue',1,1),
	(30,'2019-09-12 17:37:03.000000','2021-08-21 16:26:55.000000',29,'权限',NULL,'baseSysLog:page,baseSysLog:clear,baseSysLog:getKeep,baseSysLog:setKeep',2,NULL,1,NULL,0,1),
	(43,'2019-11-07 14:22:34.000000','2021-03-08 23:02:51.000000',45,'crud 示例','/crud',NULL,1,'icon-favor',1,'cool/modules/demo/views/crud.vue',1,1),
	(45,'2019-11-07 22:36:57.000000','2019-11-11 15:21:10.000000',1,'组件库','/ui-lib',NULL,0,'icon-common',2,NULL,1,1),
	(47,'2019-11-08 09:35:08.000000','2021-08-07 16:09:42.000000',NULL,'框架教程','/tutorial',NULL,0,'icon-task',4,NULL,1,0),
	(48,'2019-11-08 09:35:53.000000','2021-03-03 11:03:21.000000',47,'文档','/tutorial/doc',NULL,1,'icon-log',0,'https://admin.cool-js.com',1,1),
	(49,'2019-11-09 22:11:13.000000','2021-03-09 09:50:46.000000',45,'quill 富文本编辑器','/editor-quill',NULL,1,'icon-favor',2,'cool/modules/demo/views/editor-quill.vue',1,1),
	(59,'2019-11-18 16:50:27.000000','2021-08-21 16:26:55.000000',97,'部门列表',NULL,'baseSysDepartment:list',2,NULL,0,NULL,1,1),
	(60,'2019-11-18 16:50:45.000000','2021-08-21 16:26:55.000000',97,'新增部门',NULL,'baseSysDepartment:add',2,NULL,0,NULL,1,1),
	(61,'2019-11-18 16:50:59.000000','2021-08-21 16:26:55.000000',97,'更新部门',NULL,'baseSysDepartment:update',2,NULL,0,NULL,1,1),
	(62,'2019-11-18 16:51:13.000000','2021-08-21 16:26:55.000000',97,'删除部门',NULL,'baseSysDepartment:delete',2,NULL,0,NULL,1,1),
	(63,'2019-11-18 17:49:35.000000','2021-08-21 16:26:55.000000',97,'部门排序',NULL,'baseSysDepartment:order',2,NULL,0,NULL,1,1),
	(65,'2019-11-18 23:59:21.000000','2021-08-21 16:26:55.000000',97,'用户转移',NULL,'baseSysUser:move',2,NULL,0,NULL,1,1),
	(78,'2019-12-10 13:27:56.000000','2021-02-27 17:08:53.000000',2,'参数配置',NULL,NULL,0,'icon-common',4,NULL,1,1),
	(79,'1900-01-20 13:29:33.000000','2021-03-08 23:01:48.000000',78,'参数列表','/sys/param',NULL,1,'icon-menu',0,'cool/modules/base/views/param.vue',1,1),
	(80,'1900-01-20 13:29:50.146000','2021-08-21 16:26:55.000000',79,'新增',NULL,'baseSysParam:add',2,NULL,0,NULL,1,1),
	(81,'1900-01-20 13:30:10.030000','2021-08-21 16:26:55.000000',79,'修改',NULL,'baseSysParam:info,baseSysParam:update',2,NULL,0,NULL,1,1),
	(82,'1900-01-20 13:30:25.791000','2021-08-21 16:26:55.000000',79,'删除',NULL,'baseSysParam:delete',2,NULL,0,NULL,1,1),
	(83,'1900-01-20 13:30:40.469000','2021-08-21 16:26:55.000000',79,'查看',NULL,'baseSysParam:page,baseSysParam:list,baseSysParam:info',2,NULL,0,NULL,1,1),
	(84,'2020-07-25 16:21:30.000000','2020-07-25 16:21:30.000000',NULL,'通用',NULL,NULL,0,'icon-radioboxfill',99,NULL,1,0),
	(85,'2020-07-25 16:22:14.000000','2021-09-08 23:02:40.000000',84,'图片空间权限','','baseAppSpaceInfo:page,baseAppSpaceInfo:list,baseAppSpaceInfo:info,baseAppSpaceInfo:add,baseAppSpaceInfo:delete,baseAppSpaceInfo:update,baseAppSpaceType:page,baseAppSpaceType:list,baseAppSpaceType:info,baseAppSpaceType:add,baseAppSpaceType:delete,baseAppSpaceType:update',2,'',1,'',0,0),
	(86,'2020-08-12 09:56:27.000000','2021-03-08 23:03:03.000000',45,'文件上传','/upload',NULL,1,'icon-favor',3,'cool/modules/demo/views/upload.vue',1,1),
	(90,'1900-01-20 10:26:58.615000','2021-08-21 16:26:55.000000',84,'客服聊天',NULL,'baseAppImMessage:read,baseAppImMessage:page,baseAppImSession:page,baseAppImSession:list,baseAppImSession:unreadCount,baseAppImSession:delete',2,NULL,0,NULL,1,1),
	(96,'2021-01-12 14:12:20.000000','2021-03-08 23:02:40.000000',1,'组件预览','/demo',NULL,1,'icon-favor',0,'cool/modules/demo/views/demo.vue',1,1),
	(97,'1900-01-20 14:14:02.000000','2021-03-09 11:03:09.000000',27,'用户列表','/sys/user',NULL,1,'icon-user',0,'cool/modules/base/views/user.vue',1,1),
	(98,'1900-01-20 14:14:13.528000','2021-08-21 16:26:55.000000',97,'新增',NULL,'baseSysUser:add',2,NULL,0,NULL,1,1),
	(99,'1900-01-20 14:14:22.823000','2021-08-21 16:26:55.000000',97,'删除',NULL,'baseSysUser:delete',2,NULL,0,NULL,1,1),
	(100,'1900-01-20 14:14:33.973000','2021-08-21 16:26:55.000000',97,'修改',NULL,'baseSysUser:delete,baseSysUser:update',2,NULL,0,NULL,1,1),
	(101,'2021-01-12 14:14:51.000000','2021-08-21 16:26:55.000000',97,'查询',NULL,'baseSysUser:page,baseSysUser:list,baseSysUser:info',2,NULL,0,NULL,1,1),
	(105,'2021-01-21 10:42:55.000000','2021-01-21 10:42:55.000000',2,'监控管理',NULL,NULL,0,'icon-rank',6,NULL,1,1),
	(109,'2021-02-27 14:13:56.000000','2021-08-07 16:09:36.000000',NULL,'插件管理',NULL,NULL,0,'icon-menu',3,NULL,1,0),
	(110,'2021-02-27 14:14:13.000000','2021-03-08 23:01:30.000000',109,'插件列表','/plugin',NULL,1,'icon-menu',0,'cool/modules/base/views/plugin.vue',1,1),
	(111,'2021-02-27 14:24:41.877000','2021-08-21 16:26:55.000000',110,'编辑',NULL,'basePluginInfo:info,basePluginInfo:update',2,NULL,0,NULL,1,1),
	(112,'2021-02-27 14:24:52.159000','2021-08-21 16:26:55.000000',110,'列表',NULL,'basePluginInfo:list',2,NULL,0,NULL,1,1),
	(113,'2021-02-27 14:25:02.066000','2021-08-21 16:26:55.000000',110,'删除',NULL,'basePluginInfo:delete',2,NULL,0,NULL,1,1),
	(114,'2021-02-27 16:36:59.322000','2021-08-21 16:26:55.000000',110,'保存配置',NULL,'basePluginInfo:config',2,NULL,0,NULL,1,1),
	(115,'2021-02-27 16:38:21.000000','2021-08-21 16:26:55.000000',110,'获取配置',NULL,'basePluginInfo:getConfig',2,NULL,0,NULL,1,1),
	(116,'2021-02-27 17:57:42.000000','2021-08-21 16:26:55.000000',110,'开启、关闭',NULL,'basePluginInfo:enable',2,NULL,0,NULL,1,1),
	(117,'2021-03-05 10:58:25.000000','2021-03-05 10:58:25.000000',NULL,'任务管理',NULL,NULL,0,'icon-activity',5,NULL,1,1),
	(118,'2021-03-05 10:59:42.000000','2021-03-05 10:59:42.000000',117,'任务列表','/task',NULL,1,'icon-menu',0,'cool/modules/task/views/task.vue',1,1),
	(119,'2021-03-05 11:00:00.000000','2021-08-21 16:26:55.000000',118,'权限',NULL,'taskInfo:page,taskInfo:list,taskInfo:info,taskInfo:add,taskInfo:delete,taskInfo:update,taskInfo:stop,taskInfo:start,taskInfo:once,taskInfo:log',2,NULL,0,NULL,1,1),
	(120,'2021-08-03 21:31:03.924072','2021-08-26 17:31:09.000000',NULL,'代码模块区',NULL,NULL,0,'icon-dept',99,NULL,1,1),
	(121,'2021-08-03 21:31:45.000000','2021-08-23 20:32:42.000000',120,'appGoodsInfo','/appGoodsInfo',NULL,1,'icon-app',22,'cool/modules/genCode/views/appGoodsInfo.vue',1,1),
	(122,'2021-08-03 21:48:45.126203','2021-08-21 16:28:25.000000',121,'goodsInfo基本权限',NULL,'appGoodsInfo:page,appGoodsInfo:list,appGoodsInfo:info,appGoodsInfo:add,appGoodsInfo:delete,appGoodsInfo:update',2,NULL,12,NULL,1,1),
	(127,'2021-08-23 14:48:46.471687','2021-08-23 14:48:46.477268',120,'示例go','/demoGo',NULL,1,'icon-app',127,'cool/modules/genCode/views/demoGo.vue',1,1),
	(128,'2021-08-23 14:48:46.481775','2021-08-23 14:48:46.481775',127,'示例go基本权限',NULL,'demoGo:page,demoGo:list,demoGo:info,demoGo:add,demoGo:delete,demoGo:update',2,'icon-app',100,'cool/modules/genCode/views/demoGo.vue',1,1),
	(131,'2021-08-26 17:30:34.000000','2021-08-26 17:30:34.000000',NULL,'核心工具',NULL,NULL,0,'icon-workbench',99,NULL,1,1),
	(133,'2021-08-26 14:43:54.840643','2021-08-26 22:46:41.000000',131,'代码生成工具','/genCodeLog',NULL,1,'icon-app',99,'cool/modules/genCode/views/genCodeLog.vue',1,1),
	(134,'2021-08-26 14:43:54.852175','2021-08-26 14:43:54.852175',133,'代码生成日志基本权限',NULL,'genCodeLog:page,genCodeLog:list,genCodeLog:info,genCodeLog:add,genCodeLog:delete,genCodeLog:update',2,'icon-app',100,'cool/modules/genCode/views/genCodeLog.vue',1,1),
	(137,'2021-08-27 04:37:33.311735','2021-09-02 16:30:55.000000',131,'代码生成配置','/genCodeConfig','',1,'icon-app',99,'cool/modules/genCode/views/genCodeConfig.vue',1,0),
	(138,'2021-08-27 04:37:33.320918','2021-08-27 04:37:33.320918',137,'代码生成配置基本权限',NULL,'genCodeConfig:page,genCodeConfig:list,genCodeConfig:info,genCodeConfig:add,genCodeConfig:delete,genCodeConfig:update',2,'icon-app',100,'cool/modules/genCode/views/genCodeConfig.vue',1,1),
	(139,'2021-09-08 23:03:16.000000','2021-09-08 23:03:16.000000',84,'comm权限+图片上传权限','','baseComm:page,baseComm:list,baseComm:info,baseComm:add,baseComm:delete,baseComm:update,baseComm:uploadMode,baseComm:upload,baseComm:logout,baseComm:person,baseComm:personUpdate,baseComm:permmenu,baseOpen:page,baseOpen:list,baseOpen:info,baseOpen:add,baseOpen:delete,baseOpen:update',2,'',1,'',0,0),
	(140,'2021-09-08 23:14:02.000000','2021-09-08 23:14:02.000000',84,'生成代码的权限','','genCodeConfig:page,genCodeConfig:list,genCodeConfig:info,genCodeConfig:add,genCodeConfig:delete,genCodeConfig:update,genCodeConfig:firstInfo,genCodeConfig:save',2,'',12,'',0,0);

/*!40000 ALTER TABLE `base_sys_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_sys_param
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_param`;

CREATE TABLE `base_sys_param` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `keyName` varchar(255) NOT NULL COMMENT '键位',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `data` text NOT NULL COMMENT '数据',
  `dataType` tinyint(4) NOT NULL DEFAULT '0' COMMENT '数据类型 0:字符串 1：数组 2：键值对',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `IDX_cf19b5e52d8c71caa9c4534454` (`keyName`),
  KEY `IDX_7bcb57371b481d8e2d66ddeaea` (`createTime`),
  KEY `IDX_479122e3bf464112f7a7253dac` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='admin参数配置';

LOCK TABLES `base_sys_param` WRITE;
/*!40000 ALTER TABLE `base_sys_param` DISABLE KEYS */;

INSERT INTO `base_sys_param` (`id`, `createTime`, `updateTime`, `keyName`, `name`, `data`, `dataType`, `remark`)
VALUES
	(1,'2021-02-26 13:53:05.000000','2021-03-03 17:50:04.000000','text','富文本参数','<p><strong class=\"ql-size-huge\">111xxxxx2222<span class=\"ql-cursor\">﻿﻿</span></strong></p>',0,NULL),
	(2,'2021-02-26 13:53:18.000000','2021-02-26 13:53:18.000000','json','JSON参数','{\n    code: 111\n}',0,NULL);

/*!40000 ALTER TABLE `base_sys_param` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_sys_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_role`;

CREATE TABLE `base_sys_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `adminUserId` varchar(255) NOT NULL COMMENT '用户ID',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `label` varchar(50) DEFAULT NULL COMMENT '角色标签',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `relevance` int(11) NOT NULL DEFAULT '1' COMMENT '数据权限是否关联上下级',
  PRIMARY KEY (`id`),
  UNIQUE KEY `IDX_469d49a5998170e9550cf113da` (`name`),
  UNIQUE KEY `IDX_f3f24fbbccf00192b076e549a7` (`label`),
  KEY `IDX_6f01184441dec49207b41bfd92` (`createTime`),
  KEY `IDX_d64ca209f3fc52128d9b20e97b` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='admin角色';

LOCK TABLES `base_sys_role` WRITE;
/*!40000 ALTER TABLE `base_sys_role` DISABLE KEYS */;

INSERT INTO `base_sys_role` (`id`, `createTime`, `updateTime`, `adminUserId`, `name`, `label`, `remark`, `relevance`)
VALUES
	(1,'2021-02-24 21:18:39.682358','2021-02-24 21:18:39.682358','1','超管','admin','最高权限的角色',1),
	(10,'2021-02-26 14:15:38.000000','2021-02-26 14:15:38.000000','1','系统管理员','admin-sys',NULL,1),
	(11,'2021-02-26 14:16:49.044744','2021-02-26 14:16:49.044744','1','游客','visitor',NULL,0),
	(12,'2021-02-26 14:26:51.000000','2021-02-26 14:32:35.000000','1','开发','dev',NULL,0),
	(13,'2021-02-26 14:27:58.000000','2021-09-08 23:14:12.000000','1','测试','test',NULL,0);

/*!40000 ALTER TABLE `base_sys_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_sys_role_department
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_role_department`;

CREATE TABLE `base_sys_role_department` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `roleId` bigint(20) NOT NULL COMMENT '角色ID',
  `departmentId` bigint(20) NOT NULL COMMENT '部门ID',
  PRIMARY KEY (`id`),
  KEY `IDX_e881a66f7cce83ba431cf20194` (`createTime`),
  KEY `IDX_cbf48031efee5d0de262965e53` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Admin角色部门关联表';

LOCK TABLES `base_sys_role_department` WRITE;
/*!40000 ALTER TABLE `base_sys_role_department` DISABLE KEYS */;

INSERT INTO `base_sys_role_department` (`id`, `createTime`, `updateTime`, `roleId`, `departmentId`)
VALUES
	(1,'2021-02-26 12:00:23.787939','2021-02-26 12:00:23.787939',8,4),
	(2,'2021-02-26 12:01:11.525205','2021-02-26 12:01:11.525205',9,1),
	(3,'2021-02-26 12:01:11.624266','2021-02-26 12:01:11.624266',9,4),
	(4,'2021-02-26 12:01:11.721894','2021-02-26 12:01:11.721894',9,5),
	(5,'2021-02-26 12:01:11.823342','2021-02-26 12:01:11.823342',9,8),
	(6,'2021-02-26 12:01:11.922873','2021-02-26 12:01:11.922873',9,9),
	(23,'2021-02-26 14:32:40.354669','2021-02-26 14:32:40.354669',12,11),
	(25,'2021-02-26 14:32:59.726608','2021-02-26 14:32:59.726608',10,1),
	(29,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,12);

/*!40000 ALTER TABLE `base_sys_role_department` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_sys_role_menu
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_role_menu`;

CREATE TABLE `base_sys_role_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `roleId` bigint(20) NOT NULL COMMENT '角色ID',
  `menuId` bigint(20) NOT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `IDX_3641f81d4201c524a57ce2aa54` (`createTime`),
  KEY `IDX_f860298298b26e7a697be36e5b` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Admin角色菜单关联表';

LOCK TABLES `base_sys_role_menu` WRITE;
/*!40000 ALTER TABLE `base_sys_role_menu` DISABLE KEYS */;

INSERT INTO `base_sys_role_menu` (`id`, `createTime`, `updateTime`, `roleId`, `menuId`)
VALUES
	(1,'2021-02-26 12:00:18.240154','2021-02-26 12:00:18.240154',8,1),
	(2,'2021-02-26 12:00:18.342131','2021-02-26 12:00:18.342131',8,96),
	(3,'2021-02-26 12:00:18.444143','2021-02-26 12:00:18.444143',8,45),
	(4,'2021-02-26 12:00:18.545490','2021-02-26 12:00:18.545490',8,43),
	(5,'2021-02-26 12:00:18.649626','2021-02-26 12:00:18.649626',8,49),
	(6,'2021-02-26 12:00:18.752369','2021-02-26 12:00:18.752369',8,86),
	(7,'2021-02-26 12:00:18.856023','2021-02-26 12:00:18.856023',8,2),
	(8,'2021-02-26 12:00:18.956131','2021-02-26 12:00:18.956131',8,27),
	(9,'2021-02-26 12:00:19.071490','2021-02-26 12:00:19.071490',8,97),
	(10,'2021-02-26 12:00:19.171745','2021-02-26 12:00:19.171745',8,59),
	(11,'2021-02-26 12:00:19.274495','2021-02-26 12:00:19.274495',8,60),
	(12,'2021-02-26 12:00:19.374610','2021-02-26 12:00:19.374610',8,61),
	(13,'2021-02-26 12:00:19.474750','2021-02-26 12:00:19.474750',8,62),
	(14,'2021-02-26 12:00:19.573369','2021-02-26 12:00:19.573369',8,63),
	(15,'2021-02-26 12:00:19.674242','2021-02-26 12:00:19.674242',8,65),
	(16,'2021-02-26 12:00:19.772886','2021-02-26 12:00:19.772886',8,98),
	(17,'2021-02-26 12:00:19.874134','2021-02-26 12:00:19.874134',8,99),
	(18,'2021-02-26 12:00:19.972728','2021-02-26 12:00:19.972728',8,100),
	(19,'2021-02-26 12:00:20.085877','2021-02-26 12:00:20.085877',8,101),
	(20,'2021-02-26 12:00:20.192887','2021-02-26 12:00:20.192887',8,8),
	(21,'2021-02-26 12:00:20.293747','2021-02-26 12:00:20.293747',8,10),
	(22,'2021-02-26 12:00:20.393491','2021-02-26 12:00:20.393491',8,11),
	(23,'2021-02-26 12:00:20.495110','2021-02-26 12:00:20.495110',8,12),
	(24,'2021-02-26 12:00:20.594083','2021-02-26 12:00:20.594083',8,13),
	(25,'2021-02-26 12:00:20.695727','2021-02-26 12:00:20.695727',8,22),
	(26,'2021-02-26 12:00:20.794729','2021-02-26 12:00:20.794729',8,23),
	(27,'2021-02-26 12:00:20.895601','2021-02-26 12:00:20.895601',8,24),
	(28,'2021-02-26 12:00:20.994972','2021-02-26 12:00:20.994972',8,25),
	(29,'2021-02-26 12:00:21.110384','2021-02-26 12:00:21.110384',8,26),
	(30,'2021-02-26 12:00:21.210970','2021-02-26 12:00:21.210970',8,69),
	(31,'2021-02-26 12:00:21.311852','2021-02-26 12:00:21.311852',8,70),
	(32,'2021-02-26 12:00:21.411591','2021-02-26 12:00:21.411591',8,71),
	(33,'2021-02-26 12:00:21.513584','2021-02-26 12:00:21.513584',8,72),
	(34,'2021-02-26 12:00:21.612212','2021-02-26 12:00:21.612212',8,73),
	(35,'2021-02-26 12:00:21.712720','2021-02-26 12:00:21.712720',8,74),
	(36,'2021-02-26 12:00:21.812839','2021-02-26 12:00:21.812839',8,75),
	(37,'2021-02-26 12:00:21.913321','2021-02-26 12:00:21.913321',8,76),
	(38,'2021-02-26 12:00:22.013970','2021-02-26 12:00:22.013970',8,77),
	(39,'2021-02-26 12:00:22.144879','2021-02-26 12:00:22.144879',8,78),
	(40,'2021-02-26 12:00:22.246707','2021-02-26 12:00:22.246707',8,79),
	(41,'2021-02-26 12:00:22.347579','2021-02-26 12:00:22.347579',8,80),
	(42,'2021-02-26 12:00:22.446947','2021-02-26 12:00:22.446947',8,81),
	(43,'2021-02-26 12:00:22.547082','2021-02-26 12:00:22.547082',8,82),
	(44,'2021-02-26 12:00:22.647197','2021-02-26 12:00:22.647197',8,83),
	(45,'2021-02-26 12:00:22.748089','2021-02-26 12:00:22.748089',8,105),
	(46,'2021-02-26 12:00:22.847814','2021-02-26 12:00:22.847814',8,102),
	(47,'2021-02-26 12:00:22.949071','2021-02-26 12:00:22.949071',8,103),
	(48,'2021-02-26 12:00:23.047353','2021-02-26 12:00:23.047353',8,29),
	(49,'2021-02-26 12:00:23.147826','2021-02-26 12:00:23.147826',8,30),
	(50,'2021-02-26 12:00:23.246800','2021-02-26 12:00:23.246800',8,47),
	(51,'2021-02-26 12:00:23.349541','2021-02-26 12:00:23.349541',8,48),
	(52,'2021-02-26 12:00:23.463177','2021-02-26 12:00:23.463177',8,84),
	(53,'2021-02-26 12:00:23.564096','2021-02-26 12:00:23.564096',8,90),
	(54,'2021-02-26 12:00:23.663815','2021-02-26 12:00:23.663815',8,85),
	(55,'2021-02-26 12:01:05.971978','2021-02-26 12:01:05.971978',9,1),
	(56,'2021-02-26 12:01:06.085568','2021-02-26 12:01:06.085568',9,96),
	(57,'2021-02-26 12:01:06.198271','2021-02-26 12:01:06.198271',9,45),
	(58,'2021-02-26 12:01:06.309736','2021-02-26 12:01:06.309736',9,43),
	(59,'2021-02-26 12:01:06.410785','2021-02-26 12:01:06.410785',9,49),
	(60,'2021-02-26 12:01:06.510712','2021-02-26 12:01:06.510712',9,86),
	(61,'2021-02-26 12:01:06.612457','2021-02-26 12:01:06.612457',9,2),
	(62,'2021-02-26 12:01:06.710397','2021-02-26 12:01:06.710397',9,27),
	(63,'2021-02-26 12:01:06.809104','2021-02-26 12:01:06.809104',9,97),
	(64,'2021-02-26 12:01:06.907088','2021-02-26 12:01:06.907088',9,59),
	(65,'2021-02-26 12:01:07.009988','2021-02-26 12:01:07.009988',9,60),
	(66,'2021-02-26 12:01:07.122372','2021-02-26 12:01:07.122372',9,61),
	(67,'2021-02-26 12:01:07.223694','2021-02-26 12:01:07.223694',9,62),
	(68,'2021-02-26 12:01:07.325022','2021-02-26 12:01:07.325022',9,63),
	(69,'2021-02-26 12:01:07.425209','2021-02-26 12:01:07.425209',9,65),
	(70,'2021-02-26 12:01:07.522081','2021-02-26 12:01:07.522081',9,98),
	(71,'2021-02-26 12:01:07.622775','2021-02-26 12:01:07.622775',9,99),
	(72,'2021-02-26 12:01:07.721181','2021-02-26 12:01:07.721181',9,100),
	(73,'2021-02-26 12:01:07.819589','2021-02-26 12:01:07.819589',9,101),
	(74,'2021-02-26 12:01:07.920497','2021-02-26 12:01:07.920497',9,8),
	(75,'2021-02-26 12:01:08.018875','2021-02-26 12:01:08.018875',9,10),
	(76,'2021-02-26 12:01:08.135192','2021-02-26 12:01:08.135192',9,11),
	(77,'2021-02-26 12:01:08.246405','2021-02-26 12:01:08.246405',9,12),
	(78,'2021-02-26 12:01:08.346661','2021-02-26 12:01:08.346661',9,13),
	(79,'2021-02-26 12:01:08.448436','2021-02-26 12:01:08.448436',9,22),
	(80,'2021-02-26 12:01:08.547496','2021-02-26 12:01:08.547496',9,23),
	(81,'2021-02-26 12:01:08.648457','2021-02-26 12:01:08.648457',9,24),
	(82,'2021-02-26 12:01:08.750564','2021-02-26 12:01:08.750564',9,25),
	(83,'2021-02-26 12:01:08.851783','2021-02-26 12:01:08.851783',9,26),
	(84,'2021-02-26 12:01:08.950898','2021-02-26 12:01:08.950898',9,69),
	(85,'2021-02-26 12:01:09.061982','2021-02-26 12:01:09.061982',9,70),
	(86,'2021-02-26 12:01:09.165258','2021-02-26 12:01:09.165258',9,71),
	(87,'2021-02-26 12:01:09.266177','2021-02-26 12:01:09.266177',9,72),
	(88,'2021-02-26 12:01:09.366427','2021-02-26 12:01:09.366427',9,73),
	(89,'2021-02-26 12:01:09.467877','2021-02-26 12:01:09.467877',9,74),
	(90,'2021-02-26 12:01:09.568526','2021-02-26 12:01:09.568526',9,75),
	(91,'2021-02-26 12:01:09.668052','2021-02-26 12:01:09.668052',9,76),
	(92,'2021-02-26 12:01:09.766367','2021-02-26 12:01:09.766367',9,77),
	(93,'2021-02-26 12:01:09.866170','2021-02-26 12:01:09.866170',9,78),
	(94,'2021-02-26 12:01:09.963037','2021-02-26 12:01:09.963037',9,79),
	(95,'2021-02-26 12:01:10.082046','2021-02-26 12:01:10.082046',9,80),
	(96,'2021-02-26 12:01:10.185024','2021-02-26 12:01:10.185024',9,81),
	(97,'2021-02-26 12:01:10.283787','2021-02-26 12:01:10.283787',9,82),
	(98,'2021-02-26 12:01:10.382883','2021-02-26 12:01:10.382883',9,83),
	(99,'2021-02-26 12:01:10.481150','2021-02-26 12:01:10.481150',9,105),
	(100,'2021-02-26 12:01:10.579579','2021-02-26 12:01:10.579579',9,102),
	(101,'2021-02-26 12:01:10.679489','2021-02-26 12:01:10.679489',9,103),
	(102,'2021-02-26 12:01:10.777496','2021-02-26 12:01:10.777496',9,29),
	(103,'2021-02-26 12:01:10.878292','2021-02-26 12:01:10.878292',9,30),
	(104,'2021-02-26 12:01:10.977354','2021-02-26 12:01:10.977354',9,47),
	(105,'2021-02-26 12:01:11.097786','2021-02-26 12:01:11.097786',9,48),
	(106,'2021-02-26 12:01:11.201390','2021-02-26 12:01:11.201390',9,84),
	(107,'2021-02-26 12:01:11.302120','2021-02-26 12:01:11.302120',9,90),
	(108,'2021-02-26 12:01:11.402751','2021-02-26 12:01:11.402751',9,85),
	(161,'2021-02-26 14:16:49.162546','2021-02-26 14:16:49.162546',11,1),
	(162,'2021-02-26 14:16:49.257677','2021-02-26 14:16:49.257677',11,96),
	(163,'2021-02-26 14:16:49.356225','2021-02-26 14:16:49.356225',11,45),
	(164,'2021-02-26 14:16:49.450708','2021-02-26 14:16:49.450708',11,43),
	(165,'2021-02-26 14:16:49.543794','2021-02-26 14:16:49.543794',11,49),
	(166,'2021-02-26 14:16:49.636496','2021-02-26 14:16:49.636496',11,86),
	(167,'2021-02-26 14:16:49.728634','2021-02-26 14:16:49.728634',11,47),
	(168,'2021-02-26 14:16:49.824754','2021-02-26 14:16:49.824754',11,48),
	(169,'2021-02-26 14:16:49.919329','2021-02-26 14:16:49.919329',11,85),
	(170,'2021-02-26 14:16:50.015239','2021-02-26 14:16:50.015239',11,84),
	(290,'2021-02-26 14:32:35.143867','2021-02-26 14:32:35.143867',12,1),
	(291,'2021-02-26 14:32:35.239965','2021-02-26 14:32:35.239965',12,96),
	(292,'2021-02-26 14:32:35.336398','2021-02-26 14:32:35.336398',12,45),
	(293,'2021-02-26 14:32:35.435180','2021-02-26 14:32:35.435180',12,43),
	(294,'2021-02-26 14:32:35.528631','2021-02-26 14:32:35.528631',12,49),
	(295,'2021-02-26 14:32:35.623123','2021-02-26 14:32:35.623123',12,86),
	(296,'2021-02-26 14:32:35.718831','2021-02-26 14:32:35.718831',12,2),
	(297,'2021-02-26 14:32:35.812975','2021-02-26 14:32:35.812975',12,27),
	(298,'2021-02-26 14:32:35.904487','2021-02-26 14:32:35.904487',12,97),
	(299,'2021-02-26 14:32:35.998773','2021-02-26 14:32:35.998773',12,59),
	(300,'2021-02-26 14:32:36.107749','2021-02-26 14:32:36.107749',12,60),
	(301,'2021-02-26 14:32:36.213069','2021-02-26 14:32:36.213069',12,61),
	(302,'2021-02-26 14:32:36.308985','2021-02-26 14:32:36.308985',12,62),
	(303,'2021-02-26 14:32:36.404237','2021-02-26 14:32:36.404237',12,63),
	(304,'2021-02-26 14:32:36.499569','2021-02-26 14:32:36.499569',12,65),
	(305,'2021-02-26 14:32:36.593710','2021-02-26 14:32:36.593710',12,98),
	(306,'2021-02-26 14:32:36.685988','2021-02-26 14:32:36.685988',12,99),
	(307,'2021-02-26 14:32:36.778733','2021-02-26 14:32:36.778733',12,100),
	(308,'2021-02-26 14:32:36.874715','2021-02-26 14:32:36.874715',12,101),
	(309,'2021-02-26 14:32:36.973153','2021-02-26 14:32:36.973153',12,8),
	(310,'2021-02-26 14:32:37.082734','2021-02-26 14:32:37.082734',12,10),
	(311,'2021-02-26 14:32:37.176859','2021-02-26 14:32:37.176859',12,11),
	(312,'2021-02-26 14:32:37.271440','2021-02-26 14:32:37.271440',12,12),
	(313,'2021-02-26 14:32:37.365206','2021-02-26 14:32:37.365206',12,13),
	(314,'2021-02-26 14:32:37.457092','2021-02-26 14:32:37.457092',12,22),
	(315,'2021-02-26 14:32:37.549860','2021-02-26 14:32:37.549860',12,23),
	(316,'2021-02-26 14:32:37.645684','2021-02-26 14:32:37.645684',12,24),
	(317,'2021-02-26 14:32:37.743370','2021-02-26 14:32:37.743370',12,25),
	(318,'2021-02-26 14:32:37.837218','2021-02-26 14:32:37.837218',12,26),
	(319,'2021-02-26 14:32:37.930953','2021-02-26 14:32:37.930953',12,69),
	(320,'2021-02-26 14:32:38.031191','2021-02-26 14:32:38.031191',12,70),
	(321,'2021-02-26 14:32:38.130839','2021-02-26 14:32:38.130839',12,71),
	(322,'2021-02-26 14:32:38.229359','2021-02-26 14:32:38.229359',12,72),
	(323,'2021-02-26 14:32:38.323868','2021-02-26 14:32:38.323868',12,73),
	(324,'2021-02-26 14:32:38.415194','2021-02-26 14:32:38.415194',12,74),
	(325,'2021-02-26 14:32:38.505597','2021-02-26 14:32:38.505597',12,75),
	(326,'2021-02-26 14:32:38.600426','2021-02-26 14:32:38.600426',12,76),
	(327,'2021-02-26 14:32:38.698676','2021-02-26 14:32:38.698676',12,77),
	(328,'2021-02-26 14:32:38.793832','2021-02-26 14:32:38.793832',12,78),
	(329,'2021-02-26 14:32:38.889203','2021-02-26 14:32:38.889203',12,79),
	(330,'2021-02-26 14:32:38.985851','2021-02-26 14:32:38.985851',12,80),
	(331,'2021-02-26 14:32:39.092110','2021-02-26 14:32:39.092110',12,81),
	(332,'2021-02-26 14:32:39.188945','2021-02-26 14:32:39.188945',12,82),
	(333,'2021-02-26 14:32:39.280043','2021-02-26 14:32:39.280043',12,83),
	(334,'2021-02-26 14:32:39.374899','2021-02-26 14:32:39.374899',12,105),
	(335,'2021-02-26 14:32:39.473563','2021-02-26 14:32:39.473563',12,102),
	(336,'2021-02-26 14:32:39.570921','2021-02-26 14:32:39.570921',12,103),
	(337,'2021-02-26 14:32:39.665052','2021-02-26 14:32:39.665052',12,29),
	(338,'2021-02-26 14:32:39.760189','2021-02-26 14:32:39.760189',12,30),
	(339,'2021-02-26 14:32:39.852856','2021-02-26 14:32:39.852856',12,47),
	(340,'2021-02-26 14:32:39.944180','2021-02-26 14:32:39.944180',12,48),
	(341,'2021-02-26 14:32:40.038086','2021-02-26 14:32:40.038086',12,84),
	(342,'2021-02-26 14:32:40.135874','2021-02-26 14:32:40.135874',12,90),
	(343,'2021-02-26 14:32:40.234015','2021-02-26 14:32:40.234015',12,85),
	(355,'2021-02-26 14:32:54.538822','2021-02-26 14:32:54.538822',10,1),
	(356,'2021-02-26 14:32:54.634784','2021-02-26 14:32:54.634784',10,96),
	(357,'2021-02-26 14:32:54.732878','2021-02-26 14:32:54.732878',10,45),
	(358,'2021-02-26 14:32:54.826023','2021-02-26 14:32:54.826023',10,43),
	(359,'2021-02-26 14:32:54.920173','2021-02-26 14:32:54.920173',10,49),
	(360,'2021-02-26 14:32:55.019141','2021-02-26 14:32:55.019141',10,86),
	(361,'2021-02-26 14:32:55.119438','2021-02-26 14:32:55.119438',10,2),
	(362,'2021-02-26 14:32:55.211471','2021-02-26 14:32:55.211471',10,27),
	(363,'2021-02-26 14:32:55.304855','2021-02-26 14:32:55.304855',10,97),
	(364,'2021-02-26 14:32:55.397939','2021-02-26 14:32:55.397939',10,59),
	(365,'2021-02-26 14:32:55.491674','2021-02-26 14:32:55.491674',10,60),
	(366,'2021-02-26 14:32:55.584051','2021-02-26 14:32:55.584051',10,61),
	(367,'2021-02-26 14:32:55.676449','2021-02-26 14:32:55.676449',10,62),
	(368,'2021-02-26 14:32:55.774524','2021-02-26 14:32:55.774524',10,63),
	(369,'2021-02-26 14:32:55.871634','2021-02-26 14:32:55.871634',10,65),
	(370,'2021-02-26 14:32:55.964611','2021-02-26 14:32:55.964611',10,98),
	(371,'2021-02-26 14:32:56.074043','2021-02-26 14:32:56.074043',10,99),
	(372,'2021-02-26 14:32:56.169316','2021-02-26 14:32:56.169316',10,100),
	(373,'2021-02-26 14:32:56.263408','2021-02-26 14:32:56.263408',10,101),
	(374,'2021-02-26 14:32:56.356537','2021-02-26 14:32:56.356537',10,8),
	(375,'2021-02-26 14:32:56.448195','2021-02-26 14:32:56.448195',10,10),
	(376,'2021-02-26 14:32:56.544394','2021-02-26 14:32:56.544394',10,11),
	(377,'2021-02-26 14:32:56.641515','2021-02-26 14:32:56.641515',10,12),
	(378,'2021-02-26 14:32:56.735242','2021-02-26 14:32:56.735242',10,13),
	(379,'2021-02-26 14:32:56.828811','2021-02-26 14:32:56.828811',10,22),
	(380,'2021-02-26 14:32:56.922664','2021-02-26 14:32:56.922664',10,23),
	(381,'2021-02-26 14:32:57.016873','2021-02-26 14:32:57.016873',10,24),
	(382,'2021-02-26 14:32:57.123800','2021-02-26 14:32:57.123800',10,25),
	(383,'2021-02-26 14:32:57.223306','2021-02-26 14:32:57.223306',10,26),
	(384,'2021-02-26 14:32:57.328482','2021-02-26 14:32:57.328482',10,69),
	(385,'2021-02-26 14:32:57.430006','2021-02-26 14:32:57.430006',10,70),
	(386,'2021-02-26 14:32:57.521664','2021-02-26 14:32:57.521664',10,71),
	(387,'2021-02-26 14:32:57.612399','2021-02-26 14:32:57.612399',10,72),
	(388,'2021-02-26 14:32:57.705553','2021-02-26 14:32:57.705553',10,73),
	(389,'2021-02-26 14:32:57.799288','2021-02-26 14:32:57.799288',10,74),
	(390,'2021-02-26 14:32:57.893894','2021-02-26 14:32:57.893894',10,75),
	(391,'2021-02-26 14:32:57.988856','2021-02-26 14:32:57.988856',10,76),
	(392,'2021-02-26 14:32:58.090250','2021-02-26 14:32:58.090250',10,77),
	(393,'2021-02-26 14:32:58.196616','2021-02-26 14:32:58.196616',10,78),
	(394,'2021-02-26 14:32:58.288151','2021-02-26 14:32:58.288151',10,79),
	(395,'2021-02-26 14:32:58.378493','2021-02-26 14:32:58.378493',10,80),
	(396,'2021-02-26 14:32:58.471283','2021-02-26 14:32:58.471283',10,81),
	(397,'2021-02-26 14:32:58.564666','2021-02-26 14:32:58.564666',10,82),
	(398,'2021-02-26 14:32:58.658511','2021-02-26 14:32:58.658511',10,83),
	(399,'2021-02-26 14:32:58.752713','2021-02-26 14:32:58.752713',10,105),
	(400,'2021-02-26 14:32:58.849472','2021-02-26 14:32:58.849472',10,102),
	(401,'2021-02-26 14:32:58.948387','2021-02-26 14:32:58.948387',10,103),
	(402,'2021-02-26 14:32:59.042410','2021-02-26 14:32:59.042410',10,29),
	(403,'2021-02-26 14:32:59.132594','2021-02-26 14:32:59.132594',10,30),
	(404,'2021-02-26 14:32:59.226150','2021-02-26 14:32:59.226150',10,47),
	(405,'2021-02-26 14:32:59.319494','2021-02-26 14:32:59.319494',10,48),
	(406,'2021-02-26 14:32:59.413370','2021-02-26 14:32:59.413370',10,84),
	(407,'2021-02-26 14:32:59.507584','2021-02-26 14:32:59.507584',10,90),
	(408,'2021-02-26 14:32:59.604332','2021-02-26 14:32:59.604332',10,85),
	(517,'2021-08-23 14:37:40.320808','2021-08-23 14:37:40.320808',1,123),
	(518,'2021-08-23 14:37:40.320808','2021-08-23 14:37:40.320808',1,124),
	(519,'2021-08-23 14:45:55.733160','2021-08-23 14:45:55.733160',1,125),
	(520,'2021-08-23 14:45:55.733160','2021-08-23 14:45:55.733160',1,126),
	(521,'2021-08-23 14:48:46.486147','2021-08-23 14:48:46.486147',1,127),
	(522,'2021-08-23 14:48:46.486147','2021-08-23 14:48:46.486147',1,128),
	(523,'2021-08-25 07:49:30.914125','2021-08-25 07:49:30.914125',1,129),
	(524,'2021-08-25 07:49:30.914125','2021-08-25 07:49:30.914125',1,130),
	(525,'2021-08-26 14:43:54.877089','2021-08-26 14:43:54.877089',1,133),
	(526,'2021-08-26 14:43:54.877089','2021-08-26 14:43:54.877089',1,134),
	(527,'2021-08-27 03:31:14.346428','2021-08-27 03:31:14.346428',1,135),
	(528,'2021-08-27 03:31:14.346428','2021-08-27 03:31:14.346428',1,136),
	(529,'2021-08-27 04:37:33.328210','2021-08-27 04:37:33.328210',1,137),
	(530,'2021-08-27 04:37:33.328210','2021-08-27 04:37:33.328210',1,138),
	(551,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,1),
	(552,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,96),
	(553,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,45),
	(554,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,43),
	(555,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,49),
	(556,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,86),
	(557,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,84),
	(558,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,90),
	(559,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,85),
	(560,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,139),
	(561,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,140),
	(562,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,120),
	(563,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,121),
	(564,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,122),
	(565,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,127),
	(566,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,128),
	(567,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,131),
	(568,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,133),
	(569,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,134),
	(570,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,137),
	(571,'2021-09-08 23:14:12.000000','2021-09-08 23:14:12.000000',13,138);

/*!40000 ALTER TABLE `base_sys_role_menu` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_sys_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_user`;

CREATE TABLE `base_sys_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `departmentId` bigint(20) DEFAULT NULL COMMENT '部门ID',
  `name` varchar(255) DEFAULT NULL COMMENT '姓名',
  `username` varchar(100) NOT NULL COMMENT '用户名',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `passwordV` int(11) NOT NULL DEFAULT '1' COMMENT '密码版本, 作用是改完密码，让原来的token失效',
  `salt` varchar(10) DEFAULT NULL COMMENT '密码的盐',
  `nickName` varchar(255) DEFAULT NULL COMMENT '昵称',
  `headImg` varchar(255) DEFAULT NULL COMMENT '头像',
  `phone` varchar(20) DEFAULT NULL COMMENT '手机',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 0:禁用 1：启用',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `socketId` varchar(255) DEFAULT NULL COMMENT 'socketId',
  PRIMARY KEY (`id`),
  UNIQUE KEY `IDX_469ad55973f5b98930f6ad627b` (`username`),
  KEY `IDX_0cf944da378d70a94f5fefd803` (`departmentId`),
  KEY `IDX_9ec6d7ac6337eafb070e4881a8` (`phone`),
  KEY `IDX_ca8611d15a63d52aa4e292e46a` (`createTime`),
  KEY `IDX_a0f2f19cee18445998ece93ddd` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='admin用户表';

LOCK TABLES `base_sys_user` WRITE;
/*!40000 ALTER TABLE `base_sys_user` DISABLE KEYS */;

INSERT INTO `base_sys_user` (`id`, `createTime`, `updateTime`, `departmentId`, `name`, `username`, `password`, `passwordV`, `salt`, `nickName`, `headImg`, `phone`, `email`, `status`, `remark`, `socketId`)
VALUES
	(1,'2021-02-24 21:16:41.525157','2021-09-12 11:57:29.000000',1,'超级管理员','admin','e10adc3949ba59abbe56e057f20f883e',5,NULL,'管理员','http://127.0.0.1:8199/ce7ls8387c6ohcrmdk.png','18000000000','team@cool-js.com',1,'拥有最高权限的用户',NULL),
	(24,'2021-02-26 14:17:38.000000','2021-02-26 14:17:38.000000',11,'小白','xiaobai','e10adc3949ba59abbe56e057f20f883e',1,NULL,'小白',NULL,NULL,NULL,1,NULL,NULL),
	(25,'2021-02-26 14:28:25.000000','2021-02-26 14:28:25.000000',12,'小黑','xiaohei','e10adc3949ba59abbe56e057f20f883e',1,NULL,'小黑',NULL,NULL,NULL,1,NULL,NULL),
	(26,'2021-02-26 14:28:49.000000','2021-02-26 14:28:49.000000',12,'小绿','xiaolv','e10adc3949ba59abbe56e057f20f883e',1,NULL,'小绿',NULL,NULL,NULL,1,NULL,NULL),
	(27,'2021-02-26 14:29:23.000000','2021-02-26 14:29:23.000000',13,'小青','xiaoqin','e10adc3949ba59abbe56e057f20f883e',1,NULL,'小青',NULL,NULL,NULL,1,NULL,NULL),
	(30,'2021-09-08 23:12:11.000000','2021-09-08 23:12:11.000000',1,'小陈','test','24fdfac606c9dbd0a1f20bf21564853c',1,'hcXRs6','陈',NULL,NULL,NULL,1,NULL,NULL);

/*!40000 ALTER TABLE `base_sys_user` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table base_sys_user_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `base_sys_user_role`;

CREATE TABLE `base_sys_user_role` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `adminUserId` bigint(20) NOT NULL COMMENT '用户ID',
  `roleId` bigint(20) NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`id`),
  KEY `IDX_fa9555e03e42fce748c9046b1c` (`createTime`),
  KEY `IDX_3e36c0d2b1a4c659c6b4fc64b3` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Admin用户角色关联表';

LOCK TABLES `base_sys_user_role` WRITE;
/*!40000 ALTER TABLE `base_sys_user_role` DISABLE KEYS */;

INSERT INTO `base_sys_user_role` (`id`, `createTime`, `updateTime`, `adminUserId`, `roleId`)
VALUES
	(1,'2021-02-24 22:03:11.665805','2021-02-24 22:03:11.665805',1,1),
	(2,'2021-02-25 11:03:55.325988','2021-02-25 11:03:55.325988',2,1),
	(3,'2021-02-25 14:30:57.295150','2021-02-25 14:30:57.295150',3,1),
	(4,'2021-02-25 14:39:32.975014','2021-02-25 14:39:32.975014',4,1),
	(5,'2021-02-25 14:40:56.812948','2021-02-25 14:40:56.812948',5,1),
	(6,'2021-02-25 14:44:08.436555','2021-02-25 14:44:08.436555',6,1),
	(7,'2021-02-25 14:46:17.409232','2021-02-25 14:46:17.409232',7,1),
	(8,'2021-02-25 14:47:47.211749','2021-02-25 14:47:47.211749',8,1),
	(9,'2021-02-25 14:48:11.734024','2021-02-25 14:48:11.734024',9,1),
	(10,'2021-02-25 14:50:48.288616','2021-02-25 14:50:48.288616',10,1),
	(11,'2021-02-25 14:51:32.123884','2021-02-25 14:51:32.123884',11,1),
	(12,'2021-02-25 15:46:26.356943','2021-02-25 15:46:26.356943',12,1),
	(13,'2021-02-25 15:56:43.475155','2021-02-25 15:56:43.475155',13,1),
	(14,'2021-02-25 16:03:14.417784','2021-02-25 16:03:14.417784',14,1),
	(16,'2021-02-25 16:22:11.200152','2021-02-25 16:22:11.200152',16,1),
	(17,'2021-02-25 17:44:37.635550','2021-02-25 17:44:37.635550',15,1),
	(19,'2021-02-25 17:51:00.554812','2021-02-25 17:51:00.554812',18,1),
	(21,'2021-02-25 17:54:41.375113','2021-02-25 17:54:41.375113',17,1),
	(22,'2021-02-25 17:55:49.385301','2021-02-25 17:55:49.385301',20,1),
	(24,'2021-02-25 17:58:35.452363','2021-02-25 17:58:35.452363',22,1),
	(27,'2021-02-25 21:25:55.005236','2021-02-25 21:25:55.005236',19,1),
	(28,'2021-02-26 13:50:05.633242','2021-02-26 13:50:05.633242',21,8),
	(29,'2021-02-26 13:50:17.836990','2021-02-26 13:50:17.836990',23,8),
	(38,'2021-02-26 14:36:08.899046','2021-02-26 14:36:08.899046',26,13),
	(39,'2021-02-26 14:36:13.149510','2021-02-26 14:36:13.149510',25,13),
	(40,'2021-02-26 14:36:20.737073','2021-02-26 14:36:20.737073',27,11),
	(42,'2021-02-26 14:36:53.481478','2021-02-26 14:36:53.481478',24,12),
	(45,'2021-09-08 23:12:11.000000','2021-09-08 23:12:11.000000',30,13);

/*!40000 ALTER TABLE `base_sys_user_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table demo_go
# ------------------------------------------------------------

DROP TABLE IF EXISTS `demo_go`;

CREATE TABLE `demo_go` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `subTitle` varchar(255) DEFAULT NULL COMMENT '副标题##IsSearchParams',
  `pic` varchar(255) NOT NULL COMMENT '商品主图',
  `types` tinyint(4) NOT NULL COMMENT '类型#1:上架,2:下架',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态#1:启用,2:禁用',
  `tupian` varchar(300) DEFAULT NULL COMMENT '图片##IsPicColumn',
  `other` tinyint(4) DEFAULT NULL COMMENT '其他状态#1:已激活,2:未激活#IsDictColumn,IsSearchParams',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `otherStr` varchar(4) NOT NULL DEFAULT '' COMMENT '其他状态2#s1:已激活,s2a2:未激活#IsDictColumn,IsSearchParams',
  `appGoodsInfoId` int(11) NOT NULL COMMENT 'appGoodsInfoID',
  PRIMARY KEY (`id`),
  KEY `IDX_d8d0d86a0adf1001ce12aaac41` (`createTime`),
  KEY `IDX_9fe44721d1e517fa24383db56b` (`updateTime`),
  FULLTEXT KEY `IDX_91f50fa9907d5ac2c864f175bb` (`title`,`subTitle`) /*!50100 WITH PARSER `ngram` */ 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='示例go';

LOCK TABLES `demo_go` WRITE;
/*!40000 ALTER TABLE `demo_go` DISABLE KEYS */;

INSERT INTO `demo_go` (`id`, `createTime`, `updateTime`, `title`, `subTitle`, `pic`, `types`, `status`, `tupian`, `other`, `sort`, `otherStr`, `appGoodsInfoId`)
VALUES
	(2,'2021-09-06 11:51:21.000000','2021-09-08 00:49:57.000000','21','12','http://127.0.0.1:8199/ce2hday4j3js9qnemk.png',1,2,'http://127.0.0.1:8199/ce2hday4j3js9qnemk.png',2,12,'s2a2',4);

/*!40000 ALTER TABLE `demo_go` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table gen_code_config
# ------------------------------------------------------------

DROP TABLE IF EXISTS `gen_code_config`;

CREATE TABLE `gen_code_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `dbHost` varchar(255) NOT NULL DEFAULT '' COMMENT '数据库host',
  `dbPort` varchar(50) NOT NULL DEFAULT '' COMMENT '数据库端口',
  `dbUser` varchar(100) NOT NULL DEFAULT '' COMMENT '数据库用户',
  `dbPass` varchar(200) NOT NULL DEFAULT '' COMMENT '数据库密码',
  `dbName` varchar(100) NOT NULL DEFAULT '' COMMENT '数据库名',
  `dbType` varchar(30) NOT NULL DEFAULT 'mysql' COMMENT '数据库类型，如mysql/pgsql/mssql/sqlite/oracle，默认mysql',
  `adminGenPath` text NOT NULL COMMENT '生成代码前端目录',
  `serverGenPath` text NOT NULL COMMENT '生成代码服务端目录',
  `defaultCreatedAtLabel` varchar(30) NOT NULL DEFAULT 'createTime' COMMENT '默认的创建时间字段格式，如createTime',
  `defaultUpdatedAtLabel` varchar(30) NOT NULL DEFAULT 'updateTime' COMMENT '默认的更新时间字段格式，如updateTime',
  `ignoreTablePrefix` varchar(300) NOT NULL DEFAULT '' COMMENT '忽略此前缀的表，逗号分隔',
  `tableNames` text COMMENT '表名，用逗号分隔',
  `activeCode` varchar(300) DEFAULT '' COMMENT '激活码',
  PRIMARY KEY (`id`),
  KEY `IDX_d8d0d86a0adf1001ce12aaac41` (`createTime`),
  KEY `IDX_9fe44721d1e517fa24383db56b` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='代码生成配置';

LOCK TABLES `gen_code_config` WRITE;
/*!40000 ALTER TABLE `gen_code_config` DISABLE KEYS */;

INSERT INTO `gen_code_config` (`id`, `createTime`, `updateTime`, `dbHost`, `dbPort`, `dbUser`, `dbPass`, `dbName`, `dbType`, `adminGenPath`, `serverGenPath`, `defaultCreatedAtLabel`, `defaultUpdatedAtLabel`, `ignoreTablePrefix`, `tableNames`, `activeCode`)
VALUES
	(1,'2021-08-27 13:43:11.000000','2021-09-16 12:21:15.000000','127.0.0.1','3306','root','123456','go-easy','mysql','/Users/yons/Documents/develop/go/src/gitee.com/jasonLaw1015/GfEasyAdmin/','/Users/yons/Documents/develop/go/src/gitee.com/jasonLaw1015/GfEasy/','createTime','updateTime','u_,base_',NULL,''),
	(4,'2021-08-27 15:47:08.000000','2021-08-27 15:47:08.000000','1','12','2','12','1','1','12','2','12','12','1',NULL,'');

/*!40000 ALTER TABLE `gen_code_config` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table gen_code_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `gen_code_log`;

CREATE TABLE `gen_code_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `description` varchar(500) DEFAULT NULL COMMENT '描述##IsSearchParams',
  `tables` text NOT NULL COMMENT '生成的表',
  `params` text NOT NULL COMMENT '当前请求参数',
  `res` text NOT NULL COMMENT '请求结果',
  PRIMARY KEY (`id`),
  KEY `IDX_d8d0d86a0adf1001ce12aaac41` (`createTime`),
  KEY `IDX_9fe44721d1e517fa24383db56b` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='代码生成日志';

LOCK TABLES `gen_code_log` WRITE;
/*!40000 ALTER TABLE `gen_code_log` DISABLE KEYS */;

INSERT INTO `gen_code_log` (`id`, `createTime`, `updateTime`, `title`, `description`, `tables`, `params`, `res`)
VALUES
	(1,'2021-09-12 12:29:41.758618','2021-09-12 12:29:41.758618','s','s','demo_go','{\"Title\":\"s\",\"Description\":\"s\",\"ActiveCode\":\"*****\",\"TableNames\":\"demo_go\",\"ServerGenPath\":\"/Users/yons/Documents/develop/go/src/gitee.com/jasonLaw1015/GoEasy/\",\"AdminGenPath\":\"/Users/yons/Documents/develop/go/src/gitee.com/jasonLaw1015/GoEasyAdmin/\",\"DbHost\":\"127.0.0.1\",\"DbPort\":\"3306\",\"DbUser\":\"root\",\"DbPass\":\"123456\",\"DbName\":\"go-easy\",\"DbType\":\"mysql\",\"DefaultCreatedAtLabel\":\"createTime\",\"DefaultUpdatedAtLabel\":\"updateTime\",\"IgnoreTablePrefix\":\"u_,base_\"}','代码生成成功'),
	(2,'2021-09-12 14:50:05.214990','2021-09-12 14:50:05.214990','d','d','demo_go','{\"Title\":\"d\",\"Description\":\"d\",\"ActiveCode\":\"*****\",\"TableNames\":\"demo_go\",\"ServerGenPath\":\"/Users/yons/Documents/develop/go/src/gitee.com/jasonLaw1015/GoEasy/\",\"AdminGenPath\":\"/Users/yons/Documents/develop/go/src/gitee.com/jasonLaw1015/GoEasyAdmin/\",\"DbHost\":\"127.0.0.1\",\"DbPort\":\"3306\",\"DbUser\":\"root\",\"DbPass\":\"123456\",\"DbName\":\"go-easy\",\"DbType\":\"mysql\",\"DefaultCreatedAtLabel\":\"createTime\",\"DefaultUpdatedAtLabel\":\"updateTime\",\"IgnoreTablePrefix\":\"u_,base_\"}','代码生成成功');

/*!40000 ALTER TABLE `gen_code_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table task_info
# ------------------------------------------------------------

DROP TABLE IF EXISTS `task_info`;

CREATE TABLE `task_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `jobId` varchar(255) DEFAULT NULL COMMENT '任务ID',
  `repeatConf` varchar(1000) DEFAULT NULL COMMENT '任务配置',
  `name` varchar(255) NOT NULL COMMENT '名称',
  `cron` varchar(255) DEFAULT NULL COMMENT 'cron',
  `limit` int(11) DEFAULT NULL COMMENT '最大执行次数 不传为无限次',
  `every` int(11) DEFAULT NULL COMMENT '每间隔多少毫秒执行一次 如果cron设置了 这项设置就无效',
  `remark` varchar(255) DEFAULT NULL COMMENT '备注',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态 2:停止 1：运行',
  `startDate` datetime DEFAULT NULL COMMENT '开始时间',
  `endDate` datetime DEFAULT NULL COMMENT '结束时间',
  `data` varchar(255) DEFAULT NULL COMMENT '数据',
  `service` varchar(255) DEFAULT NULL COMMENT '执行的service实例ID',
  `types` tinyint(4) NOT NULL DEFAULT '2' COMMENT '状态 2:系统 1：用户',
  `nextRunTime` datetime DEFAULT NULL COMMENT '下一次执行时间',
  `taskType` tinyint(4) NOT NULL DEFAULT '2' COMMENT '状态 2:cron 1：时间间隔',
  PRIMARY KEY (`id`),
  KEY `IDX_6ced02f467e59bd6306b549bb0` (`createTime`),
  KEY `IDX_2adc6f9c241391126f27dac145` (`updateTime`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='admin定时任务表';

LOCK TABLES `task_info` WRITE;
/*!40000 ALTER TABLE `task_info` DISABLE KEYS */;

INSERT INTO `task_info` (`id`, `createTime`, `updateTime`, `jobId`, `repeatConf`, `name`, `cron`, `limit`, `every`, `remark`, `status`, `startDate`, `endDate`, `data`, `service`, `types`, `nextRunTime`, `taskType`)
VALUES
	(1,'2021-03-10 14:25:13.381172','2021-08-25 21:26:00.000000',NULL,'{\"count\":1,\"type\":1,\"limit\":5,\"name\":\"每秒执行,总共5次\",\"taskType\":1,\"every\":1000,\"service\":\"taskDemoService.test()\",\"status\":1,\"id\":1,\"createTime\":\"2021-03-10 14:25:13\",\"updateTime\":\"2021-03-10 14:25:13\",\"jobId\":1}','每秒执行,总共5次',NULL,5,1000,NULL,2,NULL,NULL,NULL,'taskDemoService.test()',1,'2021-03-10 14:25:18',1),
	(2,'2021-03-10 14:25:53.000000','2021-08-25 21:26:00.000000',NULL,'{\"count\":1,\"id\":2,\"createTime\":\"2021-03-10 14:25:53\",\"updateTime\":\"2021-03-10 14:25:55\",\"name\":\"cron任务，5秒执行一次\",\"cron\":\"0/5 * * * * ? \",\"status\":1,\"service\":\"taskDemoService.test()\",\"type\":1,\"nextRunTime\":\"2021-03-10 14:26:00\",\"taskType\":0,\"jobId\":2}','cron任务，5秒执行一次','0/5 * * * * ? ',NULL,NULL,NULL,2,NULL,NULL,NULL,'taskDemoService.test()',1,NULL,2),
	(3,'2021-08-21 17:14:59.000000','2021-08-25 21:26:26.000000',NULL,NULL,'[sys]每隔一小时监测存在定时任务，然后更新数据状态','0 */2 * * * ?',NULL,NULL,NULL,2,NULL,NULL,NULL,'SelectActiveTask2UpdateStatus',2,NULL,2);

/*!40000 ALTER TABLE `task_info` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table task_log
# ------------------------------------------------------------

DROP TABLE IF EXISTS `task_log`;

CREATE TABLE `task_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `createTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) COMMENT '创建时间',
  `updateTime` datetime(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6) COMMENT '更新时间',
  `taskId` bigint(20) DEFAULT NULL COMMENT '任务ID',
  `status` tinyint(4) NOT NULL DEFAULT '2' COMMENT '状态 2:失败 1：成功',
  `detail` text COMMENT '详情描述',
  PRIMARY KEY (`id`),
  KEY `IDX_b9af0e100be034924b270aab31` (`createTime`),
  KEY `IDX_8857d8d43d38bebd7159af1fa6` (`updateTime`),
  KEY `IDX_1142dfec452e924b346f060fda` (`taskId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Admin定时任务日志表';

LOCK TABLES `task_log` WRITE;
/*!40000 ALTER TABLE `task_log` DISABLE KEYS */;

INSERT INTO `task_log` (`id`, `createTime`, `updateTime`, `taskId`, `status`, `detail`)
VALUES
	(1,'2021-03-10 14:25:14.020930','2021-03-10 14:25:14.020930',1,1,'\"任务执行成功\"'),
	(2,'2021-03-10 14:25:15.012030','2021-03-10 14:25:15.012030',1,1,'\"任务执行成功\"'),
	(3,'2021-03-10 14:25:16.011443','2021-03-10 14:25:16.011443',1,1,'\"任务执行成功\"'),
	(4,'2021-03-10 14:25:17.009939','2021-03-10 14:25:17.009939',1,1,'\"任务执行成功\"'),
	(5,'2021-03-10 14:25:18.010410','2021-03-10 14:25:18.010410',1,1,'\"任务执行成功\"'),
	(6,'2021-03-10 14:25:55.012816','2021-03-10 14:25:55.012816',2,1,''),
	(7,'2021-03-10 14:26:00.011880','2021-03-10 14:26:00.011880',2,1,''),
	(8,'2021-03-10 14:26:05.016832','2021-03-10 14:26:05.016832',2,1,'\"任务执行成功\"'),
	(9,'2021-03-10 14:26:10.011763','2021-03-10 14:26:10.011763',2,1,'\"任务执行成功\"'),
	(10,'2021-03-10 14:26:15.010246','2021-03-10 14:26:15.010246',2,1,'\"任务执行成功\"'),
	(11,'2021-08-21 15:46:15.000000','2021-08-21 15:46:15.000000',1,2,'没有该方法:taskDemoService.test()'),
	(12,'2021-08-21 15:46:16.000000','2021-08-21 15:46:16.000000',1,2,'没有该方法:taskDemoService.test()'),
	(13,'2021-08-21 15:46:17.000000','2021-08-21 15:46:17.000000',1,2,'没有该方法:taskDemoService.test()'),
	(14,'2021-08-21 15:46:18.000000','2021-08-21 15:46:18.000000',1,2,'没有该方法:taskDemoService.test()'),
	(15,'2021-08-21 15:46:19.000000','2021-08-21 15:46:19.000000',1,2,'没有该方法:taskDemoService.test()'),
	(16,'2021-08-21 15:50:04.000000','2021-08-21 15:50:04.000000',1,2,'没有该方法:taskDemoService.test()'),
	(17,'2021-08-21 15:50:05.000000','2021-08-21 15:50:05.000000',1,2,'没有该方法:taskDemoService.test()'),
	(18,'2021-08-21 15:50:06.000000','2021-08-21 15:50:06.000000',1,2,'没有该方法:taskDemoService.test()'),
	(19,'2021-08-21 15:50:07.000000','2021-08-21 15:50:07.000000',1,2,'没有该方法:taskDemoService.test()'),
	(20,'2021-08-21 15:50:08.000000','2021-08-21 15:50:08.000000',1,2,'没有该方法:taskDemoService.test()'),
	(21,'2021-08-21 15:59:41.000000','2021-08-21 15:59:41.000000',1,2,'没有该方法:taskDemoService.test()'),
	(22,'2021-08-21 15:59:42.000000','2021-08-21 15:59:42.000000',1,2,'没有该方法:taskDemoService.test()'),
	(23,'2021-08-21 15:59:43.000000','2021-08-21 15:59:43.000000',1,2,'没有该方法:taskDemoService.test()'),
	(24,'2021-08-21 15:59:44.000000','2021-08-21 15:59:44.000000',1,2,'没有该方法:taskDemoService.test()'),
	(25,'2021-08-21 15:59:45.000000','2021-08-21 15:59:45.000000',1,2,'没有该方法:taskDemoService.test()'),
	(26,'2021-08-21 18:35:10.000000','2021-08-21 18:35:10.000000',3,1,'任务执行成功'),
	(27,'2021-08-21 18:35:11.000000','2021-08-21 18:35:11.000000',3,1,'任务执行成功'),
	(28,'2021-08-21 18:35:12.000000','2021-08-21 18:35:12.000000',3,1,'任务执行成功'),
	(29,'2021-08-21 18:35:13.000000','2021-08-21 18:35:13.000000',3,1,'任务执行成功'),
	(30,'2021-08-21 18:35:14.000000','2021-08-21 18:35:14.000000',3,1,'任务执行成功'),
	(31,'2021-08-21 18:35:15.000000','2021-08-21 18:35:15.000000',3,1,'任务执行成功'),
	(32,'2021-08-21 18:35:16.000000','2021-08-21 18:35:16.000000',3,1,'任务执行成功'),
	(33,'2021-08-21 18:35:17.000000','2021-08-21 18:35:17.000000',3,1,'任务执行成功'),
	(34,'2021-08-21 18:35:18.000000','2021-08-21 18:35:18.000000',3,1,'任务执行成功'),
	(35,'2021-08-21 18:35:19.000000','2021-08-21 18:35:19.000000',3,1,'任务执行成功'),
	(36,'2021-08-21 18:35:20.000000','2021-08-21 18:35:20.000000',3,1,'任务执行成功'),
	(37,'2021-08-21 18:35:21.000000','2021-08-21 18:35:21.000000',3,1,'任务执行成功'),
	(38,'2021-08-21 18:35:22.000000','2021-08-21 18:35:22.000000',3,1,'任务执行成功'),
	(39,'2021-08-21 18:35:23.000000','2021-08-21 18:35:23.000000',3,1,'任务执行成功'),
	(40,'2021-08-21 18:35:24.000000','2021-08-21 18:35:24.000000',3,1,'任务执行成功'),
	(41,'2021-08-21 18:35:25.000000','2021-08-21 18:35:25.000000',3,1,'任务执行成功'),
	(42,'2021-08-21 18:35:26.000000','2021-08-21 18:35:26.000000',3,1,'任务执行成功'),
	(43,'2021-08-21 18:35:27.000000','2021-08-21 18:35:27.000000',3,1,'任务执行成功'),
	(44,'2021-08-21 18:35:28.000000','2021-08-21 18:35:28.000000',3,1,'任务执行成功'),
	(45,'2021-08-21 18:35:29.000000','2021-08-21 18:35:29.000000',3,1,'任务执行成功'),
	(46,'2021-08-21 18:35:30.000000','2021-08-21 18:35:30.000000',3,1,'任务执行成功'),
	(47,'2021-08-21 18:35:31.000000','2021-08-21 18:35:31.000000',3,1,'任务执行成功'),
	(48,'2021-08-21 18:35:32.000000','2021-08-21 18:35:32.000000',3,1,'任务执行成功'),
	(49,'2021-08-21 18:35:34.000000','2021-08-21 18:35:34.000000',3,1,'任务执行成功'),
	(50,'2021-08-21 18:35:35.000000','2021-08-21 18:35:35.000000',3,1,'任务执行成功'),
	(51,'2021-08-21 18:35:36.000000','2021-08-21 18:35:36.000000',3,1,'任务执行成功'),
	(52,'2021-08-21 18:35:37.000000','2021-08-21 18:35:37.000000',3,1,'任务执行成功'),
	(53,'2021-08-21 18:35:38.000000','2021-08-21 18:35:38.000000',3,1,'任务执行成功'),
	(54,'2021-08-21 18:35:39.000000','2021-08-21 18:35:39.000000',3,1,'任务执行成功'),
	(55,'2021-08-21 18:35:40.000000','2021-08-21 18:35:40.000000',3,1,'任务执行成功'),
	(56,'2021-08-21 18:35:41.000000','2021-08-21 18:35:41.000000',3,1,'任务执行成功'),
	(57,'2021-08-21 18:35:42.000000','2021-08-21 18:35:42.000000',3,1,'任务执行成功'),
	(58,'2021-08-21 18:35:43.000000','2021-08-21 18:35:43.000000',3,1,'任务执行成功'),
	(59,'2021-08-21 18:35:44.000000','2021-08-21 18:35:44.000000',3,1,'任务执行成功'),
	(60,'2021-08-21 18:35:45.000000','2021-08-21 18:35:45.000000',3,1,'任务执行成功'),
	(61,'2021-08-21 18:35:46.000000','2021-08-21 18:35:46.000000',3,1,'任务执行成功'),
	(62,'2021-08-21 18:35:47.000000','2021-08-21 18:35:47.000000',3,1,'任务执行成功'),
	(63,'2021-08-21 18:35:48.000000','2021-08-21 18:35:48.000000',3,1,'任务执行成功'),
	(64,'2021-08-21 18:38:47.000000','2021-08-21 18:38:47.000000',3,1,'任务执行成功'),
	(65,'2021-08-21 18:38:48.000000','2021-08-21 18:38:48.000000',3,1,'任务执行成功'),
	(66,'2021-08-21 18:38:49.000000','2021-08-21 18:38:49.000000',3,1,'任务执行成功'),
	(67,'2021-08-21 18:38:50.000000','2021-08-21 18:38:50.000000',3,1,'任务执行成功'),
	(68,'2021-08-21 18:38:51.000000','2021-08-21 18:38:51.000000',3,1,'任务执行成功'),
	(69,'2021-08-21 18:38:52.000000','2021-08-21 18:38:52.000000',3,1,'任务执行成功'),
	(70,'2021-08-21 18:38:53.000000','2021-08-21 18:38:53.000000',3,1,'任务执行成功'),
	(71,'2021-08-21 18:42:53.000000','2021-08-21 18:42:53.000000',3,1,'任务执行成功'),
	(72,'2021-08-21 18:42:54.000000','2021-08-21 18:42:54.000000',3,1,'任务执行成功'),
	(73,'2021-08-21 18:42:55.000000','2021-08-21 18:42:55.000000',3,1,'任务执行成功'),
	(74,'2021-08-21 18:42:56.000000','2021-08-21 18:42:56.000000',3,1,'任务执行成功'),
	(75,'2021-08-21 18:42:57.000000','2021-08-21 18:42:57.000000',3,1,'任务执行成功'),
	(76,'2021-08-21 18:42:58.000000','2021-08-21 18:42:58.000000',3,1,'任务执行成功'),
	(77,'2021-08-21 18:42:59.000000','2021-08-21 18:42:59.000000',3,1,'任务执行成功'),
	(78,'2021-08-21 18:43:00.000000','2021-08-21 18:43:00.000000',3,1,'任务执行成功'),
	(79,'2021-08-21 18:45:32.000000','2021-08-21 18:45:32.000000',3,1,'任务执行成功'),
	(80,'2021-08-21 21:23:55.000000','2021-08-21 21:23:55.000000',3,1,'任务执行成功'),
	(81,'2021-08-21 21:23:59.000000','2021-08-21 21:23:59.000000',3,1,'任务执行成功'),
	(82,'2021-08-21 21:24:01.000000','2021-08-21 21:24:01.000000',3,1,'任务执行成功'),
	(83,'2021-08-21 21:25:39.000000','2021-08-21 21:25:39.000000',3,1,'任务执行成功'),
	(84,'2021-08-21 21:25:40.000000','2021-08-21 21:25:40.000000',3,1,'任务执行成功'),
	(85,'2021-08-21 21:25:41.000000','2021-08-21 21:25:41.000000',3,1,'任务执行成功'),
	(86,'2021-08-21 21:25:42.000000','2021-08-21 21:25:42.000000',3,1,'任务执行成功'),
	(87,'2021-08-21 21:25:43.000000','2021-08-21 21:25:43.000000',3,1,'任务执行成功'),
	(88,'2021-08-21 21:25:44.000000','2021-08-21 21:25:44.000000',3,1,'任务执行成功'),
	(89,'2021-08-21 21:25:45.000000','2021-08-21 21:25:45.000000',3,1,'任务执行成功'),
	(90,'2021-08-21 21:25:46.000000','2021-08-21 21:25:46.000000',3,1,'任务执行成功'),
	(91,'2021-08-21 21:25:47.000000','2021-08-21 21:25:47.000000',3,1,'任务执行成功'),
	(92,'2021-08-21 21:25:48.000000','2021-08-21 21:25:48.000000',3,1,'任务执行成功'),
	(93,'2021-08-21 21:25:49.000000','2021-08-21 21:25:49.000000',3,1,'任务执行成功'),
	(94,'2021-08-21 21:25:50.000000','2021-08-21 21:25:50.000000',3,1,'任务执行成功'),
	(95,'2021-08-21 21:25:51.000000','2021-08-21 21:25:51.000000',3,1,'任务执行成功'),
	(96,'2021-08-21 21:25:52.000000','2021-08-21 21:25:52.000000',3,1,'任务执行成功'),
	(97,'2021-08-21 21:25:53.000000','2021-08-21 21:25:53.000000',3,1,'任务执行成功'),
	(98,'2021-08-21 21:25:54.000000','2021-08-21 21:25:54.000000',3,1,'任务执行成功'),
	(99,'2021-08-21 21:25:55.000000','2021-08-21 21:25:55.000000',3,1,'任务执行成功'),
	(100,'2021-08-21 21:25:56.000000','2021-08-21 21:25:56.000000',3,1,'任务执行成功'),
	(101,'2021-08-21 21:25:57.000000','2021-08-21 21:25:57.000000',3,1,'任务执行成功'),
	(102,'2021-08-21 21:25:58.000000','2021-08-21 21:25:58.000000',3,1,'任务执行成功'),
	(103,'2021-08-21 21:25:59.000000','2021-08-21 21:25:59.000000',3,1,'任务执行成功'),
	(104,'2021-08-21 21:26:00.000000','2021-08-21 21:26:00.000000',3,1,'任务执行成功'),
	(105,'2021-08-21 21:26:01.000000','2021-08-21 21:26:01.000000',3,1,'任务执行成功'),
	(106,'2021-08-21 21:26:02.000000','2021-08-21 21:26:02.000000',3,1,'任务执行成功'),
	(107,'2021-08-21 21:26:03.000000','2021-08-21 21:26:03.000000',3,1,'任务执行成功'),
	(108,'2021-08-21 21:26:04.000000','2021-08-21 21:26:04.000000',3,1,'任务执行成功'),
	(109,'2021-08-21 21:26:05.000000','2021-08-21 21:26:05.000000',3,1,'任务执行成功'),
	(110,'2021-08-21 21:26:06.000000','2021-08-21 21:26:06.000000',3,1,'任务执行成功'),
	(111,'2021-08-21 21:26:07.000000','2021-08-21 21:26:07.000000',3,1,'任务执行成功'),
	(112,'2021-08-21 21:26:08.000000','2021-08-21 21:26:08.000000',3,1,'任务执行成功'),
	(113,'2021-08-21 21:26:09.000000','2021-08-21 21:26:09.000000',3,1,'任务执行成功'),
	(114,'2021-08-21 21:26:10.000000','2021-08-21 21:26:10.000000',3,1,'任务执行成功'),
	(115,'2021-08-21 21:26:11.000000','2021-08-21 21:26:11.000000',3,1,'任务执行成功'),
	(116,'2021-08-21 21:26:12.000000','2021-08-21 21:26:12.000000',3,1,'任务执行成功'),
	(117,'2021-08-21 21:26:13.000000','2021-08-21 21:26:13.000000',3,1,'任务执行成功'),
	(118,'2021-08-21 21:26:14.000000','2021-08-21 21:26:14.000000',3,1,'任务执行成功'),
	(119,'2021-08-21 21:26:15.000000','2021-08-21 21:26:15.000000',3,1,'任务执行成功'),
	(120,'2021-08-21 21:26:16.000000','2021-08-21 21:26:16.000000',3,1,'任务执行成功'),
	(121,'2021-08-21 21:26:17.000000','2021-08-21 21:26:17.000000',3,1,'任务执行成功'),
	(122,'2021-08-21 21:26:18.000000','2021-08-21 21:26:18.000000',3,1,'任务执行成功'),
	(123,'2021-08-21 21:26:19.000000','2021-08-21 21:26:19.000000',3,1,'任务执行成功'),
	(124,'2021-08-21 21:26:20.000000','2021-08-21 21:26:20.000000',3,1,'任务执行成功'),
	(125,'2021-08-21 21:26:21.000000','2021-08-21 21:26:21.000000',3,1,'任务执行成功'),
	(126,'2021-08-21 21:26:22.000000','2021-08-21 21:26:22.000000',3,1,'任务执行成功'),
	(127,'2021-08-21 21:26:23.000000','2021-08-21 21:26:23.000000',3,1,'任务执行成功'),
	(128,'2021-08-21 21:26:24.000000','2021-08-21 21:26:24.000000',3,1,'任务执行成功'),
	(129,'2021-08-21 21:26:25.000000','2021-08-21 21:26:25.000000',3,1,'任务执行成功'),
	(130,'2021-08-21 21:26:26.000000','2021-08-21 21:26:26.000000',3,1,'任务执行成功'),
	(131,'2021-08-21 21:26:27.000000','2021-08-21 21:26:27.000000',3,1,'任务执行成功'),
	(132,'2021-08-21 21:26:28.000000','2021-08-21 21:26:28.000000',3,1,'任务执行成功'),
	(133,'2021-08-21 21:26:29.000000','2021-08-21 21:26:29.000000',3,1,'任务执行成功'),
	(134,'2021-08-21 21:26:30.000000','2021-08-21 21:26:30.000000',3,1,'任务执行成功'),
	(135,'2021-08-21 21:26:31.000000','2021-08-21 21:26:31.000000',3,1,'任务执行成功'),
	(136,'2021-08-21 21:26:32.000000','2021-08-21 21:26:32.000000',3,1,'任务执行成功'),
	(137,'2021-08-21 21:26:33.000000','2021-08-21 21:26:33.000000',3,1,'任务执行成功'),
	(138,'2021-08-21 21:26:34.000000','2021-08-21 21:26:34.000000',3,1,'任务执行成功'),
	(139,'2021-08-21 21:26:35.000000','2021-08-21 21:26:35.000000',3,1,'任务执行成功'),
	(140,'2021-08-21 21:26:36.000000','2021-08-21 21:26:36.000000',3,1,'任务执行成功'),
	(141,'2021-08-21 21:26:37.000000','2021-08-21 21:26:37.000000',3,1,'任务执行成功'),
	(142,'2021-08-21 21:26:38.000000','2021-08-21 21:26:38.000000',3,1,'任务执行成功'),
	(143,'2021-08-21 21:26:39.000000','2021-08-21 21:26:39.000000',3,1,'任务执行成功'),
	(144,'2021-08-21 21:26:40.000000','2021-08-21 21:26:40.000000',3,1,'任务执行成功'),
	(145,'2021-08-21 21:26:41.000000','2021-08-21 21:26:41.000000',3,1,'任务执行成功'),
	(146,'2021-08-21 21:26:42.000000','2021-08-21 21:26:42.000000',3,1,'任务执行成功'),
	(147,'2021-08-21 21:26:43.000000','2021-08-21 21:26:43.000000',3,1,'任务执行成功'),
	(148,'2021-08-21 21:26:44.000000','2021-08-21 21:26:44.000000',3,1,'任务执行成功'),
	(149,'2021-08-21 21:26:45.000000','2021-08-21 21:26:45.000000',3,1,'任务执行成功'),
	(150,'2021-08-21 21:26:46.000000','2021-08-21 21:26:46.000000',3,1,'任务执行成功'),
	(151,'2021-08-21 21:26:47.000000','2021-08-21 21:26:47.000000',3,1,'任务执行成功'),
	(152,'2021-08-21 21:26:48.000000','2021-08-21 21:26:48.000000',3,1,'任务执行成功'),
	(153,'2021-08-21 21:26:49.000000','2021-08-21 21:26:49.000000',3,1,'任务执行成功'),
	(154,'2021-08-21 21:26:50.000000','2021-08-21 21:26:50.000000',3,1,'任务执行成功'),
	(155,'2021-08-21 21:26:51.000000','2021-08-21 21:26:51.000000',3,1,'任务执行成功'),
	(156,'2021-08-21 21:26:52.000000','2021-08-21 21:26:52.000000',3,1,'任务执行成功'),
	(157,'2021-08-21 21:26:53.000000','2021-08-21 21:26:53.000000',3,1,'任务执行成功'),
	(158,'2021-08-21 21:26:54.000000','2021-08-21 21:26:54.000000',3,1,'任务执行成功'),
	(159,'2021-08-21 21:26:55.000000','2021-08-21 21:26:55.000000',3,1,'任务执行成功'),
	(160,'2021-08-21 21:26:56.000000','2021-08-21 21:26:56.000000',3,1,'任务执行成功'),
	(161,'2021-08-21 21:26:57.000000','2021-08-21 21:26:57.000000',3,1,'任务执行成功'),
	(162,'2021-08-21 21:26:58.000000','2021-08-21 21:26:58.000000',3,1,'任务执行成功'),
	(163,'2021-08-21 21:26:59.000000','2021-08-21 21:26:59.000000',3,1,'任务执行成功'),
	(164,'2021-08-21 21:27:00.000000','2021-08-21 21:27:00.000000',3,1,'任务执行成功'),
	(165,'2021-08-21 21:27:01.000000','2021-08-21 21:27:01.000000',3,1,'任务执行成功'),
	(166,'2021-08-21 21:27:02.000000','2021-08-21 21:27:02.000000',3,1,'任务执行成功'),
	(167,'2021-08-21 21:27:03.000000','2021-08-21 21:27:03.000000',3,1,'任务执行成功'),
	(168,'2021-08-21 21:27:04.000000','2021-08-21 21:27:04.000000',3,1,'任务执行成功'),
	(169,'2021-08-21 21:27:05.000000','2021-08-21 21:27:05.000000',3,1,'任务执行成功'),
	(170,'2021-08-21 21:27:06.000000','2021-08-21 21:27:06.000000',3,1,'任务执行成功'),
	(171,'2021-08-21 21:27:07.000000','2021-08-21 21:27:07.000000',3,1,'任务执行成功'),
	(172,'2021-08-21 21:27:08.000000','2021-08-21 21:27:08.000000',3,1,'任务执行成功'),
	(173,'2021-08-21 21:27:09.000000','2021-08-21 21:27:09.000000',3,1,'任务执行成功'),
	(174,'2021-08-21 21:27:10.000000','2021-08-21 21:27:10.000000',3,1,'任务执行成功'),
	(175,'2021-08-21 21:27:11.000000','2021-08-21 21:27:11.000000',3,1,'任务执行成功'),
	(176,'2021-08-21 21:27:12.000000','2021-08-21 21:27:12.000000',3,1,'任务执行成功'),
	(177,'2021-08-21 21:27:13.000000','2021-08-21 21:27:13.000000',3,1,'任务执行成功'),
	(178,'2021-08-21 21:27:14.000000','2021-08-21 21:27:14.000000',3,1,'任务执行成功'),
	(179,'2021-08-21 21:27:15.000000','2021-08-21 21:27:15.000000',3,1,'任务执行成功'),
	(180,'2021-08-21 21:27:16.000000','2021-08-21 21:27:16.000000',3,1,'任务执行成功'),
	(181,'2021-08-21 21:27:17.000000','2021-08-21 21:27:17.000000',3,1,'任务执行成功'),
	(182,'2021-08-21 21:27:18.000000','2021-08-21 21:27:18.000000',3,1,'任务执行成功'),
	(183,'2021-08-21 21:27:19.000000','2021-08-21 21:27:19.000000',3,1,'任务执行成功'),
	(184,'2021-08-21 21:27:20.000000','2021-08-21 21:27:20.000000',3,1,'任务执行成功'),
	(185,'2021-08-21 21:27:21.000000','2021-08-21 21:27:21.000000',3,1,'任务执行成功'),
	(186,'2021-08-21 21:27:22.000000','2021-08-21 21:27:22.000000',3,1,'任务执行成功'),
	(187,'2021-08-21 21:27:23.000000','2021-08-21 21:27:23.000000',3,1,'任务执行成功'),
	(188,'2021-08-21 21:27:24.000000','2021-08-21 21:27:24.000000',3,1,'任务执行成功'),
	(189,'2021-08-21 21:27:25.000000','2021-08-21 21:27:25.000000',3,1,'任务执行成功'),
	(190,'2021-08-21 21:27:26.000000','2021-08-21 21:27:26.000000',3,1,'任务执行成功'),
	(191,'2021-08-21 21:27:27.000000','2021-08-21 21:27:27.000000',3,1,'任务执行成功'),
	(192,'2021-08-21 21:27:28.000000','2021-08-21 21:27:28.000000',3,1,'任务执行成功'),
	(193,'2021-08-21 21:27:29.000000','2021-08-21 21:27:29.000000',3,1,'任务执行成功'),
	(194,'2021-08-21 21:27:30.000000','2021-08-21 21:27:30.000000',3,1,'任务执行成功'),
	(195,'2021-08-21 21:27:31.000000','2021-08-21 21:27:31.000000',3,1,'任务执行成功'),
	(196,'2021-08-21 21:27:32.000000','2021-08-21 21:27:32.000000',3,1,'任务执行成功'),
	(197,'2021-08-21 21:27:33.000000','2021-08-21 21:27:33.000000',3,1,'任务执行成功'),
	(198,'2021-08-21 21:27:34.000000','2021-08-21 21:27:34.000000',3,1,'任务执行成功'),
	(199,'2021-08-21 21:27:35.000000','2021-08-21 21:27:35.000000',3,1,'任务执行成功'),
	(200,'2021-08-21 21:27:36.000000','2021-08-21 21:27:36.000000',3,1,'任务执行成功'),
	(201,'2021-08-21 21:27:37.000000','2021-08-21 21:27:37.000000',3,1,'任务执行成功'),
	(202,'2021-08-21 21:27:38.000000','2021-08-21 21:27:38.000000',3,1,'任务执行成功'),
	(203,'2021-08-21 21:27:39.000000','2021-08-21 21:27:39.000000',3,1,'任务执行成功'),
	(204,'2021-08-21 21:27:40.000000','2021-08-21 21:27:40.000000',3,1,'任务执行成功'),
	(205,'2021-08-21 21:27:41.000000','2021-08-21 21:27:41.000000',3,1,'任务执行成功'),
	(206,'2021-08-21 21:27:42.000000','2021-08-21 21:27:42.000000',3,1,'任务执行成功'),
	(207,'2021-08-21 21:27:43.000000','2021-08-21 21:27:43.000000',3,1,'任务执行成功'),
	(208,'2021-08-21 21:27:44.000000','2021-08-21 21:27:44.000000',3,1,'任务执行成功'),
	(209,'2021-08-21 21:27:45.000000','2021-08-21 21:27:45.000000',3,1,'任务执行成功'),
	(210,'2021-08-21 21:27:46.000000','2021-08-21 21:27:46.000000',3,1,'任务执行成功'),
	(211,'2021-08-21 21:27:47.000000','2021-08-21 21:27:47.000000',3,1,'任务执行成功'),
	(212,'2021-08-21 21:27:48.000000','2021-08-21 21:27:48.000000',3,1,'任务执行成功'),
	(213,'2021-08-21 21:27:49.000000','2021-08-21 21:27:49.000000',3,1,'任务执行成功'),
	(214,'2021-08-21 21:27:50.000000','2021-08-21 21:27:50.000000',3,1,'任务执行成功'),
	(215,'2021-08-21 21:27:51.000000','2021-08-21 21:27:51.000000',3,1,'任务执行成功'),
	(216,'2021-08-21 21:27:52.000000','2021-08-21 21:27:52.000000',3,1,'任务执行成功'),
	(217,'2021-08-21 21:27:53.000000','2021-08-21 21:27:53.000000',3,1,'任务执行成功'),
	(218,'2021-08-21 21:27:54.000000','2021-08-21 21:27:54.000000',3,1,'任务执行成功'),
	(219,'2021-08-21 21:27:55.000000','2021-08-21 21:27:55.000000',3,1,'任务执行成功'),
	(220,'2021-08-21 21:27:56.000000','2021-08-21 21:27:56.000000',3,1,'任务执行成功'),
	(221,'2021-08-21 21:27:57.000000','2021-08-21 21:27:57.000000',3,1,'任务执行成功'),
	(222,'2021-08-21 21:27:58.000000','2021-08-21 21:27:58.000000',3,1,'任务执行成功'),
	(223,'2021-08-21 21:27:59.000000','2021-08-21 21:27:59.000000',3,1,'任务执行成功'),
	(224,'2021-08-21 21:28:00.000000','2021-08-21 21:28:00.000000',3,1,'任务执行成功'),
	(225,'2021-08-21 21:28:01.000000','2021-08-21 21:28:01.000000',3,1,'任务执行成功'),
	(226,'2021-08-21 21:28:02.000000','2021-08-21 21:28:02.000000',3,1,'任务执行成功'),
	(227,'2021-08-21 21:28:03.000000','2021-08-21 21:28:03.000000',3,1,'任务执行成功'),
	(228,'2021-08-21 21:28:04.000000','2021-08-21 21:28:04.000000',3,1,'任务执行成功'),
	(229,'2021-08-21 21:28:05.000000','2021-08-21 21:28:05.000000',3,1,'任务执行成功'),
	(230,'2021-08-21 21:28:06.000000','2021-08-21 21:28:06.000000',3,1,'任务执行成功'),
	(231,'2021-08-21 21:28:07.000000','2021-08-21 21:28:07.000000',3,1,'任务执行成功'),
	(232,'2021-08-21 21:28:08.000000','2021-08-21 21:28:08.000000',3,1,'任务执行成功'),
	(233,'2021-08-21 21:28:09.000000','2021-08-21 21:28:09.000000',3,1,'任务执行成功'),
	(234,'2021-08-21 21:28:10.000000','2021-08-21 21:28:10.000000',3,1,'任务执行成功'),
	(235,'2021-08-21 21:28:11.000000','2021-08-21 21:28:11.000000',3,1,'任务执行成功'),
	(236,'2021-08-21 21:28:12.000000','2021-08-21 21:28:12.000000',3,1,'任务执行成功'),
	(237,'2021-08-21 21:28:13.000000','2021-08-21 21:28:13.000000',3,1,'任务执行成功'),
	(238,'2021-08-21 21:28:14.000000','2021-08-21 21:28:14.000000',3,1,'任务执行成功'),
	(239,'2021-08-21 21:28:15.000000','2021-08-21 21:28:15.000000',3,1,'任务执行成功'),
	(240,'2021-08-21 21:28:16.000000','2021-08-21 21:28:16.000000',3,1,'任务执行成功'),
	(241,'2021-08-21 21:28:17.000000','2021-08-21 21:28:17.000000',3,1,'任务执行成功'),
	(242,'2021-08-21 21:28:18.000000','2021-08-21 21:28:18.000000',3,1,'任务执行成功'),
	(243,'2021-08-21 21:28:19.000000','2021-08-21 21:28:19.000000',3,1,'任务执行成功'),
	(244,'2021-08-21 21:28:20.000000','2021-08-21 21:28:20.000000',3,1,'任务执行成功'),
	(245,'2021-08-21 21:28:21.000000','2021-08-21 21:28:21.000000',3,1,'任务执行成功'),
	(246,'2021-08-21 21:28:22.000000','2021-08-21 21:28:22.000000',3,1,'任务执行成功'),
	(247,'2021-08-21 21:28:23.000000','2021-08-21 21:28:23.000000',3,1,'任务执行成功'),
	(248,'2021-08-21 21:28:24.000000','2021-08-21 21:28:24.000000',3,1,'任务执行成功'),
	(249,'2021-08-21 21:28:25.000000','2021-08-21 21:28:25.000000',3,1,'任务执行成功'),
	(250,'2021-08-21 21:28:26.000000','2021-08-21 21:28:26.000000',3,1,'任务执行成功'),
	(251,'2021-08-21 21:28:27.000000','2021-08-21 21:28:27.000000',3,1,'任务执行成功'),
	(252,'2021-08-21 21:28:28.000000','2021-08-21 21:28:28.000000',3,1,'任务执行成功'),
	(253,'2021-08-21 21:28:29.000000','2021-08-21 21:28:29.000000',3,1,'任务执行成功'),
	(254,'2021-08-21 21:28:30.000000','2021-08-21 21:28:30.000000',3,1,'任务执行成功'),
	(255,'2021-08-21 21:28:31.000000','2021-08-21 21:28:31.000000',3,1,'任务执行成功'),
	(256,'2021-08-21 21:28:32.000000','2021-08-21 21:28:32.000000',3,1,'任务执行成功'),
	(257,'2021-08-21 21:28:33.000000','2021-08-21 21:28:33.000000',3,1,'任务执行成功'),
	(258,'2021-08-21 21:28:34.000000','2021-08-21 21:28:34.000000',3,1,'任务执行成功'),
	(259,'2021-08-21 21:28:35.000000','2021-08-21 21:28:35.000000',3,1,'任务执行成功'),
	(260,'2021-08-21 21:28:36.000000','2021-08-21 21:28:36.000000',3,1,'任务执行成功'),
	(261,'2021-08-21 21:28:37.000000','2021-08-21 21:28:37.000000',3,1,'任务执行成功'),
	(262,'2021-08-21 21:28:38.000000','2021-08-21 21:28:38.000000',3,1,'任务执行成功'),
	(263,'2021-08-21 21:28:39.000000','2021-08-21 21:28:39.000000',3,1,'任务执行成功'),
	(264,'2021-08-21 21:28:40.000000','2021-08-21 21:28:40.000000',3,1,'任务执行成功'),
	(265,'2021-08-21 21:28:41.000000','2021-08-21 21:28:41.000000',3,1,'任务执行成功'),
	(266,'2021-08-21 21:28:42.000000','2021-08-21 21:28:42.000000',3,1,'任务执行成功'),
	(267,'2021-08-21 21:28:43.000000','2021-08-21 21:28:43.000000',3,1,'任务执行成功'),
	(268,'2021-08-21 21:28:44.000000','2021-08-21 21:28:44.000000',3,1,'任务执行成功'),
	(269,'2021-08-21 21:28:45.000000','2021-08-21 21:28:45.000000',3,1,'任务执行成功'),
	(270,'2021-08-21 21:28:46.000000','2021-08-21 21:28:46.000000',3,1,'任务执行成功'),
	(271,'2021-08-21 21:28:47.000000','2021-08-21 21:28:47.000000',3,1,'任务执行成功'),
	(272,'2021-08-21 21:28:48.000000','2021-08-21 21:28:48.000000',3,1,'任务执行成功'),
	(273,'2021-08-21 21:28:49.000000','2021-08-21 21:28:49.000000',3,1,'任务执行成功'),
	(274,'2021-08-21 21:28:50.000000','2021-08-21 21:28:50.000000',3,1,'任务执行成功'),
	(275,'2021-08-21 21:28:51.000000','2021-08-21 21:28:51.000000',3,1,'任务执行成功'),
	(276,'2021-08-21 21:28:52.000000','2021-08-21 21:28:52.000000',3,1,'任务执行成功'),
	(277,'2021-08-21 21:28:53.000000','2021-08-21 21:28:53.000000',3,1,'任务执行成功'),
	(278,'2021-08-21 21:34:27.000000','2021-08-21 21:34:27.000000',3,1,'任务执行成功'),
	(279,'2021-08-21 21:34:28.000000','2021-08-21 21:34:28.000000',3,1,'任务执行成功'),
	(280,'2021-08-21 21:34:29.000000','2021-08-21 21:34:29.000000',3,1,'任务执行成功'),
	(281,'2021-08-21 21:34:30.000000','2021-08-21 21:34:30.000000',3,1,'任务执行成功'),
	(282,'2021-08-21 21:34:31.000000','2021-08-21 21:34:31.000000',3,1,'任务执行成功'),
	(283,'2021-08-21 21:34:32.000000','2021-08-21 21:34:32.000000',3,1,'任务执行成功'),
	(284,'2021-08-21 21:34:33.000000','2021-08-21 21:34:33.000000',3,1,'任务执行成功'),
	(285,'2021-08-21 21:34:34.000000','2021-08-21 21:34:34.000000',3,1,'任务执行成功'),
	(286,'2021-08-21 21:34:35.000000','2021-08-21 21:34:35.000000',3,1,'任务执行成功'),
	(287,'2021-08-21 21:34:36.000000','2021-08-21 21:34:36.000000',3,1,'任务执行成功'),
	(288,'2021-08-21 21:34:37.000000','2021-08-21 21:34:37.000000',3,1,'任务执行成功'),
	(289,'2021-08-21 21:34:38.000000','2021-08-21 21:34:38.000000',3,1,'任务执行成功'),
	(290,'2021-08-21 21:34:39.000000','2021-08-21 21:34:39.000000',3,1,'任务执行成功'),
	(291,'2021-08-21 21:34:40.000000','2021-08-21 21:34:40.000000',3,1,'任务执行成功'),
	(292,'2021-08-21 21:34:41.000000','2021-08-21 21:34:41.000000',3,1,'任务执行成功'),
	(293,'2021-08-21 21:34:42.000000','2021-08-21 21:34:42.000000',3,1,'任务执行成功'),
	(294,'2021-08-21 21:34:45.000000','2021-08-21 21:34:45.000000',3,1,'任务执行成功'),
	(295,'2021-08-21 21:34:46.000000','2021-08-21 21:34:46.000000',3,1,'任务执行成功'),
	(296,'2021-08-21 21:34:47.000000','2021-08-21 21:34:47.000000',3,1,'任务执行成功'),
	(297,'2021-08-21 21:34:48.000000','2021-08-21 21:34:48.000000',3,1,'任务执行成功'),
	(298,'2021-08-21 21:36:00.000000','2021-08-21 21:36:00.000000',3,1,'任务执行成功'),
	(299,'2021-08-21 21:36:01.000000','2021-08-21 21:36:01.000000',3,1,'任务执行成功'),
	(300,'2021-08-21 21:36:02.000000','2021-08-21 21:36:02.000000',3,1,'任务执行成功'),
	(301,'2021-08-21 21:36:03.000000','2021-08-21 21:36:03.000000',3,1,'任务执行成功'),
	(302,'2021-08-21 21:36:04.000000','2021-08-21 21:36:04.000000',3,1,'任务执行成功'),
	(303,'2021-08-21 21:36:05.000000','2021-08-21 21:36:05.000000',3,1,'任务执行成功'),
	(304,'2021-08-21 21:36:06.000000','2021-08-21 21:36:06.000000',3,1,'任务执行成功'),
	(305,'2021-08-21 21:36:07.000000','2021-08-21 21:36:07.000000',3,1,'任务执行成功'),
	(306,'2021-08-21 21:36:08.000000','2021-08-21 21:36:08.000000',3,1,'任务执行成功'),
	(307,'2021-08-21 21:36:09.000000','2021-08-21 21:36:09.000000',3,1,'任务执行成功'),
	(308,'2021-08-21 21:36:10.000000','2021-08-21 21:36:10.000000',3,1,'任务执行成功'),
	(309,'2021-08-21 21:36:11.000000','2021-08-21 21:36:11.000000',3,1,'任务执行成功'),
	(310,'2021-08-21 21:36:12.000000','2021-08-21 21:36:12.000000',3,1,'任务执行成功'),
	(311,'2021-08-21 21:36:13.000000','2021-08-21 21:36:13.000000',3,1,'任务执行成功'),
	(312,'2021-08-21 21:36:14.000000','2021-08-21 21:36:14.000000',3,1,'任务执行成功'),
	(313,'2021-08-21 21:36:15.000000','2021-08-21 21:36:15.000000',3,1,'任务执行成功'),
	(314,'2021-08-21 21:36:16.000000','2021-08-21 21:36:16.000000',3,1,'任务执行成功'),
	(315,'2021-08-21 21:36:17.000000','2021-08-21 21:36:17.000000',3,1,'任务执行成功'),
	(316,'2021-08-21 21:36:18.000000','2021-08-21 21:36:18.000000',3,1,'任务执行成功'),
	(317,'2021-08-21 21:36:19.000000','2021-08-21 21:36:19.000000',3,1,'任务执行成功'),
	(318,'2021-08-21 21:36:20.000000','2021-08-21 21:36:20.000000',3,1,'任务执行成功'),
	(319,'2021-08-21 21:36:21.000000','2021-08-21 21:36:21.000000',3,1,'任务执行成功'),
	(320,'2021-08-21 21:36:22.000000','2021-08-21 21:36:22.000000',3,1,'任务执行成功'),
	(321,'2021-08-21 21:36:23.000000','2021-08-21 21:36:23.000000',3,1,'任务执行成功'),
	(322,'2021-08-21 21:36:27.000000','2021-08-21 21:36:27.000000',3,1,'任务执行成功'),
	(323,'2021-08-21 21:36:28.000000','2021-08-21 21:36:28.000000',3,1,'任务执行成功'),
	(324,'2021-08-21 21:36:29.000000','2021-08-21 21:36:29.000000',3,1,'任务执行成功'),
	(325,'2021-08-21 21:36:30.000000','2021-08-21 21:36:30.000000',3,1,'任务执行成功'),
	(326,'2021-08-21 21:36:31.000000','2021-08-21 21:36:31.000000',3,1,'任务执行成功'),
	(327,'2021-08-21 21:36:32.000000','2021-08-21 21:36:32.000000',3,1,'任务执行成功'),
	(328,'2021-08-21 21:36:35.000000','2021-08-21 21:36:35.000000',3,1,'任务执行成功'),
	(329,'2021-08-21 21:36:36.000000','2021-08-21 21:36:36.000000',3,1,'任务执行成功'),
	(330,'2021-08-21 21:36:37.000000','2021-08-21 21:36:37.000000',3,1,'任务执行成功'),
	(331,'2021-08-21 21:36:38.000000','2021-08-21 21:36:38.000000',3,1,'任务执行成功'),
	(332,'2021-08-21 21:36:39.000000','2021-08-21 21:36:39.000000',3,1,'任务执行成功'),
	(333,'2021-08-21 21:36:40.000000','2021-08-21 21:36:40.000000',3,1,'任务执行成功'),
	(334,'2021-08-21 21:36:41.000000','2021-08-21 21:36:41.000000',3,1,'任务执行成功'),
	(335,'2021-08-21 21:36:42.000000','2021-08-21 21:36:42.000000',3,1,'任务执行成功'),
	(336,'2021-08-21 21:36:43.000000','2021-08-21 21:36:43.000000',3,1,'任务执行成功'),
	(337,'2021-08-21 21:36:44.000000','2021-08-21 21:36:44.000000',3,1,'任务执行成功'),
	(338,'2021-08-21 21:36:45.000000','2021-08-21 21:36:45.000000',3,1,'任务执行成功'),
	(339,'2021-08-21 21:36:46.000000','2021-08-21 21:36:46.000000',3,1,'任务执行成功'),
	(340,'2021-08-21 21:36:47.000000','2021-08-21 21:36:47.000000',3,1,'任务执行成功'),
	(341,'2021-08-21 21:36:48.000000','2021-08-21 21:36:48.000000',3,1,'任务执行成功'),
	(342,'2021-08-21 21:36:49.000000','2021-08-21 21:36:49.000000',3,1,'任务执行成功'),
	(343,'2021-08-21 21:36:50.000000','2021-08-21 21:36:50.000000',3,1,'任务执行成功'),
	(344,'2021-08-21 21:36:51.000000','2021-08-21 21:36:51.000000',3,1,'任务执行成功'),
	(345,'2021-08-21 21:36:52.000000','2021-08-21 21:36:52.000000',3,1,'任务执行成功'),
	(346,'2021-08-21 21:36:53.000000','2021-08-21 21:36:53.000000',3,1,'任务执行成功'),
	(347,'2021-08-21 21:36:54.000000','2021-08-21 21:36:54.000000',3,1,'任务执行成功'),
	(348,'2021-08-21 21:36:55.000000','2021-08-21 21:36:55.000000',3,1,'任务执行成功'),
	(349,'2021-08-21 21:36:56.000000','2021-08-21 21:36:56.000000',3,1,'任务执行成功'),
	(350,'2021-08-21 21:36:57.000000','2021-08-21 21:36:57.000000',3,1,'任务执行成功'),
	(351,'2021-08-21 21:36:58.000000','2021-08-21 21:36:58.000000',3,1,'任务执行成功'),
	(352,'2021-08-21 21:36:59.000000','2021-08-21 21:36:59.000000',3,1,'任务执行成功'),
	(353,'2021-08-21 21:38:00.000000','2021-08-21 21:38:00.000000',3,1,'任务执行成功'),
	(354,'2021-08-21 21:38:01.000000','2021-08-21 21:38:01.000000',3,1,'任务执行成功'),
	(355,'2021-08-21 21:38:02.000000','2021-08-21 21:38:02.000000',3,1,'任务执行成功'),
	(356,'2021-08-21 21:38:03.000000','2021-08-21 21:38:03.000000',3,1,'任务执行成功'),
	(357,'2021-08-21 21:38:04.000000','2021-08-21 21:38:04.000000',3,1,'任务执行成功'),
	(358,'2021-08-21 21:38:05.000000','2021-08-21 21:38:05.000000',3,1,'任务执行成功'),
	(359,'2021-08-21 21:38:06.000000','2021-08-21 21:38:06.000000',3,1,'任务执行成功'),
	(360,'2021-08-21 21:38:07.000000','2021-08-21 21:38:07.000000',3,1,'任务执行成功'),
	(361,'2021-08-21 21:38:08.000000','2021-08-21 21:38:08.000000',3,1,'任务执行成功'),
	(362,'2021-08-21 21:38:09.000000','2021-08-21 21:38:09.000000',3,1,'任务执行成功'),
	(363,'2021-08-21 22:00:01.000000','2021-08-21 22:00:01.000000',3,1,'任务执行成功'),
	(364,'2021-08-21 22:00:01.000000','2021-08-21 22:00:01.000000',3,1,'任务执行成功'),
	(365,'2021-08-21 22:00:02.000000','2021-08-21 22:00:02.000000',3,1,'任务执行成功'),
	(366,'2021-08-21 22:00:03.000000','2021-08-21 22:00:03.000000',3,1,'任务执行成功'),
	(367,'2021-08-21 22:00:05.000000','2021-08-21 22:00:05.000000',3,1,'任务执行成功'),
	(368,'2021-08-21 22:00:06.000000','2021-08-21 22:00:06.000000',3,1,'任务执行成功'),
	(369,'2021-08-21 22:00:06.000000','2021-08-21 22:00:06.000000',3,1,'任务执行成功'),
	(370,'2021-08-21 22:00:08.000000','2021-08-21 22:00:08.000000',3,1,'任务执行成功'),
	(371,'2021-08-21 22:00:09.000000','2021-08-21 22:00:09.000000',3,1,'任务执行成功'),
	(372,'2021-08-21 22:00:10.000000','2021-08-21 22:00:10.000000',3,1,'任务执行成功'),
	(373,'2021-08-21 22:00:11.000000','2021-08-21 22:00:11.000000',3,1,'任务执行成功'),
	(374,'2021-08-21 22:00:11.000000','2021-08-21 22:00:11.000000',3,1,'任务执行成功'),
	(375,'2021-08-21 22:00:13.000000','2021-08-21 22:00:13.000000',3,1,'任务执行成功'),
	(376,'2021-08-21 22:00:14.000000','2021-08-21 22:00:14.000000',3,1,'任务执行成功'),
	(377,'2021-08-21 22:00:15.000000','2021-08-21 22:00:15.000000',3,1,'任务执行成功'),
	(378,'2021-08-21 22:00:16.000000','2021-08-21 22:00:16.000000',3,1,'任务执行成功'),
	(379,'2021-08-21 22:00:16.000000','2021-08-21 22:00:16.000000',3,1,'任务执行成功'),
	(380,'2021-08-21 22:00:18.000000','2021-08-21 22:00:18.000000',3,1,'任务执行成功'),
	(381,'2021-08-21 22:00:18.000000','2021-08-21 22:00:18.000000',3,1,'任务执行成功'),
	(382,'2021-08-21 22:00:20.000000','2021-08-21 22:00:20.000000',3,1,'任务执行成功'),
	(383,'2021-08-21 22:00:21.000000','2021-08-21 22:00:21.000000',3,1,'任务执行成功'),
	(384,'2021-08-21 22:00:21.000000','2021-08-21 22:00:21.000000',3,1,'任务执行成功'),
	(385,'2021-08-21 22:00:23.000000','2021-08-21 22:00:23.000000',3,1,'任务执行成功'),
	(386,'2021-08-21 22:00:24.000000','2021-08-21 22:00:24.000000',3,1,'任务执行成功'),
	(387,'2021-08-21 22:00:25.000000','2021-08-21 22:00:25.000000',3,1,'任务执行成功'),
	(388,'2021-08-21 22:00:25.000000','2021-08-21 22:00:25.000000',3,1,'任务执行成功'),
	(389,'2021-08-21 22:00:27.000000','2021-08-21 22:00:27.000000',3,1,'任务执行成功'),
	(390,'2021-08-21 22:00:28.000000','2021-08-21 22:00:28.000000',3,1,'任务执行成功'),
	(391,'2021-08-21 22:00:29.000000','2021-08-21 22:00:29.000000',3,1,'任务执行成功'),
	(392,'2021-08-21 22:00:29.000000','2021-08-21 22:00:29.000000',3,1,'任务执行成功'),
	(393,'2021-08-21 22:00:31.000000','2021-08-21 22:00:31.000000',3,1,'任务执行成功'),
	(394,'2021-08-21 22:00:32.000000','2021-08-21 22:00:32.000000',3,1,'任务执行成功'),
	(395,'2021-08-21 22:00:33.000000','2021-08-21 22:00:33.000000',3,1,'任务执行成功'),
	(396,'2021-08-21 22:00:34.000000','2021-08-21 22:00:34.000000',3,1,'任务执行成功'),
	(397,'2021-08-21 22:00:35.000000','2021-08-21 22:00:35.000000',3,1,'任务执行成功'),
	(398,'2021-08-21 22:00:36.000000','2021-08-21 22:00:36.000000',3,1,'任务执行成功'),
	(399,'2021-08-21 22:00:37.000000','2021-08-21 22:00:37.000000',3,1,'任务执行成功'),
	(400,'2021-08-21 22:00:37.000000','2021-08-21 22:00:37.000000',3,1,'任务执行成功'),
	(401,'2021-08-21 22:00:39.000000','2021-08-21 22:00:39.000000',3,1,'任务执行成功'),
	(402,'2021-08-21 22:00:40.000000','2021-08-21 22:00:40.000000',3,1,'任务执行成功'),
	(403,'2021-08-21 22:00:40.000000','2021-08-21 22:00:40.000000',3,1,'任务执行成功'),
	(404,'2021-08-21 22:00:42.000000','2021-08-21 22:00:42.000000',3,1,'任务执行成功'),
	(405,'2021-08-21 22:00:42.000000','2021-08-21 22:00:42.000000',3,1,'任务执行成功'),
	(406,'2021-08-21 22:00:44.000000','2021-08-21 22:00:44.000000',3,1,'任务执行成功'),
	(407,'2021-08-21 22:00:45.000000','2021-08-21 22:00:45.000000',3,1,'任务执行成功'),
	(408,'2021-08-21 22:00:45.000000','2021-08-21 22:00:45.000000',3,1,'任务执行成功'),
	(409,'2021-08-21 22:00:47.000000','2021-08-21 22:00:47.000000',3,1,'任务执行成功'),
	(410,'2021-08-21 22:00:47.000000','2021-08-21 22:00:47.000000',3,1,'任务执行成功'),
	(411,'2021-08-21 22:00:49.000000','2021-08-21 22:00:49.000000',3,1,'任务执行成功'),
	(412,'2021-08-21 22:00:50.000000','2021-08-21 22:00:50.000000',3,1,'任务执行成功'),
	(413,'2021-08-21 22:00:51.000000','2021-08-21 22:00:51.000000',3,1,'任务执行成功'),
	(414,'2021-08-21 22:00:52.000000','2021-08-21 22:00:52.000000',3,1,'任务执行成功'),
	(415,'2021-08-21 22:00:52.000000','2021-08-21 22:00:52.000000',3,1,'任务执行成功'),
	(416,'2021-08-21 22:00:53.000000','2021-08-21 22:00:53.000000',3,1,'任务执行成功'),
	(417,'2021-08-21 22:00:55.000000','2021-08-21 22:00:55.000000',3,1,'任务执行成功'),
	(418,'2021-08-21 22:00:56.000000','2021-08-21 22:00:56.000000',3,1,'任务执行成功'),
	(419,'2021-08-21 22:00:57.000000','2021-08-21 22:00:57.000000',3,1,'任务执行成功'),
	(420,'2021-08-21 22:00:58.000000','2021-08-21 22:00:58.000000',3,1,'任务执行成功'),
	(421,'2021-08-21 22:00:59.000000','2021-08-21 22:00:59.000000',3,1,'任务执行成功'),
	(422,'2021-08-21 22:00:59.000000','2021-08-21 22:00:59.000000',3,1,'任务执行成功'),
	(423,'2021-08-21 22:40:00.000000','2021-08-21 22:40:00.000000',3,1,'任务执行成功'),
	(424,'2021-08-21 22:45:00.000000','2021-08-21 22:45:00.000000',3,1,'任务执行成功'),
	(425,'2021-08-21 22:50:00.000000','2021-08-21 22:50:00.000000',3,1,'任务执行成功'),
	(426,'2021-08-21 22:55:00.000000','2021-08-21 22:55:00.000000',3,1,'任务执行成功'),
	(427,'2021-08-21 23:01:00.000000','2021-08-21 23:01:00.000000',3,1,'任务执行成功'),
	(428,'2021-08-21 23:02:00.000000','2021-08-21 23:02:00.000000',3,1,'任务执行成功'),
	(429,'2021-08-21 23:03:00.000000','2021-08-21 23:03:00.000000',3,1,'任务执行成功'),
	(430,'2021-08-21 23:04:00.000000','2021-08-21 23:04:00.000000',3,1,'任务执行成功'),
	(431,'2021-08-21 23:05:00.000000','2021-08-21 23:05:00.000000',3,1,'任务执行成功'),
	(432,'2021-08-21 23:06:00.000000','2021-08-21 23:06:00.000000',3,1,'任务执行成功'),
	(433,'2021-08-21 23:07:00.000000','2021-08-21 23:07:00.000000',3,1,'任务执行成功'),
	(434,'2021-08-21 23:08:00.000000','2021-08-21 23:08:00.000000',3,1,'任务执行成功'),
	(435,'2021-08-21 23:10:00.000000','2021-08-21 23:10:00.000000',3,1,'任务执行成功'),
	(436,'2021-08-21 23:12:00.000000','2021-08-21 23:12:00.000000',3,1,'任务执行成功'),
	(437,'2021-08-21 23:14:00.000000','2021-08-21 23:14:00.000000',3,1,'任务执行成功'),
	(438,'2021-08-21 23:16:00.000000','2021-08-21 23:16:00.000000',3,1,'任务执行成功'),
	(439,'2021-08-21 23:18:00.000000','2021-08-21 23:18:00.000000',3,1,'任务执行成功'),
	(440,'2021-08-21 23:20:00.000000','2021-08-21 23:20:00.000000',3,1,'任务执行成功'),
	(441,'2021-08-21 23:22:00.000000','2021-08-21 23:22:00.000000',3,1,'任务执行成功'),
	(442,'2021-08-21 23:24:00.000000','2021-08-21 23:24:00.000000',3,1,'任务执行成功'),
	(443,'2021-08-21 23:34:00.000000','2021-08-21 23:34:00.000000',3,1,'任务执行成功'),
	(444,'2021-08-21 23:36:00.000000','2021-08-21 23:36:00.000000',3,1,'任务执行成功'),
	(445,'2021-08-21 23:38:00.000000','2021-08-21 23:38:00.000000',3,1,'任务执行成功'),
	(446,'2021-08-23 20:32:00.000000','2021-08-23 20:32:00.000000',3,1,'任务执行成功'),
	(447,'2021-08-23 22:40:00.000000','2021-08-23 22:40:00.000000',3,1,'任务执行成功'),
	(448,'2021-08-23 22:42:00.000000','2021-08-23 22:42:00.000000',3,1,'任务执行成功'),
	(449,'2021-08-23 22:44:00.000000','2021-08-23 22:44:00.000000',3,1,'任务执行成功'),
	(450,'2021-08-23 22:46:00.000000','2021-08-23 22:46:00.000000',3,1,'任务执行成功'),
	(451,'2021-08-23 22:48:00.000000','2021-08-23 22:48:00.000000',3,1,'任务执行成功'),
	(452,'2021-08-23 22:50:00.000000','2021-08-23 22:50:00.000000',3,1,'任务执行成功'),
	(453,'2021-08-23 22:52:00.000000','2021-08-23 22:52:00.000000',3,1,'任务执行成功'),
	(454,'2021-08-23 22:54:00.000000','2021-08-23 22:54:00.000000',3,1,'任务执行成功'),
	(455,'2021-08-23 22:56:00.000000','2021-08-23 22:56:00.000000',3,1,'任务执行成功'),
	(456,'2021-08-23 22:58:00.000000','2021-08-23 22:58:00.000000',3,1,'任务执行成功'),
	(457,'2021-08-23 23:00:00.000000','2021-08-23 23:00:00.000000',3,1,'任务执行成功'),
	(458,'2021-08-23 23:02:00.000000','2021-08-23 23:02:00.000000',3,1,'任务执行成功'),
	(459,'2021-08-23 23:04:00.000000','2021-08-23 23:04:00.000000',3,1,'任务执行成功'),
	(460,'2021-08-23 23:06:00.000000','2021-08-23 23:06:00.000000',3,1,'任务执行成功'),
	(461,'2021-08-23 23:08:00.000000','2021-08-23 23:08:00.000000',3,1,'任务执行成功'),
	(462,'2021-08-23 23:10:00.000000','2021-08-23 23:10:00.000000',3,1,'任务执行成功'),
	(463,'2021-08-23 23:12:00.000000','2021-08-23 23:12:00.000000',3,1,'任务执行成功'),
	(464,'2021-08-23 23:14:00.000000','2021-08-23 23:14:00.000000',3,1,'任务执行成功'),
	(465,'2021-08-24 11:38:00.000000','2021-08-24 11:38:00.000000',3,1,'任务执行成功'),
	(466,'2021-08-24 11:40:00.000000','2021-08-24 11:40:00.000000',3,1,'任务执行成功'),
	(467,'2021-08-24 11:42:00.000000','2021-08-24 11:42:00.000000',3,1,'任务执行成功'),
	(468,'2021-08-24 11:44:00.000000','2021-08-24 11:44:00.000000',3,1,'任务执行成功'),
	(469,'2021-08-24 11:46:00.000000','2021-08-24 11:46:00.000000',3,1,'任务执行成功'),
	(470,'2021-08-24 11:48:00.000000','2021-08-24 11:48:00.000000',3,1,'任务执行成功'),
	(471,'2021-08-24 11:50:00.000000','2021-08-24 11:50:00.000000',3,1,'任务执行成功'),
	(472,'2021-08-24 11:53:10.000000','2021-08-24 11:53:10.000000',3,1,'任务执行成功'),
	(473,'2021-08-24 11:54:00.000000','2021-08-24 11:54:00.000000',3,1,'任务执行成功'),
	(474,'2021-08-24 11:56:01.000000','2021-08-24 11:56:01.000000',3,1,'任务执行成功'),
	(475,'2021-08-24 11:58:00.000000','2021-08-24 11:58:00.000000',3,1,'任务执行成功'),
	(476,'2021-08-24 12:00:00.000000','2021-08-24 12:00:00.000000',3,1,'任务执行成功'),
	(477,'2021-08-24 12:02:00.000000','2021-08-24 12:02:00.000000',3,1,'任务执行成功'),
	(478,'2021-08-24 12:04:00.000000','2021-08-24 12:04:00.000000',3,1,'任务执行成功'),
	(479,'2021-08-24 12:06:00.000000','2021-08-24 12:06:00.000000',3,1,'任务执行成功'),
	(480,'2021-08-24 12:08:00.000000','2021-08-24 12:08:00.000000',3,1,'任务执行成功'),
	(481,'2021-08-24 12:10:00.000000','2021-08-24 12:10:00.000000',3,1,'任务执行成功'),
	(482,'2021-08-24 12:14:00.000000','2021-08-24 12:14:00.000000',3,1,'任务执行成功'),
	(483,'2021-08-24 12:16:00.000000','2021-08-24 12:16:00.000000',3,1,'任务执行成功'),
	(484,'2021-08-24 12:18:00.000000','2021-08-24 12:18:00.000000',3,1,'任务执行成功'),
	(485,'2021-08-24 12:20:00.000000','2021-08-24 12:20:00.000000',3,1,'任务执行成功'),
	(486,'2021-08-24 12:22:00.000000','2021-08-24 12:22:00.000000',3,1,'任务执行成功'),
	(487,'2021-08-24 12:26:00.000000','2021-08-24 12:26:00.000000',3,1,'任务执行成功'),
	(488,'2021-08-24 12:28:00.000000','2021-08-24 12:28:00.000000',3,1,'任务执行成功'),
	(489,'2021-08-24 12:30:00.000000','2021-08-24 12:30:00.000000',3,1,'任务执行成功'),
	(490,'2021-08-24 12:32:00.000000','2021-08-24 12:32:00.000000',3,1,'任务执行成功'),
	(491,'2021-08-24 12:34:00.000000','2021-08-24 12:34:00.000000',3,1,'任务执行成功'),
	(492,'2021-08-24 12:36:00.000000','2021-08-24 12:36:00.000000',3,1,'任务执行成功'),
	(493,'2021-08-24 12:38:01.000000','2021-08-24 12:38:01.000000',3,1,'任务执行成功'),
	(494,'2021-08-24 12:40:01.000000','2021-08-24 12:40:01.000000',3,1,'任务执行成功'),
	(495,'2021-08-24 12:42:01.000000','2021-08-24 12:42:01.000000',3,1,'任务执行成功'),
	(496,'2021-08-24 12:44:00.000000','2021-08-24 12:44:00.000000',3,1,'任务执行成功'),
	(497,'2021-08-24 12:46:00.000000','2021-08-24 12:46:00.000000',3,1,'任务执行成功'),
	(498,'2021-08-24 12:48:00.000000','2021-08-24 12:48:00.000000',3,1,'任务执行成功'),
	(499,'2021-08-24 12:50:00.000000','2021-08-24 12:50:00.000000',3,1,'任务执行成功'),
	(500,'2021-08-24 12:52:01.000000','2021-08-24 12:52:01.000000',3,1,'任务执行成功'),
	(501,'2021-08-24 12:54:01.000000','2021-08-24 12:54:01.000000',3,1,'任务执行成功'),
	(502,'2021-08-24 12:56:00.000000','2021-08-24 12:56:00.000000',3,1,'任务执行成功'),
	(503,'2021-08-24 12:58:00.000000','2021-08-24 12:58:00.000000',3,1,'任务执行成功'),
	(504,'2021-08-24 13:00:00.000000','2021-08-24 13:00:00.000000',3,1,'任务执行成功'),
	(505,'2021-08-24 13:02:00.000000','2021-08-24 13:02:00.000000',3,1,'任务执行成功'),
	(506,'2021-08-24 13:04:01.000000','2021-08-24 13:04:01.000000',3,1,'任务执行成功'),
	(507,'2021-08-24 13:06:01.000000','2021-08-24 13:06:01.000000',3,1,'任务执行成功'),
	(508,'2021-08-24 13:08:01.000000','2021-08-24 13:08:01.000000',3,1,'任务执行成功'),
	(509,'2021-08-24 13:10:01.000000','2021-08-24 13:10:01.000000',3,1,'任务执行成功'),
	(510,'2021-08-24 13:12:01.000000','2021-08-24 13:12:01.000000',3,1,'任务执行成功'),
	(511,'2021-08-24 13:14:01.000000','2021-08-24 13:14:01.000000',3,1,'任务执行成功'),
	(512,'2021-08-24 13:16:01.000000','2021-08-24 13:16:01.000000',3,1,'任务执行成功'),
	(513,'2021-08-24 13:18:01.000000','2021-08-24 13:18:01.000000',3,1,'任务执行成功'),
	(514,'2021-08-24 13:20:01.000000','2021-08-24 13:20:01.000000',3,1,'任务执行成功'),
	(515,'2021-08-24 13:22:00.000000','2021-08-24 13:22:00.000000',3,1,'任务执行成功'),
	(516,'2021-08-24 13:22:01.000000','2021-08-24 13:22:01.000000',3,1,'任务执行成功'),
	(517,'2021-08-24 13:24:00.000000','2021-08-24 13:24:00.000000',3,1,'任务执行成功'),
	(518,'2021-08-24 13:24:01.000000','2021-08-24 13:24:01.000000',3,1,'任务执行成功'),
	(519,'2021-08-24 13:26:00.000000','2021-08-24 13:26:00.000000',3,1,'任务执行成功'),
	(520,'2021-08-24 13:26:01.000000','2021-08-24 13:26:01.000000',3,1,'任务执行成功'),
	(521,'2021-08-24 13:28:00.000000','2021-08-24 13:28:00.000000',3,1,'任务执行成功'),
	(522,'2021-08-24 13:30:01.000000','2021-08-24 13:30:01.000000',3,1,'任务执行成功'),
	(523,'2021-08-24 13:32:00.000000','2021-08-24 13:32:00.000000',3,1,'任务执行成功'),
	(524,'2021-08-24 13:34:00.000000','2021-08-24 13:34:00.000000',3,1,'任务执行成功'),
	(525,'2021-08-24 13:36:00.000000','2021-08-24 13:36:00.000000',3,1,'任务执行成功'),
	(526,'2021-08-24 13:38:00.000000','2021-08-24 13:38:00.000000',3,1,'任务执行成功'),
	(527,'2021-08-24 13:40:00.000000','2021-08-24 13:40:00.000000',3,1,'任务执行成功'),
	(528,'2021-08-24 13:42:00.000000','2021-08-24 13:42:00.000000',3,1,'任务执行成功'),
	(529,'2021-08-24 13:44:00.000000','2021-08-24 13:44:00.000000',3,1,'任务执行成功'),
	(530,'2021-08-24 13:46:00.000000','2021-08-24 13:46:00.000000',3,1,'任务执行成功'),
	(531,'2021-08-24 13:48:02.000000','2021-08-24 13:48:02.000000',3,1,'任务执行成功'),
	(532,'2021-08-24 13:50:00.000000','2021-08-24 13:50:00.000000',3,1,'任务执行成功'),
	(533,'2021-08-24 13:52:00.000000','2021-08-24 13:52:00.000000',3,1,'任务执行成功'),
	(534,'2021-08-24 13:54:00.000000','2021-08-24 13:54:00.000000',3,1,'任务执行成功'),
	(535,'2021-08-24 13:56:00.000000','2021-08-24 13:56:00.000000',3,1,'任务执行成功'),
	(536,'2021-08-24 13:58:00.000000','2021-08-24 13:58:00.000000',3,1,'任务执行成功'),
	(537,'2021-08-24 14:00:00.000000','2021-08-24 14:00:00.000000',3,1,'任务执行成功'),
	(538,'2021-08-24 14:02:00.000000','2021-08-24 14:02:00.000000',3,1,'任务执行成功'),
	(539,'2021-08-24 14:04:00.000000','2021-08-24 14:04:00.000000',3,1,'任务执行成功'),
	(540,'2021-08-24 14:06:00.000000','2021-08-24 14:06:00.000000',3,1,'任务执行成功'),
	(541,'2021-08-24 14:08:00.000000','2021-08-24 14:08:00.000000',3,1,'任务执行成功'),
	(542,'2021-08-24 14:10:00.000000','2021-08-24 14:10:00.000000',3,1,'任务执行成功'),
	(543,'2021-08-24 14:12:00.000000','2021-08-24 14:12:00.000000',3,1,'任务执行成功'),
	(544,'2021-08-24 14:14:00.000000','2021-08-24 14:14:00.000000',3,1,'任务执行成功'),
	(545,'2021-08-24 14:16:00.000000','2021-08-24 14:16:00.000000',3,1,'任务执行成功'),
	(546,'2021-08-24 14:18:00.000000','2021-08-24 14:18:00.000000',3,1,'任务执行成功'),
	(547,'2021-08-24 14:20:00.000000','2021-08-24 14:20:00.000000',3,1,'任务执行成功'),
	(548,'2021-08-24 14:22:00.000000','2021-08-24 14:22:00.000000',3,1,'任务执行成功'),
	(549,'2021-08-24 14:24:00.000000','2021-08-24 14:24:00.000000',3,1,'任务执行成功'),
	(550,'2021-08-24 14:26:00.000000','2021-08-24 14:26:00.000000',3,1,'任务执行成功'),
	(551,'2021-08-24 14:28:00.000000','2021-08-24 14:28:00.000000',3,1,'任务执行成功'),
	(552,'2021-08-24 14:30:00.000000','2021-08-24 14:30:00.000000',3,1,'任务执行成功'),
	(553,'2021-08-24 14:32:00.000000','2021-08-24 14:32:00.000000',3,1,'任务执行成功'),
	(554,'2021-08-24 14:34:00.000000','2021-08-24 14:34:00.000000',3,1,'任务执行成功'),
	(555,'2021-08-24 14:48:00.000000','2021-08-24 14:48:00.000000',3,1,'任务执行成功'),
	(556,'2021-08-24 14:50:00.000000','2021-08-24 14:50:00.000000',3,1,'任务执行成功'),
	(557,'2021-08-24 14:52:00.000000','2021-08-24 14:52:00.000000',3,1,'任务执行成功'),
	(558,'2021-08-24 14:54:00.000000','2021-08-24 14:54:00.000000',3,1,'任务执行成功'),
	(559,'2021-08-25 15:30:00.000000','2021-08-25 15:30:00.000000',3,1,'任务执行成功'),
	(560,'2021-08-25 15:32:00.000000','2021-08-25 15:32:00.000000',3,1,'任务执行成功'),
	(561,'2021-08-25 15:34:00.000000','2021-08-25 15:34:00.000000',3,1,'任务执行成功'),
	(562,'2021-08-25 15:36:00.000000','2021-08-25 15:36:00.000000',3,1,'任务执行成功'),
	(563,'2021-08-25 15:38:00.000000','2021-08-25 15:38:00.000000',3,1,'任务执行成功'),
	(564,'2021-08-25 15:40:00.000000','2021-08-25 15:40:00.000000',3,1,'任务执行成功'),
	(565,'2021-08-25 15:42:00.000000','2021-08-25 15:42:00.000000',3,1,'任务执行成功'),
	(566,'2021-08-25 15:44:00.000000','2021-08-25 15:44:00.000000',3,1,'任务执行成功'),
	(567,'2021-08-25 15:46:00.000000','2021-08-25 15:46:00.000000',3,1,'任务执行成功'),
	(568,'2021-08-25 15:48:00.000000','2021-08-25 15:48:00.000000',3,1,'任务执行成功'),
	(569,'2021-08-25 15:50:00.000000','2021-08-25 15:50:00.000000',3,1,'任务执行成功'),
	(570,'2021-08-25 15:52:00.000000','2021-08-25 15:52:00.000000',3,1,'任务执行成功'),
	(571,'2021-08-25 15:54:00.000000','2021-08-25 15:54:00.000000',3,1,'任务执行成功'),
	(572,'2021-08-25 15:56:00.000000','2021-08-25 15:56:00.000000',3,1,'任务执行成功'),
	(573,'2021-08-25 15:58:00.000000','2021-08-25 15:58:00.000000',3,1,'任务执行成功'),
	(574,'2021-08-25 16:00:00.000000','2021-08-25 16:00:00.000000',3,1,'任务执行成功'),
	(575,'2021-08-25 16:02:00.000000','2021-08-25 16:02:00.000000',3,1,'任务执行成功'),
	(576,'2021-08-25 16:04:00.000000','2021-08-25 16:04:00.000000',3,1,'任务执行成功'),
	(577,'2021-08-25 16:06:00.000000','2021-08-25 16:06:00.000000',3,1,'任务执行成功'),
	(578,'2021-08-25 16:08:00.000000','2021-08-25 16:08:00.000000',3,1,'任务执行成功'),
	(579,'2021-08-25 16:10:00.000000','2021-08-25 16:10:00.000000',3,1,'任务执行成功'),
	(580,'2021-08-25 16:12:00.000000','2021-08-25 16:12:00.000000',3,1,'任务执行成功'),
	(581,'2021-08-25 16:14:00.000000','2021-08-25 16:14:00.000000',3,1,'任务执行成功'),
	(582,'2021-08-25 16:16:00.000000','2021-08-25 16:16:00.000000',3,1,'任务执行成功'),
	(583,'2021-08-25 16:18:00.000000','2021-08-25 16:18:00.000000',3,1,'任务执行成功'),
	(584,'2021-08-25 16:20:00.000000','2021-08-25 16:20:00.000000',3,1,'任务执行成功'),
	(585,'2021-08-25 16:22:00.000000','2021-08-25 16:22:00.000000',3,1,'任务执行成功'),
	(586,'2021-08-25 16:24:00.000000','2021-08-25 16:24:00.000000',3,1,'任务执行成功'),
	(587,'2021-08-25 16:26:00.000000','2021-08-25 16:26:00.000000',3,1,'任务执行成功'),
	(588,'2021-08-25 16:28:00.000000','2021-08-25 16:28:00.000000',3,1,'任务执行成功'),
	(589,'2021-08-25 16:30:00.000000','2021-08-25 16:30:00.000000',3,1,'任务执行成功'),
	(590,'2021-08-25 16:32:00.000000','2021-08-25 16:32:00.000000',3,1,'任务执行成功'),
	(591,'2021-08-25 16:34:00.000000','2021-08-25 16:34:00.000000',3,1,'任务执行成功'),
	(592,'2021-08-25 16:36:00.000000','2021-08-25 16:36:00.000000',3,1,'任务执行成功'),
	(593,'2021-08-25 16:38:00.000000','2021-08-25 16:38:00.000000',3,1,'任务执行成功'),
	(594,'2021-08-25 16:40:00.000000','2021-08-25 16:40:00.000000',3,1,'任务执行成功'),
	(595,'2021-08-25 16:42:00.000000','2021-08-25 16:42:00.000000',3,1,'任务执行成功'),
	(596,'2021-08-25 16:44:00.000000','2021-08-25 16:44:00.000000',3,1,'任务执行成功'),
	(597,'2021-08-25 16:46:00.000000','2021-08-25 16:46:00.000000',3,1,'任务执行成功'),
	(598,'2021-08-25 16:48:00.000000','2021-08-25 16:48:00.000000',3,1,'任务执行成功'),
	(599,'2021-08-25 16:50:00.000000','2021-08-25 16:50:00.000000',3,1,'任务执行成功'),
	(600,'2021-08-25 16:52:00.000000','2021-08-25 16:52:00.000000',3,1,'任务执行成功'),
	(601,'2021-08-25 16:54:00.000000','2021-08-25 16:54:00.000000',3,1,'任务执行成功'),
	(602,'2021-08-25 16:56:00.000000','2021-08-25 16:56:00.000000',3,1,'任务执行成功'),
	(603,'2021-08-25 16:58:00.000000','2021-08-25 16:58:00.000000',3,1,'任务执行成功'),
	(604,'2021-08-25 17:00:00.000000','2021-08-25 17:00:00.000000',3,1,'任务执行成功'),
	(605,'2021-08-25 17:02:01.000000','2021-08-25 17:02:01.000000',3,1,'任务执行成功'),
	(606,'2021-08-25 17:04:00.000000','2021-08-25 17:04:00.000000',3,1,'任务执行成功'),
	(607,'2021-08-25 17:06:00.000000','2021-08-25 17:06:00.000000',3,1,'任务执行成功'),
	(608,'2021-08-25 17:08:00.000000','2021-08-25 17:08:00.000000',3,1,'任务执行成功'),
	(609,'2021-08-25 17:10:00.000000','2021-08-25 17:10:00.000000',3,1,'任务执行成功'),
	(610,'2021-08-25 17:12:00.000000','2021-08-25 17:12:00.000000',3,1,'任务执行成功'),
	(611,'2021-08-25 17:14:00.000000','2021-08-25 17:14:00.000000',3,1,'任务执行成功'),
	(612,'2021-08-25 17:16:00.000000','2021-08-25 17:16:00.000000',3,1,'任务执行成功'),
	(613,'2021-08-25 17:18:00.000000','2021-08-25 17:18:00.000000',3,1,'任务执行成功'),
	(614,'2021-08-25 17:20:00.000000','2021-08-25 17:20:00.000000',3,1,'任务执行成功'),
	(615,'2021-08-25 17:22:00.000000','2021-08-25 17:22:00.000000',3,1,'任务执行成功'),
	(616,'2021-08-25 17:24:00.000000','2021-08-25 17:24:00.000000',3,1,'任务执行成功'),
	(617,'2021-08-25 17:26:00.000000','2021-08-25 17:26:00.000000',3,1,'任务执行成功'),
	(618,'2021-08-25 17:28:00.000000','2021-08-25 17:28:00.000000',3,1,'任务执行成功'),
	(619,'2021-08-25 17:30:00.000000','2021-08-25 17:30:00.000000',3,1,'任务执行成功'),
	(620,'2021-08-25 17:32:00.000000','2021-08-25 17:32:00.000000',3,1,'任务执行成功'),
	(621,'2021-08-25 17:34:00.000000','2021-08-25 17:34:00.000000',3,1,'任务执行成功'),
	(622,'2021-08-25 17:36:00.000000','2021-08-25 17:36:00.000000',3,1,'任务执行成功'),
	(623,'2021-08-25 17:38:00.000000','2021-08-25 17:38:00.000000',3,1,'任务执行成功'),
	(624,'2021-08-25 17:40:00.000000','2021-08-25 17:40:00.000000',3,1,'任务执行成功'),
	(625,'2021-08-25 17:42:00.000000','2021-08-25 17:42:00.000000',3,1,'任务执行成功'),
	(626,'2021-08-25 17:44:00.000000','2021-08-25 17:44:00.000000',3,1,'任务执行成功'),
	(627,'2021-08-25 17:46:00.000000','2021-08-25 17:46:00.000000',3,1,'任务执行成功'),
	(628,'2021-08-25 17:48:00.000000','2021-08-25 17:48:00.000000',3,1,'任务执行成功'),
	(629,'2021-08-25 17:50:00.000000','2021-08-25 17:50:00.000000',3,1,'任务执行成功'),
	(630,'2021-08-25 17:52:00.000000','2021-08-25 17:52:00.000000',3,1,'任务执行成功'),
	(631,'2021-08-25 17:54:00.000000','2021-08-25 17:54:00.000000',3,1,'任务执行成功'),
	(632,'2021-08-25 17:56:00.000000','2021-08-25 17:56:00.000000',3,1,'任务执行成功'),
	(633,'2021-08-25 17:58:00.000000','2021-08-25 17:58:00.000000',3,1,'任务执行成功'),
	(634,'2021-08-25 18:00:00.000000','2021-08-25 18:00:00.000000',3,1,'任务执行成功'),
	(635,'2021-08-25 18:02:00.000000','2021-08-25 18:02:00.000000',3,1,'任务执行成功'),
	(636,'2021-08-25 18:04:00.000000','2021-08-25 18:04:00.000000',3,1,'任务执行成功'),
	(637,'2021-08-25 18:06:00.000000','2021-08-25 18:06:00.000000',3,1,'任务执行成功'),
	(638,'2021-08-25 18:08:00.000000','2021-08-25 18:08:00.000000',3,1,'任务执行成功'),
	(639,'2021-08-25 18:10:00.000000','2021-08-25 18:10:00.000000',3,1,'任务执行成功'),
	(640,'2021-08-25 18:12:00.000000','2021-08-25 18:12:00.000000',3,1,'任务执行成功'),
	(641,'2021-08-25 18:14:00.000000','2021-08-25 18:14:00.000000',3,1,'任务执行成功'),
	(642,'2021-08-25 18:16:00.000000','2021-08-25 18:16:00.000000',3,1,'任务执行成功'),
	(643,'2021-08-25 18:18:00.000000','2021-08-25 18:18:00.000000',3,1,'任务执行成功'),
	(644,'2021-08-25 18:20:00.000000','2021-08-25 18:20:00.000000',3,1,'任务执行成功'),
	(645,'2021-08-25 18:22:00.000000','2021-08-25 18:22:00.000000',3,1,'任务执行成功'),
	(646,'2021-08-25 18:24:00.000000','2021-08-25 18:24:00.000000',3,1,'任务执行成功'),
	(647,'2021-08-25 18:26:00.000000','2021-08-25 18:26:00.000000',3,1,'任务执行成功'),
	(648,'2021-08-25 18:28:00.000000','2021-08-25 18:28:00.000000',3,1,'任务执行成功'),
	(649,'2021-08-25 18:30:00.000000','2021-08-25 18:30:00.000000',3,1,'任务执行成功'),
	(650,'2021-08-25 18:32:00.000000','2021-08-25 18:32:00.000000',3,1,'任务执行成功'),
	(651,'2021-08-25 18:34:00.000000','2021-08-25 18:34:00.000000',3,1,'任务执行成功'),
	(652,'2021-08-25 18:36:00.000000','2021-08-25 18:36:00.000000',3,1,'任务执行成功'),
	(653,'2021-08-25 18:38:00.000000','2021-08-25 18:38:00.000000',3,1,'任务执行成功'),
	(654,'2021-08-25 18:40:01.000000','2021-08-25 18:40:01.000000',3,1,'任务执行成功'),
	(655,'2021-08-25 18:42:00.000000','2021-08-25 18:42:00.000000',3,1,'任务执行成功'),
	(656,'2021-08-25 18:44:01.000000','2021-08-25 18:44:01.000000',3,1,'任务执行成功'),
	(657,'2021-08-25 18:46:01.000000','2021-08-25 18:46:01.000000',3,1,'任务执行成功'),
	(658,'2021-08-25 18:48:01.000000','2021-08-25 18:48:01.000000',3,1,'任务执行成功'),
	(659,'2021-08-25 18:50:00.000000','2021-08-25 18:50:00.000000',3,1,'任务执行成功'),
	(660,'2021-08-25 18:52:01.000000','2021-08-25 18:52:01.000000',3,1,'任务执行成功'),
	(661,'2021-08-25 18:54:00.000000','2021-08-25 18:54:00.000000',3,1,'任务执行成功'),
	(662,'2021-08-25 18:56:01.000000','2021-08-25 18:56:01.000000',3,1,'任务执行成功'),
	(663,'2021-08-25 18:58:01.000000','2021-08-25 18:58:01.000000',3,1,'任务执行成功'),
	(664,'2021-08-25 19:00:01.000000','2021-08-25 19:00:01.000000',3,1,'任务执行成功'),
	(665,'2021-08-25 19:02:01.000000','2021-08-25 19:02:01.000000',3,1,'任务执行成功'),
	(666,'2021-08-25 19:04:01.000000','2021-08-25 19:04:01.000000',3,1,'任务执行成功'),
	(667,'2021-08-25 19:06:01.000000','2021-08-25 19:06:01.000000',3,1,'任务执行成功'),
	(668,'2021-08-25 19:08:01.000000','2021-08-25 19:08:01.000000',3,1,'任务执行成功'),
	(669,'2021-08-25 19:10:01.000000','2021-08-25 19:10:01.000000',3,1,'任务执行成功'),
	(670,'2021-08-25 19:12:01.000000','2021-08-25 19:12:01.000000',3,1,'任务执行成功'),
	(671,'2021-08-25 19:14:01.000000','2021-08-25 19:14:01.000000',3,1,'任务执行成功'),
	(672,'2021-08-25 19:16:01.000000','2021-08-25 19:16:01.000000',3,1,'任务执行成功'),
	(673,'2021-08-25 19:18:01.000000','2021-08-25 19:18:01.000000',3,1,'任务执行成功'),
	(674,'2021-08-25 19:20:01.000000','2021-08-25 19:20:01.000000',3,1,'任务执行成功'),
	(675,'2021-08-25 19:22:01.000000','2021-08-25 19:22:01.000000',3,1,'任务执行成功'),
	(676,'2021-08-25 19:24:01.000000','2021-08-25 19:24:01.000000',3,1,'任务执行成功'),
	(677,'2021-08-25 19:26:01.000000','2021-08-25 19:26:01.000000',3,1,'任务执行成功'),
	(678,'2021-08-25 19:28:01.000000','2021-08-25 19:28:01.000000',3,1,'任务执行成功'),
	(679,'2021-08-25 19:30:01.000000','2021-08-25 19:30:01.000000',3,1,'任务执行成功'),
	(680,'2021-08-25 19:32:00.000000','2021-08-25 19:32:00.000000',3,1,'任务执行成功'),
	(681,'2021-08-25 19:34:00.000000','2021-08-25 19:34:00.000000',3,1,'任务执行成功'),
	(682,'2021-08-25 19:36:00.000000','2021-08-25 19:36:00.000000',3,1,'任务执行成功'),
	(683,'2021-08-25 19:36:01.000000','2021-08-25 19:36:01.000000',3,1,'任务执行成功'),
	(684,'2021-08-25 19:38:01.000000','2021-08-25 19:38:01.000000',3,1,'任务执行成功'),
	(685,'2021-08-25 19:40:01.000000','2021-08-25 19:40:01.000000',3,1,'任务执行成功'),
	(686,'2021-08-25 19:42:00.000000','2021-08-25 19:42:00.000000',3,1,'任务执行成功'),
	(687,'2021-08-25 19:44:00.000000','2021-08-25 19:44:00.000000',3,1,'任务执行成功'),
	(688,'2021-08-25 19:44:01.000000','2021-08-25 19:44:01.000000',3,1,'任务执行成功'),
	(689,'2021-08-25 19:46:00.000000','2021-08-25 19:46:00.000000',3,1,'任务执行成功'),
	(690,'2021-08-25 19:48:00.000000','2021-08-25 19:48:00.000000',3,1,'任务执行成功'),
	(691,'2021-08-25 19:50:00.000000','2021-08-25 19:50:00.000000',3,1,'任务执行成功'),
	(692,'2021-08-25 19:52:00.000000','2021-08-25 19:52:00.000000',3,1,'任务执行成功'),
	(693,'2021-08-25 19:54:00.000000','2021-08-25 19:54:00.000000',3,1,'任务执行成功'),
	(694,'2021-08-25 19:56:00.000000','2021-08-25 19:56:00.000000',3,1,'任务执行成功'),
	(695,'2021-08-25 19:58:00.000000','2021-08-25 19:58:00.000000',3,1,'任务执行成功'),
	(696,'2021-08-25 20:00:00.000000','2021-08-25 20:00:00.000000',3,1,'任务执行成功'),
	(697,'2021-08-25 20:02:00.000000','2021-08-25 20:02:00.000000',3,1,'任务执行成功'),
	(698,'2021-08-25 20:04:00.000000','2021-08-25 20:04:00.000000',3,1,'任务执行成功'),
	(699,'2021-08-25 20:06:00.000000','2021-08-25 20:06:00.000000',3,1,'任务执行成功'),
	(700,'2021-08-25 20:08:00.000000','2021-08-25 20:08:00.000000',3,1,'任务执行成功'),
	(701,'2021-08-25 20:10:00.000000','2021-08-25 20:10:00.000000',3,1,'任务执行成功'),
	(702,'2021-08-25 20:12:00.000000','2021-08-25 20:12:00.000000',3,1,'任务执行成功'),
	(703,'2021-08-25 20:14:00.000000','2021-08-25 20:14:00.000000',3,1,'任务执行成功'),
	(704,'2021-08-25 20:16:00.000000','2021-08-25 20:16:00.000000',3,1,'任务执行成功'),
	(705,'2021-08-25 20:18:00.000000','2021-08-25 20:18:00.000000',3,1,'任务执行成功'),
	(706,'2021-08-25 20:20:00.000000','2021-08-25 20:20:00.000000',3,1,'任务执行成功'),
	(707,'2021-08-25 20:22:00.000000','2021-08-25 20:22:00.000000',3,1,'任务执行成功'),
	(708,'2021-08-25 20:24:00.000000','2021-08-25 20:24:00.000000',3,1,'任务执行成功'),
	(709,'2021-08-25 20:26:00.000000','2021-08-25 20:26:00.000000',3,1,'任务执行成功'),
	(710,'2021-08-25 20:28:00.000000','2021-08-25 20:28:00.000000',3,1,'任务执行成功'),
	(711,'2021-08-25 20:30:00.000000','2021-08-25 20:30:00.000000',3,1,'任务执行成功'),
	(712,'2021-08-25 20:32:01.000000','2021-08-25 20:32:01.000000',3,1,'任务执行成功'),
	(713,'2021-08-25 20:34:00.000000','2021-08-25 20:34:00.000000',3,1,'任务执行成功'),
	(714,'2021-08-25 20:36:00.000000','2021-08-25 20:36:00.000000',3,1,'任务执行成功'),
	(715,'2021-08-25 20:38:00.000000','2021-08-25 20:38:00.000000',3,1,'任务执行成功'),
	(716,'2021-08-25 20:40:00.000000','2021-08-25 20:40:00.000000',3,1,'任务执行成功'),
	(717,'2021-08-25 20:42:00.000000','2021-08-25 20:42:00.000000',3,1,'任务执行成功'),
	(718,'2021-08-25 20:44:00.000000','2021-08-25 20:44:00.000000',3,1,'任务执行成功'),
	(719,'2021-08-25 20:46:00.000000','2021-08-25 20:46:00.000000',3,1,'任务执行成功'),
	(720,'2021-08-25 20:48:00.000000','2021-08-25 20:48:00.000000',3,1,'任务执行成功'),
	(721,'2021-08-25 20:50:00.000000','2021-08-25 20:50:00.000000',3,1,'任务执行成功'),
	(722,'2021-08-25 20:52:00.000000','2021-08-25 20:52:00.000000',3,1,'任务执行成功'),
	(723,'2021-08-25 20:54:00.000000','2021-08-25 20:54:00.000000',3,1,'任务执行成功'),
	(724,'2021-08-25 20:56:00.000000','2021-08-25 20:56:00.000000',3,1,'任务执行成功'),
	(725,'2021-08-25 20:58:00.000000','2021-08-25 20:58:00.000000',3,1,'任务执行成功'),
	(726,'2021-08-25 21:00:00.000000','2021-08-25 21:00:00.000000',3,1,'任务执行成功'),
	(727,'2021-08-25 21:02:00.000000','2021-08-25 21:02:00.000000',3,1,'任务执行成功'),
	(728,'2021-08-25 21:04:00.000000','2021-08-25 21:04:00.000000',3,1,'任务执行成功'),
	(729,'2021-08-25 21:06:00.000000','2021-08-25 21:06:00.000000',3,1,'任务执行成功'),
	(730,'2021-08-25 21:08:00.000000','2021-08-25 21:08:00.000000',3,1,'任务执行成功'),
	(731,'2021-08-25 21:10:00.000000','2021-08-25 21:10:00.000000',3,1,'任务执行成功'),
	(732,'2021-08-25 21:12:02.000000','2021-08-25 21:12:02.000000',3,1,'任务执行成功'),
	(733,'2021-08-25 21:14:00.000000','2021-08-25 21:14:00.000000',3,1,'任务执行成功'),
	(734,'2021-08-25 21:16:00.000000','2021-08-25 21:16:00.000000',3,1,'任务执行成功'),
	(735,'2021-08-25 21:18:00.000000','2021-08-25 21:18:00.000000',3,1,'任务执行成功'),
	(736,'2021-08-25 21:20:00.000000','2021-08-25 21:20:00.000000',3,1,'任务执行成功'),
	(737,'2021-08-25 21:26:01.000000','2021-08-25 21:26:01.000000',3,1,'任务执行成功');

/*!40000 ALTER TABLE `task_log` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table test
# ------------------------------------------------------------

DROP TABLE IF EXISTS `test`;

CREATE TABLE `test` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `createTime` datetime NOT NULL COMMENT '创建时间',
  `updateTime` datetime NOT NULL COMMENT '更新时间',
  `activeCode` varchar(200) NOT NULL DEFAULT '' COMMENT '激活码',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态#1:开启,2:关闭#IsSearchParams,IsDictColumn',
  `computerMac` text COMMENT '绑定的电脑mac码',
  `limit` int(11) NOT NULL DEFAULT '1' COMMENT '允许绑定的设备数量##IsNumberColumn',
  `isBind` tinyint(4) DEFAULT '2' COMMENT '是否绑定设备#1:绑定,未绑定#IsSearchParams,IsDictColumn',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='生成代码的权限控制';



# Dump of table test_a
# ------------------------------------------------------------

DROP TABLE IF EXISTS `test_a`;

CREATE TABLE `test_a` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `createTime` datetime NOT NULL COMMENT '创建时间',
  `updateTime` datetime NOT NULL COMMENT '更新时间',
  `activeCode` varchar(200) NOT NULL DEFAULT '' COMMENT '激活码',
  `status` tinyint(4) NOT NULL DEFAULT '1' COMMENT '状态#1:开启,2:关闭#IsSearchParams,IsDictColumn',
  `computerMac` text COMMENT '绑定的电脑mac码',
  `limit` int(11) NOT NULL DEFAULT '1' COMMENT '允许绑定的设备数量##IsNumberColumn',
  `isBind` tinyint(4) DEFAULT '2' COMMENT '是否绑定设备#1:绑定,未绑定#IsSearchParams,IsDictColumn',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='生成代码的权限控制';




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
