/*
 Navicat Premium Data Transfer

 Source Server         : coco数据库
 Source Server Type    : MySQL
 Source Server Version : 50736
 Source Host           :
 Source Schema         : mini-min-tiktok

 Target Server Type    : MySQL
 Target Server Version : 50736
 File Encoding         : 65001

 Date: 28/01/2023 21:53:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for t_comment
-- ----------------------------
DROP TABLE IF EXISTS `t_comment`;
CREATE TABLE `t_comment`  (
                              `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '评论id',
                              `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户id',
                              `content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '评论内容',
                              `video_id` int(10) UNSIGNED NOT NULL COMMENT '视频id',
                              `create_date` datetime NOT NULL COMMENT '评论发布日期，格式为mm-dd',
                              `update_date` datetime COMMENT '评论更新日期，格式为mm-dd',
                              `delete_date` datetime COMMENT '评论删除日期，格式为mm-dd',
                              PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_favorite
-- ----------------------------
DROP TABLE IF EXISTS `t_favorite`;
CREATE TABLE `t_favorite`  (
                               `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '点赞id',
                               `video_id` int(10) UNSIGNED NOT NULL COMMENT '视频id',
                               `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户id',
                               `status` tinyint(1) NOT NULL DEFAULT 0 COMMENT '点赞状态(0为未点赞, 1为已点赞)',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_follow
-- ----------------------------
DROP TABLE IF EXISTS `t_follow`;
CREATE TABLE `t_follow`  (
                             `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
                             `user_id` int(11) NOT NULL COMMENT '用户id',
                             `follower_id` int(11) NOT NULL COMMENT '关注者id',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_friend
-- ----------------------------
DROP TABLE IF EXISTS `t_friend`;
CREATE TABLE `t_friend`  (
                             `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
                             `user_id` int(10) UNSIGNED NOT NULL COMMENT '用户id',
                             `friend_id` int(10) UNSIGNED NOT NULL COMMENT '好友id',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_user
-- ----------------------------
DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user`  (
                           `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键id',
                           `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
                           `follow_count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '关注数',
                           `follower_count` int(11) UNSIGNED NOT NULL DEFAULT 0 COMMENT '粉丝数',
                           `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
                           `update_date` datetime COMMENT '用户更新日期，格式为mm-dd',
                           `delete_date` datetime COMMENT '用户删除日期，格式为mm-dd',
                           PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 6 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for t_video
-- ----------------------------
DROP TABLE IF EXISTS `t_video`;
CREATE TABLE `t_video`  (
                            `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '视频id',
                            `author_id` int(10) UNSIGNED NOT NULL COMMENT '作者id',
                            `play_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频链接',
                            `cover_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频封面链接',
                            `favorite_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '点赞数',
                            `comment_count` int(10) UNSIGNED NOT NULL DEFAULT 0 COMMENT '评论数',
                            `is_favorite` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否已点赞(0为未点赞, 1为已点赞)',
                            `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '视频标题',
                            `create_date` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '视频上传时间',
                            `delete_date` datetime COMMENT '视频删除日期，格式为mm-dd',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
