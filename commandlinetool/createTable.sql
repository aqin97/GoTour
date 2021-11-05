create database `DB4GoTour`;
use  `DB4GoTour`;
create table `blog_tag` (
`id` int(10) unsigned NOT NULL auto_increment,
    `name` varchar(100) default '' comment '标签名称',
    `create_on` int(10) unsigned not null default '0' comment '创建时间',
    `create_by` varchar(100) default '' comment '创建人',
    `modified_on` int(10) unsigned default '0' comment '修改时间',
    `modified_by` varchar(100) default '' comment '修改人',
    `delete_on` int(10) unsigned default '0' comment '删除时间',
    `is_del` tinyint(3) unsigned default '0' comment '是否删除，0为未删除，1为已删除',
    `state` tinyint(3) unsigned default '1' comment '状态，0为禁用，1为启用',
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='标签管理';