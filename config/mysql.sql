/* *****************************************************************************
// Setup the preferences
// ****************************************************************************/
SET NAMES utf8 COLLATE 'utf8_unicode_ci';
SET foreign_key_checks = 1;
SET time_zone = '+00:00';
SET sql_mode = 'NO_AUTO_VALUE_ON_ZERO';
SET default_storage_engine = InnoDB;
SET CHARACTER SET utf8;
/* *****************************************************************************
// Remove old database
// ****************************************************************************/
DROP DATABASE IF EXISTS gocbs;
/* *****************************************************************************
// Create new database
// ****************************************************************************/
CREATE DATABASE gocbs DEFAULT CHARSET = utf8 COLLATE = utf8_unicode_ci;
USE gocbs;
/* *****************************************************************************
// Create the tables
// ****************************************************************************/
CREATE TABLE user_status (
  id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  status VARCHAR(25) NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,

  PRIMARY KEY (id)
);

INSERT INTO `user_status` (`id`, `status`, `created_at`, `updated_at`, `deleted`) VALUES
  (1, 'active',   CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0),
  (2, 'inactive', CURRENT_TIMESTAMP,  CURRENT_TIMESTAMP,  0);

CREATE TABLE positions (
  id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(25) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO `positions` (`id`, `name`) VALUES
  (1, 'Бухгалтер'),
  (2, 'Гл. Бухгалтер'),
  (3, 'Стажер-бухгалтер'),
  (4, 'Ген. Директор');

CREATE TABLE client_types (
  id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(25) NOT NULL,
  short_name VARCHAR(2) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO `client_types` (`id`, `name`, `short_name`) VALUES
  (1, 'Физическое лицо', 'ФЛ'),
  (2, 'Юридическое лицо', 'ЮЛ');

CREATE TABLE currencies (
  id TINYINT(1) UNSIGNED NOT NULL AUTO_INCREMENT,
  name VARCHAR(25) NOT NULL,
  code VARCHAR(25) NOT NULL,
  PRIMARY KEY (id)
);

INSERT INTO `currencies` (`id`, `name`, `code`) VALUES
  (1, 'TJS', '972'),
  (2, 'USD', '840'),
  (3, 'RUB', '810');

CREATE TABLE user (
  id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  first_name VARCHAR(50) NOT NULL,
  last_name VARCHAR(50) NOT NULL,
  email VARCHAR(100) NOT NULL,
  password CHAR(60) NOT NULL,
  phone_number varchar (20) NULL,
  address varchar (50) NULL,
  status_id TINYINT(1) UNSIGNED NOT NULL DEFAULT 1,
  position_id TINYINT(1) UNSIGNED NOT NULL DEFAULT 1,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,

  UNIQUE KEY (email),
  CONSTRAINT `f_user_status` FOREIGN KEY (`status_id`) REFERENCES `user_status` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `f_user_position` FOREIGN KEY (`position_id`) REFERENCES `positions` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY (id)
);


CREATE TABLE note (
  id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  content TEXT NOT NULL,
  user_id INT(10) UNSIGNED NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted TINYINT(1) UNSIGNED NOT NULL DEFAULT 0,

  CONSTRAINT `f_note_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  PRIMARY KEY (id)
);

create table accounts (
  id int(10) unsigned not null auto_increment primary key,
  number varchar(255) not null,
  name varchar(255) not null,
  type enum('active', 'passive') null,
  created_at timestamp null,
  updated_at timestamp null,

  constraint accounts_number_unique unique (number)
);

create table clients (
  id int(10) unsigned not null auto_increment primary key,
  name varchar(255) not null,
  full_name varchar(255) not null,
  client_type_id tinyint(1) UNSIGNED NOT NULL DEFAULT 1,
  email VARCHAR(100) NOT NULL,
  phone_number varchar (20) NULL,
  address varchar (50) NULL,
  created_at timestamp null,
  updated_at timestamp null,
  CONSTRAINT `f_type_client` FOREIGN KEY (client_type_id) REFERENCES client_types (id)
);

create table transactions (
  id int(10) unsigned not null auto_increment primary key,
  debit_account_id int(10) unsigned not null,
  credit_account_id int(10) unsigned not null,
  amount decimal(12,2) unsigned not null,
  description varchar(255) not null,
  clients varchar(255) null,
  created_at timestamp null,
  updated_at timestamp null,
  date date null,
  constraint transactions_debit_account_id_foreign foreign key (debit_account_id) references accounts (id),
  constraint transactions_credit_account_id_foreign foreign key (credit_account_id) references accounts (id)
);

create table bank_accounts (
  id int(10) unsigned not null auto_increment primary key,
  account_id int(10) unsigned not null,
  client_id int(10) unsigned not null,
  currency_id tinyint(1) unsigned not null,

  constraint f_client_bank_acc foreign key (client_id) references clients (id),
  constraint f_acc_bank_acc foreign key (account_id) references accounts (id),
  constraint f_curr_bank_acc foreign key (currency_id) references currencies (id)
);