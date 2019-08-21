/*
Navicat MySQL Data Transfer

Source Server         : [TeamServer]192.168.1.145
Source Server Version : 50722
Source Host           : 192.168.1.145:3306
Source Database       : test

Target Server Type    : MYSQL
Target Server Version : 50722
File Encoding         : 65001

Date: 2018-05-26 18:01:44
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `app_language`
-- ----------------------------
DROP TABLE IF EXISTS `app_language`;
CREATE TABLE `app_language` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键',
  `name` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '语言',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `status` int(11) DEFAULT NULL COMMENT '状态：0不可见，1可见',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of app_language
-- ----------------------------
INSERT INTO `app_language` VALUES ('1', '英文', null, '1');
INSERT INTO `app_language` VALUES ('3', '中文', null, '1');

-- ----------------------------
-- Table structure for `certification_info`
-- ----------------------------
DROP TABLE IF EXISTS `certification_info`;
CREATE TABLE `certification_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` text COMMENT '认证名称',
  `description` text COMMENT '认证描述信息',
  `image` text COMMENT '认证图像地址',
  `priority_sort` int(11) DEFAULT NULL COMMENT '排序优先级',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of certification_info
-- ----------------------------
INSERT INTO `certification_info` VALUES ('1', '安全认证1', '安全保障1', 'http://renzheng1.png', '0', '2018-05-26 16:49:14');
INSERT INTO `certification_info` VALUES ('2', '安全认证2', '安全保障2', 'http://renzheng2.png', '1', '2018-05-26 16:49:36');
INSERT INTO `certification_info` VALUES ('3', '安全认证3', '安全保障3', 'http://renzheng3.png', '2', '2018-05-26 16:50:00');
INSERT INTO `certification_info` VALUES ('4', '安全认证4', '安全保障4', 'http://renzheng4.png', '3', '2018-05-26 16:51:02');

-- ----------------------------
-- Table structure for `country`
-- ----------------------------
DROP TABLE IF EXISTS `country`;
CREATE TABLE `country` (
  `id` int(10) NOT NULL AUTO_INCREMENT COMMENT 'ID 对应用户 Country_code',
  `country_nb` text CHARACTER SET utf8mb4 NOT NULL COMMENT '国家码',
  `mark` text CHARACTER SET utf8mb4 NOT NULL COMMENT '代码',
  `ename` text CHARACTER SET utf8mb4 NOT NULL COMMENT '国家英文名',
  `cname` text CHARACTER SET utf8mb4 NOT NULL COMMENT '国家中文名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=218 DEFAULT CHARSET=latin1;

-- ----------------------------
-- Records of country
-- ----------------------------
INSERT INTO `country` VALUES ('2', '376', 'AD', 'Andorra', '安道尔');
INSERT INTO `country` VALUES ('3', '971', 'AE', 'United Arab Emirates', '阿拉伯联合酋长国');
INSERT INTO `country` VALUES ('4', '93', 'AF', 'Afghanistan', '阿富汗');
INSERT INTO `country` VALUES ('5', '1268', 'AG', 'Antigua and Barbuda', '安提瓜和巴布达');
INSERT INTO `country` VALUES ('6', '1264', 'AI', 'Anguilla', '安圭拉');
INSERT INTO `country` VALUES ('7', '355', 'AL', 'Albania', '阿尔巴尼亚');
INSERT INTO `country` VALUES ('8', '374', 'AM', 'Armenia', '亚美尼亚');
INSERT INTO `country` VALUES ('9', '244', 'AO', 'Angola', '安哥拉');
INSERT INTO `country` VALUES ('10', '54', 'AR', 'Argentina', '阿根廷');
INSERT INTO `country` VALUES ('11', '1684', 'AS', 'American Samoa', '美属萨摩亚');
INSERT INTO `country` VALUES ('12', '43', 'AT', 'Austria', '奥地利');
INSERT INTO `country` VALUES ('13', '61', 'AU', 'Australia', '澳大利亚');
INSERT INTO `country` VALUES ('14', '297', 'AW', 'Aruba', '阿鲁巴');
INSERT INTO `country` VALUES ('15', '994', 'AZ', 'Azerbaijan', '阿塞拜疆');
INSERT INTO `country` VALUES ('16', '387', 'BA', 'Bosniaand Herzegovina', '波斯尼亚和黑塞哥维那');
INSERT INTO `country` VALUES ('17', '1246', 'BB', 'Barbados', '巴巴多斯');
INSERT INTO `country` VALUES ('18', '880', 'BD', 'Bangladesh', '孟加拉国');
INSERT INTO `country` VALUES ('19', '32', 'BE', 'Belgium', '比利时');
INSERT INTO `country` VALUES ('20', '226', 'BF', 'Burkina Faso', '布基纳法索');
INSERT INTO `country` VALUES ('21', '359', 'BG', 'Bulgaria', '保加利亚');
INSERT INTO `country` VALUES ('22', '973', 'BH', 'Bahrain', '巴林');
INSERT INTO `country` VALUES ('23', '257', 'BI', 'Burundi', '布隆迪');
INSERT INTO `country` VALUES ('24', '229', 'BJ', 'Benin', '贝宁');
INSERT INTO `country` VALUES ('25', '1441', 'BM', 'Bermuda', '百慕大群岛');
INSERT INTO `country` VALUES ('26', '673', 'BN', 'Brunei', '文莱');
INSERT INTO `country` VALUES ('27', '591', 'BO', 'Bolivia', '玻利维亚');
INSERT INTO `country` VALUES ('28', '599', 'BQ', 'Caribisch Nederland', '荷兰加勒比');
INSERT INTO `country` VALUES ('29', '55', 'BR', 'Brazil', '巴西');
INSERT INTO `country` VALUES ('30', '1242', 'BS', 'Bahamas', '巴哈马');
INSERT INTO `country` VALUES ('31', '975', 'BT', 'Bhutan', '不丹');
INSERT INTO `country` VALUES ('32', '267', 'BW', 'Botswana', '博茨瓦纳');
INSERT INTO `country` VALUES ('33', '375', 'BY', 'Belarus', '白俄罗斯');
INSERT INTO `country` VALUES ('34', '501', 'BZ', 'Belize', '伯利兹');
INSERT INTO `country` VALUES ('35', '1', 'CA', 'Canada', '加拿大');
INSERT INTO `country` VALUES ('36', '243', 'CD', 'Democratic Republic of theCongo', '刚果民主共和国');
INSERT INTO `country` VALUES ('37', '236', 'CF', 'Central African Republic', '中非共和国');
INSERT INTO `country` VALUES ('38', '242', 'CG', 'Republic Of The Congo', '刚果共和国');
INSERT INTO `country` VALUES ('39', '41', 'CH', 'Switzerland', '瑞士');
INSERT INTO `country` VALUES ('40', '225', 'CI', 'Ivory Coast', '象牙海岸');
INSERT INTO `country` VALUES ('41', '682', 'CK', 'Cook Islands', '库克群岛');
INSERT INTO `country` VALUES ('42', '56', 'CL', 'Chile', '智利');
INSERT INTO `country` VALUES ('43', '237', 'CM', 'Cameroon', '喀麦隆');
INSERT INTO `country` VALUES ('44', '57', 'CO', 'Colombia', '哥伦比亚');
INSERT INTO `country` VALUES ('45', '506', 'CR', 'CostaRica', '哥斯达黎加');
INSERT INTO `country` VALUES ('46', '53', 'CU', 'Cuba', '古巴');
INSERT INTO `country` VALUES ('47', '238', 'CV', 'Cape Verde', '开普');
INSERT INTO `country` VALUES ('48', '599', 'CW', 'Curacao', '库拉索');
INSERT INTO `country` VALUES ('49', '357', 'CY', 'Cyprus', '塞浦路斯');
INSERT INTO `country` VALUES ('50', '420', 'CZ', 'Czech', '捷克');
INSERT INTO `country` VALUES ('51', '49', 'DE', 'Germany', '德国');
INSERT INTO `country` VALUES ('52', '253', 'DJ', 'Djibouti', '吉布提');
INSERT INTO `country` VALUES ('53', '45', 'DK', 'Denmark', '丹麦');
INSERT INTO `country` VALUES ('54', '1767', 'DM', 'Dominica', '多米尼加');
INSERT INTO `country` VALUES ('55', '1809', 'DO', 'dominican republic', '多米尼加共和国');
INSERT INTO `country` VALUES ('56', '213', 'DZ', 'Algeria', '阿尔及利亚');
INSERT INTO `country` VALUES ('57', '593', 'EC', 'Ecuador', '厄瓜多尔');
INSERT INTO `country` VALUES ('58', '372', 'EE', 'Estonia', '爱沙尼亚');
INSERT INTO `country` VALUES ('59', '20', 'EG', 'Egypt', '埃及');
INSERT INTO `country` VALUES ('60', '291', 'ER', 'Eritrea', '厄立特里亚');
INSERT INTO `country` VALUES ('61', '34', 'ES', 'Spain', '西班牙');
INSERT INTO `country` VALUES ('62', '251', 'ET', 'Ethiopia', '埃塞俄比亚');
INSERT INTO `country` VALUES ('63', '358', 'FI', 'Finland', '芬兰');
INSERT INTO `country` VALUES ('64', '679', 'FJ', 'Fiji', '斐济');
INSERT INTO `country` VALUES ('65', '691', 'FM', 'Micronesia', '密克罗尼西亚');
INSERT INTO `country` VALUES ('66', '298', 'FO', 'Faroe Islands', '法罗群岛');
INSERT INTO `country` VALUES ('67', '33', 'FR', 'France', '法国');
INSERT INTO `country` VALUES ('68', '241', 'GA', 'Gabon', '加蓬');
INSERT INTO `country` VALUES ('69', '44', 'GB', 'United Kingdom', '英国');
INSERT INTO `country` VALUES ('70', '1473', 'GD', 'Grenada', '格林纳达');
INSERT INTO `country` VALUES ('71', '995', 'GE', 'Georgia', '格鲁吉亚');
INSERT INTO `country` VALUES ('72', '594', 'GF', 'French Guiana', '法属圭亚那');
INSERT INTO `country` VALUES ('73', '233', 'GH', 'Ghana', '加纳');
INSERT INTO `country` VALUES ('74', '350', 'GI', 'Gibraltar', '直布罗陀');
INSERT INTO `country` VALUES ('75', '299', 'GL', 'Greenland', '格陵兰岛');
INSERT INTO `country` VALUES ('76', '220', 'GM', 'Gambia', '冈比亚');
INSERT INTO `country` VALUES ('77', '224', 'GN', 'Guinea', '几内亚');
INSERT INTO `country` VALUES ('78', '590', 'GP', 'Guadeloupe', '瓜德罗普岛');
INSERT INTO `country` VALUES ('79', '240', 'GQ', 'Equatorial Guinea', '赤道几内亚');
INSERT INTO `country` VALUES ('80', '30', 'GR', 'Greece', '希腊');
INSERT INTO `country` VALUES ('81', '502', 'GT', 'Guatemala', '瓜地马拉');
INSERT INTO `country` VALUES ('82', '1671', 'GU', 'Guam', '关岛');
INSERT INTO `country` VALUES ('83', '245', 'GW', 'Guinea-Bissau', '几内亚比绍共和国');
INSERT INTO `country` VALUES ('84', '592', 'GY', 'Guyana', '圭亚那');
INSERT INTO `country` VALUES ('85', '852', 'HK', 'Hong Kong', '中国香港');
INSERT INTO `country` VALUES ('86', '504', 'HN', 'Honduras', '洪都拉斯');
INSERT INTO `country` VALUES ('87', '385', 'HR', 'Croatia', '克罗地亚');
INSERT INTO `country` VALUES ('88', '509', 'HT', 'Haiti', '海地');
INSERT INTO `country` VALUES ('89', '36', 'HU', 'Hungary', '匈牙利');
INSERT INTO `country` VALUES ('90', '62', 'ID', 'Indonesia', '印度尼西亚');
INSERT INTO `country` VALUES ('91', '353', 'IE', 'Ireland', '爱尔兰');
INSERT INTO `country` VALUES ('92', '972', 'IL', 'Israel', '以色列');
INSERT INTO `country` VALUES ('93', '91', 'IN', 'India', '印度');
INSERT INTO `country` VALUES ('94', '964', 'IQ', 'Iraq', '伊拉克');
INSERT INTO `country` VALUES ('95', '98', 'IR', 'Iran', '伊朗');
INSERT INTO `country` VALUES ('96', '354', 'IS', 'Iceland', '冰岛');
INSERT INTO `country` VALUES ('97', '39', 'IT', 'Italy', '意大利');
INSERT INTO `country` VALUES ('98', '1876', 'JM', 'Jamaica', '牙买加');
INSERT INTO `country` VALUES ('99', '962', 'JO', 'Jordan', '约旦');
INSERT INTO `country` VALUES ('100', '81', 'JP', 'Japan', '日本');
INSERT INTO `country` VALUES ('101', '254', 'KE', 'Kenya', '肯尼亚');
INSERT INTO `country` VALUES ('102', '996', 'KG', 'Kyrgyzstan', '吉尔吉斯斯坦');
INSERT INTO `country` VALUES ('103', '855', 'KH', 'Cambodia', '柬埔寨');
INSERT INTO `country` VALUES ('104', '686', 'KI', 'Kiribati', '基里巴斯');
INSERT INTO `country` VALUES ('105', '269', 'KM', 'Comoros', '科摩罗');
INSERT INTO `country` VALUES ('106', '1869', 'KN', 'Saint Kitts and Nevis', '圣基茨和尼维斯');
INSERT INTO `country` VALUES ('107', '850', 'KP', 'Korea Democratic Rep.', '朝鲜');
INSERT INTO `country` VALUES ('108', '82', 'KR', 'South Korea', '韩国');
INSERT INTO `country` VALUES ('109', '965', 'KW', 'Kuwait', '科威特');
INSERT INTO `country` VALUES ('110', '1345', 'KY', 'Cayman Islands', '开曼群岛');
INSERT INTO `country` VALUES ('111', '7', 'KZ', 'Kazakhstan', '哈萨克斯坦');
INSERT INTO `country` VALUES ('112', '856', 'LA', 'Laos', '老挝');
INSERT INTO `country` VALUES ('113', '961', 'LB', 'Lebanon', '黎巴嫩');
INSERT INTO `country` VALUES ('114', '1758', 'LC', 'Saint Lucia', '圣露西亚');
INSERT INTO `country` VALUES ('115', '423', 'LI', 'Liechtenstein', '列支敦士登');
INSERT INTO `country` VALUES ('116', '94', 'LK', 'Sri Lanka', '斯里兰卡');
INSERT INTO `country` VALUES ('117', '231', 'LR', 'Liberia', '利比里亚');
INSERT INTO `country` VALUES ('118', '266', 'LS', 'Lesotho', '莱索托');
INSERT INTO `country` VALUES ('119', '370', 'LT', 'Lithuania', '立陶宛');
INSERT INTO `country` VALUES ('120', '352', 'LU', 'Luxembourg', '卢森堡');
INSERT INTO `country` VALUES ('121', '371', 'LV', 'Latvia', '拉脱维亚');
INSERT INTO `country` VALUES ('122', '218', 'LY', 'Libya', '利比亚');
INSERT INTO `country` VALUES ('123', '212', 'MA', 'Morocco', '摩洛哥');
INSERT INTO `country` VALUES ('124', '377', 'MC', 'Monaco', '摩纳哥');
INSERT INTO `country` VALUES ('125', '373', 'MD', 'Moldova', '摩尔多瓦');
INSERT INTO `country` VALUES ('126', '382', 'ME', 'Montenegro', '黑山');
INSERT INTO `country` VALUES ('127', '261', 'MG', 'Madagascar', '马达加斯加');
INSERT INTO `country` VALUES ('128', '692', 'MH', 'Marshall Islands', '马绍尔群岛');
INSERT INTO `country` VALUES ('129', '389', 'MK', 'Macedonia', '马其顿');
INSERT INTO `country` VALUES ('130', '223', 'ML', 'Mali', '马里');
INSERT INTO `country` VALUES ('131', '95', 'MM', 'Myanmar', '缅甸');
INSERT INTO `country` VALUES ('132', '976', 'MN', 'Mongolia', '蒙古');
INSERT INTO `country` VALUES ('133', '853', 'MO', 'Macau', '中国澳门');
INSERT INTO `country` VALUES ('134', '222', 'MR', 'Mauritania', '毛里塔尼亚');
INSERT INTO `country` VALUES ('135', '1664', 'MS', 'Montserrat', '蒙特塞拉特岛');
INSERT INTO `country` VALUES ('136', '356', 'MT', 'Malta', '马耳他');
INSERT INTO `country` VALUES ('137', '230', 'MU', 'Mauritius', '毛里求斯');
INSERT INTO `country` VALUES ('138', '960', 'MV', 'Maldives', '马尔代夫');
INSERT INTO `country` VALUES ('139', '265', 'MW', 'Malawi', '马拉维');
INSERT INTO `country` VALUES ('140', '52', 'MX', 'Mexico', '墨西哥');
INSERT INTO `country` VALUES ('141', '60', 'MY', 'Malaysia', '马来西亚');
INSERT INTO `country` VALUES ('142', '258', 'MZ', 'Mozambique', '莫桑比克');
INSERT INTO `country` VALUES ('143', '264', 'NA', 'Namibia', '纳米比亚');
INSERT INTO `country` VALUES ('144', '687', 'NC', 'New Caledonia', '新喀里多尼亚');
INSERT INTO `country` VALUES ('145', '227', 'NE', 'Niger', '尼日尔');
INSERT INTO `country` VALUES ('146', '234', 'NG', 'Nigeria', '尼日利亚');
INSERT INTO `country` VALUES ('147', '505', 'NI', 'Nicaragua', '尼加拉瓜');
INSERT INTO `country` VALUES ('148', '31', 'NL', 'Netherlands', '荷兰');
INSERT INTO `country` VALUES ('149', '47', 'NO', 'Norway', '挪威');
INSERT INTO `country` VALUES ('150', '977', 'NP', 'Nepal', '尼泊尔');
INSERT INTO `country` VALUES ('151', '674', 'NR', 'Nauru', '拿鲁岛');
INSERT INTO `country` VALUES ('152', '64', 'NZ', 'New Zealand', '新西兰');
INSERT INTO `country` VALUES ('153', '968', 'OM', 'Oman', '阿曼');
INSERT INTO `country` VALUES ('154', '507', 'PA', 'Panama', '巴拿马');
INSERT INTO `country` VALUES ('155', '51', 'PE', 'Peru', '秘鲁');
INSERT INTO `country` VALUES ('156', '689', 'PF', 'French Polynesia', '法属波利尼西亚');
INSERT INTO `country` VALUES ('157', '675', 'PG', 'Papua New Guinea', '巴布亚新几内亚');
INSERT INTO `country` VALUES ('158', '63', 'PH', 'Philippines', '菲律宾');
INSERT INTO `country` VALUES ('159', '92', 'PK', 'Pakistan', '巴基斯坦');
INSERT INTO `country` VALUES ('160', '48', 'PL', 'Poland', '波兰');
INSERT INTO `country` VALUES ('161', '508', 'PM', 'Saint Pierreand Miquelon', '圣彼埃尔和密克隆岛');
INSERT INTO `country` VALUES ('162', '1787', 'PR', 'Puerto Rico', '波多黎各');
INSERT INTO `country` VALUES ('163', '351', 'PT', 'Portugal', '葡萄牙');
INSERT INTO `country` VALUES ('164', '680', 'PW', 'Palau', '帕劳');
INSERT INTO `country` VALUES ('165', '595', 'PY', 'Paraguay', '巴拉圭');
INSERT INTO `country` VALUES ('166', '974', 'QA', 'Qatar', '卡塔尔');
INSERT INTO `country` VALUES ('167', '262', 'RE', 'Réunion Island', '留尼汪');
INSERT INTO `country` VALUES ('168', '40', 'RO', 'Romania', '罗马尼亚');
INSERT INTO `country` VALUES ('169', '381', 'RS', 'Serbia', '塞尔维亚');
INSERT INTO `country` VALUES ('170', '7', 'RU', 'Russia', '俄罗斯');
INSERT INTO `country` VALUES ('171', '250', 'RW', 'Rwanda', '卢旺达');
INSERT INTO `country` VALUES ('172', '966', 'SA', 'Saudi Arabia', '沙特阿拉伯');
INSERT INTO `country` VALUES ('173', '677', 'SB', 'Solomon Islands', '所罗门群岛');
INSERT INTO `country` VALUES ('174', '248', 'SC', 'Seychelles', '塞舌尔');
INSERT INTO `country` VALUES ('175', '249', 'SD', 'Sudan', '苏丹');
INSERT INTO `country` VALUES ('176', '46', 'SE', 'Sweden', '瑞典');
INSERT INTO `country` VALUES ('177', '65', 'SG', 'Singapore', '新加坡');
INSERT INTO `country` VALUES ('178', '386', 'SI', 'Slovenia', '斯洛文尼亚');
INSERT INTO `country` VALUES ('179', '421', 'SK', 'Slovakia', '斯洛伐克');
INSERT INTO `country` VALUES ('180', '232', 'SL', 'Sierra Leone', '塞拉利昂');
INSERT INTO `country` VALUES ('181', '378', 'SM', 'San Marino', '圣马力诺');
INSERT INTO `country` VALUES ('182', '221', 'SN', 'Senegal', '塞内加尔');
INSERT INTO `country` VALUES ('183', '252', 'SO', 'Somalia', '索马里');
INSERT INTO `country` VALUES ('184', '597', 'SR', 'Suriname', '苏里南');
INSERT INTO `country` VALUES ('185', '239', 'ST', 'Sao Tome and Principe', '圣多美和普林西比');
INSERT INTO `country` VALUES ('186', '503', 'SV', 'ElSalvador', '萨尔瓦多');
INSERT INTO `country` VALUES ('187', '963', 'SY', 'Syria', '叙利亚');
INSERT INTO `country` VALUES ('188', '268', 'SZ', 'Swaziland', '斯威士兰');
INSERT INTO `country` VALUES ('189', '1649', 'TC', 'Turksand Caicos Islands', '特克斯和凯科斯群岛');
INSERT INTO `country` VALUES ('190', '235', 'TD', 'Chad', '乍得');
INSERT INTO `country` VALUES ('191', '228', 'TG', 'Togo', '多哥');
INSERT INTO `country` VALUES ('192', '66', 'TH', 'Thailand', '泰国');
INSERT INTO `country` VALUES ('193', '992', 'TJ', 'Tajikistan', '塔吉克斯坦');
INSERT INTO `country` VALUES ('194', '670', 'TL', 'East Timor', '东帝汶');
INSERT INTO `country` VALUES ('195', '993', 'TM', 'Turkmenistan', '土库曼斯坦');
INSERT INTO `country` VALUES ('196', '216', 'TN', 'Tunisia', '突尼斯');
INSERT INTO `country` VALUES ('197', '676', 'TO', 'Tonga', '汤加');
INSERT INTO `country` VALUES ('198', '90', 'TR', 'Turkey', '土耳其');
INSERT INTO `country` VALUES ('199', '1868', 'TT', 'Trinidadand Tobago', '特立尼达和多巴哥');
INSERT INTO `country` VALUES ('200', '886', 'TW', 'Taiwan', '中国台湾');
INSERT INTO `country` VALUES ('201', '255', 'TZ', 'Tanzania', '坦桑尼亚');
INSERT INTO `country` VALUES ('202', '380', 'UA', 'Ukraine', '乌克兰');
INSERT INTO `country` VALUES ('203', '256', 'UG', 'Uganda', '乌干达');
INSERT INTO `country` VALUES ('204', '1', 'US', 'United States', '美国');
INSERT INTO `country` VALUES ('205', '598', 'UY', 'Uruguay', '乌拉圭');
INSERT INTO `country` VALUES ('206', '998', 'UZ', 'Uzbekistan', '乌兹别克斯坦');
INSERT INTO `country` VALUES ('207', '1784', 'VC', 'Saint Vincent and The Grenadines', '圣文森特和格林纳丁斯');
INSERT INTO `country` VALUES ('208', '58', 'VE', 'Venezuela', '委内瑞拉');
INSERT INTO `country` VALUES ('209', '1284', 'VG', 'VirginIslands,British', '英属处女群岛');
INSERT INTO `country` VALUES ('210', '84', 'VN', 'Vietnam', '越南');
INSERT INTO `country` VALUES ('211', '678', 'VU', 'Vanuatu', '瓦努阿图');
INSERT INTO `country` VALUES ('212', '685', 'WS', 'Samoa', '萨摩亚');
INSERT INTO `country` VALUES ('213', '967', 'YE', 'Yemen', '也门');
INSERT INTO `country` VALUES ('214', '269', 'YT', 'Mayotte', '马约特');
INSERT INTO `country` VALUES ('215', '27', 'ZA', 'South Africa', '南非');
INSERT INTO `country` VALUES ('216', '260', 'ZM', 'Zambia', '赞比亚');
INSERT INTO `country` VALUES ('217', '263', 'ZW', 'Zimbabwe', '津巴布韦');

-- ----------------------------
-- Table structure for `currency`
-- ----------------------------
DROP TABLE IF EXISTS `currency`;
CREATE TABLE `currency` (
  `currency_id` varchar(32) CHARACTER SET utf8mb4 NOT NULL COMMENT '主键',
  `currency_name` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '币名称',
  `currency_tag` varchar(50) CHARACTER SET utf8mb4 NOT NULL COMMENT '币标签',
  `currency_transtion_commission` bigint(11) NOT NULL DEFAULT '2' COMMENT '单笔手续费（按千分数计） 默认千分之二    ',
  `currency_transtion_commission_min` bigint(11) NOT NULL DEFAULT '0' COMMENT '最小手续费（非百分数）',
  `currency_create_time` datetime DEFAULT NULL COMMENT '币创建日期',
  `currency_status` int(11) DEFAULT '0' COMMENT '币状态	0 前端不可见 1 前端可见',
  `currency_recharge_status` int(11) DEFAULT '0' COMMENT '币充值状态 0 不可充值 1可充值',
  `currency_withdraw_status` int(11) DEFAULT '0' COMMENT '币提现状态 0 不可提现 1 可提现',
  `currency_withdraw_commission` varchar(11) CHARACTER SET utf8mb4 NOT NULL DEFAULT '2' COMMENT '提现手续费(固定数值,每种代币具体值不一样)',
  `currency_withdraw_commission_min` bigint(11) NOT NULL DEFAULT '0' COMMENT '提现最小手续费',
  `currency_funding_price` varchar(50) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '众筹价格  1ETH=1231231TNB',
  `currency_white_paper_url` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '白皮书地址',
  `currency_official_url` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '官网地址',
  `currency_block_check_url` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '区块查询地址',
  `currency_issue_time` datetime DEFAULT NULL COMMENT '发行时间',
  `currency_extract_count_min` varchar(20) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '提币最小数量',
  `currency_avatar` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '币头像',
  `ratio_to_usdt` decimal(50,0) DEFAULT NULL COMMENT '当前币兑换USDT的比例(eg:value=100，则代表100个USDT=1个该币)',
  `Ethereum_usable_status` int(11) unsigned zerofill DEFAULT NULL COMMENT '是否以太坊 0:非以太坊1：以太坊 默认为0',
  `Contract_address` varchar(255) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT '合约地址',
  `Contract_Decimal_NB` int(11) DEFAULT NULL COMMENT '合约小数点位数',
  PRIMARY KEY (`currency_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of currency
-- ----------------------------
INSERT INTO `currency` VALUES ('100001', 'Bitcoin Cash', 'BCH', '0', '0', '2018-03-01 17:28:55', '1', '1', '1', '2.0', '0', '1ETH=10000', 'http://white_paper_url', 'http://white_official_url', 'http://block_check_url', '2018-03-06 00:00:00', '0.002', '0', '8970', '00000000001', '/312312', '0');
INSERT INTO `currency` VALUES ('100002', 'Bitcoin', 'BTC', '0', '0', '2018-03-01 17:28:55', '1', '1', '1', '2.0', '0', '1ETH=10000', 'http://white_paper_url', 'http://white_official_url', 'http://block_check_url', '2018-03-06 00:00:00', '0.002', '0', '8970', '00000000000', '/12312', '0');
INSERT INTO `currency` VALUES ('100003', 'Ethereum', 'ETH', '0', '0', '2018-03-01 17:28:55', '1', '1', '1', '2.0', '0', '1ETH=10000', 'http://white_paper_url', 'http://white_official_url', 'http://block_check_url', '2018-03-06 00:00:00', '0.002', '0', '8970', '00000000001', '/132123', '0');
INSERT INTO `currency` VALUES ('100004', 'Bitcoin Cash', 'BCH', '0', '0', '2018-03-01 17:28:55', '1', '1', '1', '2.0', '0', '1ETH=10000', 'http://white_paper_url', 'http://white_official_url', 'http://block_check_url', '2018-03-06 00:00:00', '0.002', '0', '8970', '00000000000', '/13123', '0');
INSERT INTO `currency` VALUES ('100006', '32', '1', '23', '1', '2018-05-25 00:00:00', '1', '0', '0', '2', '2', '23', '2', '2', '2', '2018-05-25 00:00:00', '2', '2', '2', '00000000000', '2', '2');
INSERT INTO `currency` VALUES ('100009', 'hgfh', 'ngf', '0', '1', '2018-05-25 00:00:00', '1', '0', '0', '1', '1', '12', '1', '1', '1', '2018-05-25 00:00:00', '1', '1', '1', '00000000000', '1', '1');
INSERT INTO `currency` VALUES ('100010', '2321312', '113', '1', '1', '2018-05-25 00:00:00', '0', '1', '1', '1', '1', '23123', '1', '1', '1', '2018-05-01 00:00:00', '1', '1', '1', '00000000000', '1', '1');
INSERT INTO `currency` VALUES ('155555', '11', '1', '1', '1', '2018-05-25 00:00:00', '0', '1', '1', '1', '1', '1', '1', '1', '1', '2018-05-25 00:00:00', '1', '1', '0', '00000000000', '1', '1');

-- ----------------------------
-- Table structure for `ico_order`
-- ----------------------------
DROP TABLE IF EXISTS `ico_order`;
CREATE TABLE `ico_order` (
  `id` bigint(20) NOT NULL,
  `project_id` bigint(20) DEFAULT NULL COMMENT '项目id',
  `stage_id` int(11) DEFAULT NULL COMMENT '购买阶段',
  `price` bigint(20) DEFAULT NULL COMMENT '价格',
  `status` int(11) DEFAULT NULL COMMENT '状态',
  `buy_coin` varchar(20) DEFAULT NULL COMMENT '购买的币id',
  `buy_count` bigint(20) DEFAULT NULL COMMENT '购买数量',
  `pay_coin` varchar(20) DEFAULT NULL COMMENT '支付的币id',
  `pay_count` bigint(20) DEFAULT NULL COMMENT '支付数量',
  `order_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '下单时间',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `is_settlement` int(11) DEFAULT NULL COMMENT '是否已经结算',
  `settlement_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '结算时间',
  `settlement_note` text COMMENT '结算备注',
  PRIMARY KEY (`id`),
  KEY `project_id` (`project_id`),
  CONSTRAINT `ico_order_ibfk_1` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of ico_order
-- ----------------------------

-- ----------------------------
-- Table structure for `inviter_info`
-- ----------------------------
DROP TABLE IF EXISTS `inviter_info`;
CREATE TABLE `inviter_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT '邀请人id',
  `inviter_id` bigint(20) DEFAULT NULL COMMENT '被邀请人id',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of inviter_info
-- ----------------------------
INSERT INTO `inviter_info` VALUES ('1', '1', '3', '2018-05-10 15:59:41');

-- ----------------------------
-- Table structure for `kyc_info`
-- ----------------------------
DROP TABLE IF EXISTS `kyc_info`;
CREATE TABLE `kyc_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `state` int(11) DEFAULT NULL COMMENT '状态',
  `kind` int(11) DEFAULT NULL COMMENT '用户类型',
  `real_name` text COMMENT '真实姓名',
  `country_code` int(11) DEFAULT NULL COMMENT '国家代码',
  `identity_type` int(11) DEFAULT NULL COMMENT '证件类型',
  `identity_id` text COMMENT '证件Id',
  `photo_front` text COMMENT '证件正面照',
  `photo_back` text COMMENT '证件背面照',
  `photo_hand` text COMMENT '证件手持照',
  `reason` text COMMENT '审核原因',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `rejected_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '拒绝时间',
  `passed_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT '通过时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of kyc_info
-- ----------------------------
INSERT INTO `kyc_info` VALUES ('1', '1', '0', '1', '张三', '2', '1', '3213213', null, null, null, null, '2018-05-18 11:36:27', '2018-05-18 11:36:27', '2018-05-18 11:36:27', '2018-05-18 11:36:27');
INSERT INTO `kyc_info` VALUES ('2', '2', '2', '1', '李四', '3', '2', '123213', null, null, null, null, '2018-05-18 13:06:16', '2018-05-18 13:06:15', '2018-05-24 11:17:44', '2018-05-18 13:06:15');
INSERT INTO `kyc_info` VALUES ('3', '3', '1', '1', '王五', '4', '1', '4242', null, null, null, null, '2018-05-21 11:47:04', '2018-05-09 11:10:47', null, '2018-05-24 11:17:36');
INSERT INTO `kyc_info` VALUES ('4', '5', '0', '1', 'chaoshu', '24', '1', '4567869098', '/StaticFiles/ertyjjfdhhjgj.png', '/StaticFiles/ertyjjfdhhjgj.png', '/StaticFiles/ertyjjfdhhjgj.png', '', '2018-05-24 13:53:57', null, null, null);

-- ----------------------------
-- Table structure for `kyc_requests`
-- ----------------------------
DROP TABLE IF EXISTS `kyc_requests`;
CREATE TABLE `kyc_requests` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL,
  `stage` int(11) DEFAULT NULL,
  `state` int(11) DEFAULT NULL,
  `kind` int(11) DEFAULT NULL,
  `resource` json DEFAULT NULL,
  `reason` text,
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `rejected_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `passed_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of kyc_requests
-- ----------------------------
INSERT INTO `kyc_requests` VALUES ('1', '1', '1', '0', '1', '{\"realName\": \"舒超\", \"identityid\": \"234567897098765432134567\"}', null, '2018-05-18 11:36:27', '2018-05-18 11:36:27', '2018-05-18 11:36:27');
INSERT INTO `kyc_requests` VALUES ('2', '2', '1', '0', '1', '{\"photo0\": \"http://www.baidu/photo1.png\", \"photo2\": \"http://www.baidu/photo2.png\"}', null, '2018-05-18 13:06:15', '2018-05-18 13:06:15', '2018-05-18 13:06:15');
INSERT INTO `kyc_requests` VALUES ('3', '1', '2', '0', '1', '{\"realName\": \"舒超\", \"identityid\": \"234567897098765432134567\"}', '', '2018-05-21 11:47:03', null, null);

-- ----------------------------
-- Table structure for `lock_record`
-- ----------------------------
DROP TABLE IF EXISTS `lock_record`;
CREATE TABLE `lock_record` (
  `id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `coin_id` varchar(20) DEFAULT NULL,
  `count` bigint(20) DEFAULT NULL COMMENT '锁仓数量',
  `lock_type` int(11) DEFAULT NULL COMMENT '锁仓类型',
  `lock_note` text COMMENT '锁仓原因',
  `status` int(11) DEFAULT NULL COMMENT '锁仓状态',
  `lock_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '锁仓时间',
  `unlock_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '解锁时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of lock_record
-- ----------------------------

-- ----------------------------
-- Table structure for `login_records`
-- ----------------------------
DROP TABLE IF EXISTS `login_records`;
CREATE TABLE `login_records` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `login_ip` text COMMENT '登录ip',
  `device_id` text COMMENT '登录设备id',
  `login_status` int(11) DEFAULT NULL COMMENT '登录状态',
  `login_comment` text COMMENT '登录状态信息',
  `login_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '登录时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of login_records
-- ----------------------------
INSERT INTO `login_records` VALUES ('1', '2', '192.168.0.2', 'sfddghjllkjtgfe', '1', 'Success', '2018-05-10 15:51:56');
INSERT INTO `login_records` VALUES ('2', '2', '192.168.0.2', 'sfddghjllkjtgfe', '2', 'Password not correct', '2018-05-10 15:52:50');
INSERT INTO `login_records` VALUES ('3', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 11:32:33');
INSERT INTO `login_records` VALUES ('4', '2', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 11:35:29');
INSERT INTO `login_records` VALUES ('5', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 11:38:46');
INSERT INTO `login_records` VALUES ('6', '2', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 11:41:03');
INSERT INTO `login_records` VALUES ('7', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 11:43:50');
INSERT INTO `login_records` VALUES ('8', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 11:52:18');
INSERT INTO `login_records` VALUES ('9', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 11:55:41');
INSERT INTO `login_records` VALUES ('10', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 12:55:38');
INSERT INTO `login_records` VALUES ('11', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 12:59:22');
INSERT INTO `login_records` VALUES ('12', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 13:16:30');
INSERT INTO `login_records` VALUES ('13', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 13:44:19');
INSERT INTO `login_records` VALUES ('14', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 13:45:22');
INSERT INTO `login_records` VALUES ('15', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 13:46:11');
INSERT INTO `login_records` VALUES ('16', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 13:57:14');
INSERT INTO `login_records` VALUES ('17', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 13:57:47');
INSERT INTO `login_records` VALUES ('18', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-14 18:25:32');
INSERT INTO `login_records` VALUES ('19', '1', '127.0.0.1', 'postman_001', '2', 'Password not correct', '2018-05-14 18:25:44');
INSERT INTO `login_records` VALUES ('20', '1', '127.0.0.1', 'postman_001', '2', 'Password not correct', '2018-05-14 18:30:32');
INSERT INTO `login_records` VALUES ('21', '1', '127.0.0.1', 'postman_001', '2', 'Password not correct', '2018-05-14 18:31:25');
INSERT INTO `login_records` VALUES ('22', '1', '127.0.0.1', 'postman_001', '2', 'Password not correct', '2018-05-15 12:35:54');
INSERT INTO `login_records` VALUES ('23', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-15 12:36:02');
INSERT INTO `login_records` VALUES ('24', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-18 11:36:22');
INSERT INTO `login_records` VALUES ('25', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-18 11:51:44');
INSERT INTO `login_records` VALUES ('26', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-18 15:48:18');
INSERT INTO `login_records` VALUES ('27', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-18 15:48:29');
INSERT INTO `login_records` VALUES ('28', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-21 11:38:58');
INSERT INTO `login_records` VALUES ('29', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-21 11:40:40');
INSERT INTO `login_records` VALUES ('30', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-21 11:44:12');
INSERT INTO `login_records` VALUES ('31', '1', '127.0.0.1', 'postman_001', '1', 'Success', '2018-05-21 11:59:39');
INSERT INTO `login_records` VALUES ('32', '1', '127.0.0.1', 'postman_001', '2', 'account disabled', '2018-05-24 13:35:34');
INSERT INTO `login_records` VALUES ('33', '1', '127.0.0.1', 'postman_001', '2', 'account disabled', '2018-05-24 13:36:04');
INSERT INTO `login_records` VALUES ('34', '1', '127.0.0.1', 'postman_001', '2', 'account disabled', '2018-05-24 13:36:06');
INSERT INTO `login_records` VALUES ('35', '1', '127.0.0.1', 'postman_001', '2', 'account disabled', '2018-05-24 13:36:09');
INSERT INTO `login_records` VALUES ('36', '1', '127.0.0.1', 'postman_001', '2', 'account disabled', '2018-05-24 13:36:50');
INSERT INTO `login_records` VALUES ('37', '1', '127.0.0.1', 'postman_001', '2', 'account disabled', '2018-05-24 13:36:52');
INSERT INTO `login_records` VALUES ('38', '1', '127.0.0.1', 'postman_001', '2', 'account disabled', '2018-05-24 13:48:40');
INSERT INTO `login_records` VALUES ('39', '1', '127.0.0.1', 'postman_001', '1', 'success', '2018-05-24 13:49:13');
INSERT INTO `login_records` VALUES ('40', '1', '127.0.0.1', 'postman_001', '1', 'success', '2018-05-24 17:43:49');
INSERT INTO `login_records` VALUES ('41', '1', '127.0.0.1', 'postman_001', '1', 'success', '2018-05-24 17:45:00');

-- ----------------------------
-- Table structure for `members_info`
-- ----------------------------
DROP TABLE IF EXISTS `members_info`;
CREATE TABLE `members_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `project_id` bigint(20) DEFAULT NULL COMMENT '项目id',
  `name` text COMMENT '姓名',
  `position` text COMMENT '职位',
  `description` text COMMENT '用户描述i信息',
  `member_type` int(11) DEFAULT NULL COMMENT '用户类型',
  `image` text COMMENT '头像',
  `priority_sort` int(11) DEFAULT NULL COMMENT '排序',
  `join_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '加入项目时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `project_id` (`project_id`),
  CONSTRAINT `members_info_ibfk_1` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of members_info
-- ----------------------------
INSERT INTO `members_info` VALUES ('1', '1', 'chaoshu', ' Backend dev', '后端开发', '2', 'http://www.personal.png', '1', '2018-05-26 15:54:12', '2018-05-26 15:54:12');
INSERT INTO `members_info` VALUES ('2', '1', 'guoming', 'CTO', '技术总监', '2', 'http://www.guoming.png', '0', '2018-05-26 15:54:54', '2018-05-26 15:54:54');

-- ----------------------------
-- Table structure for `msg_records`
-- ----------------------------
DROP TABLE IF EXISTS `msg_records`;
CREATE TABLE `msg_records` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `destination` text,
  `message` text,
  `send_status` int(11) DEFAULT NULL,
  `return_message` text,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of msg_records
-- ----------------------------
INSERT INTO `msg_records` VALUES ('1', '18500797779', '[Coindaq]:您在币信科技请求的验证码为:8230', '2', '网络错误', '2018-05-15 20:44:26');
INSERT INTO `msg_records` VALUES ('2', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 7113', '1', '', '2018-05-15 20:46:13');
INSERT INTO `msg_records` VALUES ('3', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 5448', '1', '', '2018-05-15 20:48:08');
INSERT INTO `msg_records` VALUES ('4', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 1942', '1', '', '2018-05-15 20:50:33');
INSERT INTO `msg_records` VALUES ('5', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 8758', '1', '', '2018-05-15 20:50:47');
INSERT INTO `msg_records` VALUES ('6', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 5419', '1', '', '2018-05-15 20:50:50');
INSERT INTO `msg_records` VALUES ('7', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 6098', '1', '', '2018-05-15 20:50:52');
INSERT INTO `msg_records` VALUES ('8', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 8343', '1', '', '2018-05-15 20:50:53');
INSERT INTO `msg_records` VALUES ('9', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 5101', '1', '', '2018-05-15 20:50:54');
INSERT INTO `msg_records` VALUES ('10', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 9775', '1', '', '2018-05-15 20:50:54');
INSERT INTO `msg_records` VALUES ('11', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 6198', '1', '', '2018-05-15 20:50:56');
INSERT INTO `msg_records` VALUES ('12', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 5870', '1', '', '2018-05-15 20:50:56');
INSERT INTO `msg_records` VALUES ('13', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 7563', '1', '', '2018-05-15 20:50:57');
INSERT INTO `msg_records` VALUES ('14', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 5734', '1', '', '2018-05-15 20:50:59');
INSERT INTO `msg_records` VALUES ('15', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 9151', '1', '', '2018-05-15 20:51:00');
INSERT INTO `msg_records` VALUES ('16', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 0656', '1', '', '2018-05-15 20:51:01');
INSERT INTO `msg_records` VALUES ('17', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 4922', '1', '', '2018-05-15 21:03:29');
INSERT INTO `msg_records` VALUES ('18', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 2195', '1', '', '2018-05-15 21:03:48');
INSERT INTO `msg_records` VALUES ('19', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 4057', '1', '', '2018-05-15 21:04:53');
INSERT INTO `msg_records` VALUES ('20', '18500797779', '[Coindaq]:您在币信科技请求的验证码为: 9225', '1', '', '2018-05-15 21:06:13');

-- ----------------------------
-- Table structure for `notice_news`
-- ----------------------------
DROP TABLE IF EXISTS `notice_news`;
CREATE TABLE `notice_news` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `project_id` bigint(20) DEFAULT NULL COMMENT '项目id',
  `title` text COMMENT '公告标题',
  `brief_introduction` text COMMENT '公告简介',
  `description` text COMMENT '公告描述信息',
  `notice_type` int(11) DEFAULT NULL,
  `send_type` int(11) DEFAULT NULL COMMENT '公告接受用户',
  `status` int(11) DEFAULT NULL COMMENT '信息状态',
  `reason` text COMMENT '审核原因',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `expire_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '过期时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `project_id` (`project_id`),
  CONSTRAINT `notice_news_ibfk_1` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of notice_news
-- ----------------------------
INSERT INTO `notice_news` VALUES ('1', '1', '项目普通公告', '项目成立公告新闻', '项目描述信息', '1', '1', '0', null, '2018-05-26 16:01:58', '2018-05-26 16:01:52', '2018-05-26 16:01:52');
INSERT INTO `notice_news` VALUES ('2', '1', '项目投票公告', '项目投票成功', '项目投票公告描述', '2', '1', '0', null, '2018-05-26 16:03:02', '2018-05-26 16:03:02', '2018-05-26 16:03:02');

-- ----------------------------
-- Table structure for `notice_vote`
-- ----------------------------
DROP TABLE IF EXISTS `notice_vote`;
CREATE TABLE `notice_vote` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `notice_id` bigint(20) DEFAULT NULL,
  `approve_vote` bigint(20) DEFAULT NULL COMMENT '赞成票',
  `disapprove_vote` bigint(20) DEFAULT NULL COMMENT '反对票',
  `abstention_vote` bigint(20) DEFAULT NULL COMMENT '弃权票',
  `platform_vote_max` bigint(20) DEFAULT NULL COMMENT '平台票（最大值）',
  `platform_vote_volumn` bigint(20) DEFAULT NULL COMMENT '平台投票数量',
  `platform_vote_type` int(11) DEFAULT NULL COMMENT '平台投票类型（赞成，反对）',
  `platform_vote_reason` text COMMENT '投票原因',
  `platform_vote_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '投票时间',
  `vote_result` int(11) DEFAULT NULL COMMENT '投票结果',
  `create_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `begin_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '投票开始时间',
  `end_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '投票结束时间',
  PRIMARY KEY (`id`),
  KEY `notice_id` (`notice_id`),
  CONSTRAINT `notice_vote_ibfk_1` FOREIGN KEY (`notice_id`) REFERENCES `notice_news` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of notice_vote
-- ----------------------------
INSERT INTO `notice_vote` VALUES ('1', '2', '0', '0', null, '200', '0', '0', null, '2018-05-26 16:04:06', null, null, '2018-05-26 16:04:06', '2018-05-26 16:04:06', null);

-- ----------------------------
-- Table structure for `notice_vote_details`
-- ----------------------------
DROP TABLE IF EXISTS `notice_vote_details`;
CREATE TABLE `notice_vote_details` (
  `id` bigint(20) NOT NULL,
  `notice_id` bigint(20) DEFAULT NULL COMMENT '提案id',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `vote_type` int(11) DEFAULT NULL COMMENT '投票类型',
  `vote_volumn` bigint(20) DEFAULT NULL COMMENT '投票量',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '投票时间',
  PRIMARY KEY (`id`),
  KEY `notice_id` (`notice_id`),
  CONSTRAINT `notice_vote_details_ibfk_1` FOREIGN KEY (`notice_id`) REFERENCES `notice_news` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of notice_vote_details
-- ----------------------------

-- ----------------------------
-- Table structure for `projects`
-- ----------------------------
DROP TABLE IF EXISTS `projects`;
CREATE TABLE `projects` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `admin_id` bigint(20) DEFAULT NULL COMMENT '管理员Id',
  `summary` text COMMENT '项目名称',
  `description` text COMMENT '项目详细介绍',
  `target_value` bigint(20) DEFAULT NULL COMMENT '项目集资总额',
  `stage_count` int(11) DEFAULT NULL COMMENT '项目分期数量',
  `classify` text COMMENT '项目类别',
  `white_paper` text COMMENT '白皮书',
  `official_site` text COMMENT '官网地址',
  `community_address` text COMMENT '社群地址',
  `status` int(11) DEFAULT NULL COMMENT '项目状态',
  `priority_sort` int(11) DEFAULT NULL COMMENT '排序优先级',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `begin_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
  `end_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `id` (`admin_id`),
  CONSTRAINT `id` FOREIGN KEY (`admin_id`) REFERENCES `users` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of projects
-- ----------------------------
INSERT INTO `projects` VALUES ('1', '2', '测试项目1', '项目描述信息', '500000', '2', 'ICO类型', 'http://www.whitepaper.com', 'http://www.officesite.com', 'http://www.community.com', '0', '1', '2018-05-26 15:55:16', '2018-05-26 15:55:16', '2018-05-31 15:48:05', '2018-05-26 15:48:05');

-- ----------------------------
-- Table structure for `projects_certification`
-- ----------------------------
DROP TABLE IF EXISTS `projects_certification`;
CREATE TABLE `projects_certification` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `project_id` bigint(20) DEFAULT NULL,
  `certification_id` bigint(20) DEFAULT NULL,
  `priority_sort` int(11) DEFAULT NULL COMMENT '排序优先级',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '添加时间',
  PRIMARY KEY (`id`),
  KEY `certification_id` (`certification_id`),
  KEY `project_id` (`project_id`),
  CONSTRAINT `projects_certification_ibfk_1` FOREIGN KEY (`certification_id`) REFERENCES `certification_info` (`id`),
  CONSTRAINT `projects_certification_ibfk_2` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of projects_certification
-- ----------------------------
INSERT INTO `projects_certification` VALUES ('1', '1', '1', '0', '2018-05-26 16:51:29');
INSERT INTO `projects_certification` VALUES ('2', '1', '2', '2', '2018-05-26 16:51:43');
INSERT INTO `projects_certification` VALUES ('3', '1', '3', '1', '2018-05-26 16:51:51');

-- ----------------------------
-- Table structure for `projects_media`
-- ----------------------------
DROP TABLE IF EXISTS `projects_media`;
CREATE TABLE `projects_media` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `project_id` bigint(20) DEFAULT NULL COMMENT '项目id',
  `title` text COMMENT '标题',
  `address` text COMMENT '地址',
  `type` int(11) DEFAULT '0' COMMENT '类型，图片或视频',
  `enable` int(11) DEFAULT '0' COMMENT '是否显示',
  `priority_sort` int(11) DEFAULT NULL COMMENT '排序优先级',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `project_id` (`project_id`),
  CONSTRAINT `projects_media_ibfk_1` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of projects_media
-- ----------------------------
INSERT INTO `projects_media` VALUES ('1', '1', 'banner1', 'http://www.png1.com', '1', '1', '0', '2018-05-26 15:49:47');
INSERT INTO `projects_media` VALUES ('2', '1', 'banner2', 'http://www.png2.com', '1', '1', '1', '2018-05-26 15:50:26');
INSERT INTO `projects_media` VALUES ('3', '1', 'video1', 'http://www.video1.avi', '2', '1', '0', '2018-05-26 15:50:54');
INSERT INTO `projects_media` VALUES ('4', '1', 'video2', 'http://www.video2.avi', '2', '1', '1', '2018-05-26 15:51:20');
INSERT INTO `projects_media` VALUES ('5', '1', 'video3', 'http://www.video3.avi', '2', '1', '2', '2018-05-26 15:53:09');

-- ----------------------------
-- Table structure for `projects_vote_details`
-- ----------------------------
DROP TABLE IF EXISTS `projects_vote_details`;
CREATE TABLE `projects_vote_details` (
  `id` bigint(20) NOT NULL,
  `project_id` bigint(20) DEFAULT NULL COMMENT '项目id',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `vote_volumn` bigint(20) DEFAULT NULL COMMENT '投票数量',
  `vote_type` int(11) DEFAULT NULL COMMENT '投票类型',
  `vote_reason` text COMMENT '投票原因',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '投票时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of projects_vote_details
-- ----------------------------

-- ----------------------------
-- Table structure for `receive_bch`
-- ----------------------------
DROP TABLE IF EXISTS `receive_bch`;
CREATE TABLE `receive_bch` (
  `receive_address` varchar(128) NOT NULL,
  `receive_user_id` varchar(64) DEFAULT NULL,
  `receive_source_server` varchar(32) DEFAULT NULL,
  `receive_bind_date` datetime DEFAULT NULL,
  `receive_status` int(11) DEFAULT '0' COMMENT '0 为启用 1为禁用',
  PRIMARY KEY (`receive_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of receive_bch
-- ----------------------------
INSERT INTO `receive_bch` VALUES ('mfcczsfpnb9ZcEGTBbQxvJjQ3fCLvhbcax', '10', null, null, '0');
INSERT INTO `receive_bch` VALUES ('mfpCk7TPpQVUykcb1G9fY8uzj8TP3TALyd', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mfwtac2LxeSCHA7AzXaWco8UhBdmp3RwdS', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mfy1cmpuKJahgfGGsQcjgxdfCPGzVRzU7L', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mfztzdwyhbFVi2rW4j33kNKDzahQFQg3Vh', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mg6fv8GVwGt85qKj9ft1uJxtpwf2Qc2sqw', 'gm1', null, null, '0');
INSERT INTO `receive_bch` VALUES ('mg9M3BWsV9CmqrC9ENjpSNMkiB3R9RceZT', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mgadY259BgLfCZHuWzikGQizTZaU4ePgLN', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mgg6Vt3N72xhF3S9e6247muH5PkpNujD3T', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mgJPfH8GaR9TTinqYXmGqVqmSvhwkFHMwQ', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mgwcG8vMiYqwrxkhYGbpiDEC1B7MHVrvM8', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mgWW9ycByv9nbpmow1tBp1uPpLTU6WDWZy', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mgXJ4MFAKEHS4e3s4m2RZ19JhaMhaWLVJZ', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mhEErMZzcdQdUqB3UpSqwrU6q2DQb6U4pe', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mhh5Cv6bEiriJwJFwyH5gVxusK7eFNwbF4', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mhjVEFvjTxYiXf9xP5KoPJYurDBxGghrFe', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mhR6spsphRQ5yD7JQ2GE3vpYb3WTzBCZh4', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mhyynuC19jQDWYft7a44QuW2FoHVYJNAzK', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mic5JGcRJcDJWoAF9EDafhATbdTHkVjNj8', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('miKKzGi7JB9PotFSPmqE46EEe1toykJA5P', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('misHk1omf6K7EirQ7MHPxCnSZk4K8caeGL', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mj1dnSwNBESR9QDzXvFbms66L44BEGJ5jY', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mj98kJ2XsJ46DRviDtXQrQg5H9G78ZYwHr', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mjE3NYJnawLtcMiXuNoeAuxdfTqodVCzS3', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mjrWut729z6zUT95nHSzyhhUhpQcbDqRc6', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mjWmKVsJ5jFWErCZnPc4mKRdQ2HJ7VJAgW', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mkcEtoRgDemQ6R9bC8YpXbQSMxz6Lh9dao', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mkF5Mtq2xRLUpV6GJkVe1fKyjgEYhS55eo', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mkKuYY9cyFquw6jHt5tsRkZgPrXtCJhYCj', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mm7EqqBqRVCW3LXy9r9TMe2jjP6zWuE1eG', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mmdQ6x4vMBdRkdoqWqjbgi1knJMVckcnBn', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mmfEfAsXPCwS6mBApqqVYMhCKpEtYuv2oh', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mmfhDy7qquJufoaHg8b4PYtZ2jaH9F7xfQ', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mmkGbBr8W6V6VaPJVKwGeXvTXxETyAujof', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mmtsEi3KEFynKsSn37eurbcgPJk6s8MaCb', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mmvR8n1dniq8bQ1tXqEkuUFaVDWrB2Kn5Q', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mngDmZMPy9BcbqifzBVedYVaJbziUX9TqD', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mo6yDkLDGcvk8mYeCD5KPctTHEkzvHBHAq', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('moi7TdsRit5iDkQwnwANrNdZtS1XkCzx6D', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('moRfpbAVtHhE4bGmhBs3zm9PhrzUPrJABS', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('moybetw1uf9yXGE54CgpFUAeEnE2RtQcNc', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mpi2n6BCbgQbRW4bqTGfo5tigi8eTV9uyL', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mpMeodDfboMtMcpdbs47zjbcotZaTL19u3', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mpuh998A7odD6BvHEJdoxBDujzcXyw4Bmz', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mq6Ah8MxYHaQNN1pTPZSRYB9YTkdbn7A6N', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mqwZWjx1iLtP5ucWD8Pfmp3DsPV8TyLE3N', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mqzsSbUmq4Jn3ebB82qJyAu7tM5AYxK1Yj', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mrbdcT8cdXED6NgtkRxzUrYmygUv2y6rvy', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mrm6jhpSR57VFQgcDntkTm3YGvWdXkgDzZ', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mrRTE3b69ssT3AtECrepDTmAqSDAhLXQjn', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mrrVr8stUD8Hj7kdPefr5SZ2GSKcqyvdAb', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('ms6r4hBNnbuCt4oF1B3Ne65b3wy44RfbCj', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('msR2obNRiFdQ4qa73EzbmU9g6GgNo2MAzi', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('msSiiAbKuQJCG6kKtDGRRokAFixFLqebJ5', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('msWeA3FtcUdr6UdfyawxyHHvVcT7suRV7J', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mt7SZt4fkfsbqyczANvvNYmxhNGvgmgQJ9', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mtBM3UPNaASdEMm97LabAHNxn8C6LmAbwH', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mtidiSZhzG81DAsCBhumjwcdqELDXaz6hY', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mtszEYW1jjmBrF7LqPkd1LT98ksSGwMY7n', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('muC6E5A9zUi9KvzkpkVVmi3og1vhybFV6L', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mufGtkwNgum5P92szg6UsLDdLENQAm3Mcd', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mukGyqwjZdT1bs4J9tB2toEQDDmsrnXwHc', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('muUyEZ44X6dnkMhVRBMA2knUMV26D4CgxA', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mvB5UpHAZHdViN6fY11HxPRJ2rwtCtnjHM', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mvEucimB9JcgCTdE2VvVcSStevYzfAbXcA', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mvjU8jfmdVAZZa4B42AxUsZPFeAJkT1MAP', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mvmyDdJ1dp6eScZYzjkY1B7JK8M59MJSW5', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mwgisoT8DcMdSaVXEdEHMQk4XX1w5xHCmU', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mwjheKyRjhVpE8Xu86RPf6sRBZbRjXHgvm', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mwKhYxqF2XsSroUmK5J1DQrMJtU2yxG7Qd', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mwqe9hFwG6A1RtkAE9edbLic5W9BD9SKv8', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mwvpDSZ2YCDNDthTqo7b3L3upoXTKQCZMc', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mwycf67qKNjCpUe92aiNHhBeSMDhAvAMyc', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mwySx3pZobofpvCBNWUcNMzWwLn4MJ4izh', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mx1EZc3NtSEyAQSx9EnJgAsj1hetBa2Dnx', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mxmA1o6mkabvDTHhWCAHjjoM4NDVqR5xQ4', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('myjWcGqAhfQBPJbk8EMnmmWXtwnYVKEnk6', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('myRzz5kwuC5FshQhNZwH2zuHDfsuSb5qbh', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mz7mzmWz7E5SncMyogjNPisY5XbfcQzfrT', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mzjR1Tn1PEoM5GXDrnzYP2YJ7DuuKtdQ79', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mzoug6ArLZRy4PpLjeckWBEn4qovfvfwiw', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('mzZJ9xs2AoUjaLu8NGkKLzduwRgxnxQ9Ju', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n1FyqXZxRUXKEArW7ficPyPffWqtog4gcA', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n1RirUNYfxbvaPAXbGstLdLnSsuq9w8j9P', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n1W3NnesbF2hx6D8mJcjZFX6FHJtPgvFQE', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n1WdfBMso1dMBuezgp7ErGXsKm2e3o13Jt', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n28KFPE9h8PqfWwPZ9MPFe3av9HmPj6P2S', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n2bb8qnvZyD6z6rATtkG5H3cyVRmisDyT9', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n2H4VcRUyhe26eGhZQWDkRM7uwrimy8PNh', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n2hFtzJq8FhbfY8FffC4dQgstZiwNQmx51', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n2q1WhmD3zqojjkHeQDA7CyF5U8udiRag8', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n2u5jHyE65nea9Qhn2xMx2vjyb1QPXsLR9', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n34eGxmztxQz5THGzGnpyPDKQkoXtkZ7wu', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n3BhfAtvfFf3JTT2vseHB5JDPmMqCarzrW', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n3p9JMg6U5T5HoW2g3mEh2wGDjQUz58sGP', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n3WMVnLDYPHHpbE4C6Wp2DdLtfNjVipMCb', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n3wwfrqCW1epwSmzinHmfRmwWHdKWLTbsz', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n4g8oHiU19upjpt2XmKisQqG6dSNsj8XfP', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n4Ky6tRgtq6sjbkHpAPxbYMmfeJ8WXEpZt', null, null, null, '0');
INSERT INTO `receive_bch` VALUES ('n4TVNTy98kkGVQEXoddmHRzdWJtsstu47U', null, null, null, '0');

-- ----------------------------
-- Table structure for `receive_btc`
-- ----------------------------
DROP TABLE IF EXISTS `receive_btc`;
CREATE TABLE `receive_btc` (
  `receive_address` varchar(128) NOT NULL,
  `receive_user_id` varchar(64) DEFAULT NULL,
  `receive_source_server` varchar(32) DEFAULT NULL,
  `receive_bind_date` datetime DEFAULT NULL,
  `receive_status` int(11) DEFAULT '0' COMMENT '0 为启用 1为禁用',
  PRIMARY KEY (`receive_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of receive_btc
-- ----------------------------
INSERT INTO `receive_btc` VALUES ('mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '', null, null, '0');
INSERT INTO `receive_btc` VALUES ('mfiWC2rQu4XxpuhKjPMwtsypMMWwcPWiAg', '10', null, null, '0');
INSERT INTO `receive_btc` VALUES ('mfktLqvV8ZYtzB7qCRuyaLnLFUpr9wUBfs', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mfziUUcGuZSXLomDRoicNEzciLGQQBTSpN', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mg6jW2aeQ6MTFYAHNAs1JkQ9hDrN7SoTHp', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mgraM26LBvsLi2wDmKD6agu6YH2hzKkHsm', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mgTfFYw567MrnhiatwVJTNuZYA6E3wLC1k', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mgVzvNeZz7BbX1CDbJBaGkgb9ZbdN8UfaV', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mgxr8txm6UDYgqTTN3xveyWKw2amCXe6J5', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', 'gm', null, null, '0');
INSERT INTO `receive_btc` VALUES ('mhDiw6s8ZtcepNyiqRdVXmEyW5CLbvrd8Y', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mhEsbiyM9UCFcXWLQR1ZpUxJWYc4c24kBE', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mhfCuUFWt5gxpFA9nqfk4zN2pZagrvpF6J', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mhobhjvUuNEtaKmBACQoyA8n4sHTCaiP9C', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mhUv4JVpBfNfkn8EVFDjvuSR8raLAmgU6L', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mhv3ckpQFf6vzY1EeJ2Ek2iX7Czt5bX9jh', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mhy8aZUiVTkQXVZf6HBSSoimstjCrmXLPG', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mhyBpRuyYUFR64wgC5QMUVkbg6UTUiofjG', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mhZTUbH5VSTdk6JFbonT1gswZFpniEjts2', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('miaqbjTKR4dWkAFWX4mYFhPfoV4prLWwSe', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mibfFSGXUo8WbqjatEBgFqCCuPS41CVktN', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mic9zdvh13mPWcPL8hPfYS6AL5hhtejT8w', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('migiPnh99UsfnvPuc5CJVpzQWhC6b57b8H', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('miLsEVUthZcpwavZeGRRpeMUEZDuZge2cL', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mimYuzFXBTZpnrkNETS3HQnmGVmDJL23ga', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('miWG5KjXUwoeyNkDm5KiwxhdvaWESjBeVB', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mjAACpntkARMB18WEWRUgPW91hEsu7NvXe', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mjDqQshx1HE9C2y3mWXAKSAXhnhh6sutx2', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mjGAKwu1CXXMsx4run8aPD1tA1PwQy2vrf', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mjmVW2dSE6Sz31TmJSayNnkgppYK78E9EQ', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mjQrvbcRjgALqKEjn3FpCADjdFnwCFD57V', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mjtnhWpkNYQZfEHo44266Vzd1Kami1NM7N', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mjTzoP7Yobx95Ygta3angy4XK1cUaLmndv', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mk5mpyLMaaxGjGLZbMo1h6UwMTGoFPk5cY', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mk9ZMmm8swFAqYiWgMxRPVoKf77J7GgJ5d', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mkbf5Miy19g1nKVjKMSN9moxqzAEmwChAs', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mkDRey5Kt928u6eno41yKnpxVTRQY6CnK1', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mkDzCchBCRj7CcdvuJZQv93L4bxqKM4y1g', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mkVvf3hFJdipwguTweV8EQWD3rkupBTDGu', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mkWR3cvh8bALQV935sWL6NV7JNr37Gm3Ur', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mkxKNyMLjEnN7HMUrYadWJ9USgLhgo4WQY', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mm3D9j9tdxzpXQDGM9j4av5Paer8pTsoKR', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mmaWnFkgqdoz3SPYzNBtqEtnhBzizppWV2', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mmEyyecPXviitWiv8LZUr6qht4udq5ZL5C', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mmi2cFHppkEc7EkDe8qCxdzreGNdxWQtaA', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mmNBPrukEY4fJxocwaK9G9XLZixPMnMCLP', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mmvdEiZSYZtS7RaX6hvHgHuCRFDSXCZHNr', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mnDVHLQtU6VTMu7tj6Xp3BZnae9USybaMJ', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mnuAouVSx3WZZiNRztVR68s9wcBBa3BihD', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mnUDGm9NoVEqi4YW5thpNLUWPZaEduHubL', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mnXQFaFcbbeLiHfhf7f8oc53AQKhFDbkGX', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mo7F2Puc5WFF3Y2DgGsVfU7AcqTxAHS8UM', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('moC8ZSWQbYE6eLvweddsNrW1KcC9W5P3Xy', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mojipom8kecgjjN4FLVLkNDChg6G4rV2FV', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mojuYQtMSBm4vCg5umKUFLWXeivvzccdoD', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('moKDPyhqqdDVi7bn7zRPCDY6E1mib7vVrt', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('moq7Aw6o1LbhAC2CwJdsejBLAiuqAG7wdj', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mosjakAXZmCPZ8pnQhgH4vvYUAPD522jfQ', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('moUBaJQufbKPtBKEDjFvgMKqLW5uemFGd1', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mox2wAJ6WiTc2hYTRpyvirJkMaPVVRVMUF', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mozjvU2KP3JVLR11mbm8pdNwFoG1fCeTV7', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mp9TFZmz7okhSvZwTEsrgjV5k46yrwFiWt', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mpBcYVDBqKCNEPmNAtEQ52RrvBTpqtudgp', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mpewhms9xdZkHu3KX995QS4YtrFfcR2nfa', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mpfjMzfAJ4fxi3giULTa6a7ke1u3VvVZy2', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqdBkQ7kCDffKfAXzPPXg4i3dYLD1PA7r2', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqfqpsS9afyD5MCy58iCjj3QA5Fm46hWtH', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqHxG7P6NLW8CgUeXEQYZczeq77Ts2e2P6', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqJGx73VpWMjoFmWY5xqkCHuSJS8v4LptF', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqjq1G7zM66yNcyDgFbCb58UZ7tuzCGf7x', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqKFQJgUEo9xrF29RAuYVq996ZtVHbUjBz', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqkKqbGPmEQKkNSrHQGzH6aJhAa74f4mFf', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqmSjByhAjEg8mNUnt24WUkfT3uVcp1LEY', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqMyzTLZYwNRwqBXxvqSHeWUGgpjq9hDYz', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqWg68EqoLcXtchkgrR9WhHjDZg7dANk7t', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqXdpMgWpnVY7GPtu8iU5SjPn8gkWvmBM2', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mqXyjC3R4S1RwZbkQHrsrvsiBmZF1gBw3T', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mr2Yt3eqfxbBX8GRTC8UvefUAHXiG3o7C1', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mr6fw3AzNei4AFfciYWrKgt8LMCL2NwzFC', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mr7LLf4DBmYKor2Jb6Sf1GgiKdBjftWf6r', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mr8kR89H58obxLc1AgcxikuY79dJYqKcGc', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mr9TJM9zXcEMDjJLDCCvufwFhJmEzpwhf4', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mrEjSG216roXmK9v4Z3j5DC2Qd8Yk4mkFT', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mrgGnH3MXVjXTE1zAexXXUr4SHvYb1Hzvv', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mrgJ7SXMc3KLHBtLyzKKx39tkFNqaLqyKF', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mriny4JWMdte2zVYupXCBdLGutAp4ECDHv', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mrjzfUqwudmhtCwW92AGf6NN3wCibRrhQ3', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mrnqNuA1eJEPV9KMrdXX5nWfiD6P3qid16', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mrSjo5mJ7eG3PCwCpQURggzcskH5mTJV9Y', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('ms1JaLNtxfN1is1ZjzaT3CjX6ScfWi2zJn', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('msbsp2MwwBwnsgkASNVHaaVJUNF96sruPm', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('msgmjTryu5MXDpk4oBYsyZiM8CUCKveQA2', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('msi5d2X8zpFfRLkE3nT1BDiWSev2y5J7ew', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('msK1rd85Ba8kteZKceCCaECTGAKLywWi5L', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('msQCPmW8FgNbdAugsPKis6sHHx4eiuxmfa', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('msqyWhPQC2a5RwtEvvRndx2xz8GqCWEEnL', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('msXttWL7g9ztqCNmQ6E2o7PdWLauzMXaj1', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('msZhgZm2P9atE2ZYjWPgsuyqa715LMDGaq', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mt1Wdgaeyd5SWThJS88ko8LDt3DpfsCH6K', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mtbHD24eWi9GnKXtqE3gAPP2nygNAJfTyK', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mtce1CDDis5FtWzAHMLegHDwgygeBWeK9R', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mtjW3F745mtW8ieZosWax3NeDvp3buUnbb', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mtQ29DPTzhPHV1jo7GCxcukDjvf2pwV7Xi', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mtweeSaDe9K2ogFvjBpcq2QjxBeuEtX9DZ', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mtWigdV8b56KfUMNRG8moQvXsSnQyZz3YQ', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mtx3jwmD1spcRmhZZc41oevVrVh7MSK594', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mu3zbx7Wpc3jv5bhf6GGhwzuArMxd6z8dS', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mu4uEXdBVYGk1hPJXMV5ArBhfaTXQ4s3wj', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mu5457bXqrdqiukPeDRHxdrUxG4odLPVq6', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mu9UwS5iMYXXRaVS8AG2cK7jGVaytMoKB5', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mudg7AqPm7yrndzuo8v2ZEWXqU3jJyZrgy', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mup53BAJj25YxK2SgvqEuh7rqNiMvTTo5f', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('muqfNY9fEVudSt7aAT265dErmRvYc6Jfa9', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('muttY1U7KqEv6AQokP3nJpSi7CKkgH8LPn', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('muxT3DJdM4Bs6mzJGpFpDXCfyjaYsNMPZq', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('muyrap5LRdbafV71dHSK1ZnL3QcJNpi11e', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('muYUurGcr3R5beiwsaMAEp8jhHMLpGmkFH', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mv1DgPYPQwHZnyWmddR5V2zwXxvt4TYkaZ', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mv68J8TXg3GATVUhU4bqn4bvtbE5myuDcy', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mvbCyB4B4XAUtkh8NnaUfysEs3LgopV4G7', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mvguQySV5AgAL4nfa5rizwZMKT1W2EWkdh', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mvPw7uK4BwA7NEk9WZXxB8rH3tXJGvqRmq', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mvqh75sQd3yKx8ZgrYQo6UHtN79ZHBtScv', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mvZUGXm31ankxqDRuM473edhkCZ1aMeEfL', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mw1KbD1jWCBPEzABipnmfAz35d214HbVEK', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mw567NAXNBBvovFdGQ5PAsuWSjPk4xUowX', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mwaHh48ZN2mvP6HboDSMD5wwNFsRuS31Ma', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mwdvi19ZKiADAf51cBr8ow2Uf6tAW6H6wT', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mwfukZ27mvfVq7fbdRruAkuvmrKcVRzW8M', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mwGSQzut8TwxLcQDmVvKavVcSerZXPUcak', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mwKfxzeUcSSsv25YErsAprStvCKdqvR3z1', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mwKNce39HZR45iQL4TDkEpCxmEJzfhDGEY', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mwsjQFwRxvS8H3P3nnUipVG1F5yFpLtFcj', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mxbeLNr6iSjmzewEJTNjuNJGQ2Umh62nju', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mxCeaiaiAAZpN3CFFM826tKc851McF52uK', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mxE7wnzVbWawQVLstg26RCaYW8QH9ENCe7', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mxeMJv9aCeFVHBg17EJ8ecQcx4CFmgNhUF', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mxkJFR6zQRnVWhZ2y5nZjRLw8ewJCcH6pD', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mxnHRT5GQgshmEa1QCtNfwWgKuGLmZHVoT', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mxr58HhfepG3ncFFPiskcknQxPS4Yk4Pxu', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mxtSqXzM2bwUe4gYAc1JiKcbA6SSgaVfEf', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mxVnsXpz75P5hTJ85dvUCnijnBs97GAb6W', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('my3YGRKinVPiAMziqyngh1G5VeZD6XXT3L', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('my6GXHpw3pv7gqjod5n2vPnA4A869oy9ZF', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('myELLFMLuNmUuB5VPmpZhesndtyBuX4qgn', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mygjSKRpSJmDK7FRnRdVywD1rJoqBQgs3H', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('myHKd9G8uJgR4UyfjYgjd7aKkpWMKjeTGu', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('myHwFnYouaYyYVsaWgN2rgmvfFHL8cL85K', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('myimJFYmBxLFX1f4vH1uNQCv9MDbaBhvsp', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('myM7Pe2ubesn7w59MtiJPkeupVUmDPtSf3', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('myQzwhVABWcC9Vn9U6B9yUdRoLW2atVoqF', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('myx8LaLdYx5mvzaNWB4EJ4So1EBLi6EREW', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('myXTL7H1XkYTRHLcfNXH8JSAnJ8Gavsj5C', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('myyoWH3Mum5bfceLephqy4pMsdNJkktbtC', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mz2GTpiMDbPfaBtn1T5PEq17tV84QpEKtD', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mzeJsRQoovLW5wgfJPdNbjeYKDYVtBX6EC', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mzFN795Z8Q25iKPrZxaZiBhdQnyBmXcNEE', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mzNv4veGiXMtUQPY3fXArHF9PjizA29LWC', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mzqkA9pA56AzEh3L2b47BRBxQgzyjx9HKE', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mzQQyBCfV3vB3ryqdjHBe7EEtm5UMaxczW', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('mzVquvRHrK9kWEkU3pHsQs3Htdm3mnPBrb', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n16P1MH8LGhrtfV84gaTbUPeRq7xazjyvK', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n177k2r77rYHqFTtKvDgNouiN8WqTx2Myz', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n1aCLPp5mC4ngSEF4TzFfcXjwu9BFr1Raq', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n1chVEUy8WGnLdTgEYsbdHxhsV4Mi4pEgq', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n1FXF45ycoG8BLrAgKe1YTX9TbrCMX9Zxp', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n1miDiVokSWJkkSTpACxSL5VTndqCw1Cu2', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n1rbBukMkPC78yMxi9FKSiFu6boFvLqafW', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n1TgXFKE2JUpFbc9EUREb74RVbbTWDnLLG', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n1TmpGFjdr8ysDVDcc4vRETzmGdeS4zkUa', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n2ekRhrCV6KoNuB5jpmzxMprMWSRUnXoUV', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n2GUVB6KSuLKrjDNu2FvJEirFTS9BZJ111', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n2jMiiaZvzbovRY3b4Hju4K3xZudnUgzPz', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n2NQm2ft8TCX68rP271nfzoSjDL1NpBsqK', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n2RJJWKnBxQA9VEsiVzFQyPMjbEwpy4aDu', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n2sYNxPSvSH1TzD9xkkqikaV6ueQi4bkS4', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n2zeQmEWU4fEJRa4P3BmPr7NntCEaGVNkr', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n35djLfMiC1nKco1XMfMPq4g6XfWpcVqSj', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n35zPCwYsY4KuUADBpXc8TGe7gbbDUrGJY', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n39xgYNThrrLwSw6SZhkgTFey4mza6RKdR', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n3Ctw4dN7CbtdLsAbrro57SEFcaTLRLnZV', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n3mDx6QYzq8qytQHsyVq4zFedSh7CejhvF', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n3P6QaAQSLPe5s4cYd4FahWTraK65j3V3Z', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n3tBmUAWTjLcW3H55fdvSg1sktiKgu8m4p', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n3vYdEMSwA4XKhqavUtxn679Q1Wxg2c4qs', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n3WLRZsLL9bLjfBRcjd3sFEvK2L5u1Brii', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n3yapRhxiYU3SuzFgVoYywWnJy9vdVfK6U', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n46wGn8hjUx5aFm4tywGgctddQvW5eSoHP', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4BaCwk9P9zuENF3BU6K8PnoiBPMWLeAu5', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4dKpo6vFz3jUbRHfXAmuhTEXqTVm1PZdA', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4dt97tB2rQFQQCKeMfZZGcpQYo5G4wo5w', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4EcQ8q1PmN85hGFpddJTYDieWRb65Go7C', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4erkPnK5Ymt3AcXKPGLwaPqmuPCMu596y', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4HwaVybtxwbk3f4GUFM8mZ262RotToy2Z', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4iLpwmNGdg8PFmbnX28JmpeyjTi9HBKHy', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4Mp2PrMGBstNe1qU1HWp97SpRJR9yPCPM', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4PV6CnLuXFNtrsSLMDJUzjBmz5iSYUU5o', null, null, null, '0');
INSERT INTO `receive_btc` VALUES ('n4TXR4sC6Pg7aY2bC7XdVCEVLF8LaKJhuC', null, null, null, '0');

-- ----------------------------
-- Table structure for `receive_eth`
-- ----------------------------
DROP TABLE IF EXISTS `receive_eth`;
CREATE TABLE `receive_eth` (
  `receive_address` varchar(128) NOT NULL,
  `receive_user_id` varchar(64) DEFAULT NULL,
  `receive_source_server` varchar(32) DEFAULT NULL,
  `receive_bind_date` datetime DEFAULT NULL,
  `receive_status` int(11) DEFAULT '0' COMMENT '0 为启用 1为禁用',
  PRIMARY KEY (`receive_address`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of receive_eth
-- ----------------------------
INSERT INTO `receive_eth` VALUES ('0x06556998aea32f3eb1642222abd80525a6ed0bc1', '10', null, null, '0');
INSERT INTO `receive_eth` VALUES ('0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', 'gm', null, null, '0');
INSERT INTO `receive_eth` VALUES ('0x1af20c47e7c3a99e338cae8edc16114eb7fba324', null, null, null, '0');
INSERT INTO `receive_eth` VALUES ('0x492c9df20dc557866c4519a39a17e7c3c440aa87', null, null, null, '0');
INSERT INTO `receive_eth` VALUES ('0x75306ecec3167eed8351afc0284412097fd3938c', null, null, null, '0');
INSERT INTO `receive_eth` VALUES ('0xbff0eca45ffe5de8d0a5a6af5b5305e258649909', null, null, null, '0');
INSERT INTO `receive_eth` VALUES ('0xd4a6474df59edeb03917d143e15760c6753123c3', null, null, null, '0');
INSERT INTO `receive_eth` VALUES ('0xdd24a02e3a2a1424fbbd772629875eb46f4449f1', null, null, null, '0');
INSERT INTO `receive_eth` VALUES ('0xf7b6fb469e865849ed82c0802ce8a5b1c4b24c6e', null, null, null, '0');
INSERT INTO `receive_eth` VALUES ('0xfb7d2cd03b61137d5a60276d1a0a7e67897e615e', null, null, null, '0');

-- ----------------------------
-- Table structure for `receive_logs`
-- ----------------------------
DROP TABLE IF EXISTS `receive_logs`;
CREATE TABLE `receive_logs` (
  `receive_hash` varchar(256) DEFAULT NULL,
  `receive_address` varchar(128) DEFAULT NULL,
  `receive_value` bigint(20) DEFAULT NULL,
  `receive_friendly` varchar(2048) DEFAULT NULL,
  `receive_create` datetime DEFAULT NULL COMMENT '区块出现时',
  `receive_block_height` int(11) DEFAULT NULL,
  `receive_status` int(11) DEFAULT '0' COMMENT '0 未确认， 1已确认，2无效',
  `receive_update` varchar(45) DEFAULT NULL COMMENT '区块最终改变时',
  `receive_check_num` int(11) DEFAULT '0' COMMENT '已经确认的区块',
  `currency_id` varchar(128) DEFAULT NULL COMMENT '收款的币种',
  `receive_id` varchar(64) NOT NULL,
  `receive_type` int(11) DEFAULT '0' COMMENT '0 比特币 1以太坊与代币',
  PRIMARY KEY (`receive_id`),
  UNIQUE KEY `receive_hash_UNIQUE` (`receive_hash`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of receive_logs
-- ----------------------------
INSERT INTO `receive_logs` VALUES ('fef6d5299d946d60210a19f264c03d0b6adbc54a5e78079cc4b2fefd51a237b4', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99999985', '0.99999985 BTC', '2018-05-12 15:38:01', '1297371', '1', '2018-05-12 15:38:01.217', '1297371', '100001', '007fb04eb920452ca243e821849e57bc', '0');
INSERT INTO `receive_logs` VALUES ('0x673770276333bf1ce2c97a03a3dd3490cfb48bc0aafee6d125e85f57979cfa2e', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '800000000', '0xe', '2018-05-15 20:55:16', '3240016', '0', null, '0', '100003', '00bf25c619f54c8ab2ba81c814f8aca3', '1');
INSERT INTO `receive_logs` VALUES ('57bd1f889b989db16ba23503ce616feba82771126c3cbb5d82ea93905bcd880d', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '9985', '0.00009985 BTC', '2018-05-12 15:09:37', '1297367', '0', null, '0', '0', '017c0625d3824714a50f4a18f4296565', '0');
INSERT INTO `receive_logs` VALUES ('0x974d2d6a9f34104f7f832aa301c82fc1f1fc024dcce789e8d4b42462e29318d3', '0xd4a6474df59edeb03917d143e15760c6753123c3', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:27:33', '3240131', '0', null, '0', '100002', '03f32100e65842968152881836bf2a87', '1');
INSERT INTO `receive_logs` VALUES ('0x302ce206c4f10a7cc9ad2a3bc67dd9995d4d9cf1d46f41daf10391e5a8b1f7a3', '0x3955ad7ff5dec0c83fe315b846b48afb8c8c7e38', '100000000', '0xde0b6b3a7640000', '2018-05-14 23:05:03', '3234161', '0', null, '0', '100002', '083abe8109ed460fb190a7c82d914881', '0');
INSERT INTO `receive_logs` VALUES ('0xf54fc90b168b89810a674aa963b1dce94452c295a6695706f8787c38d6707b1c', '0x06556998aea32f3eb1642222abd80525a6ed0bc1', '11', '0x3', '2018-05-16 17:08:55', '3244881', '0', null, '0', '100003', '089b2d6a55c544aabd5294d939b29af2', '1');
INSERT INTO `receive_logs` VALUES ('0x879246d6817e466a18c62e33d2f3997b88dc4c5876c751419570ea1717f2dd94', '0xe489aa96d8859639c53f6bf97548c740147e0503', '20000000', '0x2c68af0bb140000', '2018-05-14 18:08:29', '3232381', '0', null, '0', '100002', '1044a7a5b78e475782d972e53b60ecef', '0');
INSERT INTO `receive_logs` VALUES ('f9296ffe4419e3d145e5149c8b1f16578c1aadfbeda5c364b5b96bbd4fda9d44', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '999985', '0.00999985 BTC', '2018-05-12 15:33:20', '1297371', '1', '2018-05-12 15:33:19.772', '1297371', '100001', '106e40bdd7994a01bd718e7cd7742d60', '0');
INSERT INTO `receive_logs` VALUES ('0xf8699ce449f89aab70ce11d13301aa38b886db6a2bd7515c6ca5ec2530b53137', '0x3aa197163b476aa37d92ed69a4fdac232bab1766', '10000000', '0x9', '2018-05-15 11:54:23', '3237975', '0', null, '0', '100003', '1081ff2647504226acaae113d81bd213', '1');
INSERT INTO `receive_logs` VALUES ('0x11440ced0326bb19eb385acd07bea278f413b9bfa330b75d44742737d24adbf7', '0x0bb15077d1d2897c0f86432c2db0f5258e22cf13', '1000000000000000000', '0xde0b6b3a7640000', '2018-05-14 12:35:53', '3230387', '0', null, '0', '100002', '10a0867b8d024330aa1ded55f713ab7d', '0');
INSERT INTO `receive_logs` VALUES ('b7fac662e9e32b207f3947b4ad599e71546ce4878bc48447590ea70310297521', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '27500000', '0.275 BTC', '2018-05-12 11:56:34', '1297356', '0', null, '0', '0', '133fd0f45b5948b294a3fa8f1961c9f7', '0');
INSERT INTO `receive_logs` VALUES ('059ac9e5d087b7ccaa8767fe430b355a4f06e54acaf446f5512bd8ee718ccb7d', 'mg6fv8GVwGt85qKj9ft1uJxtpwf2Qc2sqw', '550000000', '5.50 BCH', '2018-05-22 11:19:41', '1234177', '0', null, '0', '100004', '14dfe26c2079498382eaabc1d1a82546', '0');
INSERT INTO `receive_logs` VALUES ('0x83dd2d9b880b7bee7ec553794fc9593641a6213e5fd509c12e281dc4899b7f35', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '400000000', '0xa', '2018-05-15 20:55:16', '3240016', '0', null, '0', '100003', '1689690a6ec945ff80d18c01230403e8', '1');
INSERT INTO `receive_logs` VALUES ('0x7f837329fb64ef9068ec58351e7323b70dc17bcc5e5bd960f055664bee78a7de', '0x3aa197163b476aa37d92ed69a4fdac232bab1766', '3334', '0x2', '2018-05-15 11:12:24', '3237801', '0', null, '0', '100003', '18060e343ec04721830df1657ded793d', '0');
INSERT INTO `receive_logs` VALUES ('0xa06810dbb42e22752e7f73d7538159eeba9a9092527f5ecc4c06390936ac39c3', '0x676af2bda0655eb89ccbd91d44d8ad254f019bda', '1000000', '0x2386f26fc10000', '2018-05-15 15:57:05', '3238852', '0', null, '0', '100002', '199ef1d56fc94f988772b898ec41b0ff', '1');
INSERT INTO `receive_logs` VALUES ('0x217ace7133b291d793a8c23cadc9880fcc5ee733136e808654e8822a6d7e2081', '0x310f4095b6491535132c83c7fc176be031426f4b', '1000000', '0x2386f26fc10000', '2018-05-15 15:57:05', '3238852', '0', null, '0', '100002', '19b54dd7cccb4de8a3543778fb110ace', '1');
INSERT INTO `receive_logs` VALUES ('0xf013b6c1693f84abd0d7d2e5ddc10ef8ffb0a307a92c419a56955103bce771e0', '0xfb7d2cd03b61137d5a60276d1a0a7e67897e615e', '47', '0x1', '2018-05-16 17:19:23', '3244923', '0', null, '0', '100003', '1bc04f716d3d49ceb92d7a69fe5b792c', '1');
INSERT INTO `receive_logs` VALUES ('0xa4d1c9634998892d30756471c13aef8a6a094c2f5045047edff1104c9f0eea2d', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '600000000', '0xc', '2018-05-15 20:55:16', '3240016', '0', null, '0', '100003', '1d52928ded874f41b19791561117ae18', '1');
INSERT INTO `receive_logs` VALUES ('0x58eb85f786ef8574d6cbd098419139f519a9a9f25ff8e75913c3761d275e212d', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '67900000000', '0xe', '2018-05-15 19:52:20', '3239812', '0', null, '0', '100003', '1fd1be0ff28445a8903318a6610e23a5', '1');
INSERT INTO `receive_logs` VALUES ('0xe55d6a82b99be150afa719a0fe6b9fd8936719fbc5404e7719f8f1f78b9efa30', '0xc6476508f0a167d88e321be49ec3c8d7e96badc7', '1000000', '0x2386f26fc10000', '2018-05-15 15:57:05', '3238852', '0', null, '0', '100002', '200d8b8f5b4345368aa363d59adc5158', '1');
INSERT INTO `receive_logs` VALUES ('f137c6faf7758643833dbb41a41300de6aa00f4fa20f23fc95d94e1457934284', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99985', '0.00099985 BTC', '2018-05-12 15:43:44', '1297371', '1', '2018-05-12 15:43:43.788', '1297371', '100001', '255d70c319f54253bdce96f3d8c12023', '0');
INSERT INTO `receive_logs` VALUES ('0x66931b56e34a6adb1eea0d242f4acdfc1296448a237bf3a00c3fb3ecb870f299', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '900000000', '0xf', '2018-05-15 20:55:16', '3240016', '0', null, '0', '100003', '26b1ee207b65473eb501aaf4361c3ea2', '1');
INSERT INTO `receive_logs` VALUES ('48ea1de485ca7cac4d04fffca99ce9a8ff864978e83937831ded8add284ffc2a', 'mfzVNoqaw37RpwuQcJBV7sPQCJPttRwamr', '550000000', '5.50 BCH', '2018-05-19 17:26:34', '1233077', '0', null, '0', '100004', '280c05e9bc4e455a9056212e16e6b440', '0');
INSERT INTO `receive_logs` VALUES ('0x3b8900ec5a36313b90fab86859b40eb6ac8cb5f1940e2aba55985aa2951da353', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '9800000000', '0x4', '2018-05-15 18:42:54', '3239608', '0', null, '0', '100003', '28a001199e7547bda0389510fc88648d', '1');
INSERT INTO `receive_logs` VALUES ('3785ae686843f5dba92aa0d04c320ae7fcd88e14f44ba41ef6bb38dfd99ca64f', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '55000000', '0.55 BTC', '2018-05-12 11:12:56', '1297354', '0', null, '0', '0', '299803333ead45b0896640ac9e586b69', '0');
INSERT INTO `receive_logs` VALUES ('0x0f90fda14152d7ddc1fc220d939e46ad032603db64b49e63accc6bb5bf418263', '0xda2adeeb957815d7fa702d6706a8cd9cfcf68ff2', '103474800000000', '0x2', '2018-05-15 21:49:27', '3240230', '0', null, '0', '100003', '2a8d7ae6b02348cda0c8fb29fb3b3e6b', '1');
INSERT INTO `receive_logs` VALUES ('68518d70a98eb541eddb7563d9128dd422980463ffa8335d4402c396ae30a950', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99992', '0.00099992 BTC', '2018-05-12 15:58:59', '1297372', '1', '2018-05-12 16:06:35.438', '1297373', '100001', '2e47d7c79c9747c1b432e8823ed1e46f', '0');
INSERT INTO `receive_logs` VALUES ('0x78156df091314a7836084006be6812c8e860d5ae513d7549e57be3cf8e817a19', '0xd4a6474df59edeb03917d143e15760c6753123c3', '1000000', '0x2386f26fc10000', '2018-05-15 16:14:04', '3238926', '0', null, '0', '100002', '2fc740e4475949afbb5c4bbd485fb8bc', '1');
INSERT INTO `receive_logs` VALUES ('a16e3ba610157bbcd52bbbf148b14ce5646e88d54b094000f4a0c7f008999882', 'mg6fv8GVwGt85qKj9ft1uJxtpwf2Qc2sqw', '100000000', '1.00 BCH', '2018-05-22 11:24:54', '1234177', '1', '2018-05-22 11:26:31.721', '1234178', '100004', '38708c9715b440b8b26abca6d573b8a6', '0');
INSERT INTO `receive_logs` VALUES ('0x2dcc9dff5818036af0d15a0f023d91fc94a497ef5b9d8f746eb280e453c01c41', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '200000000', '0x0', '2018-05-15 20:55:16', '3240015', '0', null, '0', '100003', '391f4c01242b4599b1e589d04c2564ac', '1');
INSERT INTO `receive_logs` VALUES ('24cbd8e2469dc2c179316b01df461e216f654752348f3a5a3f7333f9738e2d66', 'mgTfFYw567MrnhiatwVJTNuZYA6E3wLC1k', '9985', '0.00009985 BTC', '2018-05-11 22:41:05', '1297298', '0', null, '0', '0', '3a21c2a59a35421eaa8cd540ab433fcf', '0');
INSERT INTO `receive_logs` VALUES ('cca06e8cfc71e27db9218e8eda59151b43a0f440049365d2b1578b3bcf710506', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '13427', '0.00013427 BTC', '2018-05-11 19:25:12', '1297280', '0', null, '0', '0', '3c3ba9ea2e704a85ab6c8e8f0145250a', '0');
INSERT INTO `receive_logs` VALUES ('f967f47a796660be56a4976954a5af38a5c3afea019fa2190c31a87869c22d39', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '9985', '0.00009985 BTC', '2018-05-12 15:18:33', '1297370', '1', '2018-05-12 15:18:41.030', '1297370', '100001', '41b0ecc6a6524881adc0a3d7e2dfbbc0', '0');
INSERT INTO `receive_logs` VALUES ('0xc4f07e12371a60b5abb290650d82e2b6922a08baa5a834459c70b055e5148c54', '0xda2adeeb957815d7fa702d6706a8cd9cfcf68ff2', '519981566', '0x482974c9d5ca4600', '2018-05-15 22:07:39', '3240312', '0', null, '0', '100002', '455d5f4c8816470f861fc8e406ced296', '1');
INSERT INTO `receive_logs` VALUES ('156590724ac02ddcb9b50a6129bd61187c52aff280576f16c8ac067cdf7072a2', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99987', '0.00099987 BTC', '2018-05-12 15:57:52', '1297372', '1', '2018-05-12 16:06:32.695', '1297373', '100001', '4659614253b64720993e63cb5622a5fd', '0');
INSERT INTO `receive_logs` VALUES ('e156e5d768d74aecb25957004fb81322b9e40274cb70fc02020bbda29e686445', 'mg6fv8GVwGt85qKj9ft1uJxtpwf2Qc2sqw', '275000000', '2.75 BCH', '2018-05-22 11:24:48', '1234177', '1', '2018-05-22 11:26:35.699', '1234178', '100004', '48897d7cafad4c5b839f131734fa08d7', '0');
INSERT INTO `receive_logs` VALUES ('0x705abd40615e25ad650dda46a7f6fad87a340fab823c64a386c307e926c0fd0d', '0xfb7d2cd03b61137d5a60276d1a0a7e67897e615e', '4700000000', '0x4', '2018-05-16 18:32:40', '3245217', '0', null, '0', '100003', '4c9f41b7bfd5436db46cb71ae31111a8', '1');
INSERT INTO `receive_logs` VALUES ('0x02e374bcd029030be72c5f148333d16777aa66b6005eed1bb695280c6dffc98d', '0xdd24a02e3a2a1424fbbd772629875eb46f4449f1', '1000000', '0x2386f26fc10000', '2018-05-15 16:14:04', '3238926', '0', null, '0', '100002', '4d22f08072f744209be8819db4b4484d', '1');
INSERT INTO `receive_logs` VALUES ('0x553ab52be068c60cc3fcf824ddef89cc0b189623018194ba14db564e31144fa0', '0x1af20c47e7c3a99e338cae8edc16114eb7fba324', '1000000', '0x2386f26fc10000', '2018-05-15 16:14:04', '3238926', '0', null, '0', '100002', '4dc2a1d897f84d57a7d9929e65e79ad9', '1');
INSERT INTO `receive_logs` VALUES ('0x9f92cace54250af4b2b5df99bed1975673198c3e8f62119e1048dd382a5fb3ce', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '11100000000', '0x5', '2018-05-15 21:30:18', '3240151', '1', '2018-05-15 21:38:11.053', '3240171', '100003', '4dc35fca10b8481f97884f158262ddd4', '1');
INSERT INTO `receive_logs` VALUES ('e77cb67f78b50874a7472774a80d5a104705bbf7e69098eee2991f15d0078fd9', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '20000000', '0.20 BTC', '2018-05-12 14:39:33', '1297366', '1', '2018-05-12 14:39:50.487', '1297366', '0', '50cdce56f8a04edf82255d125ea6c3c1', '0');
INSERT INTO `receive_logs` VALUES ('02b8cc189708b96a19a25ce6772ea60836511340773e44d2aae7e7f0e2b8f49e', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99990', '0.0009999 BTC', '2018-05-12 15:58:27', '1297372', '1', '2018-05-12 16:06:39.323', '1297373', '100001', '5130fd2b988d400abe403de2862f674c', '0');
INSERT INTO `receive_logs` VALUES ('0x0d5ffbe572efcc6d51d97ca1263bbdb8b49af33a19ef4dcaaf77f8cfee971c7e', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '70000000', '0x9b6e64a8ec60000', '2018-05-15 21:52:42', '3240265', '1', '2018-05-15 22:07:02.450', '3240309', '100002', '52368f8ab8f84a97994c07d1abe78281', '1');
INSERT INTO `receive_logs` VALUES ('0x2501dbaca13d233e7504812ae68b0a8d891add2ba4df177907a6c4439197aee8', '0x492c9df20dc557866c4519a39a17e7c3c440aa87', '1000000', '0x2386f26fc10000', '2018-05-15 16:14:04', '3238926', '0', null, '0', '100002', '57ff09fb850946b7bd640b1be6e76b9d', '1');
INSERT INTO `receive_logs` VALUES ('0x7d91efd2879d39e5a1ef97d5b561125c0e9692e699d4a5459ce0e01be71b5fdf', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:27:33', '3240131', '0', null, '0', '100002', '65db2fdfe5ed460694753b9fbfe6fe52', '1');
INSERT INTO `receive_logs` VALUES ('d32d471ad9ff7886cdaa778c07d3c8ed6b407e9c3a3cdd95e6b705220c3e1335', 'mgTfFYw567MrnhiatwVJTNuZYA6E3wLC1k', '9988', '0.00009988 BTC', '2018-05-11 22:38:31', '1297298', '0', null, '0', '0', '6967a69bb396487e89225a6e23d54012', '0');
INSERT INTO `receive_logs` VALUES ('0x85129d9bc2e6727d9dc3284a65ed7e87fda188d7e525a876faf7f101e30e4daf', '0x06a861d752ef337ed12cfa8f0ac6b70f9047dae0', '3000000', '0x6a94d74f430000', '2018-05-14 15:54:18', '3231567', '0', null, '0', '100002', '6e28be76d19c4eaa8a78e6e9f884d949', '0');
INSERT INTO `receive_logs` VALUES ('0x7df124b75e80c74c73f0d7020a7bfd881c897b3a889a9c613964ee0ec7e29809', '0xdd24a02e3a2a1424fbbd772629875eb46f4449f1', '1', '0x0', '2018-05-16 11:34:33', '3243525', '0', null, '0', '100003', '6fbb70139d3e49fd942613b70056e9db', '1');
INSERT INTO `receive_logs` VALUES ('0x9a597fb8ece7cf0772c9fa1e2a955e58c8da6c6c54580c11e1de03fa95c18529', '0xdd24a02e3a2a1424fbbd772629875eb46f4449f1', '99999985', '0xde0b690bab1a400', '2018-05-16 11:07:33', '3243399', '0', null, '0', '100002', '73df90db5ac9496ba6b397fd9a67b7c0', '1');
INSERT INTO `receive_logs` VALUES ('0x4f2ca0d1abcde603d95994f73b99bf3abeddd46844a074c8d21bcf9cfb4fafe3', '0xdd24a02e3a2a1424fbbd772629875eb46f4449f1', '1800000000000', '0x10', '2018-05-16 16:42:10', '3244776', '0', null, '0', '100003', '75bdb375ba7348cabe6d16b6bb29e12d', '1');
INSERT INTO `receive_logs` VALUES ('0x63ea5de98667a6407e4a71eb7d1ca6247b5756cd8ca395d0a7e12a7d4549cd54', '0x5d59e1db9e2015684aa06faa9420c712071dfbbf', '100000000', '0xde0b6b3a7640000', '2018-05-14 23:05:48', '3234167', '0', null, '0', '100002', '76cd8433611c4eb9b7823c1e4d775eb0', '0');
INSERT INTO `receive_logs` VALUES ('0xd0335dbafd7a9d698ca457f765f941f7599f6d45f7652c6ad27b5a4c84f50002', '0x5c00e99c33a7d668923e60020fc8ce06c19f6c21', '1000000', '0x2386f26fc10000', '2018-05-15 15:57:05', '3238852', '0', null, '0', '100002', '7baa1f7c4910481da29317d5083aa9aa', '1');
INSERT INTO `receive_logs` VALUES ('0a01bb4ae780e77c4ec47ece50ce2fa25d4fce9039b4ac83b1335fb2b5823f07', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '9999985', '0.09999985 BTC', '2018-05-12 15:34:02', '1297371', '1', '2018-05-12 15:34:01.783', '1297371', '100001', '7c3766e70dee49a4b0e2ae82761a4972', '0');
INSERT INTO `receive_logs` VALUES ('0x115e69275132cd5201028f869b9a928dcf69450ec935ebb4b25bde3b76c41d8d', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '100000000', '0x3', '2018-05-15 21:13:20', '3240078', '1', '2018-05-15 21:18:36.396', '3240105', '100003', '7c48130c9da94662adc7dcf632203d03', '1');
INSERT INTO `receive_logs` VALUES ('4c0e4d8bfed7a6e9390add6e1c2a7667c12fc9d4ad74f9bc2322db1ab577469b', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '100000000', '1.00 BTC', '2018-05-11 19:54:42', '1297282', '0', null, '0', '0', '8182d48f8fe646429c4b3a3ae9ac4486', '0');
INSERT INTO `receive_logs` VALUES ('37783d660d72ec1322c012a6226535baf6ef1d99c5ebc795b11a0f9017a8263f', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '13750000', '0.1375 BTC', '2018-05-12 13:06:39', '1297360', '0', null, '0', '0', '82bfb977140143a8a0080aea96b8c419', '0');
INSERT INTO `receive_logs` VALUES ('0x0784f06f4e530828502329209b156e4198e2044b8eec076b90b6cd3c2e7ead2a', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '4000000000000', '0x8', '2018-05-15 17:47:00', '3239366', '0', null, '0', '100003', '8391adb9807e4959a9b1ebffdac0f237', '1');
INSERT INTO `receive_logs` VALUES ('0xbcb359c01771973d372131484a1f8c6812c6e11b76d0aa4347eddb5c827e801f', '0x3aa197163b476aa37d92ed69a4fdac232bab1766', '100000000', '0xde0b6b3a7640000', '2018-05-14 23:05:48', '3234168', '0', null, '0', '100002', '86c312598f8a427da0d5fa005f62612a', '0');
INSERT INTO `receive_logs` VALUES ('f77f6a54198b23acddadc6dd53521349e5518a2677c24f47d85e6672cdda4741', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '429687', '0.00429687 BTC', '2018-05-12 14:47:19', '1297366', '1', '2018-05-12 14:47:20.011', '1297366', '0', '87d4ad4af31148bda8bd803624ade475', '0');
INSERT INTO `receive_logs` VALUES ('0x4c9c58b1e4294e6de9aacc88fecdc6f5c9a383b0d3d6d524597810334dbfc3a9', '0x30b0ef89647e925fab572adb72b89f8b5baa2fa9', '100000000', '0xde0b6b3a7640000', '2018-05-14 23:06:48', '3234171', '0', null, '0', '100002', '88c72f9547e7457c85de66fb671688ab', '0');
INSERT INTO `receive_logs` VALUES ('0x58dc1b9433dc73cdc1e1de8d8ca1a14935c5ef9f7c4f06c0442c8394ee5700e8', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:44:52', '3240213', '1', '2018-05-15 21:58:23.454', '3240258', '100002', '8da6775ab3aa4f42b152b277e308fd47', '1');
INSERT INTO `receive_logs` VALUES ('0xa55b7f70c77442fe195b8e9189d5dc70f418865fab7f2c9b825846cd89311db4', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '50000000', '0x6f05b59d3b20000', '2018-05-15 21:47:07', '3240219', '1', '2018-05-15 21:58:32.270', '3240259', '100002', '9284dfd4d0fe4936bcde0fc882282681', '1');
INSERT INTO `receive_logs` VALUES ('0x1af85bd828f2459ccc31fea7753964c28763bb5f6dcbf559b5dc74c9aba25c85', '0x75306ecec3167eed8351afc0284412097fd3938c', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:27:33', '3240131', '0', null, '0', '100002', '96caa43c816546308f46ed4c93cbce85', '1');
INSERT INTO `receive_logs` VALUES ('78c8de6cf3a1f048e3c65f2c72c750aa98e95563a67dff888fc82550cb2b7233', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '6875000', '0.06875 BTC', '2018-05-12 14:17:57', '1297364', '0', null, '0', '0', '98a9557bc0ad4f3185f108f01521fc6f', '0');
INSERT INTO `receive_logs` VALUES ('0x4f9af0e5c91121f0054e5ea8cf6b3f8096a8a2b336c015f0045900e2d70680dc', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '99000000000000', '0x21', '2018-05-15 21:19:33', '3240112', '1', '2018-05-15 21:24:04.073', '3240125', '100003', '9c78d735b6f4456696cfa9fcdaeaf122', '1');
INSERT INTO `receive_logs` VALUES ('0xa0725fbfd1c8f109e3c20cc3148fedd77b64bda1cd677cee56cc3820ccbfc2db', '0x0c3b8cccc02119687744c05a21bc5224f824bc67', '2000000', '0x470de4df820000', '2018-05-14 13:51:28', '3230821', '0', null, '0', '100002', '9dd4f1f260594bf4bb06319bfa4de9b2', '0');
INSERT INTO `receive_logs` VALUES ('0xcf6383fa85e15d2df0582cc162924e7d91901237b05f9cb7e2a182de259d1a0a', '0xdd24a02e3a2a1424fbbd772629875eb46f4449f1', '498', '0x0', '2018-05-16 11:34:17', '3243500', '0', null, '0', '100003', '9e75424cefb24ee499ccbaa22f3b9921', '1');
INSERT INTO `receive_logs` VALUES ('7286cc9632ac4f7925246b68086756f6d11f7f07469c641c0bb1e67b6a45568e', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '107420', '0.0010742 BTC 0.0010742', '2018-05-11 19:17:11', '0', '0', null, '0', '0', 'a151fc45cfdb490a8f8d04dbefa83008', '0');
INSERT INTO `receive_logs` VALUES ('b5eea7cb64c93bfffad24ed2668534685ed453506d21e552c4d1bdb798d68f0f', 'mgTfFYw567MrnhiatwVJTNuZYA6E3wLC1k', '9988', '0.00009988 BTC', '2018-05-11 22:37:53', '1297298', '0', null, '0', '0', 'a42e8ec2e14f4922b8dadd5f83537044', '0');
INSERT INTO `receive_logs` VALUES ('0x171d8f1010cc2e35ef22f2ea2bc0f6ddb1ef03708e7f6709f9fd7cc037596f5a', '0xdd24a02e3a2a1424fbbd772629875eb46f4449f1', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:26:03', '3240130', '0', null, '0', '100002', 'a5f9ca00e29141eb81102de337422bd9', '1');
INSERT INTO `receive_logs` VALUES ('0x8ebd6c9cc57d6487196a00abc54e8f31d82139a6572e43cb96c3b6df9b4e4dfb', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '900000000', '0xb', '2018-05-15 17:54:22', '3239401', '0', null, '0', '100003', 'a682e902a6ab4b47bc54596f3e86b539', '1');
INSERT INTO `receive_logs` VALUES ('0x1104af85ac6d9f6ac0d59815bea564e602dab84e27885ed73f7f16365c9bcb32', '0x06e36dd21a8b0857288631bf9a9f0ba08398c177', '1000000', '0x2386f26fc10000', '2018-05-15 15:57:05', '3238852', '0', null, '0', '100002', 'a7567794bd7d45c4a318b0a49b725db8', '1');
INSERT INTO `receive_logs` VALUES ('0x93e7b5b510d86edad9512525b165ed2ef65778e398315f1df4501e2d417a1844', '0xbff0eca45ffe5de8d0a5a6af5b5305e258649909', '1000000', '0x2386f26fc10000', '2018-05-15 16:14:04', '3238926', '0', null, '0', '100002', 'a7c0dd85b27148e7abf43295bb528bd6', '1');
INSERT INTO `receive_logs` VALUES ('0xfe59264941140518d6613d022c98b9fecf8fda6679f8abc3e90a0cc159950a4a', '0xfb7d2cd03b61137d5a60276d1a0a7e67897e615e', '1000000', '0x2386f26fc10000', '2018-05-15 16:14:04', '3238926', '0', null, '0', '100002', 'a90fdfb533c24680bf230c3945627e96', '1');
INSERT INTO `receive_logs` VALUES ('0x5dcb3593c9eb110c3bef6e13b2d3bb1ed94154596bf87e5355eb34bf507e7ea3', '0x06a861d752ef337ed12cfa8f0ac6b70f9047dae0', '1000', '0x9184e72a000', '2018-05-14 15:17:16', '3231356', '0', null, '0', '100002', 'abfaa046dbdc4c02aadf55e20892792a', '0');
INSERT INTO `receive_logs` VALUES ('dea37ffa663ff99801c58955b1d2088c336defbd69842607b3b5bef5c6457882', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99991', '0.00099991 BTC', '2018-05-12 15:58:43', '1297372', '1', '2018-05-12 16:06:38.038', '1297373', '100001', 'ad9f6389bd504e5bb48328c9d14c0678', '0');
INSERT INTO `receive_logs` VALUES ('1799d72cc4a497d1b45b8b97c61608cd1ffae9ed9aae94845cec3d17be093387', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '107420', '0.0010742 BTC', '2018-05-12 15:55:22', '1297372', '1', '2018-05-12 16:06:08.861', '1297373', '100001', 'adc2b37777c241fcbc1925d2a223e650', '0');
INSERT INTO `receive_logs` VALUES ('0x21323fb3ac72d8d16a26b115c4dde12ec7e6a9f85fee48a109bfa0f57d815147', '0x0741b266f72c39011b430144dd196e14fdd42051', '30000000', '0x429d069189e0000', '2018-05-14 16:00:31', '3231601', '1', '2018-05-14 16:02:21.164', '3231614', '100002', 'b0fa73531a0945adb7dc3abfc188a9aa', '0');
INSERT INTO `receive_logs` VALUES ('0xbcb998f12f449e533d48435b30f414566a42525fb57ec66fe1adc56cd707053e', '0x06556998aea32f3eb1642222abd80525a6ed0bc1', '100000000', '0xde0b6b3a7640000', '2018-05-16 17:12:25', '3244892', '0', null, '0', '100002', 'b22c089db75f44a7ab28068ac824a31f', '1');
INSERT INTO `receive_logs` VALUES ('f269ad654b8b1b94656fee6916a20814962751172cc5f784bbf278845d588b0c', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '3437500', '0.034375 BTC', '2018-05-12 14:31:05', '1297365', '0', null, '0', '0', 'b3e06e3d7856482ca6479af032a215e7', '0');
INSERT INTO `receive_logs` VALUES ('0x885d9ad38d60b239cb8df11f09ddefcdc4ccaa8f327e54ce00ea2ce151a35622', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '52500000000', '0x4', '2018-05-15 21:42:37', '3240202', '1', '2018-05-15 21:49:07.750', '3240226', '100003', 'b493fdc7acda40129d67de6017b6d559', '1');
INSERT INTO `receive_logs` VALUES ('621704d0855e4da72f9b0785628a040d1a8d0e9316129e043e3f7456b4c84b24', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99985', '0.00099985 BTC', '2018-05-12 15:41:45', '1297371', '1', '2018-05-12 15:41:45.026', '1297371', '100001', 'b5e5b4cca2c14d2bb31c4b6c195a913b', '0');
INSERT INTO `receive_logs` VALUES ('0891af9ade6d2a2d767061869279de4459b5219c29ea83a104794bc7a54aeea2', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '3355', '0.00003355 BTC', '2018-05-11 19:27:55', '1297280', '0', null, '0', '0', 'b5f47c6969fe486aa3ea1774058bd32b', '0');
INSERT INTO `receive_logs` VALUES ('0x3df6a697cf7666251923d18372aa3f707ed706dea191fe4105ca789825f6e36c', '0x0c3b8cccc02119687744c05a21bc5224f824bc67', '8000000000', '0xb1a2bc2ec500000', '2018-05-14 13:41:27', '3230762', '0', null, '0', '100002', 'bc64e3cfc1244c40bf394c625ae390a4', '0');
INSERT INTO `receive_logs` VALUES ('0x6f6edbfe465e1e8b4a03102d4488879fff3191a6519e53a685c1735e18b5c4c3', '0x50e76ea65728dbe7616ca4ba72bb68265cb34040', '1000000', '0x2386f26fc10000', '2018-05-15 15:57:05', '3238852', '0', null, '0', '100002', 'bffe24f190394f6cb414deebd231c172', '1');
INSERT INTO `receive_logs` VALUES ('15258291011fd47dc2ea5aecc90103773b33f2484ba9aba9db697b144b8e21c6', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99986', '0.00099986 BTC', '2018-05-12 15:57:37', '1297372', '1', '2018-05-12 16:06:37.199', '1297373', '100001', 'c0e9da7df01f4d70ac518a3a286790db', '0');
INSERT INTO `receive_logs` VALUES ('0x3af256c7f7f0f5401f195c4923710daa3622bb61516dc69e1efc5e8ce21afaa5', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '1000000000', '0x10', '2018-05-15 20:55:16', '3240016', '0', null, '0', '100003', 'c1941bab693a49599acfc1ce06f3dc0b', '1');
INSERT INTO `receive_logs` VALUES ('137cb109f521095fb628fb5ed8aedf74a7bf7581d9a4ee07fe70517ea5777f91', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '9985', '0.00009985 BTC', '2018-05-12 15:14:45', '1297368', '0', null, '0', '0', 'c1d3056c55624c98b5bbcde1bc235ee1', '0');
INSERT INTO `receive_logs` VALUES ('e62a77f526aac1332a18d8719ed5a97b84cdd68173f85ee3685cfeafd7ef893a', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99988', '0.00099988 BTC', '2018-05-12 15:58:07', '1297372', '1', '2018-05-12 16:06:21.341', '1297373', '100001', 'c1f092bd5ad04235b31ff2ee94214235', '0');
INSERT INTO `receive_logs` VALUES ('0x3b21a1d9876a6405f5b3fed2d62b7ef72b6fc4bf964e46d997f5747e2ca77857', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '500000000', '0xb', '2018-05-15 20:55:16', '3240016', '0', null, '0', '100003', 'c492e92ae51f403c93c7e3a5c6cf1ff5', '1');
INSERT INTO `receive_logs` VALUES ('0x01824cc12f6de7e7806377254ef79d17ce291f0a736ab230f9b19bb61c4b5d9a', '0x9f3cdfa973c442c82859ded8d90eae7434047347', '1000000', '0x2386f26fc10000', '2018-05-15 15:57:05', '3238852', '0', null, '0', '100002', 'c5894239664149b5ae1fd4f4e7db2cb2', '1');
INSERT INTO `receive_logs` VALUES ('0x03962a2285d669e790220aac3b161b6316dbd4acab1f8bd1f4108138838c3e97', '0xdd24a02e3a2a1424fbbd772629875eb46f4449f1', '498', '0x3', '2018-05-16 11:11:18', '3243417', '0', null, '0', '100003', 'c6068f1db9434dcb9dee9b1a1aa252fc', '1');
INSERT INTO `receive_logs` VALUES ('0x09c566aff40569de344262a94ac5c9cd4b649beaedd8396455004a5c2ee34e8f', '0x1af20c47e7c3a99e338cae8edc16114eb7fba324', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:27:33', '3240131', '0', null, '0', '100002', 'c63221d34084497bbe73fae9d63887c8', '1');
INSERT INTO `receive_logs` VALUES ('0xc5d217f2500717c8c53506a4f668461963573e63ab38aeb8a1aa7932ccd00df5', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:43:22', '3240205', '0', null, '0', '100002', 'c6cb7650f0894d98b59ca7bd96095c12', '1');
INSERT INTO `receive_logs` VALUES ('5b7148a10ae31cba63dd67b0c2a396c0ac483490eb80314cfd797921707e91d6', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '214842', '0.00214842 BTC', '2018-05-12 14:53:07', '1297367', '1', '2018-05-12 14:53:22.558', '1297367', '0', 'cd55b982ed3e4bea9689c91253bb889d', '0');
INSERT INTO `receive_logs` VALUES ('0x7f46c690ab7b8cbaedd6bcf05ecba2b4900b0715ab07d27505d2a79ed07760ee', '0x48aca8ecfcd7db272bf754b39f652714ea734a5d', '1000000000000000000', '0xde0b6b3a7640000', '2018-05-14 12:32:11', '3230352', '0', null, '0', '100002', 'd214b41a8f4f46659008d24d89cd4d97', '0');
INSERT INTO `receive_logs` VALUES ('da2a2ca7cf4c2fca420ec5eef1879984e871d050668608495f1a7fee75afe085', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99985', '0.00099985 BTC', '2018-05-12 15:31:19', '1297371', '1', '2018-05-12 15:31:23.784', '1297371', '100001', 'd2cb4bcb40ed46deb6d6b098aab59112', '0');
INSERT INTO `receive_logs` VALUES ('0xd49d8ad3d79f04a1cfe1dbb2d1958073b6d9ab161f0f78bbd78c34a2ac74c4ce', '0x0741b266f72c39011b430144dd196e14fdd42051', '980000', '0x22d10c4ecc8000', '2018-05-14 18:42:23', '3232592', '0', null, '0', '100002', 'd74bff8760d14001ba711f54fe4ddf37', '0');
INSERT INTO `receive_logs` VALUES ('0x54e5dfff32ad9ef18af34d1b121edc6b43ffb7ae3e9c97071209ce38934160c6', '0x0741b266f72c39011b430144dd196e14fdd42051', '100000000', '0xde0b6b3a7640000', '2018-05-14 15:00:29', '3231278', '1', '2018-05-14 15:01:51.620', '3231291', '100002', 'd867313c8def427c9d866e7cddc7915d', '0');
INSERT INTO `receive_logs` VALUES ('0xe0fa3014f9a482eed209b325fb2b1328b88edc88561ebfea102c46787e9fee7c', '0xfb7d2cd03b61137d5a60276d1a0a7e67897e615e', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:27:33', '3240131', '0', null, '0', '100002', 'da7141d20e474ae389e116012d5efd60', '1');
INSERT INTO `receive_logs` VALUES ('852e232abe652e4f9ae715a36f6f2b09f078ba61842eef798086d5ac1856867c', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '26855', '0.00026855 BTC', '2018-05-11 19:24:22', '1297280', '0', null, '0', '0', 'dbf34e2f3ecf4ffeaced0372952eaf0f', '0');
INSERT INTO `receive_logs` VALUES ('0x12c077af95958d62be1fc788df6d2f46aa6081980d6cf4ce5feea4fdaf759776', '0x06556998aea32f3eb1642222abd80525a6ed0bc1', '1000000', '0x2386f26fc10000', '2018-05-15 16:14:04', '3238926', '0', null, '0', '100002', 'dca651886ad14dce8f05f935b110ceb4', '1');
INSERT INTO `receive_logs` VALUES ('0x9438aa2cc2701bac90aa282a18492817b8ddf8a3b2af14c779bc94ed0e8d1db7', '0xfaed8b092dc81780dc7e1f5c0872feb397f2d216', '1000000', '0x2386f26fc10000', '2018-05-15 15:57:05', '3238852', '0', null, '0', '100002', 'df3f3cc31e334115b7f29875e5259a12', '1');
INSERT INTO `receive_logs` VALUES ('0xb8511319124ad117c72d02ab8d233db1de8d22613c4868fd057b1a7356dbb7f7', '0x3aa197163b476aa37d92ed69a4fdac232bab1766', '1', '0x2', '2018-05-15 12:39:00', '3238144', '1', '2018-05-15 12:43:34.065', '3238157', '100003', 'e00744b909c74821a6e6f8bf20f31455', '1');
INSERT INTO `receive_logs` VALUES ('2e6541d6293b32f59eb78b6a276ad728dbcb83f30a8e8b64cc0935a2470e03a0', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '6713', '0.00006713 BTC', '2018-05-11 19:26:28', '1297280', '0', null, '0', '0', 'e18c9dd962ae4c579cae9c8dc2af1d7b', '0');
INSERT INTO `receive_logs` VALUES ('90cda1ae23abff49525ff24ad07492f7a361698f358a60366b2c48228a58b69d', 'mgTfFYw567MrnhiatwVJTNuZYA6E3wLC1k', '9986', '0.00009986 BTC', '2018-05-11 22:40:06', '1297298', '0', null, '0', '0', 'e5ae2bef1de74a579b3d1898c4a85476', '0');
INSERT INTO `receive_logs` VALUES ('0x81b0373e628034a04218980675b96260e5bd917a1c7ab05cf7118f0326fc5cee', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:32:18', '3240172', '0', null, '0', '100002', 'e706eb85e3cb453f85fe4e32a4b7ba67', '1');
INSERT INTO `receive_logs` VALUES ('e54ba40e18adcee1e9d7b304450a89f100076b029b6cb141ecdf4d25b0914ac2', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '109967', '0.00109967 BTC', '2018-05-12 16:12:23', '1297373', '1', '2018-05-12 16:32:19.981', '1297375', '100001', 'e7c185f82b84413ebdafaf07c219f839', '0');
INSERT INTO `receive_logs` VALUES ('0x6757c71009bfb6f828bcd8f0cb0c6372d6eb833178cda25fcb0ddec30c21c313', '0xbff0eca45ffe5de8d0a5a6af5b5305e258649909', '100000000', '0xde0b6b3a7640000', '2018-05-15 21:27:33', '3240131', '0', null, '0', '100002', 'eaaef8d186da41608d4dd778576f47ed', '1');
INSERT INTO `receive_logs` VALUES ('0xd7bdbab8aa72df5b919db703507f1b9573f884c29171b3a0a2f10c7349bb59c2', '0xfb7d2cd03b61137d5a60276d1a0a7e67897e615e', '470000000000', '0x0', '2018-05-16 18:34:25', '3245225', '0', null, '0', '100003', 'ef49bfb12b83462cab701fba2ab8ffbb', '1');
INSERT INTO `receive_logs` VALUES ('0x985b80de70e559fe03d06101eeb39e1dabd16898646ea059c6943ae361b5b345', '0xc3f0fe44e968e3aee643eecec59a296028355308', '100000000', '0xde0b6b3a7640000', '2018-05-14 23:05:18', '3234164', '0', null, '0', '100002', 'f0290461b0a44c36bb8872fe62521911', '0');
INSERT INTO `receive_logs` VALUES ('9d122ea107419cc48b99f519c11a47c174f299dd9be1b8315b2e4b8ae0f984ed', 'mfef2yZurqi11URC9z71YbvXNy6S5o1nft', '53710', '0.0005371 BTC', '2018-05-11 19:21:53', '1297280', '0', null, '0', '0', 'f20849572ed544feb219af58fa9d333b', '0');
INSERT INTO `receive_logs` VALUES ('6c2a27d9a096cb4669d76c387d0a155bcf7d965c44218775536aeb3db588b0f6', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99985', '0.00099985 BTC', '2018-05-12 15:39:55', '1297371', '1', '2018-05-12 15:39:55.158', '1297371', '100001', 'f269595d25f34248a0b2bcbf6a65353d', '0');
INSERT INTO `receive_logs` VALUES ('0xb18a5a58a568928047a35c61c492ef26be81802d7165f142282118d49d682377', '0xf7b6fb469e865849ed82c0802ce8a5b1c4b24c6e', '1000000', '0x2386f26fc10000', '2018-05-15 16:14:04', '3238926', '0', null, '0', '100002', 'f4ea6edbff324583b848e889b2a00097', '1');
INSERT INTO `receive_logs` VALUES ('0x7072ad2f937d58919fb65c601dc76bb88d72f77dccb94e31c46828cfe44a7ad5', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '300000000', '0x2', '2018-05-15 20:55:16', '3240016', '0', null, '0', '100003', 'f63b1e795ba740e98eec91c01d3a33e4', '1');
INSERT INTO `receive_logs` VALUES ('80045087799091737c126bd9550e86fdc06fdef6092dfcb610ddd7fc8e52cec4', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99985', '0.00099985 BTC', '2018-05-12 15:47:04', '1297372', '1', '2018-05-12 16:06:38.686', '1297373', '100001', 'f96fb8013fee49fe980f72484c70f017', '0');
INSERT INTO `receive_logs` VALUES ('c6f4e7e38887c99138f4bb76ebe04927393ae52189e30626abcd92e8257f1b40', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', '99989', '0.00099989 BTC', '2018-05-12 15:58:17', '1297372', '1', '2018-05-12 16:06:25.775', '1297373', '100001', 'f9c14e3cc2f840ad8efdd4cbb2c9cc18', '0');
INSERT INTO `receive_logs` VALUES ('0x002cd5d2603cc1aacb321b4fef51917fcbce967430ed751cdff87a9d6dffcfb9', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '700000000', '0xd', '2018-05-15 20:55:16', '3240016', '0', null, '0', '100003', 'fb22bd8553cc49c1872257f0cb68099f', '1');
INSERT INTO `receive_logs` VALUES ('0x1aad022ff66a68a85b19b625f0f94c4ad2cb78f61139c4e72655da2faea20d0c', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', '100000000', '0x7', '2018-05-15 20:55:01', '3240014', '0', null, '0', '100003', 'febe9130458c4fe68c0302bc808ee06b', '1');

-- ----------------------------
-- Table structure for `stage_coin_info`
-- ----------------------------
DROP TABLE IF EXISTS `stage_coin_info`;
CREATE TABLE `stage_coin_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `stage_id` bigint(20) DEFAULT NULL COMMENT '关联阶段id',
  `coin_id` varchar(20) DEFAULT NULL COMMENT '币种id',
  `complete_value` bigint(20) DEFAULT NULL COMMENT '完成金额',
  `min_value` bigint(20) DEFAULT NULL,
  `max_value` bigint(20) DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `stage_id` (`stage_id`),
  CONSTRAINT `stage_coin_info_ibfk_1` FOREIGN KEY (`stage_id`) REFERENCES `stage_info` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of stage_coin_info
-- ----------------------------
INSERT INTO `stage_coin_info` VALUES ('1', '1', '10001', '0', '5', '10', '100', '2018-05-26 15:59:49', null);
INSERT INTO `stage_coin_info` VALUES ('2', '1', '10002', '0', '10', '20', '50', '2018-05-26 16:00:13', null);
INSERT INTO `stage_coin_info` VALUES ('3', '2', '10001', '0', '3', '5', '50', '2018-05-26 16:00:43', null);

-- ----------------------------
-- Table structure for `stage_info`
-- ----------------------------
DROP TABLE IF EXISTS `stage_info`;
CREATE TABLE `stage_info` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `project_id` bigint(20) DEFAULT NULL COMMENT '项目Id',
  `stage_number` int(11) DEFAULT NULL COMMENT '项目期数信息',
  `stage_name` text COMMENT '阶段名称',
  `coin_count` int(11) DEFAULT NULL COMMENT '募资币种',
  `target_value` bigint(20) DEFAULT NULL COMMENT '募集资金',
  `soft_value` bigint(20) DEFAULT NULL COMMENT '软顶',
  `discount` float DEFAULT NULL COMMENT '折扣',
  `complete_value` bigint(20) DEFAULT NULL COMMENT '当前完成金额',
  `stage_status` int(11) DEFAULT NULL COMMENT '分期状态',
  `begin_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `end_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `project_id` (`project_id`),
  CONSTRAINT `stage_info_ibfk_1` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of stage_info
-- ----------------------------
INSERT INTO `stage_info` VALUES ('1', '1', '1', '天使轮', '2', '3000', '2500', '0.8', '0', '0', '2018-05-26 15:56:45', '2018-05-26 15:56:45', '2018-05-26 15:56:45', '2018-05-26 15:56:45');
INSERT INTO `stage_info` VALUES ('2', '1', '2', 'A轮', '1', '2000', '1800', '0.9', '0', '0', '2018-05-26 15:57:31', '2018-05-26 15:57:31', '2018-05-26 15:57:31', '2018-05-26 15:57:31');

-- ----------------------------
-- Table structure for `sys_log`
-- ----------------------------
DROP TABLE IF EXISTS `sys_log`;
CREATE TABLE `sys_log` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `UserId` varchar(50) DEFAULT NULL,
  `UserName` varchar(50) DEFAULT NULL,
  `Type` int(11) NOT NULL,
  `IP` varchar(50) DEFAULT NULL,
  `IPAddress` varchar(50) DEFAULT NULL,
  `ModuleId` varchar(50) DEFAULT NULL,
  `ModuleName` varchar(50) DEFAULT NULL,
  `Result` bit(1) DEFAULT NULL,
  `Description` varchar(500) DEFAULT NULL,
  `CreationTime` datetime NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB AUTO_INCREMENT=155 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_log
-- ----------------------------
INSERT INTO `sys_log` VALUES ('3', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-17 16:55:03');
INSERT INTO `sys_log` VALUES ('4', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-17 16:56:31');
INSERT INTO `sys_log` VALUES ('5', '975247111765495808', '超级管理员', '6', null, null, null, 'User', '', '用户[975247111765495808]修改用户[975308390198808576]状态为：Disabled', '2018-05-17 16:56:47');
INSERT INTO `sys_log` VALUES ('6', '975247111765495808', '超级管理员', '6', null, null, null, 'User', '', '用户[975247111765495808]修改用户[975308390198808576]状态为：Normal', '2018-05-17 16:56:50');
INSERT INTO `sys_log` VALUES ('7', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-17 17:00:18');
INSERT INTO `sys_log` VALUES ('8', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-17 17:01:05');
INSERT INTO `sys_log` VALUES ('9', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-17 17:02:15');
INSERT INTO `sys_log` VALUES ('10', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-17 17:05:50');
INSERT INTO `sys_log` VALUES ('11', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-17 17:20:55');
INSERT INTO `sys_log` VALUES ('12', '975247111765495808', '超级管理员', '6', null, null, null, 'User', '', '用户[975247111765495808]修改用户[975308390198808576]状态为：Closed', '2018-05-17 17:26:48');
INSERT INTO `sys_log` VALUES ('13', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-17 20:45:26');
INSERT INTO `sys_log` VALUES ('14', '975247111765495808', '超级管理员', '6', null, null, null, 'User', '', '用户[975247111765495808]修改用户[997045295579795456]状态为：Disabled', '2018-05-17 20:48:18');
INSERT INTO `sys_log` VALUES ('15', '975247111765495808', '超级管理员', '6', null, null, null, 'User', '', '用户[975247111765495808]修改用户[997045295579795456]状态为：Normal', '2018-05-17 20:48:21');
INSERT INTO `sys_log` VALUES ('16', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 11:59:21');
INSERT INTO `sys_log` VALUES ('17', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 12:12:05');
INSERT INTO `sys_log` VALUES ('18', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 12:14:52');
INSERT INTO `sys_log` VALUES ('19', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 12:22:18');
INSERT INTO `sys_log` VALUES ('20', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 12:47:02');
INSERT INTO `sys_log` VALUES ('21', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 12:48:15');
INSERT INTO `sys_log` VALUES ('22', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 12:53:38');
INSERT INTO `sys_log` VALUES ('23', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 12:57:47');
INSERT INTO `sys_log` VALUES ('24', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 13:01:00');
INSERT INTO `sys_log` VALUES ('25', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 13:05:38');
INSERT INTO `sys_log` VALUES ('26', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 13:33:44');
INSERT INTO `sys_log` VALUES ('27', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 15:24:10');
INSERT INTO `sys_log` VALUES ('28', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 15:49:45');
INSERT INTO `sys_log` VALUES ('29', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 15:54:21');
INSERT INTO `sys_log` VALUES ('30', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 16:11:43');
INSERT INTO `sys_log` VALUES ('31', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 16:19:39');
INSERT INTO `sys_log` VALUES ('32', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 16:31:48');
INSERT INTO `sys_log` VALUES ('33', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-18 19:54:16');
INSERT INTO `sys_log` VALUES ('34', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 10:33:38');
INSERT INTO `sys_log` VALUES ('35', '975247111765495808', '超级管理员', '6', null, null, null, 'User', '', '用户[975247111765495808]修改用户[997045295579795456]状态为：Disabled', '2018-05-21 10:35:35');
INSERT INTO `sys_log` VALUES ('36', '975247111765495808', '超级管理员', '6', null, null, null, 'User', '', '用户[975247111765495808]修改用户[997045295579795456]状态为：Normal', '2018-05-21 10:35:50');
INSERT INTO `sys_log` VALUES ('37', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 10:39:21');
INSERT INTO `sys_log` VALUES ('38', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 10:57:53');
INSERT INTO `sys_log` VALUES ('39', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 11:20:32');
INSERT INTO `sys_log` VALUES ('40', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 11:22:18');
INSERT INTO `sys_log` VALUES ('41', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 11:33:11');
INSERT INTO `sys_log` VALUES ('42', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 12:01:39');
INSERT INTO `sys_log` VALUES ('43', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 12:09:03');
INSERT INTO `sys_log` VALUES ('44', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 13:41:41');
INSERT INTO `sys_log` VALUES ('45', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 13:44:47');
INSERT INTO `sys_log` VALUES ('46', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 13:54:59');
INSERT INTO `sys_log` VALUES ('47', null, null, '1', '127.0.0.1', null, null, '系统登录', '', '用户[admin]登录失败：密码不正确，请重新输入', '2018-05-21 15:12:28');
INSERT INTO `sys_log` VALUES ('48', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 15:12:36');
INSERT INTO `sys_log` VALUES ('49', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 15:33:04');
INSERT INTO `sys_log` VALUES ('50', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 16:30:39');
INSERT INTO `sys_log` VALUES ('51', null, null, '1', '127.0.0.1', null, null, '系统登录', '', '用户[admin]登录失败：密码不正确，请重新输入', '2018-05-21 17:11:38');
INSERT INTO `sys_log` VALUES ('52', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 17:11:42');
INSERT INTO `sys_log` VALUES ('53', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 17:19:42');
INSERT INTO `sys_log` VALUES ('54', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 17:29:19');
INSERT INTO `sys_log` VALUES ('55', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 17:35:12');
INSERT INTO `sys_log` VALUES ('56', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 17:38:11');
INSERT INTO `sys_log` VALUES ('57', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 17:41:44');
INSERT INTO `sys_log` VALUES ('58', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 17:45:02');
INSERT INTO `sys_log` VALUES ('59', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 17:59:26');
INSERT INTO `sys_log` VALUES ('60', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 18:03:38');
INSERT INTO `sys_log` VALUES ('61', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 18:10:54');
INSERT INTO `sys_log` VALUES ('62', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 18:24:59');
INSERT INTO `sys_log` VALUES ('63', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 18:30:21');
INSERT INTO `sys_log` VALUES ('64', null, null, '1', '127.0.0.1', null, null, '系统登录', '', '用户[admin]登录失败：密码不正确，请重新输入', '2018-05-21 18:44:16');
INSERT INTO `sys_log` VALUES ('65', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 18:44:20');
INSERT INTO `sys_log` VALUES ('66', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 18:48:03');
INSERT INTO `sys_log` VALUES ('67', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 18:53:18');
INSERT INTO `sys_log` VALUES ('68', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 19:18:56');
INSERT INTO `sys_log` VALUES ('69', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 19:28:03');
INSERT INTO `sys_log` VALUES ('70', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 20:01:08');
INSERT INTO `sys_log` VALUES ('71', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 20:45:21');
INSERT INTO `sys_log` VALUES ('72', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 21:08:32');
INSERT INTO `sys_log` VALUES ('73', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 21:17:55');
INSERT INTO `sys_log` VALUES ('74', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-21 21:36:14');
INSERT INTO `sys_log` VALUES ('75', null, null, '1', '127.0.0.1', null, null, '系统登录', '', '用户[admin]登录失败：密码不正确，请重新输入', '2018-05-22 10:05:59');
INSERT INTO `sys_log` VALUES ('76', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-22 10:06:03');
INSERT INTO `sys_log` VALUES ('77', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-22 10:51:41');
INSERT INTO `sys_log` VALUES ('78', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-22 10:53:42');
INSERT INTO `sys_log` VALUES ('79', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 11:07:46');
INSERT INTO `sys_log` VALUES ('80', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 11:11:27');
INSERT INTO `sys_log` VALUES ('81', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 11:17:16');
INSERT INTO `sys_log` VALUES ('82', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 12:13:54');
INSERT INTO `sys_log` VALUES ('83', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 12:23:00');
INSERT INTO `sys_log` VALUES ('84', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 12:30:58');
INSERT INTO `sys_log` VALUES ('85', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 12:36:15');
INSERT INTO `sys_log` VALUES ('86', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 12:42:13');
INSERT INTO `sys_log` VALUES ('87', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 12:43:30');
INSERT INTO `sys_log` VALUES ('88', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 13:17:48');
INSERT INTO `sys_log` VALUES ('89', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 13:23:49');
INSERT INTO `sys_log` VALUES ('90', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 14:14:57');
INSERT INTO `sys_log` VALUES ('91', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 14:55:50');
INSERT INTO `sys_log` VALUES ('92', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 15:05:39');
INSERT INTO `sys_log` VALUES ('93', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 15:11:02');
INSERT INTO `sys_log` VALUES ('94', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 15:14:20');
INSERT INTO `sys_log` VALUES ('95', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 16:16:14');
INSERT INTO `sys_log` VALUES ('96', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 16:47:22');
INSERT INTO `sys_log` VALUES ('97', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 17:55:21');
INSERT INTO `sys_log` VALUES ('98', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 18:35:39');
INSERT INTO `sys_log` VALUES ('99', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 18:46:22');
INSERT INTO `sys_log` VALUES ('100', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 19:07:42');
INSERT INTO `sys_log` VALUES ('101', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 20:09:46');
INSERT INTO `sys_log` VALUES ('102', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 20:19:27');
INSERT INTO `sys_log` VALUES ('103', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 20:27:00');
INSERT INTO `sys_log` VALUES ('104', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 20:29:25');
INSERT INTO `sys_log` VALUES ('105', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 21:36:39');
INSERT INTO `sys_log` VALUES ('106', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 21:42:52');
INSERT INTO `sys_log` VALUES ('107', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 21:45:30');
INSERT INTO `sys_log` VALUES ('108', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 21:57:42');
INSERT INTO `sys_log` VALUES ('109', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 22:12:47');
INSERT INTO `sys_log` VALUES ('110', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 22:17:59');
INSERT INTO `sys_log` VALUES ('111', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 22:21:40');
INSERT INTO `sys_log` VALUES ('112', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-24 22:27:06');
INSERT INTO `sys_log` VALUES ('113', null, null, '1', '127.0.0.1', null, null, '系统登录', '', '用户[admin]登录失败：密码不正确，请重新输入', '2018-05-25 11:49:10');
INSERT INTO `sys_log` VALUES ('114', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 11:49:15');
INSERT INTO `sys_log` VALUES ('115', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 12:26:52');
INSERT INTO `sys_log` VALUES ('116', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 13:32:48');
INSERT INTO `sys_log` VALUES ('117', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 13:43:22');
INSERT INTO `sys_log` VALUES ('118', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 13:45:43');
INSERT INTO `sys_log` VALUES ('119', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 13:48:29');
INSERT INTO `sys_log` VALUES ('120', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 14:20:14');
INSERT INTO `sys_log` VALUES ('121', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 14:46:20');
INSERT INTO `sys_log` VALUES ('122', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 14:47:44');
INSERT INTO `sys_log` VALUES ('123', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 14:51:04');
INSERT INTO `sys_log` VALUES ('124', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 14:52:44');
INSERT INTO `sys_log` VALUES ('125', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 15:06:01');
INSERT INTO `sys_log` VALUES ('126', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 15:13:15');
INSERT INTO `sys_log` VALUES ('127', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 15:18:00');
INSERT INTO `sys_log` VALUES ('128', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 15:29:38');
INSERT INTO `sys_log` VALUES ('129', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 15:37:29');
INSERT INTO `sys_log` VALUES ('130', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 15:38:46');
INSERT INTO `sys_log` VALUES ('131', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 15:40:21');
INSERT INTO `sys_log` VALUES ('132', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 15:44:40');
INSERT INTO `sys_log` VALUES ('133', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 16:08:05');
INSERT INTO `sys_log` VALUES ('134', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 16:13:06');
INSERT INTO `sys_log` VALUES ('135', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 16:14:38');
INSERT INTO `sys_log` VALUES ('136', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 16:33:07');
INSERT INTO `sys_log` VALUES ('137', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 16:55:50');
INSERT INTO `sys_log` VALUES ('138', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 16:59:38');
INSERT INTO `sys_log` VALUES ('139', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 17:20:58');
INSERT INTO `sys_log` VALUES ('140', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 17:27:03');
INSERT INTO `sys_log` VALUES ('141', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 17:56:49');
INSERT INTO `sys_log` VALUES ('142', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 17:58:45');
INSERT INTO `sys_log` VALUES ('143', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 18:05:24');
INSERT INTO `sys_log` VALUES ('144', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 18:44:02');
INSERT INTO `sys_log` VALUES ('145', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 18:50:56');
INSERT INTO `sys_log` VALUES ('146', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 18:54:17');
INSERT INTO `sys_log` VALUES ('147', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 19:04:11');
INSERT INTO `sys_log` VALUES ('148', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 19:15:38');
INSERT INTO `sys_log` VALUES ('149', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 19:18:25');
INSERT INTO `sys_log` VALUES ('150', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 19:22:11');
INSERT INTO `sys_log` VALUES ('151', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 19:26:35');
INSERT INTO `sys_log` VALUES ('152', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 19:40:23');
INSERT INTO `sys_log` VALUES ('153', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-25 20:22:32');
INSERT INTO `sys_log` VALUES ('154', '975247111765495808', '超级管理员', '1', '127.0.0.1', null, null, '系统登录', '', '登录成功', '2018-05-26 10:03:05');

-- ----------------------------
-- Table structure for `sys_org`
-- ----------------------------
DROP TABLE IF EXISTS `sys_org`;
CREATE TABLE `sys_org` (
  `Id` varchar(50) NOT NULL,
  `Name` varchar(50) DEFAULT NULL,
  `OrgType` int(11) DEFAULT NULL,
  `ManagerId` varchar(50) DEFAULT NULL,
  `ParentId` varchar(50) DEFAULT NULL,
  `Description` varchar(100) DEFAULT NULL,
  `CreationTime` datetime NOT NULL,
  `IsDeleted` int(11) NOT NULL,
  `DeletionTime` datetime DEFAULT NULL,
  `DeleteUserId` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_org
-- ----------------------------
INSERT INTO `sys_org` VALUES ('975257850106482688', '币信', '1', null, null, null, '2018-05-16 14:29:00', '0', null, null);
INSERT INTO `sys_org` VALUES ('975258020424585216', '开发部', '2', null, '975257850106482688', null, '2018-05-16 14:30:00', '0', null, null);
INSERT INTO `sys_org` VALUES ('975258520352067584', '产品部', '2', null, '975257850106482688', null, '2018-05-16 14:32:00', '0', null, null);

-- ----------------------------
-- Table structure for `sys_orgpermission`
-- ----------------------------
DROP TABLE IF EXISTS `sys_orgpermission`;
CREATE TABLE `sys_orgpermission` (
  `Id` varchar(50) NOT NULL,
  `OrgId` varchar(50) NOT NULL,
  `PermissionId` varchar(50) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_orgpermission
-- ----------------------------

-- ----------------------------
-- Table structure for `sys_orgtype`
-- ----------------------------
DROP TABLE IF EXISTS `sys_orgtype`;
CREATE TABLE `sys_orgtype` (
  `Id` int(11) NOT NULL,
  `Name` varchar(100) DEFAULT NULL,
  `ParentId` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_orgtype
-- ----------------------------
INSERT INTO `sys_orgtype` VALUES ('1', '公司', null);
INSERT INTO `sys_orgtype` VALUES ('2', '部门', '1');

-- ----------------------------
-- Table structure for `sys_permission`
-- ----------------------------
DROP TABLE IF EXISTS `sys_permission`;
CREATE TABLE `sys_permission` (
  `Id` varchar(50) NOT NULL,
  `Name` varchar(50) DEFAULT NULL,
  `Code` varchar(50) DEFAULT NULL,
  `ParentId` varchar(50) DEFAULT NULL,
  `Type` int(11) NOT NULL,
  `Url` varchar(50) DEFAULT NULL,
  `Icon` varchar(50) DEFAULT NULL,
  `Description` varchar(50) DEFAULT NULL,
  `SortCode` int(11) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_permission
-- ----------------------------
INSERT INTO `sys_permission` VALUES ('1000195888754331648', '项目管理', null, null, '4', null, 'fa fa-file-picture-o ', null, '4');
INSERT INTO `sys_permission` VALUES ('1000198530486374400', '项目模块', null, '1000195888754331648', '4', null, null, null, '1');
INSERT INTO `sys_permission` VALUES ('1000198840529326080', '项目发起', null, '1000198530486374400', '1', null, null, null, '1');
INSERT INTO `sys_permission` VALUES ('1000200990722494464', '添加', null, '1000198530486374400', '3', null, null, null, '2');
INSERT INTO `sys_permission` VALUES ('1000201180971929600', '修改', null, '1000198530486374400', '3', null, null, null, '3');
INSERT INTO `sys_permission` VALUES ('975247111765495809', '系统管理', null, null, '4', null, 'fa fa-gears', null, '1');
INSERT INTO `sys_permission` VALUES ('975247111765495812', '组织管理', 'system.org', '976093351549997056', '1', '/System/Org/Index', null, null, '11');
INSERT INTO `sys_permission` VALUES ('975247111769690113', '用户管理', 'system.user', '976092881406267392', '1', '/System/User/Index', null, null, '3');
INSERT INTO `sys_permission` VALUES ('975247111769690114', '角色管理', 'system.role', '976094018385612800', '1', '/System/Role/Index', null, null, '22');
INSERT INTO `sys_permission` VALUES ('975247111769690117', '权限管理', 'system.permission', '976094340222947328', '1', '/System/Permission/Index', null, null, '28');
INSERT INTO `sys_permission` VALUES ('975247111769690119', '岗位管理', 'system.post', '976093808829796352', '1', '/System/Post/Index', null, null, '17');
INSERT INTO `sys_permission` VALUES ('975302681314856960', '添加权限', 'system.permission.add', '976094340222947328', '3', null, null, null, '29');
INSERT INTO `sys_permission` VALUES ('975302907190710272', '更新权限', 'system.permission.update', '976094340222947328', '3', null, null, null, '30');
INSERT INTO `sys_permission` VALUES ('975303061264273408', '删除权限', 'system.permission.delete', '976094340222947328', '3', null, null, null, '31');
INSERT INTO `sys_permission` VALUES ('975303702510440448', '添加用户', 'system.user.add', '976092881406267392', '3', null, null, null, '4');
INSERT INTO `sys_permission` VALUES ('975303794801905664', '修改用户', 'system.user.update', '976092881406267392', '3', null, null, null, '5');
INSERT INTO `sys_permission` VALUES ('975304146464935936', '重置用户密码', 'system.user.revise_password', '976092881406267392', '3', null, null, null, '6');
INSERT INTO `sys_permission` VALUES ('975304246784299008', '设置用户权限', 'system.user.set_permission', '976092881406267392', '3', null, null, null, '7');
INSERT INTO `sys_permission` VALUES ('975304393958232064', '修改用户状态', 'system.user.change_state', '976092881406267392', '3', null, null, null, '8');
INSERT INTO `sys_permission` VALUES ('975304681167392768', '切换用户部门权限状态', 'system.user.change_user_org_permission_state', '976092881406267392', '3', null, null, null, '9');
INSERT INTO `sys_permission` VALUES ('975305328730181632', '添加角色', 'system.role.add', '976094018385612800', '3', null, null, null, '23');
INSERT INTO `sys_permission` VALUES ('975305551389003776', '修改角色', 'system.role.update', '976094018385612800', '3', null, null, null, '24');
INSERT INTO `sys_permission` VALUES ('975305767588597760', '删除角色', 'system.role.delete', '976094018385612800', '3', null, null, null, '25');
INSERT INTO `sys_permission` VALUES ('975305860647620608', '设置权限', 'system.role.set_permission', '976094018385612800', '3', null, null, null, '26');
INSERT INTO `sys_permission` VALUES ('975306168010412032', '添加组织', 'system.org.add', '976093351549997056', '3', null, null, null, '12');
INSERT INTO `sys_permission` VALUES ('975306251208626176', '更新组织', 'system.org.update', '976093351549997056', '3', null, null, null, '13');
INSERT INTO `sys_permission` VALUES ('975306444578623488', '删除组织', 'system.org.delete', '976093351549997056', '3', null, null, null, '14');
INSERT INTO `sys_permission` VALUES ('975306615915941888', '设置权限', 'system.org.set_permission', '976093351549997056', '3', null, null, null, '15');
INSERT INTO `sys_permission` VALUES ('975306875098763264', '添加岗位', 'system.post.add', '976093808829796352', '3', null, null, null, '18');
INSERT INTO `sys_permission` VALUES ('975306937950408704', '更新岗位', 'system.post.update', '976093808829796352', '3', null, null, null, '19');
INSERT INTO `sys_permission` VALUES ('975307061833371648', '删除岗位', 'system.post.delete', '976093808829796352', '3', null, null, null, '20');
INSERT INTO `sys_permission` VALUES ('976092881406267392', '用户模块', null, '975247111765495809', '4', null, null, null, '2');
INSERT INTO `sys_permission` VALUES ('976093351549997056', '组织模块', null, '975247111765495809', '4', null, null, null, '10');
INSERT INTO `sys_permission` VALUES ('976093808829796352', '岗位模块', null, '975247111765495809', '4', null, null, null, '16');
INSERT INTO `sys_permission` VALUES ('976094018385612800', '角色模块', null, '975247111765495809', '4', null, null, null, '21');
INSERT INTO `sys_permission` VALUES ('976094340222947328', '权限模块', null, '975247111765495809', '4', null, null, null, '27');
INSERT INTO `sys_permission` VALUES ('976094603197419520', '文档模块', null, '975247111765495810', '4', null, null, null, '33');
INSERT INTO `sys_permission` VALUES ('976094863944716288', '菜单模块', null, '975247111765495810', '4', null, null, null, '38');
INSERT INTO `sys_permission` VALUES ('997039457519669248', '客户管理', null, null, '4', null, 'fa fa-users', null, '2');
INSERT INTO `sys_permission` VALUES ('997040019879366656', '客户禁用', 'client.customers', '998747605611712512', '1', '/Client/Customers/Index', null, null, '1');
INSERT INTO `sys_permission` VALUES ('998747605611712512', '启禁模块', null, '997039457519669248', '4', null, null, null, '1');
INSERT INTO `sys_permission` VALUES ('998748551368544256', '禁用/启用', 'Client.Customers.update', '998747605611712512', '3', null, null, null, '2');
INSERT INTO `sys_permission` VALUES ('998749192409190400', '审核模块', null, '997039457519669248', '4', null, null, null, '2');
INSERT INTO `sys_permission` VALUES ('998749737735819264', '客户审核', 'Client.audit', '998749192409190400', '1', '/Client/Audit/Index', null, null, '1');
INSERT INTO `sys_permission` VALUES ('999473375124525056', '审核', 'Client.Audit.update', '998749192409190400', '3', null, null, null, '2');
INSERT INTO `sys_permission` VALUES ('999476204467785728', '钱包管理', null, null, '4', null, 'fa fa-bitcoin', null, '3');
INSERT INTO `sys_permission` VALUES ('999476546450362368', '货币模块', null, '999476204467785728', '4', null, null, null, '1');
INSERT INTO `sys_permission` VALUES ('999476758115913728', '货币管理', 'Wallet.Currency', '999476546450362368', '1', '/Wallet/Currency/Index', null, null, '1');
INSERT INTO `sys_permission` VALUES ('999478920178962432', '货币添加', 'Wallet.Currency.add', '999476546450362368', '3', null, null, null, '2');
INSERT INTO `sys_permission` VALUES ('999479007017832448', '货币修改', 'Wallet.Currency.update', '999476546450362368', '3', null, null, null, '3');

-- ----------------------------
-- Table structure for `sys_post`
-- ----------------------------
DROP TABLE IF EXISTS `sys_post`;
CREATE TABLE `sys_post` (
  `Id` varchar(50) NOT NULL,
  `Name` varchar(50) DEFAULT NULL,
  `OrgId` varchar(50) DEFAULT NULL,
  `Description` varchar(200) DEFAULT NULL,
  `CreationTime` datetime NOT NULL,
  `IsDeleted` int(1) NOT NULL,
  `DeletionTime` datetime DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_post
-- ----------------------------
INSERT INTO `sys_post` VALUES ('975258691823603712', '.net工程师', '975258020424585216', null, '2018-05-16 14:32:00', '0', null);
INSERT INTO `sys_post` VALUES ('975258734123159552', '产品经理', '975258520352067584', null, '2018-05-16 14:32:00', '0', null);

-- ----------------------------
-- Table structure for `sys_role`
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `Id` varchar(50) NOT NULL,
  `Name` varchar(50) DEFAULT NULL,
  `SortCode` int(11) DEFAULT NULL,
  `IsEnabled` bit(1) DEFAULT NULL,
  `Description` varchar(500) DEFAULT NULL,
  `CreationTime` datetime DEFAULT NULL,
  `CreateUserId` varchar(50) DEFAULT NULL,
  `LastModifyTime` datetime DEFAULT NULL,
  `LastModifyUserId` varchar(50) DEFAULT NULL,
  `IsDeleted` int(1) DEFAULT NULL,
  `DeletionTime` datetime DEFAULT NULL,
  `DeleteUserId` varchar(500) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('975256202294464512', '群友', null, '', 'ewrqwrwr23123123234234324', '2018-05-16 14:22:00', '975247111765495808', null, null, '0', null, null);

-- ----------------------------
-- Table structure for `sys_rolepermission`
-- ----------------------------
DROP TABLE IF EXISTS `sys_rolepermission`;
CREATE TABLE `sys_rolepermission` (
  `Id` varchar(50) NOT NULL,
  `RoleId` varchar(50) NOT NULL,
  `PermissionId` varchar(50) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_rolepermission
-- ----------------------------
INSERT INTO `sys_rolepermission` VALUES ('975308019413946368', '975256202294464512', '975247111765495809');
INSERT INTO `sys_rolepermission` VALUES ('975308019413946369', '975256202294464512', '975247111769690113');
INSERT INTO `sys_rolepermission` VALUES ('975308019413946370', '975256202294464512', '975247111765495812');
INSERT INTO `sys_rolepermission` VALUES ('975308019413946371', '975256202294464512', '975247111769690119');
INSERT INTO `sys_rolepermission` VALUES ('975308019413946372', '975256202294464512', '975247111769690114');
INSERT INTO `sys_rolepermission` VALUES ('975308019413946373', '975256202294464512', '975247111769690117');
INSERT INTO `sys_rolepermission` VALUES ('975308019413946374', '975256202294464512', '975247111765495810');
INSERT INTO `sys_rolepermission` VALUES ('975308019413946375', '975256202294464512', '975247111765495817');
INSERT INTO `sys_rolepermission` VALUES ('975308019413946376', '975256202294464512', '975247111769690122');

-- ----------------------------
-- Table structure for `sys_user`
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `Id` varchar(50) NOT NULL,
  `AccountName` varchar(50) DEFAULT NULL,
  `Name` varchar(50) DEFAULT NULL,
  `HeadIcon` varchar(50) DEFAULT NULL,
  `Gender` int(11) DEFAULT NULL,
  `Birthday` datetime DEFAULT NULL,
  `MobilePhone` varchar(50) DEFAULT NULL,
  `Email` varchar(50) DEFAULT NULL,
  `WeChat` varchar(50) DEFAULT NULL,
  `Description` varchar(500) DEFAULT NULL,
  `CreationTime` datetime NOT NULL,
  `CreateUserId` varchar(50) DEFAULT NULL,
  `LastModifyTime` datetime DEFAULT NULL,
  `LastModifyUserId` varchar(50) DEFAULT NULL,
  `State` int(11) NOT NULL,
  `OrgIds` varchar(1024) DEFAULT NULL,
  `PostIds` varchar(1024) DEFAULT NULL,
  `RoleIds` varchar(1024) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('975247111765495808', 'admin', '超级管理员', null, '1', '2016-11-29 00:00:00', '1557985464', '123@163.com', 'so9527', '系统内置账户', '2016-07-20 00:00:00', null, null, null, '1', null, null, null);
INSERT INTO `sys_user` VALUES ('975308390198808576', 'hb', '海滨', null, null, '2018-04-21 00:00:00', null, null, null, null, '2018-05-16 17:50:00', '975247111765495808', null, null, '3', '975258520352067584', '', '975256202294464512');
INSERT INTO `sys_user` VALUES ('997045295579795456', 'a123456', '123', null, '1', null, '13693650515', null, null, null, '2018-05-17 17:24:57', '975247111765495808', null, null, '1', '975258020424585216', '975258691823603712', '975256202294464512');
INSERT INTO `sys_user` VALUES ('999869975504228352', 'hbdddddddddddd', '123', null, '2', '2018-05-25 00:00:00', '15869254569', '123@555.com', '123', '123', '2018-05-25 12:29:14', '975247111765495808', null, null, '1', '975257850106482688', '', '975256202294464512');

-- ----------------------------
-- Table structure for `sys_userlogon`
-- ----------------------------
DROP TABLE IF EXISTS `sys_userlogon`;
CREATE TABLE `sys_userlogon` (
  `Id` varchar(50) NOT NULL,
  `UserId` varchar(50) DEFAULT NULL,
  `UserPassword` varchar(50) DEFAULT NULL,
  `UserSecretkey` varchar(50) DEFAULT NULL,
  `PreviousVisitTime` datetime DEFAULT NULL,
  `LastVisitTime` datetime DEFAULT NULL,
  `LogOnCount` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_userlogon
-- ----------------------------
INSERT INTO `sys_userlogon` VALUES ('975247111765495808', '975247111765495808', '161222fd9eb63a5f47b9a70136999dc7', '04efe209a923cdb4', '2018-05-25 20:22:32', '2018-05-26 10:03:04', '512');
INSERT INTO `sys_userlogon` VALUES ('975308390207197184', '975308390198808576', '161222fd9eb63a5f47b9a70136999dc7', '04efe209a923cdb4', '2018-03-21 17:06:00', '2018-03-21 20:44:00', '91');
INSERT INTO `sys_userlogon` VALUES ('997045295588184064', '997045295579795456', '66d2d55a610e9d4b2eabc3e7a565f38a', '16f327f5a49f936e', null, null, '0');
INSERT INTO `sys_userlogon` VALUES ('999869975516811264', '999869975504228352', 'b75adc998c2bfa0e1f6fbb20b58ad183', 'd21297dc238a1fbd', null, null, '0');

-- ----------------------------
-- Table structure for `sys_userorg`
-- ----------------------------
DROP TABLE IF EXISTS `sys_userorg`;
CREATE TABLE `sys_userorg` (
  `Id` varchar(50) NOT NULL,
  `UserId` varchar(50) NOT NULL,
  `OrgId` varchar(50) NOT NULL,
  `DisablePermission` int(11) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_userorg
-- ----------------------------
INSERT INTO `sys_userorg` VALUES ('990423540840927232', '975308390198808576', '975258520352067584', '0');
INSERT INTO `sys_userorg` VALUES ('997045604045688832', '997045295579795456', '975258020424585216', '1');
INSERT INTO `sys_userorg` VALUES ('999869975525199872', '999869975504228352', '975257850106482688', '0');

-- ----------------------------
-- Table structure for `sys_userpermission`
-- ----------------------------
DROP TABLE IF EXISTS `sys_userpermission`;
CREATE TABLE `sys_userpermission` (
  `Id` varchar(50) NOT NULL,
  `UserId` varchar(50) NOT NULL,
  `PermissionId` varchar(50) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_userpermission
-- ----------------------------
INSERT INTO `sys_userpermission` VALUES ('997325873218719744', '975308390198808576', '975247111765495809');
INSERT INTO `sys_userpermission` VALUES ('997325873222914048', '975308390198808576', '976092881406267392');
INSERT INTO `sys_userpermission` VALUES ('997325873222914049', '975308390198808576', '975247111769690113');
INSERT INTO `sys_userpermission` VALUES ('997325873222914050', '975308390198808576', '975303702510440448');
INSERT INTO `sys_userpermission` VALUES ('997325873222914051', '975308390198808576', '975303794801905664');
INSERT INTO `sys_userpermission` VALUES ('997325873222914052', '975308390198808576', '975304146464935936');
INSERT INTO `sys_userpermission` VALUES ('997325873222914053', '975308390198808576', '975304246784299008');
INSERT INTO `sys_userpermission` VALUES ('997325873222914054', '975308390198808576', '975304393958232064');
INSERT INTO `sys_userpermission` VALUES ('997325873222914055', '975308390198808576', '975304681167392768');
INSERT INTO `sys_userpermission` VALUES ('997325873222914056', '975308390198808576', '976093351549997056');
INSERT INTO `sys_userpermission` VALUES ('997325873222914057', '975308390198808576', '975247111765495812');
INSERT INTO `sys_userpermission` VALUES ('997325873222914058', '975308390198808576', '975306168010412032');
INSERT INTO `sys_userpermission` VALUES ('997325873222914059', '975308390198808576', '975306251208626176');
INSERT INTO `sys_userpermission` VALUES ('997325873222914060', '975308390198808576', '975306444578623488');
INSERT INTO `sys_userpermission` VALUES ('997325873222914061', '975308390198808576', '975306615915941888');
INSERT INTO `sys_userpermission` VALUES ('997325873222914062', '975308390198808576', '976093808829796352');
INSERT INTO `sys_userpermission` VALUES ('997325873222914063', '975308390198808576', '975247111769690119');
INSERT INTO `sys_userpermission` VALUES ('997325873222914064', '975308390198808576', '975306875098763264');
INSERT INTO `sys_userpermission` VALUES ('997325873222914065', '975308390198808576', '975306937950408704');
INSERT INTO `sys_userpermission` VALUES ('997325873222914066', '975308390198808576', '975307061833371648');
INSERT INTO `sys_userpermission` VALUES ('997325873222914067', '975308390198808576', '976094018385612800');
INSERT INTO `sys_userpermission` VALUES ('997325873222914068', '975308390198808576', '975247111769690114');
INSERT INTO `sys_userpermission` VALUES ('997325873222914069', '975308390198808576', '975305328730181632');
INSERT INTO `sys_userpermission` VALUES ('997325873222914070', '975308390198808576', '975305551389003776');
INSERT INTO `sys_userpermission` VALUES ('997325873222914071', '975308390198808576', '975305767588597760');
INSERT INTO `sys_userpermission` VALUES ('997325873222914072', '975308390198808576', '975305860647620608');
INSERT INTO `sys_userpermission` VALUES ('997325873222914073', '975308390198808576', '976094340222947328');
INSERT INTO `sys_userpermission` VALUES ('997325873222914074', '975308390198808576', '975247111769690117');
INSERT INTO `sys_userpermission` VALUES ('997325873222914075', '975308390198808576', '975302681314856960');
INSERT INTO `sys_userpermission` VALUES ('997325873222914076', '975308390198808576', '975302907190710272');
INSERT INTO `sys_userpermission` VALUES ('997325873222914077', '975308390198808576', '975303061264273408');
INSERT INTO `sys_userpermission` VALUES ('997325873222914078', '975308390198808576', '997039457519669248');
INSERT INTO `sys_userpermission` VALUES ('997325873222914079', '975308390198808576', '997040019879366656');

-- ----------------------------
-- Table structure for `sys_userpost`
-- ----------------------------
DROP TABLE IF EXISTS `sys_userpost`;
CREATE TABLE `sys_userpost` (
  `Id` varchar(50) NOT NULL,
  `UserId` varchar(50) DEFAULT NULL,
  `PostId` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_userpost
-- ----------------------------
INSERT INTO `sys_userpost` VALUES ('975259300899459072', '975259300828155904', '975258691823603712');
INSERT INTO `sys_userpost` VALUES ('997045604045688833', '997045295579795456', '975258691823603712');

-- ----------------------------
-- Table structure for `sys_userrole`
-- ----------------------------
DROP TABLE IF EXISTS `sys_userrole`;
CREATE TABLE `sys_userrole` (
  `Id` varchar(50) NOT NULL,
  `UserId` varchar(50) NOT NULL,
  `RoleId` varchar(50) NOT NULL,
  PRIMARY KEY (`Id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sys_userrole
-- ----------------------------
INSERT INTO `sys_userrole` VALUES ('990423540828344320', '975308390198808576', '975256202294464512');
INSERT INTO `sys_userrole` VALUES ('997045603999551488', '997045295579795456', '975256202294464512');
INSERT INTO `sys_userrole` VALUES ('999869975521005568', '999869975504228352', '975256202294464512');

-- ----------------------------
-- Table structure for `two_factors`
-- ----------------------------
DROP TABLE IF EXISTS `two_factors`;
CREATE TABLE `two_factors` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '二次认证，短信，邮件',
  `user_id` bigint(20) DEFAULT NULL COMMENT '用户id',
  `otp_secret` text COMMENT '验证码信息',
  `activated` int(11) DEFAULT NULL COMMENT '是否已激活',
  `verify_type` int(11) DEFAULT NULL COMMENT '验证类型',
  `last_verify_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最晚激活时间',
  `refreshed_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '刷新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=58 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of two_factors
-- ----------------------------
INSERT INTO `two_factors` VALUES ('1', '1', '9760', '1', '1', '2018-05-10 15:54:29', '2018-05-10 15:50:37');
INSERT INTO `two_factors` VALUES ('2', '1', '169073', '2', '2', '2018-05-23 19:35:23', '2018-05-10 15:50:37');
INSERT INTO `two_factors` VALUES ('3', '1', 'EOXXVRGSZYDKCJGJ', '1', '3', '2018-05-23 10:52:22', '2018-05-10 15:50:37');
INSERT INTO `two_factors` VALUES ('4', '2', '8202', '2', '1', '2018-05-11 15:51:40', '2018-05-10 15:51:40');
INSERT INTO `two_factors` VALUES ('5', '2', '7452', '2', '2', '2018-05-10 16:21:40', '2018-05-10 15:51:40');
INSERT INTO `two_factors` VALUES ('6', '2', '3111', '2', '3', '2018-05-10 15:51:40', '2018-05-10 15:51:40');
INSERT INTO `two_factors` VALUES ('7', '3', '0933', '2', '1', '2018-05-11 15:59:41', '2018-05-10 15:59:41');
INSERT INTO `two_factors` VALUES ('8', '3', '7664', '2', '2', '2018-05-10 16:29:41', '2018-05-10 15:59:41');
INSERT INTO `two_factors` VALUES ('9', '3', '1417', '2', '3', '2018-05-10 15:59:41', '2018-05-10 15:59:42');
INSERT INTO `two_factors` VALUES ('10', '4', '4458', '2', '1', '2018-05-12 16:40:01', '2018-05-11 16:40:01');
INSERT INTO `two_factors` VALUES ('11', '4', '0021', '2', '2', '2018-05-11 17:10:01', '2018-05-11 16:40:01');
INSERT INTO `two_factors` VALUES ('12', '4', '8066', '2', '3', '2018-05-11 16:40:01', '2018-05-11 16:40:01');
INSERT INTO `two_factors` VALUES ('13', '5', '8313', '2', '1', '2018-05-17 10:28:52', '2018-05-16 10:28:52');
INSERT INTO `two_factors` VALUES ('14', '5', '2420', '2', '2', '2018-05-16 10:58:52', '2018-05-16 10:28:52');
INSERT INTO `two_factors` VALUES ('15', '5', '4862', '2', '3', '2018-05-16 10:28:52', '2018-05-16 10:28:52');
INSERT INTO `two_factors` VALUES ('16', '6', '4540', '2', '1', '2018-05-19 14:59:24', '2018-05-18 14:59:25');
INSERT INTO `two_factors` VALUES ('17', '6', '5255', '2', '2', '2018-05-18 15:29:25', '2018-05-18 14:59:25');
INSERT INTO `two_factors` VALUES ('18', '6', '9080', '2', '3', '2018-05-18 14:59:26', '2018-05-18 14:59:26');
INSERT INTO `two_factors` VALUES ('19', '7', '9889', '2', '1', '2018-05-19 15:01:12', '2018-05-18 15:01:13');
INSERT INTO `two_factors` VALUES ('20', '7', '5581', '2', '2', '2018-05-18 15:31:13', '2018-05-18 15:01:14');
INSERT INTO `two_factors` VALUES ('21', '7', '4545', '2', '3', '2018-05-18 15:01:15', '2018-05-18 15:01:15');
INSERT INTO `two_factors` VALUES ('22', '8', '1155', '2', '1', '2018-05-19 15:02:26', '2018-05-18 15:02:26');
INSERT INTO `two_factors` VALUES ('23', '8', '8662', '2', '2', '2018-05-18 15:32:27', '2018-05-18 15:02:28');
INSERT INTO `two_factors` VALUES ('24', '8', '2539', '2', '3', '2018-05-18 15:02:29', '2018-05-18 15:02:29');
INSERT INTO `two_factors` VALUES ('25', '9', '0328', '2', '1', '2018-05-19 15:05:04', '2018-05-18 15:05:04');
INSERT INTO `two_factors` VALUES ('26', '9', '9873', '2', '2', '2018-05-18 15:35:04', '2018-05-18 15:05:04');
INSERT INTO `two_factors` VALUES ('27', '9', '4026', '2', '3', '2018-05-18 15:05:04', '2018-05-18 15:05:04');
INSERT INTO `two_factors` VALUES ('28', '10', '6154', '2', '1', '2018-05-19 15:06:42', '2018-05-18 15:06:42');
INSERT INTO `two_factors` VALUES ('29', '10', '2064', '2', '2', '2018-05-18 15:36:42', '2018-05-18 15:06:43');
INSERT INTO `two_factors` VALUES ('30', '10', '0996', '2', '3', '2018-05-18 15:06:43', '2018-05-18 15:06:44');
INSERT INTO `two_factors` VALUES ('31', '11', '8451', '2', '1', '2018-05-19 15:08:26', '2018-05-18 15:08:26');
INSERT INTO `two_factors` VALUES ('32', '11', '4318', '2', '2', '2018-05-18 15:38:26', '2018-05-18 15:08:26');
INSERT INTO `two_factors` VALUES ('33', '11', '9543', '2', '3', '2018-05-18 15:08:27', '2018-05-18 15:08:27');
INSERT INTO `two_factors` VALUES ('34', '12', '9926', '1', '1', '2018-05-18 15:27:44', '2018-05-18 15:14:07');
INSERT INTO `two_factors` VALUES ('35', '12', '7185', '2', '2', '2018-05-18 15:44:07', '2018-05-18 15:14:08');
INSERT INTO `two_factors` VALUES ('36', '12', '5378', '2', '3', '2018-05-18 15:14:08', '2018-05-18 15:14:08');
INSERT INTO `two_factors` VALUES ('37', '13', '9130', '2', '1', '2018-05-19 15:28:21', '2018-05-18 15:28:21');
INSERT INTO `two_factors` VALUES ('38', '13', '1260', '2', '2', '2018-05-18 15:58:21', '2018-05-18 15:28:21');
INSERT INTO `two_factors` VALUES ('39', '13', '4409', '2', '3', '2018-05-18 15:28:21', '2018-05-18 15:28:21');
INSERT INTO `two_factors` VALUES ('40', '14', '3450', '1', '1', '2018-05-18 15:38:58', '2018-05-18 15:34:05');
INSERT INTO `two_factors` VALUES ('41', '14', '3417', '2', '2', '2018-05-18 16:04:05', '2018-05-18 15:34:05');
INSERT INTO `two_factors` VALUES ('42', '14', '2010', '2', '3', '2018-05-18 15:34:05', '2018-05-18 15:34:05');
INSERT INTO `two_factors` VALUES ('43', '15', '1399', '2', '1', '2018-05-19 15:45:29', '2018-05-18 15:41:50');
INSERT INTO `two_factors` VALUES ('44', '15', '6196', '2', '2', '2018-05-18 16:11:50', '2018-05-18 15:41:50');
INSERT INTO `two_factors` VALUES ('45', '15', 'W6EZROQAVXEUVFAM', '2', '3', '2018-05-21 17:40:21', '2018-05-18 15:41:50');
INSERT INTO `two_factors` VALUES ('46', '16', '765183', '2', '1', '2018-05-20 16:47:43', '2018-05-19 16:47:43');
INSERT INTO `two_factors` VALUES ('47', '16', '946262', '2', '2', '2018-05-19 17:17:43', '2018-05-19 16:47:43');
INSERT INTO `two_factors` VALUES ('48', '16', '775760', '2', '3', '2018-05-19 16:47:43', '2018-05-19 16:47:43');
INSERT INTO `two_factors` VALUES ('49', '17', '292915', '2', '1', '2018-05-20 17:03:29', '2018-05-19 17:03:29');
INSERT INTO `two_factors` VALUES ('50', '17', '186384', '2', '2', '2018-05-19 17:33:29', '2018-05-19 17:03:29');
INSERT INTO `two_factors` VALUES ('51', '17', '824348', '2', '3', '2018-05-19 17:03:29', '2018-05-19 17:03:29');
INSERT INTO `two_factors` VALUES ('52', '18', '697349', '2', '1', '2018-05-20 17:26:24', '2018-05-19 17:26:24');
INSERT INTO `two_factors` VALUES ('53', '18', '691879', '2', '2', '2018-05-19 17:56:24', '2018-05-19 17:26:24');
INSERT INTO `two_factors` VALUES ('54', '18', '242703', '2', '3', '2018-05-19 17:26:24', '2018-05-19 17:26:24');
INSERT INTO `two_factors` VALUES ('55', '19', '547173', '1', '1', '2018-05-21 10:32:56', '2018-05-21 10:31:27');
INSERT INTO `two_factors` VALUES ('56', '19', '783639', '2', '2', '2018-05-21 11:01:27', '2018-05-21 10:31:27');
INSERT INTO `two_factors` VALUES ('57', '19', '456867', '2', '3', '2018-05-21 10:31:27', '2018-05-21 10:31:27');

-- ----------------------------
-- Table structure for `users`
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `email` text COMMENT '用户邮箱',
  `password` text COMMENT '用户密码',
  `salt` text COMMENT '密码随机数',
  `display_name` text COMMENT '显示名称',
  `phone_number` text COMMENT '电话号码',
  `kind` int(11) DEFAULT NULL COMMENT '用户类型',
  `activated` int(11) DEFAULT NULL COMMENT '是否激活',
  `country_code` text COMMENT '国家代码',
  `identity_type` int(11) DEFAULT NULL COMMENT '证件类型',
  `identity_id` text COMMENT '身份证id',
  `real_name` text NOT NULL COMMENT '真实姓名',
  `disabled` int(11) DEFAULT NULL COMMENT '是否被禁',
  `photo_front` text COMMENT '证件正面照',
  `photo_back` text COMMENT '证件背面照',
  `photo_hand` text COMMENT '手持照片',
  `invite_code` text COMMENT '邀请码',
  `register_ip` text NOT NULL COMMENT '注册ip',
  `device_id` text COMMENT '注册设备id',
  `access_token` text COMMENT '用户登录token',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  `activated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '激活时间',
  `disabled_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '被禁或解禁时间',
  `description` text COMMENT '后台修改描述',
  `payment_password` text COMMENT '支付密码（默认6位，纯数字）',
  `google_qrimage` text COMMENT 'Google二维码图片',
  `google_is_bind` int(11) DEFAULT NULL COMMENT '是否绑定',
  `lock_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '锁定时间',
  `user_icon` text COMMENT '用户头像',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('1', 'blackeye@163.com', 'NGTH32xxYevH6ItxHSoNMFd6pjZlL2SUYpYEcASGGec=', '6640', '163_Account', '18500797779', '1', '1', '', null, '', '', '2', 'null', 'null', '0', '0UK4M8E3DA6CWY1U', '', '', '48221f33-fc0e-466f-b3c1-87e473f99a14', '2018-05-24 17:45:00', '2018-05-24 17:45:00', '2018-05-10 15:54:28', '2018-05-10 15:50:38', '213123123123', 'NGTH32xxYevH6ItxHSoNMFd6pjZlL2SUYpYEcASGGec=', null, '2', '2018-05-24 17:45:00', null);
INSERT INTO `users` VALUES ('2', 'blackeye@gmail.com', '3UPyQg6dFzhlgyE0lo3HTdubCH/lYuYFIMNoCdoe+Pw=', '1851', 'Blackeye_Account', '', '1', '2', '', null, '', '', '2', 'null', 'null', '0', '', '', '', null, '2018-05-23 16:31:48', '2018-05-23 16:31:48', '2018-05-10 15:51:41', '2018-05-10 15:51:41', null, null, '56715ec7-fd39-4138-8e71-ff5c2525b7fd.png', '1', '2018-05-23 16:31:48', null);
INSERT INTO `users` VALUES ('3', 'blackeye124@gmail.com', 'lewKEPxyMYGl2fG8tsZhmuInpYL9XKIgYbDw8W64lk0=', '7411', 'Test Gmail', '', '1', '2', '', null, '', '', '1', 'null', 'null', '0', '', '', '', null, '2018-05-23 15:49:52', '2018-05-10 15:59:43', '2018-05-10 15:59:43', '2018-05-10 15:59:43', 'tfdfsdf3213123123', null, null, '2', '2018-05-23 15:49:52', null);
INSERT INTO `users` VALUES ('4', 'visualstudio_test@126.com', 'rkskNIqIKGnps9G04pnRUsSQXeXnj850CJvyUxsKNDs=', '2112', 'VisualStudio Test', '', '1', '2', 'cn', null, '', '', '1', 'null', 'null', '0', '', '192.168.1.105', 'Win10 OS', null, '2018-05-23 15:49:51', '2018-05-11 16:39:59', '2018-05-11 16:39:59', '2018-05-11 16:39:59', '禁用时间TEST', null, null, '2', '2018-05-23 15:49:51', null);
INSERT INTO `users` VALUES ('5', 'test1@126.com', 'gjO9HfApuVGnDMCFFcYZJcK4IWw4S+fezGemCAe9ghs=', '5030', '', '', '1', '2', '', null, '', '', '2', 'null', 'null', '0', '', '221.221.139.58', 'postmain_id', '', '2018-05-23 15:49:50', '2018-05-16 10:28:53', '2018-05-16 10:28:53', '2018-05-16 10:28:53', '1231231231231', null, null, '2', '2018-05-23 15:49:50', null);
INSERT INTO `users` VALUES ('15', 'blackeye28@126.com', 'PX7o/yKKJfCh9D66q8MWGVyY8qsNL1euVnHVt1asiyw=', '9930', '', '', '1', '1', '', null, '', '', '2', 'null', 'null', '0', 'K8Z2NTZ1YS642CD7', '221.221.139.58', 'postmain_id', '', '2018-05-23 15:50:06', '2018-05-23 15:48:03', '2018-05-18 15:42:45', '2018-05-18 15:41:50', null, null, '', '2', '2018-05-23 15:50:06', null);
INSERT INTO `users` VALUES ('16', 'blackeye1@126.com', '6nnHqBR9zIdR9WX8BoBeoyA/ZgA4hLhwCy1fbqN+cXI=', '041780', '', '', '1', '2', '', null, '', '', '2', 'null', 'null', '0', '', '221.221.139.58', 'postmain_id', '', '2018-05-23 15:49:55', '2018-05-19 16:47:42', '2018-05-19 16:47:42', '2018-05-19 16:47:42', null, null, null, '2', '2018-05-23 15:49:55', null);
INSERT INTO `users` VALUES ('17', 'blackeye284@126.com', 'aNqgcymd+4UTiW1AyJx7ocwtgMhDbI8qEzRTXKQ/UXg=', '339525', '', '', '1', '2', '', null, '', '', '1', 'null', 'null', '0', '', '221.221.139.58', 'postmain_id', '', '2018-05-23 16:19:43', '2018-05-23 16:19:43', '2018-05-19 17:03:28', '2018-05-19 17:03:28', '231123123123', null, '99f444bf-2cf9-4205-81f6-9e77e65d883e.png', '1', '2018-05-23 16:19:43', null);
INSERT INTO `users` VALUES ('19', 'blackeye286@126.com', 'yG5B5k2LqjPVUFZSlORF/FZ+c7cHxO+kpB3UZiTlW+8=', '107140', '', '', '1', '1', '', null, '', '', '1', 'null', 'null', '0', '8ZURRYDYEFWQ3TZ6', '221.221.139.58', 'postmain_id', '', '2018-05-23 15:49:57', '2018-05-21 10:31:26', '2018-05-21 10:32:57', '2018-05-21 10:31:26', null, null, null, '2', '2018-05-23 15:49:57', null);

-- ----------------------------
-- Table structure for `user_balance`
-- ----------------------------
DROP TABLE IF EXISTS `user_balance`;
CREATE TABLE `user_balance` (
  `balance_id` varchar(32) NOT NULL COMMENT '余额主键',
  `user_id` varchar(32) NOT NULL,
  `currency_id` varchar(32) NOT NULL COMMENT '币表主键',
  `balance_value` bigint(50) unsigned DEFAULT NULL COMMENT '余额数字 1表示 0.00000001',
  `recharge_address` varchar(100) DEFAULT NULL,
  `QRCode_address` varchar(100) DEFAULT NULL COMMENT '二维码地址',
  `total_balance` bigint(20) DEFAULT NULL COMMENT '账户总额',
  `charge_un_account` bigint(20) DEFAULT NULL COMMENT '充值未到账',
  `withdraw_un_transfer` bigint(20) DEFAULT NULL COMMENT '提现未转出',
  `lock_position` bigint(20) DEFAULT NULL COMMENT '锁仓不可卖',
  `ico_undue` bigint(20) DEFAULT NULL COMMENT 'ico未到期',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`balance_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_balance
-- ----------------------------
INSERT INTO `user_balance` VALUES ('6af6c83ee32f705d3a80e5e8ab83ad7a', '10', '100003', '0', '0x06556998aea32f3eb1642222abd80525a6ed0bc1', '', null, null, null, null, null, '2018-05-25 15:13:54');
INSERT INTO `user_balance` VALUES ('7e9df0daf8bcceaf95c25690bd18cf8d', '10', '100001', '0', 'mfcczsfpnb9ZcEGTBbQxvJjQ3fCLvhbcax', '', null, null, null, null, null, '2018-05-25 15:13:54');
INSERT INTO `user_balance` VALUES ('fa42fee62e23cd2b2465ec57798ad937', '10', '100002', '0', 'mfiWC2rQu4XxpuhKjPMwtsypMMWwcPWiAg', '', null, null, null, null, null, '2018-05-25 15:13:54');
INSERT INTO `user_balance` VALUES ('gmbalance', 'gm', '100001', '212427175', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', null, null, null, null, null, null, '2018-05-25 15:13:54');
INSERT INTO `user_balance` VALUES ('gmbalance1', 'gm', '100002', '360000001', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', null, null, null, null, null, null, '2018-05-25 15:13:54');
INSERT INTO `user_balance` VALUES ('gmbalance2', 'gm', '100003', '99063700000000', '0x0707aa888bf2b72917b6ab09b7e4e2fea213df64', null, null, null, null, null, null, '2018-05-25 15:13:54');
INSERT INTO `user_balance` VALUES ('gmbalance3', 'gm', '100004', '375000000', 'mg6fv8GVwGt85qKj9ft1uJxtpwf2Qc2sqw', null, null, null, null, null, null, '2018-05-25 15:13:54');

-- ----------------------------
-- Table structure for `user_wallet`
-- ----------------------------
DROP TABLE IF EXISTS `user_wallet`;
CREATE TABLE `user_wallet` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL COMMENT '用户Id',
  `currency` text COMMENT '币种',
  `wallet_address` text CHARACTER SET utf8mb4 COMMENT '用户钱包地址',
  `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `is_deleted` int(11) DEFAULT NULL COMMENT '是否删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user_wallet
-- ----------------------------
INSERT INTO `user_wallet` VALUES ('1', '10', 'ETH', 'dfvghnjkljyuhgtrfd', '2018-05-22 18:43:25', '2018-05-22 18:43:25', '2');
INSERT INTO `user_wallet` VALUES ('2', '10', 'BTC', 'dfvghnjkljydfgdshgfhgfjhgjfgdguhgtrfd', '2018-05-22 19:00:38', '2018-05-22 19:00:38', '1');

-- ----------------------------
-- Table structure for `withdraw_cash`
-- ----------------------------
DROP TABLE IF EXISTS `withdraw_cash`;
CREATE TABLE `withdraw_cash` (
  `withdraw_id` varchar(64) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT ' 提现id',
  `user_id` varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '提现的用户id',
  `withdraw_comment` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '人工订单备注',
  `withdraw_money` bigint(19) NOT NULL COMMENT '提现金额',
  `withdraw_status` int(11) NOT NULL COMMENT '提现的状态 1 提出提现请求，2 取消提现请求 ，3 提现请求完成',
  `withdraw_date` datetime NOT NULL COMMENT '提现日期',
  `withdraw_actual_date` datetime DEFAULT NULL COMMENT '完成提现时间',
  `out_trade_no` varchar(32) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `user_send_ip` varchar(40) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `withdraw_commission` int(11) DEFAULT NULL COMMENT '手续费',
  `withdraw_address` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '提现地址',
  `withdraw_hash` varchar(500) COLLATE utf8mb4_unicode_ci DEFAULT NULL COMMENT '提现返回hash码，根据 txHash 可以区块链上查询转账进度',
  `withdraw_is_auto` int(1) DEFAULT '0' COMMENT '默认为0手动提现,1自动提现',
  `withdraw_coin` varchar(64) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `withdraw_type` int(11) DEFAULT '0' COMMENT '0 比特币 1以太坊系列',
  PRIMARY KEY (`withdraw_id`),
  KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- ----------------------------
-- Records of withdraw_cash
-- ----------------------------
INSERT INTO `withdraw_cash` VALUES ('1', '1', '1', '109982', '3', '2018-05-11 00:00:00', '2018-05-12 16:12:19', null, null, '15', 'mh8J9HR1K19PtqrWoqUpaZrX6qEsYtEXgT', 'e54ba40e18adcee1e9d7b304450a89f100076b029b6cb141ecdf4d25b0914ac2', '0', '100001', '0');
INSERT INTO `withdraw_cash` VALUES ('2', '1', '1', '100000000', '3', '2018-05-11 00:00:00', '2018-05-16 16:25:55', null, null, '15', '0xdd24A02e3A2A1424FBBd772629875EB46F4449F1', '0xa2548cd7fcde89edcaf442a32e1f6b67880d1d721431b19b7a29a1b08cf60da9', '0', '100002', '1');
INSERT INTO `withdraw_cash` VALUES ('3', '1', '1', '470000000000', '3', '2018-05-11 00:00:00', '2018-05-16 18:33:54', null, null, '0', '0xfb7d2cd03b61137d5A60276D1A0a7E67897e615E', '0xd7bdbab8aa72df5b919db703507f1b9573f884c29171b3a0a2f10c7349bb59c2', '0', '100003', '1');
