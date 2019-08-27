DROP TABLE IF EXISTS `salary`;

CREATE TABLE `salary` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `salary` int(10) unsigned DEFAULT NULL,
  `people` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_salary` (`salary`,`people`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

INSERT INTO `salary` (id, salary, people) VALUES (1, 100, "张三");
INSERT INTO `salary` (id, salary, people) VALUES (2, 200, "李四");
INSERT INTO `salary` (id, salary, people) VALUES (3, 300, "王五");
INSERT INTO `salary` (id, salary, people) VALUES (4, 400, "赵六");