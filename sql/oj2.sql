drop database if exists `cugoj`;
create database `cugoj`;
USE `cugoj`;
CREATE TABLE `user` (
    user_id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY # 用户ID(唯一)
);
CREATE TABLE `new` (
    ID INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,    # 新闻ID(唯一)
    author VARCHAR(48) NOT NULL,                              # 作者
    updateTime DATETIME NOT NULL,                                # 更新时间
    title VARCHAR(255) NOT NULL,                                # 标题
    content MEDIUMTEXT NOT NULL,                                # 内容
    importance  BOOLEAN NOT NULL,                           # 是否置顶
    foreign key(author) references user(user_id)            # 外键
);
