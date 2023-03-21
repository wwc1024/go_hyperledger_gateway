-- phpMyAdmin SQL Dump
-- version 4.7.0
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: 2020-04-27 13:59:47
-- 服务器版本： 5.6.25-log
-- PHP Version: 7.0.18

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `blockchain_gateway`
--

-- --------------------------------------------------------

--
-- 表的结构 `gateway_admin`
--

CREATE TABLE `gateway_admin` (
  `id` bigint(20) NOT NULL COMMENT '自增id',
  `user_name` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `salt` varchar(50) NOT NULL DEFAULT '' COMMENT '盐',
  `password` varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
  `create_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '新增时间',
  `update_at` datetime NOT NULL DEFAULT '1971-01-01 00:00:00' COMMENT '更新时间',
  `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='管理员表';

--
-- 转存表中的数据 `gateway_admin`
--

INSERT INTO `gateway_admin` (`id`, `user_name`, `salt`, `password`, `create_at`, `update_at`, `is_delete`) VALUES
(1, 'admin', 'admin', '2823d896e9822c0833d41d4904f0c00756d718570fce49b9a379a62c804689d3', '2020-04-10 16:42:05', '2020-04-21 06:35:08', 0);

-- --------------------------------------------------------

--
-- 表的结构 `gateway_peerlist`
--

CREATE TABLE `gateway_peerlist` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '自增id',
  `name` varchar(255) NOT NULL DEFAULT '' COMMENT '节点名称',
  `create_at` datetime NOT NULL COMMENT '创建时间',
  `update_at` datetime NOT NULL COMMENT '更新时间',
  `org` VARCHAR(255) NOT NULL COMMENT '所属组织',
  `ip` varchar(255) NOT NULL DEFAULT '' COMMENT '所属ip',
  `state` varchar(255) NOT NULL DEFAULT '' COMMENT '状态',
  `port` bigint(20) NOT NULL DEFAULT '0' COMMENT '端口',
  `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 1=删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='节点管理表';

--
-- 转存表中的数据 `gateway_peerlist`
--

INSERT INTO `gateway_peerlist` (`id`, `peer`, `channelID`, `chaincode`,`detail`, `peerNum`) VALUES
('0-00', 'peer_test0','channeltest0', '2020-04-21 07:23:34', 'org1', '192.127.0.1', 'good', ':8080', 0),
('0-01', 'peer_test1','2019-03-11 21:35:01', '2021-02-22 03:23:54', 'org1', '192.127.0.2', 'good', ':8081', 0);

-- --------------------------------------------------------

--
-- 表的结构 `gateway_channellist`
--

CREATE TABLE `gateway_channellist` (
   `id` varchar(255) NOT NULL COMMENT 'id',
  `peer` varchar(255) NOT NULL  COMMENT '节点名称',
  `channel_id` varchar(255) NOT NULL COMMENT '通道名称',
  `chaincode` varchar(255) NOT NULL COMMENT '链码',
  `detail` varchar(255) NOT NULL COMMENT '细琐',
  `peerNum` varchar(255) NOT NULL  COMMENT '通道节点数'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通道管理表';


--
-- 转存表中的数据 `gateway_channellist`
--

INSERT INTO `gateway_channellist` (`id`, `peer`, `channel_id`, `chaincode`,`detail`, `peerNum`) VALUES
('0-00', 'peer_test0','channeltest0', 'chaincode test0', '测试', '0');


-- --------------------------------------------------------

--
-- 表的结构 `gateway_contentlist`
--

CREATE TABLE `gateway_contentlist` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '自增id',
  `content_name` varchar(255) NOT NULL DEFAULT '' COMMENT '合约名称',
  `tap` varchar(255) NOT NULL DEFAULT '' COMMENT '版本号',
  `service_type` varchar(255) NOT NULL DEFAULT '' COMMENT '服务类型',
  `service_name` varchar(255) NOT NULL DEFAULT '' COMMENT '服务名称',
  `detail` varchar(255) NOT NULL DEFAULT '' COMMENT '通道描述'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='通道合约表';

--
-- 转存表中的数据 `gateway_contentlist`
--

INSERT INTO `gateway_contentlist` (`id`, `content_name`, `tap`, `service_type`,`service_name`, `detail`) VALUES
(1, 'channel_content0','v1','传感器类','自动采集合约','该合约实现了温湿度传感器的自动采集服务');


CREATE TABLE `gateway_servicelist` (
  `id` bigint(20) UNSIGNED NOT NULL COMMENT '自增id',
  `content_name` varchar(255) NOT NULL DEFAULT '' COMMENT '合约名称',
  `tag` varchar(255) NOT NULL COMMENT '版本号',
  `service_type` varchar(255) NOT NULL COMMENT '服务类型',
  `service_name` VARCHAR(255) NOT NULL COMMENT '服务名称',
  `detail` varchar(255) NOT NULL DEFAULT '' COMMENT '服务简介',
  `port` bigint(20) NOT NULL DEFAULT '0' COMMENT '端口',
  `is_delete` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否删除 1=删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='合约服务管理表';

--
-- 转存表中的数据 `gateway_peerlist`
--

INSERT INTO `gateway_servicelist` (`id`, `content_name`, `tag`, `service_type`,`service_name`, `detail`, `port`, `is_delete`) VALUES
(1, 'service_test0','0', '传感类', '温度采样', '该服务实现了温度传感器设定采样频率下的数据上链采集', '8080', 0),
(2, 'service_test1','0', '传感类', '湿度采样', '该服务实现了湿度传感器设定采样频率下的数据上链采集', '8080', 0),
(3, 'service_test0','0', '其他', '其他', '其他', '8082', 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `gateway_admin`
--
ALTER TABLE `gateway_admin`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `gateway_peerlist`
--
ALTER TABLE `gateway_peerlist`
  ADD PRIMARY KEY (`id`);


ALTER TABLE `gateway_channellist`
  ADD PRIMARY KEY (`id`);

ALTER TABLE `gateway_servicelist`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `gateway_service_access_control`
--
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
