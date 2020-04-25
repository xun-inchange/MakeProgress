create table IF NOT EXISTS `blog_tag`(
    `id` int(10) unsigned not null auto_increment,
    `name` varchar(100) default '' comment '标签名称',
    `created_on` int(10) unsigned default '0' comment '创建人',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `state` tinyint(3) unsigned default '1' comment '状态 0为禁用,1为启用',
    primary key (`id`)
)engine=innoDB default charset=utf8 comment='文章标签管理';

create table IF NOT EXISTS `blog_article`(
    `id` int(10) unsigned not null auto_increment,
    `tag_id` int(10) unsigned default '0' comment '标签ID',
    `title` varchar(100) default '' comment '文章标题',
    `desc` varchar(255) default '' comment '简述',
    `content` text,
    `created_on` int(11) default null,
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(255) default '' comment '修改人',
    `state` tinyint(3) unsigned default '1' comment '状态 0为禁用1启用',
    primary key(`id`)
)engine=innoDB default charset=utf8 comment='文章管理';

CREATE TABLE IF not exists `blog_auth`(
    `id` int(10) unsigned not null auto_increment,
    `username` varchar(50) default '' comment '账号',
    `password` varchar(50) default '' comment '密码',
    primary key (`id`)
)engine=innoDB default charset=utf8;