drop database if exists `cugoj`;
create database `cugoj`;
USE `cugoj`;
CREATE TABLE `user` (
    username VARCHAR(48) NOT NULL UNIQUE,                # 用户名(唯一)
    password VARCHAR(32) NOT NULL,                       # 密码
    email VARCHAR(100) NOT NULL,                         # 邮箱
    black_list boolean NOT NULL,                         # 是否是黑名单
    admin TINYINT NOT NULL,                              # 权限(0:普通用户, 1:助理, 2:教练)
    register_time DATETIME NOT NULL,                     # 注册时间
    privilege_end_time DATETIME                          # 权限结束时间
);

CREATE TABLE `new` (
    id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,    # 新闻ID(唯一)
    post_time DATETIME NOT NULL,                                # 最后更新时间
    title VARCHAR(255) NOT NULL,                                # 标题
    author VARCHAR(48) NOT NULL,                                # 作者
    content MEDIUMTEXT NOT NULL,                                # 内容
    importance  BOOLEAN NOT NULL,                               # 是否置顶
    foreign key(author) references user(username)          # 外键
);