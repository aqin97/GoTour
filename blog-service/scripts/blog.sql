create table common (
	`created_on` int(10) unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `delete_on` int(100) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除 0为未删除 1为删除'
);

use blog_service;

create table blog_tag (
	`id` int unsigned not null auto_increment,
    `name` varchar(100) default '' comment '标签名称',
    -- 公共字段
    `created_on` int unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int unsigned default '0' comment '删除时间',
    `is_del` tinyint unsigned default '0' comment '是否删除 0为未删除 1为删除',
    -- 公共字段
    `state` tinyint unsigned default '1' comment '状态 0为禁用 1为启用',
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment=' 标签管理';

create table blog_article (
	`id` int unsigned not null auto_increment,
    `title` varchar(100) default '' comment '文章名称',
    `desc` varchar(255) default '' comment '文章简述',
    `cover_image_url` varchar(255) default '' comment '封面图片地址',
    `content` longtext comment '文章内容',
    -- 公共字段
    `created_on` int unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int unsigned default '0' comment '删除时间',
    `is_del` tinyint unsigned default '0' comment '是否删除 0为未删除 1为删除',
    -- 公共字段
    `state` tinyint unsigned default '1' comment '状态 0为禁用 1为启用',
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment='文章管理';

drop table blog_service_tag;
drop table blog_tag;
drop table blog_article;

create table blog_article_tag (
	`id` int unsigned not null auto_increment,
    `article_id` int not null comment '文章ID',
    `tag_id` int unsigned not null default '0' comment '标签ID',
     -- 公共字段
     `created_on` int unsigned default '0' comment '创建时间',
    `created_by` varchar(100) default '' comment '创建人',
    `modified_on` int unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `deleted_on` int unsigned default '0' comment '删除时间',
    `is_del` tinyint unsigned default '0' comment '是否删除 0为未删除 1为删除',
    -- 公共字段
    primary key (`id`)
) engine=InnoDB default charset=utf8mb4 comment='文章标签关联';