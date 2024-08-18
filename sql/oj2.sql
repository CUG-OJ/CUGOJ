drop database if exists `cugoj`;
create database `cugoj`;
USE `cugoj`;
CREATE TABLE `user` (
    user_id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, # 用户ID(唯一)
    privilege TINYINT NOT NULL,                          # 权限(0:普通用户, 1:助理, 2:教练)
    privilege_end_time DATETIME NOT NULL,                # 权限结束时间
    password VARCHAR(32) NOT NULL,                       # 密码
    email VARCHAR(100) NOT NULL,                         # 邮箱
    register_time DATETIME NOT NULL,                     # 注册时间
    captain_name VARCHAR(48),                            # 队长ID
    school VARCHAR(32),                                  # 学校
    faculty VARCHAR(32),                                 # 院系
    valid boolean NOT NULL,                              # 是否有效(封禁)
    foreign key (captain_name) references user(username) # 外键
);
create table `team` (
    captain_name VARCHAR(48),                                   # 队长ID(唯一)
    team_name VARCHAR(48) NOT NULL,                             # 队伍名
    member1_name VARCHAR(48),                                     # 队员2
    member2_name VARCHAR(48),                                     # 队员3
    foreign key(captain_name) references user(username),           # 外键
    foreign key (member1_name) references user(username),          # 外键
    foreign key (member2_name) references user(username)           # 外键
);
create table `class`(
    class_id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,   # 班级ID(唯一)
    class_name VARCHAR(48) NOT NULL,                            # 班级名
    teacher_name VARCHAR(48) NOT NULL,                            # 教师(群主)ID
    register_time DATETIME NOT NULL,                            # 注册时间
    foreign key (teacher_name) references user(username)          # 外键
);
CREATE TABLE `new` (
    news_id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,    # 新闻ID(唯一)
    writer_name VARCHAR(48) NOT NULL,                              # 发布者ID
    post_time DATETIME NOT NULL,                                # 发布时间
    title VARCHAR(255) NOT NULL,                                # 标题
    content MEDIUMTEXT NOT NULL,                                # 内容
    importance  TINYINT UNSIGNED UNIQUE NOT NULL,               # 重要性[0, 255]
    foreign key(writer_name) references user(username)            # 外键
);
CREATE TABLE `problem`(
    problem_id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, # 题目ID(唯一)
    title VARCHAR(255) NOT NULL,                                # 题目标题
    username VARCHAR(48) NOT NULL,                              # 发布者ID
    time_limit INTEGER(11) NOT NULL,                            # 时间限制
    memory_limit INTEGER(11) NOT NULL,                          # 内存限制
    description TEXT NOT NULL,                                  # 题目描述
    input TEXT NOT NULL,                                        # 输入描述
    output TEXT NOT NULL,                                       # 输出描述
    sample_input TEXT NOT NULL,                                 # 样例输入
    sample_output TEXT NOT NULL,                                # 样例输出
    hint TEXT,                                                  # 提示
    source VARCHAR(255),                                        # 来源
    foreign key(username) references user(username)             # 外键
);
CREATE TABLE `solution`(
    solution_id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,# 提交ID(唯一)
    problem_id INTEGER(11) NOT NULL,                            # 题目ID
    user_id INTEGER(11) NOT NULL,                               # 提交者ID
    time DATETIME NOT NULL,                                     # 提交时间
    language INTEGER(11) NOT NULL,                              # 编程语言
    code TEXT NOT NULL,                                         # 代码
    result INTEGER(11) NOT NULL,                                # 结果
    time_used INTEGER(11),                                      # 运行时长(ms)
    memory_used INTEGER(11),                                    # 内存
    foreign key (problem_id) references problem(problem_id)     # 外键
);
CREATE TABLE `similarity`(
    similarity_id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,  # 相似度ID(唯一)
    solution_id INTEGER(11) NOT NULL,                               # 提交ID
    similarity FLOAT(11) NOT NULL,                                  # 相似度
    foreign key (solution_id) references solution(solution_id)      # 外键
);
CREATE TABLE `contest`(
    contest_id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, # 比赛ID(唯一)
    title VARCHAR(255) NOT NULL,                                # 竞赛标题
    host_id INTEGER(11) NOT NULL,                               # 发布者ID
    start_time DATETIME NOT NULL,                               # 开始时间
    end_time DATETIME NOT NULL,                                 # 结束时间
    password VARCHAR(32),                                       # 密码(可选)
    language_mask BIGINT(64) NOT NULL,                          # 编程语言(多选)
    description TEXT NOT NULL,                                  # 竞赛描述
    private boolean NOT NULL,                                   # 是否私有
    foreign key (host_id) references user(username)             # 外键
);
CREATE TABLE `contest_problem`(
    problem_id INTEGER(11) NOT NULL AUTO_INCREMENT PRIMARY KEY, # 题目ID(唯一)
    title VARCHAR(255) NOT NULL,                                # 题目标题
    username VARCHAR(48) NOT NULL,                              # 发布者ID
    time_limit INTEGER(11) NOT NULL,                            # 时间限制
    memory_limit INTEGER(11) NOT NULL,                          # 内存限制
    description TEXT NOT NULL,                                  # 题目描述
    input TEXT NOT NULL,                                        # 输入描述
    output TEXT NOT NULL,                                       # 输出描述
    sample_input TEXT NOT NULL,                                 # 样例输入
    sample_output TEXT NOT NULL,                                # 样例输出
    hint TEXT,                                                  # 提示
    source VARCHAR(255),                                        # 来源
    contest_id INTEGER(11) NOT NULL,                            # 比赛ID
    foreign key(username) references user(username),            # 外键
    foreign key(contest_id) references contest(contest_id)      # 外键
);
