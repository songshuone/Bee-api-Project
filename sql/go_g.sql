/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 50717
Source Host           : localhost:3306
Source Database       : go_g

Target Server Type    : MYSQL
Target Server Version : 50717
File Encoding         : 65001

Date: 2017-08-29 11:27:23
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `banner`
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `image_url` varchar(60) DEFAULT NULL,
  `adver_url` varchar(60) DEFAULT NULL,
  `banner_desc` varchar(10) DEFAULT NULL,
  `create_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of banner
-- ----------------------------
INSERT INTO `banner` VALUES ('1', 'fdfds', 'fdfds', 'fdfsfs', '2017-08-26 13:21:34', '2017-09-08 13:21:38');
INSERT INTO `banner` VALUES ('2', 'fdfd', 'ffdfd', 'fdfd', '2017-08-26 13:21:51', '2017-09-10 13:21:54');

-- ----------------------------
-- Table structure for `post`
-- ----------------------------
DROP TABLE IF EXISTS `post`;
CREATE TABLE `post` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `title` varchar(255) CHARACTER SET utf8 NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of post
-- ----------------------------
INSERT INTO `post` VALUES ('1', 'java是一门跨平台的语言');
INSERT INTO `post` VALUES ('2', '上一篇介绍Go反射的时候，提到了如何通过反射获取Struct的Tag，这一篇文章主要就是介绍这个的使用和原理，在介绍之前我们先看一下JSON字符串和Struct类型相互转换的例子。');

-- ----------------------------
-- Table structure for `post_tag_rel`
-- ----------------------------
DROP TABLE IF EXISTS `post_tag_rel`;
CREATE TABLE `post_tag_rel` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `post_id` int(11) NOT NULL,
  `tag_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of post_tag_rel
-- ----------------------------
INSERT INTO `post_tag_rel` VALUES ('1', '1', '1');
INSERT INTO `post_tag_rel` VALUES ('2', '2', '2');
INSERT INTO `post_tag_rel` VALUES ('3', '2', '1');
INSERT INTO `post_tag_rel` VALUES ('4', '1', '2');

-- ----------------------------
-- Table structure for `tag`
-- ----------------------------
DROP TABLE IF EXISTS `tag`;
CREATE TABLE `tag` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of tag
-- ----------------------------
INSERT INTO `tag` VALUES ('1', 'java');
INSERT INTO `tag` VALUES ('2', 'go');
INSERT INTO `tag` VALUES ('3', 'python');
INSERT INTO `tag` VALUES ('4', null);

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(10) NOT NULL DEFAULT '',
  `password` varchar(255) NOT NULL DEFAULT '',
  `address` varchar(20) DEFAULT NULL,
  `age` int(11) DEFAULT NULL,
  `email` varchar(20) DEFAULT NULL,
  `birthday` varchar(20) DEFAULT NULL,
  `post_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'wp123', '88d31af5cda7dc0962b7d1d9c64184b1', '', '0', '', '', '1');
INSERT INTO `user` VALUES ('2', '1231231', '90d47a191e15cf5ab3f576707f61b57b', '', '0', '', '', '2');
